-- create a trade and execute settlement between clients
-- open amounts for trade orders are updated
-- price levels are updated
CREATE OR REPLACE FUNCTION
    create_trade(
        instrument_param instrument,
        price_param DECIMAL,
        amount_param DECIMAL,
        seller_trade_order_param trade_order,
        buyer_trade_order_param trade_order,
        taker_trade_order_param trade_order
    )
    RETURNS void
    LANGUAGE 'plpgsql'
AS $$
DECLARE
    price_level_instance price_level%ROWTYPE;
    trade_instance trade%ROWTYPE;
    master_application_entity_instance application_entity%ROWTYPE;
    seller_application_entity_instance application_entity%ROWTYPE;
    buyer_application_entity_instance application_entity%ROWTYPE;
    update_price_level_var BOOLEAN := TRUE;
BEGIN

    INSERT INTO trade (
        instrument_id,
        price,
        amount,
        seller_order_id,
        buyer_order_id,
        taker_order_id
    )    
    VALUES (instrument_param.id,
            price_param,
            amount_param,
            seller_trade_order_param.id,
            buyer_trade_order_param.id,
            taker_trade_order_param.id
    )
    RETURNING * INTO trade_instance;

    IF NOT FOUND THEN
        RAISE EXCEPTION 'trade_instance_not_created';
    END IF;

    -- set up master
    SELECT * FROM application_entity
    WHERE type = 'MASTER'
    INTO master_application_entity_instance;

    IF NOT FOUND THEN
        RAISE EXCEPTION 'master_application_entity_instance_not_found';
    END IF;   

    -- set up buyer
    SELECT * FROM application_entity
    INNER JOIN trading_account
        ON trading_account.application_entity_id = application_entity.id
    INNER JOIN trade_order
        ON trade_order.trading_account_id = trading_account.id
    WHERE trade_order.id = seller_trade_order_param.id
    INTO seller_application_entity_instance;

    IF NOT FOUND THEN
        RAISE EXCEPTION 'seller_application_entity_instance_not_found';
    END IF;  

    -- set up buyer
    SELECT * FROM application_entity
    INNER JOIN trading_account
        ON trading_account.application_entity_id = application_entity.id
    INNER JOIN trade_order
        ON trade_order.trading_account_id = trading_account.id
    WHERE trade_order.id = buyer_trade_order_param.id
    INTO buyer_application_entity_instance;

    IF NOT FOUND THEN
        RAISE EXCEPTION 'buyer_application_entity_instance_not_found';
    END IF;       

    -- EXECUTE PAYMENTS FOR SELLER
    PERFORM create_payment(
        'INSTRUMENT_SELL'::payment_type,
        seller_application_entity_instance.pub_id,
        amount_param,
        instrument_param.base_currency,
        'MASTER',
        trade_instance.pub_id,
        trade_instance.pub_id
    );

    PERFORM create_payment(
        'INSTRUMENT_BUY'::payment_type,
        'MASTER',
        amount_param * price_param,
        instrument_param.quote_currency,
        seller_application_entity_instance.pub_id,
        trade_instance.pub_id,
        trade_instance.pub_id
    );

    -- EXECUTE PAYMENTS FOR BUYER
    PERFORM create_payment(
        'INSTRUMENT_BUY'::payment_type,
        buyer_application_entity_instance.pub_id,
        amount_param * price_param,
        instrument_param.quote_currency,
        'MASTER',
        trade_instance.pub_id,
        trade_instance.pub_id
    );

    PERFORM create_payment(
        'INSTRUMENT_BUY'::payment_type,
        'MASTER',
        amount_param,
        instrument_param.base_currency,
        buyer_application_entity_instance.pub_id,
        trade_instance.pub_id,
        trade_instance.pub_id
    );

    -- update open amounts
    UPDATE trade_order 
    SET open_amount = open_amount - amount_param,
        status = 
            CASE open_amount - amount_param = 0 
                WHEN TRUE THEN 'FILLED'::trade_order_status
                WHEN FALSE THEN 'PARTIALLY_FILLED'::trade_order_status
            END
    WHERE id = buyer_trade_order_param.id;

    UPDATE trade_order 
    SET open_amount = open_amount - amount_param,
        status = 
            CASE open_amount - amount_param = 0 
                WHEN TRUE THEN 'FILLED'::trade_order_status
                WHEN FALSE THEN 'PARTIALLY_FILLED'::trade_order_status
            END
    WHERE id = seller_trade_order_param.id;
    
    IF buyer_trade_order_param = taker_trade_order_param THEN
        
        IF seller_trade_order_param.order_type = 'LIMIT'::order_type THEN
            PERFORM update_price_level(
                instrument_param.id,
                'SELL',
                price_param,
                amount_param,
                FALSE
            );
        END IF;
    ELSE
        IF buyer_trade_order_param.order_type = 'LIMIT'::order_type THEN
            PERFORM update_price_level(
                instrument_param.id,
                'BUY',
                price_param,
                amount_param,
                FALSE
            );
        END IF;
    END IF;
END;
$$;