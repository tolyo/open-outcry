package services

import "open-outcry/pkg/models"

type TestData struct {
	side                     models.OrderSide
	price                    float64
	amount                   float64
	expectedBtcReserveAmount float64
	expectedEurReserveAmount float64
}

var testData = []TestData{
	{
		side:                     models.Sell,
		price:                    10,
		amount:                   100,
		expectedEurReserveAmount: 0,
		expectedBtcReserveAmount: 100,
	},

	{
		side:                     models.Buy,
		price:                    10,
		amount:                   1,
		expectedEurReserveAmount: 10,
		expectedBtcReserveAmount: 0,
	},
}

func (assert *ServiceTestSuite) TestCancelTradeOrder() {

	for _, data := range testData {
		// given: an existing trade limit order
		tradeOrder, err := ProcessTradeOrder(assert.tradingAccount1,
			"BTC_EUR",
			"LIMIT",
			data.side,
			models.OrderPrice(data.price),
			data.amount,
			"GTC",
		)
		assert.Nil(err)
		assert.Equal(1, GetBookOrderCount(data.side))
		assert.Equal(
			[]PriceVolume{{Price: data.price, Volume: data.amount}},
			GetVolumes("BTC_EUR", data.side),
		)

		res := models.GetTradeOrder(tradeOrder)
		assert.Equal(models.Open, res.Status)
		assert.Equal(data.expectedBtcReserveAmount,
			models.FindPaymentAccountByAppEntityIdAndCurrencyName(assert.appEntity1, "BTC").AmountReserved)
		assert.Equal(data.expectedEurReserveAmount,
			models.FindPaymentAccountByAppEntityIdAndCurrencyName(assert.appEntity1, "EUR").AmountReserved)

		// when: order is cancelled
		CancelTradeOrder(tradeOrder)

		// then: it is removed from the order book
		assert.Equal(0, GetBookOrderCount(data.side))
		assert.Equal([]PriceVolume{}, GetVolumes("BTC_EUR", data.side))

		// and: its status is cancelled
		res = models.GetTradeOrder(tradeOrder)
		assert.Equal(models.Cancelled, res.Status)
		assert.Equal(0.00,
			models.FindPaymentAccountByAppEntityIdAndCurrencyName(assert.appEntity1, "BTC").AmountReserved)
		assert.Equal(0.00,
			models.FindPaymentAccountByAppEntityIdAndCurrencyName(assert.appEntity1, "EUR").AmountReserved)
	}

}

func (assert *ServiceTestSuite) TestCancelTradeOrderWithMultipleOrders() {
	// given
	assert.Equal(0.0, models.FindPaymentAccountByAppEntityIdAndCurrencyName(assert.appEntity1, "BTC").AmountReserved)

	tradeOrderId, _ := ProcessTradeOrder(assert.tradingAccount1,
		"BTC_EUR", "LIMIT", models.Sell, 10, 100, "GTC")

	tradeOrderId2, _ := ProcessTradeOrder(assert.tradingAccount1,
		"BTC_EUR", "LIMIT", models.Sell, 5, 100, "GTC")

	assert.Equal(2, GetSellBookOrderCount())

	assert.Equal(200.00, models.FindPaymentAccountByAppEntityIdAndCurrencyName(
		assert.appEntity1,
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
		assert.appEntity1,
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
		assert.appEntity1,
		"BTC",
	).AmountReserved)
}

func (assert *ServiceTestSuite) TestCancelTradeOrderWithPartiallyExecutedOrder() {
	// given: an existing trade limit order

	assert.Equal(0.0, models.FindPaymentAccountByAppEntityIdAndCurrencyName(
		assert.appEntity1,
		"BTC",
	).AmountReserved)

	tradeOrderId, _ := ProcessTradeOrder(assert.tradingAccount1,
		"BTC_EUR", "LIMIT", models.Sell, 10, 100, "GTC")

	assert.Equal(1, GetSellBookOrderCount())

	assert.Equal(100.00, models.FindPaymentAccountByAppEntityIdAndCurrencyName(
		assert.appEntity1,
		"BTC",
	).AmountReserved)

	// when: it is partially matched against another order and then cancelled
	ProcessTradeOrder(assert.tradingAccount2, "BTC_EUR", "LIMIT", models.Buy, 10, 50, "GTC")

	assert.Equal(50.00, models.FindPaymentAccountByAppEntityIdAndCurrencyName(
		assert.appEntity1,
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
		assert.appEntity1,
		"BTC",
	).AmountReserved)
}
