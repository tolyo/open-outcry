package services

import "open-outcry/pkg/models"

//	These tests apply to balance constraints of order matching as we never want to allow
//	order into the order book that do no have sufficient leverage.
//	For limit orders, we ensure that the seller has sufficient amount in base currency or instument
//	and that the buys has sufficient amount in quote currency which is limit price times base currency or instrument.
//
//	For market order we ensure that a seller instrument amount is valid or buyer quote currency amount is valid.
//
// `
func (assert *ServiceTestSuite) TestProcessLimitSellOrderSaveWithInsufficientFunds() {
	// given:
	appEntity1 := CreateAppEntity("test3")
	models.CreatePaymentAccount(appEntity1, "BTC")
	CreatePaymentDeposit(appEntity1, 100, "BTC", "test", "Test")
	tradingAccountId := models.FindTradingAccountByApplicationEntityId(appEntity1).Id
	ProcessTradeOrder(tradingAccountId, "BTC_EUR", "LIMIT", models.Sell, 10, 100, "GTC")
	CreatePaymentDeposit(appEntity1, 100, "BTC", "test", "Test")

	// when: a limit order is sent with insufficient funds
	_, err := ProcessTradeOrder(tradingAccountId, "BTC_EUR", "LIMIT", models.Sell, 10, 101, "GTC")

	// then: err
	assert.NotNil(err)
}

func (assert *ServiceTestSuite) TestProcessLimitBuyOrderSaveWithInsufficientFunds() {
	// given:
	appEntity1 := CreateAppEntity("test3")
	CreatePaymentDeposit(appEntity1, 100, "EUR", "test", "Test")
	tradingAccountId := models.FindTradingAccountByApplicationEntityId(appEntity1).Id
	ProcessTradeOrder(tradingAccountId, "BTC_EUR", "LIMIT", models.Buy, 10, 10, "GTC")
	CreatePaymentDeposit(appEntity1, 100, "EUR", "test", "Test")
	// when: a limit order is sent with insufficient funds

	// then: err
	_, err := ProcessTradeOrder(tradingAccountId, "BTC_EUR", "LIMIT", models.Buy, 10, 11, "GTC")
	assert.NotNil(err)
}

func (assert *ServiceTestSuite) TestProcessMarketSellOrderSaveWithInsufficientFunds() {
	// given:
	appEntity1 := CreateAppEntity("test3")
	models.CreatePaymentAccount(appEntity1, "BTC")
	CreatePaymentDeposit(appEntity1, 100, "BTC", "test", "Test")
	tradingAccountId := models.FindTradingAccountByApplicationEntityId(appEntity1).Id
	ProcessTradeOrder(tradingAccountId, "BTC_EUR", models.Market, models.Sell, 0, 100, "GTC")
	CreatePaymentDeposit(appEntity1, 100, "BTC", "test", "Test")
	// when: a market order is sent with insufficient funds

	// then: err
	_, err := ProcessTradeOrder(tradingAccountId, "BTC_EUR", models.Market, models.Sell, 0, 101, "GTC")
	assert.NotNil(err)
}

func (assert *ServiceTestSuite) TestProcessMarketBuyOrderSaveWithInsufficientFunds() {
	// given:
	appEntity1 := CreateAppEntity("test3")
	CreatePaymentDeposit(appEntity1, 100, "EUR", "test", "Test")
	tradingAccountId := models.FindTradingAccountByApplicationEntityId(appEntity1).Id
	ProcessTradeOrder(tradingAccountId, "BTC_EUR", models.Market, models.Buy, 0, 100, "GTC")
	CreatePaymentDeposit(appEntity1, 100, "EUR", "test", "Test")
	// when: a market order is sent with insufficient funds
	// then: exception is raised
	_, err := ProcessTradeOrder(tradingAccountId, "BTC_EUR", models.Market, models.Buy, 0, 101, "GTC")
	assert.NotNil(err)
}
