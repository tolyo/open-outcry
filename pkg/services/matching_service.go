package services

import (
	"context"
	"database/sql"
	"open-outcry/pkg/db"
	"open-outcry/pkg/models"

	log "github.com/sirupsen/logrus"
)

// ProcessTradeOrder Main entry point for processing a trade order.
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
	tx, err := db.Instance().BeginTx(context.Background(), &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return "", err
	}

	err = tx.QueryRow("SELECT process_trade_order($1, $2, $3, $4, $5, $6, $7, 0)",
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
		tx.Rollback()
		return "", err
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return "", err
	}

	return models.TradeOrderId(tradeOrderId), nil
}

func CancelTradeOrder(tradeOrderId models.TradeOrderId) error {
	tx, err := db.Instance().BeginTx(context.Background(), &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return err
	}

	res, err := tx.Exec("SELECT cancel_trade_order($1)", tradeOrderId)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return err
	}

	count, _ := res.RowsAffected()
	if count != 1 {
		log.Fatal(count)
	}
	return nil
}
