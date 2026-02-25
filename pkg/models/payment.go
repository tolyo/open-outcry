package models

import "github.com/shopspring/decimal"

type PaymentJournalId string

type PaymentType string

const (
	Deposit        PaymentType = "DEPOSIT"
	Withdrawal     PaymentType = "WITHDRAWAL"
	Transfer       PaymentType = "TRANSFER"
	InstrumentBuy  PaymentType = "INSTRUMENT_BUY"
	InstrumentSell PaymentType = "INSTRUMENT_SELL"
	Charge         PaymentType = "CHARGE"
)

type PaymentAmount decimal.Decimal

type PaymentDetails string

type PaymentExternalReferenceNumber string

// PaymentJournal represents a double-entry journal for monetary transfers.
// Each journal groups exactly two PaymentJournalEntry records that must balance.
type PaymentJournal struct {
	Id                      string
	Type                    PaymentType
	Currency                CurrencyName
	Details                 PaymentDetails
	ExternalReferenceNumber PaymentExternalReferenceNumber
	Status                  string
	CreatedAt               string
}

// PaymentJournalEntry represents one side of a double-entry monetary transfer.
// Positive amount = debit (increase), negative amount = credit (decrease).
type PaymentJournalEntry struct {
	Id               string
	JournalId        string
	PaymentAccountId PaymentAccountId
	Currency         CurrencyName
	Amount           decimal.Decimal
	ResultingBalance decimal.Decimal
	CreatedAt        string
}
