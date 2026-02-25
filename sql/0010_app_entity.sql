
-- +goose Up
CREATE TABLE app_entity
(
    id                  BIGSERIAL PRIMARY KEY,
    pub_id              TEXT default uuid_generate_v4() UNIQUE NOT NULL,
    type                TEXT NOT NULL,   -- CLIENT, MASTER
    external_id         TEXT NOT NULL,
    updated_at          TIMESTAMPTZ default current_timestamp NOT NULL,
    created_at          TIMESTAMPTZ default current_timestamp NOT NULL
);

CREATE INDEX idx_app_entity_type ON app_entity(type);
CREATE INDEX idx_app_entity_external_id ON app_entity(external_id);

-- +goose Down
DROP TABLE app_entity;