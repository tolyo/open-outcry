package models

// // `trade.pub_id` db reference
type TradeId string

type Trade struct {
	Id            TradeId
	InstrumentId  InstrumentId
	Price         float64
	Amount        float64
	SellerOrderId TradeOrderId
	BuyerOrderId  TradeOrderId
	TakerOrderId  TradeOrderId
}
