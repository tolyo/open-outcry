package services

import "open-outcry/pkg/models"

func (assert *ServiceTestSuite) TestCreateTradeOrderBook() {
	// given

	// when given a new limit order
	ProcessTradeOrder(assert.tradingAccount1,
		"BTC_EUR",
		"LIMIT", models.Sell,
		20.10,
		10, "GTC")

	// then should be saved
	assert.Equal(1, GetSellBookOrderCount())

	// when given a new market order
	ProcessTradeOrder(assert.tradingAccount1,
		"BTC_EUR", "MARKET", models.Sell, 0, 10, "GTC")

	// then should be saved
	assert.Equal(2, GetSellBookOrderCount())

	// when given a stop loss order
	ProcessTradeOrder(assert.tradingAccount1,
		"BTC_EUR", "STOPLOSS", models.Sell, 20.10, 10, "GTC")

	// then should be not be saved to order book
	assert.Equal(2, GetSellBookOrderCount())

	// when given a stop limit order
	ProcessTradeOrder(assert.tradingAccount1, "BTC_EUR", "STOPLIMIT", models.Sell, 20.10, 10, "GTC")

	// then should be not be saved to order book
	assert.Equal(2, GetSellBookOrderCount())
}
