-- +goose Up

-- +goose StatementBegin

-- activates all crossing stop limit orders and return true if any have been activated
CREATE OR REPLACE FUNCTION
    process_crossing_stop_orders(
        instrument_id_param BIGINT,
        order_side_param order_side,
        price_param NUMERIC
    )
    RETURNS void
    LANGUAGE 'plpgsql'
AS $$
DECLARE
    matching_stop_loss_order_instance trade_order%ROWTYPE;
    matching_stop_limit_order_instance trade_order%ROWTYPE;
BEGIN

    <<stop_loss_loop>>
    FOR matching_stop_loss_order_instance
        IN SELECT * FROM trade_order t
        INNER JOIN stop_order s
            ON s.trade_order_id = t.id
        WHERE t.instrument_id = instrument_id_param
        AND s.price >= price_param
        AND t.side = order_side_param
        AND t.order_type = 'STOPLOSS'::order_type

        LOOP
            DELETE FROM stop_order
            WHERE trade_order_id = matching_stop_loss_order_instance.id;
            -- update order type
            UPDATE trade_order
            SET order_type = 'MARKET'::order_type,
                price = 0.00
            WHERE id = matching_stop_loss_order_instance.id;
            PERFORM process_trade_order(
                'VOID',
                (SELECT name FROM instrument WHERE id = instrument_id_param),
                'MARKET',
                order_side_param,
                0,
                matching_stop_loss_order_instance.amount,
                matching_stop_loss_order_instance.time_in_force::text,
                matching_stop_loss_order_instance.id
            );
        END LOOP;

    <<stop_limit_loop>>
    FOR matching_stop_limit_order_instance
        IN SELECT * FROM trade_order t
        INNER JOIN stop_order s
            ON s.trade_order_id = t.id
        WHERE t.instrument_id = instrument_id_param
        AND s.price >= price_param
        AND t.side = order_side_param
        AND t.order_type = 'STOPLIMIT'::order_type

        LOOP
            DELETE FROM stop_order
            WHERE trade_order_id = matching_stop_limit_order_instance.id;
            -- update order type
            UPDATE trade_order
            SET order_type = 'LIMIT'::order_type
            WHERE id = matching_stop_limit_order_instance.id;
            PERFORM process_trade_order(
                'VOID',
                (SELECT name FROM instrument WHERE id = instrument_id_param),
                'LIMIT',
                order_side_param,
                matching_stop_limit_order_instance.price,
                matching_stop_limit_order_instance.amount,
                matching_stop_limit_order_instance.time_in_force::text,
                matching_stop_limit_order_instance.id
            );
        END LOOP;

    RETURN;
END;
$$;
-- +goose StatementEnd

-- +goose Down
DROP FUNCTION process_crossing_stop_orders(BIGINT, order_side, NUMERIC);