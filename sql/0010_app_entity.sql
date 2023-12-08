
-- +goose Up
CREATE TABLE app_entity
(
    id                  BIGSERIAL PRIMARY KEY,
    pub_id              TEXT default uuid_generate_v4() UNIQUE NOT NULL,
    type                TEXT NOT NULL,   -- CLIENT, MASTER
    external_id         TEXT NOT NULL,
    updated_at          TIMESTAMP default current_timestamp NOT NULL,
    created_at          TIMESTAMP default current_timestamp NOT NULL
);

-- +goose Down
DROP TABLE app_entity;