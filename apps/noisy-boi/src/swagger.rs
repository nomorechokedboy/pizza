use crate::misc::HealthCheckResponse;
use utoipa::{
    openapi::{
        security::{HttpAuthScheme, HttpBuilder, SecurityScheme},
        Server,
    },
    Modify, OpenApi,
};

#[derive(OpenApi)]
#[openapi(
    info(description = "My Api description"), 
    paths(
        crate::misc::health_check,
        crate::notification::handlers::get_notifications,
        crate::notification::handlers::notify
    ),
    modifiers(&SecurityAddon),
    components(
        schemas(HealthCheckResponse,
                crate::notification::entities::user::User,
                crate::notification::entities::post::Post,
                crate::notification::entities::notification_object::NotificationObject,
                crate::notification::entities::notification::Notification,
                crate::notification::entities::notification_change::NotificationChange
        )
    )
)]
pub struct ApiDoc;

pub struct SecurityAddon;

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
