use crate::{
    common::PaginationRes,
    misc::HealthCheckResponse,
    notification::entities::{
        notification::Notification, notification_change::NotificationChange,
        notification_object::NotificationObject, post::Post, user::User,
    },
};
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
        crate::notification::handlers::notify,
        crate::notification::handlers::read_at
    ),
    modifiers(&SecurityAddon),
    components(
        schemas(HealthCheckResponse,
                PaginationRes<NotificationObject>,
                User,
                Post,
                NotificationObject,
                Notification,
                NotificationChange
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
