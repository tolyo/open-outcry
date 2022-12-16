-- +goose Up
CREATE TYPE order_side AS ENUM (
    'BUY',
    'SELL'
    );

-- +goose Down
DROP TYPE  order_side CASCADE;
