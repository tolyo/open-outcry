package services

import (
	"open-outcry/pkg/db"
	"open-outcry/pkg/models"

	log "github.com/sirupsen/logrus"
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
	err := db.Instance().QueryRow(
		"SELECT process_payment('DEPOSIT', 'MASTER', $2, $3, $1, $4, $5, $6)",
		appEntityId, amount, currency, reference, details, feeType,
	).Scan(&id)
	if err != nil {
		log.Fatal(err)
	}
	return models.PaymentId(id)
}
