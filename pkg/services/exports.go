package services

import (
	appentity "open-outcry/pkg/models/app_entity"
	"open-outcry/pkg/models/currency"
	"open-outcry/pkg/models/instrument"
	instrumentaccount "open-outcry/pkg/models/instrument_account"
	orderbook "open-outcry/pkg/models/order_book"
	tradeorder "open-outcry/pkg/models/trade_order"
	"open-outcry/pkg/models/transfer"
	"open-outcry/pkg/services/matching"
	orderbooksvc "open-outcry/pkg/services/order_book"
	"open-outcry/pkg/services/registration"
	tradeordersvc "open-outcry/pkg/services/trade_order"
	transfersvc "open-outcry/pkg/services/transfer"
)

func ProcessTradeOrder(
	instrumentAccountId instrumentaccount.InstrumentAccountId,
	instrumentName instrument.InstrumentName,
	orderType tradeorder.OrderType,
	side tradeorder.OrderSide,
	price tradeorder.OrderPrice,
	amount float64,
	timeInForce tradeorder.OrderTimeInForce,
) (tradeorder.TradeOrderId, error) {
	return tradeordersvc.ProcessTradeOrder(instrumentAccountId, instrumentName, orderType, side, price, amount, timeInForce)
}

func CancelTradeOrder(tradeOrderId tradeorder.TradeOrderId) error {
	return tradeordersvc.CancelTradeOrder(tradeOrderId)
}

func GetVolumeAtPrice(instrumentName instrument.InstrumentName, side tradeorder.OrderSide, price tradeorder.OrderPrice) float64 {
	return orderbooksvc.GetVolumeAtPrice(instrumentName, side, price)
}

func GetVolumes(instrumentName instrument.InstrumentName, side tradeorder.OrderSide) []orderbook.PriceVolume {
	return orderbooksvc.GetVolumes(instrumentName, side)
}

func GetOrderBook(instrumentName instrument.InstrumentName) orderbook.OrderBook {
	return orderbooksvc.GetOrderBook(instrumentName)
}

func CreateAppEntity(id appentity.AppEntityExternalId) appentity.AppEntityId {
	return registration.CreateAppEntity(id)
}

func CreateTransferDeposit(appEntityId appentity.AppEntityId, amount float64, currencyName currency.CurrencyName, reference string, details string) transfer.TransferJournalId {
	return transfersvc.CreateTransferDeposit(appEntityId, amount, currencyName, reference, details)
}

func CreateTransferDepositCustomFee(appEntityId appentity.AppEntityId, amount float64, currencyName currency.CurrencyName, reference string, details string, feeType any) transfer.TransferJournalId {
	return transfersvc.CreateTransferDepositCustomFee(appEntityId, amount, currencyName, reference, details, feeType)
}

func GetBookOrderCount(side tradeorder.OrderSide) int {
	return matching.GetBookOrderCount(side)
}

func GetSellBookOrderCount() int {
	return matching.GetSellBookOrderCount()
}

func GetBuyBookOrderCount() int {
	return matching.GetBuyBookOrderCount()
}

func GetTradeCount() int {
	return matching.GetTradeCount()
}

func GetTradePrices() []float64 {
	return matching.GetTradePrices()
}

func GetCrossingLimitOrders(instrumentId int, side tradeorder.OrderSide, price tradeorder.OrderPrice) int {
	return orderbooksvc.GetCrossingLimitOrders(instrumentId, side, price)
}

func GetAvailableLimitVolume(side tradeorder.OrderSide, price tradeorder.OrderPrice) float64 {
	return orderbooksvc.GetAvailableLimitVolume(side, price)
}
