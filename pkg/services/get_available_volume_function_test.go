package services

func (assert *ServiceTestSuite) TestGetAvailableLimitVolumeEmpty() {
	// expect return none
	assert.Equal(0.0, GetAvailableLimitVolume("SELL", 10.00))
	// expect return none
	assert.Equal(0.0, GetAvailableLimitVolume("BUY", 10.00))
}

// Test for available volume on the sell side. Available volume should increase
// if the order is on the sell side and order limit price is below or equal the query limit price.
func (assert *ServiceTestSuite) TestGetAvailableLimitVolumeSellSingleOrder() {
	// when given a new sell order
	ProcessTradeOrder(Acc(), "BTC_EUR", "LIMIT", "SELL", 10.00, 100, "GTC")

	// then expect the available volume to increase
	assert.Equal(100.00, GetAvailableLimitVolume("SELL", 10.00))
	assert.Equal(100.0, GetAvailableLimitVolume("SELL", 11.00))
	assert.Equal(0.0, GetAvailableLimitVolume("SELL", 9.00))
	assert.Equal(0.0, GetAvailableLimitVolume("BUY", 10.00))
}

func (assert *ServiceTestSuite) TestGetAvailableLimitVolumeSellSideMultipleOrdersSamePrice() {
	// when given a new sell order
	tradingAccount := Acc()
	ProcessTradeOrder(tradingAccount, "BTC_EUR", "LIMIT", "SELL", 10, 100, "GTC")
	ProcessTradeOrder(tradingAccount, "BTC_EUR", "LIMIT", "SELL", 10, 100, "GTC")

	// then expect the available volume to increase
	assert.Equal(200.0, GetAvailableLimitVolume("SELL", 10.00))
	assert.Equal(200.0, GetAvailableLimitVolume("SELL", 11.00))
	assert.Equal(0.0, GetAvailableLimitVolume("SELL", 9.00))
	assert.Equal(0.0, GetAvailableLimitVolume("BUY", 10.00))

}

func (assert *ServiceTestSuite) TestGetAvailableLimitVolumeSellSideMultipleOrdersDifferentPrices() {
	// when given multiple new sell orders
	tradingAccount := Acc()
	ProcessTradeOrder(tradingAccount, "BTC_EUR", "LIMIT", "SELL", 10, 100, "GTC")
	ProcessTradeOrder(tradingAccount, "BTC_EUR", "LIMIT", "SELL", 9, 100, "GTC")

	// then expect the available volume to increase
	assert.Equal(200.0, GetAvailableLimitVolume("SELL", 10.00))
	assert.Equal(100.0, GetAvailableLimitVolume("SELL", 9.00))
	assert.Equal(200.0, GetAvailableLimitVolume("SELL", 11.00))
	assert.Equal(0.0, GetAvailableLimitVolume("SELL", 8.99))
	assert.Equal(0.0, GetAvailableLimitVolume("BUY", 10.00))
}

//	Test for available volume on the buy side. Available volume should increase
//	if the order is on the buy side and order limit price is above
//	or equal the query limit price.
//
// `
func (assert *ServiceTestSuite) TestGetAvailableLimitVolumeBuySideSingleOrder() {
	// when given a new buy order
	tradingAccount := Acc()
	ProcessTradeOrder(tradingAccount, "BTC_EUR", "LIMIT", "BUY", 10, 10, "GTC")

	// then expect the available volume to increase
	assert.Equal(10.0, GetAvailableLimitVolume("BUY", 10.00))
	assert.Equal(10.0, GetAvailableLimitVolume("BUY", 9.00))
	assert.Equal(0.0, GetAvailableLimitVolume("BUY", 11.00))
	assert.Equal(0.0, GetAvailableLimitVolume("SELL", 10.00))
}

func (assert *ServiceTestSuite) TestGetAvailableLimitVolumeBuySideMultipleOrdersSamePrice() {
	// when given 2 new buy orders
	tradingAccount := Acc()
	ProcessTradeOrder(tradingAccount, "BTC_EUR", "LIMIT", "BUY", 10, 10, "GTC")
	ProcessTradeOrder(tradingAccount, "BTC_EUR", "LIMIT", "BUY", 10, 10, "GTC")

	// then expect the available volume to increase
	assert.Equal(20.0, GetAvailableLimitVolume("BUY", 10.00))
	assert.Equal(20.0, GetAvailableLimitVolume("BUY", 9.00))
	assert.Equal(0.0, GetAvailableLimitVolume("BUY", 11.00))
	assert.Equal(0.0, GetAvailableLimitVolume("SELL", 10.00))
}

func (assert *ServiceTestSuite) TestGetAvailableLimitVolumeBuySideMultipleOrdersDifferentPrices() {
	// when given a new sell order
	tradingAccount := Acc()
	ProcessTradeOrder(tradingAccount, "BTC_EUR", "LIMIT", "BUY", 10, 10, "GTC")
	ProcessTradeOrder(tradingAccount, "BTC_EUR", "LIMIT", "BUY", 9, 10, "GTC")

	// then expect the available volume to increase
	assert.Equal(10.0, GetAvailableLimitVolume("BUY", 10.00))
	assert.Equal(20.0, GetAvailableLimitVolume("BUY", 9.00))
	assert.Equal(0.0, GetAvailableLimitVolume("BUY", 11.00))
	assert.Equal(10.0, GetAvailableLimitVolume("BUY", 9.99))
	assert.Equal(0.0, GetAvailableLimitVolume("BUY", 10.000001))
	assert.Equal(0.0, GetAvailableLimitVolume("SELL", 10.00))

}
