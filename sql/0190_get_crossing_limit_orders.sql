-- +goose Up
-- +goose StatementBegin

CREATE OR REPLACE FUNCTION
    get_crossing_limit_orders(
        instrument_id_param BIGINT,
        side_param order_side,
        price_param NUMERIC,
        trade_account_id_param BIGINT
    )
    RETURNS setof book_order
    LANGUAGE 'plpgsql'
AS $$
BEGIN
    IF instrument_id_param IS NULL 
    OR side_param IS NULL
    OR price_param IS NULL
    OR trade_account_id_param IS NULL 
    THEN
        RAISE EXCEPTION 'param_cannot_be_null %, %, %, %',
            instrument_id_param,
            side_param,
            price_param,
            trade_account_id_param;
    END IF;
    
    IF price_param = 0 THEN
        -- handle market orders
        IF side_param = 'SELL' THEN
            RETURN QUERY SELECT b.id, b.pub_id, b.trade_order_id FROM book_order b
            INNER JOIN trade_order t
                ON b.trade_order_id = t.id
            WHERE t.instrument_id = instrument_id_param
                AND t.side = side_param
                AND t.trading_account_id != trade_account_id_param
                AND t.order_type = 'LIMIT'::order_type
            -- order first by price then by date created
            ORDER BY t.price, t.created_at;
        ELSE
            RETURN QUERY SELECT b.id, b.pub_id, b.trade_order_id FROM book_order b
            INNER JOIN trade_order t
                ON b.trade_order_id = t.id
            WHERE t.instrument_id = instrument_id_param
                AND t.side = side_param
                AND t.trading_account_id != trade_account_id_param
                AND t.order_type = 'LIMIT'::order_type
            -- order first by price then by date created
            ORDER BY t.price DESC, t.created_at;
        END IF;
    ELSE
        -- handle limit order
        IF side_param = 'SELL' THEN
            RETURN QUERY SELECT b.id, b.pub_id, b.trade_order_id FROM book_order b
            INNER JOIN trade_order t
                ON b.trade_order_id = t.id
            WHERE t.instrument_id = instrument_id_param
                AND t.side = side_param
                AND t.trading_account_id != trade_account_id_param
                AND t.order_type = 'LIMIT'::order_type
                AND t.price <= price_param
            -- order first by price then by date created
            ORDER BY t.price, t.created_at;
        ELSE
            RETURN QUERY SELECT b.id, b.pub_id, b.trade_order_id FROM book_order b
            INNER JOIN trade_order t
                ON b.trade_order_id = t.id
            WHERE t.instrument_id = instrument_id_param
                AND t.side = side_param
                AND t.trading_account_id != trade_account_id_param
                AND t.order_type = 'LIMIT'::order_type
                AND t.price >= price_param
            -- order first by price then by date created
            ORDER BY t.price DESC, t.created_at;
        END IF;
    END IF;
END;
$$;
-- +goose StatementEnd

-- +goose Down
DROP FUNCTION get_crossing_limit_orders(BIGINT, order_side, NUMERIC, BIGINT);