-- +goose Up
-- +goose StatementBegin

-- Create payment account for application entity; return an existing account if already exists
CREATE OR REPLACE FUNCTION
    create_payment_account(
        app_entity_id_param text,
        currency_param text
    )
    RETURNS TEXT -- payment_account.pub_id
    LANGUAGE 'plpgsql'
AS $$
DECLARE
    app_entity_instance app_entity%ROWTYPE;
    payment_account_instance payment_account%ROWTYPE;
    currency_instance currency%ROWTYPE;
BEGIN

    SELECT * FROM app_entity
    WHERE pub_id = app_entity_id_param
    INTO app_entity_instance;

    IF NOT FOUND THEN
        RAISE EXCEPTION 'app_entity_instance_not_found';
    END IF;

    SELECT * FROM currency
    WHERE name = currency_param
    INTO currency_instance;

    IF NOT FOUND THEN
        RAISE EXCEPTION 'currency_instance_not_found';
    END IF;

    -- prevent duplicates
    SELECT * FROM payment_account
    WHERE currency_name = currency_instance.name
    AND app_entity_id = app_entity_instance.id
    INTO payment_account_instance;

    IF NOT FOUND THEN
        -- create payment_account
        INSERT INTO payment_account (
            app_entity_id,
            currency_name
        )
        VALUES (
           app_entity_instance.id,
           currency_instance.name
        )
        RETURNING * INTO payment_account_instance;
    END IF;


    RETURN payment_account_instance.pub_id;
END;
$$;
-- +goose StatementEnd

-- +goose Down
DROP FUNCTION create_payment_account(TEXT, TEXT);
