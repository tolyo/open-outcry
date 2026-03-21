package orderbook

import (
	"open-outcry/pkg/db"
	tradeorder "open-outcry/pkg/models/trade_order"

	log "github.com/sirupsen/logrus"
)

func GetCrossingLimitOrders(instrumentId int, side tradeorder.OrderSide, price tradeorder.OrderPrice) int {
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

func GetAvailableLimitVolume(side tradeorder.OrderSide, price tradeorder.OrderPrice) float64 {
	return db.QueryVal[float64]("SELECT get_available_limit_volume(1, $1::order_side, $2)", side, price)
}
