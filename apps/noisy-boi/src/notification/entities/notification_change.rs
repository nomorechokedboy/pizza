use super::user::User;
use sea_query::Iden;
use serde::{Deserialize, Serialize};
use utoipa::ToSchema;

#[derive(Debug, Iden)]
pub enum NotificationChangeIden {
    #[iden(rename = "notification_changes")]
    Table,
    Id,
    ActorId,
    NotificationObjectId,
}

#[derive(Debug, Deserialize, Serialize, ToSchema)]
#[serde(rename_all = "camelCase")]
pub struct NotificationChange {
    pub id: i64,
    pub actor: User,
}
