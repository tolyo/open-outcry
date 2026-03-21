package orderbook

import (
	"open-outcry/pkg/db"
	"open-outcry/pkg/models/instrument"
	order "open-outcry/pkg/models/order_book"
	tradeorder "open-outcry/pkg/models/trade_order"
)

func GetVolumeAtPrice(instrumentName instrument.InstrumentName, side tradeorder.OrderSide, price tradeorder.OrderPrice) float64 {
	res := db.QueryVal[float64](`
	 SELECT volume
	 FROM price_level
	 WHERE side = $2
	   AND instrument_id = (SELECT id FROM instrument WHERE name = $1)
	   AND price =  $3
   `, instrumentName, side, price)
	return res
}

func GetVolumes(instrumentName instrument.InstrumentName, side tradeorder.OrderSide) []order.PriceVolume {
	var orderBy string
	switch side {
	case tradeorder.Sell:
		orderBy = "ASC"
	case tradeorder.Buy:
		orderBy = "DESC"
	}
	res := db.QueryList[order.PriceVolume](`
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

func GetOrderBook(instrumentName instrument.InstrumentName) order.OrderBook {
	res := db.QueryList[order.PriceVolume](`
		SELECT price, volume, side
		FROM price_level
		WHERE price > 0
		AND instrument_id = (SELECT id FROM instrument WHERE name = $1)
		ORDER BY price ASC, side DESC
	`, instrumentName)

	book := order.OrderBook{
		SellSide: make([]order.PriceVolume, 0),
		BuySide:  make([]order.PriceVolume, 0),
	}
	for _, entry := range res {
		switch entry.Side {
		case tradeorder.Sell:
			book.SellSide = append(book.SellSide, order.PriceVolume{Price: entry.Price, Volume: entry.Volume})
		case tradeorder.Buy:
			book.BuySide = append(book.BuySide, order.PriceVolume{Price: entry.Price, Volume: entry.Volume})
		}
	}

	for i, j := 0, len(book.BuySide)-1; i < j; i, j = i+1, j-1 {
		book.BuySide[i], book.BuySide[j] = book.BuySide[j], book.BuySide[i]
	}

	return book
}
