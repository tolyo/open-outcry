defmodule Trade do
  @typedoc """
    `trade.pub_id` db reference
  """
  @type id :: String.t()

  @type t :: %Trade{
          id: id(),
          instrument_id: Instrument.id(),
          price: number(),
          amount: number(),
          seller_order_id: TradeOrder.id() | nil,
          buyer_order_id: TradeOrder.id() | nil,
          taker_order_id: TradeOrder.id() | nil,
          created_at: String.t()
        }

  defstruct id: nil,
            instrument_id: nil,
            price: nil,
            amount: nil,
            seller_order_id: nil,
            buyer_order_id: nil,
            taker_order_id: nil,
            created_at: nil
end
