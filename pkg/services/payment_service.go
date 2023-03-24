package services

import (
	"open-outcry/pkg/db"
	"open-outcry/pkg/models"
)

func CreatePaymentDeposit(appEntityId models.AppEntityId,
	amount float64,
	currency models.CurrencyName,
	reference string,
	details string,
) models.PaymentId {
	return CreatePaymentDepositCustomFee(appEntityId, amount, currency, reference, details, "DEPOSIT_FEE")
}

func CreatePaymentDepositCustomFee(appEntityId models.AppEntityId,
	amount float64,
	currency models.CurrencyName,
	reference string,
	details string,
	feeType any,
) models.PaymentId {
	var id string
	db.Instance().QueryRow(
		"SELECT process_payment('DEPOSIT', 'MASTER', $2, $3, $1, $4, $5, $6)",
		appEntityId, amount, currency, reference, details, feeType,
	).Scan(&id)
	return models.PaymentId(id)
}
