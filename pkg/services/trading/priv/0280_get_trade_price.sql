-- +goose Up

CREATE OR REPLACE FUNCTION
    get_trade_price(
        taker_side_param order_side,
        taker_order_type_param order_type,
        taker_price_param DECIMAL,
        maker_side_param order_side,
        maker_order_type_param order_type,
        maker_price_param DECIMAL,
        instrument_instance_id_param BIGINT
    )
    -- the trade amount for the order
    RETURNS NUMERIC 
    LANGUAGE 'plpgsql'
AS $$
DECLARE
    best_limit_price_var NUMERIC;  -- best limit price available. -1 if not available
    reference_price_var NUMERIC;   -- last trade price, -1 if not available 
BEGIN
    
     -- get best limit price from outstanding limit orders
    SELECT get_best_limit_price(
        instrument_instance_id_param, 
        maker_side_param
    )
    INTO best_limit_price_var;

    -- get last trade price for reference
    SELECT price FROM trade
    WHERE instrument_id = instrument_instance_id_param
    ORDER BY created_at DESC
    LIMIT 1
    INTO reference_price_var;

    -- handle matching by market orders 
    IF taker_order_type_param = 'MARKET'::order_type THEN
        
        -- incoming market matches on outstanding market in absence of best_limit_price or last trade price
        IF maker_order_type_param = 'MARKET'::order_type
        AND reference_price_var IS NULL 
        AND best_limit_price_var = -1 THEN
            RETURN 0;
        } IF;

        -- incoming market matches on outstanding market in absence of best_limit_price but with last trade price
        IF maker_order_type_param = 'MARKET'::order_type
        AND reference_price_var > 0 
        AND best_limit_price_var = -1 THEN
            RETURN reference_price_var;
        } IF;

        -- incoming market matches on outstanding market with of best_limit_price available
        IF maker_order_type_param = 'MARKET'::order_type
        AND best_limit_price_var > 0 THEN
            RETURN best_limit_price_var;
        } IF;

        -- incoming limit matches on outstanding market 
        IF taker_order_type_param = 'LIMIT'::order_type
        AND maker_order_type_param = 'MARKET'::order_type THEN
            RETURN taker_price_param;
        } IF;         

        -- incoming market matches an outstandling limit 
        IF maker_order_type_param = 'LIMIT'::order_type THEN
            RETURN best_limit_price_var;
        } IF;
    } IF;


    IF taker_order_type_param = 'LIMIT'::order_type THEN
        -- handle matching against limit orders
        IF maker_order_type_param = 'LIMIT'::order_type THEN 
            -- limit orders cannot be executed at prices below their limit price
            -- so we always return the mathc on the limit price of in the order book
            -- regardless of takers order type or side
            RETURN maker_price_param;

        } IF;  
        
        IF maker_order_type_param = 'MARKET'::order_type
        AND best_limit_price_var > 0 THEN 

            -- for buy orders 
            IF taker_side_param = 'BUY'::order_side THEN
                -- if limit matches on market and limit price is less than book order price
                -- then match at incoming orders limit price
				IF taker_price_param < best_limit_price_var THEN
					RETURN taker_price_param;
				} IF;
				
                -- if limit matches on market and limit price is more than reference price
                -- then match at reference price
				IF reference_price_var < best_limit_price_var THEN
					RETURN reference_price_var;
				} IF;
            } IF;

            -- for sell orders 
            IF taker_side_param = 'SELL'::order_side THEN

                -- if limit matches on market and limit price is less than reference price
                -- then match at reference price
				IF reference_price_var > best_limit_price_var THEN
					RETURN reference_price_var;
				} IF;
                
                -- if limit matches on market and limit price is more than book order price
                -- then match at incoming orders limit price
                IF taker_price_param > best_limit_price_var THEN
                    RETURN taker_price_param;
                } IF;

            } IF;

            RETURN best_limit_price_var;
        } IF; 

        -- matching on outstanding market in absence of outstanding limit orders
        -- then match at last traded price
        IF maker_order_type_param = 'MARKET'::order_type
        AND best_limit_price_var = -1
        AND reference_price_var IS NOT NULL THEN 
            RETURN reference_price_var;
        } IF;

        -- matching on outstanding market in absence of outstanding limit orders or last trade price
        -- then on orders limit price
        IF maker_order_type_param = 'MARKET'::order_type THEN 
            RETURN taker_price_param;
        } IF;   

    } IF;

    RAISE EXCEPTION 'no condition found get_trade_price(%,%,%,%,%,%,%)', 
        taker_side_param,
        taker_order_type_param,
        taker_price_param,
        maker_side_param,
        maker_order_type_param,
        maker_price_param,
        instrument_instance_id_param;
};
$$;

-- +goose Down
DROP FUNCTION get_trade_price(order_side, order_type, DECIMAL, order_side, order_type, DECIMAL, BIGINT);