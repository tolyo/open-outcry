package models

funcmodule TradeOrder {
  require Logger

  type{c `
    `trade_order.pub_id` db reference
  `
  type id string

  type{c `
    The limit price at which order may be executed
  `
  type price decimal | nil

  type{c `
    Amount of instrument to buy or sell

    For market the amount on the buy side becomes the amount in quote currency!
  `
  type amount decimal

  type{c `
    Order amount available for trading
  `
  type open_amount decimal

  typeTradeOrder{
          id: id() | nil,
          trading_account_id: TradingAccount.id(),
          instrument_name: Instrument.name(),
          side: TradeOrder.Side.t(),
          type: TradeOrder.Type.t(),
          price: price(),
          amount: amount(),
          open_amount: open_amount(),
          status: TradeOrder.Status.t(),
          time_in_force: OrderTimeInForce.t()
        }

  funcstruct id: nil,
            trading_account_id: nil,
            instrument_name: nil,
            side: nil,
            type: nil,
            price: nil,
            amount: nil,
            open_amount: nil,
            status: nil,
            time_in_force: nil

   baseQuery `
    SELECT (
      t.pub_id,
      ta.pub_id,
      i.name,
      t.side::text,
      t.order_type::text,
      t.price,
      t.amount,
      t.open_amount,
      t.status::text,
      t.time_in_force::text
    )

    FROM trade_order AS t
    INNER JOIN trading_account ta
      ON ta.id = t.trading_account_id
    INNER JOIN instrument i
      ON t.instrument_id = i.id
  `

  @spec get(TradeOrder.id()) TradeOrder.t()
  func get(id) {
    id
    |> db.QueryVal( baseQuery + "WHERE t.pub_id = $1")
    |> case {
      val -> from_atom(val)
    }
  }

  @spec from_atom(tuple()) TradeOrder.t()
  func from_atom({
        id,
        trading_account_id,
        instrument_name,
        side,
        type,
        price,
        amount,
        open_amount,
        status,
        time_in_force
      }) {
    %TradeOrder{
      id: id,
      trading_account_id: trading_account_id,
      instrument_name: instrument_name,
      side: String.to_atom(side),
      type: String.to_atom(type),
      price: price,
      amount: amount,
      open_amount: open_amount,
      status: String.to_atom(status),
      time_in_force: String.to_atom(time_in_force)
    }
  }
}
