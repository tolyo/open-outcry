-- +goose Up
-- for non monetary instruments, instruments must be credited directly to instrument account
CREATE TABLE instrument_account_holding (
    id                      BIGSERIAL PRIMARY KEY,
    pub_id                  TEXT default uuid_generate_v4() UNIQUE NOT NULL,
    instrument_account      BIGINT REFERENCES instrument_account(id) NOT NULL,
    amount                  BIGINT default 0 NOT NULL
        CHECK (amount >= 0),
    amount_reserved         BIGINT default 0 NOT NULL
        CHECK (amount >= 0),

    instrument_id           BIGINT REFERENCES instrument(id) NOT NULL,
    updated_at              TIMESTAMPTZ default current_timestamp NOT NULL,
    created_at              TIMESTAMPTZ default current_timestamp NOT NULL,

    -- enforce one instrument_account_holding per instrument_account
    UNIQUE (instrument_account, instrument_id)
);

CREATE INDEX idx_tai_instrument_id ON instrument_account_holding(instrument_id);

-- +goose Down
DROP TABLE instrument_account_holding CASCADE;

