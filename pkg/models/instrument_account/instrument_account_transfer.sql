-- +goose Up

-- Journal entry: one record per transfer linking two ledger entries
CREATE TABLE instrument_account_transfer (
       id                                 BIGSERIAL PRIMARY KEY,
       pub_id                             TEXT default uuid_generate_v4() UNIQUE NOT NULL,
       instrument_id                      BIGINT REFERENCES instrument(id) NOT NULL,
       amount                             INTEGER default 0 NOT NULL CHECK (amount > 0),
       details                            TEXT NULL,
       external_reference_number          TEXT NULL,
       updated_at                         TIMESTAMPTZ default current_timestamp NOT NULL,
       created_at                         TIMESTAMPTZ default current_timestamp NOT NULL
);


-- Double-entry ledger: every transfer produces exactly two entries (DEBIT + CREDIT)
CREATE TABLE instrument_account_ledger_entry (
       id                                 BIGSERIAL PRIMARY KEY,
       pub_id                             TEXT default uuid_generate_v4() UNIQUE NOT NULL,
       transfer_id                        BIGINT REFERENCES instrument_account_transfer(id) NOT NULL,
       instrument_account_holding_id      BIGINT REFERENCES instrument_account_holding(id) NOT NULL,
       entry_type                         ledger_entry_type NOT NULL,
       amount                             INTEGER default 0 NOT NULL CHECK (amount > 0),
       -- Resulting balance after this entry
       resulting_balance                  INTEGER default 0 NOT NULL CHECK (resulting_balance >= 0),
       created_at                         TIMESTAMPTZ default current_timestamp NOT NULL
);

CREATE INDEX idx_tat_instrument_id ON instrument_account_transfer(instrument_id);
CREATE INDEX idx_tale_transfer_id ON instrument_account_ledger_entry(transfer_id);
CREATE INDEX idx_tale_tai_id ON instrument_account_ledger_entry(instrument_account_holding_id);

-- +goose Down
DROP TABLE instrument_account_ledger_entry CASCADE;
DROP TABLE instrument_account_transfer CASCADE;

