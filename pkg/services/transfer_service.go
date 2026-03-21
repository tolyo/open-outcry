package services

import (
	"open-outcry/pkg/db"

	log "github.com/sirupsen/logrus"
)

func CreateTransferDeposit(appEntityId AppEntityId,
	amount float64,
	currency CurrencyName,
	reference string,
	details string,
) TransferJournalId {
	return CreateTransferDepositCustomFee(appEntityId, amount, currency, reference, details, "DEPOSIT_FEE")
}

func CreateTransferDepositCustomFee(appEntityId AppEntityId,
	amount float64,
	currency CurrencyName,
	reference string,
	details string,
	feeType any,
) TransferJournalId {
	var id string
	err := db.Instance().QueryRow(
		"SELECT process_transfer('DEPOSIT', 'MASTER', $2, $3, $1, $4, $5, $6)",
		appEntityId, amount, currency, reference, details, feeType,
	).Scan(&id)
	if err != nil {
		log.Fatal(err)
	}
	return TransferJournalId(id)
}
