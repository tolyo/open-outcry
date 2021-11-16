-- Add the order to the database while increasing the respective price level of
-- This function is meant to be used internally during trade order processing
CREATE OR REPLACE FUNCTION
    create_book_order(
        trade_order_param trade_order
    )
    RETURNS TEXT
    LANGUAGE 'plpgsql' VOLATILE
AS $$
BEGIN
    -- create book_order
    INSERT INTO book_order (trade_order_id)
    VALUES (trade_order_param.id);

    -- update price levels
    PERFORM update_price_level(
        trade_order_param.instrument_id,
        trade_order_param.side,
        trade_order_param.price,
        trade_order_param.open_amount,
        TRUE
    );
    
    RETURN trade_order_param.pub_id;
END;
$$;