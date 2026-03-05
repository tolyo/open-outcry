package models

import (
	"open-outcry/pkg/db"

	log "github.com/sirupsen/logrus"
)

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

// TradeOrderId `trade_order.pub_id` db reference
type TradeOrderId string

// OrderPrice The limit price at which order may be executed
type OrderPrice float64

type TradeOrder struct {
	Id               TradeOrderId
	InstrumentAccountId InstrumentAccountId
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
    INNER JOIN instrument_account ta
      ON ta.id = t.instrument_account_id
    INNER JOIN instrument i
      ON t.instrument_id = i.id
  `

func GetTradeOrder(id TradeOrderId) TradeOrder {
	var order TradeOrder
	err := db.Instance().QueryRow(tradeOrderBaseQuery+`WHERE t.pub_id = $1`, id).Scan(
		&order.Id,
		&order.InstrumentAccountId,
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
	if err != nil {
		log.Fatal(err)
	}
	return order
}

func GetTradeOrdersByInstrumentAccount(instrumentAccountId InstrumentAccountId) []TradeOrder {
	query := tradeOrderBaseQuery + `WHERE ta.pub_id = $1 ORDER BY t.created_at DESC`
	rows, err := db.Instance().Query(query, instrumentAccountId)
	if err != nil {
		log.Error(err)
		return nil
	}
	defer rows.Close()

	var orders []TradeOrder
	for rows.Next() {
		var o TradeOrder
		err := rows.Scan(
			&o.Id, &o.InstrumentAccountId, &o.InstrumentName,
			&o.Side, &o.Type, &o.Price, &o.Amount, &o.OpenAmount,
			&o.Status, &o.TimeInForce, &o.Created,
		)
		if err != nil {
			log.Error(err)
			continue
		}
		orders = append(orders, o)
	}
	return orders
}

func GetBookOrdersByInstrumentAccount(instrumentAccountId InstrumentAccountId) []TradeOrder {
	query := tradeOrderBaseQuery + `
		INNER JOIN book_order b ON b.trade_order_id = t.id
		WHERE ta.pub_id = $1
		ORDER BY t.created_at DESC
	`
	rows, err := db.Instance().Query(query, instrumentAccountId)
	if err != nil {
		log.Error(err)
		return nil
	}
	defer rows.Close()

	var orders []TradeOrder
	for rows.Next() {
		var o TradeOrder
		err := rows.Scan(
			&o.Id, &o.InstrumentAccountId, &o.InstrumentName,
			&o.Side, &o.Type, &o.Price, &o.Amount, &o.OpenAmount,
			&o.Status, &o.TimeInForce, &o.Created,
		)
		if err != nil {
			log.Error(err)
			continue
		}
		orders = append(orders, o)
	}
	return orders
}

// Trade represents an executed trade
type Trade struct {
	Id             string
	InstrumentName InstrumentName
	Price          float64
	Amount         float64
	SellerOrderId  TradeOrderId
	BuyerOrderId   TradeOrderId
	Created        string
}

const tradeBaseQuery = `
	SELECT
	  tr.pub_id,
	  i.name,
	  tr.price,
	  tr.amount,
	  so.pub_id,
	  bo.pub_id,
	  tr.created_at
	FROM trade tr
	INNER JOIN instrument i ON tr.instrument_id = i.id
	INNER JOIN trade_order so ON tr.seller_order_id = so.id
	INNER JOIN trade_order bo ON tr.buyer_order_id = bo.id
`

func GetTrade(id string) *Trade {
	var t Trade
	err := db.Instance().QueryRow(tradeBaseQuery+`WHERE tr.pub_id = $1`, id).Scan(
		&t.Id, &t.InstrumentName, &t.Price, &t.Amount,
		&t.SellerOrderId, &t.BuyerOrderId, &t.Created,
	)
	if err != nil {
		log.Error(err)
		return nil
	}
	return &t
}

func GetTradesByInstrumentAccount(instrumentAccountId InstrumentAccountId) []Trade {
	query := tradeBaseQuery + `
		WHERE so.instrument_account_id = (SELECT id FROM instrument_account WHERE pub_id = $1)
		   OR bo.instrument_account_id = (SELECT id FROM instrument_account WHERE pub_id = $1)
		ORDER BY tr.created_at DESC
	`
	rows, err := db.Instance().Query(query, instrumentAccountId)
	if err != nil {
		log.Error(err)
		return nil
	}
	defer rows.Close()

	var trades []Trade
	for rows.Next() {
		var t Trade
		err := rows.Scan(
			&t.Id, &t.InstrumentName, &t.Price, &t.Amount,
			&t.SellerOrderId, &t.BuyerOrderId, &t.Created,
		)
		if err != nil {
			log.Error(err)
			continue
		}
		trades = append(trades, t)
	}
	return trades
}
