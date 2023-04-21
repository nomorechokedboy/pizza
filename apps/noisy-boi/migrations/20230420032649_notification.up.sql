CREATE TABLE IF NOT EXISTS notifications (
  id SERIAL PRIMARY KEY,
  notification_object_id INTEGER NOT NULL REFERENCES notification_objects (id),
  notifier_id INTEGER NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  read_at TIMESTAMP
);

CREATE INDEX IF NOT EXISTS fk_notification_object_idx_1 ON notifications (notification_object_id ASC);
CREATE INDEX IF NOT EXISTS fk_notification_notifier_id_idx ON notifications (notifier_id ASC);