package models

import (
	"log"
	"open-outcry/pkg/db"

	"github.com/shopspring/decimal"
)

// InstrumentAccountId `instrument_account.pub_id` db reference
type InstrumentAccountId string

type InstrumentAccountHolding struct {
	Amount          decimal.Decimal
	AmountAvailable float64
	AmountReserved  float64
	Name            InstrumentName
	Value           float64
	Currency        CurrencyName
}

type InstrumentAccount struct {
	Id          InstrumentAccountId
	AppEntityId AppEntityExternalId
}

const instrumentAccountBaseQuery = `
    SELECT t.pub_id, ae.pub_id
    FROM instrument_account AS t
    INNER JOIN app_entity ae
          ON ae.id = t.app_entity_id
`

func GetInstrumentAccount(id InstrumentAccountId) *InstrumentAccount {
	return helper(instrumentAccountBaseQuery+"WHERE t.pub_id = $1", id)
}

func FindInstrumentAccountByApplicationEntityId(appEntityId AppEntityId) *InstrumentAccount {
	return helper(instrumentAccountBaseQuery+"WHERE ae.pub_id = $1", appEntityId)
}

func helper(query string, arg any) *InstrumentAccount {
	var instrumentAccount InstrumentAccount
	err := db.Instance().QueryRow(
		query, arg,
	).Scan(&instrumentAccount.Id, &instrumentAccount.AppEntityId)
	if err != nil {
		log.Fatal(err)
	}

	return &instrumentAccount
}

func GetInstrumentAccountHoldings(instrumentAccountId InstrumentAccountId) []InstrumentAccountHolding {
	query := `
		SELECT
		  tai.amount,
		  tai.amount - tai.amount_reserved,
		  tai.amount_reserved,
		  i.name,
		  0,
		  i.quote_currency
		FROM instrument_account_holding tai
		INNER JOIN instrument i ON tai.instrument_id = i.id
		INNER JOIN instrument_account ta ON tai.instrument_account = ta.id
		WHERE ta.pub_id = $1
	`
	rows, err := db.Instance().Query(query, instrumentAccountId)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	instruments := make([]InstrumentAccountHolding, 0)
	for rows.Next() {
		var inst InstrumentAccountHolding
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
