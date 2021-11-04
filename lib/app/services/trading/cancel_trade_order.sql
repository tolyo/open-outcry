CREATE OR REPLACE FUNCTION
    cancel_trade_order(
        trade_order_id_param text
    )
    RETURNS void
    LANGUAGE 'plpgsql'
AS $$
DECLARE
    instrument_instance instrument%ROWTYPE;
    trade_order_instance trade_order%ROWTYPE; 
    book_order_instance book_order%ROWTYPE; 
    price_level_instance price_level%ROWTYPE;
    order_currency_var text;
BEGIN
    
    SELECT * FROM trade_order
    WHERE pub_id = trade_order_id_param
    INTO trade_order_instance;

    IF NOT FOUND THEN
        RAISE EXCEPTION 'trade_order_instance_not_found';
    END IF;

    SELECT * FROM instrument
    WHERE id = trade_order_instance.instrument_id
    INTO instrument_instance;

    SELECT * FROM book_order
    WHERE trade_order_id = trade_order_instance.id
    INTO book_order_instance;

    IF NOT FOUND THEN
        RAISE EXCEPTION 'book_order_instance_not_found';
    END IF;

    -- update price level if order is limit
    IF trade_order_instance.order_type = 'LIMIT'::order_type THEN

        PERFORM update_price_level(
            trade_order_instance.instrument_id,
            trade_order_instance.side,
            trade_order_instance.price,
            trade_order_instance.open_amount,
            FALSE
        );
    END IF;

    -- update trade order status
    IF trade_order_instance.amount != trade_order_instance.open_amount THEN
        UPDATE trade_order
        SET status = 'PARTIALLY_CANCELLED'
        WHERE id = trade_order_instance.id;
    ELSE
        UPDATE trade_order
        SET status = 'CANCELLED'
        WHERE id = trade_order_instance.id;
    END IF;

    -- release funds
    IF trade_order_instance.side = 'SELL'::order_side THEN
        order_currency_var = instrument_instance.base_currency;
    ELSE
        order_currency_var = instrument_instance.quote_currency;
    END IF;
    
    UPDATE payment_account
    SET amount_reserved = amount_reserved - trade_order_instance.open_amount
    WHERE currency_name = order_currency_var 
    AND application_entity_id = (
        SELECT application_entity_id FROM trading_account ta
        WHERE ta.id = trade_order_instance.trading_account_id
    );     

    -- delete book order
    DELETE FROM book_order WHERE id = book_order_instance.id; 
END;
$$;