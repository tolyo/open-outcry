package models

// `trade.pub_id` db reference

  `
  type id string

  type Trade struct{
          id: id()
          instrument_id: Instrument.id()
          price: decimal,
          amount: decimal,
          seller_order_id: TradeOrder.id() | nil,
          buyer_order_id: TradeOrder.id() | nil,
          taker_order_id: TradeOrder.id() | nil,
          created_at: String.t()
        }

  funcstruct id: nil,
            instrument_id: nil,
            price: nil,
            amount: nil,
            seller_order_id: nil,
            buyer_order_id: nil,
            taker_order_id: nil,
            created_at: nil
}
