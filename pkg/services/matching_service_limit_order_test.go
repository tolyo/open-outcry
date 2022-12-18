package services

func (assert *ServiceTestSuite) TestProcessLimitSellOrderSave() {
	// when: a limit order is sent to an empty matching unit
	res, _ := ProcessTradeOrder(Acc(), "BTC_EUR", "LIMIT", "SELL", 10.0, 100.0, "GTC")

	// then: a matching unit should save the trade order on save order to the order book
	assert.NotNil(res)
	assert.Equal(1, GetSellBookOrderCount())
	assert.Equal([]PriceVolume{{10.0, 100.0}}, GetVolumes("BTC_EUR", "SELL"))
}

func (assert *ServiceTestSuite) TestProcessLimitBuyOrderSave() {
	// when: a limit order is sent to an empty matching unit
	ProcessTradeOrder(Acc(), "BTC_EUR", "LIMIT", "BUY", 10, 100, "GTC")

	// then: a matching unit should save the orderSELL
	assert.Equal(1, GetBuyBookOrderCount())
	assert.Equal([]PriceVolume{{10.0, 100.0}}, GetVolumes("BTC_EUR", "BUY"))
}

func (assert *ServiceTestSuite) TestProcessLimitNoMatchCaseIncomingBuy() {
	// when: there is a SELL order in the book and a BUY limit order arrives that {es not cross
	ProcessTradeOrder(Acc(), "BTC_EUR", "LIMIT", "SELL", 10, 100, "GTC")
	ProcessTradeOrder(Acc2(), "BTC_EUR", "LIMIT", "BUY", 9, 100, "GTC")

	//then: the book should have both orders and no trade should be generated
	assert.Equal(1, GetSellBookOrderCount())
	assert.Equal(1, GetBuyBookOrderCount())
	assert.Equal(0, GetTradeCount())
	assert.Equal([]PriceVolume{{10.0, 100.0}}, GetVolumes("BTC_EUR", "SELL"))
	assert.Equal([]PriceVolume{{9.0, 100.0}}, GetVolumes("BTC_EUR", "BUY"))
}

func (assert *ServiceTestSuite) TestProcessLimitNoMatchCaseIncomingSell() {
	// when: there is a BUY order in the book and a SELL limit order arrives that {es not cross
	ProcessTradeOrder(Acc(), "BTC_EUR", "LIMIT", "BUY", 9, 100, "GTC")
	ProcessTradeOrder(Acc2(), "BTC_EUR", "LIMIT", "SELL", 10, 100, "GTC")

	// then: the book should have both orders and no trade should be generated
	assert.Equal(1, GetSellBookOrderCount())
	assert.Equal(1, GetBuyBookOrderCount())
	assert.Equal(0, GetTradeCount())
//
   assert.Equal([]PriceVolume{{10.0, 100.0}}, GetVolumes( "BTC_EUR", "SELL"))
   assert.Equal([]PriceVolume{{9.0, 100.0}}, GetVolumes("BTC_EUR", "BUY"))
}

func (assert *ServiceTestSuite) TestProcessLimitExactMatchIncomingBuy() {
	// when: there is a SELL order in the book
	ProcessTradeOrder(Acc(), "BTC_EUR", "LIMIT", "SELL", 10, 100, "GTC")

	// then:
	assert.Equal(1, GetSellBookOrderCount())
	assert.Equal(0, GetBuyBookOrderCount())
	assert.Equal(0, GetTradeCount())

	assert.Equal(100.0, GetAvailableLimitVolume("SELL", 10))
	assert.Equal([]PriceVolume{{10.0, 100.0}}, GetVolumes("BTC_EUR", "SELL"))

	// when: a BUY limit order arrives that crossed
	ProcessTradeOrder(Acc2(), "BTC_EUR", "LIMIT", "BUY", 10, 100, "GTC")
	// then: the book should have no orders and a single trade should be generated
	assert.Equal(0, GetSellBookOrderCount())
	assert.Equal(0, GetBuyBookOrderCount())
	assert.Equal(1, GetTradeCount())
	assert.Equal(0.0, GetAvailableLimitVolume("SELL", 10))
	assert.Equal([]PriceVolume{}, GetVolumes("BTC_EUR", "SELL"))
	assert.Equal([]PriceVolume{}, GetVolumes("BTC_EUR", "BUY"))
}

func (assert *ServiceTestSuite) TestProcessLimitPartialMatchIncomingBuySingleTrade() {
	ProcessTradeOrder(Acc(), "BTC_EUR", "LIMIT", "SELL", 10.00, 100.00, "GTC")

	//then:
	assert.Equal(1, GetSellBookOrderCount())
	assert.Equal(0, GetBuyBookOrderCount())
	assert.Equal(0, GetTradeCount())

	assert.Equal([]PriceVolume{{10.0, 100.0}}, GetVolumes("BTC_EUR", "SELL"))

	// when: incoming buy order is only partially matched
   	ProcessTradeOrder(Acc2(), "BTC_EUR", "LIMIT", "BUY", 10.00, 50.00, "GTC")

 	// then: the book should have one sell order and a single trade should be generated
	assert.Equal(1, GetSellBookOrderCount())
	assert.Equal(0, GetBuyBookOrderCount())
	assert.Equal(1, GetTradeCount())

	assert.Equal([]PriceVolume{{10.0, 50.0}}, GetVolumes("BTC_EUR", "SELL"))
   	assert.Equal([]PriceVolume{}, GetVolumes("BTC_EUR", "BUY"))
}

func (assert *ServiceTestSuite) TestProcessLimitOverflowMatchIncomingBuySingleTrade() {
	// when: there is a SELL order in the book
	ProcessTradeOrder(Acc(), "BTC_EUR", "LIMIT", "SELL", 10.00, 10.00, "GTC")

	// then:
	assert.Equal(1, GetSellBookOrderCount())
	assert.Equal(0, GetBuyBookOrderCount())
	assert.Equal(0, GetTradeCount())
	assert.Equal(10.0, GetAvailableLimitVolume("SELL", 10))

	// when: a BUY limit order arrives that crosses and is more that the book amount
   ProcessTradeOrder(Acc2(), "BTC_EUR", "LIMIT", "BUY", 10.00, 15, "GTC")

	// then: the book should be one buy order,  no sell orders and a one trade
	assert.Equal(0, GetSellBookOrderCount())
	assert.Equal(1, GetBuyBookOrderCount())
	assert.Equal(1, GetTradeCount())
   	assert.Equal([]PriceVolume{}, GetVolumes("BTC_EUR", "SELL"))


	assert.Equal([]PriceVolume{{10.00, 5.0}}, GetVolumes("BTC_EUR", "BUY"))
}

func (assert *ServiceTestSuite) TestProcessLimitPartialMatchIncomingSellSingleTrade() {
	// when: there is a BUY order
   ProcessTradeOrder(Acc(), "BTC_EUR", "LIMIT", "BUY", 10.00, 100.00, "GTC")

	// then:
	assert.Equal(0, GetSellBookOrderCount())
	assert.Equal(1, GetBuyBookOrderCount())
	assert.Equal(0, GetTradeCount())

	assert.Equal(100.0, GetAvailableLimitVolume("BUY", 10))

	// when: incoming SELL order is only partially matched
   ProcessTradeOrder(Acc2(), "BTC_EUR", "LIMIT", "SELL", 10.00, 50.00, "GTC")

	// then: the book should have one BUY order and a 1 trade should be generated
	assert.Equal(0, GetSellBookOrderCount())
	assert.Equal(1, GetBuyBookOrderCount())
	assert.Equal(1, GetTradeCount())

	assert.Equal(50.0, GetAvailableLimitVolume("BUY", 10))
   assert.Equal([]PriceVolume{}, GetVolumes("BTC_EUR", "SELL"))

   assert.Equal([]PriceVolume{{10.00, 50.0}}, GetVolumes("BTC_EUR", "BUY"))
}

func (assert *ServiceTestSuite) TestProcessLimitOverflowMatchIncomingSellSingleTrade() {
	// when: there is a BUY order in the book
	ProcessTradeOrder(Acc(), "BTC_EUR", "LIMIT", "BUY", 10.00, 100.00, "GTC")

	// then:
	assert.Equal(0, GetSellBookOrderCount())
	assert.Equal(1, GetBuyBookOrderCount())
	assert.Equal(0, GetTradeCount())
	assert.Equal(100.0, GetAvailableLimitVolume("BUY", 10))

	// when: a SELL limit order arrives that crosses and is more that the book amount
	ProcessTradeOrder(Acc2(), "BTC_EUR", "LIMIT", "SELL", 10.00, 150.00, "GTC")

	// then: the book should be one SELL order,  no BUY orders and 1 trade
	assert.Equal(1, GetSellBookOrderCount())
	assert.Equal(0, GetBuyBookOrderCount())
	assert.Equal(1, GetTradeCount())

	assert.Equal(50.0, GetAvailableLimitVolume("SELL", 10))

	assert.Equal([]PriceVolume{{10.00, 50.0}}, GetVolumes("BTC_EUR", "SELL"))

	assert.Equal([]PriceVolume{}, GetVolumes("BTC_EUR", "BUY"))
}

func (assert *ServiceTestSuite) TestProcessLimitExactMatchIncomingBuysMultipleTrades() {
	// when: there is a SELL order in the book and 2 BUY limit order arrive that cross
	ProcessTradeOrder(Acc(), "BTC_EUR", "LIMIT", "SELL", 10.00, 10.00, "GTC")

	// then:
	assert.Equal(1, GetSellBookOrderCount())
	assert.Equal(0, GetBuyBookOrderCount())
	assert.Equal(0, GetTradeCount())

	assert.Equal(10.0, GetAvailableLimitVolume("SELL", 10))

   // when: incoming buy order that are partially matched
   tradingAccount2 := Acc2()
   ProcessTradeOrder(tradingAccount2, "BTC_EUR", "LIMIT", "BUY", 10.00, 5.00, "GTC")
   ProcessTradeOrder(tradingAccount2, "BTC_EUR", "LIMIT", "BUY", 10.00, 5.00, "GTC")

   // then: the book should have 2 trades should be generated
	assert.Equal(0, GetSellBookOrderCount())
	assert.Equal(0, GetBuyBookOrderCount())
	assert.Equal(2, GetTradeCount())
	assert.Equal(0.0, GetAvailableLimitVolume("SELL", 10))
	assert.Equal([]PriceVolume{}, GetVolumes("BTC_EUR", "SELL"))
	assert.Equal([]PriceVolume{}, GetVolumes("BTC_EUR", "BUY"))
}

func (assert *ServiceTestSuite) TestProcessLimitExactMatchIncomingSellMultipleTrades() {
	tradingAccount := Acc()
	tradingAccount2 := Acc2()

	// when: there is a BUY order in the book and 2 SELL limit order arrive that cross
	ProcessTradeOrder(tradingAccount, "BTC_EUR", "LIMIT", "BUY", 10.00, 10.00, "GTC")

	// then:
	assert.Equal(0, GetSellBookOrderCount())
	assert.Equal(1, GetBuyBookOrderCount())
	assert.Equal(0, GetTradeCount())

	assert.Equal(10.0, GetAvailableLimitVolume("BUY", 10))

	// when: incoming buy order is only partially matched
	ProcessTradeOrder(tradingAccount2, "BTC_EUR", "LIMIT", "SELL", 10.00, 5.00, "GTC")
	ProcessTradeOrder(tradingAccount2, "BTC_EUR", "LIMIT", "SELL", 10.00, 5.00, "GTC")

	// then: the book should have 2 trades should be generated
	assert.Equal(0, GetSellBookOrderCount())
	assert.Equal(0, GetBuyBookOrderCount())
	assert.Equal(2, GetTradeCount())
	assert.Equal([]PriceVolume{}, GetVolumes("BTC_EUR", "SELL"))
	assert.Equal([]PriceVolume{}, GetVolumes("BTC_EUR", "BUY"))
}

func (assert *ServiceTestSuite) TestProcessLimitPartialMatchMultipleBookSellsToMultipleTrades() {
	// given:
	tradingAccount := Acc()
	tradingAccount2 := Acc2()

	// when: there are multiple SELL orders in the book and BUY limit order arrive that cross but can fill only partially one of the orders
	ProcessTradeOrder(tradingAccount, "BTC_EUR", "LIMIT", "SELL", 10.00, 50.00, "GTC")
	ProcessTradeOrder(tradingAccount, "BTC_EUR", "LIMIT", "SELL", 10.00, 50.00, "GTC")

	// then:
	assert.Equal(2, GetSellBookOrderCount())
	assert.Equal(0, GetBuyBookOrderCount())
	assert.Equal(0, GetTradeCount())

	assert.Equal(100.0, GetAvailableLimitVolume("SELL", 10))

	// when: incoming buy order is only partially matched
	ProcessTradeOrder(tradingAccount2, "BTC_EUR", "LIMIT", "BUY", 10.00, 75.00, "GTC")

	// then: the book should have 2 trades should be generated
	assert.Equal(1, GetSellBookOrderCount())
	assert.Equal(0, GetBuyBookOrderCount())
	assert.Equal(2, GetTradeCount())

	assert.Equal([]PriceVolume{{10.00, 25.0}}, GetVolumes("BTC_EUR", "SELL"))
	assert.Equal([]PriceVolume{}, GetVolumes("BTC_EUR", "BUY"))
}

func (assert *ServiceTestSuite) TestProcessLimitPartialMatchMultipleBookBuysToMultipleTrades() {
	// given:
	tradingAccount := Acc()
	tradingAccount2 := Acc2()

	// when: there are multiple BUY orders in the book and SELL limit order arrive that cross but can fill only partially one of the orders
	ProcessTradeOrder(tradingAccount, "BTC_EUR", "LIMIT", "BUY", 10.00, 50.00, "GTC")
	ProcessTradeOrder(tradingAccount, "BTC_EUR", "LIMIT", "BUY", 10.00, 50.00, "GTC")

	// then:
	assert.Equal(0, GetSellBookOrderCount())
	assert.Equal(2, GetBuyBookOrderCount())
	assert.Equal(0, GetTradeCount())
	assert.Equal(100.0, GetAvailableLimitVolume("BUY", 10))

	// when: incoming SELL order is only partially matched
	ProcessTradeOrder(tradingAccount2, "BTC_EUR", "LIMIT", "SELL", 10.00, 75.00, "GTC")

	// then: the book should have 2 trades should be generated
	assert.Equal(0, GetSellBookOrderCount())
	assert.Equal(1, GetBuyBookOrderCount())
	assert.Equal(2, GetTradeCount())

	assert.Equal([]PriceVolume{}, GetVolumes("BTC_EUR", "SELL"))

	assert.Equal([]PriceVolume{{10.00, 25.0}}, GetVolumes("BTC_EUR", "BUY"))
}

func (assert *ServiceTestSuite) TestProcessLimitExactMatchMultipleBookSellsToMultipleTrades() {
	// given:
	tradingAccount := Acc()
	tradingAccount2 := Acc2()

	// when:
	ProcessTradeOrder(tradingAccount, "BTC_EUR", "LIMIT", "SELL", 10.00, 10.00, "GTC")
	ProcessTradeOrder(tradingAccount, "BTC_EUR", "LIMIT", "SELL", 10.00, 10.00, "GTC")
	ProcessTradeOrder(tradingAccount, "BTC_EUR", "LIMIT", "SELL", 10.00, 10.00, "GTC")
	ProcessTradeOrder(tradingAccount, "BTC_EUR", "LIMIT", "SELL", 10.00, 10.00, "GTC")

	// then:
	assert.Equal(4, GetSellBookOrderCount())
	assert.Equal(0, GetBuyBookOrderCount())
	assert.Equal(0, GetTradeCount())

	assert.Equal(40.0, GetAvailableLimitVolume("SELL", 10))

	// when: incoming buy order is only partially matched
	ProcessTradeOrder(tradingAccount2, "BTC_EUR", "LIMIT", "BUY", 10.00, 40.00, "GTC")

	// then: the book should have 2 trades should be generated
	assert.Equal(0, GetSellBookOrderCount())
	assert.Equal(0, GetBuyBookOrderCount())
	assert.Equal(4, GetTradeCount())
	assert.Equal([]PriceVolume{}, GetVolumes("BTC_EUR", "SELL"))
	assert.Equal([]PriceVolume{}, GetVolumes("BTC_EUR", "BUY"))
}

func (assert *ServiceTestSuite) TestProcessLimitExactMatchMultipleBookBuyToMultipleTrades() {
	// given:
	tradingAccount := Acc()
	tradingAccount2 := Acc2()

	// when:
	ProcessTradeOrder(tradingAccount, "BTC_EUR", "LIMIT", "BUY", 10.00, 5.00, "GTC")
	ProcessTradeOrder(tradingAccount, "BTC_EUR", "LIMIT", "BUY", 10.00, 5.00, "GTC")
	ProcessTradeOrder(tradingAccount, "BTC_EUR", "LIMIT", "BUY", 10.00, 5.00, "GTC")
	ProcessTradeOrder(tradingAccount, "BTC_EUR", "LIMIT", "BUY", 10.00, 5.00, "GTC")

	// then:
	assert.Equal(0, GetSellBookOrderCount())
	assert.Equal(4, GetBuyBookOrderCount())
	assert.Equal(0, GetTradeCount())
	assert.Equal(20.0, GetAvailableLimitVolume("BUY", 10))

	// when: incoming buy order is only partially matched
	ProcessTradeOrder(tradingAccount2, "BTC_EUR", "LIMIT", "SELL", 10.00, 20.00, "GTC")

	// then: the book should have 2 trades should be generated
	assert.Equal(0, GetSellBookOrderCount())
	assert.Equal(0, GetBuyBookOrderCount())
	assert.Equal(4, GetTradeCount())
	assert.Equal([]PriceVolume{}, GetVolumes("BTC_EUR", "SELL"))
	assert.Equal([]PriceVolume{}, GetVolumes("BTC_EUR", "BUY"))
}

func (assert *ServiceTestSuite) TestProcessLimitIncompleteMatchMultipleBookSellsToMultipleTrades() {
	// given:
	tradingAccount := Acc()
	tradingAccount2 := Acc2()

	// when: there are multiple SELL orders in the book and BUY limit order arrive that cross but can be only partially filled
	ProcessTradeOrder(tradingAccount, "BTC_EUR", "LIMIT", "SELL", 10.00, 5.00, "GTC")
	ProcessTradeOrder(tradingAccount, "BTC_EUR", "LIMIT", "SELL", 10.00, 5.00, "GTC")

	// then:
	assert.Equal(2, GetSellBookOrderCount())
	assert.Equal(0, GetBuyBookOrderCount())
	assert.Equal(0, GetTradeCount())
	assert.Equal(10.0, GetAvailableLimitVolume("SELL", 10))

	// when: incoming buy order is only partially matched
	ProcessTradeOrder(tradingAccount2, "BTC_EUR", "LIMIT", "BUY", 10.00, 17, "GTC")

	// then: the book should have 2 trades should be generated
	assert.Equal(0, GetSellBookOrderCount())
	assert.Equal(1, GetBuyBookOrderCount())
	assert.Equal(2, GetTradeCount())
	assert.Equal([]PriceVolume{}, GetVolumes("BTC_EUR", "SELL"))

	assert.Equal([]PriceVolume{{10.00, 7.0}},GetVolumes("BTC_EUR", "BUY"))
}

func (assert *ServiceTestSuite) TestProcessLimitIncompleteMatchMultipleBookBuysToMultipleTrades() {
	// given:
	tradingAccount := Acc()
	tradingAccount2 := Acc2()

	// when: there are multiple BUY orders in the book and SELL limit order arrive that cross but can be only partially filled
	ProcessTradeOrder(tradingAccount, "BTC_EUR", "LIMIT", "BUY", 10.00, 50.00, "GTC")
	ProcessTradeOrder(tradingAccount, "BTC_EUR", "LIMIT", "BUY", 10.00, 50.00, "GTC")

	// then:
	assert.Equal(0, GetSellBookOrderCount())
	assert.Equal(2, GetBuyBookOrderCount())
	assert.Equal(0, GetTradeCount())

	assert.Equal(100.0, GetAvailableLimitVolume("BUY", 10))

	// when: incoming SELL order is only partially matched
	ProcessTradeOrder(tradingAccount2, "BTC_EUR", "LIMIT", "SELL", 10.00, 175.00, "GTC")

	// then: the book should have 2 trades should be generated
	assert.Equal(1, GetSellBookOrderCount())
	assert.Equal(0, GetBuyBookOrderCount())
	assert.Equal(2, GetTradeCount())

	assert.Equal([]PriceVolume{{10.00, 75.0}}, GetVolumes("BTC_EUR", "SELL"))

	assert.Equal([]PriceVolume{}, GetVolumes("BTC_EUR", "BUY"))
}

func (assert *ServiceTestSuite) TestProcessLimitPartialMatchMultipleBookSellsToMultipleTradesMultiplePrices() {
	// given:
	tradingAccount := Acc()
	tradingAccount2 := Acc2()

	// when: there are multiple SELL orders in the book and BUY limit order arrive that cross but can fill only partially one of the orders
	ProcessTradeOrder(tradingAccount, "BTC_EUR", "LIMIT", "SELL", 9.00, 50.00, "GTC")
	ProcessTradeOrder(tradingAccount, "BTC_EUR", "LIMIT", "SELL", 10.00, 50.00, "GTC")

	// then:
	assert.Equal(2, GetSellBookOrderCount())
	assert.Equal(0, GetBuyBookOrderCount())
	assert.Equal(0, GetTradeCount())

	// when: incoming buy order is only partially matched
	ProcessTradeOrder(tradingAccount2, "BTC_EUR", "LIMIT", "BUY", 11.00, 75.00, "GTC")

	// then: the book should have 2 trades should be generated
	assert.Equal(1, GetSellBookOrderCount())
	assert.Equal(0, GetBuyBookOrderCount())
	assert.Equal(2, GetTradeCount())

	assert.Equal([]PriceVolume{{10.00, 25.0}}, GetVolumes("BTC_EUR", "SELL"))

	assert.Equal([]PriceVolume{}, GetVolumes("BTC_EUR", "BUY"))
}

func (assert *ServiceTestSuite) TestProcessLimitPartialMatchMultipleBookBuysToMultipleTradesMultiplePrices() {
	// given:
	tradingAccount := Acc()
	tradingAccount2 := Acc2()

	// when: there are multiple BUY orders in the book and SELL limit order arrive that cross but can fill only partially one of the orders
	ProcessTradeOrder(tradingAccount, "BTC_EUR", "LIMIT", "BUY", 9.00, 50.00, "GTC")
	ProcessTradeOrder(tradingAccount, "BTC_EUR", "LIMIT", "BUY", 10.00, 50.00, "GTC")

	// then:
	assert.Equal(0, GetSellBookOrderCount())
	assert.Equal(2, GetBuyBookOrderCount())
	assert.Equal(0, GetTradeCount())
	assert.Equal(50.0, GetAvailableLimitVolume("BUY", 10))
	assert.Equal(100.0, GetAvailableLimitVolume("BUY", 9))

	// when: incoming SELL order is only partially matched
	ProcessTradeOrder(tradingAccount2, "BTC_EUR", "LIMIT", "SELL", 8.00, 75.00, "GTC")

	// then: the book should have 2 trades should be generated
	assert.Equal(0, GetSellBookOrderCount())
	assert.Equal(1, GetBuyBookOrderCount())
	assert.Equal(2, GetTradeCount())
	assert.Equal(25.0, GetAvailableLimitVolume("BUY", 9))
	assert.Equal([]PriceVolume{}, GetVolumes("BTC_EUR", "SELL"))

	assert.Equal([]PriceVolume{{9.00, 25.0}}, GetVolumes("BTC_EUR", "BUY"))
}

func (assert *ServiceTestSuite) TestProcessLimitSelfTradePreventions() {
	// given:
	tradingAccount := Acc()

	//when:
	ProcessTradeOrder(tradingAccount, "BTC_EUR", "LIMIT", "BUY", 10.00, 50.00, "GTC")
	assert.Equal(1, GetBuyBookOrderCount())
	// then:
	ProcessTradeOrder(tradingAccount, "BTC_EUR", "LIMIT", "SELL", 10.00, 50.00, "GTC")
	assert.Equal(1, GetSellBookOrderCount())
	assert.Equal(0, GetTradeCount())
	assert.Equal(50.0, GetAvailableLimitVolume("BUY", 10))
	assert.Equal(50.0, GetAvailableLimitVolume("SELL", 10))

}
