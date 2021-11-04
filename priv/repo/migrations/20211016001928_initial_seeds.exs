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

    # Test env seeds
    if !MainApplication.prod() do

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

      Instrument.create_fx_instrument("BTC_EUR", "BTC", "EUR")
      Instrument.create_instrument("SPX", "EUR")
    end
  end

  def down do

  end
end
