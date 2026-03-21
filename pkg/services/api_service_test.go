package services

import (
	"open-outcry/pkg/models/app_entity"
	"open-outcry/pkg/models/currency_account"
	"open-outcry/pkg/models/instrument"
	"open-outcry/pkg/models/instrument_account"
	"open-outcry/pkg/models/trade"
	"open-outcry/pkg/models/trade_order"
	"open-outcry/pkg/models/transfer"
)

func (assert *ServiceTestSuite) TestGetTradeOrdersByInstrumentAccount() {
	orderId, err := ProcessTradeOrder(assert.instrumentAccount1, "BTC_EUR", "LIMIT", tradeorder.Sell, 10.6, 100, "GTC")
	assert.Nil(err)

	orders := tradeorder.GetTradeOrdersByInstrumentAccount(assert.instrumentAccount1)
	assert.GreaterOrEqual(len(orders), 1)
	assert.Equal(orderId, orders[0].Id)
	assert.Equal(assert.instrumentAccount1, orders[0].InstrumentAccountId)
	assert.Equal(instrument.InstrumentName("BTC_EUR"), orders[0].InstrumentName)
	assert.Equal(tradeorder.Sell, orders[0].Side)
	assert.Equal(tradeorder.Limit, orders[0].Type)

	orders2 := tradeorder.GetTradeOrdersByInstrumentAccount(assert.instrumentAccount2)
	assert.Equal(0, len(orders2))
}

func (assert *ServiceTestSuite) TestGetBookOrdersByInstrumentAccount() {
	ProcessTradeOrder(assert.instrumentAccount1, "BTC_EUR", "LIMIT", tradeorder.Sell, 10.6, 100, "GTC")

	bookOrders := tradeorder.GetBookOrdersByInstrumentAccount(assert.instrumentAccount1)
	assert.Equal(1, len(bookOrders))
	assert.Equal(tradeorder.Sell, bookOrders[0].Side)

	bookOrders2 := tradeorder.GetBookOrdersByInstrumentAccount(assert.instrumentAccount2)
	assert.Equal(0, len(bookOrders2))
}

func (assert *ServiceTestSuite) TestGetTradesByInstrumentAccount() {
	ProcessTradeOrder(assert.instrumentAccount1, "BTC_EUR", "LIMIT", tradeorder.Sell, 10.0, 100, "GTC")
	ProcessTradeOrder(assert.instrumentAccount2, "BTC_EUR", "LIMIT", tradeorder.Buy, 10.0, 100, "GTC")

	trades1 := trade.GetTradesByInstrumentAccount(assert.instrumentAccount1)
	assert.GreaterOrEqual(len(trades1), 1)

	trades2 := trade.GetTradesByInstrumentAccount(assert.instrumentAccount2)
	assert.GreaterOrEqual(len(trades2), 1)

	assert.Equal(trades1[0].Id, trades2[0].Id)
	assert.Equal(10.0, trades1[0].Price)
	assert.Equal(100.0, trades1[0].Amount)
}

func (assert *ServiceTestSuite) TestGetTrade() {
	ProcessTradeOrder(assert.instrumentAccount1, "BTC_EUR", "LIMIT", tradeorder.Sell, 10.0, 100, "GTC")
	ProcessTradeOrder(assert.instrumentAccount2, "BTC_EUR", "LIMIT", tradeorder.Buy, 10.0, 100, "GTC")

	trades := trade.GetTradesByInstrumentAccount(assert.instrumentAccount1)
	assert.GreaterOrEqual(len(trades), 1)

	entry := trade.GetTrade(trades[0].Id)
	assert.NotNil(entry)
	assert.Equal(trades[0].Id, entry.Id)
	assert.Equal(instrument.InstrumentName("BTC_EUR"), entry.InstrumentName)
}

func (assert *ServiceTestSuite) TestGetTradeNotFound() {
	entry := trade.GetTrade("nonexistent-id")
	assert.Nil(entry)
}

func (assert *ServiceTestSuite) TestGetAppEntities() {
	entities := appentity.GetAppEntities()
	assert.GreaterOrEqual(len(entities), 3)
}

func (assert *ServiceTestSuite) TestGetAppEntity() {
	entity := appentity.GetAppEntity(assert.appEntity1)
	assert.NotNil(entity)
	assert.Equal(assert.appEntity1, entity.Id)
}

func (assert *ServiceTestSuite) TestGetAppEntityNotFound() {
	entity := appentity.GetAppEntity("nonexistent-id")
	assert.Nil(entity)
}

func (assert *ServiceTestSuite) TestGetTransfersByAppEntity() {
	transfers := transfer.GetTransfersByAppEntity(assert.appEntity1)
	assert.GreaterOrEqual(len(transfers), 1)
	assert.NotEmpty(transfers[0].Id)
	assert.NotEmpty(transfers[0].Currency)
}

func (assert *ServiceTestSuite) TestGetTransfer() {
	transfers := transfer.GetTransfersByAppEntity(assert.appEntity1)
	assert.GreaterOrEqual(len(transfers), 1)

	entry := transfer.GetTransfer(transfers[0].Id)
	assert.NotNil(entry)
	assert.Equal(transfers[0].Id, entry.Id)
	assert.Equal(transfers[0].Currency, entry.Currency)
}

func (assert *ServiceTestSuite) TestGetTransferNotFound() {
	entry := transfer.GetTransfer("nonexistent-id")
	assert.Nil(entry)
}

func (assert *ServiceTestSuite) TestGetCurrencyAccountsByAppEntity() {
	accounts := currencyaccount.GetCurrencyAccountsByAppEntity(assert.appEntity1)
	assert.GreaterOrEqual(len(accounts), 2)
}

func (assert *ServiceTestSuite) TestGetInstrumentAccountHoldings() {
	instruments := instrumentaccount.GetInstrumentAccountHoldings(assert.instrumentAccount1)
	assert.NotNil(instruments)
}
