use super::{
    entities::{
        comment::CommentIden,
        notification::{Notification, NotificationIden},
        notification_change::NotificationChangeIden,
        notification_object::{NotificationObject, NotificationObjectDB, NotificationObjectDBIden},
        post::PostIden,
        user::UserIden,
    },
    handlers::GetNotificationQuery,
};
use crate::common::from_db_flatten::FromDBFlatten;
use actix_web::error::{ErrorConflict, ErrorInternalServerError};
use sea_query::{Alias, Expr, JoinType, PostgresQueryBuilder, Query};
use sea_query_binder::{SqlxBinder, SqlxValues};
use sqlx::{Pool, Postgres};
use tracing::error;

#[derive(Clone, Debug)]
pub struct GetNotificationRepository {
    conn: Pool<Postgres>,
}

impl GetNotificationRepository {
    fn get_query(&self, user_id: i64, query: &GetNotificationQuery) -> (String, SqlxValues) {
        let actor_alias = Alias::new("actor");
        let notifier_alias = Alias::new("notifier");
        Query::select()
            .columns([
                (
                    NotificationObjectDBIden::Table,
                    NotificationObjectDBIden::Id,
                ),
                (
                    NotificationObjectDBIden::Table,
                    NotificationObjectDBIden::EntityId,
                ),
                (
                    NotificationObjectDBIden::Table,
                    NotificationObjectDBIden::CreatedAt,
                ),
                (
                    NotificationObjectDBIden::Table,
                    NotificationObjectDBIden::ActionType,
                ),
                (
                    NotificationObjectDBIden::Table,
                    NotificationObjectDBIden::EntityType,
                ),
            ])
            .expr_as(
                Expr::col((PostIden::Table, PostIden::Title)),
                Alias::new("posts.title"),
            )
            .expr_as(
                Expr::col((CommentIden::Table, CommentIden::Id)),
                Alias::new("comments.id"),
            )
            .expr_as(
                Expr::col((CommentIden::Table, CommentIden::Content)),
                Alias::new("comments.content"),
            )
            .expr_as(
                Expr::col((NotificationIden::Table, NotificationIden::Id)),
                Alias::new("notifications.id"),
            )
            .expr_as(
                Expr::col((NotificationIden::Table, NotificationIden::ReadAt)),
                Alias::new("notifications.read_at"),
            )
            .expr_as(
                Expr::col((NotificationChangeIden::Table, NotificationChangeIden::Id)),
                Alias::new("notification_changes.id"),
            )
            .expr_as(
                Expr::col((actor_alias.clone(), UserIden::Id)),
                Alias::new("actor.id"),
            )
            .expr_as(
                Expr::col((actor_alias.clone(), UserIden::Avatar)),
                Alias::new("actor.avatar"),
            )
            .expr_as(
                Expr::col((actor_alias.clone(), UserIden::Fullname)),
                Alias::new("actor.fullname"),
            )
            .expr_as(
                Expr::col((actor_alias.clone(), UserIden::UserName)),
                Alias::new("actor.username"),
            )
            .expr_as(
                Expr::col((actor_alias.clone(), UserIden::Identifier)),
                Alias::new("actor.identifier"),
            )
            .expr_as(
                Expr::col((notifier_alias.clone(), UserIden::Id)),
                Alias::new("notifier.id"),
            )
            .expr_as(
                Expr::col((notifier_alias.clone(), UserIden::Avatar)),
                Alias::new("notifier.avatar"),
            )
            .expr_as(
                Expr::col((notifier_alias.clone(), UserIden::Fullname)),
                Alias::new("notifier.fullname"),
            )
            .expr_as(
                Expr::col((notifier_alias.clone(), UserIden::UserName)),
                Alias::new("notifier.username"),
            )
            .expr_as(
                Expr::col((notifier_alias.clone(), UserIden::Identifier)),
                Alias::new("notifier.identifier"),
            )
            .from(NotificationObjectDBIden::Table)
            .inner_join(
                NotificationIden::Table,
                Expr::col((
                    NotificationObjectDBIden::Table,
                    NotificationObjectDBIden::Id,
                ))
                .equals((
                    NotificationIden::Table,
                    NotificationIden::NotificationObjectId,
                )),
            )
            .inner_join(
                NotificationChangeIden::Table,
                Expr::col((
                    NotificationObjectDBIden::Table,
                    NotificationObjectDBIden::Id,
                ))
                .equals((
                    NotificationChangeIden::Table,
                    NotificationChangeIden::NotificationObjectId,
                )),
            )
            .left_join(
                CommentIden::Table,
                Expr::col((CommentIden::Table, CommentIden::Id))
                    .equals((
                        NotificationObjectDBIden::Table,
                        NotificationObjectDBIden::EntityId,
                    ))
                    .and(
                        Expr::col((
                            NotificationObjectDBIden::Table,
                            NotificationObjectDBIden::EntityType,
                        ))
                        .eq("comment"),
                    ),
            )
            .left_join(
                PostIden::Table,
                Expr::col((PostIden::Table, PostIden::Id))
                    .equals((
                        NotificationObjectDBIden::Table,
                        NotificationObjectDBIden::EntityId,
                    ))
                    .and(
                        Expr::col((
                            NotificationObjectDBIden::Table,
                            NotificationObjectDBIden::EntityType,
                        ))
                        .eq("post"),
                    ),
            )
            .join_as(
                JoinType::InnerJoin,
                UserIden::Table,
                notifier_alias.clone(),
                Expr::col((NotificationIden::Table, NotificationIden::NotifierId))
                    .equals((notifier_alias.clone(), UserIden::Id)),
            )
            .join_as(
                JoinType::InnerJoin,
                UserIden::Table,
                actor_alias.clone(),
                Expr::col((
                    NotificationChangeIden::Table,
                    NotificationChangeIden::ActorId,
                ))
                .equals((actor_alias.clone(), UserIden::Id)),
            )
            .cond_where(
                Expr::col((NotificationIden::Table, NotificationIden::NotifierId))
                    .eq(user_id)
                    .and(
                        Expr::col((actor_alias, UserIden::Id))
                            .not_equals((notifier_alias, UserIden::Id)),
                    ),
            )
            .offset(query.get_offset())
            .limit(query.get_page_size())
            .order_by(
                (
                    NotificationObjectDBIden::Table,
                    NotificationObjectDBIden::Id,
                ),
                query.get_sort(),
            )
            .build_sqlx(PostgresQueryBuilder)
    }

    pub async fn exec(
        &self,
        user_id: i64,
        query: &GetNotificationQuery,
    ) -> Result<Vec<NotificationObject>, actix_web::Error> {
        let conn = &self.conn;
        let (sql, values) = self.get_query(user_id, query);
        let res = sqlx::query_with(&sql, values)
            .map(|row| NotificationObjectDB::try_from(row))
            .fetch_all(conn)
            .await
            .map_err(|e| {
                error!("Error getting notification_objects: {e}");
                ErrorInternalServerError(e)
            })?;
        let result_notifications_db: Result<Vec<NotificationObjectDB>, anyhow::Error> =
            res.into_iter().collect();

        result_notifications_db
            .map(FromDBFlatten::flatten)
            .map_err(|e| {
                error!("Error getting notification_objects: {e}");
                ErrorInternalServerError(e)
            })
    }

    pub fn new(conn: Pool<Postgres>) -> Self {
        Self { conn }
    }

    fn update_read_at_query(&self, notification_id: i64) -> (String, SqlxValues) {
        let (update_query, values) = Query::update()
            .table(NotificationIden::Table)
            .values([(NotificationIden::ReadAt, Expr::current_timestamp().into())])
            .and_where(
                Expr::col(NotificationIden::Id)
                    .eq(notification_id)
                    .and(Expr::col(NotificationIden::ReadAt).is_null()),
            )
            .returning(Query::returning().columns([
                NotificationIden::Id,
                NotificationIden::ReadAt,
                NotificationIden::NotifierId,
                NotificationIden::NotificationObjectId,
            ]))
            .build_sqlx(PostgresQueryBuilder);
        let sql = format!(
            r#"
                WITH updated_notification AS ({update_query})
                SELECT 
                users.id AS "users.id", users.avatar AS "users.avatar",
                users.identifier AS "users.identifier", users.username AS "users.username",
                users.fullname AS "users.fullname",
                updated_notification.id, updated_notification.read_at
                FROM updated_notification
                INNER JOIN users
                ON users.id = updated_notification.notifier_id;
            "#
        );
        (sql, values)
    }

    pub async fn exec_read_at(
        &self,
        notification_id: i64,
    ) -> Result<Notification, actix_web::Error> {
        let conn = &self.conn;
        let (sql, values) = self.update_read_at_query(notification_id);
        let row = sqlx::query_with(&sql, values)
            .fetch_optional(conn)
            .await
            .map_err(|e| {
                error!("Error getting notification: {e}");
                ErrorInternalServerError(e)
            })?;
        match row {
            Some(row) => Notification::try_from(row).map_err(|e| {
                error!("Error mapping notification: {e}");
                ErrorInternalServerError(e)
            }),
            None => Err(ErrorConflict("Notification already read")),
        }
    }
}
