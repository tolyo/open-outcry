package models

import (
	"open-outcry/pkg/db"

	log "github.com/sirupsen/logrus"
)

// CurrencyAccountId `currency_account.pub_id` db reference
type CurrencyAccountId string

type CurrencyAccount struct {
	Id              CurrencyAccountId
	AppEntityId     AppEntityId
	Amount          float64
	AmountAvailable float64
	AmountReserved  float64
	Currency        CurrencyName
}

const baseCurrencyAccountQuery = `
	SELECT
	  ta.pub_id,
	  ae.pub_id,
	  ta.amount,
	  ta.amount_reserved,
	  ta.amount - ta.amount_reserved,
	  c.name

	FROM currency_account AS ta

	INNER JOIN app_entity ae
	  ON ta.app_entity_id = ae.id

	INNER JOIN currency c
	  ON ta.currency_name = c.name
  `

func GetCurrencyAccount(id CurrencyAccountId) *CurrencyAccount {
	var res CurrencyAccount
	err := db.Instance().QueryRow(baseCurrencyAccountQuery+` WHERE ta.pub_id = $1`, id).Scan(
		&res.Id,
		&res.AppEntityId,
		&res.Amount,
		&res.AmountReserved,
		&res.AmountAvailable,
		&res.Currency,
	)
	if err != nil {
		log.Error(err)
	}
	return &res
}

func FindCurrencyAccountByAppEntityIdAndCurrencyName(
	appEntityId AppEntityId,
	currencyName CurrencyName,
) *CurrencyAccount {
	var currencyAccount CurrencyAccount
	err := db.Instance().QueryRow(
		baseCurrencyAccountQuery+`WHERE ae.pub_id = $1 AND c.name = $2`,
		appEntityId,
		currencyName,
	).Scan(
		&currencyAccount.Id,
		&currencyAccount.AppEntityId,
		&currencyAccount.Amount,
		&currencyAccount.AmountReserved,
		&currencyAccount.AmountAvailable,
		&currencyAccount.Currency,
	)
	if err != nil {
		log.Fatal(err)
	}
	return &currencyAccount
}

func CreateCurrencyAccount(appEntityId AppEntityId, currencyName CurrencyName) CurrencyAccountId {
	var id string
	err := db.Instance().QueryRow("SELECT create_currency_account($1, $2)", appEntityId, currencyName).Scan(&id)
	if err != nil {
		log.Fatal(err)
	}
	return CurrencyAccountId(id)
}

func GetCurrencyAccountsByAppEntity(appEntityId AppEntityId) []CurrencyAccount {
	query := baseCurrencyAccountQuery + `WHERE ae.pub_id = $1`
	rows, err := db.Instance().Query(query, appEntityId)
	if err != nil {
		log.Error(err)
		return nil
	}
	defer rows.Close()

	var accounts []CurrencyAccount
	for rows.Next() {
		var a CurrencyAccount
		err := rows.Scan(
			&a.Id, &a.AppEntityId, &a.Amount,
			&a.AmountReserved, &a.AmountAvailable, &a.Currency,
		)
		if err != nil {
			log.Error(err)
			continue
		}
		accounts = append(accounts, a)
	}
	return accounts
}
