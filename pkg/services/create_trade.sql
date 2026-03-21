-- +goose Up

-- +goose StatementBegin
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
    trade_instance trade%ROWTYPE;
    master_app_entity_instance app_entity%ROWTYPE;
    seller_app_entity_instance app_entity%ROWTYPE;
    buyer_app_entity_instance app_entity%ROWTYPE;
BEGIN

    INSERT INTO trade (
        instrument_id,
        price,
        amount,
        seller_order_id,
        buyer_order_id,
        taker_order_id
    )
    VALUES (
        instrument_param.id,
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
    SELECT * FROM app_entity
    WHERE type = 'MASTER'
    INTO master_app_entity_instance;

    IF NOT FOUND THEN
        RAISE EXCEPTION 'master_app_entity_instance_not_found';
    END IF;

    -- set up seller
    SELECT * FROM app_entity
    INNER JOIN instrument_account
        ON instrument_account.app_entity_id = app_entity.id
    INNER JOIN trade_order
        ON trade_order.instrument_account_id = instrument_account.id
    WHERE trade_order.id = seller_trade_order_param.id
    INTO seller_app_entity_instance;

    IF NOT FOUND THEN
        RAISE EXCEPTION 'seller_app_entity_instance_not_found';
    END IF;

    -- set up buyer
    SELECT * FROM app_entity
    INNER JOIN instrument_account
        ON instrument_account.app_entity_id = app_entity.id
    INNER JOIN trade_order
        ON trade_order.instrument_account_id = instrument_account.id
    WHERE trade_order.id = buyer_trade_order_param.id
    INTO buyer_app_entity_instance;

    IF NOT FOUND THEN
        RAISE EXCEPTION 'buyer_app_entity_instance_not_found';
    END IF;

    -- EXECUTE TRANSFERS FOR SELLER IF INSTRUMENT IS FX
    IF instrument_param.fx_instrument IS TRUE THEN
        PERFORM process_transfer(
            'INSTRUMENT_SELL'::transfer_type,
            seller_app_entity_instance.pub_id,
            amount_param,
            instrument_param.base_currency,
            'MASTER',
            trade_instance.pub_id,
            trade_instance.pub_id,
            NULL
        );

        PERFORM process_transfer(
            'INSTRUMENT_BUY'::transfer_type,
            'MASTER',
            amount_param * price_param,
            instrument_param.quote_currency,
            seller_app_entity_instance.pub_id,
            trade_instance.pub_id,
            trade_instance.pub_id,
            (CASE seller_trade_order_param = taker_trade_order_param
                WHEN TRUE THEN 'TAKER_FEE'
                WHEN FALSE THEN 'MAKER_FEE'
            END)
        );
    ELSE
        -- transfer instruments directly between two accounts
        PERFORM create_instrument_account_transfer(
            (SELECT pub_id FROM instrument_account WHERE app_entity_id = seller_app_entity_instance.id),
            (SELECT pub_id FROM instrument_account WHERE app_entity_id = buyer_app_entity_instance.id),
            instrument_param,
            amount_param
        );

        -- TODO release any funds that are insufficient for buying an single instument
    END IF;



    -- EXECUTE TRANSFERS FOR BUYER
    PERFORM process_transfer(
        'INSTRUMENT_BUY'::transfer_type,
        buyer_app_entity_instance.pub_id,
        amount_param * price_param,
        instrument_param.quote_currency,
        'MASTER',
        trade_instance.pub_id,
        trade_instance.pub_id,
        NULL
    );

    PERFORM process_transfer(
        'INSTRUMENT_BUY'::transfer_type,
        'MASTER',
        amount_param,
        instrument_param.base_currency,
        buyer_app_entity_instance.pub_id,
        trade_instance.pub_id,
        trade_instance.pub_id,
        (CASE buyer_trade_order_param = taker_trade_order_param
             WHEN TRUE THEN 'TAKER_FEE'
             WHEN FALSE THEN 'MAKER_FEE'
        END)
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

    IF buyer_trade_order_param = taker_trade_order_param
    AND seller_trade_order_param.order_type = 'LIMIT'::order_type THEN
            PERFORM update_price_level(
                instrument_param.id,
                'SELL',
                price_param,
                amount_param,
                FALSE
            );
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

-- +goose StatementEnd

-- +goose Down
DROP FUNCTION create_trade(instrument, NUMERIC, NUMERIC, trade_order, trade_order, trade_order);