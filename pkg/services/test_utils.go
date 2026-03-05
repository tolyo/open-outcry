package services

import (
	"open-outcry/pkg/models"
)

// AppState represents the expected transfer account state for both test entities
type AppState struct {
	entity1         []models.CurrencyAccount
	entity2         []models.CurrencyAccount
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

// Acc shorthand methods
func Acc(v string) (models.AppEntityId, models.InstrumentAccountId) {
	appEntityId := CreateAppEntity(models.AppEntityExternalId(v))
	models.CreateCurrencyAccount(appEntityId, "BTC")
	CreateTransferDeposit(appEntityId, 1000, "BTC", "Test", "Test")
	CreateTransferDeposit(appEntityId, 1000, "EUR", "Test", "Test")
	instrumentAccount := models.FindInstrumentAccountByApplicationEntityId(appEntityId)
	return appEntityId, instrumentAccount.Id
}
