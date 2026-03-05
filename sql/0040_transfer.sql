-- +goose Up

-- Journal entry: one record per transfer linking two ledger entries
CREATE TABLE transfer (
       id                                   BIGSERIAL PRIMARY KEY,
       pub_id                               TEXT default uuid_generate_v4() UNIQUE NOT NULL,
       type                                 transfer_type NOT NULL,
       amount                               NUMERIC default 0.00 NOT NULL CHECK (amount > 0),
       currency_name                        TEXT REFERENCES currency(name) NOT NULL,
       details                              TEXT NOT NULL,
       external_reference_number            TEXT NULL,
       status                               TEXT NOT NULL,
       updated_at                           TIMESTAMPTZ default current_timestamp NOT NULL,
       created_at                           TIMESTAMPTZ default current_timestamp NOT NULL
);

CREATE INDEX idx_transfer_currency_name ON transfer(currency_name);

-- Double-entry ledger: every transfer produces exactly two entries (DEBIT + CREDIT)
CREATE TABLE transfer_ledger_entry (
       id                                   BIGSERIAL PRIMARY KEY,
       pub_id                               TEXT default uuid_generate_v4() UNIQUE NOT NULL,
       transfer_id                          BIGINT REFERENCES transfer(id) ON DELETE CASCADE NOT NULL,
       currency_account_id                  BIGINT REFERENCES currency_account(id) NOT NULL,
       entry_type                           ledger_entry_type NOT NULL,
       amount                               NUMERIC default 0.00 NOT NULL CHECK (amount > 0),
       -- Resulting balance after this entry
       resulting_balance                    NUMERIC default 0.00 NOT NULL
                                            CHECK (resulting_balance >= 0),
       created_at                           TIMESTAMPTZ default current_timestamp NOT NULL
);

CREATE INDEX idx_tle_transfer_id ON transfer_ledger_entry(transfer_id);
CREATE INDEX idx_tle_currency_account_id ON transfer_ledger_entry(currency_account_id);

-- +goose Down
DROP TABLE transfer_ledger_entry CASCADE;
DROP TABLE transfer CASCADE;

