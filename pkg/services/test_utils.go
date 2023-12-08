package services

import (
	"open-outcry/pkg/models"
)

// AppState represents the expected payment account state for both test entities
type AppState struct {
	entity1         []models.PaymentAccount
	entity2         []models.PaymentAccount
	tradeCount      int
	orderBookStates models.OrderBook
}

// TestStep is a representation of initial and final account states with orders to be executed in between
type TestStep struct {
	initialState  AppState
	orders        []models.TradeOrder
	expectedState AppState
}

// MatchingServiceTestCase represents a series of steps that need to be taken within each test case
type MatchingServiceTestCase struct {
	steps []TestStep
}

// // shorthand methods
func Acc(v string) (models.AppEntityId, models.TradingAccountId) {
	appEntityId := CreateAppEntity(models.AppEntityExternalId(v))
	models.CreatePaymentAccount(appEntityId, models.CurrencyName("BTC"))
	CreatePaymentDeposit(appEntityId, 1000, "BTC", "Test", "Test")
	CreatePaymentDeposit(appEntityId, 1000, "EUR", "Test", "Test")
	tradingAccount := models.FindTradingAccountByApplicationEntityId(appEntityId)
	return appEntityId, tradingAccount.Id
}
