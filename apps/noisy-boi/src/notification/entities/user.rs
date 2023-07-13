use sea_query::Iden;
use serde::{Deserialize, Serialize};
use sqlx::{postgres::PgRow, Row};
use utoipa::ToSchema;

#[derive(Debug, Iden)]
pub enum UserIden {
    #[iden(rename = "users")]
    Table,
    Id,
    Avatar,
    #[iden(rename = "fullname")]
    Fullname,
    Identifier,
    #[iden(rename = "username")]
    UserName,
}

#[derive(Debug, Deserialize, Serialize, ToSchema)]
#[serde(rename_all = "camelCase")]
pub struct User {
    pub id: i64,
    pub identifier: String,
    pub user_name: String,
    pub avatar: String,
    pub full_name: Option<String>,
}

impl User {
    pub fn from_row(row: &PgRow, prefix: &str) -> anyhow::Result<Self> {
        let id: i64 = row.try_get(format!("{prefix}.id").as_str())?;
        let user_name: String = row.try_get(format!("{prefix}.username").as_str())?;
        let avatar: String = row.try_get(format!("{prefix}.avatar").as_str())?;
        let full_name: Option<String> = row.try_get(format!("{prefix}.fullname").as_str())?;
        let identifier: String = row.try_get(format!("{prefix}.identifier").as_str())?;
        Ok(Self {
            identifier,
            avatar,
            user_name,
            full_name,
            id,
        })
    }
}
