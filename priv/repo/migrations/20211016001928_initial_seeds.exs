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

    DB.execute("""
      SELECT create_payment_account(
        'MASTER',
        'EUR'
      )
    """)

    DB.execute("""
      SELECT create_payment_account(
        'MASTER',
        'BTC'
      )
    """)

    DB.execute("INSERT INTO instrument(name, precision, base_currency, quote_currency) VALUES ('BTC_EUR', 5, 'BTC', 'EUR');") # DEFAULT
  end

  def down do

  end
end
