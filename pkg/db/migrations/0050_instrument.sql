-- +goose Up
CREATE TABLE instrument(
    id              BIGSERIAL PRIMARY KEY,
    pub_id          TEXT default uuid_generate_v4() UNIQUE NOT NULL,
    name            TEXT UNIQUE NOT NULL,
    quote_currency  TEXT NOT NULL,
    -- fx instruments involve an exchange of currencies 
    fx_instrument   BOOLEAN NOT NULL default FALSE,
    base_currency   TEXT NULL,
    active          BOOLEAN NOT NULL default TRUE
);

-- +goose Down
DROP TABLE instrument CASCADE;
