
-- +goose Up
CREATE TYPE payment_type AS ENUM (
    'DEPOSIT',
    'WITHDRAWAL',
    'TRANSFER',
    'INSTRUMENT_BUY',
    'INSTRUMENT_SELL'
);

-- +goose Down
DROP TYPE   payment_type CASCADE;