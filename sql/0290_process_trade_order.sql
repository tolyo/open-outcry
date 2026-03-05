-- +goose Up
-- +goose StatementBegin

CREATE OR REPLACE FUNCTION
    process_trade_order(
        instrument_account_id_param text,
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
instrument_account_instance instrument_account%ROWTYPE;
    currency_account_instance currency_account%ROWTYPE;
    taker_trade_order_instance trade_order%ROWTYPE; -- saved
    maker_book_order_instance trade_order%ROWTYPE;
    book_order_instance book_order%ROWTYPE;
    instrument_instance instrument%ROWTYPE;

    base_currency_precision INTEGER;
    quote_currency_precision INTEGER;

    opposite_side_var order_side;
    book_order_volume_var NUMERIC;
    total_available_volume_var NUMERIC;
    trade_amount_var NUMERIC;
    order_currency_var text;
    trade_price_var NUMERIC;

    trigger_loop_restart BOOLEAN := FALSE;

    original_amount NUMERIC;
    remaining_amount NUMERIC;

    reserve_amount NUMERIC;
    release_amount NUMERIC;
BEGIN
    -- basic validation
    IF instrument_name_param IS NULL OR length(instrument_name_param) = 0 THEN
        RAISE EXCEPTION 'invalid_instrument';
END IF;

    IF amount_param IS NULL OR amount_param <= 0 THEN
        RAISE EXCEPTION 'invalid_amount';
END IF;

    IF order_type_param IN ('LIMIT','STOPLIMIT') AND (price_param IS NULL OR price_param <= 0) THEN
        RAISE EXCEPTION 'invalid_price';
END IF;

    original_amount := amount_param;
    remaining_amount := amount_param;

    -- load trading account (external mode)
    IF instrument_account_id_param != 'VOID' THEN
SELECT *
FROM instrument_account
WHERE pub_id = instrument_account_id_param
    INTO instrument_account_instance;

IF NOT FOUND THEN
            RAISE EXCEPTION 'instrument_account_instance_not_found';
END IF;
END IF;

    -- load instrument
SELECT *
FROM instrument
WHERE name = instrument_name_param
    INTO instrument_instance;

IF NOT FOUND THEN
        RAISE EXCEPTION 'instrument_instance_not_found';
END IF;

    -- side derived vars (fix := assignment)
    IF side_param = 'SELL' THEN
        opposite_side_var := 'BUY'::order_side;
        order_currency_var := instrument_instance.base_currency;
ELSE
        opposite_side_var := 'SELL'::order_side;
        order_currency_var := instrument_instance.quote_currency;
END IF;

    -- deterministic precision lookups
SELECT c.precision
INTO base_currency_precision
FROM currency c
WHERE c.name = instrument_instance.base_currency;

IF NOT FOUND THEN
        RAISE EXCEPTION 'base_currency_precision_not_found';
END IF;

SELECT c.precision
INTO quote_currency_precision
FROM currency c
WHERE c.name = instrument_instance.quote_currency;

IF NOT FOUND THEN
        RAISE EXCEPTION 'quote_currency_precision_not_found';
END IF;

    IF instrument_account_id_param != 'VOID' THEN
        -- find and lock transfer account row to avoid reservation races
SELECT pa.*
FROM currency_account pa
         INNER JOIN app_entity ae
                    ON pa.app_entity_id = ae.id
WHERE ae.id = instrument_account_instance.app_entity_id
  AND pa.currency_name = order_currency_var
    FOR UPDATE
    INTO currency_account_instance;

IF NOT FOUND THEN
            RAISE EXCEPTION 'currency_account_instance_not_found';
END IF;

        -- normalize inputs to currency precisions
        remaining_amount := round(remaining_amount, base_currency_precision);
        original_amount := remaining_amount;

        IF price_param IS NOT NULL THEN
            price_param := round(price_param, quote_currency_precision);
END IF;

        -- reserve funds (fix precision usage and stable error token)
        IF side_param = 'SELL' OR (side_param = 'BUY' AND order_type_param = 'MARKET') THEN
            reserve_amount := remaining_amount; -- base for SELL, quote for BUY\+MARKET in this model
            IF currency_account_instance.amount - currency_account_instance.amount_reserved < reserve_amount THEN
                RAISE EXCEPTION 'insufficient_funds'
                    USING DETAIL = format(
                        'available=%s required=%s',
                        currency_account_instance.amount - currency_account_instance.amount_reserved,
                        reserve_amount
                    );
END IF;

UPDATE currency_account
SET amount_reserved = round(currency_account.amount_reserved + reserve_amount, base_currency_precision)
WHERE id = currency_account_instance.id;
ELSE
            reserve_amount := banker_round(remaining_amount * price_param, quote_currency_precision);
            IF currency_account_instance.amount - currency_account_instance.amount_reserved < reserve_amount THEN
                RAISE EXCEPTION 'insufficient_funds'
                    USING DETAIL = format(
                        'available=%s required=%s',
                        currency_account_instance.amount - currency_account_instance.amount_reserved,
                        reserve_amount
                    );
END IF;

UPDATE currency_account
SET amount_reserved = round(currency_account.amount_reserved + reserve_amount, quote_currency_precision)
WHERE id = currency_account_instance.id;
END IF;

        -- create trade order
INSERT INTO trade_order (
    instrument_account_id,
    instrument_id,
    order_type,
    side,
    price,
    amount,
    open_amount,
    time_in_force
)
VALUES (
           instrument_account_instance.id,
           instrument_instance.id,
           order_type_param::order_type,
           side_param,
           price_param,
           original_amount,
           original_amount,
           time_in_force_param::order_time_in_force
       )
    RETURNING * INTO taker_trade_order_instance;

ELSE
        -- internal mode
SELECT *
FROM trade_order
WHERE id = trade_order_id_param
    INTO taker_trade_order_instance;

SELECT *
FROM instrument_account
WHERE id = taker_trade_order_instance.instrument_account_id
    INTO instrument_account_instance;

-- in internal mode, remaining\_amount should start from persisted open\_amount
remaining_amount := taker_trade_order_instance.open_amount;
        original_amount := taker_trade_order_instance.amount;
END IF;

    -- stop orders: persist and exit (no matching)
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
        total_available_volume_var =
            get_available_market_volume(instrument_instance.id, opposite_side_var)
            + get_available_limit_volume(instrument_instance.id, opposite_side_var, price_param)
            - get_potential_self_trade_volume(instrument_instance.id, opposite_side_var, instrument_account_instance.id, price_param);

        IF taker_trade_order_instance.time_in_force = 'FOK'::order_time_in_force
           AND total_available_volume_var < remaining_amount THEN

UPDATE trade_order
SET status = 'REJECTED'::trade_order_status
WHERE id = taker_trade_order_instance.id;

IF instrument_account_id_param != 'VOID' THEN
                IF side_param = 'SELL' OR (side_param = 'BUY' AND order_type_param = 'MARKET') THEN
                    release_amount := remaining_amount;
UPDATE currency_account
SET amount_reserved = round(currency_account.amount_reserved - release_amount, base_currency_precision)
WHERE id = currency_account_instance.id;
ELSE
                    release_amount := banker_round(remaining_amount * price_param, quote_currency_precision);
UPDATE currency_account
SET amount_reserved = round(currency_account.amount_reserved - release_amount, quote_currency_precision)
WHERE id = currency_account_instance.id;
END IF;
END IF;

RETURN taker_trade_order_instance.pub_id;
END IF;

        -- execute market order trades
<<market_matching_loop>>
        FOR maker_book_order_instance
            IN SELECT t.*
               FROM trade_order t
                        INNER JOIN book_order b
                                   ON b.trade_order_id = t.id
               WHERE t.instrument_id = instrument_instance.id
                 AND t.instrument_account_id != instrument_account_instance.id
                 AND t.side = opposite_side_var
                 AND t.order_type = 'MARKET'::order_type
               ORDER BY t.created_at
                   LOOP
                   trade_price_var := get_trade_price(
                   side_param::order_side,
                   order_type_param::order_type,
                   price_param,
                   opposite_side_var,
                   'MARKET'::order_type,
                   0,
                   instrument_instance.id
                   );

IF trade_price_var IS NULL OR trade_price_var <= 0 OR remaining_amount <= 0 THEN
                EXIT market_matching_loop;
END IF;

            IF side_param = 'SELL' THEN
                trade_amount_var :=
                    banker_round(maker_book_order_instance.open_amount / trade_price_var, quote_currency_precision);
                book_order_volume_var :=
                    maker_book_order_instance.open_amount
                    - banker_round(trade_amount_var * trade_price_var, quote_currency_precision);

                remaining_amount := remaining_amount - trade_amount_var;
ELSE
                IF order_type_param = 'MARKET' THEN
                    IF maker_book_order_instance.open_amount < banker_round(remaining_amount / trade_price_var, quote_currency_precision) THEN
                        trade_amount_var := maker_book_order_instance.open_amount;
ELSE
                        trade_amount_var := banker_round(remaining_amount / trade_price_var, quote_currency_precision);
END IF;

                    remaining_amount := remaining_amount
                        - banker_round(trade_amount_var * trade_price_var, quote_currency_precision);
END IF;

                IF order_type_param = 'LIMIT' THEN
                    IF maker_book_order_instance.open_amount < remaining_amount THEN
                        trade_amount_var := maker_book_order_instance.open_amount;
ELSE
                        trade_amount_var := remaining_amount;
END IF;

                    remaining_amount := remaining_amount - trade_amount_var;
END IF;

                book_order_volume_var := maker_book_order_instance.open_amount - trade_amount_var;
END IF;

            IF book_order_volume_var = 0 THEN
DELETE FROM book_order
WHERE trade_order_id = maker_book_order_instance.id;
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

            trigger_loop_restart := activate_crossing_stop_orders(
                instrument_instance.id,
                opposite_side_var::order_side,
                trade_price_var
            );

            IF trigger_loop_restart IS TRUE OR remaining_amount = 0 THEN
                EXIT market_matching_loop;
END IF;
END LOOP;

        -- process limit orders
        IF remaining_amount > 0 THEN
            <<limit_matching_loop>>
            FOR book_order_instance
                IN SELECT *
                   FROM get_crossing_limit_orders(
                                instrument_instance.id,
                                opposite_side_var,
                                price_param,
                                instrument_account_instance.id
                        )
                            LOOP
SELECT *
FROM trade_order
WHERE id = book_order_instance.trade_order_id
    INTO maker_book_order_instance;

trade_price_var := get_trade_price(
                    side_param,
                    order_type_param::order_type,
                    price_param,
                    opposite_side_var,
                    'LIMIT'::order_type,
                    maker_book_order_instance.price,
                    instrument_instance.id
                );

                IF trade_price_var IS NULL OR trade_price_var <= 0 THEN
                    EXIT limit_matching_loop;
END IF;

                IF side_param = 'BUY' AND order_type_param = 'MARKET' THEN
                    trade_amount_var := banker_round(remaining_amount / maker_book_order_instance.price, quote_currency_precision);

                    IF maker_book_order_instance.open_amount < trade_amount_var THEN
                        trade_amount_var := maker_book_order_instance.open_amount;
                        book_order_volume_var := 0;
ELSE
                        book_order_volume_var := maker_book_order_instance.open_amount - trade_amount_var;
END IF;

                    remaining_amount := remaining_amount
                        - banker_round(trade_amount_var * maker_book_order_instance.price, quote_currency_precision);
ELSE
                    IF maker_book_order_instance.open_amount < remaining_amount THEN
                        trade_amount_var := maker_book_order_instance.open_amount;
ELSE
                        trade_amount_var := remaining_amount;
END IF;

                    book_order_volume_var := maker_book_order_instance.open_amount - trade_amount_var;
                    remaining_amount := remaining_amount - trade_amount_var;
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

                trigger_loop_restart := activate_crossing_stop_orders(
                    instrument_instance.id,
                    opposite_side_var::order_side,
                    trade_price_var
                );

                IF trigger_loop_restart IS TRUE OR remaining_amount = 0 THEN
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

    -- persist leftover or reject per TIF
    IF remaining_amount > 0 THEN
        IF taker_trade_order_instance.time_in_force = 'IOC'::order_time_in_force THEN
            IF taker_trade_order_instance.open_amount != remaining_amount THEN
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

            IF instrument_account_id_param != 'VOID' THEN
                IF side_param = 'SELL' OR (side_param = 'BUY' AND order_type_param = 'MARKET') THEN
UPDATE currency_account
SET amount_reserved = round(currency_account.amount_reserved - remaining_amount, base_currency_precision)
WHERE id = currency_account_instance.id;
ELSE
UPDATE currency_account
SET amount_reserved =
        round(currency_account.amount_reserved - banker_round(remaining_amount * price_param, quote_currency_precision), quote_currency_precision)
WHERE id = currency_account_instance.id;
END IF;
END IF;
ELSE
SELECT *
FROM trade_order
WHERE id = taker_trade_order_instance.id
    INTO taker_trade_order_instance;

PERFORM create_book_order(taker_trade_order_instance);
END IF;
END IF;

    PERFORM process_crossing_stop_orders(instrument_instance.id, side_param::order_side, trade_price_var);

RETURN taker_trade_order_instance.pub_id;
END;
$$;
-- +goose StatementEnd

-- +goose Down
DROP FUNCTION process_trade_order(TEXT, TEXT, TEXT, order_side, DECIMAL, DECIMAL, TEXT, BIGINT);