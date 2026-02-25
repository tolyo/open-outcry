-- +goose Up

-- Journal entry: one record per transfer linking two ledger entries
CREATE TABLE trading_account_transfer (
       id                                 BIGSERIAL PRIMARY KEY,
       pub_id                             TEXT default uuid_generate_v4() UNIQUE NOT NULL,
       instrument_id                      BIGINT REFERENCES instrument(id) NOT NULL,
       amount                             INTEGER default 0 NOT NULL CHECK (amount > 0),
       details                            TEXT NULL,
       external_reference_number          TEXT NULL,
       updated_at                         TIMESTAMPTZ default current_timestamp NOT NULL,
       created_at                         TIMESTAMPTZ default current_timestamp NOT NULL
);

-- +goose StatementBegin
DO $$ BEGIN
    CREATE TYPE ledger_entry_type AS ENUM ('DEBIT', 'CREDIT');
EXCEPTION
    WHEN duplicate_object THEN null;
END $$;
-- +goose StatementEnd

-- Double-entry ledger: every transfer produces exactly two entries (DEBIT + CREDIT)
CREATE TABLE trading_account_ledger_entry (
       id                                 BIGSERIAL PRIMARY KEY,
       pub_id                             TEXT default uuid_generate_v4() UNIQUE NOT NULL,
       transfer_id                        BIGINT REFERENCES trading_account_transfer(id) NOT NULL,
       trading_account_instrument_id      BIGINT REFERENCES trading_account_instrument(id) NOT NULL,
       entry_type                         ledger_entry_type NOT NULL,
       amount                             INTEGER default 0 NOT NULL CHECK (amount > 0),
       -- Resulting balance after this entry
       resulting_balance                  INTEGER default 0 NOT NULL CHECK (resulting_balance >= 0),
       created_at                         TIMESTAMPTZ default current_timestamp NOT NULL
);

-- +goose Down
DROP TABLE trading_account_ledger_entry CASCADE;
DROP TABLE trading_account_transfer CASCADE;
DROP TYPE IF EXISTS ledger_entry_type;
