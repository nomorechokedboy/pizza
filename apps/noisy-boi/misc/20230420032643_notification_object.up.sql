CREATE TABLE IF NOT EXISTS notification_objects (
  id SERIAL PRIMARY KEY,
  entity_type_id SMALLINT NOT NULL,
  entity_id INTEGER NOT NULL
);
