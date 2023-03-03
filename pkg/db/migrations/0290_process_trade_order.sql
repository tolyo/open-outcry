-- +goose Up
-- +goose StatementBegin

CREATE OR REPLACE FUNCTION
    process_trade_order(
        trading_account_id_param text,
        instrument_name_param text,
        order_type_param text,
        side_param order_side,
        price_param NUMERIC,
        amount_param NUMERIC,
        time_in_force_param text,
        trade_order_id_param BIGINT
    )
    RETURNS TEXT -- taker_trade_order.pub_id
    LANGUAGE 'plpgsql'
AS $$
DECLARE
    trading_account_instance trading_account%ROWTYPE;
    payment_account_instance payment_account%ROWTYPE;
    taker_trade_order_instance trade_order%ROWTYPE; -- saved
    taker_book_order_instance trade_order%ROWTYPE;  
    maker_book_order_instance trade_order%ROWTYPE;
    book_order_instance book_order%ROWTYPE;
    instrument_instance instrument%ROWTYPE;
    base_currency_precision INTEGER;
    quote_currency_precision INTEGER;
    opposite_side_var order_side;
    book_order_volume_var NUMERIC;
    available_limit_volume_var NUMERIC;
    available_market_volume_var NUMERIC;
    total_available_volume_var NUMERIC;
    trade_amount_var NUMERIC;
    fill_type_var order_fill;
    order_currency_var text;
    trade_price_var NUMERIC;
    trigger_loop_restart BOOLEAN := FALSE;
BEGIN

    IF trading_account_id_param != 'VOID' THEN
        -- trading account check
        SELECT * FROM trading_account
        WHERE pub_id = trading_account_id_param
        INTO trading_account_instance;

        IF NOT FOUND THEN
            RAISE EXCEPTION 'trading_account_instance_not_found';
        END IF;
    END IF;

    SELECT * FROM instrument
    WHERE name = instrument_name_param
    INTO instrument_instance;

    IF NOT FOUND THEN
        RAISE EXCEPTION 'instrument_instance_not_found';
    END IF;

    IF side_param = 'SELL' THEN
        opposite_side_var = 'BUY'::order_side;
        order_currency_var = instrument_instance.base_currency;
    ELSE
        opposite_side_var = 'SELL'::order_side;
        order_currency_var = instrument_instance.quote_currency;
    END IF;

    SELECT precision 
    FROM currency 
    WHERE name = instrument_instance.base_currency 
    INTO base_currency_precision;

    SELECT precision 
    FROM currency 
    WHERE name = instrument_instance.quote_currency 
    INTO quote_currency_precision;

    IF NOT FOUND THEN
        RAISE EXCEPTION 'precision_not_found';
    END IF;

    IF trading_account_id_param != 'VOID' THEN
        -- find trade order's payment account
        SELECT * FROM payment_account
        INNER JOIN app_entity
            ON payment_account.app_entity_id = app_entity.id
        WHERE app_entity.id = trading_account_instance.app_entity_id
        AND currency_name = order_currency_var
        INTO payment_account_instance;
        
        IF NOT FOUND THEN
            RAISE EXCEPTION 'payment_account_instance_not_found';
        END IF;

        amount_param = round(amount_param, base_currency_precision);
        price_param = round(price_param, quote_currency_precision);

        IF side_param = 'SELL' OR (side_param = 'BUY' and order_type_param = 'MARKET') THEN
            -- check sufficiency of funds
            IF payment_account_instance.amount - payment_account_instance.amount_reserved < amount_param THEN
                RAISE EXCEPTION 'insufficient_funds available: % required %', 
                    payment_account_instance.amount - payment_account_instance.amount_reserved,
                    amount_param;
            END IF;
            -- reserve required amount
            UPDATE payment_account
            SET amount_reserved =
                round(payment_account_instance.amount_reserved + amount_param, quote_currency_precision)
            WHERE id = payment_account_instance.id;
        ELSE
            -- check sufficiency of funds
            IF payment_account_instance.amount - payment_account_instance.amount_reserved < banker_round(amount_param * price_param, base_currency_precision)
            THEN
                RAISE EXCEPTION 'insufficient_funds available: % required: %', 
                        payment_account_instance.amount - payment_account_instance.amount_reserved,
                        amount_param * price_param;
            END IF;
            -- reserve required amount
            UPDATE payment_account
            SET amount_reserved =
                payment_account_instance.amount_reserved +
                    banker_round(amount_param * price_param, base_currency_precision)
            WHERE id = payment_account_instance.id;
        END IF;

        -- create trade order
        INSERT INTO trade_order (
            trading_account_id, 
            instrument_id, 
            order_type, 
            side, 
            price, 
            amount,
            open_amount,
            time_in_force
        )
        VALUES (
            trading_account_instance.id, 
            instrument_instance.id,
            order_type_param::order_type,
            side_param,
            price_param,
            amount_param,
            amount_param,
            time_in_force_param::order_time_in_force
        )
        RETURNING * INTO taker_trade_order_instance;
    
    ELSE

        SELECT * FROM trade_order
        WHERE id = trade_order_id_param
        INTO taker_trade_order_instance;

        SELECT * FROM trading_account
        WHERE id = taker_trade_order_instance.trading_account_id
        INTO trading_account_instance;

    END IF;

    IF order_type_param = 'STOPLOSS' OR order_type_param = 'STOPLIMIT' THEN
        INSERT INTO stop_order (
            trade_order_id,
            price
        ) 
        VALUES (
            taker_trade_order_instance.id,
            price_param
        );

        RETURN taker_trade_order_instance.pub_id;
    END IF;

    -- begin matching
    <<matching_loop>>
    LOOP
        -- check if order can be filled
        total_available_volume_var = 
            get_available_market_volume(instrument_instance.id, opposite_side_var) 
            + get_available_limit_volume(instrument_instance.id, opposite_side_var, price_param)
            - get_potential_self_trade_volume(instrument_instance.id, opposite_side_var, trading_account_instance.id, price_param);

        IF taker_trade_order_instance.time_in_force = 'FOK'::order_time_in_force 
        AND total_available_volume_var < amount_param THEN
            UPDATE trade_order
            SET status = 'REJECTED'::trade_order_status
            WHERE id = taker_trade_order_instance.id;
            -- release the funds
            -- reserve required amount
            IF side_param = 'SELL' OR
               (side_param = 'BUY' AND order_type_param = 'MARKET') THEN
                -- release reserved  amount
                UPDATE payment_account 
                SET amount_reserved = round(amount_reserved - amount_param, quote_currency_precision)
                WHERE id = payment_account_instance.id;
            ELSE
                -- release reserved amount
                UPDATE payment_account
                SET amount_reserved = amount_reserved - banker_round(amount_param * price_param, base_currency_precision)
                WHERE id = payment_account_instance.id;
            END IF;
            
            RETURN taker_trade_order_instance.pub_id;
        END IF;

        -- execute market order trades
        <<market_matching_loop>>
        FOR maker_book_order_instance 
            IN SELECT * FROM trade_order t
            INNER JOIN book_order b
                ON b.trade_order_id = t.id
            WHERE t.instrument_id = instrument_instance.id 
            AND t.trading_account_id != trading_account_instance.id
            AND t.side = opposite_side_var
            AND t.order_type = 'MARKET'::order_type
            ORDER BY t.created_at ASC
            
            LOOP
                IF amount_param > 0 THEN
                    trade_price_var = get_trade_price(
                            side_param::order_side,
                            order_type_param::order_type,
                            price_param,
                            opposite_side_var,
                            'MARKET'::order_type,
                            0,
                            instrument_instance.id 
                        );
                    
                    IF trade_price_var = 0 THEN
                        EXIT market_matching_loop;
                    END IF;

                    IF side_param = 'SELL' THEN
                        -- market buy in quote currency
                        trade_amount_var = 
                            banker_round(maker_book_order_instance.open_amount / trade_price_var, quote_currency_precision);
                        book_order_volume_var = 
                            maker_book_order_instance.open_amount - banker_round(trade_amount_var * trade_price_var, quote_currency_precision);

                        -- incoming order leftover
                        amount_param = amount_param - trade_amount_var;
                    ELSE
                        IF order_type_param = 'MARKET' THEN
                            -- market on market order handling
                            IF  maker_book_order_instance.open_amount < banker_round(amount_param / trade_price_var, quote_currency_precision) THEN
                                trade_amount_var = maker_book_order_instance.open_amount;
                            ELSE 
                                trade_amount_var = banker_round(amount_param / trade_price_var, quote_currency_precision);
                            END IF;
                            amount_param = amount_param - banker_round(trade_amount_var * trade_price_var, quote_currency_precision);
                        END IF;

                        IF order_type_param = 'LIMIT' THEN
                            IF maker_book_order_instance.open_amount < amount_param THEN
                                trade_amount_var = maker_book_order_instance.open_amount;
                                amount_param = amount_param - trade_amount_var;
                            ELSE 
                                trade_amount_var = amount_param;
                                amount_param = amount_param - trade_amount_var;
                            END IF;
                        END IF;

                        book_order_volume_var = maker_book_order_instance.open_amount - trade_amount_var;
                    END IF;

                    -- book order is greater than incoming order. book order must remain in book
                    IF book_order_volume_var = 0 THEN
                        -- delete trade
                        DELETE FROM book_order
                        WHERE trade_order_id = maker_book_order_instance.id;
                    END IF;

                    -- update open amount
                    IF side_param = 'SELL' THEN
                        PERFORM create_trade(
                            instrument_instance,
                            trade_price_var,
                            trade_amount_var,
                            taker_trade_order_instance,
                            maker_book_order_instance,
                            taker_trade_order_instance
                        );
                    ELSE
                        PERFORM create_trade(
                            instrument_instance,
                            trade_price_var,
                            trade_amount_var,
                            maker_book_order_instance,
                            taker_trade_order_instance,
                            taker_trade_order_instance
                        );
                    END IF;

                    -- activate crossing stop orders
                    trigger_loop_restart = activate_crossing_stop_orders(
                        instrument_instance.id, 
                        opposite_side_var::order_side, 
                        trade_price_var
                    );

                    -- immediately trigger exit from loop
                    IF trigger_loop_restart IS TRUE THEN
                        EXIT market_matching_loop;
                    END IF;

                    IF amount_param = 0 THEN
                        EXIT market_matching_loop;
                    END IF;
                
                ELSE
                    exit market_matching_loop;
                END IF;
            END LOOP;
        
            -- process limit orders
            IF amount_param > 0 THEN
                -- execute trades for amount
                <<limit_matching_loop>>
                FOR book_order_instance
                    IN SELECT * FROM get_crossing_limit_orders(
                        instrument_instance.id, 
                        opposite_side_var, 
                        price_param, 
                        trading_account_instance.id
                    )
                    LOOP
                        SELECT * FROM trade_order
                        WHERE id = book_order_instance.trade_order_id
                        INTO maker_book_order_instance;

                        trade_price_var = get_trade_price(
                            side_param,
                            order_type_param::order_type,
                            price_param,
                            opposite_side_var,
                            'LIMIT'::order_type,
                            maker_book_order_instance.price,
                            instrument_instance.id
                        );

                        IF side_param = 'BUY' AND order_type_param = 'MARKET' THEN
                            -- market buy in quote currency
                            trade_amount_var = banker_round(amount_param / maker_book_order_instance.price, quote_currency_precision);
                            
                            -- ensure to trade at available book order amount
                            IF maker_book_order_instance.open_amount < trade_amount_var THEN
                                trade_amount_var = maker_book_order_instance.open_amount;
                                book_order_volume_var = 0;
                            ELSE
                                book_order_volume_var = 
                                    maker_book_order_instance.open_amount - trade_amount_var;
                            END IF;
                            -- incoming order leftover
                            amount_param = amount_param - banker_round(trade_amount_var * maker_book_order_instance.price, quote_currency_precision);
                        ELSE
                            IF maker_book_order_instance.open_amount < amount_param THEN
                                trade_amount_var = maker_book_order_instance.open_amount;
                            ELSE 
                                trade_amount_var = amount_param;
                            END IF;

                            book_order_volume_var = maker_book_order_instance.open_amount - trade_amount_var;
            
                            -- incoming order leftover
                            amount_param = amount_param - trade_amount_var;
                        END IF;
                        
                        IF book_order_volume_var = 0 THEN
                            DELETE FROM book_order
                            WHERE id = book_order_instance.id;
                        END IF;
                        
                        IF side_param = 'SELL' THEN
                            PERFORM create_trade(
                                instrument_instance,
                                trade_price_var,
                                trade_amount_var,
                                taker_trade_order_instance,
                                maker_book_order_instance,
                                taker_trade_order_instance
                            );
                        ELSE
                            PERFORM create_trade(
                                instrument_instance,
                                trade_price_var,
                                trade_amount_var,
                                maker_book_order_instance,
                                taker_trade_order_instance,
                                taker_trade_order_instance
                            );
                        END IF;

                        -- activate crossing stop orders
                        trigger_loop_restart = activate_crossing_stop_orders(
                            instrument_instance.id, 
                            opposite_side_var::order_side, 
                            trade_price_var
                        );

                        -- immediately trigger exit from loop
                        IF trigger_loop_restart IS TRUE THEN
                            EXIT limit_matching_loop;
                        END IF;

                        IF amount_param = 0 THEN
                            EXIT limit_matching_loop;
                        END IF;
                    END LOOP;
            END IF;
            
            IF trigger_loop_restart IS TRUE THEN
                trigger_loop_restart := FALSE;
            ELSE 
                EXIT matching_loop;  
            END IF;
    END LOOP;
    

    -- in case there anything remains in the open amount, persist it in the order book
    IF amount_param > 0 THEN
        -- unless specific order type conditions apply
        IF taker_trade_order_instance.time_in_force = 'IOC'::order_time_in_force THEN
            IF taker_trade_order_instance.open_amount != amount_param THEN
                UPDATE trade_order
                SET status = 'PARTIALLY_REJECTED'::trade_order_status
                WHERE id = taker_trade_order_instance.id
                RETURNING * INTO taker_trade_order_instance;
            ELSE
                UPDATE trade_order
                SET status = 'REJECTED'::trade_order_status
                WHERE id = taker_trade_order_instance.id
                RETURNING * INTO taker_trade_order_instance;
            END IF;

            -- release the funds
            -- reserve required amount
            IF side_param = 'SELL' OR (side_param = 'BUY' and order_type_param = 'MARKET') THEN
                -- release reserved  amount
                UPDATE payment_account 
                SET amount_reserved = amount_reserved - amount_param
                WHERE id = payment_account_instance.id;
            ELSE
                -- release reserved amount
                UPDATE payment_account
                SET amount_reserved = amount_reserved - banker_round(amount_param * price_param, base_currency_precision)
                WHERE id = payment_account_instance.id;
            END IF;
        ELSE 
            -- ensure an update order gets passed
            SELECT * FROM trade_order
            WHERE id = taker_trade_order_instance.id
            INTO taker_trade_order_instance;
            
            -- save orderbook_order
            PERFORM create_book_order(
                taker_trade_order_instance
            );
        END IF;
    END IF;

    PERFORM process_crossing_stop_orders(instrument_instance.id, side_param::order_side, trade_price_var);
 
    RETURN taker_trade_order_instance.pub_id;
END;
$$;
-- +goose StatementEnd


-- +goose Down
DROP FUNCTION process_trade_order(TEXT, TEXT, TEXT, order_side, DECIMAL, DECIMAL, TEXT, BIGINT);