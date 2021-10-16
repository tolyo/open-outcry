-- price levels for instrument allow for fast access to volumes of orders
CREATE TABLE IF NOT EXISTS price_level(
    id              BIGSERIAL PRIMARY KEY,
    price           NUMERIC NOT NULL,
    side            order_side NOT NULL,
    instrument_id   BIGINT REFERENCES instrument(id),
    volume          NUMERIC NOT NULL DEFAULT 0
);
