-- +goose Up

-- for non monetary instruments, instruments must be credited directly to trading account
CREATE TABLE trading_account_instrument (
    id                      BIGSERIAL PRIMARY KEY,
    pub_id                  TEXT default uuid_generate_v4() UNIQUE NOT NULL,
    trading_account         BIGINT REFERENCES trading_account(id) NOT NULL,
    amount                  BIGINT default 0 NOT NULL 
        CHECK (amount >= 0),
    amount_reserved         BIGINT default 0 NOT NULL
        CHECK (amount >= 0),
    
    instrument_id           BIGINT REFERENCES instrument(id) NOT NULL,
    updated_at              TIMESTAMP default current_timestamp NOT NULL,
    created_at              TIMESTAMP default current_timestamp NOT NULL,
    
    -- enforce one trading_account_instrument per trading_account
    UNIQUE (trading_account, instrument_id) 
);

-- +goose Down
DROP TABLE  trading_account_instrument CASCADE;
