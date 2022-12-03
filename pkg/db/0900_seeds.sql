-- +goose Up

 INSERT INTO application_entity (
        pub_id,
        external_id,
        type
      )
      VALUES (
        'MASTER',
        'MASTER',
        'MASTER'
      );


        SELECT create_payment_account(
          'MASTER',
          'EUR'
        )

        SELECT create_payment_account(
          'MASTER',
          'BTC'
        )