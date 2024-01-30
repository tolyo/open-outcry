-- +goose Up

CREATE TABLE trading_account_transfer (
       id                                 BIGSERIAL PRIMARY KEY,
       pub_id                             TEXT default uuid_generate_v4() UNIQUE NOT NULL,
       type                               TEXT NOT NULL,
       amount                             INTEGER default 0 NOT NULL CHECK (amount > 0),
       instrument_id                      BIGINT REFERENCES instrument(id) NOT NULL,
       sender_trading_account_id          BIGINT REFERENCES trading_account(id) NOT NULL,
       beneficiary_trading_account_id     BIGINT REFERENCES trading_account(id) NOT NULL,
       details                            TEXT NOT NULL,
       external_reference_number          TEXT NULL,
       -- Resulting debit balance amount 
       debit_instrument_amount            INTEGER default 0 NOT NULL CHECK (debit_instrument_amount >= 0), 
       -- Resulting credit instrument amount
       credit_instrument_amount           INTEGER default 0 NOT NULL CHECK (credit_instrument_amount >= 0),
       updated_at                         TIMESTAMP default current_timestamp NOT NULL,
       created_at                         TIMESTAMP default current_timestamp NOT NULL
);

-- +goose Down
DROP TABLE trading_account_transfer CASCADE;
