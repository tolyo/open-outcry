package services

func (assert *ServiceTestSuite) TestGetTradeOrdersByInstrumentAccount() {
	// Place an order
	orderId, err := ProcessTradeOrder(assert.instrumentAccount1, "BTC_EUR", "LIMIT", Sell, 10.6, 100, "GTC")
	assert.Nil(err)

	// Fetch orders for this trading account
	orders := GetTradeOrdersByInstrumentAccount(assert.instrumentAccount1)
	assert.GreaterOrEqual(len(orders), 1)
	assert.Equal(orderId, orders[0].Id)
	assert.Equal(assert.instrumentAccount1, orders[0].InstrumentAccountId)
	assert.Equal(InstrumentName("BTC_EUR"), orders[0].InstrumentName)
	assert.Equal(Sell, orders[0].Side)
	assert.Equal(Limit, orders[0].Type)

	// Other account should have no orders
	orders2 := GetTradeOrdersByInstrumentAccount(assert.instrumentAccount2)
	assert.Equal(0, len(orders2))
}

func (assert *ServiceTestSuite) TestGetBookOrdersByInstrumentAccount() {
	// Place a limit order (goes to book)
	ProcessTradeOrder(assert.instrumentAccount1, "BTC_EUR", "LIMIT", Sell, 10.6, 100, "GTC")

	bookOrders := GetBookOrdersByInstrumentAccount(assert.instrumentAccount1)
	assert.Equal(1, len(bookOrders))
	assert.Equal(Sell, bookOrders[0].Side)

	// Other account has no book orders
	bookOrders2 := GetBookOrdersByInstrumentAccount(assert.instrumentAccount2)
	assert.Equal(0, len(bookOrders2))
}

func (assert *ServiceTestSuite) TestGetTradesByInstrumentAccount() {
	// Create a matching pair that produces a trade
	ProcessTradeOrder(assert.instrumentAccount1, "BTC_EUR", "LIMIT", Sell, 10.0, 100, "GTC")
	ProcessTradeOrder(assert.instrumentAccount2, "BTC_EUR", "LIMIT", Buy, 10.0, 100, "GTC")

	// Both accounts should see the trade
	trades1 := GetTradesByInstrumentAccount(assert.instrumentAccount1)
	assert.GreaterOrEqual(len(trades1), 1)

	trades2 := GetTradesByInstrumentAccount(assert.instrumentAccount2)
	assert.GreaterOrEqual(len(trades2), 1)

	// Same trade
	assert.Equal(trades1[0].Id, trades2[0].Id)
	assert.Equal(10.0, trades1[0].Price)
	assert.Equal(100.0, trades1[0].Amount)
}

func (assert *ServiceTestSuite) TestGetTrade() {
	// Create a trade
	ProcessTradeOrder(assert.instrumentAccount1, "BTC_EUR", "LIMIT", Sell, 10.0, 100, "GTC")
	ProcessTradeOrder(assert.instrumentAccount2, "BTC_EUR", "LIMIT", Buy, 10.0, 100, "GTC")

	trades := GetTradesByInstrumentAccount(assert.instrumentAccount1)
	assert.GreaterOrEqual(len(trades), 1)

	// Fetch by id
	trade := GetTrade(trades[0].Id)
	assert.NotNil(trade)
	assert.Equal(trades[0].Id, trade.Id)
	assert.Equal(InstrumentName("BTC_EUR"), trade.InstrumentName)
}

func (assert *ServiceTestSuite) TestGetTradeNotFound() {
	trade := GetTrade("nonexistent-id")
	assert.Nil(trade)
}

func (assert *ServiceTestSuite) TestGetAppEntities() {
	entities := GetAppEntities()
	// MASTER + 2 test entities
	assert.GreaterOrEqual(len(entities), 3)
}

func (assert *ServiceTestSuite) TestGetAppEntity() {
	entity := GetAppEntity(assert.appEntity1)
	assert.NotNil(entity)
	assert.Equal(assert.appEntity1, entity.Id)
}

func (assert *ServiceTestSuite) TestGetAppEntityNotFound() {
	entity := GetAppEntity("nonexistent-id")
	assert.Nil(entity)
}

func (assert *ServiceTestSuite) TestGetTransfersByAppEntity() {
	// SetupTest creates deposits, so there should be transfers
	transfers := GetTransfersByAppEntity(assert.appEntity1)
	assert.GreaterOrEqual(len(transfers), 1)
	assert.NotEmpty(transfers[0].Id)
	assert.NotEmpty(transfers[0].Currency)
}

func (assert *ServiceTestSuite) TestGetTransfer() {
	transfers := GetTransfersByAppEntity(assert.appEntity1)
	assert.GreaterOrEqual(len(transfers), 1)

	transfer := GetTransfer(transfers[0].Id)
	assert.NotNil(transfer)
	assert.Equal(transfers[0].Id, transfer.Id)
	assert.Equal(transfers[0].Currency, transfer.Currency)
}

func (assert *ServiceTestSuite) TestGetTransferNotFound() {
	transfer := GetTransfer("nonexistent-id")
	assert.Nil(transfer)
}

func (assert *ServiceTestSuite) TestGetCurrencyAccountsByAppEntity() {
	accounts := GetCurrencyAccountsByAppEntity(assert.appEntity1)
	// BTC + EUR accounts
	assert.GreaterOrEqual(len(accounts), 2)
}

func (assert *ServiceTestSuite) TestGetInstrumentAccountHoldings() {
	// After placing and executing a trade, trading account instruments should exist
	// First, just check it doesn't panic with an empty result
	instruments := GetInstrumentAccountHoldings(assert.instrumentAccount1)
	// May be empty if no instrument positions yet — that's fine
	assert.NotNil(instruments)
}
