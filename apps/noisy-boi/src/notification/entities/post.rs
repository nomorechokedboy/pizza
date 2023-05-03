use sea_query::Iden;
use serde::Serialize;
use utoipa::ToSchema;

#[derive(Debug, Iden)]
pub enum PostIden {
    #[iden(rename = "posts")]
    Table,
    Id,
    Title,
}

#[derive(Debug, Serialize, ToSchema)]
#[serde(rename_all = "camelCase")]
pub struct Post {
    pub title: String,
    // pub slug: String,
}
