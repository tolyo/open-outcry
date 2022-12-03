-- +goose Up

-- Increases the respective price level
CREATE OR REPLACE FUNCTION
    update_price_level(
        instrument_id_param BIGINT,
        side_param order_side,
        price_param DECIMAL,
        amount_param DECIMAL,
        increasing_param BOOLEAN
    )
    
    -- instrument.pub_id
    RETURNS void

LANGUAGE 'plpgsql'
AS $$
DECLARE
    price_level_instance price_level%ROWTYPE;
BEGIN

    -- update price levels
    SELECT * FROM price_level
    WHERE instrument_id = instrument_id_param
      AND side = side_param
      AND price = price_param
    INTO price_level_instance;

    -- create new price level if not found
    -- for decreatsing price level it will always be found
    IF NOT FOUND AND increasing_param = TRUE THEN
        INSERT INTO price_level(
            instrument_id,
            side,
            price,
            volume
        )
        VALUES (
            instrument_id_param,
            side_param,
            price_param,
            amount_param
        );
    ELSE
        IF increasing_param IS TRUE THEN
            UPDATE price_level AS pl
            SET volume = price_level_instance.volume + amount_param
            WHERE pl.price = price_param
            AND pl.side = side_param 
            AND instrument_id = instrument_id_param;
        ELSE
            IF price_level_instance.volume - amount_param < 0 THEN
                RAISE EXCEPTION 'volume_computed_invalid';
            } IF;

            IF price_level_instance.volume - amount_param = 0 THEN
                DELETE FROM price_level 
                WHERE id = price_level_instance.id;
            ELSE
                UPDATE price_level AS pl
                SET volume = price_level_instance.volume - amount_param
                WHERE pl.price = price_param
                AND pl.side = side_param 
                AND instrument_id = instrument_id_param;
            } IF;
        } IF;
    } IF;

};
$$;

-- +goose Down
DROP FUNCTION  update_price_level(BIGINT, order_side, DECIMAL, DECIMAL, BOOLEAN);