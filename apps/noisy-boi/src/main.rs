mod common;
mod config;
mod notification;

use crate::{
    config::AppSettings,
    notification::{
        entities::notification_object::NotificationObject, handlers::get_notifications,
        repository::GetNotificationRepository,
    },
};
use actix_cors::Cors;
use actix_web::{
    error::{ErrorInternalServerError, ErrorUnauthorized},
    get, http,
    middleware::Logger,
    web::{self, Redirect},
    App, FromRequest, HttpServer, Responder, Result, Scope,
};
use anyhow::Context;
use futures::{future, StreamExt};
use jsonwebtoken::{decode, DecodingKey, Validation};
use notification::{broadcaster::Broadcaster, handlers::notify};
use redis_async::{
    client::{self, pubsub::PubsubStream, PubsubConnection},
    resp::FromResp,
};
use serde::{Deserialize, Serialize};
use sqlx::postgres::PgPoolOptions;
use std::sync::Arc;
use std::{
    future::{ready, Ready},
    time::{Duration, Instant, SystemTime, UNIX_EPOCH},
};
use tracing::{error, info};
use utoipa::{
    openapi::{
        security::{HttpAuthScheme, HttpBuilder, SecurityScheme},
        Server,
    },
    Modify, OpenApi,
};
use utoipa_swagger_ui::SwaggerUi;

#[derive(Debug)]
pub struct AuthGuard {
    pub user_id: i64,
}

#[derive(Debug, Deserialize, Serialize)]
pub struct TokenClaims {
    pub sub: i64,
    pub exp: usize,
}

impl FromRequest for AuthGuard {
    type Error = actix_web::Error;
    type Future = Ready<Result<Self, Self::Error>>;

    fn from_request(req: &actix_web::HttpRequest, _: &mut actix_web::dev::Payload) -> Self::Future {
        let res = match req.cookie("token").map(|c| c.value().to_string()) {
            Some(token) => Ok(token),
            None => req
                .headers()
                .get(http::header::AUTHORIZATION)
                .ok_or(ErrorUnauthorized("Missing token"))
                .and_then(|auth_value| {
                    auth_value
                        .to_str()
                        .map_err(|e| {
                            error!("{e}");
                            ErrorUnauthorized("Missing token")
                        })
                        .and_then(|bearer| {
                            bearer
                                .split(" ")
                                .last()
                                .ok_or(ErrorUnauthorized("Missing token"))
                                .map(|token| token.to_string())
                        })
                }),
        }
        .map(|token| {
            decode::<TokenClaims>(
                &token,
                &DecodingKey::from_secret("token-secret".as_ref()),
                &Validation::default(),
            )
            .map_err(|e| {
                error!("Decode error: {e}");
                ErrorUnauthorized("Invalid token")
            })
            .and_then(|decoded| Ok(decoded.claims))
        })
        .and_then(|claims| claims.map(|TokenClaims { sub: user_id, .. }| Self { user_id }));

        ready(res)
    }
}

#[derive(OpenApi)]
#[openapi(
    info(description = "My Api description"), 
    paths(
        health_check,
        notification::handlers::get_notifications,
        notification::handlers::notify
    ),
    modifiers(&SecurityAddon),
    components(
        schemas(HealthCheckResponse,
                notification::entities::user::User,
                notification::entities::post::Post,
                notification::entities::notification_object::NotificationObject,
                notification::entities::notification::Notification,
                notification::entities::notification_change::NotificationChange
        )
    )
)]
struct ApiDoc;

struct SecurityAddon;

impl Modify for SecurityAddon {
    fn modify(&self, openapi: &mut utoipa::openapi::OpenApi) {
        let components = openapi
            .components
            .as_mut()
            .expect("Should have register at least a component");
        openapi.servers = Some(vec![Server::new("/api/v1")]);
        components.add_security_scheme(
            "bearer",
            SecurityScheme::Http(
                HttpBuilder::new()
                    .scheme(HttpAuthScheme::Bearer)
                    .bearer_format("JWT")
                    .build(),
            ),
        )
    }
}

#[get("/")]
pub async fn index() -> impl Responder {
    Redirect::to("/docs/")
}

#[derive(utoipa::ToSchema, Serialize)]
struct HealthCheckResponse {
    message: String,
    uptime: u64,
    timestamp: u128,
}

#[utoipa::path(
    responses(
        (status = 200, description = "OK", body = HealthCheckResponse),
    ),
    tag = "HealthCheck",
)]
#[get("/healthz")]
pub async fn health_check(app_state: web::Data<AppState>) -> Result<impl Responder> {
    let elapsed_time = app_state.start_time.elapsed();
    let uptime = Duration::from_secs_f64(elapsed_time.as_secs_f64()).as_secs();
    let now = SystemTime::now();
    let since_epoch = now
        .duration_since(UNIX_EPOCH)
        .map_err(|e| ErrorInternalServerError(e))?;
    let timestamp = since_epoch.as_millis();
    let res = HealthCheckResponse {
        message: String::from("Not dead yet"),
        uptime,
        timestamp,
    };
    Ok(web::Json(res))
}

#[derive(Clone, Debug)]
pub struct AppState {
    pub start_time: Instant,
    pub get_notification_repo: GetNotificationRepository,
    pub sse_broadcaster: Arc<Broadcaster>,
    pub pubsub_con: PubsubConnection,
}

#[actix_web::main]
async fn main() -> anyhow::Result<()> {
    std::env::set_var("RUST_LOG", "debug");
    tracing_subscriber::fmt().init();
    let settings = AppSettings::new().context("Failed to parse settings")?;
    let server_url = settings.server_url();
    let db_url = settings.db_url();

    let pool = PgPoolOptions::new()
        .max_connections(100)
        .connect(&db_url)
        .await
        .context(format!("Failed to connect database, {db_url}"))?;
    /* sqlx::migrate!()
    .run(&pool)
    .await
    .map_err(|e| std::io::Error::new(std::io::ErrorKind::Other, e.to_string()))?; */

    let openapi = ApiDoc::openapi();
    let start_time = Instant::now();

    let get_notification_repo = GetNotificationRepository::new(pool);

    let broadcaster = Broadcaster::create();
    let pubsub_con = client::pubsub_connect(settings.redis.host, settings.redis.port)
        .await
        .context("Cannot connect to Redis")?;
    let msgs = pubsub_con
        .subscribe("notification")
        .await
        .context(format!("Cannot subscribe to topic"))?;
    let app_state = Arc::new(AppState {
        start_time,
        get_notification_repo,
        pubsub_con,
        sse_broadcaster: Arc::clone(&broadcaster),
    });

    tokio::spawn(subscriber_task(broadcaster, msgs));
    let server = HttpServer::new(move || {
        App::new()
            .app_data(web::Data::from(Arc::clone(&app_state)))
            .wrap(Logger::default())
            .wrap(Cors::permissive())
            .service(
                Scope::new("/api/v1")
                    .service(health_check)
                    .service(get_notifications)
                    .service(notify),
            )
            .service(index)
            .service(SwaggerUi::new("/docs/{_:.*}").url("/api-docs/openapi.json", openapi.clone()))
    });

    info!("Starting server at {server_url}");
    info!("Api documentation is at {server_url}/docs/");
    server.bind(server_url)?.run().await?;

    Ok(())
}

async fn subscriber_task(
    broadcaster: Arc<Broadcaster>,
    mut msgs: PubsubStream,
) -> anyhow::Result<()> {
    while let Some(message) = msgs.next().await {
        match message {
            Ok(message) => {
                let payload_serialize =
                    String::from_resp(message).context("Failed to get string from resp")?;
                let payload = serde_json::from_str::<NotificationObject>(&payload_serialize)
                    .context("Failed to deserialized payload");
                match payload {
                    Err(e) => {
                        error!("{e}");
                    }
                    Ok(payload) => {
                        info!("Payload: {payload:?}");
                        let futures = payload
                            .notifications
                            .iter()
                            .map(|n| broadcaster.broadcast(n.notifier.id, &payload_serialize));
                        let _ = future::join_all(futures).await;
                    }
                }
            }
            Err(e) => {
                error!("Subscriber err: {e}");
                break;
            }
        }
    }

    info!("Shut down subscriber task");
    Ok(())
}
