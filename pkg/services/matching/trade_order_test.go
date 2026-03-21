package matching

func (assert *ServiceTestSuite) TestCreateTradeOrderBook() {
	// given

	// when given a new limit order
	ProcessTradeOrder(assert.instrumentAccount1,
		"BTC_EUR",
		"LIMIT", Sell,
		20.10,
		10, "GTC")

	// then should be saved
	assert.Equal(1, GetSellBookOrderCount())

	// when given a new market order
	ProcessTradeOrder(assert.instrumentAccount1,
		"BTC_EUR", "MARKET", Sell, 0, 10, "GTC")

	// then should be saved
	assert.Equal(2, GetSellBookOrderCount())

	// when given a stop loss order
	ProcessTradeOrder(assert.instrumentAccount1,
		"BTC_EUR", "STOPLOSS", Sell, 20.10, 10, "GTC")

	// then should be not be saved to order book
	assert.Equal(2, GetSellBookOrderCount())

	// when given a stop limit order
	ProcessTradeOrder(assert.instrumentAccount1, "BTC_EUR", "STOPLIMIT", Sell, 20.10, 10, "GTC")

	// then should be not be saved to order book
	assert.Equal(2, GetSellBookOrderCount())
}
