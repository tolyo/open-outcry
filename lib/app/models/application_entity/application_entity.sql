CREATE TABLE IF NOT EXISTS application_entity
(
    id                  BIGSERIAL PRIMARY KEY,
    pub_id              TEXT DEFAULT uuid_generate_v4() UNIQUE NOT NULL,
    type                TEXT NOT NULL,   -- CLIENT, MASTER
    external_id         TEXT NOT NULL,
    updated_at          TIMESTAMP DEFAULT current_timestamp NOT NULL,
    created_at          TIMESTAMP DEFAULT current_timestamp NOT NULL
);