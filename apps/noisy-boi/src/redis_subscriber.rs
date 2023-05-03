use crate::notification::{
    broadcaster::Broadcaster, entities::notification_object::NotificationObject,
};
use anyhow::Context;
use futures::{future, StreamExt};
use redis_async::{client::pubsub::PubsubStream, resp::FromResp};
use std::sync::Arc;
use tracing::{error, info};

pub async fn subscriber_task(
    broadcaster: Arc<Broadcaster>,
    mut msgs: PubsubStream,
) -> anyhow::Result<()> {
    while let Some(message) = msgs.next().await {
        match message {
            Ok(message) => {
                let payload_serialize =
                    String::from_resp(message).context("Failed to get string from resp")?;
                let payload = serde_json::from_str::<NotificationObject>(&payload_serialize)
                    .context("Failed to deserialized payload");
                match payload {
                    Err(e) => {
                        error!("{e}");
                    }
                    Ok(payload) => {
                        let futures = payload
                            .notifications
                            .iter()
                            .map(|n| broadcaster.broadcast(n.notifier.id, &payload_serialize));
                        let _ = future::join_all(futures).await;
                    }
                }
            }
            Err(e) => {
                error!("Subscriber err: {e}");
                break;
            }
        }
    }

    info!("Shut down subscriber task");
    Ok(())
}
