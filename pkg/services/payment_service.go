package services

import (
	"open-outcry/pkg/db"
	"open-outcry/pkg/models"
)

func CreatePaymentDeposit(appEntityId models.AppEntityId,
	amount float64,
	currency models.CurrencyName,
	reference string,
	details string) models.PaymentId {
	var id string
	db.Instance().QueryRow(
		"SELECT process_payment('DEPOSIT', 'MASTER', $2, $3, $1, $4, $5)",
		appEntityId, amount, currency, reference, details,
	).Scan(&id)
	return models.PaymentId(id)
}
