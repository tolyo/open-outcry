defmodule MatchingServiceTestHelpers do
  @spec get_sell_book_order_count :: number()
  def get_sell_book_order_count() do
    DB.query_val(
      "SELECT COUNT(*) FROM trade_order t INNER JOIN  book_order b ON t.id = b.trade_order_id WHERE t.side = 'SELL'"
    )
  end

  @spec get_buy_book_order_count :: number()
  def get_buy_book_order_count() do
    DB.query_val(
      "SELECT COUNT(*) FROM trade_order t INNER JOIN  book_order b ON t.id = b.trade_order_id WHERE t.side = 'BUY'"
    )
  end

  @spec get_trade_count :: number()
  def get_trade_count() do
    DB.query_val("SELECT COUNT(*) FROM trade")
  end

  def get_trade_prices() do
    DB.query_list("SELECT (price) FROM trade ORDER BY created_at ASC")
    |> Enum.map(&Decimal.to_float(&1))
  end

  @spec get_crossing_limit_orders(number(), TradeOrder.Side.t(), Decimal.t()) :: [any]
  def get_crossing_limit_orders(instrument_id, side, price) do
    [instrument_id, side |> Atom.to_string(), price, 100_000]
    |> DB.query_list("SELECT get_crossing_limit_orders($1, $2, $3, $4)")
  end

  @spec get_available_limit_volume(TradeOrder.Side.t(), Decimal.t()) :: float()
  def get_available_limit_volume(side, price) do
    [side |> Atom.to_string(), price]
    |> DB.query_val("SELECT get_available_limit_volume(1, $1::order_side, $2)")
    |> case do
      val -> Decimal.to_float(val)
    end
  end

  @spec get_payment_account(TradingAccount.id(), PaymentAccount.currency()) ::
          :none | PaymentAccount.t()
  def get_payment_account(trading_account, currency) do
    TradingAccount.get(trading_account).application_entity_id
    |> PaymentAccount.find_by_application_entity_and_currency(currency)
  end
end
