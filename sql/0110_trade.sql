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
    updated_at      TIMESTAMPTZ default current_timestamp NOT NULL,
    created_at      TIMESTAMPTZ default current_timestamp NOT NULL
);

CREATE INDEX idx_trade_instrument_created ON trade(instrument_id, created_at DESC);
CREATE INDEX idx_trade_seller_order ON trade(seller_order_id);
CREATE INDEX idx_trade_buyer_order ON trade(buyer_order_id);
CREATE INDEX idx_trade_taker_order ON trade(taker_order_id);

-- +goose Down
DROP TABLE trade;
