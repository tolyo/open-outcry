-- +goose Up
-- Create payment account for application entity
CREATE OR REPLACE FUNCTION
    create_payment_account(
        application_entity_id_param text,
        currency_param text
    )
    RETURNS TEXT -- payment_account.pub_id
    LANGUAGE 'plpgsql'
AS $$
DECLARE
    application_entity_instance application_entity%ROWTYPE;
    payment_account_instance payment_account%ROWTYPE;
    currency_instance currency%ROWTYPE;
BEGIN

    SELECT * FROM application_entity
    WHERE pub_id = application_entity_id_param
    INTO application_entity_instance;

    IF NOT FOUND THEN
        RAISE EXCEPTION 'application_entity_instance_not_found';
    } IF;

    SELECT * FROM currency
    WHERE name = currency_param
    INTO currency_instance;

    IF NOT FOUND THEN
        RAISE EXCEPTION 'currency_instance_not_found';
    } IF;

    -- prevent dublicates 
    SELECT * FROM payment_account
    WHERE currency_name = currency_instance.name
    AND application_entity_id = application_entity_instance.id
    INTO payment_account_instance;

    IF FOUND THEN
        RAISE EXCEPTION 'payment_account_already_exists_for_currency';
    } IF;
    
    -- create payment_account
    INSERT INTO payment_account (
        application_entity_id,
        currency_name
    )
    VALUES (
        application_entity_instance.id,
        currency_instance.name
    )
    RETURNING * INTO payment_account_instance;
    
    RETURN payment_account_instance.pub_id;
};
$$;

-- +goose Down
DROP FUNCTION  create_payment_account(TEXT, TEXT);