-- +goose Up
CREATE TABLE currency_account (
    id                      BIGSERIAL PRIMARY KEY,
    pub_id                  TEXT default uuid_generate_v4() UNIQUE NOT NULL,
    app_entity_id           BIGINT REFERENCES app_entity(id) NOT NULL,
    amount                  NUMERIC default 0.00 NOT NULL
                            CHECK (amount >= 0),
    amount_reserved         NUMERIC default 0.00 NOT NULL
                            CHECK (amount_reserved >= 0 and amount_reserved <= amount),
    currency_name           TEXT REFERENCES currency(name) NOT NULL,
    updated_at              TIMESTAMPTZ default current_timestamp NOT NULL,
    created_at              TIMESTAMPTZ default current_timestamp NOT NULL,

    -- enforce one currency account per application entity
    UNIQUE (app_entity_id, currency_name)
);

CREATE INDEX idx_currency_account_currency_name ON currency_account(currency_name);

-- +goose Down
DROP TABLE currency_account CASCADE;

