use actix_web::{
    error::{ErrorInternalServerError, ErrorUnauthorized},
    http, web, FromRequest,
};
use jsonwebtoken::{decode, DecodingKey, Validation};
use serde::{Deserialize, Serialize};
use std::future::{ready, Ready};
use tracing::error;

use super::app_state::AppState;

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
        let app_state = req.app_data::<web::Data<AppState>>();
        match app_state {
            None => ready(Err(ErrorInternalServerError("Internal error"))),
            Some(app_state) => {
                let settings = &app_state.settings;
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
                        &DecodingKey::from_secret(&settings.server.token_secret.as_ref()),
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
    }
}
