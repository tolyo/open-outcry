
-- +goose Up
CREATE TABLE trade (
    id              BIGSERIAL PRIMARY KEY,
    pub_id          TEXT default uuid_generate_v4() UNIQUE NOT NULL,
    instrument_id   BIGINT REFERENCES instrument(id) NOT NULL,
    price           DECIMAL NOT NULL,
    amount          DECIMAL NOT NULL,
    seller_order_id BIGINT REFERENCES trade_order(id) NOT NULL,
    buyer_order_id  BIGINT REFERENCES trade_order(id) NOT NULL,
    taker_order_id  BIGINT REFERENCES trade_order(id) NOT NULL, -- the order acting opposite of maker
    updated_at      TIMESTAMP default current_timestamp NOT NULL,
    created_at      TIMESTAMP default current_timestamp NOT NULL
);