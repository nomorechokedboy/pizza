use super::{notification::Notification, notification_change::NotificationChange, user::User};
use crate::common::FromDBFlatten;
use chrono::{DateTime, Utc};
use sea_query::Iden;
use serde::{Deserialize, Serialize};
use sqlx::{postgres::PgRow, FromRow, Row};
use utoipa::ToSchema;

#[derive(Debug, Iden)]
pub enum NotificationObjectDBIden {
    #[iden(rename = "notification_objects")]
    Table,
    Id,
    EntityId,
    CreatedAt,
    ActionType,
    EntityType,
}

#[derive(Debug, FromRow)]
pub struct NotificationObjectDB {
    pub id: i64,
    pub entity_id: i64,
    pub entity_data: String,
    pub action_type: String,
    pub created_at: DateTime<Utc>,
    pub notification: Notification,
    pub notification_change: NotificationChange,
}

#[derive(Debug, Deserialize, Serialize, ToSchema)]
#[serde(rename_all = "camelCase")]
pub struct NotificationObject {
    pub id: i64,
    pub entity_id: i64,
    pub entity_data: String,
    pub action_type: String,
    pub created_at: DateTime<Utc>,
    pub notifications: Vec<Notification>,
    pub notification_change: NotificationChange,
}

impl TryFrom<PgRow> for NotificationObjectDB {
    type Error = anyhow::Error;

    fn try_from(row: PgRow) -> Result<Self, Self::Error> {
        let id: i64 = row.try_get("notification_changes.id")?;
        let actor = User::from_row(&row, "actor")?;
        let notification_change = NotificationChange { id, actor };

        let id: i64 = row.try_get("notifications.id")?;
        let read_at: Option<DateTime<Utc>> = row.try_get("notifications.read_at")?;
        let notifier = User::from_row(&row, "notifier")?;
        let notification = Notification {
            id,
            notifier,
            read_at,
        };

        let entity_type: String = row.try_get("entity_type")?;
        let entity_data = match entity_type.as_str() {
            "comment" => {
                let data: String = row.try_get("comments.content")?;
                data
            }
            "post" => {
                let data: String = row.try_get("posts.title")?;
                data
            }
            _ => anyhow::bail!("Unknown entity"),
        };

        let id: i64 = row.try_get("id")?;
        let action_type: String = row.try_get("action_type")?;
        let entity_id: i64 = row.try_get("entity_id")?;
        let created_at: DateTime<Utc> = row.try_get("created_at")?;

        Ok(Self {
            id,
            action_type,
            entity_id,
            notification_change,
            notification,
            created_at,
            entity_data,
        })
    }
}

impl From<NotificationObjectDB> for NotificationObject {
    fn from(
        NotificationObjectDB {
            action_type,
            notification_change,
            notification,
            created_at,
            entity_id,
            entity_data,
            id,
        }: NotificationObjectDB,
    ) -> Self {
        Self {
            id,
            action_type,
            entity_id,
            notification_change,
            entity_data,
            notifications: vec![notification],
            created_at,
        }
    }
}

impl FromDBFlatten for NotificationObject {
    type DatabaseType = NotificationObjectDB;

    type GroupBy = i64;

    type VecType = Notification;

    fn group_by_field(db_item: &Self::DatabaseType) -> Self::GroupBy {
        db_item.id
    }

    fn vec_field(&mut self) -> &mut Vec<Self::VecType> {
        &mut self.notifications
    }
}
