defmodule OrderBookService do
  @spec get_volume_at_price(Instrument.name(), TradeOrder.Side.t(), Decimal.t()) :: Decimal.t()
  def get_volume_at_price(instument_name, side, price) do
    [instument_name, side |> Atom.to_string(), price]
    |> DB.query_val("""
      SELECT volume
      FROM price_level
      WHERE side = $2
        AND instrument_id = (SELECT id FROM instrument WHERE name = $1)
        AND price =  $3
    """)
    |> case do
      nil -> 0.0
      val -> val
    end
  end

  @spec get_volumes(Instrument.name(), TradeOrder.Side.t()) :: [
          {PriceLevel.price(), PriceLevel.volume()}
        ]
  def get_volumes(instrument_name, side) do
    [instrument_name, side |> Atom.to_string()]
    |> DB.query_list("""
      SELECT (price, volume)
      FROM price_level
      WHERE side = $2
      AND price > 0
      AND instrument_id = (SELECT id FROM instrument WHERE name = $1)
      ORDER BY price #{case side do
      :SELL -> "ASC"
      :BUY -> "DESC"
    end}
    """)
    |> Enum.map(fn {a, b} -> {Decimal.to_float(a), Decimal.to_float(b)} end)
  end
end
