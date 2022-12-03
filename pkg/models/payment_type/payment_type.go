package models

type PaymentType string

const (
	Deposit        PaymentType = "DEPOSIT"
	Withdrawal     PaymentType = "WITHDRAWAL"
	Transfer       PaymentType = "TRANSFER"
	InstrumentBuy  PaymentType = "INSTRUMENT_BUY"
	InstrumentSell PaymentType = "INSTRUMENT_SELL"
)
