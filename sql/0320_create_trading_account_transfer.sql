-- +goose Up

-- +goose StatementBegin
CREATE OR REPLACE FUNCTION
  create_trading_account_transfer(
      from_trading_account_id_param text,
      to_trading_account_id_param text,
      instrument_param instrument,
      amount_param integer
  )
  RETURNS TEXT -- trading_account_transfer.pub_id
LANGUAGE 'plpgsql'
AS $$
DECLARE
  from_trading_account_instance trading_account%ROWTYPE;
  to_trading_account_instance trading_account%ROWTYPE;
  from_tai trading_account_instrument%ROWTYPE;
  to_tai trading_account_instrument%ROWTYPE;
  transfer_instance trading_account_transfer%ROWTYPE;
BEGIN

  -- Validate: no self-transfer
  IF from_trading_account_id_param = to_trading_account_id_param THEN
    RAISE EXCEPTION 'Self-transfer not allowed --> (%, %)',
        from_trading_account_id_param,
        to_trading_account_id_param;
  END IF;

  -- Look up sender trading account
  SELECT * FROM trading_account
  WHERE pub_id = from_trading_account_id_param
  INTO from_trading_account_instance;

  IF NOT FOUND THEN
      RAISE EXCEPTION 'from_trading_account_instance_not_found';
  END IF;

  -- Look up receiver trading account
  SELECT * FROM trading_account
  WHERE pub_id = to_trading_account_id_param
  INTO to_trading_account_instance;

  IF NOT FOUND THEN
      RAISE EXCEPTION 'to_trading_account_instance_not_found';
  END IF;

  -- Look up sender instrument holding (must exist and have sufficient balance)
  SELECT * FROM trading_account_instrument
  WHERE trading_account = from_trading_account_instance.id
    AND instrument_id = instrument_param.id
  INTO from_tai;

  IF NOT FOUND THEN
      RAISE EXCEPTION 'from_trading_account_instrument_not_found';
  END IF;

  IF from_tai.amount < amount_param THEN
      RAISE EXCEPTION 'insufficient_instrument_balance: available %, required %',
          from_tai.amount, amount_param;
  END IF;

  -- Look up or create receiver instrument holding
  SELECT * FROM trading_account_instrument
  WHERE trading_account = to_trading_account_instance.id
    AND instrument_id = instrument_param.id
  INTO to_tai;

  IF NOT FOUND THEN
      INSERT INTO trading_account_instrument (trading_account, instrument_id, amount, amount_reserved)
      VALUES (to_trading_account_instance.id, instrument_param.id, 0, 0)
      RETURNING * INTO to_tai;
  END IF;

  -- 1. Create the journal entry (transfer header)
  INSERT INTO trading_account_transfer (
      instrument_id,
      amount,
      details,
      external_reference_number
  ) VALUES (
      instrument_param.id,
      amount_param,
      'Transfer of ' || instrument_param.name,
      NULL
  ) RETURNING * INTO transfer_instance;

  -- 2. Create DEBIT ledger entry (sender side: decrease)
  INSERT INTO trading_account_ledger_entry (
      transfer_id,
      trading_account_instrument_id,
      entry_type,
      amount,
      resulting_balance
  ) VALUES (
      transfer_instance.id,
      from_tai.id,
      'DEBIT',
      amount_param,
      from_tai.amount - amount_param
  );

  -- 3. Create CREDIT ledger entry (receiver side: increase)
  INSERT INTO trading_account_ledger_entry (
      transfer_id,
      trading_account_instrument_id,
      entry_type,
      amount,
      resulting_balance
  ) VALUES (
      transfer_instance.id,
      to_tai.id,
      'CREDIT',
      amount_param,
      to_tai.amount + amount_param
  );

  -- 4. Update sender balance (debit reduces the balance)
  UPDATE trading_account_instrument
  SET amount = from_tai.amount - amount_param,
      amount_reserved = (
        CASE WHEN from_tai.amount_reserved >= amount_param
        THEN from_tai.amount_reserved - amount_param
        ELSE 0 END
      ),
      updated_at = current_timestamp
  WHERE id = from_tai.id;

  -- 5. Update receiver balance (credit increases the balance)
  UPDATE trading_account_instrument
  SET amount = to_tai.amount + amount_param,
      updated_at = current_timestamp
  WHERE id = to_tai.id;

  RETURN transfer_instance.pub_id;
END;
$$;
-- +goose StatementEnd

-- +goose Down
DROP FUNCTION create_trading_account_transfer(TEXT, TEXT, instrument, INTEGER);
