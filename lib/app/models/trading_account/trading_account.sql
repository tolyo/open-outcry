-- Trading account belonging to application entity
CREATE TABLE IF NOT EXISTS trading_account (
    id                      BIGSERIAL PRIMARY KEY,
    pub_id                  TEXT DEFAULT uuid_generate_v4() UNIQUE NOT NULL,
    application_entity_id   BIGINT REFERENCES application_entity(id) NOT NULL,
    updated_at              TIMESTAMP DEFAULT current_timestamp NOT NULL,
    created_at              TIMESTAMP DEFAULT current_timestamp NOT NULL
);