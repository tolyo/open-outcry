package services

import (
	"open-outcry/pkg/db"
	"open-outcry/pkg/models"
)

//  Primary market order test cases:
//    * An incoming market order hits an outstanding limit order in the opposite book
//    * An incoming limit order hits an outstanding market order in the opposite book
//    * An incoming market order hits another market order of the opposite side
//
//  Additional use cases to test:
//    * If the opposite book {es not have outstanding limit orders,
//      then the trade settles at the incoming order’s limit price
//    * If the opposite book {es have limit orders, the trade settles at the better of two prices –
//      the incoming order’s limit or the best limit from the opposite book –
//      the term “better of two prices” is from the point of view of the incoming limit order.
//      In other words, if incoming limit order would have crossed with outstanding opposite
//      “best limit” order in the absence of market order,
//      then trade at that, potentially improved, “best limit” price.
//
//    * when a market order matches with another market order in the opposite order book, there are three possibilites to test:
//      - the opposite order book with the resting market order also contains one or more outstanding limit orders –
//        in this case the opposite book has a “best limit” price and this price becomes the price for the trade
//      - the opposte order book {es not have outstanding limit orders, so the “best limit” price is not funcined.
//        In this case the trade occurs at the “reference price”. Most often reference price is the last traded price for a security.
//      - if no trades have occured (upon security launch), market orders simply rest

func (assert *ServiceTestSuite) TestProcessMarketSellOrderSave() {
	// when: a market order is sent to an empty matching unit
	ProcessTradeOrder(assert.tradingAccount1, "BTC_EUR", models.Market, models.Sell, 0, 100, "GTC")
	// then: a matching unit should save the trade order but it should not be visible to the order book
	assert.Equal(1, GetSellBookOrderCount())
	assert.Equal([]PriceVolume{}, GetVolumes("BTC_EUR", models.Sell))
}

func (assert *ServiceTestSuite) TestProcessMarketBuyOrderSave() {
	// when: a market order is sent to an empty matching unit
	ProcessTradeOrder(assert.tradingAccount1, "BTC_EUR", models.Market, models.Buy, 0, 100, "GTC")

	// then: a matching unit should save the trade order but it should not be visible to the order book
	assert.Equal(1, GetBuyBookOrderCount())
	assert.Equal([]PriceVolume{}, GetVolumes("BTC_EUR", models.Buy))
}

func (assert *ServiceTestSuite) TestProcessMarketSellOrderWithExistingSellLimit() {

	// when: a market order is sent to an non empty matching unit
	ProcessTradeOrder(assert.tradingAccount1, "BTC_EUR", "LIMIT", models.Sell, 10, 100, "GTC")
	ProcessTradeOrder(assert.tradingAccount1, "BTC_EUR", models.Market, models.Sell, 0, 100, "GTC")

	// then: a matching unit should save the trade order but it should not be visible to the order book
	assert.Equal(2, GetSellBookOrderCount())
	assert.Equal([]PriceVolume{{10, 100}}, GetVolumes("BTC_EUR", models.Sell))
}

func (assert *ServiceTestSuite) TestProcessMarketBuyOrderWithExistingBuyLimit() {

	// when: a market order is sent to an non empty matching unit
	ProcessTradeOrder(assert.tradingAccount1, "BTC_EUR", "LIMIT", models.Buy, 10, 10, "GTC")
	ProcessTradeOrder(assert.tradingAccount1, "BTC_EUR", models.Market, models.Buy, 0, 10, "GTC")

	// then: a matching unit should save the trade order but it should not be visible to the order book
	assert.Equal(2, GetBuyBookOrderCount())
	assert.Equal([]PriceVolume{{10.0, 10.0}}, GetVolumes("BTC_EUR", models.Buy))
}

func (assert *ServiceTestSuite) TestProcessMarketSellOrderWithExistingBuyLimit() {
	// when:
	ProcessTradeOrder(assert.tradingAccount1, "BTC_EUR", "LIMIT", models.Buy, 10, 100, "GTC")
	ProcessTradeOrder(assert.tradingAccount2, "BTC_EUR", models.Market, models.Sell, 0, 100, "GTC")

	// then:
	assert.Equal(1, GetTradeCount())
	assert.Equal([]float64{10.0}, GetTradePrices())
	assert.Equal(0, GetBuyBookOrderCount())
	assert.Equal(0, GetSellBookOrderCount())
	assert.Equal([]PriceVolume{}, GetVolumes("BTC_EUR", models.Buy))
	assert.Equal([]PriceVolume{}, GetVolumes("BTC_EUR", models.Sell))
}

func (assert *ServiceTestSuite) TestProcessMarketBuyOrderWithExistingSellLimit() {
	// when:
	ProcessTradeOrder(assert.tradingAccount1, "BTC_EUR", "LIMIT", models.Sell, 10, 100, "GTC")
	ProcessTradeOrder(assert.tradingAccount2, "BTC_EUR", models.Market, models.Buy, 0, 100, "GTC")
	// then:
	assert.Equal(1, GetTradeCount())
	assert.Equal([]float64{10.0}, GetTradePrices())
	assert.Equal(0, GetBuyBookOrderCount())
	assert.Equal(1, GetSellBookOrderCount())
	assert.Equal([]PriceVolume{}, GetVolumes("BTC_EUR", models.Buy))
	assert.Equal([]PriceVolume{{10.0, 90.0}}, GetVolumes("BTC_EUR", models.Sell))
}

func (assert *ServiceTestSuite) TestProcessMarketSellOrderWithMultipleExistingBuyLimits() {
	// when:
	ProcessTradeOrder(assert.tradingAccount1, "BTC_EUR", "LIMIT", models.Buy, 5, 10, "GTC")
	ProcessTradeOrder(assert.tradingAccount1, "BTC_EUR", "LIMIT", models.Buy, 4, 10, "GTC")
	ProcessTradeOrder(assert.tradingAccount1, "BTC_EUR", "LIMIT", models.Buy, 7, 10, "GTC")
	ProcessTradeOrder(assert.tradingAccount2, "BTC_EUR", models.Market, models.Sell, 0, 30, "GTC")

	// then:
	assert.Equal(3, GetTradeCount())
	assert.Equal([]float64{7.0, 5.0, 4.0}, GetTradePrices())
	assert.Equal(0, GetBuyBookOrderCount())
	assert.Equal(0, GetSellBookOrderCount())
	assert.Equal([]PriceVolume{}, GetVolumes("BTC_EUR", models.Buy))
	assert.Equal([]PriceVolume{}, GetVolumes("BTC_EUR", models.Sell))
}

func (assert *ServiceTestSuite) TestProcessMarketBuyOrderWithMultipleExistingSellLimits() {

	// when:
	ProcessTradeOrder(assert.tradingAccount1, "BTC_EUR", "LIMIT", models.Sell, 5, 10, "GTC")
	ProcessTradeOrder(assert.tradingAccount1, "BTC_EUR", "LIMIT", models.Sell, 4, 10, "GTC")
	ProcessTradeOrder(assert.tradingAccount1, "BTC_EUR", "LIMIT", models.Sell, 7, 10, "GTC")
	ProcessTradeOrder(assert.tradingAccount2, "BTC_EUR", models.Market, models.Buy, 0, 160, "GTC")

	// then:
	assert.Equal(3, GetTradeCount())
	assert.Equal([]float64{4.0, 5.0, 7.0}, GetTradePrices())
	assert.Equal(0, GetBuyBookOrderCount())
	assert.Equal(0, GetSellBookOrderCount())
	assert.Equal([]PriceVolume{}, GetVolumes("BTC_EUR", models.Buy))
	assert.Equal([]PriceVolume{}, GetVolumes("BTC_EUR", models.Sell))
}

func (assert *ServiceTestSuite) TestProcessLimitBuyOrderWithExistingMarketSell() {

	// when: the opposite book {es not have outstanding limit orders
	ProcessTradeOrder(assert.tradingAccount2, "BTC_EUR", models.Market, models.Sell, 0, 10, "GTC")
	ProcessTradeOrder(assert.tradingAccount1, "BTC_EUR", "LIMIT", models.Buy, 5, 10, "GTC")

	// then: the trade settles at the incoming order’s limit price
	assert.Equal(1, GetTradeCount())
	assert.Equal([]float64{5.0}, GetTradePrices())
	assert.Equal(0, GetBuyBookOrderCount())
	assert.Equal(0, GetSellBookOrderCount())
	assert.Equal([]PriceVolume{}, GetVolumes("BTC_EUR", models.Buy))
	assert.Equal([]PriceVolume{}, GetVolumes("BTC_EUR", models.Sell))
}

func (assert *ServiceTestSuite) TestProcessLimitSellOrderWithExistingMarketBuy() {

	// when: the opposite book {es not have outstanding limit orders
	ProcessTradeOrder(assert.tradingAccount2, "BTC_EUR", models.Market, models.Buy, 0, 10, "GTC")
	ProcessTradeOrder(assert.tradingAccount1, "BTC_EUR", "LIMIT", models.Sell, 5, 10, "GTC")

	// then: the trade settles at the incoming order’s limit price
	assert.Equal(1, GetTradeCount())
	assert.Equal([]float64{5.0}, GetTradePrices())
	assert.Equal(0, GetBuyBookOrderCount())
	assert.Equal(1, GetSellBookOrderCount())
	assert.Equal([]PriceVolume{}, GetVolumes("BTC_EUR", models.Buy))
	assert.Equal([]PriceVolume{{5.0, 8.0}}, GetVolumes("BTC_EUR", models.Sell))
}

func (assert *ServiceTestSuite) TestProcessLimitBuyOrderWithExistingMarketSellAndBetterLimitSell() {

	// when: the opposite book has a market and a limit order
	ProcessTradeOrder(assert.tradingAccount2, "BTC_EUR", "LIMIT", models.Sell, 5, 10, "GTC")
	ProcessTradeOrder(assert.tradingAccount2, "BTC_EUR", models.Market, models.Sell, 0, 10, "GTC")
	ProcessTradeOrder(assert.tradingAccount1, "BTC_EUR", "LIMIT", models.Buy, 6, 20, "GTC")

	// then: the trade settles at the book order’s limit price as if market order did not exist
	assert.Equal(2, GetTradeCount())
	assert.Equal([]float64{5.0, 5.0}, GetTradePrices())
	assert.Equal(0, GetBuyBookOrderCount())
	assert.Equal(0, GetSellBookOrderCount())
	assert.Equal([]PriceVolume{}, GetVolumes("BTC_EUR", models.Buy))
	assert.Equal([]PriceVolume{}, GetVolumes("BTC_EUR", models.Sell))
}

func (assert *ServiceTestSuite) TestProcessLimitSellOrderWithExistingMarketBuyAndBetterLimitBuy() {

	// when: the opposite book has a market and a limit order
	ProcessTradeOrder(assert.tradingAccount2, "BTC_EUR", "LIMIT", models.Buy, 5, 10, "GTC")
	ProcessTradeOrder(assert.tradingAccount2, "BTC_EUR", models.Market, models.Buy, 0, 50, "GTC")
	ProcessTradeOrder(assert.tradingAccount1, "BTC_EUR", "LIMIT", models.Sell, 4, 20, "GTC")

	// then: the trade settles at the book order’s limit price as if market order did not exist
	assert.Equal(2, GetTradeCount())
	assert.Equal([]float64{5.0, 5.0}, GetTradePrices())
	assert.Equal(0, GetBuyBookOrderCount())
	assert.Equal(0, GetSellBookOrderCount())
	assert.Equal([]PriceVolume{}, GetVolumes("BTC_EUR", models.Buy))
	assert.Equal([]PriceVolume{}, GetVolumes("BTC_EUR", models.Sell))
}

func (assert *ServiceTestSuite) TestProcessLimitBuyOrderWithExistingMarketSellAndWorseLimitSell() {

	// when: the opposite book has a market and a limit order
	ProcessTradeOrder(assert.tradingAccount2, "BTC_EUR", "LIMIT", models.Sell, 7, 10, "GTC")
	ProcessTradeOrder(assert.tradingAccount2, "BTC_EUR", models.Market, models.Sell, 0, 10, "GTC")
	ProcessTradeOrder(assert.tradingAccount1, "BTC_EUR", "LIMIT", models.Buy, 6, 10, "GTC")

	// then: the trade settles at the incoming order’s limit price
	assert.Equal(1, GetTradeCount())
	assert.Equal([]float64{6.0}, GetTradePrices())
	assert.Equal(0, GetBuyBookOrderCount())
	assert.Equal(1, GetSellBookOrderCount())
	assert.Equal([]PriceVolume{}, GetVolumes("BTC_EUR", models.Buy))
	assert.Equal([]PriceVolume{{7.0, 10.0}}, GetVolumes("BTC_EUR", models.Sell))
}

func (assert *ServiceTestSuite) TestProcessLimitSellOrderWithExistingMarketBuyAndWorseLimitBuy() {

	// when: the opposite book has a market and a limit order
	ProcessTradeOrder(assert.tradingAccount2, "BTC_EUR", "LIMIT", models.Buy, 4, 10, "GTC")
	ProcessTradeOrder(assert.tradingAccount2, "BTC_EUR", models.Market, models.Buy, 0, 100, "GTC")
	ProcessTradeOrder(assert.tradingAccount1, "BTC_EUR", "LIMIT", models.Sell, 5, 20, "GTC")

	// then: the trade settles at the incoming order’s limit price
	assert.Equal(1, GetTradeCount())
	assert.Equal([]float64{5.0}, GetTradePrices())
	assert.Equal(1, GetBuyBookOrderCount())
	assert.Equal(0, GetSellBookOrderCount())
	assert.Equal([]PriceVolume{{4.0, 10.0}}, GetVolumes("BTC_EUR", models.Buy))
	assert.Equal([]PriceVolume{}, GetVolumes("BTC_EUR", models.Sell))
}

func (assert *ServiceTestSuite) TestProcessMarketBuyOrderWithExistingMarketSellAndNoReferencePrice() {

	// when: there is not a trade with a reference price
	ProcessTradeOrder(assert.tradingAccount2, "BTC_EUR", models.Market, models.Sell, 0, 10, "GTC")
	ProcessTradeOrder(assert.tradingAccount1, "BTC_EUR", models.Market, models.Buy, 0, 50, "GTC")

	// then: the trade settles at a reference price
	assert.Equal(0, GetTradeCount())
	assert.Equal([]float64{}, GetTradePrices())
	assert.Equal(1, GetBuyBookOrderCount())
	assert.Equal(1, GetSellBookOrderCount())
	assert.Equal([]PriceVolume{}, GetVolumes("BTC_EUR", models.Buy))
	assert.Equal([]PriceVolume{}, GetVolumes("BTC_EUR", models.Sell))
}

func (assert *ServiceTestSuite) TestProcessMarketSellOrderWithExistingMarketBuyAndNoReferencePrice() {

	// when: there is not a trade with a reference price
	ProcessTradeOrder(assert.tradingAccount2, "BTC_EUR", models.Market, models.Buy, 0, 50, "GTC")
	ProcessTradeOrder(assert.tradingAccount1, "BTC_EUR", models.Market, models.Sell, 0, 10, "GTC")

	// then: the trade settles at a reference price
	assert.Equal(0, GetTradeCount())
	assert.Equal([]float64{}, GetTradePrices())
	assert.Equal(1, GetBuyBookOrderCount())
	assert.Equal(1, GetSellBookOrderCount())
	assert.Equal([]PriceVolume{}, GetVolumes("BTC_EUR", models.Buy))
	assert.Equal([]PriceVolume{}, GetVolumes("BTC_EUR", models.Sell))
}

func (assert *ServiceTestSuite) TestProcessMarketBuOrderWithExistingMarketSellAndLimitSell() {

	// when: the opposite book has a market and a limit order
	ProcessTradeOrder(assert.tradingAccount2, "BTC_EUR", "LIMIT", models.Sell, 5, 10, "GTC")
	ProcessTradeOrder(assert.tradingAccount2, "BTC_EUR", models.Market, models.Sell, 0, 10, "GTC")
	ProcessTradeOrder(assert.tradingAccount1, "BTC_EUR", models.Market, models.Buy, 0, 50, "GTC")

	// then: the trade settles at the book order’s limit price
	assert.Equal(1, GetTradeCount())
	assert.Equal([]float64{5.0}, GetTradePrices())
	assert.Equal(0, GetBuyBookOrderCount())
	assert.Equal(1, GetSellBookOrderCount())
	assert.Equal([]PriceVolume{}, GetVolumes("BTC_EUR", models.Buy))
	assert.Equal([]PriceVolume{{5.0, 10.0}}, GetVolumes("BTC_EUR", models.Sell))
}

func (assert *ServiceTestSuite) TestProcessMarketSellOrderWithExistingMarketBuyAndLimitBuy() {

	// when: the opposite book has a market and a limit order
	ProcessTradeOrder(assert.tradingAccount2, "BTC_EUR", "LIMIT", models.Buy, 5, 10, "GTC")
	ProcessTradeOrder(assert.tradingAccount2, "BTC_EUR", models.Market, models.Buy, 0, 50, "GTC")
	ProcessTradeOrder(assert.tradingAccount1, "BTC_EUR", models.Market, models.Sell, 0, 10, "GTC")

	// then: the trade settles at the book order’s limit price
	assert.Equal(1, GetTradeCount())
	assert.Equal([]float64{5.0}, GetTradePrices())
	assert.Equal(1, GetBuyBookOrderCount())
	assert.Equal(0, GetSellBookOrderCount())
	assert.Equal([]PriceVolume{{5.0, 10.0}}, GetVolumes("BTC_EUR", models.Buy))
	assert.Equal([]PriceVolume{}, GetVolumes("BTC_EUR", models.Sell))
}

func (assert *ServiceTestSuite) TestProcessMarketBuyOrderWithExistingMarketSellAnReferencePrice() {

	// when: there is a trade with a reference price
	ProcessTradeOrder(assert.tradingAccount2, "BTC_EUR", "LIMIT", models.Sell, 5, 1, "GTC")
	ProcessTradeOrder(assert.tradingAccount1, "BTC_EUR", "LIMIT", models.Buy, 5, 1, "GTC")
	ProcessTradeOrder(assert.tradingAccount2, "BTC_EUR", models.Market, models.Sell, 0, 10, "GTC")
	ProcessTradeOrder(assert.tradingAccount1, "BTC_EUR", models.Market, models.Buy, 0, 50, "GTC")

	// then: the trade settles at a reference price
	assert.Equal(2, GetTradeCount())
	assert.Equal([]float64{5.0, 5.0}, GetTradePrices())
	assert.Equal(0, GetBuyBookOrderCount())
	assert.Equal(0, GetSellBookOrderCount())
	assert.Equal([]PriceVolume{}, GetVolumes("BTC_EUR", models.Buy))
	assert.Equal([]PriceVolume{}, GetVolumes("BTC_EUR", models.Sell))
}

func (assert *ServiceTestSuite) TestProcessMarketSellOrderWithExistingMarketBuyAndReferencePrice() {

	// when: the opposite book has a market and a limit order
	ProcessTradeOrder(assert.tradingAccount2, "BTC_EUR", "LIMIT", models.Sell, 5, 1, "GTC")
	ProcessTradeOrder(assert.tradingAccount1, "BTC_EUR", "LIMIT", models.Buy, 5, 1, "GTC")
	ProcessTradeOrder(assert.tradingAccount2, "BTC_EUR", models.Market, models.Buy, 0, 50, "GTC")
	ProcessTradeOrder(assert.tradingAccount1, "BTC_EUR", models.Market, models.Sell, 0, 10, "GTC")

	// then: the trade settles at a reference price
	assert.Equal(2, GetTradeCount())
	assert.Equal([]float64{5.0, 5.0}, GetTradePrices())
	assert.Equal(0, GetBuyBookOrderCount())
	assert.Equal(0, GetSellBookOrderCount())
	assert.Equal([]PriceVolume{}, GetVolumes("BTC_EUR", models.Buy))
	assert.Equal([]PriceVolume{}, GetVolumes("BTC_EUR", models.Sell))
}

func (assert *ServiceTestSuite) TestProcessLimitBuyOrderWithExistingMarketSellAndReferencePrice() {

	// when: there is a trade with a reference price
	ProcessTradeOrder(assert.tradingAccount2, "BTC_EUR", "LIMIT", models.Sell, 5, 1, "GTC")
	ProcessTradeOrder(assert.tradingAccount1, "BTC_EUR", "LIMIT", models.Buy, 5, 1, "GTC")
	ProcessTradeOrder(assert.tradingAccount2, "BTC_EUR", models.Market, models.Sell, 0, 10, "GTC")
	ProcessTradeOrder(assert.tradingAccount1, "BTC_EUR", "LIMIT", models.Buy, 6, 10, "GTC")

	// then: the trade settles at a reference price
	assert.Equal(2, GetTradeCount())
	assert.Equal([]float64{5.0, 5.0}, GetTradePrices())
	assert.Equal(0, GetBuyBookOrderCount())
	assert.Equal(0, GetSellBookOrderCount())
	assert.Equal([]PriceVolume{}, GetVolumes("BTC_EUR", models.Buy))
	assert.Equal([]PriceVolume{}, GetVolumes("BTC_EUR", models.Sell))
}

// ---- TESTS BELOW REQUIRE VARIOUS HACKS TO SIMULATE TIMESTAMP INCREMENTS
func (assert *ServiceTestSuite) TestProcessMarketBuyOrderWithExistingMarketSellAndMultipleReferencePrices() {

	// when: there is a trade with a reference price
	ProcessTradeOrder(assert.tradingAccount2, "BTC_EUR", "LIMIT", models.Sell, 4, 1, "GTC")
	ProcessTradeOrder(assert.tradingAccount1, "BTC_EUR", "LIMIT", models.Buy, 4, 1, "GTC")

	db.Instance().Exec("UPDATE trade SET created_at = current_timestamp + interval '1 second' WHERE price = 4")

	ProcessTradeOrder(assert.tradingAccount2, "BTC_EUR", "LIMIT", models.Sell, 6, 1, "GTC")
	ProcessTradeOrder(assert.tradingAccount1, "BTC_EUR", "LIMIT", models.Buy, 6, 1, "GTC")

	db.Instance().Exec("UPDATE trade SET created_at = current_timestamp + interval '2 second' WHERE price = 6")

	ProcessTradeOrder(assert.tradingAccount2, "BTC_EUR", "LIMIT", models.Sell, 5, 1, "GTC")
	ProcessTradeOrder(assert.tradingAccount1, "BTC_EUR", "LIMIT", models.Buy, 5, 1, "GTC")

	db.Instance().Exec("UPDATE trade SET created_at = current_timestamp + interval '3 second' WHERE price = 5")

	ProcessTradeOrder(assert.tradingAccount2, "BTC_EUR", models.Market, models.Sell, 0, 10, "GTC")
	ProcessTradeOrder(assert.tradingAccount1, "BTC_EUR", models.Market, models.Buy, 0, 50, "GTC")

	// then: the trade settles at last reference price
	assert.Equal(4, GetTradeCount())
	assert.Equal(5.0, GetTradePrices()[0])
	assert.Equal(0, GetBuyBookOrderCount())
	assert.Equal(0, GetSellBookOrderCount())
	assert.Equal([]PriceVolume{}, GetVolumes("BTC_EUR", models.Buy))
	assert.Equal([]PriceVolume{}, GetVolumes("BTC_EUR", models.Sell))
}

func (assert *ServiceTestSuite) TestProcessMarketSellOrderWithExistingMarketBuyAndMultipleReferencePrices() {

	// when: the opposite book has a market and a limit order
	ProcessTradeOrder(assert.tradingAccount2, "BTC_EUR", "LIMIT", models.Sell, 4, 1, "GTC")
	ProcessTradeOrder(assert.tradingAccount1, "BTC_EUR", "LIMIT", models.Buy, 4, 1, "GTC")

	db.Instance().Exec("UPDATE trade SET created_at = current_timestamp + interval '1 second' WHERE price = 4")

	ProcessTradeOrder(assert.tradingAccount2, "BTC_EUR", "LIMIT", models.Sell, 6, 1, "GTC")
	ProcessTradeOrder(assert.tradingAccount1, "BTC_EUR", "LIMIT", models.Buy, 6, 1, "GTC")

	db.Instance().Exec("UPDATE trade SET created_at = current_timestamp + interval '2 second' WHERE price = 6")

	ProcessTradeOrder(assert.tradingAccount2, "BTC_EUR", "LIMIT", models.Sell, 5, 1, "GTC")
	ProcessTradeOrder(assert.tradingAccount1, "BTC_EUR", "LIMIT", models.Buy, 5, 1, "GTC")
	db.Instance().Exec("UPDATE trade SET created_at = current_timestamp + interval '3 second' WHERE price = 5")

	ProcessTradeOrder(assert.tradingAccount2, "BTC_EUR", models.Market, models.Buy, 0, 50, "GTC")
	ProcessTradeOrder(assert.tradingAccount1, "BTC_EUR", models.Market, models.Sell, 0, 10, "GTC")

	// then: the trade settles at last reference price
	assert.Equal(4, GetTradeCount())
	assert.Equal(5.0, GetTradePrices()[0])
	assert.Equal(0, GetBuyBookOrderCount())
	assert.Equal(0, GetSellBookOrderCount())
	assert.Equal([]PriceVolume{}, GetVolumes("BTC_EUR", models.Buy))
	assert.Equal([]PriceVolume{}, GetVolumes("BTC_EUR", models.Sell))
}
