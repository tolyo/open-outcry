package orderbook

import tradeorder "open-outcry/pkg/models/trade_order"

type PriceVolume struct {
	Price  float64
	Volume float64
	Side   tradeorder.OrderSide
}

type OrderBook struct {
	SellSide []PriceVolume
	BuySide  []PriceVolume
}
