CREATE TABLE IF NOT EXISTS instrument(
    id              BIGSERIAL PRIMARY KEY,
    pub_id          TEXT DEFAULT uuid_generate_v4() UNIQUE NOT NULL,
    name            TEXT UNIQUE NOT NULL,
    precision       INTEGER NOT NULL,
    quote_currency  TEXT NOT NULL,
    -- currency instrument where dual exchange of currencies is involved
    currency_instrument BOOLEAN NOT NULL DEFAULT FALSE,
    base_currency   TEXT NOT NULL,
    active          BOOLEAN NOT NULL DEFAULT TRUE
);