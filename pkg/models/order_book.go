package models

type PriceVolume struct {
	Price  float64
	Volume float64
	Side   OrderSide
}

type OrderBook struct {
	SellSide []PriceVolume
	BuySide  []PriceVolume
}
