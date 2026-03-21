package transfer

import (
	"open-outcry/pkg/db"
	appentity "open-outcry/pkg/models/app_entity"
	currency "open-outcry/pkg/models/currency"
	transfermodel "open-outcry/pkg/models/transfer"

	log "github.com/sirupsen/logrus"
)

func CreateTransferDeposit(appEntityId appentity.AppEntityId,
	amount float64,
	currencyName currency.CurrencyName,
	reference string,
	details string,
) transfermodel.TransferJournalId {
	return CreateTransferDepositCustomFee(appEntityId, amount, currencyName, reference, details, "DEPOSIT_FEE")
}

func CreateTransferDepositCustomFee(appEntityId appentity.AppEntityId,
	amount float64,
	currencyName currency.CurrencyName,
	reference string,
	details string,
	feeType any,
) transfermodel.TransferJournalId {
	var id string
	err := db.Instance().QueryRow(
		"SELECT process_transfer('DEPOSIT', 'MASTER', $2, $3, $1, $4, $5, $6)",
		appEntityId, amount, currencyName, reference, details, feeType,
	).Scan(&id)
	if err != nil {
		log.Fatal(err)
	}
	return transfermodel.TransferJournalId(id)
}
