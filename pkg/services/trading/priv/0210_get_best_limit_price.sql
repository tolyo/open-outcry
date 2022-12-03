-- +goose Up
CREATE OR REPLACE FUNCTION get_best_limit_price(
    instrument_id_param BIGINT,
    side_param order_side
)
RETURNS NUMERIC
LANGUAGE 'plpgsql'

AS $$
DECLARE
    acc NUMERIC := 0;
BEGIN
    IF side_param = 'SELL'::order_side THEN
        SELECT price
        FROM price_level
        WHERE side = side_param
        AND instrument_id = instrument_id_param
        AND price > 0
        ORDER BY price ASC
        LIMIT 1
        INTO acc;
    ELSEIF side_param = 'BUY'::order_side THEN
        SELECT price
        FROM price_level
        WHERE side = side_param
        AND instrument_id = instrument_id_param
        AND price > 0
        ORDER BY price DESC
        LIMIT 1
        INTO acc;        
    ELSE
        RAISE EXCEPTION 'invalid side %', side_param;
    } IF;

    IF acc IS NULL THEN
        RETURN -1;
    } IF;

    RETURN acc;
};
$$;

-- +goose Down
DROP FUNCTION  get_best_limit_price(BIGINT, order_side);