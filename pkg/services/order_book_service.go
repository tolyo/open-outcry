package services

import (
	"open-outcry/pkg/db"
	"open-outcry/pkg/models"
)

func GetVolumeAtPrice(instrumentName models.InstrumentName, side models.OrderSide, price models.OrderPrice) float64 {
	res := db.QueryVal[float64](`
	 SELECT volume
	 FROM price_level
	 WHERE side = $2
	   AND instrument_id = (SELECT id FROM instrument WHERE name = $1)
	   AND price =  $3
   `, instrumentName, side, price)
	return res
}

func GetVolumes(instrumentName models.InstrumentName, side models.OrderSide) []models.PriceVolume {
	var orderBy string
	switch side {
	case models.Sell:
		orderBy = "ASC"
	case models.Buy:
		orderBy = "DESC"
	}
	res := db.QueryList[models.PriceVolume](`
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

func GetOrderBook(instrumentName models.InstrumentName) models.OrderBook {
	res := db.QueryList[models.PriceVolume](`
		SELECT price, volume, side
		FROM price_level
		WHERE price > 0
		AND instrument_id = (SELECT id FROM instrument WHERE name = $1)
		ORDER BY price ASC, side DESC
	`, instrumentName)

	orderBook := models.OrderBook{}
	for _, entry := range res {
		switch entry.Side {
		case models.Sell:
			orderBook.SellSide = append(orderBook.SellSide, models.PriceVolume{Price: entry.Price, Volume: entry.Volume})
		case models.Buy:
			orderBook.BuySide = append(orderBook.BuySide, models.PriceVolume{Price: entry.Price, Volume: entry.Volume})
		}
	}

	// reverse array
	for i, j := 0, len(orderBook.BuySide)-1; i < j; i, j = i+1, j-1 {
		orderBook.BuySide[i], orderBook.BuySide[j] = orderBook.BuySide[j], orderBook.BuySide[i]
	}

	return orderBook
}
