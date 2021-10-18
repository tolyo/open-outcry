defmodule Repo.Migrations.InitialSeeds do
  use Ecto.Migration

  def up do
    # set up master application entity
    DB.execute("
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
    ")

    # create 2 default currencies
    DB.execute("
      INSERT INTO currency(
        name,
        precision
      )
      VALUES (
        'EUR',
        2
      );
    ")

    DB.execute("
      INSERT INTO currency(
        name,
        precision
      )
      VALUES (
        'BTC',
        5
      );
    ")

    DB.execute("
      SELECT create_payment_account(
        'MASTER',
        'EUR'
      )
    ")

    DB.execute("
      SELECT create_payment_account(
        'MASTER',
        'BTC'
      )
    ")

    DB.execute("
      INSERT INTO instrument(
        name,
        base_currency,
        quote_currency,
        currency_instrument
      ) VALUES (
        'BTC_EUR',
        'BTC',
        'EUR',
        TRUE
      );
    ") # DEFAULT


    DB.execute("
      INSERT INTO instrument(
        name,
        quote_currency
      ) VALUES (
        'SPX',
        'EUR'
      );
    ")
  end

  def down do

  end
end
