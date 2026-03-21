package services

import (
	"open-outcry/pkg/db"
	"open-outcry/pkg/models/trade_order"
)

func (assert *ServiceTestSuite) TestCreatePriceLevel() {
	_, err := ProcessTradeOrder(assert.instrumentAccount1, "BTC_EUR", "LIMIT", tradeorder.Buy, 10.00, 10.00, tradeorder.GTC)
	assert.Nil(err)
	assert.Equal(1, db.GetCount("price_level"))
	assert.Equal(10.0, db.QueryVal[float64]("SELECT volume FROM price_level LIMIT 1"))
	ProcessTradeOrder(assert.instrumentAccount1, "BTC_EUR", "LIMIT", tradeorder.Buy, 10.00, 5, tradeorder.GTC)
	assert.Equal(1, db.GetCount("price_level"))
	assert.Equal(15.0, db.QueryVal[float64]("SELECT volume FROM price_level LIMIT 1"))
	ProcessTradeOrder(assert.instrumentAccount1, "BTC_EUR", "LIMIT", tradeorder.Buy, 5.00, 5, tradeorder.GTC)
	assert.Equal(2, db.GetCount("price_level"))
}

func (assert *ServiceTestSuite) TestCancelWithSingle() {
	id, _ := ProcessTradeOrder(assert.instrumentAccount1, "BTC_EUR", "LIMIT", tradeorder.Buy, 10.00, 10.00, tradeorder.GTC)
	assert.Equal(1, db.GetCount("price_level"))
	assert.Equal(10.0, db.QueryVal[float64]("SELECT volume FROM price_level LIMIT 1"))
	CancelTradeOrder(id)
	assert.Equal(0, db.GetCount("price_level"))
}

func (assert *ServiceTestSuite) TestCancelWithTwoOrdersOfSameSize() {
	id, _ := ProcessTradeOrder(assert.instrumentAccount1, "BTC_EUR", "LIMIT", tradeorder.Buy, 10.00, 10.00, tradeorder.GTC)
	ProcessTradeOrder(assert.instrumentAccount1, "BTC_EUR", "LIMIT", tradeorder.Buy, 10.00, 10.00, tradeorder.GTC)
	assert.Equal(1, db.GetCount("price_level"))
	assert.Equal(20.0, db.QueryVal[float64]("SELECT volume FROM price_level LIMIT 1"))
	CancelTradeOrder(id)
	assert.Equal(1, db.GetCount("price_level"))
	assert.Equal(10.0, db.QueryVal[float64]("SELECT volume FROM price_level LIMIT 1"))
}

func (assert *ServiceTestSuite) TestCancelWithTwoOrdersWithDiffPrice() {
	id, _ := ProcessTradeOrder(assert.instrumentAccount1, "BTC_EUR", "LIMIT", tradeorder.Buy, 20.00, 10.00, tradeorder.GTC)
	ProcessTradeOrder(assert.instrumentAccount1, "BTC_EUR", "LIMIT", tradeorder.Buy, 10.00, 10.00, tradeorder.GTC)
	assert.Equal(2, db.GetCount("price_level"))
	CancelTradeOrder(id)
	assert.Equal(1, db.GetCount("price_level"))
	assert.Equal(10.0, db.QueryVal[float64]("SELECT volume FROM price_level LIMIT 1"))
}
