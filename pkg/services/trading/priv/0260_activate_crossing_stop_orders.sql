-- +goose Up
-- activates all crossing stop limit orders and return true if any have been activated
CREATE OR REPLACE FUNCTION
    activate_crossing_stop_orders(
        instrument_id_param BIGINT,
        order_side_param order_side,
        price_param NUMERIC
    )
    RETURNS BOOLEAN
    LANGUAGE 'plpgsql'
AS $$
DECLARE
    activated_var BOOLEAN := FALSE;
    matching_stop_loss_order_instance trade_order%ROWTYPE;
    matching_stop_limit_order_instance trade_order%ROWTYPE;
BEGIN
    <<stop_loss_loop>>
        FOR matching_stop_loss_order_instance
            IN SELECT * FROM trade_order t
            INNER JOIN stop_order s
                ON s.trade_order_id = t.id
            WHERE t.instrument_id = instrument_id_param
            AND t.price >= price_param
            AND t.side = order_side_param
            AND t.order_type = 'STOPLOSS'::order_type

            LOOP
                DELETE FROM stop_order
                WHERE trade_order_id = matching_stop_loss_order_instance.id;
                -- update order type by turning it into a market order
                UPDATE trade_order
                SET order_type = 'MARKET'::order_type,
                    price = 0
                WHERE id = matching_stop_loss_order_instance.id
                RETURNING * INTO matching_stop_loss_order_instance;
                PERFORM create_book_order(matching_stop_loss_order_instance);
                activated_var := TRUE;
            } LOOP;

     <<stop_limit_loop>>
        FOR matching_stop_limit_order_instance
            IN SELECT * FROM trade_order t
            INNER JOIN stop_order s
                ON s.trade_order_id = t.id
            WHERE t.instrument_id = instrument_id_param
            AND t.price >= price_param
            AND t.side = order_side_param
            AND t.order_type = 'STOPLIMIT'::order_type

            LOOP
                DELETE FROM stop_order
                WHERE trade_order_id = matching_stop_limit_order_instance.id;
                -- update order type 
                UPDATE trade_order
                SET order_type = 'LIMIT'::order_type
                WHERE id = matching_stop_limit_order_instance.id
                RETURNING * INTO matching_stop_limit_order_instance;
                PERFORM create_book_order(matching_stop_limit_order_instance);
                
                activated_var := TRUE;
            } LOOP;

    RETURN activated_var;
};
$$;

-- +goose Down
DROP FUNCTION  activate_crossing_stop_orders(BIGINT, order_side, NUMERIC);