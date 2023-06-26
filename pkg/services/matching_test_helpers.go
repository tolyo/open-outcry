package services

import (
	"open-outcry/pkg/db"
	"open-outcry/pkg/models"

	log "github.com/sirupsen/logrus"
)

func GetSellBookOrderCount() int {
	return db.QueryVal[int](
		`
		SELECT COUNT(*)
		FROM trade_order t
		    INNER JOIN book_order b
			ON t.id = b.trade_order_id
		WHERE t.side = 'SELL'
		`,
	)
}
func GetBuyBookOrderCount() int {
	return db.QueryVal[int](
		`
		SELECT COUNT(*)
		FROM trade_order t
		    INNER JOIN book_order b
			ON t.id = b.trade_order_id
		WHERE t.side = 'BUY'
		`,
	)
}

func GetTradeCount() int {
	return db.QueryVal[int]("SELECT COUNT(*) FROM trade")
}

func GetTradePrices() []float64 {
	return db.QueryList[float64]("SELECT (price) FROM trade ORDER BY created_at ASC")
}

func GetCrossingLimitOrders(instrumentId int, side models.OrderSide, price models.OrderPrice) int {
	rows, err := db.Instance().Query("SELECT get_crossing_limit_orders($1, $2, $3, $4)",
		instrumentId,
		side,
		price,
		0,
	)

	if err != nil {
		log.Fatal(err)
	}
	var count int
	for rows.Next() {
		count++
	}
	return count
}

func GetAvailableLimitVolume(side models.OrderSide, price models.OrderPrice) float64 {
	return db.QueryVal[float64]("SELECT get_available_limit_volume(1, $1::order_side, $2)", side, price)
}
