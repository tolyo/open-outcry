-- +goose Up

-- +goose StatementBegin
CREATE OR REPLACE FUNCTION
  create_client(
    external_id_param text
  )  
  
  -- app_entity.pub_id
  RETURNS TEXT 

LANGUAGE 'plpgsql'
AS $$
DECLARE
  app_entity_instance app_entity%ROWTYPE;
BEGIN
  -- create app_entity
  INSERT INTO app_entity (external_id, type)
  VALUES (external_id_param, 'CUSTOMER')
  RETURNING * INTO app_entity_instance;

  -- attach payment account for some default currency - assume 'EUR' for now
  PERFORM create_payment_account(app_entity_instance.pub_id, 'EUR');

  -- attach trading account
  INSERT INTO trading_account(app_entity_id)
  VALUES (app_entity_instance.id);

  -- TO{ onboarding state
  RETURN app_entity_instance.pub_id;
END;
$$;

-- +goose StatementEnd

-- +goose Down
DROP FUNCTION create_client(TEXT);