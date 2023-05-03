use actix_web::{
    error::ErrorInternalServerError,
    get,
    web::{self, Redirect},
    Responder, Result,
};
use serde::Serialize;
use std::time::{Duration, SystemTime, UNIX_EPOCH};

use crate::common::app_state::AppState;

#[get("/")]
pub async fn index() -> impl Responder {
    Redirect::to("/docs/")
}

#[derive(utoipa::ToSchema, Serialize)]
pub struct HealthCheckResponse {
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
