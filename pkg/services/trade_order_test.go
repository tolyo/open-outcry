package services

func (assert *ServiceTestSuite) TestCreateTradeOrderBook() {
	// given
	tradingAccountId := Acc()

	// when given a new limit order
	ProcessTradeOrder(tradingAccountId,
		"BTC_EUR",
		"LIMIT", "SELL",
		20.10,
		10, "GTC")

	// then should be saved
	assert.Equal(1, GetSellBookOrderCount())

	// when given a new market order
	ProcessTradeOrder(tradingAccountId,
		"BTC_EUR", "MARKET", "SELL", 0, 10, "GTC")

	// then should be saved
	assert.Equal(2, GetSellBookOrderCount())

	// when given a stop loss order
	ProcessTradeOrder(tradingAccountId,
		"BTC_EUR", "STOPLOSS", "SELL", 20.10, 10, "GTC")

	// then should be not be saved to order book
	assert.Equal(2, GetSellBookOrderCount())

	// when given a stop limit order
	ProcessTradeOrder(tradingAccountId, "BTC_EUR", "STOPLIMIT", "SELL", 20.10, 10, "GTC")

	// then should be not be saved to order book
	assert.Equal(2, GetSellBookOrderCount())
}
