package services

import (
	"open-outcry/pkg/db"
	"open-outcry/pkg/models"
)

func (assert *ServiceTestSuite) TestGtc() {
	// given

	entity := assert.appEntity1

	// when: given a new order
	tradeOrder, _ := ProcessTradeOrder(assert.tradingAccount1, "BTC_EUR", "LIMIT", models.Sell, 1.0, 10.0, "GTC")

	// then: it remains in the book until cancelled
	assert.Equal(1, db.QueryVal[int]("SELECT COUNT(*) FROM book_order"))
	assert.Equal(models.Open, models.GetTradeOrder(tradeOrder).Status)
	assert.Equal(10.0, models.FindPaymentAccountByAppEntityIdAndCurrencyName(entity, "BTC").AmountReserved)
}

func (assert *ServiceTestSuite) TestFok() {
	// given
	entity := assert.appEntity1

	// when: given a new order that cannot be filled
	tradeOrder, _ := ProcessTradeOrder(assert.tradingAccount1, "BTC_EUR", "LIMIT", models.Sell, 1.0, 1.0, models.FOK)

	// then: it is rejected
	assert.Equal(0, db.QueryVal[int]("SELECT COUNT(*) FROM book_order"))
	assert.Equal(models.Rejected, models.GetTradeOrder(tradeOrder).Status)
	assert.Equal(0.0, GetVolumeAtPrice("BTC_EUR", models.Sell, 1.0))

	assert.Equal(0.0, models.FindPaymentAccountByAppEntityIdAndCurrencyName(entity, "BTC").AmountReserved)

	//when: given a new order that cannot be filled even when other orders present
	ProcessTradeOrder(assert.tradingAccount2, "BTC_EUR", "LIMIT", models.Buy, 1.0, 1.0, "GTC")

	tradeOrder, _ = ProcessTradeOrder(assert.tradingAccount1, "BTC_EUR", "LIMIT", models.Sell, 1.0, 2.0, models.FOK)

	//then: it is rejected
	assert.Equal(models.Rejected, models.GetTradeOrder(tradeOrder).Status)

	assert.Equal(0.0, models.FindPaymentAccountByAppEntityIdAndCurrencyName(entity, "BTC").AmountReserved)

	//when: added another market order that can fill
	ProcessTradeOrder(assert.tradingAccount2, "BTC_EUR", models.Market, models.Buy, 0.0, 2.0, "GTC")

	tradeOrder, _ = ProcessTradeOrder(assert.tradingAccount1, "BTC_EUR", "LIMIT", models.Sell, 1.0, 2.0, models.FOK)

	//then: it is not reject
	assert.Equal(models.Filled, models.GetTradeOrder(tradeOrder).Status)
	assert.Equal(0.0, GetVolumeAtPrice("BTC_EUR", models.Sell, 1.0))

	assert.Equal(0.0, models.FindPaymentAccountByAppEntityIdAndCurrencyName(entity, "BTC").AmountReserved)
}

func (assert *ServiceTestSuite) TestIoc() {
	// given
	entity := assert.appEntity1

	// when: given a new order that cannot be filled
	tradeOrder, _ := ProcessTradeOrder(assert.tradingAccount1, "BTC_EUR", "LIMIT", models.Sell, 1.0, 1.0, models.IOC)

	//	then: it is rejected
	assert.Equal(0, db.QueryVal[int]("SELECT COUNT(*) FROM book_order"))
	assert.Equal(models.Rejected, models.GetTradeOrder(tradeOrder).Status)
	assert.Equal(0.0, GetVolumeAtPrice("BTC_EUR", models.Sell, 1.0))

	assert.Equal(0.0, models.FindPaymentAccountByAppEntityIdAndCurrencyName(entity, "BTC").AmountReserved)

	//when: given a new order that can only be partially filled by a standing order in the order book
	ProcessTradeOrder(assert.tradingAccount2, "BTC_EUR", "LIMIT", models.Buy, 1.0, 1, "GTC")
	tradeOrder, _ = ProcessTradeOrder(assert.tradingAccount1, "BTC_EUR", "LIMIT", models.Sell, 1.0, 2, models.IOC)

	// then: it is partially rejected
	assert.Equal(1, GetTradeCount())
	assert.Equal(0, db.QueryVal[int]("SELECT COUNT(*) FROM book_order"))
	assert.Equal(models.PartiallyRejected, models.GetTradeOrder(tradeOrder).Status)
	assert.Equal(0.0, GetVolumeAtPrice("BTC_EUR", models.Sell, 1.0))

	assert.Equal(0.0, models.FindPaymentAccountByAppEntityIdAndCurrencyName(entity, "BTC").AmountReserved)
}
