package models

import "github.com/shopspring/decimal"

type FeeType string

const (
	DepositFee    FeeType = "DEPOSIT_FEE"
	WithdrawalFee FeeType = "WITHDRAWAL_FEE"
	TransferFee   FeeType = "TRANSFER_FEE"
	TakerFee      FeeType = "TAKER_FEE"
	MakerFee      FeeType = "MAKER_FEE"
)

type Fee struct {
	Type        FeeType
	Min         decimal.Decimal
	Max         decimal.Decimal
	Percentatge int
	Currency    CurrencyName
}
