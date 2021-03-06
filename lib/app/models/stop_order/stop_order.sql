CREATE TABLE IF NOT EXISTS stop_order(
    id                  BIGSERIAL PRIMARY KEY,
    pub_id              TEXT DEFAULT uuid_generate_v4() UNIQUE NOT NULL,
    price               DECIMAL NOT NULL,           -- trigger price
    trade_order_id      BIGINT REFERENCES trade_order(id) NOT NULL
);
