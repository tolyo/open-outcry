CREATE TABLE IF NOT EXISTS instrument(
    id              BIGSERIAL PRIMARY KEY,
    pub_id          TEXT DEFAULT uuid_generate_v4() UNIQUE NOT NULL,
    name            TEXT UNIQUE NOT NULL,
    quote_currency  TEXT NOT NULL,
    -- fx instruments involve an exchange of currencies 
    fx_instrument BOOLEAN NOT NULL DEFAULT FALSE,
    base_currency   TEXT NULL,
    active          BOOLEAN NOT NULL DEFAULT TRUE
);