package services

func (assert *ServiceTestSuite) TestGetCrossingLimitOrdersSellSidePrice() {
	// given:
	tradingAccount := Acc()
	// then should return none
	assert.Equal(0, GetCrossingLimitOrders(1, "SELL", 10.00))

	// when given a new order
	ProcessTradeOrder(tradingAccount, "BTC_EUR", "LIMIT", "SELL", 10.00, 1, "GTC")
	// then count should increase
	assert.Equal(1, GetCrossingLimitOrders(1, "SELL", 10.00))

	// when given another new order
	ProcessTradeOrder(tradingAccount, "BTC_EUR", "LIMIT", "SELL", 10.00, 1, "GTC")

	// then count should increase
	assert.Equal(2, GetCrossingLimitOrders(1, "SELL", 10.00))

	// when given another new order with crossing price
	ProcessTradeOrder(tradingAccount, "BTC_EUR", "LIMIT", "SELL", 9.00, 1, "GTC")

	// then count should increase
	assert.Equal(3, GetCrossingLimitOrders(1, "SELL", 10.00))

	// when given another new order non crossing price
	ProcessTradeOrder(tradingAccount, "BTC_EUR", "LIMIT", "SELL", 19.00, 1, "GTC")

	// then count should not change
	assert.Equal(3, GetCrossingLimitOrders(1, "SELL", 10.00))

	// when given another new order with crossing price for buy side
	tradingAccount2 := Acc2()
	ProcessTradeOrder(tradingAccount2, "BTC_EUR", "LIMIT", "BUY", 10.00, 1, "GTC")

	// then count should decrease
	assert.Equal(2, GetCrossingLimitOrders(1, "SELL", 10.00))

	// when given another new order non crossing price
	ProcessTradeOrder(tradingAccount, "BTC_EUR", "LIMIT", "SELL", 10.01, 1, "GTC")

	// then count should not increase
	assert.Equal(2, GetCrossingLimitOrders(1, "SELL", 10.00))

	// when given another new order with crossing price
	ProcessTradeOrder(tradingAccount, "BTC_EUR", "LIMIT", "SELL", 10.000-0.000001, 1, "GTC")

	// then count should increase because buy side is emtpy
	assert.Equal(3, GetCrossingLimitOrders(1, "SELL", 10.00))
}

func (assert *ServiceTestSuite)	TestGetCrossingLimitOrdersPriceBuySide() {
	//  should return none
	assert.Equal(0, GetCrossingLimitOrders(1, "BUY", 10.00))
	// when given a new order
	tradingAccount := Acc()
	// given:
	ProcessTradeOrder(tradingAccount, "BTC_EUR", "LIMIT", "BUY", 10.00, 1, "GTC")

	// then count should increase
	assert.Equal(1, GetCrossingLimitOrders(1, "BUY", 10.00))

	// when given another new order
	ProcessTradeOrder(tradingAccount, "BTC_EUR", "LIMIT", "BUY", 10.00, 1, "GTC")

	//  #then count should increase
	assert.Equal(2, GetCrossingLimitOrders(1, "BUY", 10.00))

	// when given another new order with crossing price
	ProcessTradeOrder(tradingAccount, "BTC_EUR", "LIMIT", "BUY", 11.00, 1, "GTC")

	// then count should increase
	assert.Equal(3, GetCrossingLimitOrders(1, "BUY", 10.00))

	//  #when given another new order non crossing price
	ProcessTradeOrder(tradingAccount, "BTC_EUR", "LIMIT", "BUY", 9.00, 1, "GTC")

	// then count should not change
	assert.Equal(3, GetCrossingLimitOrders(1, "BUY", 10.00))

	//  when given another new order with crossing price for sell side
	ProcessTradeOrder(tradingAccount, "BTC_EUR", "LIMIT", "SELL", 10.00, 1, "GTC")

	//  then count should not change
	assert.Equal(3, GetCrossingLimitOrders(1, "BUY", 10.00))

	//  #when given another new order non crossing price
	//  #   ProcessTradeOrder(%TradeOrder{order | price: 9.99999})
	ProcessTradeOrder(tradingAccount, "BTC_EUR", "LIMIT", "SELL", 9.99999, 1, "GTC")

	// then count should not change
	assert.Equal(3, GetCrossingLimitOrders(1, "BUY", 10.00))

	// when given another new order with crossing price
	ProcessTradeOrder(tradingAccount, "BTC_EUR", "LIMIT", "BUY", 10.000001, 1, "GTC")

	//  #then count should increase
	assert.Equal(4, GetCrossingLimitOrders(1, "BUY", 10.00))
}
