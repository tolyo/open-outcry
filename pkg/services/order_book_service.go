package services

import (
	"open-outcry/pkg/db"
)

func GetVolumeAtPrice(instrumentName InstrumentName, side OrderSide, price OrderPrice) float64 {
	res := db.QueryVal[float64](`
	 SELECT volume
	 FROM price_level
	 WHERE side = $2
	   AND instrument_id = (SELECT id FROM instrument WHERE name = $1)
	   AND price =  $3
   `, instrumentName, side, price)
	return res
}

func GetVolumes(instrumentName InstrumentName, side OrderSide) []PriceVolume {
	var orderBy string
	switch side {
	case Sell:
		orderBy = "ASC"
	case Buy:
		orderBy = "DESC"
	}
	res := db.QueryList[PriceVolume](`
		SELECT price, volume
		FROM price_level
		WHERE side = $2
		AND price > 0
		AND instrument_id = (SELECT id FROM instrument WHERE name = $1)
		ORDER BY price `+orderBy,
		instrumentName,
		side,
	)
	return res
}

func GetOrderBook(instrumentName InstrumentName) OrderBook {
	res := db.QueryList[PriceVolume](`
		SELECT price, volume, side
		FROM price_level
		WHERE price > 0
		AND instrument_id = (SELECT id FROM instrument WHERE name = $1)
		ORDER BY price ASC, side DESC
	`, instrumentName)

	orderBook := OrderBook{
		SellSide: make([]PriceVolume, 0),
		BuySide:  make([]PriceVolume, 0),
	}
	for _, entry := range res {
		switch entry.Side {
		case Sell:
			orderBook.SellSide = append(orderBook.SellSide, PriceVolume{Price: entry.Price, Volume: entry.Volume})
		case Buy:
			orderBook.BuySide = append(orderBook.BuySide, PriceVolume{Price: entry.Price, Volume: entry.Volume})
		}
	}

	// reverse array
	for i, j := 0, len(orderBook.BuySide)-1; i < j; i, j = i+1, j-1 {
		orderBook.BuySide[i], orderBook.BuySide[j] = orderBook.BuySide[j], orderBook.BuySide[i]
	}

	return orderBook
}
