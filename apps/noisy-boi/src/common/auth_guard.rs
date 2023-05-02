use actix_web::{error::ErrorUnauthorized, http, FromRequest};
use jsonwebtoken::{decode, DecodingKey, Validation};
use serde::{Deserialize, Serialize};
use std::future::{ready, Ready};
use tracing::error;

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
