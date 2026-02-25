package models

import (
	"log"
	"open-outcry/pkg/db"

	"github.com/shopspring/decimal"
)

// TradingAccountId `trading_account.pub_id` db reference
type TradingAccountId string

type TradingAccountInstrument struct {
	Amount          decimal.Decimal
	AmountAvailable float64
	AmountReserved  float64
	Name            InstrumentName
	Value           float64
	Currency        CurrencyName
}

type TradingAccount struct {
	Id          TradingAccountId
	AppEntityId AppEntityExternalId
}

const tradingAccountBaseQuery = `
    SELECT t.pub_id, ae.pub_id
    FROM trading_account AS t
    INNER JOIN app_entity ae
          ON ae.id = t.app_entity_id
`

func GetTradingAccount(id TradingAccountId) *TradingAccount {
	return helper(tradingAccountBaseQuery+"WHERE t.pub_id = $1", id)
}

func FindTradingAccountByApplicationEntityId(appEntityId AppEntityId) *TradingAccount {
	return helper(tradingAccountBaseQuery+"WHERE ae.pub_id = $1", appEntityId)
}

func helper(query string, arg any) *TradingAccount {
	var tradingAccount TradingAccount
	err := db.Instance().QueryRow(
		query, arg,
	).Scan(&tradingAccount.Id, &tradingAccount.AppEntityId)
	if err != nil {
		log.Fatal(err)
	}

	return &tradingAccount
}

func GetTradingAccountInstruments(tradingAccountId TradingAccountId) []TradingAccountInstrument {
	query := `
		SELECT
		  tai.amount,
		  tai.amount - tai.amount_reserved,
		  tai.amount_reserved,
		  i.name,
		  0,
		  i.quote_currency
		FROM trading_account_instrument tai
		INNER JOIN instrument i ON tai.instrument_id = i.id
		INNER JOIN trading_account ta ON tai.trading_account = ta.id
		WHERE ta.pub_id = $1
	`
	rows, err := db.Instance().Query(query, tradingAccountId)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	instruments := make([]TradingAccountInstrument, 0)
	for rows.Next() {
		var inst TradingAccountInstrument
		err := rows.Scan(
			&inst.Amount, &inst.AmountAvailable, &inst.AmountReserved,
			&inst.Name, &inst.Value, &inst.Currency,
		)
		if err != nil {
			log.Fatal(err)
		}
		instruments = append(instruments, inst)
	}
	return instruments
}
