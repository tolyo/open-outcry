package services

import (
	appentity "open-outcry/pkg/models/app_entity"
	currency "open-outcry/pkg/models/currency"
	currencyaccount "open-outcry/pkg/models/currency_account"
	instrument "open-outcry/pkg/models/instrument"
	instrumentaccount "open-outcry/pkg/models/instrument_account"
	orderbook "open-outcry/pkg/models/order_book"
	trade "open-outcry/pkg/models/trade"
	tradeorder "open-outcry/pkg/models/trade_order"
	transfer "open-outcry/pkg/models/transfer"
)

type AppEntityId = appentity.AppEntityId
type AppEntityExternalId = appentity.AppEntityExternalId

type CurrencyName = currency.CurrencyName
type Currency = currency.Currency

type CurrencyAccountId = currencyaccount.CurrencyAccountId
type CurrencyAccount = currencyaccount.CurrencyAccount

type InstrumentName = instrument.InstrumentName
type Instrument = instrument.Instrument

type InstrumentAccountId = instrumentaccount.InstrumentAccountId
type InstrumentAccountHolding = instrumentaccount.InstrumentAccountHolding
type InstrumentAccount = instrumentaccount.InstrumentAccount

type PriceVolume = orderbook.PriceVolume
type OrderBook = orderbook.OrderBook

type OrderTimeInForce = tradeorder.OrderTimeInForce
type OrderFill = tradeorder.OrderFill
type OrderSide = tradeorder.OrderSide
type OrderStatus = tradeorder.OrderStatus
type OrderType = tradeorder.OrderType
type TradeOrderId = tradeorder.TradeOrderId
type OrderPrice = tradeorder.OrderPrice

type TradeOrder = tradeorder.TradeOrder
type Trade = trade.Trade

type TransferJournalId = transfer.TransferJournalId
type TransferEntry = transfer.TransferEntry

const (
	GTC = tradeorder.GTC
	IOC = tradeorder.IOC
	FOK = tradeorder.FOK
	GTD = tradeorder.GTD
	GTT = tradeorder.GTT

	Full    = tradeorder.Full
	Partial = tradeorder.Partial
	None    = tradeorder.None

	Buy  = tradeorder.Buy
	Sell = tradeorder.Sell

	Open               = tradeorder.Open
	PartiallyFilled    = tradeorder.PartiallyFilled
	Cancelled          = tradeorder.Cancelled
	PartiallyCancelled = tradeorder.PartiallyCancelled
	PartiallyRejected  = tradeorder.PartiallyRejected
	Filled             = tradeorder.Filled
	Rejected           = tradeorder.Rejected

	Limit     = tradeorder.Limit
	Market    = tradeorder.Market
	StopLoss  = tradeorder.StopLoss
	StopLimit = tradeorder.StopLimit
)

func GetCurrencies() []Currency                        { return currency.GetCurrencies() }
func GetAppEntities() []appentity.AppEntity            { return appentity.GetAppEntities() }
func GetAppEntity(id AppEntityId) *appentity.AppEntity { return appentity.GetAppEntity(id) }
func GetCurrencyAccount(id CurrencyAccountId) *CurrencyAccount {
	return currencyaccount.GetCurrencyAccount(id)
}
func FindCurrencyAccountByAppEntityIdAndCurrencyName(appEntityId AppEntityId, currencyName CurrencyName) *CurrencyAccount {
	return currencyaccount.FindCurrencyAccountByAppEntityIdAndCurrencyName(appEntityId, currencyName)
}
func CreateCurrencyAccount(appEntityId AppEntityId, currencyName CurrencyName) CurrencyAccountId {
	return currencyaccount.CreateCurrencyAccount(appEntityId, currencyName)
}
func GetCurrencyAccountsByAppEntity(appEntityId AppEntityId) []CurrencyAccount {
	return currencyaccount.GetCurrencyAccountsByAppEntity(appEntityId)
}
func FindInstrumentAccountByApplicationEntityId(appEntityId AppEntityId) *InstrumentAccount {
	return instrumentaccount.FindInstrumentAccountByApplicationEntityId(appEntityId)
}
func GetInstrumentAccount(id InstrumentAccountId) *InstrumentAccount {
	return instrumentaccount.GetInstrumentAccount(id)
}
func GetInstrumentAccountHoldings(instrumentAccountId InstrumentAccountId) []InstrumentAccountHolding {
	return instrumentaccount.GetInstrumentAccountHoldings(instrumentAccountId)
}
func GetTradeOrder(id TradeOrderId) TradeOrder { return tradeorder.GetTradeOrder(id) }
func GetTradeOrdersByInstrumentAccount(instrumentAccountId InstrumentAccountId) []TradeOrder {
	return tradeorder.GetTradeOrdersByInstrumentAccount(instrumentAccountId)
}
func GetBookOrdersByInstrumentAccount(instrumentAccountId InstrumentAccountId) []TradeOrder {
	return tradeorder.GetBookOrdersByInstrumentAccount(instrumentAccountId)
}
func GetTrade(id string) *Trade { return trade.GetTrade(id) }
func GetTradesByInstrumentAccount(instrumentAccountId InstrumentAccountId) []Trade {
	return trade.GetTradesByInstrumentAccount(instrumentAccountId)
}
func GetTransfer(id string) *TransferEntry { return transfer.GetTransfer(id) }
func GetTransfersByAppEntity(appEntityId AppEntityId) []TransferEntry {
	return transfer.GetTransfersByAppEntity(appEntityId)
}
