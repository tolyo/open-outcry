package services

import (
	"open-outcry/pkg/db"
	order "open-outcry/pkg/models/order_book"
	tradeorder "open-outcry/pkg/models/trade_order"
)

func (assert *ServiceTestSuite) TestGetVolumeAtPrice() {
	ProcessTradeOrder(assert.instrumentAccount1, "BTC_EUR", "LIMIT", tradeorder.Sell, 10.6, 100, "GTC")
	assert.Equal(100.0, GetVolumeAtPrice("BTC_EUR", tradeorder.Sell, 10.6))
	assert.Equal(100.0, db.QueryVal[float64](`
	  SELECT SUM(volume)
		FROM price_level
		WHERE side = 'SELL'
		  AND instrument_id = (SELECT id FROM instrument WHERE name = 'BTC_EUR')
		  AND price =  10.6
	  `))
	ProcessTradeOrder(assert.instrumentAccount2, "BTC_EUR", "LIMIT", tradeorder.Buy, 9.5, 100, "GTC")
	assert.Equal(100.0, GetVolumeAtPrice("BTC_EUR", tradeorder.Buy, 9.5))
	assert.Equal(100.00, db.QueryVal[float64](`
	            SELECT SUM(volume)
	              FROM price_level
	              WHERE side = 'BUY'
	                AND instrument_id = (SELECT id FROM instrument WHERE name = 'BTC_EUR')
	                AND price =  9.5
	`))
}

func (assert *ServiceTestSuite) TestGetVolumeSellSide() {
	ProcessTradeOrder(assert.instrumentAccount1, "BTC_EUR", "LIMIT", tradeorder.Sell, 10.7, 100, "GTC")
	ProcessTradeOrder(assert.instrumentAccount1, "BTC_EUR", "LIMIT", tradeorder.Sell, 10.6, 100, "GTC")
	ProcessTradeOrder(assert.instrumentAccount1, "BTC_EUR", "LIMIT", tradeorder.Sell, 10.7, 100, "GTC")
	ProcessTradeOrder(assert.instrumentAccount1, "BTC_EUR", "LIMIT", tradeorder.Sell, 10.4, 100, "GTC")
	assert.Equal([]order.PriceVolume{{Price: 10.4, Volume: 100}, {Price: 10.6, Volume: 100}, {Price: 10.7, Volume: 200}}, GetVolumes("BTC_EUR", tradeorder.Sell))
}

func (assert *ServiceTestSuite) TestGetVolumeBuySide() {
	ProcessTradeOrder(assert.instrumentAccount1, "BTC_EUR", "LIMIT", tradeorder.Buy, 1.7, 10, "GTC")
	ProcessTradeOrder(assert.instrumentAccount1, "BTC_EUR", "LIMIT", tradeorder.Buy, 1.6, 10, "GTC")
	ProcessTradeOrder(assert.instrumentAccount1, "BTC_EUR", "LIMIT", tradeorder.Buy, 1.7, 10, "GTC")
	ProcessTradeOrder(assert.instrumentAccount1, "BTC_EUR", "LIMIT", tradeorder.Buy, 1.4, 10, "GTC")
	assert.Equal([]order.PriceVolume{{Price: 1.7, Volume: 20}, {Price: 1.6, Volume: 10}, {Price: 1.4, Volume: 10}}, GetVolumes("BTC_EUR", tradeorder.Buy))
}

func (assert *ServiceTestSuite) TestGetOrderBook() {
	res := GetOrderBook("BTC_EUR")
	assert.Len(res.BuySide, 0)
	assert.Len(res.BuySide, 0)
	ProcessTradeOrder(assert.instrumentAccount1, "BTC_EUR", "LIMIT", tradeorder.Buy, 1.7, 10, "GTC")
	ProcessTradeOrder(assert.instrumentAccount1, "BTC_EUR", "LIMIT", tradeorder.Buy, 1.6, 10, "GTC")
	ProcessTradeOrder(assert.instrumentAccount1, "BTC_EUR", "LIMIT", tradeorder.Buy, 1.7, 10, "GTC")
	ProcessTradeOrder(assert.instrumentAccount1, "BTC_EUR", "LIMIT", tradeorder.Buy, 1.4, 10, "GTC")
	assert.Equal(order.OrderBook{BuySide: []order.PriceVolume{{Price: 1.7, Volume: 20}, {Price: 1.6, Volume: 10}, {Price: 1.4, Volume: 10}}, SellSide: []order.PriceVolume{}}, GetOrderBook("BTC_EUR"))
	ProcessTradeOrder(assert.instrumentAccount1, "BTC_EUR", "LIMIT", tradeorder.Sell, 10.7, 100, "GTC")
	ProcessTradeOrder(assert.instrumentAccount1, "BTC_EUR", "LIMIT", tradeorder.Sell, 10.6, 100, "GTC")
	ProcessTradeOrder(assert.instrumentAccount1, "BTC_EUR", "LIMIT", tradeorder.Sell, 10.7, 100, "GTC")
	ProcessTradeOrder(assert.instrumentAccount1, "BTC_EUR", "LIMIT", tradeorder.Sell, 10.4, 100, "GTC")
	assert.Equal(order.OrderBook{BuySide: []order.PriceVolume{{Price: 1.7, Volume: 20}, {Price: 1.6, Volume: 10}, {Price: 1.4, Volume: 10}}, SellSide: []order.PriceVolume{{Price: 10.4, Volume: 100}, {Price: 10.6, Volume: 100}, {Price: 10.7, Volume: 200}}}, GetOrderBook("BTC_EUR"))
}
