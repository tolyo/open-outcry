-- +goose Up
-- +goose StatementBegin
CREATE OR REPLACE FUNCTION
  create_transfer(
      type_param transfer_type,
      from_customer_id_param text,
      amount_param numeric,
      currency_param text,
      to_customer_id_param text,
      reference_param text,
      details_param text
  )
  RETURNS TEXT -- transfer_instance.pub_id
LANGUAGE 'plpgsql'
AS $$
DECLARE
    from_currency_account_instance currency_account%ROWTYPE;
    to_currency_account_instance currency_account%ROWTYPE;
    transfer_instance transfer%ROWTYPE;
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
  SELECT * FROM currency_account
  WHERE app_entity_id =
        (SELECT id FROM app_entity
         WHERE pub_id = from_customer_id_param)
  AND currency_name = currency_instance.name
  INTO from_currency_account_instance;
  IF NOT FOUND THEN
      RAISE EXCEPTION 'from_currency_account_instance_not_found';
  END IF;
  -- check sufficiency of funds in case of non-master accounts
  IF from_customer_id_param != 'MASTER' THEN
    IF from_currency_account_instance.amount < amount_param THEN
      RAISE EXCEPTION 'insufficient_funds available: %, required % ', from_currency_account_instance.amount, amount_param;
    END IF;
  END IF;
  SELECT * FROM currency_account
  WHERE app_entity_id =
        (SELECT id FROM app_entity
         WHERE pub_id = to_customer_id_param)
  AND currency_name = currency_instance.name
  INTO to_currency_account_instance;
  IF NOT FOUND THEN
      RAISE EXCEPTION 'to_currency_account_instance_not_found';
  END IF;
  -- 1. Create the journal entry (transfer header)
  INSERT INTO transfer (
      type,
      amount,
      currency_name,
      details,
      external_reference_number,
      status
  ) VALUES (
      type_param,
      amount_param,
      currency_instance.name,
      details_param,
      reference_param,
      'COMPLETE'
  ) RETURNING * INTO transfer_instance;
  -- 2. Create DEBIT ledger entry (sender side: decrease)
  INSERT INTO transfer_ledger_entry (
      transfer_id,
      currency_account_id,
      entry_type,
      amount,
      resulting_balance
  ) VALUES (
      transfer_instance.id,
      from_currency_account_instance.id,
      'DEBIT',
      amount_param,
      (CASE WHEN from_customer_id_param = 'MASTER' THEN 0 ELSE from_currency_account_instance.amount - amount_param END)
  );
  -- 3. Create CREDIT ledger entry (receiver side: increase)
  INSERT INTO transfer_ledger_entry (
      transfer_id,
      currency_account_id,
      entry_type,
      amount,
      resulting_balance
  ) VALUES (
      transfer_instance.id,
      to_currency_account_instance.id,
      'CREDIT',
      amount_param,
      (CASE WHEN to_customer_id_param = 'MASTER' THEN 0 ELSE to_currency_account_instance.amount + amount_param END)
  );
  -- 4. Update sender balance (debit reduces the balance)
  IF from_customer_id_param != 'MASTER' THEN
    UPDATE currency_account
    SET amount = from_currency_account_instance.amount - amount_param,
        amount_reserved = (
          CASE WHEN type_param = 'INSTRUMENT_SELL'::transfer_type
               OR type_param = 'INSTRUMENT_BUY'::transfer_type
          THEN from_currency_account_instance.amount_reserved - amount_param
          ELSE from_currency_account_instance.amount_reserved END
        ),
        updated_at = current_timestamp
    WHERE id = from_currency_account_instance.id;
  END IF;
  -- 5. Update receiver balance (credit increases the balance)
  IF to_customer_id_param != 'MASTER' THEN
    UPDATE currency_account
    SET amount = to_currency_account_instance.amount + amount_param,
        updated_at = current_timestamp
    WHERE id = to_currency_account_instance.id;
  END IF;
  RETURN transfer_instance.pub_id;
END;
$$;
-- +goose StatementEnd
-- +goose Down
DROP FUNCTION create_transfer(transfer_type, TEXT, NUMERIC, TEXT, TEXT, TEXT, TEXT);
