package services

import (
	"open-outcry/pkg/models"
	"open-outcry/pkg/utils"
)

func (assert *ServiceTestSuite) TestCreateStopLossSellOrderSave() {
	tradingAccount := Acc()

	// when: a stop loss order is created
	res := ProcessTradeOrder(tradingAccount, "BTC_EUR", models.StopLoss, "SELL", 10, 100, "GTC")

	// then:
	assert.NotNil(res)
	assert.Equal(1, utils.GetCount("stop_order"))
	assert.Equal(0, GetSellBookOrderCount())
	assert.Equal([]PriceVolume{}, GetVolumes("BTC_EUR", "SELL"))

	assert.Equal(100.00, models.FindPaymentAccountByAppEntityIdAndCurrencyName(GetAppEntityId(), "BTC").AmountReserved)
}

func (assert *ServiceTestSuite) TestCreateStoLossBuyOrderBuy() {
	tradingAccount := Acc()
	// when: a stop loss order is created
	res := ProcessTradeOrder(tradingAccount, "BTC_EUR", models.StopLoss, "BUY", 10, 100, "GTC")

	// then:
	assert.NotNil(res)
	assert.Equal(1, utils.GetCount("stop_order"))
	assert.Equal(0, GetBuyBookOrderCount())
	assert.Equal([]PriceVolume{}, GetVolumes("BTC_EUR", "BUY"))
	assert.Equal(1000.00, models.FindPaymentAccountByAppEntityIdAndCurrencyName(GetAppEntityId(), "EUR").AmountReserved)
}
func (assert *ServiceTestSuite) TestCreateStopLossSellOrderActivate() {
	tradingAccount := Acc()
	tradingAccount2 := Acc2()
	// when: a stop loss order is created and then a crossing trade occurs
	ProcessTradeOrder(tradingAccount, "BTC_EUR", models.StopLoss, "SELL", 10, 2, "GTC")

	// then
	assert.Equal(2.0, models.FindPaymentAccountByAppEntityIdAndCurrencyName(GetAppEntityId(), "BTC").AmountReserved)

	ProcessTradeOrder(tradingAccount, "BTC_EUR", "LIMIT", "SELL", 10, 1, "GTC")
	ProcessTradeOrder(tradingAccount2, "BTC_EUR", "LIMIT", "BUY", 10, 1, "GTC")

	// then: the order becomes activated as a market order which is invisible to the order book
	assert.Equal(0, utils.GetCount("stop_order"))
	assert.Equal(1, GetSellBookOrderCount())
	assert.Equal([]PriceVolume{}, GetVolumes("BTC_EUR", "SELL"))

	// activation has no affect on AmountReserved
	assert.Equal(2.0, models.FindPaymentAccountByAppEntityIdAndCurrencyName(GetAppEntityId(), "BTC").AmountReserved)
}

func (assert *ServiceTestSuite) TestCreateStopLosBuyOrderActivate() {
	tradingAccount := Acc()
	tradingAccount2 := Acc2()

	// when: a stop loss order is created and then a crossing trade occurs
	ProcessTradeOrder(tradingAccount, "BTC_EUR", models.StopLoss, "BUY", 10, 2, "GTC")
	ProcessTradeOrder(tradingAccount2, "BTC_EUR", "LIMIT", "BUY", 10, 1, "GTC")
	ProcessTradeOrder(tradingAccount, "BTC_EUR", "LIMIT", "SELL", 10, 1, "GTC")

	assert.Equal(20.0, models.FindPaymentAccountByAppEntityIdAndCurrencyName(GetAppEntityId(), "EUR").AmountReserved)

	// then: the order becomes activated as a market order which is invisible to the order book
	assert.Equal(0, utils.GetCount("stop_order"))
	assert.Equal(1, GetBuyBookOrderCount())
	assert.Equal([]PriceVolume{}, GetVolumes("BTC_EUR", "BUY"))

	assert.Equal(20.0, models.FindPaymentAccountByAppEntityIdAndCurrencyName(GetAppEntityId(), "EUR").AmountReserved)
}

func (assert *ServiceTestSuite) TestCreateStopLossSellOrderNonCrossing() {
	tradingAccount := Acc()
	tradingAccount2 := Acc2()

	// when: a stop loss order is created and then a non-crossing trade occurs
	ProcessTradeOrder(tradingAccount, "BTC_EUR", models.StopLoss, "SELL", 10, 2, "GTC")
	ProcessTradeOrder(tradingAccount, "BTC_EUR", "LIMIT", "SELL", 11, 1, "GTC")
	ProcessTradeOrder(tradingAccount2, "BTC_EUR", "LIMIT", "BUY", 11, 1, "GTC")

	// then: the order remains deactivated
	assert.Equal(1, utils.GetCount("stop_order"))
	assert.Equal(0, GetSellBookOrderCount())
	assert.Equal([]PriceVolume{}, GetVolumes("BTC_EUR", "SELL"))

	// when: a crossing trade occurs
	ProcessTradeOrder(tradingAccount, "BTC_EUR", "LIMIT", "SELL", 9, 1, "GTC")
	ProcessTradeOrder(tradingAccount2, "BTC_EUR", "LIMIT", "BUY", 9, 1, "GTC")

	// then: it becomes activated
	assert.Equal(0, utils.GetCount("stop_order"))
	assert.Equal(1, GetSellBookOrderCount())
	assert.Equal([]PriceVolume{}, GetVolumes("BTC_EUR", "SELL"))
}

func (assert *ServiceTestSuite) TestCreateStopLossBuyOrderNonCrossing() {
	tradingAccount := Acc()
	tradingAccount2 := Acc2()

	// when: a stop loss order is created and then a non-crossing trade occurs
	ProcessTradeOrder(tradingAccount, "BTC_EUR", models.StopLoss, "BUY", 10, 2, "GTC")
	ProcessTradeOrder(tradingAccount2, "BTC_EUR", "LIMIT", "BUY", 11, 1, "GTC")
	ProcessTradeOrder(tradingAccount, "BTC_EUR", "LIMIT", "SELL", 11, 1, "GTC")

	// then: the order remains deactivated
	assert.Equal(1, utils.GetCount("stop_order"))
	assert.Equal(0, GetSellBookOrderCount())
	assert.Equal([]PriceVolume{}, GetVolumes("BTC_EUR", "SELL"))

	// when: a crossing trade occurs
	ProcessTradeOrder(tradingAccount2, "BTC_EUR", "LIMIT", "BUY", 9, 1, "GTC")
	ProcessTradeOrder(tradingAccount, "BTC_EUR", "LIMIT", "SELL", 9, 1, "GTC")

	// then: it becomes activated
	assert.Equal(0, utils.GetCount("stop_order"))
	assert.Equal(1, GetBuyBookOrderCount())
	assert.Equal([]PriceVolume{}, GetVolumes("BTC_EUR", "BUY"))
}

func (assert *ServiceTestSuite) TestCreateStopLossSellOrderActivateAndSettle() {
	tradingAccount := Acc()
	tradingAccount2 := Acc2()

	// when: a stop loss order is created and then a crossing trade occurs
	ProcessTradeOrder(tradingAccount, "BTC_EUR", models.StopLoss, "SELL", 10, 2, "GTC")
	ProcessTradeOrder(tradingAccount, "BTC_EUR", "LIMIT", "SELL", 10, 1, "GTC")
	ProcessTradeOrder(tradingAccount2, "BTC_EUR", "LIMIT", "BUY", 10, 3, "GTC")

	// then: the order becomes activated and settled
	assert.Equal(0, utils.GetCount("stop_order"))
	assert.Equal(0, GetSellBookOrderCount())
	assert.Equal([]PriceVolume{}, GetVolumes("BTC_EUR", "BUY"))
	assert.Equal([]PriceVolume{}, GetVolumes("BTC_EUR", "SELL"))
}

func (assert *ServiceTestSuite) TestCreateStopLossBuyOrderActivateAndSettle() {
	tradingAccount := Acc()
	tradingAccount2 := Acc2()

	// when: a stop loss order is created and then a crossing trade occurs
	ProcessTradeOrder(tradingAccount, "BTC_EUR", models.StopLoss, "BUY", 10, 20, "GTC")
	ProcessTradeOrder(tradingAccount, "BTC_EUR", "LIMIT", "BUY", 10, 1, "GTC")
	ProcessTradeOrder(tradingAccount2, "BTC_EUR", "LIMIT", "SELL", 10, 3, "GTC")

	// then: the order becomes activated and settled
	assert.Equal(0, utils.GetCount("stop_order"))
	assert.Equal(0, GetSellBookOrderCount())
	assert.Equal([]PriceVolume{}, GetVolumes("BTC_EUR", "BUY"))
	assert.Equal([]PriceVolume{}, GetVolumes("BTC_EUR", "SELL"))
}

func (assert *ServiceTestSuite) TestCreateStopLossSellOrderActivateAndSettleOppositeSide() {
	tradingAccount := Acc()
	tradingAccount2 := Acc2()

	// when: a stop loss order is created and then a crossing trade occurs
	ProcessTradeOrder(tradingAccount, "BTC_EUR", models.StopLoss, "SELL", 10, 2, "GTC")
	ProcessTradeOrder(tradingAccount2, "BTC_EUR", "LIMIT", "BUY", 10, 3, "GTC")
	ProcessTradeOrder(tradingAccount, "BTC_EUR", "LIMIT", "SELL", 10, 1, "GTC")

	// then: the order becomes activated and settled
	assert.Equal(0, utils.GetCount("stop_order"))
	assert.Equal(0, GetSellBookOrderCount())
	assert.Equal([]PriceVolume{}, GetVolumes("BTC_EUR", "BUY"))
	assert.Equal([]PriceVolume{}, GetVolumes("BTC_EUR", "SELL"))
}

func (assert *ServiceTestSuite) TestCreateStopLossBuyOrderActivateAndSettleOppositeSide() {
	tradingAccount := Acc()
	tradingAccount2 := Acc2()

	// when: a stop loss order is created and then a crossing trade occurs
	ProcessTradeOrder(tradingAccount, "BTC_EUR", models.StopLoss, "BUY", 10, 20, "GTC")
	ProcessTradeOrder(tradingAccount2, "BTC_EUR", "LIMIT", "SELL", 10, 3, "GTC")
	ProcessTradeOrder(tradingAccount, "BTC_EUR", "LIMIT", "BUY", 10, 1, "GTC")

	// then: the order becomes activated and settled
	assert.Equal(0, utils.GetCount("stop_order"))
	assert.Equal(0, GetSellBookOrderCount())
	assert.Equal([]PriceVolume{}, GetVolumes("BTC_EUR", "BUY"))
	assert.Equal([]PriceVolume{}, GetVolumes("BTC_EUR", "SELL"))
}

func (assert *ServiceTestSuite) TestCreateStopLossSellOrderActivateAndSettleBeforeWorsePriceOrders() {
	tradingAccount := Acc()
	tradingAccount2 := Acc2()

	// when: a stop loss order is created and then a crossing trade occurs
	ProcessTradeOrder(tradingAccount, "BTC_EUR", models.StopLoss, "SELL", 10, 1, "GTC")
	ProcessTradeOrder(tradingAccount, "BTC_EUR", "LIMIT", "SELL", 11, 1, "GTC")
	ProcessTradeOrder(tradingAccount, "BTC_EUR", "LIMIT", "SELL", 10, 1, "GTC")
	ProcessTradeOrder(tradingAccount2, "BTC_EUR", "LIMIT", "BUY", 11, 3, "GTC")

	// then: the order becomes activated and settled at last trade price which is 10
	assert.Equal([]float64{10.0, 10.0, 11.0}, GetTradePrices())
	assert.Equal(0, utils.GetCount("stop_order"))
	assert.Equal(0, GetSellBookOrderCount())
	assert.Equal([]PriceVolume{}, GetVolumes("BTC_EUR", "BUY"))
	assert.Equal([]PriceVolume{}, GetVolumes("BTC_EUR", "SELL"))
}

func (assert *ServiceTestSuite) TestCreateStopLossBuyOrderActivateAndSettleBeforeWorsePriceOrders() {
	tradingAccount := Acc()
	tradingAccount2 := Acc2()

	// when: a stop loss order is created and then a crossing trade occurs
	ProcessTradeOrder(tradingAccount, "BTC_EUR", models.StopLoss, "BUY", 10, 10, "GTC")
	ProcessTradeOrder(tradingAccount, "BTC_EUR", "LIMIT", "BUY", 9, 1, "GTC")
	ProcessTradeOrder(tradingAccount, "BTC_EUR", "LIMIT", "BUY", 10, 1, "GTC")
	ProcessTradeOrder(tradingAccount2, "BTC_EUR", "LIMIT", "SELL", 9, 3, "GTC")

	// then: the order becomes activated and settled
	assert.Equal([]float64{10.0, 10.0, 9.0}, GetTradePrices())
	assert.Equal(0, utils.GetCount("stop_order"))
	assert.Equal(0, GetSellBookOrderCount())
	assert.Equal([]PriceVolume{}, GetVolumes("BTC_EUR", "BUY"))
	assert.Equal([]PriceVolume{}, GetVolumes("BTC_EUR", "SELL"))
}

func (assert *ServiceTestSuite) TestCreateStopLossSellOrderActivateByMarketAndSettle() {
	tradingAccount := Acc()
	tradingAccount2 := Acc2()

	// when: a stop loss order is created and then a crossing trade occurs
	ProcessTradeOrder(tradingAccount, "BTC_EUR", models.StopLoss, "SELL", 10, 2, "GTC")
	ProcessTradeOrder(tradingAccount, "BTC_EUR", models.Market, "SELL", 0, 1, "GTC")
	ProcessTradeOrder(tradingAccount2, "BTC_EUR", "LIMIT", "BUY", 10, 3, "GTC")

	// then: the order becomes activated and settled
	assert.Equal(0, utils.GetCount("stop_order"))
	assert.Equal(0, GetSellBookOrderCount())
	assert.Equal([]PriceVolume{}, GetVolumes("BTC_EUR", "BUY"))
	assert.Equal([]PriceVolume{}, GetVolumes("BTC_EUR", "SELL"))

	// when: a stop loss order is created and then a crossing trade occurs
	ProcessTradeOrder(tradingAccount, "BTC_EUR", models.StopLoss, "SELL", 10, 2, "GTC")
	ProcessTradeOrder(tradingAccount2, "BTC_EUR", "LIMIT", "BUY", 10, 3, "GTC")
	ProcessTradeOrder(tradingAccount, "BTC_EUR", models.Market, "SELL", 0, 1, "GTC")

	// then: the order becomes activated and settled
	assert.Equal(0, utils.GetCount("stop_order"))
	assert.Equal(0, GetSellBookOrderCount())
	assert.Equal([]PriceVolume{}, GetVolumes("BTC_EUR", "BUY"))
	assert.Equal([]PriceVolume{}, GetVolumes("BTC_EUR", "SELL"))
}

func (assert *ServiceTestSuite) TestCreateStopLossBuyOrderActivateByMarketAndSettle() {
	tradingAccount := Acc()
	tradingAccount2 := Acc2()

	//    # when: a stop loss order is created and then a crossing trade occurs
	ProcessTradeOrder(tradingAccount, "BTC_EUR", models.StopLoss, "BUY", 10, 20, "GTC")
	ProcessTradeOrder(tradingAccount, "BTC_EUR", models.Market, "BUY", 0, 10, "GTC")
	ProcessTradeOrder(tradingAccount2, "BTC_EUR", "LIMIT", "SELL", 10, 3, "GTC")

	// then: the order becomes activated and settled
	assert.Equal(0, utils.GetCount("stop_order"))
	assert.Equal(0, GetSellBookOrderCount())
	assert.Equal([]PriceVolume{}, GetVolumes("BTC_EUR", "BUY"))
	assert.Equal([]PriceVolume{}, GetVolumes("BTC_EUR", "SELL"))

	//  when: a stop loss order is created and then a crossing trade occurs
	ProcessTradeOrder(tradingAccount, "BTC_EUR", models.StopLoss, "BUY", 10, 20, "GTC")
	ProcessTradeOrder(tradingAccount2, "BTC_EUR", "LIMIT", "SELL", 10, 3, "GTC")
	ProcessTradeOrder(tradingAccount, "BTC_EUR", models.Market, "BUY",  0, 10, "GTC")

	// then: the order becomes activated and settled
	assert.Equal(0, utils.GetCount("stop_order"))
	assert.Equal(0, GetSellBookOrderCount())
	assert.Equal([]PriceVolume{}, GetVolumes("BTC_EUR", "BUY"))
	assert.Equal([]PriceVolume{}, GetVolumes("BTC_EUR", "SELL"))
}