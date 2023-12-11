package models

import "open-outcry/pkg/db"

type OrderTimeInForce string

const (
	GTC OrderTimeInForce = "GTC"
	IOC OrderTimeInForce = "IOC"
	FOK OrderTimeInForce = "FOK"
	GTD OrderTimeInForce = "GTD"
	GTT OrderTimeInForce = "GTT"
)

type OrderFill string

const (
	Full    OrderFill = "FULL"
	Partial OrderFill = "PARTIAL"
	None    OrderFill = "NONE"
)

type OrderSide string

const (
	Buy  OrderSide = "BUY"
	Sell OrderSide = "SELL"
)

type OrderStatus string

const (
	Open               OrderStatus = "OPEN"
	PartiallyFilled    OrderStatus = "PARTIALLY_FILLED"
	Cancelled          OrderStatus = "CANCELLED"
	PartiallyCancelled OrderStatus = "PARTIALLY_CANCELLED"
	PartiallyRejected  OrderStatus = "PARTIALLY_REJECTED"
	Filled             OrderStatus = "FILLED"
	Rejected           OrderStatus = "REJECTED"
)

// OrderType See SQL for atom funcinitions
type OrderType string

const (
	Limit     OrderType = "LIMIT"
	Market    OrderType = "MARKET"
	StopLoss  OrderType = "STOPLOSS"
	StopLimit OrderType = "STOPLIMIT"
)

// `trade_order.pub_id` db reference
type TradeOrderId string

// The limit price at which order may be executed
type OrderPrice float64

type TradeOrder struct {
	Id               TradeOrderId
	TradingAccountId TradingAccountId
	InstrumentName   InstrumentName
	Side             OrderSide
	Type             OrderType
	Price            OrderPrice
	Amount           float64 //     Amount of instrument to buy or sell
	//     For market the amount on the buy side becomes the amount in quote currency!
	OpenAmount  float64 //     Order amount available for trading
	Status      OrderStatus
	TimeInForce OrderTimeInForce
	Created     string
}

const tradeOrderBaseQuery = `
    SELECT
      t.pub_id,
      ta.pub_id,
      i.name,
      t.side::text,
      t.order_type::text,
      t.price,
      t.amount,
      t.open_amount,
      t.status::text,
      t.time_in_force::text,
      t.created_at

    FROM trade_order AS t
    INNER JOIN trading_account ta
      ON ta.id = t.trading_account_id
    INNER JOIN instrument i
      ON t.instrument_id = i.id
  `

func GetTradeOrder(id TradeOrderId) TradeOrder {
	var order TradeOrder
	db.Instance().QueryRow(tradeOrderBaseQuery+`WHERE t.pub_id = $1`, id).Scan(
		&order.Id,
		&order.TradingAccountId,
		&order.InstrumentName,
		&order.Side,
		&order.Type,
		&order.Price,
		&order.Amount,
		&order.OpenAmount,
		&order.Status,
		&order.TimeInForce,
		&order.Created,
	)
	return order
}
