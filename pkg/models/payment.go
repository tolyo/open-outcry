package models

import "github.com/shopspring/decimal"

type PaymentId string

type PaymentType string

const (
	Deposit        PaymentType = "DEPOSIT"
	Withdrawal     PaymentType = "WITHDRAWAL"
	Transfer       PaymentType = "TRANSFER"
	InstrumentBuy  PaymentType = "INSTRUMENT_BUY"
	InstrumentSell PaymentType = "INSTRUMENT_SELL"
)

type PaymentAmount decimal.Decimal

type PaymentDetails string

type PaymentExternaReferenceNumber string

type Payment struct {
	Id                      string
	Number                  string
	Type                    string
	Amount                  string
	Currency                CurrencyName
	SenderAccountId         string
	BeneficiaryAccountId    string
	Details                 string
	ExternalReferenceNumber string
	FeeSender               string
	FeeBeneficiary          string
	Status                  string
	TotalAmount             string
	DebitBalanceAmount      string
	CreditBalanceAmount     string
}
