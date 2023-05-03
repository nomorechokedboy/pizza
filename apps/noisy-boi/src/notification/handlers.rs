use super::entities::notification_object::NotificationObject;
use crate::common::{app_state::AppState, auth_guard::AuthGuard, PaginationRes};
use actix_web::{get, web, Responder, Result};
use sea_query::Order;
use serde::Deserialize;
use utoipa::{IntoParams, ToSchema};

#[derive(Clone, Debug, Deserialize, ToSchema)]
pub enum Sort {
    Asc,
    Desc,
}

impl From<&Sort> for Order {
    fn from(s: &Sort) -> Self {
        match s {
            Sort::Desc => Self::Desc,
            Sort::Asc => Self::Asc,
        }
    }
}

#[derive(Clone, Debug, Deserialize, IntoParams)]
#[serde(rename_all = "camelCase")]
pub struct GetNotificationQuery {
    cursor: Option<u64>,
    page_size: Option<u64>,
    #[param(inline)]
    sort: Option<Sort>,
}

impl GetNotificationQuery {
    pub fn get_cursor(&self) -> u64 {
        self.cursor.unwrap_or_default() + self.get_page_size()
    }

    pub fn get_page_size(&self) -> u64 {
        self.page_size.unwrap_or(10)
    }

    pub fn get_sort(&self) -> Order {
        Order::from(self.sort.as_ref().unwrap_or(&Sort::Desc))
    }
}

#[utoipa::path(
    params(
        GetNotificationQuery
    ),
    responses(
        (status = 200, description = "OK", body = PaginationResNotificationObject),
        (status = 500, description = "Internal error", body = String)
    ),
    tag = "Notification",
    security(
        ("bearer" = [])
    )
)]
#[get("/notifications")]
pub async fn get_notifications(
    app_state: web::Data<AppState>,
    AuthGuard { user_id }: AuthGuard,
    query: web::Query<GetNotificationQuery>,
) -> Result<web::Json<PaginationRes<NotificationObject>>> {
    let repo = &app_state.get_notification_repo;
    let query = query.into_inner();
    let data = repo.exec(user_id, &query).await?;
    let res = PaginationRes::new(query.get_cursor(), data, query.get_page_size());
    Ok(web::Json(res))
}

#[utoipa::path(
    responses(
        (status = 200, description = "Notification event sent", body = NotificationObject),
    ),
    tag = "Notification",
    security(
        ("bearer" = [])
    )
)]
#[get("/notify")]
pub async fn notify(
    app_state: web::Data<AppState>,
    AuthGuard { user_id }: AuthGuard,
) -> impl Responder {
    let broadcaster = &app_state.sse_broadcaster;
    broadcaster.new_client(user_id).await
}
