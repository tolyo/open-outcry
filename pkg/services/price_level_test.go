package services

import (
	"open-outcry/pkg/db"
	"open-outcry/pkg/models"
	"open-outcry/pkg/utils"
)

func (assert *ServiceTestSuite) TestCreatePriceLevel() {
	// given:
	// when given a new saved limit order
	_, err := ProcessTradeOrder(
		assert.tradingAccount1,
		"BTC_EUR",
		"LIMIT",
		models.Buy,
		10.00,
		10.00,
		models.GTC,
	)

	assert.Nil(err)
	// then a price level is created
	assert.Equal(1, utils.GetCount("price_level"))
	assert.Equal(10.0, db.QueryVal[float64]("SELECT volume FROM price_level LIMIT 1"))

	// when give another order for smaller amount
	ProcessTradeOrder(assert.tradingAccount1,
		"BTC_EUR",
		"LIMIT",
		models.Buy,
		10.00,
		5,
		models.GTC,
	)

	// then price level is updated
	assert.Equal(1, utils.GetCount("price_level"))
	assert.Equal(15.0, db.QueryVal[float64]("SELECT volume FROM price_level LIMIT 1"))

	// when give another order for different price
	ProcessTradeOrder(assert.tradingAccount1,
		"BTC_EUR",
		"LIMIT", models.Buy,
		5.00,
		5,
		models.GTC,
	)

	// then another price level is created
	assert.Equal(2, utils.GetCount("price_level"))
}

func (assert *ServiceTestSuite) TestCancelWithSingle() {
	// when given a new saved limit order

	id, _ := ProcessTradeOrder(
		assert.tradingAccount1,
		"BTC_EUR",
		"LIMIT",
		models.Buy,
		10.00,
		10.00,
		models.GTC,
	)

	// then a price level is created
	assert.Equal(1, utils.GetCount("price_level"))
	assert.Equal(10.0, db.QueryVal[float64]("SELECT volume FROM price_level LIMIT 1"))

	// when the order is deleted
	CancelTradeOrder(id)

	// then price level is deleted also
	assert.Equal(0, utils.GetCount("price_level"))
}

func (assert *ServiceTestSuite) TestCancelWithTwoOrdersOfSameSize() {
	// when given a new saved limit order

	id, _ := ProcessTradeOrder(
		assert.tradingAccount1,
		"BTC_EUR",
		"LIMIT",
		models.Buy,
		10.00,
		10.00,
		models.GTC,
	)

	ProcessTradeOrder(
		assert.tradingAccount1,
		"BTC_EUR",
		"LIMIT",
		models.Buy,
		10.00,
		10.00,
		models.GTC,
	)

	// then a price level is created
	assert.Equal(1, utils.GetCount("price_level"))
	assert.Equal(20.0, db.QueryVal[float64]("SELECT volume FROM price_level LIMIT 1"))

	// when the order is deleted
	CancelTradeOrder(id)

	// then price level updated
	assert.Equal(1, utils.GetCount("price_level"))
	assert.Equal(10.0, db.QueryVal[float64]("SELECT volume FROM price_level LIMIT 1"))
}

func (assert *ServiceTestSuite) TestCancelWithTwoOrdersWithDiffPrice() {
	// when given a new saved limit order

	id, _ := ProcessTradeOrder(
		assert.tradingAccount1,
		"BTC_EUR",
		"LIMIT",
		models.Buy,
		20.00,
		10.00,
		models.GTC,
	)

	ProcessTradeOrder(
		assert.tradingAccount1,
		"BTC_EUR",
		"LIMIT",
		models.Buy,
		10.00,
		10.00,
		models.GTC,
	)

	// then a price level is created
	assert.Equal(2, utils.GetCount("price_level"))

	// when the order is deleted
	CancelTradeOrder(id)

	// then price levels are updated
	assert.Equal(1, utils.GetCount("price_level"))
	assert.Equal(10.0, db.QueryVal[float64]("SELECT volume FROM price_level LIMIT 1"))
}
