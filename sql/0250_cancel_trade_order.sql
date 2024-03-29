-- +goose Up
-- +goose StatementBegin
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
    order_currency_var text;
    update_amount_var NUMERIC;
BEGIN

    SELECT * FROM trade_order
    WHERE pub_id = trade_order_id_param
    -- allow cancellation only of active orders
    AND status IN ('OPEN'::trade_order_status, 'PARTIALLY_FILLED'::trade_order_status)
    INTO trade_order_instance;

    IF NOT FOUND THEN
        RAISE EXCEPTION 'trade_order_instance_not_found';
    END IF;

    SELECT * FROM instrument
    WHERE id = trade_order_instance.instrument_id
    INTO instrument_instance;

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
        update_amount_var = trade_order_instance.open_amount;
    ELSE
        order_currency_var = instrument_instance.quote_currency;
        update_amount_var = trade_order_instance.open_amount * trade_order_instance.price;
    END IF;

    UPDATE payment_account
    SET amount_reserved = amount_reserved - update_amount_var
    WHERE currency_name = order_currency_var
    AND app_entity_id = (
        SELECT app_entity_id FROM trading_account ta
        WHERE ta.id = trade_order_instance.trading_account_id
    );

    -- delete book order
    DELETE FROM book_order WHERE trade_order_id = trade_order_instance.id;
END;
$$;

-- +goose StatementEnd

-- +goose Down
DROP FUNCTION cancel_trade_order(TEXT);
