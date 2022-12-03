
-- +goose Up
CREATE TYPE order_fill AS ENUM (
    'FULL',
    'PARTIAL',
    'NONE'
    );


-- +goose Down
DROP TYPE  order_fill CASCADE;

    