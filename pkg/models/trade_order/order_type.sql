-- +goose Up
CREATE TYPE order_type AS ENUM (
    'LIMIT',
    'MARKET',
    'STOPLOSS',
    'STOPLIMIT'
);

-- +goose Down
DROP TYPE  order_type CASCADE;
