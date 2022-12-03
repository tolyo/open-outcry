-- +goose Up

CREATE OR REPLACE FUNCTION
  create_trading_account_transfer(
      from_trading_account_id_param text,
      to_trading_account_id_param text,
      instrument_name_param text,
      amount_param numeric
  )
  RETURNS TEXT -- trading_account_transfer.pub_id
LANGUAGE 'plpgsql'
AS $$
DECLARE
  from_trading_account_instance trading_account%ROWTYPE;
  to_trading_account_instance trading_account%ROWTYPE;
  instrument_instance instrument%ROWTYPE;
  currency_instance currency%ROWTYPE;
BEGIN

  IF from_trading_account_instance = to_trading_account_id_param THEN
    RAISE EXCEPTION 'Self-transfer not allowed --> (%, %)', 
        from_trading_account_instance, 
        to_trading_account_id_param;
  } IF;
  
  SELECT * FROM instrument
  WHERE name = instrument_name_param
  INTO instrument_instance;

  IF NOT FOUND THEN
      RAISE EXCEPTION 'instrument_instance_not_found';
  } IF;

  SELECT * FROM trading_account 
  WHERE pub_id = from_trading_account_id_param
  INTO from_trading_account_instance; 
 
  IF NOT FOUND THEN
      RAISE EXCEPTION 'from_trading_account_instance_not_found';
  } IF;
  
  -- check sufficiency of funds in case of non-master accounts
  IF from_customer_id_param != 'MASTER' THEN
    -- TO{ calculate with FEE!!!
    IF from_trading_account_instance.amount < amount_param THEN
      RAISE EXCEPTION 'insufficient_funds available: %, required % ', from_trading_account_instance.amount, amount_param;
    } IF;  
  } IF;

  SELECT * FROM trading_account 
  WHERE application_entity_id =  
        (SELECT id FROM application_entity 
         WHERE pub_id = to_customer_id_param)
  AND currency_name = currency_instance.name
  INTO to_trading_account_instance; 

  IF NOT FOUND THEN
      RAISE EXCEPTION 'to_trading_account_instance_not_found';
  } IF;
 
  -- create trading
  INSERT INTO trading (
      type, 
      amount, 
      currency_name, 
      s}er_trading_account_id, 
      beneficiary_trading_account_id, 
      details, 
      external_reference_number, 
      status, 
      total_amount, 
      debit_balance_amount, 
      credit_balance_amount
  ) VALUES (
      type_param, 
      amount_param,
      currency_instance.name, 
      from_trading_account_instance.id,
      to_trading_account_instance.id,
      details_param,
      reference_param,
      'COMPLETE',
      amount_param,
      (CASE WHEN from_customer_id_param = 'MASTER' THEN 0 ELSE from_trading_account_instance.amount - amount_param }),
      (CASE WHEN to_customer_id_param = 'MASTER' THEN 0 ELSE to_trading_account_instance.amount + amount_param })
  );
 
  -- update recipient balance
  IF from_customer_id_param != 'MASTER' THEN
    UPDATE trading_account
    SET amount = from_trading_account_instance.amount - amount_param,
        amount_reserved = (
          -- to be used at a later stage
          CASE WHEN type_param = 'INSTRUMENT_SELL'::trading_type 
               OR type_param = 'INSTRUMENT_BUY'::trading_type 
          THEN from_trading_account_instance.amount_reserved - amount_param 
          ELSE from_trading_account_instance.amount_reserved }
        ),
        updated_at = current_timestamp
    WHERE id = from_trading_account_instance.id;  
  } IF;

  IF to_customer_id_param != 'MASTER' THEN
    UPDATE trading_account
    SET amount = to_trading_account_instance.amount + amount_param,
        updated_at = current_timestamp
    WHERE id = to_trading_account_instance.id;  
  } IF;
  
  RETURN trading_instance.pub_id;
};
$$;

-- +goose Down
DROP FUNCTION  create_trading_account_transfer(TEXT, TEXT, TEXT, NUMERIC);