package services

import (
	"open-outcry/pkg/db"
	"open-outcry/pkg/models"
)

func (assert *ServiceTestSuite) TestGetVolumeAtPrice() {
	// when: a single sell limit order is added to the order book
	ProcessTradeOrder(assert.tradingAccount1, "BTC_EUR", "LIMIT", models.Sell, 10.6, 100, "GTC")

	// then:
	assert.Equal(100.0, GetVolumeAtPrice("BTC_EUR", models.Sell, 10.6))

	assert.Equal(100.0, db.QueryVal[float64](`
	  SELECT SUM(volume)
		FROM price_level
		WHERE side = 'SELL'
		  AND instrument_id = (SELECT id FROM instrument WHERE name = 'BTC_EUR')
		  AND price =  10.6
	  `))

	// when: a single buy limit order is added to the order book
	ProcessTradeOrder(assert.tradingAccount2, "BTC_EUR", "LIMIT", models.Buy, 9.5, 100, "GTC")

	// then:
	assert.Equal(100.0, GetVolumeAtPrice("BTC_EUR", models.Buy, 9.5))

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
	ProcessTradeOrder(assert.tradingAccount1, "BTC_EUR", "LIMIT", models.Sell, 10.7, 100, "GTC")
	ProcessTradeOrder(assert.tradingAccount1, "BTC_EUR", "LIMIT", models.Sell, 10.6, 100, "GTC")
	ProcessTradeOrder(assert.tradingAccount1, "BTC_EUR", "LIMIT", models.Sell, 10.7, 100, "GTC")
	ProcessTradeOrder(assert.tradingAccount1, "BTC_EUR", "LIMIT", models.Sell, 10.4, 100, "GTC")

	// then should be sorted with cheapest orders first
	assert.Equal([]models.PriceVolume{
		{Price: 10.4, Volume: 100},
		{Price: 10.6, Volume: 100},
		{Price: 10.7, Volume: 200},
	}, GetVolumes("BTC_EUR", models.Sell))
}

func (assert *ServiceTestSuite) TestGetVolumeBuySide() {

	// when
	ProcessTradeOrder(assert.tradingAccount1, "BTC_EUR", "LIMIT", models.Buy, 1.7, 10, "GTC")
	ProcessTradeOrder(assert.tradingAccount1, "BTC_EUR", "LIMIT", models.Buy, 1.6, 10, "GTC")
	ProcessTradeOrder(assert.tradingAccount1, "BTC_EUR", "LIMIT", models.Buy, 1.7, 10, "GTC")
	ProcessTradeOrder(assert.tradingAccount1, "BTC_EUR", "LIMIT", models.Buy, 1.4, 10, "GTC")

	// then should be sorted with most expensive orders first
	assert.Equal([]models.PriceVolume{
		{Price: 1.7, Volume: 20},
		{Price: 1.6, Volume: 10},
		{Price: 1.4, Volume: 10},
	}, GetVolumes("BTC_EUR", models.Buy))

}

func (assert *ServiceTestSuite) TestGetOrderBook() {

	res := GetOrderBook("BTC_EUR")
	assert.Len(res.BuySide, 0)
	assert.Len(res.BuySide, 0)

	// when
	ProcessTradeOrder(assert.tradingAccount1, "BTC_EUR", "LIMIT", models.Buy, 1.7, 10, "GTC")
	ProcessTradeOrder(assert.tradingAccount1, "BTC_EUR", "LIMIT", models.Buy, 1.6, 10, "GTC")
	ProcessTradeOrder(assert.tradingAccount1, "BTC_EUR", "LIMIT", models.Buy, 1.7, 10, "GTC")
	ProcessTradeOrder(assert.tradingAccount1, "BTC_EUR", "LIMIT", models.Buy, 1.4, 10, "GTC")

	// then should be sorted with most expensive orders first
	assert.Equal(models.OrderBook{
		BuySide: []models.PriceVolume{
			{Price: 1.7, Volume: 20},
			{Price: 1.6, Volume: 10},
			{Price: 1.4, Volume: 10},
		},
		SellSide: []models.PriceVolume{},
	}, GetOrderBook("BTC_EUR"))

	ProcessTradeOrder(assert.tradingAccount1, "BTC_EUR", "LIMIT", models.Sell, 10.7, 100, "GTC")
	ProcessTradeOrder(assert.tradingAccount1, "BTC_EUR", "LIMIT", models.Sell, 10.6, 100, "GTC")
	ProcessTradeOrder(assert.tradingAccount1, "BTC_EUR", "LIMIT", models.Sell, 10.7, 100, "GTC")
	ProcessTradeOrder(assert.tradingAccount1, "BTC_EUR", "LIMIT", models.Sell, 10.4, 100, "GTC")

	assert.Equal(models.OrderBook{
		BuySide: []models.PriceVolume{
			{Price: 1.7, Volume: 20},
			{Price: 1.6, Volume: 10},
			{Price: 1.4, Volume: 10},
		},
		SellSide: []models.PriceVolume{
			{Price: 10.4, Volume: 100},
			{Price: 10.6, Volume: 100},
			{Price: 10.7, Volume: 200},
		},
	}, GetOrderBook("BTC_EUR"))
}
