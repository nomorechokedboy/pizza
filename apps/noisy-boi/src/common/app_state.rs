use crate::notification::{broadcaster::Broadcaster, repository::GetNotificationRepository};
use redis_async::client::PubsubConnection;
use std::{sync::Arc, time::Instant};

#[derive(Clone, Debug)]
pub struct AppState {
    pub start_time: Instant,
    pub get_notification_repo: GetNotificationRepository,
    pub sse_broadcaster: Arc<Broadcaster>,
    pub pubsub_con: PubsubConnection,
}
