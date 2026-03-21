package matching

import (
	instrument "open-outcry/pkg/models/instrument"
	instrumentaccount "open-outcry/pkg/models/instrument_account"
	tradeorder "open-outcry/pkg/models/trade_order"
	tradeordersvc "open-outcry/pkg/services/trade_order"
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
