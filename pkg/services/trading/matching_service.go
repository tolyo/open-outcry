package services

  // Main entry point for processing an market order.
  //  - For BUY side, the amount must be alocated in quote currency.
  //  - For SELL side the amount must be allocared in base currency
  
  @spec create(
          TradingAccountId(),
          InstrumentName(),
          TradeOrderType,
          TradeOrderSide,
          TradeOrder.amount(),
          OrderTimeInForce.t()
        )
  func Create(
    tradeingAccountId TradingAccountId, 
    instrument_name InstrumentName, :MARKET, side, amount, time_in_force) TradeOrderId {
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
  }

  @spec create(
          TradingAccount.id(),
          Instrument.name(),
          TradeOrder.Type.t(),
          TradeOrder.Side.t(),
          TradeOrder.price(),
          TradeOrder.amount(),
          OrderTimeInForce.t()
        ) TradeOrder.id()
  func create(trading_account_id, instrument_name, order_type, side, price, amount, time_in_force) {
    case order_type {
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
    }
  }

  func CancelTradeOrder(tradeOrderId TradeOrderId) {
    db.QueryVal("SELECT cancel_trade_order($1)", tradeOrderId)
  }

  func ProcessTradeOrder(tradeOrder TradeOrder) {
    db.QueryVal("SELECT process_trade_order($1, $2, $3, $4, $5, $6, $7, 0)", 
      tradeOrder.TraingAccountId,
      tradeOrder.InstrumentName,
      tradeOrder.Type,
      tradeOrder.Side,
      tradeOrder.Price,
      tradeOrder.Amount,
      tradeOrder.TimeInForce
    )
  }
}
