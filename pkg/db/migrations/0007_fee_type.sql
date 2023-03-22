
-- +goose Up
CREATE TYPE fee_type AS ENUM (
    'DEPOSIT_FEE',
    'WITHDRAWAL_FEE',
    'TRANSFER_FEE',
    'TAKER_FEE',
    'MAKER_FEE'
);

-- +goose Down
DROP TYPE fee_type CASCADE;