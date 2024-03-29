
-- +goose Up
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TYPE payment_type AS ENUM (
    'DEPOSIT',
    'WITHDRAWAL',
    'TRANSFER',
    'INSTRUMENT_BUY',
    'INSTRUMENT_SELL',
    'CHARGE'
);

-- +goose Down
DROP TYPE payment_type CASCADE;