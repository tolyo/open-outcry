-- +goose Up
-- Trading account belonging to application entity
CREATE TABLE trading_account (
    id                      BIGSERIAL PRIMARY KEY,
    pub_id                  TEXT default uuid_generate_v4() UNIQUE NOT NULL,
    app_entity_id           BIGINT REFERENCES app_entity(id) NOT NULL,
    updated_at              TIMESTAMPTZ default current_timestamp NOT NULL,
    created_at              TIMESTAMPTZ default current_timestamp NOT NULL
);

CREATE INDEX idx_trading_account_app_entity_id ON trading_account(app_entity_id);

-- +goose Down
DROP TABLE  trading_account CASCADE;
