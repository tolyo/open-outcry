-- +goose Up
CREATE TABLE payment (
       id                                 BIGSERIAL PRIMARY KEY,
       pub_id                             TEXT default uuid_generate_v4() UNIQUE NOT NULL,
       type                               payment_type NOT NULL,
       amount                             NUMERIC default 0.00 NOT NULL CHECK (amount > 0),
       currency_name                      TEXT REFERENCES currency(name) NOT NULL,
       sender_payment_account_id          BIGINT REFERENCES payment_account(id) NOT NULL,
       beneficiary_payment_account_id     BIGINT REFERENCES payment_account(id) NOT NULL,
       details                            TEXT NOT NULL,
       external_reference_number          TEXT NULL,
       fee_sender                         NUMERIC default 0.00 NOT NULL CHECK (fee_sender >= 0),
       fee_beneficiary                    NUMERIC default 0.00 NOT NULL CHECK (fee_beneficiary >= 0),
       status                             TEXT NOT NULL,
       total_amount                       NUMERIC default 0.00 NOT NULL CHECK (total_amount >= 0),
       -- Resulting debit balance amount including reserved funds
       debit_balance_amount               NUMERIC default 0.00 NOT NULL CHECK (debit_balance_amount >= 0), 
       -- Resulting credit balance amount including reserved funds
       credit_balance_amount              NUMERIC default 0.00 NOT NULL CHECK (credit_balance_amount >= 0),
       updated_at                         TIMESTAMP default current_timestamp NOT NULL,
       created_at                         TIMESTAMP default current_timestamp NOT NULL
);

-- +goose Down
DROP TABLE payment CASCADE;
