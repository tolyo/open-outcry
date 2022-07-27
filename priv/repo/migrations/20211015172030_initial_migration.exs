defmodule Repo.Migrations.InitialMigration do
  use Ecto.Migration

  @doc """
    Returns a list of SQL ENUMS types
  """
  @spec types() :: [String.t()]
  def types() do
    [
      "payment_type",
      "order_fill",
      "order_side",
      "order_type",
      "trade_order_status",
      "order_time_in_force"
    ]
  end

  @doc """
    Returns a list of SQL TABLES
  """
  @spec models() :: [String.t()]
  def models() do
    [
      "application_entity",
      "currency",
      "payment_account",
      "payment",
      "instrument",
      "price_level",
      "trading_account",
      "trading_account_instrument",
      "trading_account_transfer",
      "trade_order",
      "trade",
      "book_order",
      "stop_order"
    ]
  end

  @doc """
    Returns a list of SQL FUNCTIONS
  """
  @spec functions() :: [String.t()]
  def functions() do
    [
      "banker_round(NUMERIC, INTEGER)",
      "create_payment_account(TEXT, TEXT)",
      "create_payment(payment_type, TEXT, NUMERIC, TEXT, TEXT, TEXT, TEXT)",
      "create_client(TEXT)",
      "get_crossing_limit_orders(BIGINT, order_side, NUMERIC, BIGINT)",
      "get_available_limit_volume(BIGINT, order_side, DECIMAL)",
      "get_available_market_volume(BIGINT, order_side)",
      "get_best_limit_price(BIGINT, order_side)",
      "get_fill_type(NUMERIC, NUMERIC)",
      "create_trade(BIGINT, DECIMAL, DECIMAL, BIGINT, BIGINT, BIGINT)",
      "create_book_order(TEXT)",
      "cancel_trade_order(TEXT)",
      "activate_crossing_stop_orders(BIGINT, order_side, NUMERIC)",
      "process_crossing_stop_orders(BIGINT, order_side, NUMERIC)",
      "get_trade_price(order_side, order_type, DECIMAL, order_side, order_type, DECIMAL, BIGINT)",
      "process_trade_order(TEXT, TEXT, TEXT, order_side, DECIMAL, DECIMAL, TEXT, BIGINT)",
      "update_price_level(BIGINT, order_side, DECIMAL, DECIMAL, BOOLEAN)",
      "get_potential_self_trade_volume(BIGINT, order_side, BIGINT, NUMERIC)",
      "create_trading_account_transfer(TEXT, TEXT, TEXT, NUMERIC)"
    ]
  end

  def up do
    # plugins
    DB.execute("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"")

    # types
    types()
    |> Enum.each(fn type -> DB.migrate_up_type(type) end)

    # models
    models()
    |> Enum.each(fn model -> DB.migrate_up_table(model) end)

    functions()
    |> Enum.each(fn function ->
      function
      |> String.split("(")
      |> List.first()
      |> DB.migrate_up_function()
    end)
  end

  def down do
    # migrate down functions
    functions()
    |> Enum.reverse()
    |> Enum.each(fn function -> DB.migrate_down_function(function) end)

    # migrate down models
    models()
    |> Enum.reverse()
    |> Enum.each(fn model -> DB.migrate_down_table(model) end)

    # mirgate down types
    types()
    |> Enum.each(fn type -> DB.migrate_down_type(type) end)

    DB.execute("DROP EXTENSION \"uuid-ossp\"")
  end
end
