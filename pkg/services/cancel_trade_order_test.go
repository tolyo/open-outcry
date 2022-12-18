package services

import "open-outcry/pkg/models"

func (assert *ServiceTestSuite) TestCancelTradeOrder() {
	// given: an existing trade limit order
	tradingAccount := Acc()
	tradeOrder, _ := ProcessTradeOrder(tradingAccount, "BTC_EUR", "LIMIT", "SELL", 10, 100, "GTC")

	assert.Equal(1, GetSellBookOrderCount())
	assert.Equal([]PriceVolume{
		{Price: 10, Volume: 100},
	}, GetVolumes("BTC_EUR", "SELL"))

	// when: order is cancelled
	CancelTradeOrder(tradeOrder)

	// then: it is removed from the order book
	assert.Equal(0, GetSellBookOrderCount())
	assert.Equal([]PriceVolume{}, GetVolumes("BTC_EUR", "SELL"))

	// and: its status is cancelled
	res := models.GetTradeOrder(tradeOrder)
	assert.Equal(models.Cancelled, res.Status)
	assert.Equal(0.00,
		models.FindPaymentAccountByAppEntityIdAndCurrencyName(
			GetAppEntityId(), "BTC").AmountReserved,
	)
}

func (assert *ServiceTestSuite) TestCancelTradeOrderWithMultipleOrders() {
	// given
	tradingAccount := Acc()

	assert.Equal(0.0, models.FindPaymentAccountByAppEntityIdAndCurrencyName(GetAppEntityId(), "BTC").AmountReserved)

	tradeOrderId, _ := ProcessTradeOrder(tradingAccount,
		"BTC_EUR", "LIMIT", "SELL", 10, 100, "GTC")

	tradeOrderId2, _ := ProcessTradeOrder(tradingAccount,
		"BTC_EUR", "LIMIT", "SELL", 5, 100, "GTC")

	assert.Equal(2, GetSellBookOrderCount())

	assert.Equal(200.00, models.FindPaymentAccountByAppEntityIdAndCurrencyName(
		GetAppEntityId(),
		"BTC",
	).AmountReserved)

	// when: first order is cancelled
	CancelTradeOrder(tradeOrderId)

	// then: it is removed from the order book
	assert.Equal(1, GetSellBookOrderCount())

	// and: its status is cancelled
	assert.Equal(models.Cancelled, models.GetTradeOrder(tradeOrderId).Status)

	//  and: funds are released
	assert.Equal(100.00, models.FindPaymentAccountByAppEntityIdAndCurrencyName(
		GetAppEntityId(),
		"BTC",
	).AmountReserved)

	// when: second order is cancelled
	CancelTradeOrder(tradeOrderId2)
	//  then: it is removed from the order book
	assert.Equal(0, GetSellBookOrderCount())
	// and: its status is cancelled
	assert.Equal(models.Cancelled, models.GetTradeOrder(tradeOrderId2).Status)
	// and: funds are released
	assert.Equal(0.0, models.FindPaymentAccountByAppEntityIdAndCurrencyName(
		GetAppEntityId(),
		"BTC",
	).AmountReserved)
}

func (assert *ServiceTestSuite) TestCancelTradeOrderWithPartiallyExecdOrder() {
	// given: an existing trade limit order
	tradingAccount := Acc()
	tradingAccount2 := Acc2()

	assert.Equal(0.0, models.FindPaymentAccountByAppEntityIdAndCurrencyName(
		GetAppEntityId(),
		"BTC",
	).AmountReserved)

	tradeOrderId, _ := ProcessTradeOrder(tradingAccount,
		"BTC_EUR", "LIMIT", "SELL", 10, 100, "GTC")

	assert.Equal(1, GetSellBookOrderCount())

	assert.Equal(100.00, models.FindPaymentAccountByAppEntityIdAndCurrencyName(
		GetAppEntityId(),
		"BTC",
	).AmountReserved)

	// when: it is partially matched against another order and then cancelled
	ProcessTradeOrder(tradingAccount2, "BTC_EUR", "LIMIT", "BUY", 10, 50, "GTC")

	assert.Equal(50.00, models.FindPaymentAccountByAppEntityIdAndCurrencyName(
		GetAppEntityId(),
		"BTC",
	).AmountReserved)

	assert.Equal(models.PartiallyFilled, models.GetTradeOrder(tradeOrderId).Status)

	CancelTradeOrder(tradeOrderId)
	// then: it is removed from the order book
	assert.Equal(0, GetSellBookOrderCount())
	// and: its status is partially cancelled
	assert.Equal(models.PartiallyCancelled, models.GetTradeOrder(tradeOrderId).Status)
	// and: funds are released
	assert.Equal(0.00, models.FindPaymentAccountByAppEntityIdAndCurrencyName(
		GetAppEntityId(),
		"BTC",
	).AmountReserved)
}
