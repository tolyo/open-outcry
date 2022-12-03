
-- Returns available market order volume in base currency for sell 
-- and quote currency for buy
CREATE OR REPLACE FUNCTION get_available_market_volume(
    instrument_id_param BIGINT,
    side_param order_side
)
    RETURNS NUMERIC
    LANGUAGE 'plpgsql'

AS $$
DECLARE
    acc NUMERIC;
BEGIN

    IF side_param = 'SELL' THEN
        SELECT SUM(volume)
        FROM price_level
        WHERE side = side_param
        AND instrument_id = instrument_id_param
        AND price = 0.00
        INTO acc;     
    ELSE
        SELECT SUM(volume)
        FROM price_level
        WHERE side = side_param
        AND instrument_id = instrument_id_param
        AND price = 0.00
        INTO acc;
    END IF;

    IF acc IS NULL THEN
        RETURN 0.00;
    END IF;

    RETURN acc;
END;
$$;