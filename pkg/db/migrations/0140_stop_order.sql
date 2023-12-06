-- +goose Up
CREATE TABLE stop_order(
    id                  BIGSERIAL PRIMARY KEY,
    pub_id              TEXT default uuid_generate_v4() UNIQUE NOT NULL,
    price               DECIMAL NOT NULL,           -- trigger price
    trade_order_id      BIGINT REFERENCES trade_order(id) NOT NULL
);

-- +goose Down
DROP TABLE stop_order CASCADE;

