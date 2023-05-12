use super::{notification_object::NotificationObjectDB, user::User};
use chrono::{DateTime, Utc};
use sea_query::Iden;
use serde::{Deserialize, Serialize};
use sqlx::{postgres::PgRow, Row};
use utoipa::ToSchema;

#[derive(Debug, Iden)]
pub enum NotificationIden {
    #[iden(rename = "notifications")]
    Table,
    Id,
    NotifierId,
    NotificationObjectId,
    ReadAt,
}

#[derive(Debug, Deserialize, Serialize, ToSchema)]
#[serde(rename_all = "camelCase")]
pub struct Notification {
    pub id: i64,
    pub notifier: User,
    pub read_at: Option<DateTime<Utc>>,
}

impl From<NotificationObjectDB> for Notification {
    fn from(
        NotificationObjectDB {
            notification:
                Notification {
                    id,
                    read_at,
                    notifier,
                },
            ..
        }: NotificationObjectDB,
    ) -> Self {
        Self {
            id,
            read_at,
            notifier,
        }
    }
}

impl TryFrom<PgRow> for Notification {
    type Error = anyhow::Error;

    fn try_from(row: PgRow) -> Result<Self, Self::Error> {
        let id: i64 = row.try_get("id")?;
        let read_at: Option<DateTime<Utc>> = row.try_get("read_at")?;
        let notifier = User::from_row(&row, "users")?;

        Ok(Self {
            id,
            notifier,
            read_at,
        })
    }
}
