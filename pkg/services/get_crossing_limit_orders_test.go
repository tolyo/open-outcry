package services

import "open-outcry/pkg/models/trade_order"

func (assert *ServiceTestSuite) TestGetCrossingLimitOrdersSellSidePrice() {
	assert.Equal(0, GetCrossingLimitOrders(1, tradeorder.Sell, 10.00))
	ProcessTradeOrder(assert.instrumentAccount1, "BTC_EUR", "LIMIT", tradeorder.Sell, 10.00, 1, "GTC")
	assert.Equal(1, GetCrossingLimitOrders(1, tradeorder.Sell, 10.00))
	ProcessTradeOrder(assert.instrumentAccount1, "BTC_EUR", "LIMIT", tradeorder.Sell, 10.00, 1, "GTC")
	assert.Equal(2, GetCrossingLimitOrders(1, tradeorder.Sell, 10.00))
	ProcessTradeOrder(assert.instrumentAccount1, "BTC_EUR", "LIMIT", tradeorder.Sell, 9.00, 1, "GTC")
	assert.Equal(3, GetCrossingLimitOrders(1, tradeorder.Sell, 10.00))
	ProcessTradeOrder(assert.instrumentAccount1, "BTC_EUR", "LIMIT", tradeorder.Sell, 19.00, 1, "GTC")
	assert.Equal(3, GetCrossingLimitOrders(1, tradeorder.Sell, 10.00))
	ProcessTradeOrder(assert.instrumentAccount2, "BTC_EUR", "LIMIT", tradeorder.Buy, 10.00, 1, "GTC")
	assert.Equal(2, GetCrossingLimitOrders(1, tradeorder.Sell, 10.00))
	ProcessTradeOrder(assert.instrumentAccount1, "BTC_EUR", "LIMIT", tradeorder.Sell, 10.01, 1, "GTC")
	assert.Equal(2, GetCrossingLimitOrders(1, tradeorder.Sell, 10.00))
	ProcessTradeOrder(assert.instrumentAccount1, "BTC_EUR", "LIMIT", tradeorder.Sell, 9.999999, 1, "GTC")
	assert.Equal(3, GetCrossingLimitOrders(1, tradeorder.Sell, 10.00))
}

func (assert *ServiceTestSuite) TestGetCrossingLimitOrdersPriceBuySide() {
	assert.Equal(0, GetCrossingLimitOrders(1, tradeorder.Buy, 10.00))
	ProcessTradeOrder(assert.instrumentAccount1, "BTC_EUR", "LIMIT", tradeorder.Buy, 10.00, 1, "GTC")
	assert.Equal(1, GetCrossingLimitOrders(1, tradeorder.Buy, 10.00))
	ProcessTradeOrder(assert.instrumentAccount1, "BTC_EUR", "LIMIT", tradeorder.Buy, 10.00, 1, "GTC")
	assert.Equal(2, GetCrossingLimitOrders(1, tradeorder.Buy, 10.00))
	ProcessTradeOrder(assert.instrumentAccount1, "BTC_EUR", "LIMIT", tradeorder.Buy, 11.00, 1, "GTC")
	assert.Equal(3, GetCrossingLimitOrders(1, tradeorder.Buy, 10.00))
	ProcessTradeOrder(assert.instrumentAccount1, "BTC_EUR", "LIMIT", tradeorder.Buy, 9.00, 1, "GTC")
	assert.Equal(3, GetCrossingLimitOrders(1, tradeorder.Buy, 10.00))
	ProcessTradeOrder(assert.instrumentAccount1, "BTC_EUR", "LIMIT", tradeorder.Sell, 10.00, 1, "GTC")
	assert.Equal(3, GetCrossingLimitOrders(1, tradeorder.Buy, 10.00))
	ProcessTradeOrder(assert.instrumentAccount1, "BTC_EUR", "LIMIT", tradeorder.Sell, 9.99999, 1, "GTC")
	assert.Equal(3, GetCrossingLimitOrders(1, tradeorder.Buy, 10.00))
	ProcessTradeOrder(assert.instrumentAccount1, "BTC_EUR", "LIMIT", tradeorder.Buy, 10.000001, 1, "GTC")
	assert.Equal(4, GetCrossingLimitOrders(1, tradeorder.Buy, 10.00))
}
