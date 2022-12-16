package services

import (
	"open-outcry/pkg/db"
	"open-outcry/pkg/models"
)

func CreateClient() models.AppEntityId {
	return CreateAppEntity("test")
}

func CreateClient2() models.AppEntityId {
	return CreateAppEntity("test2")
}

func CreateTradingAccountId() models.TradingAccountId {
	appEntityId := CreateClient()
	models.CreatePaymentAccount(appEntityId, models.CurrencyName("BTC"))
	CreatePaymentDeposit(appEntityId, 1000, "BTC", "Test", "Test")
	CreatePaymentDeposit(appEntityId, 1000, "EUR", "Test", "Test")
	tradingAccount := models.FindTradingAccountByApplicationEntityId(appEntityId)
	return tradingAccount.Id
}

func CreateTradingAccountId2() models.TradingAccountId {
	appEntityId := CreateClient2()
	models.CreatePaymentAccount(appEntityId, models.CurrencyName("BTC"))
	CreatePaymentDeposit(appEntityId, 1000, "BTC", "Test", "Test")
	CreatePaymentDeposit(appEntityId, 1000, "EUR", "Test", "Test")
	tradingAccount := models.FindTradingAccountByApplicationEntityId(appEntityId)
	return tradingAccount.Id
}

// // shorthand methods
func Acc() models.TradingAccountId {
	return CreateTradingAccountId()
}

func Acc2() models.TradingAccountId {
	return CreateTradingAccountId2()
}

func GetAppEntityId() models.AppEntityId {
	var res string
	db.Instance().QueryRow(`
	   SELECT pub_id FROM app_entity
	   WHERE external_id = 'test';
 	`).Scan(&res)
	return models.AppEntityId(res)
}

func GetAppEntityId2() models.AppEntityId {
	var res string
	db.Instance().QueryRow(`
	   SELECT pub_id FROM app_entity
	   WHERE external_id = 'test2';
 	`).Scan(&res)
	return models.AppEntityId(res)
}
