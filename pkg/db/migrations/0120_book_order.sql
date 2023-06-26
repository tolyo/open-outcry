-- +goose Up
CREATE TABLE book_order(
    id                  BIGSERIAL PRIMARY KEY,
    pub_id              TEXT default uuid_generate_v4() UNIQUE NOT NULL,
    trade_order_id      BIGINT REFERENCES trade_order(id) ON DELETE CASCADE NOT NULL
);

-- +goose Down
DROP TABLE book_order;