package services

import (
	"open-outcry/pkg/models"
)

func (assert *ServiceTestSuite) TestGetTradeOrdersByTradingAccount() {
	// Place an order
	orderId, err := ProcessTradeOrder(assert.tradingAccount1, "BTC_EUR", "LIMIT", models.Sell, 10.6, 100, "GTC")
	assert.Nil(err)

	// Fetch orders for this trading account
	orders := models.GetTradeOrdersByTradingAccount(assert.tradingAccount1)
	assert.GreaterOrEqual(len(orders), 1)
	assert.Equal(orderId, orders[0].Id)
	assert.Equal(assert.tradingAccount1, orders[0].TradingAccountId)
	assert.Equal(models.InstrumentName("BTC_EUR"), orders[0].InstrumentName)
	assert.Equal(models.Sell, orders[0].Side)
	assert.Equal(models.Limit, orders[0].Type)

	// Other account should have no orders
	orders2 := models.GetTradeOrdersByTradingAccount(assert.tradingAccount2)
	assert.Equal(0, len(orders2))
}

func (assert *ServiceTestSuite) TestGetBookOrdersByTradingAccount() {
	// Place a limit order (goes to book)
	ProcessTradeOrder(assert.tradingAccount1, "BTC_EUR", "LIMIT", models.Sell, 10.6, 100, "GTC")

	bookOrders := models.GetBookOrdersByTradingAccount(assert.tradingAccount1)
	assert.Equal(1, len(bookOrders))
	assert.Equal(models.Sell, bookOrders[0].Side)

	// Other account has no book orders
	bookOrders2 := models.GetBookOrdersByTradingAccount(assert.tradingAccount2)
	assert.Equal(0, len(bookOrders2))
}

func (assert *ServiceTestSuite) TestGetTradesByTradingAccount() {
	// Create a matching pair that produces a trade
	ProcessTradeOrder(assert.tradingAccount1, "BTC_EUR", "LIMIT", models.Sell, 10.0, 100, "GTC")
	ProcessTradeOrder(assert.tradingAccount2, "BTC_EUR", "LIMIT", models.Buy, 10.0, 100, "GTC")

	// Both accounts should see the trade
	trades1 := models.GetTradesByTradingAccount(assert.tradingAccount1)
	assert.GreaterOrEqual(len(trades1), 1)

	trades2 := models.GetTradesByTradingAccount(assert.tradingAccount2)
	assert.GreaterOrEqual(len(trades2), 1)

	// Same trade
	assert.Equal(trades1[0].Id, trades2[0].Id)
	assert.Equal(10.0, trades1[0].Price)
	assert.Equal(100.0, trades1[0].Amount)
}

func (assert *ServiceTestSuite) TestGetTrade() {
	// Create a trade
	ProcessTradeOrder(assert.tradingAccount1, "BTC_EUR", "LIMIT", models.Sell, 10.0, 100, "GTC")
	ProcessTradeOrder(assert.tradingAccount2, "BTC_EUR", "LIMIT", models.Buy, 10.0, 100, "GTC")

	trades := models.GetTradesByTradingAccount(assert.tradingAccount1)
	assert.GreaterOrEqual(len(trades), 1)

	// Fetch by id
	trade := models.GetTrade(trades[0].Id)
	assert.NotNil(trade)
	assert.Equal(trades[0].Id, trade.Id)
	assert.Equal(models.InstrumentName("BTC_EUR"), trade.InstrumentName)
}

func (assert *ServiceTestSuite) TestGetTradeNotFound() {
	trade := models.GetTrade("nonexistent-id")
	assert.Nil(trade)
}

func (assert *ServiceTestSuite) TestGetAppEntities() {
	entities := models.GetAppEntities()
	// MASTER + 2 test entities
	assert.GreaterOrEqual(len(entities), 3)
}

func (assert *ServiceTestSuite) TestGetAppEntity() {
	entity := models.GetAppEntity(assert.appEntity1)
	assert.NotNil(entity)
	assert.Equal(assert.appEntity1, entity.Id)
}

func (assert *ServiceTestSuite) TestGetAppEntityNotFound() {
	entity := models.GetAppEntity("nonexistent-id")
	assert.Nil(entity)
}

func (assert *ServiceTestSuite) TestGetPaymentsByAppEntity() {
	// SetupTest creates deposits, so there should be payments
	payments := models.GetPaymentsByAppEntity(assert.appEntity1)
	assert.GreaterOrEqual(len(payments), 1)
	assert.NotEmpty(payments[0].Id)
	assert.NotEmpty(payments[0].Currency)
}

func (assert *ServiceTestSuite) TestGetPayment() {
	payments := models.GetPaymentsByAppEntity(assert.appEntity1)
	assert.GreaterOrEqual(len(payments), 1)

	payment := models.GetPayment(payments[0].Id)
	assert.NotNil(payment)
	assert.Equal(payments[0].Id, payment.Id)
	assert.Equal(payments[0].Currency, payment.Currency)
}

func (assert *ServiceTestSuite) TestGetPaymentNotFound() {
	payment := models.GetPayment("nonexistent-id")
	assert.Nil(payment)
}

func (assert *ServiceTestSuite) TestGetPaymentAccountsByAppEntity() {
	accounts := models.GetPaymentAccountsByAppEntity(assert.appEntity1)
	// BTC + EUR accounts
	assert.GreaterOrEqual(len(accounts), 2)
}

func (assert *ServiceTestSuite) TestGetTradingAccountInstruments() {
	// After placing and executing a trade, trading account instruments should exist
	// First, just check it doesn't panic with an empty result
	instruments := models.GetTradingAccountInstruments(assert.tradingAccount1)
	// May be empty if no instrument positions yet — that's fine
	assert.NotNil(instruments)
}
