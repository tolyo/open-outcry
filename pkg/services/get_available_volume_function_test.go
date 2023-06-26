package services

import "open-outcry/pkg/models"

func (assert *ServiceTestSuite) TestGetAvailableLimitVolumeEmpty() {
	// expect return none
	assert.Equal(0.0, GetAvailableLimitVolume(models.Sell, 10.00))
	// expect return none
	assert.Equal(0.0, GetAvailableLimitVolume(models.Buy, 10.00))
}

// Test for available volume on the sell side. Available volume should increase
// if the order is on the sell side and order limit price is below or equal the query limit price.
func (assert *ServiceTestSuite) TestGetAvailableLimitVolumeSellSingleOrder() {
	// when given a new sell order
	ProcessTradeOrder(assert.tradingAccount1, "BTC_EUR", "LIMIT", models.Sell, 10.00, 100, "GTC")

	// then expect the available volume to increase
	assert.Equal(100.00, GetAvailableLimitVolume(models.Sell, 10.00))
	assert.Equal(100.0, GetAvailableLimitVolume(models.Sell, 11.00))
	assert.Equal(0.0, GetAvailableLimitVolume(models.Sell, 9.00))
	assert.Equal(0.0, GetAvailableLimitVolume(models.Buy, 10.00))
}

func (assert *ServiceTestSuite) TestGetAvailableLimitVolumeSellSideMultipleOrdersSamePrice() {
	// when given a new sell order
	ProcessTradeOrder(assert.tradingAccount1, "BTC_EUR", "LIMIT", models.Sell, 10, 100, "GTC")
	ProcessTradeOrder(assert.tradingAccount1, "BTC_EUR", "LIMIT", models.Sell, 10, 100, "GTC")

	// then expect the available volume to increase
	assert.Equal(200.0, GetAvailableLimitVolume(models.Sell, 10.00))
	assert.Equal(200.0, GetAvailableLimitVolume(models.Sell, 11.00))
	assert.Equal(0.0, GetAvailableLimitVolume(models.Sell, 9.00))
	assert.Equal(0.0, GetAvailableLimitVolume(models.Buy, 10.00))

}

func (assert *ServiceTestSuite) TestGetAvailableLimitVolumeSellSideMultipleOrdersDifferentPrices() {
	// when given multiple new sell orders
	ProcessTradeOrder(assert.tradingAccount1, "BTC_EUR", "LIMIT", models.Sell, 10, 100, "GTC")
	ProcessTradeOrder(assert.tradingAccount1, "BTC_EUR", "LIMIT", models.Sell, 9, 100, "GTC")

	// then expect the available volume to increase
	assert.Equal(200.0, GetAvailableLimitVolume(models.Sell, 10.00))
	assert.Equal(100.0, GetAvailableLimitVolume(models.Sell, 9.00))
	assert.Equal(200.0, GetAvailableLimitVolume(models.Sell, 11.00))
	assert.Equal(0.0, GetAvailableLimitVolume(models.Sell, 8.99))
	assert.Equal(0.0, GetAvailableLimitVolume(models.Buy, 10.00))
}

//	Test for available volume on the buy side. Available volume should increase
//	if the order is on the buy side and order limit price is above
//	or equal the query limit price.
//
// `
func (assert *ServiceTestSuite) TestGetAvailableLimitVolumeBuySideSingleOrder() {
	// when given a new buy order
	ProcessTradeOrder(assert.tradingAccount1, "BTC_EUR", "LIMIT", models.Buy, 10, 10, "GTC")

	// then expect the available volume to increase
	assert.Equal(10.0, GetAvailableLimitVolume(models.Buy, 10.00))
	assert.Equal(10.0, GetAvailableLimitVolume(models.Buy, 9.00))
	assert.Equal(0.0, GetAvailableLimitVolume(models.Buy, 11.00))
	assert.Equal(0.0, GetAvailableLimitVolume(models.Sell, 10.00))
}

func (assert *ServiceTestSuite) TestGetAvailableLimitVolumeBuySideMultipleOrdersSamePrice() {
	// when given 2 new buy orders
	ProcessTradeOrder(assert.tradingAccount1, "BTC_EUR", "LIMIT", models.Buy, 10, 10, "GTC")
	ProcessTradeOrder(assert.tradingAccount1, "BTC_EUR", "LIMIT", models.Buy, 10, 10, "GTC")

	// then expect the available volume to increase
	assert.Equal(20.0, GetAvailableLimitVolume(models.Buy, 10.00))
	assert.Equal(20.0, GetAvailableLimitVolume(models.Buy, 9.00))
	assert.Equal(0.0, GetAvailableLimitVolume(models.Buy, 11.00))
	assert.Equal(0.0, GetAvailableLimitVolume(models.Sell, 10.00))
}

func (assert *ServiceTestSuite) TestGetAvailableLimitVolumeBuySideMultipleOrdersDifferentPrices() {
	// when given a new sell order
	ProcessTradeOrder(assert.tradingAccount1, "BTC_EUR", "LIMIT", models.Buy, 10, 10, "GTC")
	ProcessTradeOrder(assert.tradingAccount1, "BTC_EUR", "LIMIT", models.Buy, 9, 10, "GTC")

	// then expect the available volume to increase
	assert.Equal(10.0, GetAvailableLimitVolume(models.Buy, 10.00))
	assert.Equal(20.0, GetAvailableLimitVolume(models.Buy, 9.00))
	assert.Equal(0.0, GetAvailableLimitVolume(models.Buy, 11.00))
	assert.Equal(10.0, GetAvailableLimitVolume(models.Buy, 9.99))
	assert.Equal(0.0, GetAvailableLimitVolume(models.Buy, 10.000001))
	assert.Equal(0.0, GetAvailableLimitVolume(models.Sell, 10.00))

}
