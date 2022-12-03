package models

// See SQL for atom funcinitions
type TradeOrderType string

const (
	Limit     TradeOrderType = "LIMIT"
	Market    TradeOrderType = "MARKET"
	StopLoss  TradeOrderType = "STOPLOSS"
	StopLimit TradeOrderType = "STOPLIMIT"
)
