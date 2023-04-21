mod config;
mod notification;
use crate::config::AppSettings;
use actix_web::{
    get,
    middleware::Logger,
    web::{self, Redirect},
    App, HttpServer, Responder, Result, Scope,
};
use serde::Serialize;
use sqlx::{postgres::PgPoolOptions, Pool, Postgres};
use std::{
    sync::Arc,
    time::{Duration, Instant, SystemTime, UNIX_EPOCH},
};
use tracing::info;
use utoipa::{
    openapi::{
        security::{ApiKey, ApiKeyValue, SecurityScheme},
        Server,
    },
    Modify, OpenApi,
};
use utoipa_swagger_ui::SwaggerUi;

#[derive(OpenApi)]
#[openapi(
    info(description = "My Api description"), 
    paths(
        health_check
    ),
    modifiers(&SecurityAddon),
    components(
        schemas(HealthCheckResponse)
    )
)]
struct ApiDoc;

struct SecurityAddon;

impl Modify for SecurityAddon {
    fn modify(&self, openapi: &mut utoipa::openapi::OpenApi) {
        let components = openapi.components.as_mut().unwrap();
        openapi.servers = Some(vec![Server::new("/api/v1")]);
        components.add_security_scheme(
            "api_key",
            SecurityScheme::ApiKey(ApiKey::Header(ApiKeyValue::new("todo_apikey"))),
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
    /* security(
        ("api_key" = [])
    ), */
)]
#[get("/healthz")]
pub async fn health_check(app_state: web::Data<AppState>) -> Result<impl Responder> {
    let elapsed_time = app_state.start_time.elapsed();
    let uptime = Duration::from_secs_f64(elapsed_time.as_secs_f64()).as_secs();
    let now = SystemTime::now();
    let since_epoch = now.duration_since(UNIX_EPOCH).unwrap();
    let timestamp = since_epoch.as_millis();
    let res = HealthCheckResponse {
        message: String::from("Not dead yet"),
        uptime,
        timestamp,
    };
    Ok(web::Json(res))
}

pub struct AppState {
    start_time: Instant,
    pool: Arc<Pool<Postgres>>,
}

#[actix_web::main]
async fn main() -> std::io::Result<()> {
    std::env::set_var("RUST_LOG", "debug");
    tracing_subscriber::fmt().pretty().init();
    let settings = AppSettings::new()
        .map_err(|e| std::io::Error::new(std::io::ErrorKind::Other, e.to_string()))?;
    let server_url = settings.server_url();
    let db_url = settings.db_url();

    let pool = PgPoolOptions::new()
        .max_connections(5)
        .connect(&db_url)
        .await
        .map_err(|e| std::io::Error::new(std::io::ErrorKind::TimedOut, e.to_string()))?;
    sqlx::migrate!()
        .run(&pool)
        .await
        .map_err(|e| std::io::Error::new(std::io::ErrorKind::Other, e.to_string()))?;
    let pool = Arc::new(pool);

    let openapi = ApiDoc::openapi();
    let start_time = Instant::now();

    let server = HttpServer::new(move || {
        App::new()
            .app_data(web::Data::new(AppState {
                start_time,
                pool: pool.clone(),
            }))
            .wrap(Logger::default())
            .service(Scope::new("/api/v1").service(health_check))
            .service(index)
            .service(SwaggerUi::new("/docs/{_:.*}").url("/api-docs/openapi.json", openapi.clone()))
    });

    info!("Starting server at {server_url}");
    info!("Api documentation is at {server_url}/docs/");
    server.bind(server_url)?.run().await?;

    Ok(())
}
