package models

type PriceVolume struct {
	Price  float64
	Volume float64
}

type OrderBook struct {
	SellSide []PriceVolume
	BuySide  []PriceVolume
}
