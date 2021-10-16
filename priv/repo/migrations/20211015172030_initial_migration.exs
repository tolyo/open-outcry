defmodule Repo.Migrations.InitialMigration do
  use Ecto.Migration

  @doc """
    Returns a list of SQL ENUMS types
  """
  @spec types() :: [String.t]
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
  @spec models() :: [String.t]
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
