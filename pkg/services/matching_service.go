package services

import (
	log "github.com/sirupsen/logrus"
	"open-outcry/pkg/db"
	"open-outcry/pkg/models"
)

// Main entry point for processing an market order.
//   - For BUY side, the amount must be alocated in quote currency.
//   - For SELL side the amount must be allocared in base currency
func ProcessTradeOrder(
	tradingAccountId models.TradingAccountId,
	instrumentName models.InstrumentName,
	orderType models.OrderType,
	side models.OrderSide,
	price models.OrderPrice,
	amount float64,
	timeInForce models.OrderTimeInForce,
) (models.TradeOrderId, error) {
	var tradeOrderId string
	err := db.Instance().QueryRow("SELECT process_trade_order($1, $2, $3, $4, $5, $6, $7, 0)",
		tradingAccountId,
		instrumentName,
		orderType,
		side,
		price,
		amount,
		timeInForce,
	).Scan(&tradeOrderId)

	if err != nil {
		log.Error(err)
		return "", err
	}
	return models.TradeOrderId(tradeOrderId), nil
}

func CancelTradeOrder(tradeOrderId models.TradeOrderId) {
	res, err := db.Instance().Exec("SELECT cancel_trade_order($1)", tradeOrderId)
	if err != nil {
		log.Fatal(err)
	}
	count, _ := res.RowsAffected()
	if count != 1 {
		log.Fatal(count)
	}
	return
}
