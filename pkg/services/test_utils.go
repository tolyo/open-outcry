package services

// AppState represents the expected transfer account state for both test entities
type AppState struct {
	entity1         []CurrencyAccount
	entity2         []CurrencyAccount
	tradeCount      int
	orderBookStates OrderBook
}

// TestStep is a representation of initial and final account states with orders to be executed in between
type TestStep struct {
	initialState  AppState
	orders        []TradeOrder
	expectedState AppState
}

// MatchingServiceTestCase represents a series of steps that need to be taken within each test case
type MatchingServiceTestCase struct {
	steps []TestStep
}

// Acc shorthand methods
func Acc(v string) (AppEntityId, InstrumentAccountId) {
	appEntityId := CreateAppEntity(AppEntityExternalId(v))
	CreateCurrencyAccount(appEntityId, "BTC")
	CreateTransferDeposit(appEntityId, 1000, "BTC", "Test", "Test")
	CreateTransferDeposit(appEntityId, 1000, "EUR", "Test", "Test")
	instrumentAccount := FindInstrumentAccountByApplicationEntityId(appEntityId)
	return appEntityId, instrumentAccount.Id
}
