-- +goose Up
-- +goose StatementBegin

CREATE OR REPLACE FUNCTION
    process_payment(
        type_param payment_type,
        from_customer_id_param TEXT,
        amount_param NUMERIC,
        currency_param TEXT,
        to_customer_id_param TEXT,
        reference_param TEXT,
        details_param TEXT,
        fee_type_param TEXT
    )
    RETURNS TEXT -- payment_instance.pub_id
    LANGUAGE 'plpgsql'
AS $$
DECLARE
    payment_instance_pub_id_var TEXT;
    currency_instance currency%ROWTYPE;
    fee_instance fee%ROWTYPE;
    fee_amount_var NUMERIC = 0.00;
BEGIN

    payment_instance_pub_id_var = create_payment(type_param, from_customer_id_param, amount_param, currency_param, to_customer_id_param, reference_param, details_param);

    -- create fee payment fee
    IF fee_type_param IS NOT NULL THEN
        SELECT * FROM fee
        WHERE type = fee_type_param AND currency_name = currency_param
        INTO fee_instance;

        IF FOUND THEN
            -- priority is given to percentage
            IF fee_instance.percentage IS NOT NULL THEN
                fee_amount_var = banker_round(amount_param * fee_instance.percentage / 100,  currency_instance.precision);
            END IF;

            -- check min
            IF fee_instance.min IS NOT NULL THEN
                IF fee_amount_var < fee_instance.min THEN
                    fee_amount_var = fee_instance.min;
                END IF;
            END IF;

            -- check max
            IF fee_instance.max IS NOT NULL THEN
                IF fee_amount_var > fee_instance.max THEN
                    fee_amount_var = fee_instance.max;
                END IF;
            END IF;

            PERFORM create_payment(
                'CHARGE'::payment_type,
                to_customer_id_param,
                fee_amount_var,
                currency_param,
                'MASTER',
                payment_instance_pub_id_var,
                payment_instance_pub_id_var
            );
        END IF;
    END IF;


    RETURN payment_instance_pub_id_var;
END;
$$;
-- +goose StatementEnd

-- +goose Down
DROP FUNCTION  process_payment(payment_type, TEXT, NUMERIC, TEXT, TEXT, TEXT, TEXT, TEXT);