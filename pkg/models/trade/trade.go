package trade

import (
	"open-outcry/pkg/db"
	"open-outcry/pkg/models/instrument"
	instrumentaccount "open-outcry/pkg/models/instrument_account"
	tradeorder "open-outcry/pkg/models/trade_order"

	log "github.com/sirupsen/logrus"
)

// Trade represents an executed trade
type Trade struct {
	Id             string
	InstrumentName instrument.InstrumentName
	Price          float64
	Amount         float64
	SellerOrderId  tradeorder.TradeOrderId
	BuyerOrderId   tradeorder.TradeOrderId
	Created        string
}

const tradeBaseQuery = `
	SELECT
	  tr.pub_id,
	  i.name,
	  tr.price,
	  tr.amount,
	  so.pub_id,
	  bo.pub_id,
	  tr.created_at
	FROM trade tr
	INNER JOIN instrument i ON tr.instrument_id = i.id
	INNER JOIN trade_order so ON tr.seller_order_id = so.id
	INNER JOIN trade_order bo ON tr.buyer_order_id = bo.id
`

func GetTrade(id string) *Trade {
	var t Trade
	err := db.Instance().QueryRow(tradeBaseQuery+`WHERE tr.pub_id = $1`, id).Scan(
		&t.Id, &t.InstrumentName, &t.Price, &t.Amount,
		&t.SellerOrderId, &t.BuyerOrderId, &t.Created,
	)
	if err != nil {
		log.Error(err)
		return nil
	}
	return &t
}

func GetTradesByInstrumentAccount(instrumentAccountId instrumentaccount.InstrumentAccountId) []Trade {
	query := tradeBaseQuery + `
		WHERE so.instrument_account_id = (SELECT id FROM instrument_account WHERE pub_id = $1)
		   OR bo.instrument_account_id = (SELECT id FROM instrument_account WHERE pub_id = $1)
		ORDER BY tr.created_at DESC
	`
	rows, err := db.Instance().Query(query, instrumentAccountId)
	if err != nil {
		log.Error(err)
		return nil
	}
	defer rows.Close()

	var trades []Trade
	for rows.Next() {
		var t Trade
		err := rows.Scan(
			&t.Id, &t.InstrumentName, &t.Price, &t.Amount,
			&t.SellerOrderId, &t.BuyerOrderId, &t.Created,
		)
		if err != nil {
			log.Error(err)
			continue
		}
		trades = append(trades, t)
	}
	return trades
}
