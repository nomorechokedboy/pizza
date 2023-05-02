use sea_query::Iden;

#[derive(Debug, Iden)]
pub enum CommentIden {
    #[iden(rename = "comments")]
    Table,
    Id,
    Content,
}
