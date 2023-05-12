mod common;
mod config;
mod misc;
mod notification;
mod redis_subscriber;
mod swagger;

use crate::{
    common::app_state::AppState,
    config::AppSettings,
    misc::{health_check, index},
    notification::{
        handlers::{get_notifications, read_at},
        repository::GetNotificationRepository,
    },
    redis_subscriber::subscriber_task,
    swagger::ApiDoc,
};
use actix_cors::Cors;
use actix_web::{middleware::Logger, web, App, HttpServer, Scope};
use anyhow::Context;
use notification::{broadcaster::Broadcaster, handlers::notify};
use redis_async::client;
use sqlx::postgres::PgPoolOptions;
use std::sync::Arc;
use std::time::Instant;
use tracing::info;
use utoipa::OpenApi;
use utoipa_swagger_ui::SwaggerUi;

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
    let pubsub_con = client::pubsub_connect(&settings.redis.host, 6379)
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
        settings,
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
                    .service(notify)
                    .service(read_at),
            )
            .service(index)
            .service(SwaggerUi::new("/docs/{_:.*}").url("/api-docs/openapi.json", openapi.clone()))
    });

    info!("Starting server at {server_url}");
    info!("Api documentation is at {server_url}/docs/");
    server.bind(server_url)?.run().await?;

    Ok(())
}
