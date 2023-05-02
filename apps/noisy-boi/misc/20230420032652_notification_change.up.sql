CREATE TABLE IF NOT EXISTS notification_changes (
  id SERIAL PRIMARY KEY,
  notification_object_id INTEGER NOT NULL REFERENCES notification_objects (id),
  actor_id INTEGER NOT NULL
);

CREATE INDEX IF NOT EXISTS fk_notification_object_idx_2 ON notification_changes (notification_object_id ASC);
CREATE INDEX IF NOT EXISTS fk_notification_actor_id_idx ON notification_changes (actor_id ASC);
