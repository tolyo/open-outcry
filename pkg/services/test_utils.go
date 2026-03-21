package services

import (
	appentity "open-outcry/pkg/models/app_entity"
	currencyaccount "open-outcry/pkg/models/currency_account"
	instrumentaccount "open-outcry/pkg/models/instrument_account"
	orderbook "open-outcry/pkg/models/order_book"
	tradeorder "open-outcry/pkg/models/trade_order"
)

type AppState struct {
	entity1         []currencyaccount.CurrencyAccount
	tradeCount      int
	orderBookStates orderbook.OrderBook
}

type TestStep struct {
	initialState  AppState
	orders        []tradeorder.TradeOrder
	expectedState AppState
}

type MatchingServiceTestCase struct {
	steps []TestStep
}

func Acc(v string) (appentity.AppEntityId, instrumentaccount.InstrumentAccountId) {
	appEntityId := CreateAppEntity(appentity.AppEntityExternalId(v))
	currencyaccount.CreateCurrencyAccount(appEntityId, "BTC")
	CreateTransferDeposit(appEntityId, 1000, "BTC", "Test", "Test")
	CreateTransferDeposit(appEntityId, 1000, "EUR", "Test", "Test")
	instrumentAccount := instrumentaccount.FindInstrumentAccountByApplicationEntityId(appEntityId)
	return appEntityId, instrumentAccount.Id
}
