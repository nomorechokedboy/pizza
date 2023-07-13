use crate::notification::entities::notification_object::NotificationObject;
use serde::Serialize;
use utoipa::ToSchema;

pub mod app_state;
pub mod auth_guard;
pub mod from_db_flatten;

pub trait Pagination {
    fn pagination() -> () {
        unimplemented!();
    }
}

#[derive(Clone, Debug, Serialize, ToSchema)]
#[aliases(PaginationResNotificationObject = PaginationRes<NotificationObject>)]
#[serde(rename_all = "camelCase")]
pub struct PaginationRes<T> {
    page: u64,
    data: Vec<T>,
    page_size: u64,
    // total: u64,
}

impl<T> PaginationRes<T> {
    pub fn new(page: u64, data: Vec<T>, page_size: u64) -> Self {
        Self {
            page,
            data,
            page_size,
        }
    }
}
