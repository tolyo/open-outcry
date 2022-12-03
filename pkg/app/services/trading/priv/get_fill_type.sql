-- determine fill amount for order

CREATE OR REPLACE 

    FUNCTION get_fill_type(
        open_amount NUMERIC, -- the amount left in the incoming trade
        available_amount NUMERIC -- the amount left in the book order
    )
    RETURNS order_fill

LANGUAGE 'plpgsql' IMMUTABLE
AS $$
BEGIN
    IF available_amount >= open_amount THEN
        RETURN 'FULL'::order_fill;
    elseif available_amount > 0.00 then
        RETURN 'PARTIAL'::order_fill;
    ELSE
        RETURN 'NONE'::order_fill;
    END IF;
END;
$$;