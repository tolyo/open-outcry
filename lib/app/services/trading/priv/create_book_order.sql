-- Add the order to the database while increasing the respective price level of
-- This function is meant to be used internally during trade order processing
CREATE OR REPLACE FUNCTION
    create_book_order(
        trade_order_id_param text
    )
    RETURNS TEXT
    LANGUAGE 'plpgsql'
AS $$
DECLARE
    trade_order_instance trade_order%ROWTYPE;
    instrument_instance instrument%ROWTYPE;
    price_level_instance price_level%ROWTYPE;
BEGIN

    SELECT * FROM trade_order
    WHERE pub_id = trade_order_id_param
    INTO trade_order_instance;

    IF NOT FOUND THEN
        RAISE EXCEPTION 'trading_accountorder_instance_not_found';
    END IF;
    
    -- create book_order
    INSERT INTO book_order (trade_order_id)
    VALUES (trade_order_instance.id);

    -- update price levels
    PERFORM update_price_level(
        trade_order_instance.instrument_id,
        trade_order_instance.side,
        trade_order_instance.price,
        trade_order_instance.open_amount,
        TRUE
    );
    
    RETURN trade_order_instance.pub_id;
END;
$$;