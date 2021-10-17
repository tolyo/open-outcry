-- Returns available limit order volume in base currency
CREATE OR REPLACE FUNCTION get_available_limit_volume(
    instrument_id_param BIGINT,
    side_param order_side,
    price_param DECIMAL
)
    RETURNS NUMERIC
    LANGUAGE 'plpgsql'

AS $$
DECLARE
    acc NUMERIC := 0;
BEGIN
    IF price_param = 0 THEN
        IF side_param = 'SELL' THEN
            SELECT SUM(volume)
            FROM price_level
            WHERE side = side_param
            AND instrument_id = instrument_id_param
            AND price != 0.00
            INTO acc;
        ELSE
            SELECT SUM(volume)
            FROM price_level
            WHERE side = side_param
            AND instrument_id = instrument_id_param
            AND price != 0.00
            INTO acc;        
        END IF;
    ELSE
        IF side_param = 'SELL' THEN
            SELECT SUM(volume)
            FROM price_level
            WHERE side = side_param
            AND instrument_id = instrument_id_param
            AND price <=  price_param 
            AND price != 0.00
            INTO acc;
        ELSE
            SELECT SUM(volume)
            FROM price_level
            WHERE side = side_param
            AND instrument_id = instrument_id_param
            AND price >=  price_param 
            INTO acc;        
        END IF;
    END IF;    

    IF acc IS NULL THEN
        RETURN 0.00;
    END IF;

    RETURN acc;
END;
$$;