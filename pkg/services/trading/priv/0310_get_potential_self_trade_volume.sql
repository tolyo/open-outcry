-- +goose Up
CREATE OR REPLACE FUNCTION
    get_potential_self_trade_volume(
        instrument_id_param BIGINT, 
        side_param order_side, 
        trading_account_id_param BIGINT, 
        price_param NUMERIC
    )
    RETURNS NUMERIC
    LANGUAGE 'plpgsql'
AS $$
DECLARE
    acc NUMERIC := 0;
BEGIN
    IF side_param = 'SELL' THEN
        SELECT SUM(open_amount)
        FROM trade_order t
        INNER JOIN book_order b
            ON b.trade_order_id = t.id
        WHERE t.side = side_param
        AND t.trading_account_id = trading_account_id_param
        AND t.price <=  price_param 
        INTO acc;
    ELSE
        SELECT SUM(open_amount)
        FROM trade_order t
        INNER JOIN book_order b
            ON b.trade_order_id = t.id
        WHERE t.side = side_param
        AND t.trading_account_id = trading_account_id_param
        AND (t.price >=  price_param OR t.price = 0.00)
        INTO acc;        
    } IF;

    IF acc IS NULL THEN
        RETURN 0.00;
    } IF;

    RETURN acc;
};
$$;


-- +goose Down
DROP FUNCTION  get_potential_self_trade_volume(BIGINT, order_side, BIGINT, NUMERIC);