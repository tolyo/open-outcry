package tradeorder

import (
	"context"
	"database/sql"

	"open-outcry/pkg/db"
	instrument "open-outcry/pkg/models/instrument"
	instrumentaccount "open-outcry/pkg/models/instrument_account"
	tradeorder "open-outcry/pkg/models/trade_order"

	log "github.com/sirupsen/logrus"
)

// ProcessTradeOrder Main entry point for processing a trade order.
//   - For BUY side, the amount must be allocated in quote currency.
//   - For SELL side the amount must be allocated in base currency
func ProcessTradeOrder(
	instrumentAccountId instrumentaccount.InstrumentAccountId,
	instrumentName instrument.InstrumentName,
	orderType tradeorder.OrderType,
	side tradeorder.OrderSide,
	price tradeorder.OrderPrice,
	amount float64,
	timeInForce tradeorder.OrderTimeInForce,
) (tradeorder.TradeOrderId, error) {
	var tradeOrderId string
	tx, err := db.Instance().BeginTx(context.Background(), &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return "", err
	}

	err = tx.QueryRow("SELECT process_trade_order($1, $2, $3, $4, $5, $6, $7, 0)",
		instrumentAccountId,
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

	return tradeorder.TradeOrderId(tradeOrderId), nil
}

func CancelTradeOrder(tradeOrderId tradeorder.TradeOrderId) error {
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
