package matching

import (
	"open-outcry/pkg/db"
	tradeorder "open-outcry/pkg/models/trade_order"
	"open-outcry/pkg/utils"
)

func GetBookOrderCount(side tradeorder.OrderSide) int {
	return db.QueryVal[int](utils.Format(
		`
		SELECT COUNT(*)
		FROM trade_order t
		    INNER JOIN book_order b
			ON t.id = b.trade_order_id
		WHERE t.side = '{{.}}'
		`, side,
	))
}

func GetSellBookOrderCount() int {
	return GetBookOrderCount(tradeorder.Sell)
}

func GetBuyBookOrderCount() int {
	return GetBookOrderCount(tradeorder.Buy)
}

func GetTradeCount() int {
	return db.QueryVal[int]("SELECT COUNT(*) FROM trade")
}

func GetTradePrices() []float64 {
	return db.QueryList[float64]("SELECT (price) FROM trade ORDER BY created_at ASC")
}
