-- +goose Up
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TYPE transfer_type AS ENUM (
    'DEPOSIT',
    'WITHDRAWAL',
    'TRANSFER',
    'INSTRUMENT_BUY',
    'INSTRUMENT_SELL',
    'CHARGE'
);

-- +goose Down
DROP TYPE transfer_type CASCADE;

