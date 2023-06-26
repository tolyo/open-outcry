package services

import (
	"open-outcry/pkg/models"
	"open-outcry/pkg/utils"
)

func (assert *ServiceTestSuite) TestCreateStopLimitSellOrderSave() {
	// when: a stop limit order is created
	res, _ := ProcessTradeOrder(assert.tradingAccount1, "BTC_EUR", models.StopLimit, models.Sell, 10, 100, "GTC")

	// then:
	assert.NotNil(res)
	assert.Equal(1, utils.GetCount("stop_order"))
	assert.Equal(0, GetSellBookOrderCount())
	assert.Equal([]PriceVolume{}, GetVolumes("BTC_EUR", models.Sell))

	assert.Equal(100.0, models.FindPaymentAccountByAppEntityIdAndCurrencyName(assert.appEntity1, "BTC").AmountReserved)
}

func (assert *ServiceTestSuite) TestCreateStopLimitBuyOrderBuy() {

	//

	// when: a stop limit order is created
	ProcessTradeOrder(assert.tradingAccount1, "BTC_EUR", models.StopLimit, models.Buy, 10, 100, "GTC")

	// then:

	assert.Equal(1, utils.GetCount("stop_order"))
	assert.Equal(0, GetSellBookOrderCount())
	assert.Equal([]PriceVolume{}, GetVolumes("BTC_EUR", models.Buy))
	assert.Equal([]PriceVolume{}, GetVolumes("BTC_EUR", models.Sell))

	assert.Equal(1000.0, models.FindPaymentAccountByAppEntityIdAndCurrencyName(assert.appEntity1, "EUR").AmountReserved)
}

func (assert *ServiceTestSuite) TestCreateStopLimitSellOrderActivate() {

	// when: a stop limit order is created and then a crossing trade occurs
	ProcessTradeOrder(assert.tradingAccount1, "BTC_EUR", models.StopLimit, models.Sell, 10, 2, "GTC")
	ProcessTradeOrder(assert.tradingAccount1, "BTC_EUR", "LIMIT", models.Sell, 10, 1, "GTC")
	ProcessTradeOrder(assert.tradingAccount2, "BTC_EUR", "LIMIT", models.Buy, 10, 1, "GTC")
	// then: the order becomes activated

	assert.Equal(0, utils.GetCount("stop_order"))
	assert.Equal(1, GetSellBookOrderCount())
	assert.Equal([]PriceVolume{}, GetVolumes("BTC_EUR", models.Buy))
	assert.Equal([]PriceVolume{{10.0, 2.0}}, GetVolumes("BTC_EUR", models.Sell))
}

func (assert *ServiceTestSuite) TestCreateStopLimitBuyOrderActivate() {

	// when: a stop limit order is created and then a crossing trade occurs
	ProcessTradeOrder(assert.tradingAccount1, "BTC_EUR", models.StopLimit, models.Buy, 10, 2, "GTC")
	ProcessTradeOrder(assert.tradingAccount2, "BTC_EUR", "LIMIT", models.Buy, 10, 1, "GTC")
	ProcessTradeOrder(assert.tradingAccount1, "BTC_EUR", "LIMIT", models.Sell, 10, 1, "GTC")

	// then: the order becomes activated
	assert.Equal(0, utils.GetCount("stop_order"))
	assert.Equal(1, GetBuyBookOrderCount())
	assert.Equal([]PriceVolume{{10.0, 2.0}}, GetVolumes("BTC_EUR", models.Buy))
	assert.Equal([]PriceVolume{}, GetVolumes("BTC_EUR", models.Sell))
}

func (assert *ServiceTestSuite) TestCreateStopLimitSellOrderNonCrossing() {

	// when: a stop limit order is created and then a non-crossing trade occurs
	ProcessTradeOrder(assert.tradingAccount1, "BTC_EUR", models.StopLimit, models.Sell, 10, 2, "GTC")
	ProcessTradeOrder(assert.tradingAccount1, "BTC_EUR", "LIMIT", models.Sell, 11, 1, "GTC")
	ProcessTradeOrder(assert.tradingAccount2, "BTC_EUR", "LIMIT", models.Buy, 11, 1, "GTC")

	// then: the order remains deactivated
	assert.Equal(1, utils.GetCount("stop_order"))
	assert.Equal(0, GetSellBookOrderCount())
	assert.Equal([]PriceVolume{}, GetVolumes("BTC_EUR", models.Buy))
	assert.Equal([]PriceVolume{}, GetVolumes("BTC_EUR", models.Sell))

	// when: a crossing trade occurs
	ProcessTradeOrder(assert.tradingAccount1, "BTC_EUR", "LIMIT", models.Sell, 9, 1, "GTC")
	ProcessTradeOrder(assert.tradingAccount2, "BTC_EUR", "LIMIT", models.Buy, 9, 1, "GTC")

	// then: it becomes activated
	assert.Equal(0, utils.GetCount("stop_order"))
	assert.Equal(1, GetSellBookOrderCount())
	assert.Equal([]PriceVolume{}, GetVolumes("BTC_EUR", models.Buy))
	assert.Equal([]PriceVolume{{10.0, 2.0}}, GetVolumes("BTC_EUR", models.Sell))
}

func (assert *ServiceTestSuite) TestCreateStopLimitBuyOrderNonCrossing() {

	// when: a stop limit order is created and then a non-crossing trade occurs
	ProcessTradeOrder(assert.tradingAccount1, "BTC_EUR", models.StopLimit, models.Buy, 10, 2, "GTC")
	ProcessTradeOrder(assert.tradingAccount2, "BTC_EUR", "LIMIT", models.Buy, 11, 1, "GTC")
	ProcessTradeOrder(assert.tradingAccount1, "BTC_EUR", "LIMIT", models.Sell, 11, 1, "GTC")

	// then: the order remains deactivated
	assert.Equal(1, utils.GetCount("stop_order"))
	assert.Equal(0, GetSellBookOrderCount())
	assert.Equal([]PriceVolume{}, GetVolumes("BTC_EUR", models.Buy))
	assert.Equal([]PriceVolume{}, GetVolumes("BTC_EUR", models.Sell))

	// when: a crossing trade occurs
	ProcessTradeOrder(assert.tradingAccount2, "BTC_EUR", "LIMIT", models.Buy, 9, 1, "GTC")
	ProcessTradeOrder(assert.tradingAccount1, "BTC_EUR", "LIMIT", models.Sell, 9, 1, "GTC")

	// then: it becomes activated
	assert.Equal(0, utils.GetCount("stop_order"))
	assert.Equal(1, GetBuyBookOrderCount())
	assert.Equal([]PriceVolume{{10.0, 2.0}}, GetVolumes("BTC_EUR", models.Buy))
	assert.Equal([]PriceVolume{}, GetVolumes("BTC_EUR", models.Sell))
}

func (assert *ServiceTestSuite) TestCreateStopLimitSellOrderActivateAndSettle() {

	// when: a stop limit order is created and then a crossing trade occurs
	ProcessTradeOrder(assert.tradingAccount1, "BTC_EUR", models.StopLimit, models.Sell, 10, 2, "GTC")
	ProcessTradeOrder(assert.tradingAccount1, "BTC_EUR", "LIMIT", models.Sell, 10, 1, "GTC")
	ProcessTradeOrder(assert.tradingAccount2, "BTC_EUR", "LIMIT", models.Buy, 10, 3, "GTC")

	// then: the order becomes activated and settled
	assert.Equal(0, utils.GetCount("stop_order"))
	assert.Equal(0, GetSellBookOrderCount())
	assert.Equal([]PriceVolume{}, GetVolumes("BTC_EUR", models.Buy))
	assert.Equal([]PriceVolume{}, GetVolumes("BTC_EUR", models.Sell))

}

func (assert *ServiceTestSuite) TestCreateStopLimitBuyOrderActivateAndSettle() {

	// when: a stop limit order is created and then a crossing trade occurs
	ProcessTradeOrder(assert.tradingAccount1, "BTC_EUR", models.StopLimit, models.Buy, 10, 2, "GTC")
	ProcessTradeOrder(assert.tradingAccount1, "BTC_EUR", "LIMIT", models.Buy, 10, 1, "GTC")
	ProcessTradeOrder(assert.tradingAccount2, "BTC_EUR", "LIMIT", models.Sell, 10, 3, "GTC")

	// then: the order becomes activated and settled
	assert.Equal(0, utils.GetCount("stop_order"))
	assert.Equal(0, GetSellBookOrderCount())
	assert.Equal([]PriceVolume{}, GetVolumes("BTC_EUR", models.Buy))
	assert.Equal([]PriceVolume{}, GetVolumes("BTC_EUR", models.Sell))
}

func (assert *ServiceTestSuite) TestCreateStopLimitSellOrderActivateAndSettleOppositeSide() {

	// when: a stop limit order is created and then a crossing trade occurs
	ProcessTradeOrder(assert.tradingAccount1, "BTC_EUR", models.StopLimit, models.Sell, 10, 2, "GTC")
	ProcessTradeOrder(assert.tradingAccount2, "BTC_EUR", "LIMIT", models.Buy, 10, 3, "GTC")
	ProcessTradeOrder(assert.tradingAccount1, "BTC_EUR", "LIMIT", models.Sell, 10, 1, "GTC")

	// then: the order becomes activated and settled
	assert.Equal(0, utils.GetCount("stop_order"))
	assert.Equal(0, GetSellBookOrderCount())
	assert.Equal([]PriceVolume{}, GetVolumes("BTC_EUR", models.Buy))
	assert.Equal([]PriceVolume{}, GetVolumes("BTC_EUR", models.Sell))
}

func (assert *ServiceTestSuite) TestCreateStopLimitBuyOrderActivateAndSettleOppositeSide() {

	// when: a stop limit order is created and then a crossing trade occurs
	ProcessTradeOrder(assert.tradingAccount1, "BTC_EUR", models.StopLimit, models.Buy, 10, 2, "GTC")
	ProcessTradeOrder(assert.tradingAccount2, "BTC_EUR", "LIMIT", models.Sell, 10, 3, "GTC")
	ProcessTradeOrder(assert.tradingAccount1, "BTC_EUR", "LIMIT", models.Buy, 10, 1, "GTC")

	// then: the order becomes activated and settled
	assert.Equal(0, utils.GetCount("stop_order"))
	assert.Equal(0, GetSellBookOrderCount())
	assert.Equal([]PriceVolume{}, GetVolumes("BTC_EUR", models.Buy))
	assert.Equal([]PriceVolume{}, GetVolumes("BTC_EUR", models.Sell))
}

func (assert *ServiceTestSuite) TestCreateStopLimitSellOrderActivateAndSettleBeforeWorsePriceOrders() {

	// when: a stop limit order is created and then a crossing trade occurs
	ProcessTradeOrder(assert.tradingAccount1, "BTC_EUR", models.StopLimit, models.Sell, 10, 1, "GTC")
	ProcessTradeOrder(assert.tradingAccount1, "BTC_EUR", "LIMIT", models.Sell, 11, 1, "GTC")
	ProcessTradeOrder(assert.tradingAccount1, "BTC_EUR", "LIMIT", models.Sell, 10, 1, "GTC")
	ProcessTradeOrder(assert.tradingAccount2, "BTC_EUR", "LIMIT", models.Buy, 11, 3, "GTC")

	// then: the order becomes activated and settled
	assert.Equal([]float64{10.0, 10.0, 11.0}, GetTradePrices())
	assert.Equal(0, utils.GetCount("stop_order"))
	assert.Equal(0, GetSellBookOrderCount())
	assert.Equal([]PriceVolume{}, GetVolumes("BTC_EUR", models.Buy))
	assert.Equal([]PriceVolume{}, GetVolumes("BTC_EUR", models.Sell))
}

func (assert *ServiceTestSuite) TestCreateStopLimitBuyOrderActivateAndSettleBeforeWorsePriceOrders() {

	// when: a stop limit order is created and then a crossing trade occurs
	ProcessTradeOrder(assert.tradingAccount1, "BTC_EUR", models.StopLimit, models.Buy, 10, 1, "GTC")
	ProcessTradeOrder(assert.tradingAccount1, "BTC_EUR", "LIMIT", models.Buy, 9, 1, "GTC")
	ProcessTradeOrder(assert.tradingAccount1, "BTC_EUR", "LIMIT", models.Buy, 10, 1, "GTC")
	ProcessTradeOrder(assert.tradingAccount2, "BTC_EUR", "LIMIT", models.Sell, 9, 3, "GTC")
	//
	// then: the order becomes activated and settled
	assert.Equal([]float64{10.0, 10.0, 9.0}, GetTradePrices())
	assert.Equal(0, utils.GetCount("stop_order"))
	assert.Equal(0, GetSellBookOrderCount())
	assert.Equal([]PriceVolume{}, GetVolumes("BTC_EUR", models.Buy))
	assert.Equal([]PriceVolume{}, GetVolumes("BTC_EUR", models.Sell))
}

func (assert *ServiceTestSuite) TestCreateStopLimitSellOrderActivateByMarketAndSettle() {

	// when: a stop limit order is created and then a crossing trade occurs
	ProcessTradeOrder(assert.tradingAccount1, "BTC_EUR", models.StopLimit, models.Sell, 10, 2, "GTC")
	ProcessTradeOrder(assert.tradingAccount1, "BTC_EUR", models.Market, models.Sell, 0, 1, "GTC")
	ProcessTradeOrder(assert.tradingAccount2, "BTC_EUR", "LIMIT", models.Buy, 10, 3, "GTC")

	// then: the order becomes activated and settled
	assert.Equal(0, utils.GetCount("stop_order"))
	assert.Equal(0, GetSellBookOrderCount())
	assert.Equal([]PriceVolume{}, GetVolumes("BTC_EUR", models.Buy))
	assert.Equal([]PriceVolume{}, GetVolumes("BTC_EUR", models.Sell))

	// when: a stop limit order is created and then a crossing trade occurs
	ProcessTradeOrder(assert.tradingAccount1, "BTC_EUR", models.StopLimit, models.Sell, 10, 2, "GTC")
	ProcessTradeOrder(assert.tradingAccount2, "BTC_EUR", "LIMIT", models.Buy, 10, 3, "GTC")
	ProcessTradeOrder(assert.tradingAccount1, "BTC_EUR", models.Market, models.Sell, 0, 1, "GTC")

	// then: the order becomes activated and settled
	assert.Equal(0, utils.GetCount("stop_order"))
	assert.Equal(0, GetSellBookOrderCount())
	assert.Equal([]PriceVolume{}, GetVolumes("BTC_EUR", models.Buy))
	assert.Equal([]PriceVolume{}, GetVolumes("BTC_EUR", models.Sell))
}

func (assert *ServiceTestSuite) TestCreateStopLimitBuyOrderActivateByMarketAndSettle() {

	// when: a stop limit order is created and then a crossing trade occurs
	ProcessTradeOrder(assert.tradingAccount1, "BTC_EUR", models.StopLimit, models.Buy, 10, 2, "GTC")
	ProcessTradeOrder(assert.tradingAccount1, "BTC_EUR", models.Market, models.Buy, 0, 10, "GTC")
	ProcessTradeOrder(assert.tradingAccount2, "BTC_EUR", "LIMIT", models.Sell, 10, 3, "GTC")
	//
	// then: the order becomes activated and settled
	assert.Equal(0, utils.GetCount("stop_order"))
	assert.Equal(0, GetSellBookOrderCount())
	assert.Equal([]PriceVolume{}, GetVolumes("BTC_EUR", models.Buy))
	assert.Equal([]PriceVolume{}, GetVolumes("BTC_EUR", models.Sell))
	//
	// when: a stop limit order is created and then a crossing trade occurs
	ProcessTradeOrder(assert.tradingAccount1, "BTC_EUR", models.StopLimit, models.Buy, 10, 2, "GTC")
	ProcessTradeOrder(assert.tradingAccount2, "BTC_EUR", "LIMIT", models.Sell, 10, 3, "GTC")
	ProcessTradeOrder(assert.tradingAccount1, "BTC_EUR", models.Market, models.Buy, 0, 10, "GTC")
	//
	// then: the order becomes activated and settled
	assert.Equal(0, utils.GetCount("stop_order"))

	assert.Equal(0, GetSellBookOrderCount())
	assert.Equal([]PriceVolume{}, GetVolumes("BTC_EUR", models.Buy))
	assert.Equal([]PriceVolume{}, GetVolumes("BTC_EUR", models.Sell))
}
