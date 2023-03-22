-- +goose Up
-- +goose StatementBegin

CREATE OR REPLACE FUNCTION
  create_payment(
      type_param payment_type,
      from_customer_id_param text,
      amount_param numeric,
      currency_param text,
      to_customer_id_param text,
      reference_param text,
      details_param text
  )
  RETURNS TEXT -- payment_instance.pub_id
LANGUAGE 'plpgsql'
AS $$
DECLARE
    from_payment_account_instance payment_account%ROWTYPE;
    to_payment_account_instance payment_account%ROWTYPE;
    payment_instance payment%ROWTYPE;
    currency_instance currency%ROWTYPE;
BEGIN

  IF from_customer_id_param = to_customer_id_param THEN
    RAISE EXCEPTION 'Self-transfer not allowed --> (%, %)', from_customer_id_param, to_customer_id_param;
  END IF;
  
  SELECT * FROM currency
  WHERE name = currency_param
  INTO currency_instance;

  IF NOT FOUND THEN
      RAISE EXCEPTION 'currency_instance_not_found';
  END IF;

  SELECT * FROM payment_account 
  WHERE app_entity_id =
        (SELECT id FROM app_entity
         WHERE pub_id = from_customer_id_param)
  AND currency_name = currency_instance.name 
  INTO from_payment_account_instance; 
 
  IF NOT FOUND THEN
      RAISE EXCEPTION 'from_payment_account_instance_not_found';
  END IF;
  
  -- check sufficiency of funds in case of non-master accounts
  IF from_customer_id_param != 'MASTER' THEN
    IF from_payment_account_instance.amount < amount_param THEN
      RAISE EXCEPTION 'insufficient_funds available: %, required % ', from_payment_account_instance.amount, amount_param;
    END IF;
  END IF;

  SELECT * FROM payment_account 
  WHERE app_entity_id =
        (SELECT id FROM app_entity
         WHERE pub_id = to_customer_id_param)
  AND currency_name = currency_instance.name
  INTO to_payment_account_instance; 

  IF NOT FOUND THEN
      RAISE EXCEPTION 'to_payment_account_instance_not_found';
  END IF;
 
  -- create payment
  INSERT INTO payment (
      type, 
      amount, 
      currency_name, 
      sender_payment_account_id,
      beneficiary_payment_account_id, 
      details, 
      external_reference_number, 
      status, 
      debit_balance_amount,
      credit_balance_amount
  ) VALUES (
      type_param, 
      amount_param,
      currency_instance.name, 
      from_payment_account_instance.id,
      to_payment_account_instance.id,
      details_param,
      reference_param,
      'COMPLETE',
      (CASE WHEN from_customer_id_param = 'MASTER' THEN 0 ELSE from_payment_account_instance.amount - amount_param END),
      (CASE WHEN to_customer_id_param = 'MASTER' THEN 0 ELSE to_payment_account_instance.amount + amount_param END)
  ) RETURNING * INTO payment_instance;

  -- update recipient balance
  IF from_customer_id_param != 'MASTER' THEN
    UPDATE payment_account
    SET amount = from_payment_account_instance.amount - amount_param,
        amount_reserved = (
          -- to be used at a later stage
          CASE WHEN type_param = 'INSTRUMENT_SELL'::payment_type 
               OR type_param = 'INSTRUMENT_BUY'::payment_type 
          THEN from_payment_account_instance.amount_reserved - amount_param
          ELSE from_payment_account_instance.amount_reserved END
        ),
        updated_at = current_timestamp
    WHERE id = from_payment_account_instance.id;  
  END IF;

  IF to_customer_id_param != 'MASTER' THEN
    UPDATE payment_account
    SET amount = to_payment_account_instance.amount + amount_param,
        updated_at = current_timestamp
    WHERE id = to_payment_account_instance.id;  
  END IF;
  
  RETURN payment_instance.pub_id;
END;
$$;
-- +goose StatementEnd

-- +goose Down
DROP FUNCTION  create_payment(payment_type, TEXT, NUMERIC, TEXT, TEXT, TEXT, TEXT);