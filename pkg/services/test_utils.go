package services

import (
	"open-outcry/pkg/models"
)

// // shorthand methods
func Acc(v string) (models.AppEntityId, models.TradingAccountId) {
	appEntityId := CreateAppEntity(models.AppEntityExternalId(v))
	models.CreatePaymentAccount(appEntityId, models.CurrencyName("BTC"))
	CreatePaymentDeposit(appEntityId, 1000, "BTC", "Test", "Test")
	CreatePaymentDeposit(appEntityId, 1000, "EUR", "Test", "Test")
	tradingAccount := models.FindTradingAccountByApplicationEntityId(appEntityId)
	return appEntityId, tradingAccount.Id
}
