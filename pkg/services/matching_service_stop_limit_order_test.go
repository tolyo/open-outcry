package services


func (assert *ServiceTestSuite) TestCreateStopLimitSellOrderSave() {
// when: a stop limit order is created
//    res, _ := ProcessTradeOrder(c.tradingAccount, "BTC_EUR", models.StopLimit, "SELL", 10, 100, "GTC")
//
// then:
//    assert.Equal(res != nil
//    assert.Equal(utils.GetCount("stop_order") == 1
//    assert.Equal(GetSellBookOrderCount() == 0
//    assert.Equal(GetVolumes("BTC_EUR", "SELL"))
//
//    assert.Equal(FindPaymentAccountByAppEntityIdAndCurrencyName(GetAppEntityId(), "BTC").AmountReserved
//           |> Decimal.to_float() == 100
}

func (assert *ServiceTestSuite) TestCreateStopLimitBuyOrderBuy() {
// when: a stop limit order is created
//    res, _ := ProcessTradeOrder(c.tradingAccount, "BTC_EUR", models.StopLimit, "BUY", 10, 100, "GTC")
//
// then:
//    assert.Equal(res != nil
//    assert.Equal(utils.GetCount("stop_order") == 1
//    assert.Equal(GetBuyBookOrderCount() == 0
//    assert.Equal(GetVolumes("BTC_EUR", "BUY"))
//
//    assert.Equal(FindPaymentAccountByAppEntityIdAndCurrencyName(GetAppEntityId(), "EUR").AmountReserved
//           |> Decimal.to_float() == 1000
}

func (assert *ServiceTestSuite) TestCreateStopLimitSellOrderActivate() {
// when: a stop limit order is created and then a crossing trade occurs
//    ProcessTradeOrder(c.tradingAccount, "BTC_EUR", models.StopLimit, "SELL", 10, 2, "GTC")
//    ProcessTradeOrder(c.tradingAccount, "BTC_EUR", "LIMIT", "SELL", 10, 1, "GTC")
//    ProcessTradeOrder(c.tradingAccount2, "BTC_EUR", "LIMIT", "BUY", 10, 1, "GTC")
//
// then: the order becomes activated
//    assert.Equal(utils.GetCount("stop_order") == 0
//    assert.Equal(GetSellBookOrderCount() == 1
//    assert.Equal(GetVolumes("BTC_EUR", "SELL") == [{10, 2}]
}

func (assert *ServiceTestSuite) TestCreateStopLimitBuyOrderActivate() {
// when: a stop limit order is created and then a crossing trade occurs
//    ProcessTradeOrder(c.tradingAccount, "BTC_EUR", models.StopLimit, "BUY", 10, 2, "GTC")
//    ProcessTradeOrder(c.tradingAccount2, "BTC_EUR", "LIMIT", "BUY", 10, 1, "GTC")
//    ProcessTradeOrder(c.tradingAccount, "BTC_EUR", "LIMIT", "SELL", 10, 1, "GTC")
//
// then: the order becomes activated
//    assert.Equal(utils.GetCount("stop_order") == 0
//    assert.Equal(GetBuyBookOrderCount() == 1
//    assert.Equal(GetVolumes("BTC_EUR", "BUY") == [{10, 2}]
}

func (assert *ServiceTestSuite) TestCreateStopLimitSellOrderNonCrossing() {
// when: a stop limit order is created and then a non-crossing trade occurs
//    ProcessTradeOrder(c.tradingAccount, "BTC_EUR", models.StopLimit, "SELL", 10, 2, "GTC")
//    ProcessTradeOrder(c.tradingAccount, "BTC_EUR", "LIMIT", "SELL", 11, 1, "GTC")
//    ProcessTradeOrder(c.tradingAccount2, "BTC_EUR", "LIMIT", "BUY", 11, 1, "GTC")
//
// then: the order remains deactivated
//    assert.Equal(utils.GetCount("stop_order") == 1
//    assert.Equal(GetSellBookOrderCount() == 0
//    assert.Equal(GetVolumes("BTC_EUR", "SELL"))
//
// when: a crossing trade occurs
//    ProcessTradeOrder(c.tradingAccount, "BTC_EUR", "LIMIT", "SELL", 9, 1, "GTC")
//    ProcessTradeOrder(c.tradingAccount2, "BTC_EUR", "LIMIT", "BUY", 9, 1, "GTC")
//
// then: it becomes activated
//    assert.Equal(utils.GetCount("stop_order") == 0
//    assert.Equal(GetSellBookOrderCount() == 1
//    assert.Equal(GetVolumes("BTC_EUR", "SELL") == [{10, 2}]
}

func (assert *ServiceTestSuite) TestCreateStopLimitBuyOrderNonCrossing() {
// when: a stop limit order is created and then a non-crossing trade occurs
//    ProcessTradeOrder(c.tradingAccount, "BTC_EUR", models.StopLimit, "BUY", 10, 2, "GTC")
//    ProcessTradeOrder(c.tradingAccount2, "BTC_EUR", "LIMIT", "BUY", 11, 1, "GTC")
//    ProcessTradeOrder(c.tradingAccount, "BTC_EUR", "LIMIT", "SELL", 11, 1, "GTC")
//
// then: the order remains deactivated
//    assert.Equal(utils.GetCount("stop_order") == 1
//    assert.Equal(GetSellBookOrderCount() == 0
//    assert.Equal(GetVolumes("BTC_EUR", "SELL"))
//
// when: a crossing trade occurs
//    ProcessTradeOrder(c.tradingAccount2, "BTC_EUR", "LIMIT", "BUY", 9, 1, "GTC")
//    ProcessTradeOrder(c.tradingAccount, "BTC_EUR", "LIMIT", "SELL", 9, 1, "GTC")
//
// then: it becomes activated
//    assert.Equal(utils.GetCount("stop_order") == 0
//    assert.Equal(GetBuyBookOrderCount() == 1
//    assert.Equal(GetVolumes("BTC_EUR", "BUY") == [{10, 2}]
}

func (assert *ServiceTestSuite) TestCreateStopLimitSellOrderActivateAndSettle() {
// when: a stop limit order is created and then a crossing trade occurs
//    ProcessTradeOrder(c.tradingAccount, "BTC_EUR", models.StopLimit, "SELL", 10, 2, "GTC")
//    ProcessTradeOrder(c.tradingAccount, "BTC_EUR", "LIMIT", "SELL", 10, 1, "GTC")
//    ProcessTradeOrder(c.tradingAccount2, "BTC_EUR", "LIMIT", "BUY", 10, 3, "GTC")
//
// then: the order becomes activated and settled
//    assert.Equal(utils.GetCount("stop_order") == 0
//    assert.Equal(GetSellBookOrderCount() == 0
//    assert.Equal(GetVolumes("BTC_EUR", "BUY"))
//    assert.Equal(GetVolumes("BTC_EUR", "SELL"))
}

func (assert *ServiceTestSuite) TestCreateStopLimitBuyOrderActivateAndSettle() {
// when: a stop limit order is created and then a crossing trade occurs
//    ProcessTradeOrder(c.tradingAccount, "BTC_EUR", models.StopLimit, "BUY", 10, 2, "GTC")
//    ProcessTradeOrder(c.tradingAccount, "BTC_EUR", "LIMIT", "BUY", 10, 1, "GTC")
//    ProcessTradeOrder(c.tradingAccount2, "BTC_EUR", "LIMIT", "SELL", 10, 3, "GTC")
//
// then: the order becomes activated and settled
//    assert.Equal(utils.GetCount("stop_order") == 0
//    assert.Equal(GetSellBookOrderCount() == 0
//    assert.Equal(GetVolumes("BTC_EUR", "BUY"))
//    assert.Equal(GetVolumes("BTC_EUR", "SELL"))
}

func (assert *ServiceTestSuite) TestCreateStopLimitSellOrderActivateAndSettleOppositeSide() {
// when: a stop limit order is created and then a crossing trade occurs
//    ProcessTradeOrder(c.tradingAccount, "BTC_EUR", models.StopLimit, "SELL", 10, 2, "GTC")
//    ProcessTradeOrder(c.tradingAccount2, "BTC_EUR", "LIMIT", "BUY", 10, 3, "GTC")
//    ProcessTradeOrder(c.tradingAccount, "BTC_EUR", "LIMIT", "SELL", 10, 1, "GTC")
//
// then: the order becomes activated and settled
//    assert.Equal(utils.GetCount("stop_order") == 0
//    assert.Equal(GetSellBookOrderCount() == 0
//    assert.Equal(GetVolumes("BTC_EUR", "BUY"))
//    assert.Equal(GetVolumes("BTC_EUR", "SELL"))
}

func (assert *ServiceTestSuite) TestCreateStopLimitBuyOrderActivateAndSettleOppositeSide() {
// when: a stop limit order is created and then a crossing trade occurs
//    ProcessTradeOrder(c.tradingAccount, "BTC_EUR", models.StopLimit, "BUY", 10, 2, "GTC")
//    ProcessTradeOrder(c.tradingAccount2, "BTC_EUR", "LIMIT", "SELL", 10, 3, "GTC")
//    ProcessTradeOrder(c.tradingAccount, "BTC_EUR", "LIMIT", "BUY", 10, 1, "GTC")
//
// then: the order becomes activated and settled
//    assert.Equal(utils.GetCount("stop_order") == 0
//    assert.Equal(GetSellBookOrderCount() == 0
//    assert.Equal(GetVolumes("BTC_EUR", "BUY"))
//    assert.Equal(GetVolumes("BTC_EUR", "SELL"))
}

func (assert *ServiceTestSuite) TestCreateStopLimitSellOrderActivateAndSettleBeforeWorsePriceOrders() {
// when: a stop limit order is created and then a crossing trade occurs
//    ProcessTradeOrder(c.tradingAccount, "BTC_EUR", models.StopLimit, "SELL", 10, 1, "GTC")
//    ProcessTradeOrder(c.tradingAccount, "BTC_EUR", "LIMIT", "SELL", 11, 1, "GTC")
//    ProcessTradeOrder(c.tradingAccount, "BTC_EUR", "LIMIT", "SELL", 10, 1, "GTC")
//    ProcessTradeOrder(c.tradingAccount2, "BTC_EUR", "LIMIT", "BUY", 11, 3, "GTC")
//
// then: the order becomes activated and settled
//    assert.Equal(GetTradePrices() == [10, 10, 11]
//    assert.Equal(utils.GetCount("stop_order") == 0
//    assert.Equal(GetSellBookOrderCount() == 0
//    assert.Equal(GetVolumes("BTC_EUR", "BUY"))
//    assert.Equal(GetVolumes("BTC_EUR", "SELL"))
}

func (assert *ServiceTestSuite) TestCreateStopLimitBuyOrderActivateAndSettleBeforeWorsePriceOrders() {
// when: a stop limit order is created and then a crossing trade occurs
//    ProcessTradeOrder(c.tradingAccount, "BTC_EUR", models.StopLimit, "BUY", 10, 1, "GTC")
//    ProcessTradeOrder(c.tradingAccount, "BTC_EUR", "LIMIT", "BUY", 9, 1, "GTC")
//    ProcessTradeOrder(c.tradingAccount, "BTC_EUR", "LIMIT", "BUY", 10, 1, "GTC")
//    ProcessTradeOrder(c.tradingAccount2, "BTC_EUR", "LIMIT", "SELL", 9, 3, "GTC")
//
// then: the order becomes activated and settled
//    assert.Equal(GetTradePrices() == [10, 10, 9]
//    assert.Equal(utils.GetCount("stop_order") == 0
//    assert.Equal(GetSellBookOrderCount() == 0
//    assert.Equal(GetVolumes("BTC_EUR", "BUY"))
//    assert.Equal(GetVolumes("BTC_EUR", "SELL"))
}

func (assert *ServiceTestSuite) TestCreateStopLimitSellOrderActivateByMarketAndSettle() {
// when: a stop limit order is created and then a crossing trade occurs
//    ProcessTradeOrder(c.tradingAccount, "BTC_EUR", models.StopLimit, "SELL", 10, 2, "GTC")
//    ProcessTradeOrder(c.tradingAccount, "BTC_EUR", models.Market, "SELL", 1, "GTC")
//    ProcessTradeOrder(c.tradingAccount2, "BTC_EUR", "LIMIT", "BUY", 10, 3, "GTC")
//
// then: the order becomes activated and settled
//    assert.Equal(utils.GetCount("stop_order") == 0
//    assert.Equal(GetSellBookOrderCount() == 0
//    assert.Equal(GetVolumes("BTC_EUR", "BUY"))
//    assert.Equal(GetVolumes("BTC_EUR", "SELL"))
//
// when: a stop limit order is created and then a crossing trade occurs
//    ProcessTradeOrder(c.tradingAccount, "BTC_EUR", models.StopLimit, "SELL", 10, 2, "GTC")
//    ProcessTradeOrder(c.tradingAccount2, "BTC_EUR", "LIMIT", "BUY", 10, 3, "GTC")
//    ProcessTradeOrder(c.tradingAccount, "BTC_EUR", models.Market, "SELL", 1, "GTC")
//
// then: the order becomes activated and settled
//    assert.Equal(utils.GetCount("stop_order") == 0
//    assert.Equal(GetSellBookOrderCount() == 0
//    assert.Equal(GetVolumes("BTC_EUR", "BUY"))
//    assert.Equal(GetVolumes("BTC_EUR", "SELL"))
}

func (assert *ServiceTestSuite) TestCreateStopLimitBuyOrderActivateByMarketAndSettle() {
// when: a stop limit order is created and then a crossing trade occurs
//    ProcessTradeOrder(c.tradingAccount, "BTC_EUR", models.StopLimit, "BUY", 10, 2, "GTC")
//    ProcessTradeOrder(c.tradingAccount, "BTC_EUR", models.Market, "BUY", 10, "GTC")
//    ProcessTradeOrder(c.tradingAccount2, "BTC_EUR", "LIMIT", "SELL", 10, 3, "GTC")
//
// then: the order becomes activated and settled
//    assert.Equal(utils.GetCount("stop_order") == 0
//    assert.Equal(GetSellBookOrderCount() == 0
//    assert.Equal(GetVolumes("BTC_EUR", "BUY"))
//    assert.Equal(GetVolumes("BTC_EUR", "SELL"))
//
// when: a stop limit order is created and then a crossing trade occurs
//    ProcessTradeOrder(c.tradingAccount, "BTC_EUR", models.StopLimit, "BUY", 10, 2, "GTC")
//    ProcessTradeOrder(c.tradingAccount2, "BTC_EUR", "LIMIT", "SELL", 10, 3, "GTC")
//    ProcessTradeOrder(c.tradingAccount, "BTC_EUR", models.Market, "BUY", 10, "GTC")
//
// then: the order becomes activated and settled
//    assert.Equal(utils.GetCount("stop_order") == 0
//
//    assert.Equal(GetSellBookOrderCount() == 0
//    assert.Equal(GetVolumes("BTC_EUR", "BUY"))
//    assert.Equal(GetVolumes("BTC_EUR", "SELL"))
//  }
}