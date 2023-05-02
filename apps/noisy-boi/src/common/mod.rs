pub mod app_state;
pub mod auth_guard;
pub mod from_db_flatten;

pub trait Pagination {
    fn pagination() -> () {
        unimplemented!();
    }
}
