-- +goose Up
-- price levels for instrument allow for fast access to volumes of orders
CREATE TABLE price_level(
    id              BIGSERIAL PRIMARY KEY,
    price           NUMERIC NOT NULL,
    side            order_side NOT NULL,
    instrument_id   BIGINT REFERENCES instrument(id),
    volume          NUMERIC NOT NULL default 0
);

CREATE INDEX idx_price_level_instrument_id ON price_level(instrument_id);
CREATE INDEX idx_price_level_instrument_price_side ON price_level(instrument_id, price, side);

-- +goose Down
DROP TABLE price_level;
