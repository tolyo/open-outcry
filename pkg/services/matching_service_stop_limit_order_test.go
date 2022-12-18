package services

import (
	"open-outcry/pkg/models"
	"open-outcry/pkg/utils"
)

func (assert *ServiceTestSuite) TestCreateStopLimitSellOrderSave() {
	// when: a stop limit order is created
	res, _ := ProcessTradeOrder(Acc(), "BTC_EUR", models.StopLimit, "SELL", 10, 100, "GTC")

	// then:
	assert.NotNil(res)
	assert.Equal(1, utils.GetCount("stop_order"))
	assert.Equal(0, GetSellBookOrderCount())
	assert.Equal([]PriceVolume{}, GetVolumes("BTC_EUR", "SELL"))

	assert.Equal(100.0, models.FindPaymentAccountByAppEntityIdAndCurrencyName(GetAppEntityId(), "BTC").AmountReserved)
}

func (assert *ServiceTestSuite) TestCreateStopLimitBuyOrderBuy() {
	tradingAccount := Acc()
	// tradingAccount2 := Acc2()

	// when: a stop limit order is created
	ProcessTradeOrder(tradingAccount, "BTC_EUR", models.StopLimit, "BUY", 10, 100, "GTC")

	// then:

	assert.Equal(1, utils.GetCount("stop_order"))
	assert.Equal(0, GetSellBookOrderCount())
	assert.Equal([]PriceVolume{}, GetVolumes("BTC_EUR", "BUY"))
	assert.Equal([]PriceVolume{}, GetVolumes("BTC_EUR", "SELL"))

	assert.Equal(1000.0, models.FindPaymentAccountByAppEntityIdAndCurrencyName(GetAppEntityId(), "EUR").AmountReserved)
}

func (assert *ServiceTestSuite) TestCreateStopLimitSellOrderActivate() {
	tradingAccount := Acc()
	tradingAccount2 := Acc2()

	// when: a stop limit order is created and then a crossing trade occurs
	ProcessTradeOrder(tradingAccount, "BTC_EUR", models.StopLimit, "SELL", 10, 2, "GTC")
	ProcessTradeOrder(tradingAccount, "BTC_EUR", "LIMIT", "SELL", 10, 1, "GTC")
	ProcessTradeOrder(tradingAccount2, "BTC_EUR", "LIMIT", "BUY", 10, 1, "GTC")
	// then: the order becomes activated

	assert.Equal(0, utils.GetCount("stop_order"))
	assert.Equal(1, GetSellBookOrderCount())
	assert.Equal([]PriceVolume{}, GetVolumes("BTC_EUR", "BUY"))
	assert.Equal([]PriceVolume{{10.0, 2.0}}, GetVolumes("BTC_EUR", "SELL"))
}

func (assert *ServiceTestSuite) TestCreateStopLimitBuyOrderActivate() {
	tradingAccount := Acc()
	tradingAccount2 := Acc2()

	// when: a stop limit order is created and then a crossing trade occurs
	ProcessTradeOrder(tradingAccount, "BTC_EUR", models.StopLimit, "BUY", 10, 2, "GTC")
	ProcessTradeOrder(tradingAccount2, "BTC_EUR", "LIMIT", "BUY", 10, 1, "GTC")
	ProcessTradeOrder(tradingAccount, "BTC_EUR", "LIMIT", "SELL", 10, 1, "GTC")

	// then: the order becomes activated
	assert.Equal(0, utils.GetCount("stop_order"))
	assert.Equal(1, GetBuyBookOrderCount())
	assert.Equal([]PriceVolume{{10.0, 2.0}}, GetVolumes("BTC_EUR", "BUY"))
	assert.Equal([]PriceVolume{}, GetVolumes("BTC_EUR", "SELL"))
}

func (assert *ServiceTestSuite) TestCreateStopLimitSellOrderNonCrossing() {
	tradingAccount := Acc()
	tradingAccount2 := Acc2()

	// when: a stop limit order is created and then a non-crossing trade occurs
	ProcessTradeOrder(tradingAccount, "BTC_EUR", models.StopLimit, "SELL", 10, 2, "GTC")
	ProcessTradeOrder(tradingAccount, "BTC_EUR", "LIMIT", "SELL", 11, 1, "GTC")
	ProcessTradeOrder(tradingAccount2, "BTC_EUR", "LIMIT", "BUY", 11, 1, "GTC")

	// then: the order remains deactivated
	assert.Equal(1, utils.GetCount("stop_order"))
	assert.Equal(0, GetSellBookOrderCount())
	assert.Equal([]PriceVolume{}, GetVolumes("BTC_EUR", "BUY"))
	assert.Equal([]PriceVolume{}, GetVolumes("BTC_EUR", "SELL"))

	// when: a crossing trade occurs
	ProcessTradeOrder(tradingAccount, "BTC_EUR", "LIMIT", "SELL", 9, 1, "GTC")
	ProcessTradeOrder(tradingAccount2, "BTC_EUR", "LIMIT", "BUY", 9, 1, "GTC")

	// then: it becomes activated
	assert.Equal(0, utils.GetCount("stop_order"))
	assert.Equal(1, GetSellBookOrderCount())
	assert.Equal([]PriceVolume{}, GetVolumes("BTC_EUR", "BUY"))
	assert.Equal([]PriceVolume{{10.0, 2.0}}, GetVolumes("BTC_EUR", "SELL"))
}

func (assert *ServiceTestSuite) TestCreateStopLimitBuyOrderNonCrossing() {
	tradingAccount := Acc()
	tradingAccount2 := Acc2()

	// when: a stop limit order is created and then a non-crossing trade occurs
	ProcessTradeOrder(tradingAccount, "BTC_EUR", models.StopLimit, "BUY", 10, 2, "GTC")
	ProcessTradeOrder(tradingAccount2, "BTC_EUR", "LIMIT", "BUY", 11, 1, "GTC")
	ProcessTradeOrder(tradingAccount, "BTC_EUR", "LIMIT", "SELL", 11, 1, "GTC")

	// then: the order remains deactivated
	assert.Equal(1, utils.GetCount("stop_order"))
	assert.Equal(0, GetSellBookOrderCount())
	assert.Equal([]PriceVolume{}, GetVolumes("BTC_EUR", "BUY"))
	assert.Equal([]PriceVolume{}, GetVolumes("BTC_EUR", "SELL"))

	// when: a crossing trade occurs
	ProcessTradeOrder(tradingAccount2, "BTC_EUR", "LIMIT", "BUY", 9, 1, "GTC")
	ProcessTradeOrder(tradingAccount, "BTC_EUR", "LIMIT", "SELL", 9, 1, "GTC")

	// then: it becomes activated
	assert.Equal(0, utils.GetCount("stop_order"))
	assert.Equal(1, GetBuyBookOrderCount())
	assert.Equal([]PriceVolume{{10.0, 2.0}}, GetVolumes("BTC_EUR", "BUY"))
	assert.Equal([]PriceVolume{}, GetVolumes("BTC_EUR", "SELL"))
}

func (assert *ServiceTestSuite) TestCreateStopLimitSellOrderActivateAndSettle() {
	tradingAccount := Acc()
	tradingAccount2 := Acc2()

	// when: a stop limit order is created and then a crossing trade occurs
	ProcessTradeOrder(tradingAccount, "BTC_EUR", models.StopLimit, "SELL", 10, 2, "GTC")
	ProcessTradeOrder(tradingAccount, "BTC_EUR", "LIMIT", "SELL", 10, 1, "GTC")
	ProcessTradeOrder(tradingAccount2, "BTC_EUR", "LIMIT", "BUY", 10, 3, "GTC")

	// then: the order becomes activated and settled
	assert.Equal(0, utils.GetCount("stop_order"))
	assert.Equal(0, GetSellBookOrderCount())
	assert.Equal([]PriceVolume{}, GetVolumes("BTC_EUR", "BUY"))
	assert.Equal([]PriceVolume{}, GetVolumes("BTC_EUR", "SELL"))

}

func (assert *ServiceTestSuite) TestCreateStopLimitBuyOrderActivateAndSettle() {
	tradingAccount := Acc()
	tradingAccount2 := Acc2()

	// when: a stop limit order is created and then a crossing trade occurs
	ProcessTradeOrder(tradingAccount, "BTC_EUR", models.StopLimit, "BUY", 10, 2, "GTC")
	ProcessTradeOrder(tradingAccount, "BTC_EUR", "LIMIT", "BUY", 10, 1, "GTC")
	ProcessTradeOrder(tradingAccount2, "BTC_EUR", "LIMIT", "SELL", 10, 3, "GTC")

	// then: the order becomes activated and settled
	assert.Equal(0, utils.GetCount("stop_order"))
	assert.Equal(0, GetSellBookOrderCount())
	assert.Equal([]PriceVolume{}, GetVolumes("BTC_EUR", "BUY"))
	assert.Equal([]PriceVolume{}, GetVolumes("BTC_EUR", "SELL"))
}

func (assert *ServiceTestSuite) TestCreateStopLimitSellOrderActivateAndSettleOppositeSide() {
	tradingAccount := Acc()
	tradingAccount2 := Acc2()

	// when: a stop limit order is created and then a crossing trade occurs
	ProcessTradeOrder(tradingAccount, "BTC_EUR", models.StopLimit, "SELL", 10, 2, "GTC")
	ProcessTradeOrder(tradingAccount2, "BTC_EUR", "LIMIT", "BUY", 10, 3, "GTC")
	ProcessTradeOrder(tradingAccount, "BTC_EUR", "LIMIT", "SELL", 10, 1, "GTC")

	// then: the order becomes activated and settled
	assert.Equal(0, utils.GetCount("stop_order"))
	assert.Equal(0, GetSellBookOrderCount())
	assert.Equal([]PriceVolume{}, GetVolumes("BTC_EUR", "BUY"))
	assert.Equal([]PriceVolume{}, GetVolumes("BTC_EUR", "SELL"))
}

func (assert *ServiceTestSuite) TestCreateStopLimitBuyOrderActivateAndSettleOppositeSide() {
	tradingAccount := Acc()
	tradingAccount2 := Acc2()

	// when: a stop limit order is created and then a crossing trade occurs
	ProcessTradeOrder(tradingAccount, "BTC_EUR", models.StopLimit, "BUY", 10, 2, "GTC")
	ProcessTradeOrder(tradingAccount2, "BTC_EUR", "LIMIT", "SELL", 10, 3, "GTC")
	ProcessTradeOrder(tradingAccount, "BTC_EUR", "LIMIT", "BUY", 10, 1, "GTC")

	// then: the order becomes activated and settled
	assert.Equal(0, utils.GetCount("stop_order"))
	assert.Equal(0, GetSellBookOrderCount())
	assert.Equal([]PriceVolume{}, GetVolumes("BTC_EUR", "BUY"))
	assert.Equal([]PriceVolume{}, GetVolumes("BTC_EUR", "SELL"))
}

func (assert *ServiceTestSuite) TestCreateStopLimitSellOrderActivateAndSettleBeforeWorsePriceOrders() {
	tradingAccount := Acc()
	tradingAccount2 := Acc2()

	// when: a stop limit order is created and then a crossing trade occurs
	ProcessTradeOrder(tradingAccount, "BTC_EUR", models.StopLimit, "SELL", 10, 1, "GTC")
	ProcessTradeOrder(tradingAccount, "BTC_EUR", "LIMIT", "SELL", 11, 1, "GTC")
	ProcessTradeOrder(tradingAccount, "BTC_EUR", "LIMIT", "SELL", 10, 1, "GTC")
	ProcessTradeOrder(tradingAccount2, "BTC_EUR", "LIMIT", "BUY", 11, 3, "GTC")

	// then: the order becomes activated and settled
	assert.Equal([]float64{10.0, 10.0, 11.0}, GetTradePrices())
	assert.Equal(0, utils.GetCount("stop_order"))
	assert.Equal(0, GetSellBookOrderCount())
	assert.Equal([]PriceVolume{}, GetVolumes("BTC_EUR", "BUY"))
	assert.Equal([]PriceVolume{}, GetVolumes("BTC_EUR", "SELL"))
}

func (assert *ServiceTestSuite) TestCreateStopLimitBuyOrderActivateAndSettleBeforeWorsePriceOrders() {
	tradingAccount := Acc()
	tradingAccount2 := Acc2()

	// when: a stop limit order is created and then a crossing trade occurs
	ProcessTradeOrder(tradingAccount, "BTC_EUR", models.StopLimit, "BUY", 10, 1, "GTC")
	ProcessTradeOrder(tradingAccount, "BTC_EUR", "LIMIT", "BUY", 9, 1, "GTC")
	ProcessTradeOrder(tradingAccount, "BTC_EUR", "LIMIT", "BUY", 10, 1, "GTC")
	ProcessTradeOrder(tradingAccount2, "BTC_EUR", "LIMIT", "SELL", 9, 3, "GTC")
	//
	// then: the order becomes activated and settled
	assert.Equal([]float64{10.0, 10.0, 9.0}, GetTradePrices())
	assert.Equal(0, utils.GetCount("stop_order"))
	assert.Equal(0, GetSellBookOrderCount())
	assert.Equal([]PriceVolume{}, GetVolumes("BTC_EUR", "BUY"))
	assert.Equal([]PriceVolume{}, GetVolumes("BTC_EUR", "SELL"))
}

func (assert *ServiceTestSuite) TestCreateStopLimitSellOrderActivateByMarketAndSettle() {
	tradingAccount := Acc()
	tradingAccount2 := Acc2()

	// when: a stop limit order is created and then a crossing trade occurs
	ProcessTradeOrder(tradingAccount, "BTC_EUR", models.StopLimit, "SELL", 10, 2, "GTC")
	ProcessTradeOrder(tradingAccount, "BTC_EUR", models.Market, "SELL", 0, 1, "GTC")
	ProcessTradeOrder(tradingAccount2, "BTC_EUR", "LIMIT", "BUY", 10, 3, "GTC")

	// then: the order becomes activated and settled
	assert.Equal(0, utils.GetCount("stop_order"))
	assert.Equal(0, GetSellBookOrderCount())
	assert.Equal([]PriceVolume{}, GetVolumes("BTC_EUR", "BUY"))
	assert.Equal([]PriceVolume{}, GetVolumes("BTC_EUR", "SELL"))

	// when: a stop limit order is created and then a crossing trade occurs
	ProcessTradeOrder(tradingAccount, "BTC_EUR", models.StopLimit, "SELL", 10, 2, "GTC")
	ProcessTradeOrder(tradingAccount2, "BTC_EUR", "LIMIT", "BUY", 10, 3, "GTC")
	ProcessTradeOrder(tradingAccount, "BTC_EUR", models.Market, "SELL", 0, 1, "GTC")

	// then: the order becomes activated and settled
	assert.Equal(0, utils.GetCount("stop_order"))
	assert.Equal(0, GetSellBookOrderCount())
	assert.Equal([]PriceVolume{}, GetVolumes("BTC_EUR", "BUY"))
	assert.Equal([]PriceVolume{}, GetVolumes("BTC_EUR", "SELL"))
}

func (assert *ServiceTestSuite) TestCreateStopLimitBuyOrderActivateByMarketAndSettle() {
	tradingAccount := Acc()
	tradingAccount2 := Acc2()
	// when: a stop limit order is created and then a crossing trade occurs
	ProcessTradeOrder(tradingAccount, "BTC_EUR", models.StopLimit, "BUY", 10, 2, "GTC")
	ProcessTradeOrder(tradingAccount, "BTC_EUR", models.Market, "BUY", 0, 10, "GTC")
	ProcessTradeOrder(tradingAccount2, "BTC_EUR", "LIMIT", "SELL", 10, 3, "GTC")
	//
	// then: the order becomes activated and settled
	assert.Equal(0, utils.GetCount("stop_order"))
	assert.Equal(0, GetSellBookOrderCount())
	assert.Equal([]PriceVolume{}, GetVolumes("BTC_EUR", "BUY"))
	assert.Equal([]PriceVolume{}, GetVolumes("BTC_EUR", "SELL"))
	//
	// when: a stop limit order is created and then a crossing trade occurs
	ProcessTradeOrder(tradingAccount, "BTC_EUR", models.StopLimit, "BUY", 10, 2, "GTC")
	ProcessTradeOrder(tradingAccount2, "BTC_EUR", "LIMIT", "SELL", 10, 3, "GTC")
	ProcessTradeOrder(tradingAccount, "BTC_EUR", models.Market, "BUY", 0, 10, "GTC")
	//
	// then: the order becomes activated and settled
	assert.Equal(0, utils.GetCount("stop_order"))

	assert.Equal(0, GetSellBookOrderCount())
	assert.Equal([]PriceVolume{}, GetVolumes("BTC_EUR", "BUY"))
	assert.Equal([]PriceVolume{}, GetVolumes("BTC_EUR", "SELL"))
}