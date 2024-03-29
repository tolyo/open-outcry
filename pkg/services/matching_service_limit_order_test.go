package services

import (
	"open-outcry/pkg/models"
)

func (assert *ServiceTestSuite) TestProcessLimitSellOrderSave() {
	// when: a limit order is sent to an empty matching unit
	res, _ := ProcessTradeOrder(assert.tradingAccount1, "BTC_EUR", "LIMIT", models.Sell, 10.0, 100.0, "GTC")

	// then: a matching unit should save the trade order on save order to the order book
	assert.NotNil(res)
	assert.Equal(1, GetSellBookOrderCount())
	assert.Equal([]models.PriceVolume{{Price: 10.0, Volume: 100.0}}, GetVolumes("BTC_EUR", models.Sell))
}

func (assert *ServiceTestSuite) TestProcessLimitBuyOrderSave() {
	// when: a limit order is sent to an empty matching unit
	ProcessTradeOrder(assert.tradingAccount1, "BTC_EUR", "LIMIT", models.Buy, 10, 100, "GTC")

	// then: a matching unit should save the orderSELL
	assert.Equal(1, GetBuyBookOrderCount())
	assert.Equal([]models.PriceVolume{{Price: 10.0, Volume: 100.0}}, GetVolumes("BTC_EUR", models.Buy))
}

func (assert *ServiceTestSuite) TestProcessLimitNoMatchCaseIncomingBuy() {
	// when: there is a SELL order in the book and a BUY limit order arrives that {es not cross
	ProcessTradeOrder(assert.tradingAccount1, "BTC_EUR", "LIMIT", models.Sell, 10, 100, "GTC")
	ProcessTradeOrder(assert.tradingAccount2, "BTC_EUR", "LIMIT", models.Buy, 9, 100, "GTC")

	//then: the book should have both orders and no trade should be generated
	assert.Equal(1, GetSellBookOrderCount())
	assert.Equal(1, GetBuyBookOrderCount())
	assert.Equal(0, GetTradeCount())
	assert.Equal([]models.PriceVolume{{Price: 10.0, Volume: 100.0}}, GetVolumes("BTC_EUR", models.Sell))
	assert.Equal([]models.PriceVolume{{Price: 9.0, Volume: 100.0}}, GetVolumes("BTC_EUR", models.Buy))
}

func (assert *ServiceTestSuite) TestProcessLimitNoMatchCaseIncomingSell() {
	// when: there is a BUY order in the book and a SELL limit order arrives that {es not cross
	ProcessTradeOrder(assert.tradingAccount1, "BTC_EUR", "LIMIT", models.Buy, 9, 100, "GTC")
	ProcessTradeOrder(assert.tradingAccount2, "BTC_EUR", "LIMIT", models.Sell, 10, 100, "GTC")

	// then: the book should have both orders and no trade should be generated
	assert.Equal(1, GetSellBookOrderCount())
	assert.Equal(1, GetBuyBookOrderCount())
	assert.Equal(0, GetTradeCount())
	//
	assert.Equal([]models.PriceVolume{{Price: 10.0, Volume: 100.0}}, GetVolumes("BTC_EUR", models.Sell))
	assert.Equal([]models.PriceVolume{{Price: 9.0, Volume: 100.0}}, GetVolumes("BTC_EUR", models.Buy))
}

func (assert *ServiceTestSuite) TestProcessLimitExactMatchIncomingBuy() {
	// when: there is a SELL order in the book
	ProcessTradeOrder(assert.tradingAccount1, "BTC_EUR", "LIMIT", models.Sell, 10, 100, "GTC")

	// then:
	assert.Equal(1, GetSellBookOrderCount())
	assert.Equal(0, GetBuyBookOrderCount())
	assert.Equal(0, GetTradeCount())

	assert.Equal(100.0, GetAvailableLimitVolume(models.Sell, 10))
	assert.Equal([]models.PriceVolume{{Price: 10.0, Volume: 100.0}}, GetVolumes("BTC_EUR", models.Sell))

	// when: a BUY limit order arrives that crossed
	ProcessTradeOrder(assert.tradingAccount2, "BTC_EUR", "LIMIT", models.Buy, 10, 100, "GTC")
	// then: the book should have no orders and a single trade should be generated
	assert.Equal(0, GetSellBookOrderCount())
	assert.Equal(0, GetBuyBookOrderCount())
	assert.Equal(1, GetTradeCount())
	assert.Equal(0.0, GetAvailableLimitVolume(models.Sell, 10))
	assert.Equal([]models.PriceVolume{}, GetVolumes("BTC_EUR", models.Sell))
	assert.Equal([]models.PriceVolume{}, GetVolumes("BTC_EUR", models.Buy))
}

func (assert *ServiceTestSuite) TestProcessLimitPartialMatchIncomingBuySingleTrade() {
	ProcessTradeOrder(assert.tradingAccount1, "BTC_EUR", "LIMIT", models.Sell, 10.00, 100.00, "GTC")

	//then:
	assert.Equal(1, GetSellBookOrderCount())
	assert.Equal(0, GetBuyBookOrderCount())
	assert.Equal(0, GetTradeCount())

	assert.Equal([]models.PriceVolume{{Price: 10.0, Volume: 100.0}}, GetVolumes("BTC_EUR", models.Sell))

	// when: incoming buy order is only partially matched
	ProcessTradeOrder(assert.tradingAccount2, "BTC_EUR", "LIMIT", models.Buy, 10.00, 50.00, "GTC")

	// then: the book should have one sell order and a single trade should be generated
	assert.Equal(1, GetSellBookOrderCount())
	assert.Equal(0, GetBuyBookOrderCount())
	assert.Equal(1, GetTradeCount())

	assert.Equal([]models.PriceVolume{{Price: 10.0, Volume: 50.0}}, GetVolumes("BTC_EUR", models.Sell))
	assert.Equal([]models.PriceVolume{}, GetVolumes("BTC_EUR", models.Buy))
}

func (assert *ServiceTestSuite) TestProcessLimitOverflowMatchIncomingBuySingleTrade() {
	// when: there is a SELL order in the book
	ProcessTradeOrder(assert.tradingAccount1, "BTC_EUR", "LIMIT", models.Sell, 10.00, 10.00, "GTC")

	// then:
	assert.Equal(1, GetSellBookOrderCount())
	assert.Equal(0, GetBuyBookOrderCount())
	assert.Equal(0, GetTradeCount())
	assert.Equal(10.0, GetAvailableLimitVolume(models.Sell, 10))

	// when: a BUY limit order arrives that crosses and is more that the book amount
	ProcessTradeOrder(assert.tradingAccount2, "BTC_EUR", "LIMIT", models.Buy, 10.00, 15, "GTC")

	// then: the book should be one buy order,  no sell orders and a one trade
	assert.Equal(0, GetSellBookOrderCount())
	assert.Equal(1, GetBuyBookOrderCount())
	assert.Equal(1, GetTradeCount())
	assert.Equal([]models.PriceVolume{}, GetVolumes("BTC_EUR", models.Sell))

	assert.Equal([]models.PriceVolume{{Price: 10.00, Volume: 5.0}}, GetVolumes("BTC_EUR", models.Buy))
}

func (assert *ServiceTestSuite) TestProcessLimitPartialMatchIncomingSellSingleTrade() {
	// when: there is a BUY order
	ProcessTradeOrder(assert.tradingAccount1, "BTC_EUR", "LIMIT", models.Buy, 10.00, 100.00, "GTC")

	// then:
	assert.Equal(0, GetSellBookOrderCount())
	assert.Equal(1, GetBuyBookOrderCount())
	assert.Equal(0, GetTradeCount())

	assert.Equal(100.0, GetAvailableLimitVolume(models.Buy, 10))

	// when: incoming SELL order is only partially matched
	ProcessTradeOrder(assert.tradingAccount2, "BTC_EUR", "LIMIT", models.Sell, 10.00, 50.00, "GTC")

	// then: the book should have one BUY order and a 1 trade should be generated
	assert.Equal(0, GetSellBookOrderCount())
	assert.Equal(1, GetBuyBookOrderCount())
	assert.Equal(1, GetTradeCount())

	assert.Equal(50.0, GetAvailableLimitVolume(models.Buy, 10))
	assert.Equal([]models.PriceVolume{}, GetVolumes("BTC_EUR", models.Sell))

	assert.Equal([]models.PriceVolume{{Price: 10.00, Volume: 50.0}}, GetVolumes("BTC_EUR", models.Buy))
}

func (assert *ServiceTestSuite) TestProcessLimitOverflowMatchIncomingSellSingleTrade() {
	// when: there is a BUY order in the book
	ProcessTradeOrder(assert.tradingAccount1, "BTC_EUR", "LIMIT", models.Buy, 10.00, 100.00, "GTC")

	// then:
	assert.Equal(0, GetSellBookOrderCount())
	assert.Equal(1, GetBuyBookOrderCount())
	assert.Equal(0, GetTradeCount())
	assert.Equal(100.0, GetAvailableLimitVolume(models.Buy, 10))

	// when: a SELL limit order arrives that crosses and is more that the book amount
	ProcessTradeOrder(assert.tradingAccount2, "BTC_EUR", "LIMIT", models.Sell, 10.00, 150.00, "GTC")

	// then: the book should be one SELL order,  no BUY orders and 1 trade
	assert.Equal(1, GetSellBookOrderCount())
	assert.Equal(0, GetBuyBookOrderCount())
	assert.Equal(1, GetTradeCount())

	assert.Equal(50.0, GetAvailableLimitVolume(models.Sell, 10))

	assert.Equal([]models.PriceVolume{{Price: 10.00, Volume: 50.0}}, GetVolumes("BTC_EUR", models.Sell))

	assert.Equal([]models.PriceVolume{}, GetVolumes("BTC_EUR", models.Buy))
}

func (assert *ServiceTestSuite) TestProcessLimitExactMatchIncomingBuysMultipleTrades() {
	// when: there is a SELL order in the book and 2 BUY limit order arrive that cross
	ProcessTradeOrder(assert.tradingAccount1, "BTC_EUR", "LIMIT", models.Sell, 10.00, 10.00, "GTC")

	// then:
	assert.Equal(1, GetSellBookOrderCount())
	assert.Equal(0, GetBuyBookOrderCount())
	assert.Equal(0, GetTradeCount())

	assert.Equal(10.0, GetAvailableLimitVolume(models.Sell, 10))

	// when: incoming buy order that are partially matched

	ProcessTradeOrder(assert.tradingAccount2, "BTC_EUR", "LIMIT", models.Buy, 10.00, 5.00, "GTC")
	ProcessTradeOrder(assert.tradingAccount2, "BTC_EUR", "LIMIT", models.Buy, 10.00, 5.00, "GTC")

	// then: the book should have 2 trades should be generated
	assert.Equal(0, GetSellBookOrderCount())
	assert.Equal(0, GetBuyBookOrderCount())
	assert.Equal(2, GetTradeCount())
	assert.Equal(0.0, GetAvailableLimitVolume(models.Sell, 10))
	assert.Equal([]models.PriceVolume{}, GetVolumes("BTC_EUR", models.Sell))
	assert.Equal([]models.PriceVolume{}, GetVolumes("BTC_EUR", models.Buy))
}

func (assert *ServiceTestSuite) TestProcessLimitExactMatchIncomingSellMultipleTrades() {

	// when: there is a BUY order in the book and 2 SELL limit order arrive that cross
	ProcessTradeOrder(assert.tradingAccount1, "BTC_EUR", "LIMIT", models.Buy, 10.00, 10.00, "GTC")

	// then:
	assert.Equal(0, GetSellBookOrderCount())
	assert.Equal(1, GetBuyBookOrderCount())
	assert.Equal(0, GetTradeCount())

	assert.Equal(10.0, GetAvailableLimitVolume(models.Buy, 10))

	// when: incoming buy order is only partially matched
	ProcessTradeOrder(assert.tradingAccount2, "BTC_EUR", "LIMIT", models.Sell, 10.00, 5.00, "GTC")
	ProcessTradeOrder(assert.tradingAccount2, "BTC_EUR", "LIMIT", models.Sell, 10.00, 5.00, "GTC")

	// then: the book should have 2 trades should be generated
	assert.Equal(0, GetSellBookOrderCount())
	assert.Equal(0, GetBuyBookOrderCount())
	assert.Equal(2, GetTradeCount())
	assert.Equal([]models.PriceVolume{}, GetVolumes("BTC_EUR", models.Sell))
	assert.Equal([]models.PriceVolume{}, GetVolumes("BTC_EUR", models.Buy))
}

func (assert *ServiceTestSuite) TestProcessLimitPartialMatchMultipleBookSellsToMultipleTrades() {
	// given:

	// when: there are multiple SELL orders in the book and BUY limit order arrive that cross but can fill only partially one of the orders
	ProcessTradeOrder(assert.tradingAccount1, "BTC_EUR", "LIMIT", models.Sell, 10.00, 50.00, "GTC")
	ProcessTradeOrder(assert.tradingAccount1, "BTC_EUR", "LIMIT", models.Sell, 10.00, 50.00, "GTC")

	// then:
	assert.Equal(2, GetSellBookOrderCount())
	assert.Equal(0, GetBuyBookOrderCount())
	assert.Equal(0, GetTradeCount())

	assert.Equal(100.0, GetAvailableLimitVolume(models.Sell, 10))

	// when: incoming buy order is only partially matched
	ProcessTradeOrder(assert.tradingAccount2, "BTC_EUR", "LIMIT", models.Buy, 10.00, 75.00, "GTC")

	// then: the book should have 2 trades should be generated
	assert.Equal(1, GetSellBookOrderCount())
	assert.Equal(0, GetBuyBookOrderCount())
	assert.Equal(2, GetTradeCount())

	assert.Equal([]models.PriceVolume{{Price: 10.00, Volume: 25.0}}, GetVolumes("BTC_EUR", models.Sell))
	assert.Equal([]models.PriceVolume{}, GetVolumes("BTC_EUR", models.Buy))
}

func (assert *ServiceTestSuite) TestProcessLimitPartialMatchMultipleBookBuysToMultipleTrades() {
	// given:

	// when: there are multiple BUY orders in the book and SELL limit order arrive that cross but can fill only partially one of the orders
	ProcessTradeOrder(assert.tradingAccount1, "BTC_EUR", "LIMIT", models.Buy, 10.00, 50.00, "GTC")
	ProcessTradeOrder(assert.tradingAccount1, "BTC_EUR", "LIMIT", models.Buy, 10.00, 50.00, "GTC")

	// then:
	assert.Equal(0, GetSellBookOrderCount())
	assert.Equal(2, GetBuyBookOrderCount())
	assert.Equal(0, GetTradeCount())
	assert.Equal(100.0, GetAvailableLimitVolume(models.Buy, 10))

	// when: incoming SELL order is only partially matched
	ProcessTradeOrder(assert.tradingAccount2, "BTC_EUR", "LIMIT", models.Sell, 10.00, 75.00, "GTC")

	// then: the book should have 2 trades should be generated
	assert.Equal(0, GetSellBookOrderCount())
	assert.Equal(1, GetBuyBookOrderCount())
	assert.Equal(2, GetTradeCount())

	assert.Equal([]models.PriceVolume{}, GetVolumes("BTC_EUR", models.Sell))

	assert.Equal([]models.PriceVolume{{Price: 10.00, Volume: 25.0}}, GetVolumes("BTC_EUR", models.Buy))
}

func (assert *ServiceTestSuite) TestProcessLimitExactMatchMultipleBookSellsToMultipleTrades() {
	// given:

	// when:
	ProcessTradeOrder(assert.tradingAccount1, "BTC_EUR", "LIMIT", models.Sell, 10.00, 10.00, "GTC")
	ProcessTradeOrder(assert.tradingAccount1, "BTC_EUR", "LIMIT", models.Sell, 10.00, 10.00, "GTC")
	ProcessTradeOrder(assert.tradingAccount1, "BTC_EUR", "LIMIT", models.Sell, 10.00, 10.00, "GTC")
	ProcessTradeOrder(assert.tradingAccount1, "BTC_EUR", "LIMIT", models.Sell, 10.00, 10.00, "GTC")

	// then:
	assert.Equal(4, GetSellBookOrderCount())
	assert.Equal(0, GetBuyBookOrderCount())
	assert.Equal(0, GetTradeCount())

	assert.Equal(40.0, GetAvailableLimitVolume(models.Sell, 10))

	// when: incoming buy order is only partially matched
	ProcessTradeOrder(assert.tradingAccount2, "BTC_EUR", "LIMIT", models.Buy, 10.00, 40.00, "GTC")

	// then: the book should have 2 trades should be generated
	assert.Equal(0, GetSellBookOrderCount())
	assert.Equal(0, GetBuyBookOrderCount())
	assert.Equal(4, GetTradeCount())
	assert.Equal([]models.PriceVolume{}, GetVolumes("BTC_EUR", models.Sell))
	assert.Equal([]models.PriceVolume{}, GetVolumes("BTC_EUR", models.Buy))
}

func (assert *ServiceTestSuite) TestProcessLimitExactMatchMultipleBookBuyToMultipleTrades() {
	// given:

	// when:
	ProcessTradeOrder(assert.tradingAccount1, "BTC_EUR", "LIMIT", models.Buy, 10.00, 5.00, "GTC")
	ProcessTradeOrder(assert.tradingAccount1, "BTC_EUR", "LIMIT", models.Buy, 10.00, 5.00, "GTC")
	ProcessTradeOrder(assert.tradingAccount1, "BTC_EUR", "LIMIT", models.Buy, 10.00, 5.00, "GTC")
	ProcessTradeOrder(assert.tradingAccount1, "BTC_EUR", "LIMIT", models.Buy, 10.00, 5.00, "GTC")

	// then:
	assert.Equal(0, GetSellBookOrderCount())
	assert.Equal(4, GetBuyBookOrderCount())
	assert.Equal(0, GetTradeCount())
	assert.Equal(20.0, GetAvailableLimitVolume(models.Buy, 10))

	// when: incoming buy order is only partially matched
	ProcessTradeOrder(assert.tradingAccount2, "BTC_EUR", "LIMIT", models.Sell, 10.00, 20.00, "GTC")

	// then: the book should have 2 trades should be generated
	assert.Equal(0, GetSellBookOrderCount())
	assert.Equal(0, GetBuyBookOrderCount())
	assert.Equal(4, GetTradeCount())
	assert.Equal([]models.PriceVolume{}, GetVolumes("BTC_EUR", models.Sell))
	assert.Equal([]models.PriceVolume{}, GetVolumes("BTC_EUR", models.Buy))
}

func (assert *ServiceTestSuite) TestProcessLimitIncompleteMatchMultipleBookSellsToMultipleTrades() {
	// given:

	// when: there are multiple SELL orders in the book and BUY limit order arrive that cross but can be only partially filled
	ProcessTradeOrder(assert.tradingAccount1, "BTC_EUR", "LIMIT", models.Sell, 10.00, 5.00, "GTC")
	ProcessTradeOrder(assert.tradingAccount1, "BTC_EUR", "LIMIT", models.Sell, 10.00, 5.00, "GTC")

	// then:
	assert.Equal(2, GetSellBookOrderCount())
	assert.Equal(0, GetBuyBookOrderCount())
	assert.Equal(0, GetTradeCount())
	assert.Equal(10.0, GetAvailableLimitVolume(models.Sell, 10))

	// when: incoming buy order is only partially matched
	ProcessTradeOrder(assert.tradingAccount2, "BTC_EUR", "LIMIT", models.Buy, 10.00, 17, "GTC")

	// then: the book should have 2 trades should be generated
	assert.Equal(0, GetSellBookOrderCount())
	assert.Equal(1, GetBuyBookOrderCount())
	assert.Equal(2, GetTradeCount())
	assert.Equal([]models.PriceVolume{}, GetVolumes("BTC_EUR", models.Sell))

	assert.Equal([]models.PriceVolume{{Price: 10.00, Volume: 7.0}}, GetVolumes("BTC_EUR", models.Buy))
}

func (assert *ServiceTestSuite) TestProcessLimitIncompleteMatchMultipleBookBuysToMultipleTrades() {
	// given:
	// when: there are multiple BUY orders in the book and SELL limit order arrive that cross but can be only partially filled
	ProcessTradeOrder(assert.tradingAccount1, "BTC_EUR", "LIMIT", models.Buy, 10.00, 50.00, "GTC")
	ProcessTradeOrder(assert.tradingAccount1, "BTC_EUR", "LIMIT", models.Buy, 10.00, 50.00, "GTC")

	// then:
	assert.Equal(0, GetSellBookOrderCount())
	assert.Equal(2, GetBuyBookOrderCount())
	assert.Equal(0, GetTradeCount())

	assert.Equal(100.0, GetAvailableLimitVolume(models.Buy, 10))

	// when: incoming SELL order is only partially matched
	ProcessTradeOrder(assert.tradingAccount2, "BTC_EUR", "LIMIT", models.Sell, 10.00, 175.00, "GTC")

	// then: the book should have 2 trades should be generated
	assert.Equal(1, GetSellBookOrderCount())
	assert.Equal(0, GetBuyBookOrderCount())
	assert.Equal(2, GetTradeCount())

	assert.Equal([]models.PriceVolume{{Price: 10.00, Volume: 75.0}}, GetVolumes("BTC_EUR", models.Sell))

	assert.Equal([]models.PriceVolume{}, GetVolumes("BTC_EUR", models.Buy))
}

func (assert *ServiceTestSuite) TestProcessLimitPartialMatchMultipleBookSellsToMultipleTradesMultiplePrices() {
	// given:

	// when: there are multiple SELL orders in the book and BUY limit order arrive that cross but can fill only partially one of the orders
	ProcessTradeOrder(assert.tradingAccount1, "BTC_EUR", "LIMIT", models.Sell, 9.00, 50.00, "GTC")
	ProcessTradeOrder(assert.tradingAccount1, "BTC_EUR", "LIMIT", models.Sell, 10.00, 50.00, "GTC")

	// then:
	assert.Equal(2, GetSellBookOrderCount())
	assert.Equal(0, GetBuyBookOrderCount())
	assert.Equal(0, GetTradeCount())

	// when: incoming buy order is only partially matched
	ProcessTradeOrder(assert.tradingAccount2, "BTC_EUR", "LIMIT", models.Buy, 11.00, 75.00, "GTC")

	// then: the book should have 2 trades should be generated
	assert.Equal(1, GetSellBookOrderCount())
	assert.Equal(0, GetBuyBookOrderCount())
	assert.Equal(2, GetTradeCount())

	assert.Equal([]models.PriceVolume{{Price: 10.00, Volume: 25.0}}, GetVolumes("BTC_EUR", models.Sell))

	assert.Equal([]models.PriceVolume{}, GetVolumes("BTC_EUR", models.Buy))
}

func (assert *ServiceTestSuite) TestProcessLimitPartialMatchMultipleBookBuysToMultipleTradesMultiplePrices() {
	// given:

	// when: there are multiple BUY orders in the book and SELL limit order arrive that cross but can fill only partially one of the orders
	ProcessTradeOrder(assert.tradingAccount1, "BTC_EUR", "LIMIT", models.Buy, 9.00, 50.00, "GTC")
	ProcessTradeOrder(assert.tradingAccount1, "BTC_EUR", "LIMIT", models.Buy, 10.00, 50.00, "GTC")

	// then:
	assert.Equal(0, GetSellBookOrderCount())
	assert.Equal(2, GetBuyBookOrderCount())
	assert.Equal(0, GetTradeCount())
	assert.Equal(50.0, GetAvailableLimitVolume(models.Buy, 10))
	assert.Equal(100.0, GetAvailableLimitVolume(models.Buy, 9))

	// when: incoming SELL order is only partially matched
	ProcessTradeOrder(assert.tradingAccount2, "BTC_EUR", "LIMIT", models.Sell, 8.00, 75.00, "GTC")

	// then: the book should have 2 trades should be generated
	assert.Equal(0, GetSellBookOrderCount())
	assert.Equal(1, GetBuyBookOrderCount())
	assert.Equal(2, GetTradeCount())
	assert.Equal(25.0, GetAvailableLimitVolume(models.Buy, 9))
	assert.Equal([]models.PriceVolume{}, GetVolumes("BTC_EUR", models.Sell))

	assert.Equal([]models.PriceVolume{{Price: 9.00, Volume: 25.0}}, GetVolumes("BTC_EUR", models.Buy))
}

func (assert *ServiceTestSuite) TestProcessLimitSelfTradePreventions() {
	// given:

	//when:
	ProcessTradeOrder(assert.tradingAccount1, "BTC_EUR", "LIMIT", models.Buy, 10.00, 50.00, "GTC")
	assert.Equal(1, GetBuyBookOrderCount())
	// then:
	ProcessTradeOrder(assert.tradingAccount1, "BTC_EUR", "LIMIT", models.Sell, 10.00, 50.00, "GTC")
	assert.Equal(1, GetSellBookOrderCount())
	assert.Equal(0, GetTradeCount())
	assert.Equal(50.0, GetAvailableLimitVolume(models.Buy, 10))
	assert.Equal(50.0, GetAvailableLimitVolume(models.Sell, 10))

}
