defmodule MatchingService do
  require Logger

  @doc """
    Main entry point for processing an market order.
      - For BUY side, the amount must be alocated in quote currency.
      - For SELL side the amount must be allocared in base currency
  """
  @spec create(
          TradingAccount.id(),
          Instrument.name(),
          TradeOrderType.t(),
          OrderSide.t(),
          TradeOrder.amount(),
          OrderTimeInForce.t()
        ) :: TradeOrder.id()
  def create(trading_account_id, instrument_name, :MARKET, side, amount, time_in_force) do
    %TradeOrder{
      trading_account_id: trading_account_id,
      instrument_name: instrument_name,
      type: :MARKET,
      side: side,
      price: 0,
      amount: amount,
      time_in_force: time_in_force
    }
    |> process_trade_order()
  end

  @spec create(
          TradingAccount.id(),
          Instrument.name(),
          TradeOrderType.t(),
          OrderSide.t(),
          TradeOrder.price(),
          TradeOrder.amount(),
          OrderTimeInForce.t()
        ) :: TradeOrder.id()
  def create(trading_account_id, instrument_name, order_type, side, price, amount, time_in_force) do
    case order_type do
      order_type_val when order_type_val in [:LIMIT, :STOPLOSS, :STOPLIMIT] ->
        %TradeOrder{
          trading_account_id: trading_account_id,
          instrument_name: instrument_name,
          type: order_type,
          side: side,
          price: price,
          amount: amount,
          time_in_force: time_in_force
        }
        |> process_trade_order()

      _ ->
        raise ArgumentError, "Invalid order type: #{order_type}"
    end
  end

  @spec cancel_trade_order(TradeOrder.id()) :: :ok
  def cancel_trade_order(trade_order_id) do
    trade_order_id
    |> DB.query_val("SELECT cancel_trade_order($1)")

    :ok
  end

  defp process_trade_order(%TradeOrder{
         trading_account_id: trading_account_id,
         instrument_name: instrument_name,
         type: type,
         side: side,
         price: price,
         amount: amount,
         time_in_force: time_in_force
       }) do
    [
      trading_account_id,
      instrument_name,
      type |> Atom.to_string(),
      side |> Atom.to_string(),
      price,
      amount,
      time_in_force |> Atom.to_string()
    ]
    |> DB.query_val("SELECT process_trade_order($1, $2, $3, $4, $5, $6, $7, 0)")
  end
end
