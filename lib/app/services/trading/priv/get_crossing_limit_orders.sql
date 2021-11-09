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
            ORDER BY t.order_type DESC, t.price ASC, t.created_at ASC;
        ELSE
            RETURN QUERY SELECT b.id, b.pub_id, b.trade_order_id FROM book_order b
            INNER JOIN trade_order t
                ON b.trade_order_id = t.id
            WHERE t.instrument_id = instrument_id_param
                AND t.side = side_param
                AND t.trading_account_id != trade_account_id_param
                AND t.order_type = 'LIMIT'::order_type
            -- order first by price then by date created
            ORDER BY t.order_type DESC, t.price DESC, t.created_at ASC;
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
            ORDER BY t.order_type DESC, t.price ASC, t.created_at ASC;
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
            ORDER BY t.order_type DESC, t.price DESC, t.created_at ASC;
        END IF;
    END IF;
END;
$$;