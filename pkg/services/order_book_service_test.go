package services

import (
	"open-outcry/pkg/db"
)

func (assert *ServiceTestSuite) TestGetVolumeAtPrice() {
	// when: a single sell limit order is added to the order book
	ProcessTradeOrder(assert.instrumentAccount1, "BTC_EUR", "LIMIT", Sell, 10.6, 100, "GTC")

	// then:
	assert.Equal(100.0, GetVolumeAtPrice("BTC_EUR", Sell, 10.6))

	assert.Equal(100.0, db.QueryVal[float64](`
	  SELECT SUM(volume)
		FROM price_level
		WHERE side = 'SELL'
		  AND instrument_id = (SELECT id FROM instrument WHERE name = 'BTC_EUR')
		  AND price =  10.6
	  `))

	// when: a single buy limit order is added to the order book
	ProcessTradeOrder(assert.instrumentAccount2, "BTC_EUR", "LIMIT", Buy, 9.5, 100, "GTC")

	// then:
	assert.Equal(100.0, GetVolumeAtPrice("BTC_EUR", Buy, 9.5))

	assert.Equal(100.00, db.QueryVal[float64](`
	            SELECT SUM(volume)
	              FROM price_level
	              WHERE side = 'BUY'
	                AND instrument_id = (SELECT id FROM instrument WHERE name = 'BTC_EUR')
	                AND price =  9.5
	`))
}

func (assert *ServiceTestSuite) TestGetVolumeSellSide() {

	// when
	ProcessTradeOrder(assert.instrumentAccount1, "BTC_EUR", "LIMIT", Sell, 10.7, 100, "GTC")
	ProcessTradeOrder(assert.instrumentAccount1, "BTC_EUR", "LIMIT", Sell, 10.6, 100, "GTC")
	ProcessTradeOrder(assert.instrumentAccount1, "BTC_EUR", "LIMIT", Sell, 10.7, 100, "GTC")
	ProcessTradeOrder(assert.instrumentAccount1, "BTC_EUR", "LIMIT", Sell, 10.4, 100, "GTC")

	// then should be sorted with cheapest orders first
	assert.Equal([]PriceVolume{
		{Price: 10.4, Volume: 100},
		{Price: 10.6, Volume: 100},
		{Price: 10.7, Volume: 200},
	}, GetVolumes("BTC_EUR", Sell))
}

func (assert *ServiceTestSuite) TestGetVolumeBuySide() {

	// when
	ProcessTradeOrder(assert.instrumentAccount1, "BTC_EUR", "LIMIT", Buy, 1.7, 10, "GTC")
	ProcessTradeOrder(assert.instrumentAccount1, "BTC_EUR", "LIMIT", Buy, 1.6, 10, "GTC")
	ProcessTradeOrder(assert.instrumentAccount1, "BTC_EUR", "LIMIT", Buy, 1.7, 10, "GTC")
	ProcessTradeOrder(assert.instrumentAccount1, "BTC_EUR", "LIMIT", Buy, 1.4, 10, "GTC")

	// then should be sorted with most expensive orders first
	assert.Equal([]PriceVolume{
		{Price: 1.7, Volume: 20},
		{Price: 1.6, Volume: 10},
		{Price: 1.4, Volume: 10},
	}, GetVolumes("BTC_EUR", Buy))

}

func (assert *ServiceTestSuite) TestGetOrderBook() {

	res := GetOrderBook("BTC_EUR")
	assert.Len(res.BuySide, 0)
	assert.Len(res.BuySide, 0)

	// when
	ProcessTradeOrder(assert.instrumentAccount1, "BTC_EUR", "LIMIT", Buy, 1.7, 10, "GTC")
	ProcessTradeOrder(assert.instrumentAccount1, "BTC_EUR", "LIMIT", Buy, 1.6, 10, "GTC")
	ProcessTradeOrder(assert.instrumentAccount1, "BTC_EUR", "LIMIT", Buy, 1.7, 10, "GTC")
	ProcessTradeOrder(assert.instrumentAccount1, "BTC_EUR", "LIMIT", Buy, 1.4, 10, "GTC")

	// then should be sorted with most expensive orders first
	assert.Equal(OrderBook{
		BuySide: []PriceVolume{
			{Price: 1.7, Volume: 20},
			{Price: 1.6, Volume: 10},
			{Price: 1.4, Volume: 10},
		},
		SellSide: []PriceVolume{},
	}, GetOrderBook("BTC_EUR"))

	ProcessTradeOrder(assert.instrumentAccount1, "BTC_EUR", "LIMIT", Sell, 10.7, 100, "GTC")
	ProcessTradeOrder(assert.instrumentAccount1, "BTC_EUR", "LIMIT", Sell, 10.6, 100, "GTC")
	ProcessTradeOrder(assert.instrumentAccount1, "BTC_EUR", "LIMIT", Sell, 10.7, 100, "GTC")
	ProcessTradeOrder(assert.instrumentAccount1, "BTC_EUR", "LIMIT", Sell, 10.4, 100, "GTC")

	assert.Equal(OrderBook{
		BuySide: []PriceVolume{
			{Price: 1.7, Volume: 20},
			{Price: 1.6, Volume: 10},
			{Price: 1.4, Volume: 10},
		},
		SellSide: []PriceVolume{
			{Price: 10.4, Volume: 100},
			{Price: 10.6, Volume: 100},
			{Price: 10.7, Volume: 200},
		},
	}, GetOrderBook("BTC_EUR"))
}
