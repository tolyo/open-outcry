package models

type TradeOrderStatus string

const (
	Open               TradeOrderStatus = "OPEN"
	PartiallyFilled    TradeOrderStatus = "PARTIALLY_FILLED"
	Cancelled          TradeOrderStatus = "CANCELLED"
	PartiallyCancelled TradeOrderStatus = "PARTIALLY_CANCELLEd"
	PartiallyRejected  TradeOrderStatus = "PARTIALLY_REJECTED"
	Filled             TradeOrderStatus = "FILLED"
	Rejected           TradeOrderStatus = "REJECTED"
)
