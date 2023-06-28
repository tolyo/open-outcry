package models

import (
	"open-outcry/pkg/db"

	log "github.com/sirupsen/logrus"
)

// `payment_account.pub_id` db reference
type PaymentAccountId string

type PaymentAccount struct {
	Id              PaymentAccountId
	AppEntityId     AppEntityId
	Amount          float64
	AmountAvailable float64
	AmountReserved  float64
	Currency        CurrencyName
}

const basePaymentAccountQuery = `
	SELECT
	  pa.pub_id,
	  ae.pub_id,
	  pa.amount,
	  pa.amount_reserved,
	  pa.amount - pa.amount_reserved,
	  c.name

	FROM payment_account AS pa

	INNER JOIN app_entity ae
	  ON pa.app_entity_id = ae.id

	INNER JOIN currency c
	  ON pa.currency_name = c.name
  `

func GetPaymentAccount(id PaymentAccountId) *PaymentAccount {
	var res PaymentAccount
	err := db.Instance().QueryRow(basePaymentAccountQuery+` WHERE pa.pub_id = $1`, id).Scan(
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

func FindPaymentAccountByAppEntityIdAndCurrencyName(
	appEntityId AppEntityId,
	currencyName CurrencyName,
) *PaymentAccount {
	var paymentAccount PaymentAccount
	db.Instance().QueryRow(
		basePaymentAccountQuery+`WHERE ae.pub_id = $1 AND c.name = $2`,
		appEntityId,
		currencyName,
	).Scan(
		&paymentAccount.Id,
		&paymentAccount.AppEntityId,
		&paymentAccount.Amount,
		&paymentAccount.AmountReserved,
		&paymentAccount.AmountAvailable,
		&paymentAccount.Currency,
	)
	return &paymentAccount
}

func CreatePaymentAccount(appEntityId AppEntityId, currencyName CurrencyName) PaymentAccountId {
	var id string
	db.Instance().QueryRow("SELECT create_payment_account($1, $2)", appEntityId, currencyName).Scan(&id)
	return PaymentAccountId(id)
}
