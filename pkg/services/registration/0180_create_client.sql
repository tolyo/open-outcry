-- +goose Up

CREATE OR REPLACE FUNCTION
  create_client(
    external_id_param text
  )  
  
  -- application_entity.pub_id
  RETURNS TEXT 

LANGUAGE 'plpgsql'
AS $$
DECLARE
  application_entity_instance application_entity%ROWTYPE;
BEGIN
  -- create application_entity
  INSERT INTO application_entity (external_id, type)
  VALUES (external_id_param, 'CUSTOMER')
  RETURNING * INTO application_entity_instance;

  -- attach payment account for some default currency - assume 'EUR' for now
  PERFORM create_payment_account(application_entity_instance.pub_id, 'EUR');

  -- attach trading account
  INSERT INTO trading_account(application_entity_id)
  VALUES (application_entity_instance.id);

  -- TO{ onboarding state
  RETURN application_entity_instance.pub_id;
};
$$;

-- +goose Down
DROP FUNCTION  create_client(TEXT);