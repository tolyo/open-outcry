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

//
//  test "GetAvailableLimitVolume/3 sell side multiple orders same price" {
// when given a new sell order
//    tradingAccount = Acc()
//    ProcessTradeOrder(tradingAccount, "BTC_EUR", "LIMIT", "SELL", 10, 100, "GTC")
//    ProcessTradeOrder(tradingAccount, "BTC_EUR", "LIMIT", "SELL", 10, 100, "GTC")
// then expect the available volume to increase
//    assert.Equal(GetAvailableLimitVolume("SELL", 10.00) == 200
//    assert.Equal(GetAvailableLimitVolume("SELL", 11.00) == 200
//    assert.Equal(GetAvailableLimitVolume("SELL", 9.00) == 0
//    assert.Equal(GetAvailableLimitVolume("BUY", 10.00) == 0
//  }
//
//  test "GetAvailableLimitVolume/3 sell side multiple orders different prices" {
// when given multiple new sell orders
//    tradingAccount = Acc()
//    ProcessTradeOrder(tradingAccount, "BTC_EUR", "LIMIT", "SELL", 10, 100, "GTC")
//    ProcessTradeOrder(tradingAccount, "BTC_EUR", "LIMIT", "SELL", 9, 100, "GTC")
//
// then expect the available volume to increase
//    assert.Equal(GetAvailableLimitVolume("SELL", 10.00) == 200
//    assert.Equal(GetAvailableLimitVolume("SELL", 9.00) == 100
//    assert.Equal(GetAvailableLimitVolume("SELL", 11.00) == 200
//    assert.Equal(GetAvailableLimitVolume("SELL", 8.99) == 0
//    assert.Equal(GetAvailableLimitVolume("BUY", 10.00) == 0
//  }
//
//  @{c `
//    Test for available volume on the buy side. Available volume should increase
//    if the order is on the buy side and order limit price is above
//    or equal the query limit price.
//  `
//  test "GetAvailableLimitVolume/3 buy side single order" {
// when given a new buy order
//    tradingAccount = Acc()
//    ProcessTradeOrder(tradingAccount, "BTC_EUR", "LIMIT", "BUY", 10, 10, "GTC")
//
// then expect the available volume to increase
//    assert.Equal(GetAvailableLimitVolume("BUY", 10.00) == 10
//    assert.Equal(GetAvailableLimitVolume("BUY", 9.00) == 10
//    assert.Equal(GetAvailableLimitVolume("BUY", 11.00) == 0
//    assert.Equal(GetAvailableLimitVolume("SELL", 10.00) == 0
//  }
//
//  test "GetAvailableLimitVolume/3 buy side multiple orders same price" {
// when given 2 new buy orders
//    tradingAccount = Acc()
//    ProcessTradeOrder(tradingAccount, "BTC_EUR", "LIMIT", "BUY", 10, 10, "GTC")
//    ProcessTradeOrder(tradingAccount, "BTC_EUR", "LIMIT", "BUY", 10, 10, "GTC")
//
// then expect the available volume to increase
//    assert.Equal(GetAvailableLimitVolume("BUY", 10.00) == 20
//    assert.Equal(GetAvailableLimitVolume("BUY", 9.00) == 20
//    assert.Equal(GetAvailableLimitVolume("BUY", 11.00) == 0
//    assert.Equal(GetAvailableLimitVolume("SELL", 10.00) == 0
//  }
//
//  test "GetAvailableLimitVolume/3 buy side multiple orders different prices" {
// when given a new sell order
//    tradingAccount = Acc()
//    ProcessTradeOrder(tradingAccount, "BTC_EUR", "LIMIT", "BUY", 10, 10, "GTC")
//    ProcessTradeOrder(tradingAccount, "BTC_EUR", "LIMIT", "BUY", 9, 10, "GTC")
//
// then expect the available volume to increase
//    assert.Equal(GetAvailableLimitVolume("BUY", 10.00) == 10
//    assert.Equal(GetAvailableLimitVolume("BUY", 9.00) == 20
//    assert.Equal(GetAvailableLimitVolume("BUY", 11.00) == 0
//    assert.Equal(GetAvailableLimitVolume("BUY", 9.99) == 10
//    assert.Equal(GetAvailableLimitVolume("BUY", 10.000001) == 0
//    assert.Equal(GetAvailableLimitVolume("SELL", 10.00) == 0
//  }
//}
