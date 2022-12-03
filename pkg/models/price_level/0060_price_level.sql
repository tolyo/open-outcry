-- +goose Up
-- price levels for instrument allow for fast access to volumes of orders
CREATE TABLE price_level(
    id              BIGSERIAL PRIMARY KEY,
    price           NUMERIC NOT NULL,
    side            order_side NOT NULL,
    instrument_id   BIGINT REFERENCES instrument(id),
    volume          NUMERIC NOT NULL default 0
);

-- +goose Down
DROP TABLE price_level;
