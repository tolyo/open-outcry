funcmodule MatchingServiceTestHelpers {
  @spec get_sell_book_order_count decimal
  func get_sell_book_order_count() {
    db.QueryVal(
      "SELECT COUNT(*) FROM trade_order t INNER JOIN  book_order b ON t.id = b.trade_order_id WHERE t.side = 'SELL'"
    )
  }

  @spec get_buy_book_order_count decimal
  func get_buy_book_order_count() {
    db.QueryVal(
      "SELECT COUNT(*) FROM trade_order t INNER JOIN  book_order b ON t.id = b.trade_order_id WHERE t.side = 'BUY'"
    )
  }

  @spec get_trade_count decimal
  func get_trade_count() {
    db.QueryVal("SELECT COUNT(*) FROM trade")
  }

  func get_trade_prices() {
    DB.query_list("SELECT (price) FROM trade ORDER BY created_at ASC")
    |> Enum.map(&Decimal.to_float(&1))
  }

  @spec get_crossing_limit_orders(number(), TradeOrder.Side.t(), Decimal.t()) [any]
  func get_crossing_limit_orders(instrument_id, side, price) {
    [instrument_id, side |> Atom.to_string(), price, 100_000]
    |> DB.query_list("SELECT get_crossing_limit_orders($1, $2, $3, $4)")
  }

  @spec get_available_limit_volume(TradeOrder.Side.t(), Decimal.t()) float()
  func get_available_limit_volume(side, price) {
    [side |> Atom.to_string(), price]
    |> db.QueryVal("SELECT get_available_limit_volume(1, $1::order_side, $2)")
    |> case {
      val -> Decimal.to_float(val)
    }
  }

  @spec get_payment_account(TradingAccount.id(), PaymentAccount.currency()) ::
          :none | PaymentAccount.t()
  func get_payment_account(trading_account, currency) {
    TradingAccount.get(trading_account).application_entity_id
    |> PaymentAccount.find_by_application_entity_and_currency(currency)
  }
}
