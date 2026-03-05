package models

import (
	"open-outcry/pkg/db"

	"github.com/shopspring/decimal"
	log "github.com/sirupsen/logrus"
)

type TransferJournalId string

type TransferType string

const (
	Deposit        TransferType = "DEPOSIT"
	Withdrawal     TransferType = "WITHDRAWAL"
	Transfer       TransferType = "TRANSFER"
	InstrumentBuy  TransferType = "INSTRUMENT_BUY"
	InstrumentSell TransferType = "INSTRUMENT_SELL"
	Charge         TransferType = "CHARGE"
)

type TransferAmount decimal.Decimal

type TransferDetails string

type TransferExternalReferenceNumber string

// TransferEntry represents a transfer journal entry with its two ledger sides resolved.
type TransferEntry struct {
	Id                      string
	Type                    TransferType
	Amount                  float64
	Currency                CurrencyName
	SenderAccountId         CurrencyAccountId
	BeneficiaryAccountId    CurrencyAccountId
	Details                 TransferDetails
	ExternalReferenceNumber TransferExternalReferenceNumber
	Status                  string
	DebitBalanceAmount      float64
	CreditBalanceAmount     float64
}

const transferBaseQuery = `
	SELECT
	  t.pub_id,
	  t.type::text,
	  t.amount,
	  t.currency_name,
	  debit_ta.pub_id,
	  credit_ta.pub_id,
	  t.details,
	  COALESCE(t.external_reference_number, ''),
	  t.status,
	  debit_le.resulting_balance,
	  credit_le.resulting_balance
	FROM transfer t
	INNER JOIN transfer_ledger_entry debit_le ON debit_le.transfer_id = t.id AND debit_le.entry_type = 'DEBIT'
	INNER JOIN currency_account debit_ta ON debit_le.currency_account_id = debit_ta.id
	INNER JOIN transfer_ledger_entry credit_le ON credit_le.transfer_id = t.id AND credit_le.entry_type = 'CREDIT'
	INNER JOIN currency_account credit_ta ON credit_le.currency_account_id = credit_ta.id
`

func GetTransfer(id string) *TransferEntry {
	var t TransferEntry
	err := db.Instance().QueryRow(transferBaseQuery+`WHERE t.pub_id = $1`, id).Scan(
		&t.Id, &t.Type, &t.Amount, &t.Currency,
		&t.SenderAccountId, &t.BeneficiaryAccountId,
		&t.Details, &t.ExternalReferenceNumber,
		&t.Status, &t.DebitBalanceAmount, &t.CreditBalanceAmount,
	)
	if err != nil {
		log.Error(err)
		return nil
	}
	return &t
}

func GetTransfersByAppEntity(appEntityId AppEntityId) []TransferEntry {
	query := transferBaseQuery + `
		WHERE debit_ta.app_entity_id = (SELECT id FROM app_entity WHERE pub_id = $1)
		   OR credit_ta.app_entity_id = (SELECT id FROM app_entity WHERE pub_id = $1)
		ORDER BY t.created_at DESC
	`
	rows, err := db.Instance().Query(query, appEntityId)
	if err != nil {
		log.Error(err)
		return nil
	}
	defer rows.Close()

	var transfers []TransferEntry
	for rows.Next() {
		var t TransferEntry
		err := rows.Scan(
			&t.Id, &t.Type, &t.Amount, &t.Currency,
			&t.SenderAccountId, &t.BeneficiaryAccountId,
			&t.Details, &t.ExternalReferenceNumber,
			&t.Status, &t.DebitBalanceAmount, &t.CreditBalanceAmount,
		)
		if err != nil {
			log.Error(err)
			continue
		}
		transfers = append(transfers, t)
	}
	return transfers
}

// TransferLedgerEntry represents one side of a double-entry monetary transfer.
type TransferLedgerEntry struct {
	Id                string
	TransferId        string
	CurrencyAccountId CurrencyAccountId
	EntryType         string
	Amount            decimal.Decimal
	ResultingBalance  decimal.Decimal
	CreatedAt         string
}

// GetTransferLedgerEntries returns all ledger entries for a given transfer.
func GetTransferLedgerEntries(transferPubId string) []TransferLedgerEntry {
	rows, err := db.Instance().Query(`
		SELECT
			tle.pub_id,
			t.pub_id,
			ta.pub_id,
			tle.entry_type::text,
			tle.amount,
			tle.resulting_balance,
			tle.created_at
		FROM transfer_ledger_entry tle
		INNER JOIN transfer t ON tle.transfer_id = t.id
		INNER JOIN currency_account ta ON tle.currency_account_id = ta.id
		WHERE t.pub_id = $1
		ORDER BY tle.id
	`, transferPubId)
	if err != nil {
		log.Error(err)
		return nil
	}
	defer rows.Close()

	var entries []TransferLedgerEntry
	for rows.Next() {
		var e TransferLedgerEntry
		err := rows.Scan(
			&e.Id, &e.TransferId, &e.CurrencyAccountId,
			&e.EntryType, &e.Amount, &e.ResultingBalance,
			&e.CreatedAt,
		)
		if err != nil {
			log.Error(err)
			continue
		}
		entries = append(entries, e)
	}
	return entries
}

// GetTransferLedgerEntriesByAccount returns all ledger entries for a given currency account.
func GetTransferLedgerEntriesByAccount(currencyAccountPubId CurrencyAccountId) []TransferLedgerEntry {
	rows, err := db.Instance().Query(`
		SELECT
			tle.pub_id,
			t.pub_id,
			ta.pub_id,
			tle.entry_type::text,
			tle.amount,
			tle.resulting_balance,
			tle.created_at
		FROM transfer_ledger_entry tle
		INNER JOIN transfer t ON tle.transfer_id = t.id
		INNER JOIN currency_account ta ON tle.currency_account_id = ta.id
		WHERE ta.pub_id = $1
		ORDER BY tle.id
	`, currencyAccountPubId)
	if err != nil {
		log.Error(err)
		return nil
	}
	defer rows.Close()

	var entries []TransferLedgerEntry
	for rows.Next() {
		var e TransferLedgerEntry
		err := rows.Scan(
			&e.Id, &e.TransferId, &e.CurrencyAccountId,
			&e.EntryType, &e.Amount, &e.ResultingBalance,
			&e.CreatedAt,
		)
		if err != nil {
			log.Error(err)
			continue
		}
		entries = append(entries, e)
	}
	return entries
}
