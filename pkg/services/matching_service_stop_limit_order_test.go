package services

//funcmodule MatchingServiceStopLimitOrderTest {
//  use DataCase
//
//  import TestUtils
//
//  setup {
//    [
//      tradingAccount: Acc(),
//      entity: get_appEntityId(),
//      tradingAccount2: Acc2()
//    ]
//  }
//
//  test "create/1 stop limit sell order save", c {
// when: a stop limit order is created
//    res = ProcessTradeOrder(c.tradingAccount, "BTC_EUR", models.StopLimit, "SELL", 10, 100, "GTC")
//
// then:
//    assert.Equal(res != nil
//    assert.Equal(utils.GetCount("stop_order") == 1
//    assert.Equal(GetSellBookOrderCount() == 0
//    assert.Equal(GetVolumes("BTC_EUR", "SELL") == []
//
//    assert.Equal(FindPaymentAccountByAppEntityIdAndCurrencyName(GetAppEntityId(), "BTC").AmountReserved
//           |> Decimal.to_float() == 100
//  }
//
//  test "create/1 stop limit buy order buy", c {
// when: a stop limit order is created
//    res = ProcessTradeOrder(c.tradingAccount, "BTC_EUR", models.StopLimit, "BUY", 10, 100, "GTC")
//
// then:
//    assert.Equal(res != nil
//    assert.Equal(utils.GetCount("stop_order") == 1
//    assert.Equal(GetBuyBookOrderCount() == 0
//    assert.Equal(GetVolumes("BTC_EUR", "BUY") == []
//
//    assert.Equal(FindPaymentAccountByAppEntityIdAndCurrencyName(GetAppEntityId(), "EUR").AmountReserved
//           |> Decimal.to_float() == 1000
//  }
//
//  test "create/1 stop limit sell order activate", c {
// when: a stop limit order is created and then a crossing trade occurs
//    ProcessTradeOrder(c.tradingAccount, "BTC_EUR", models.StopLimit, "SELL", 10, 2, "GTC")
//    ProcessTradeOrder(c.tradingAccount, "BTC_EUR", "LIMIT", "SELL", 10, 1, "GTC")
//    ProcessTradeOrder(c.tradingAccount2, "BTC_EUR", "LIMIT", "BUY", 10, 1, "GTC")
//
// then: the order becomes activated
//    assert.Equal(utils.GetCount("stop_order") == 0
//    assert.Equal(GetSellBookOrderCount() == 1
//    assert.Equal(GetVolumes("BTC_EUR", "SELL") == [{10, 2}]
//  }
//
//  test "create/1 stop limit buy order activate", c {
// when: a stop limit order is created and then a crossing trade occurs
//    ProcessTradeOrder(c.tradingAccount, "BTC_EUR", models.StopLimit, "BUY", 10, 2, "GTC")
//    ProcessTradeOrder(c.tradingAccount2, "BTC_EUR", "LIMIT", "BUY", 10, 1, "GTC")
//    ProcessTradeOrder(c.tradingAccount, "BTC_EUR", "LIMIT", "SELL", 10, 1, "GTC")
//
// then: the order becomes activated
//    assert.Equal(utils.GetCount("stop_order") == 0
//    assert.Equal(GetBuyBookOrderCount() == 1
//    assert.Equal(GetVolumes("BTC_EUR", "BUY") == [{10, 2}]
//  }
//
//  test "create/1 stop limit sell order non crossing ", c {
// when: a stop limit order is created and then a non-crossing trade occurs
//    ProcessTradeOrder(c.tradingAccount, "BTC_EUR", models.StopLimit, "SELL", 10, 2, "GTC")
//    ProcessTradeOrder(c.tradingAccount, "BTC_EUR", "LIMIT", "SELL", 11, 1, "GTC")
//    ProcessTradeOrder(c.tradingAccount2, "BTC_EUR", "LIMIT", "BUY", 11, 1, "GTC")
//
// then: the order remains deactivated
//    assert.Equal(utils.GetCount("stop_order") == 1
//    assert.Equal(GetSellBookOrderCount() == 0
//    assert.Equal(GetVolumes("BTC_EUR", "SELL") == []
//
// when: a crossing trade occurs
//    ProcessTradeOrder(c.tradingAccount, "BTC_EUR", "LIMIT", "SELL", 9, 1, "GTC")
//    ProcessTradeOrder(c.tradingAccount2, "BTC_EUR", "LIMIT", "BUY", 9, 1, "GTC")
//
// then: it becomes activated
//    assert.Equal(utils.GetCount("stop_order") == 0
//    assert.Equal(GetSellBookOrderCount() == 1
//    assert.Equal(GetVolumes("BTC_EUR", "SELL") == [{10, 2}]
//  }
//
//  test "create/1 stop limit buy order non crossing ", c {
// when: a stop limit order is created and then a non-crossing trade occurs
//    ProcessTradeOrder(c.tradingAccount, "BTC_EUR", models.StopLimit, "BUY", 10, 2, "GTC")
//    ProcessTradeOrder(c.tradingAccount2, "BTC_EUR", "LIMIT", "BUY", 11, 1, "GTC")
//    ProcessTradeOrder(c.tradingAccount, "BTC_EUR", "LIMIT", "SELL", 11, 1, "GTC")
//
// then: the order remains deactivated
//    assert.Equal(utils.GetCount("stop_order") == 1
//    assert.Equal(GetSellBookOrderCount() == 0
//    assert.Equal(GetVolumes("BTC_EUR", "SELL") == []
//
// when: a crossing trade occurs
//    ProcessTradeOrder(c.tradingAccount2, "BTC_EUR", "LIMIT", "BUY", 9, 1, "GTC")
//    ProcessTradeOrder(c.tradingAccount, "BTC_EUR", "LIMIT", "SELL", 9, 1, "GTC")
//
// then: it becomes activated
//    assert.Equal(utils.GetCount("stop_order") == 0
//    assert.Equal(GetBuyBookOrderCount() == 1
//    assert.Equal(GetVolumes("BTC_EUR", "BUY") == [{10, 2}]
//  }
//
//  test "create/1 stop limit sell order activate and settle", c {
// when: a stop limit order is created and then a crossing trade occurs
//    ProcessTradeOrder(c.tradingAccount, "BTC_EUR", models.StopLimit, "SELL", 10, 2, "GTC")
//    ProcessTradeOrder(c.tradingAccount, "BTC_EUR", "LIMIT", "SELL", 10, 1, "GTC")
//    ProcessTradeOrder(c.tradingAccount2, "BTC_EUR", "LIMIT", "BUY", 10, 3, "GTC")
//
// then: the order becomes activated and settled
//    assert.Equal(utils.GetCount("stop_order") == 0
//    assert.Equal(GetSellBookOrderCount() == 0
//    assert.Equal(GetVolumes("BTC_EUR", "BUY") == []
//    assert.Equal(GetVolumes("BTC_EUR", "SELL") == []
//  }
//
//  test "create/1 stop limit buy order activate and settle", c {
// when: a stop limit order is created and then a crossing trade occurs
//    ProcessTradeOrder(c.tradingAccount, "BTC_EUR", models.StopLimit, "BUY", 10, 2, "GTC")
//    ProcessTradeOrder(c.tradingAccount, "BTC_EUR", "LIMIT", "BUY", 10, 1, "GTC")
//    ProcessTradeOrder(c.tradingAccount2, "BTC_EUR", "LIMIT", "SELL", 10, 3, "GTC")
//
// then: the order becomes activated and settled
//    assert.Equal(utils.GetCount("stop_order") == 0
//    assert.Equal(GetSellBookOrderCount() == 0
//    assert.Equal(GetVolumes("BTC_EUR", "BUY") == []
//    assert.Equal(GetVolumes("BTC_EUR", "SELL") == []
//  }
//
//  test "create/1 stop limit sell order activate and settle opposite side", c {
// when: a stop limit order is created and then a crossing trade occurs
//    ProcessTradeOrder(c.tradingAccount, "BTC_EUR", models.StopLimit, "SELL", 10, 2, "GTC")
//    ProcessTradeOrder(c.tradingAccount2, "BTC_EUR", "LIMIT", "BUY", 10, 3, "GTC")
//    ProcessTradeOrder(c.tradingAccount, "BTC_EUR", "LIMIT", "SELL", 10, 1, "GTC")
//
// then: the order becomes activated and settled
//    assert.Equal(utils.GetCount("stop_order") == 0
//    assert.Equal(GetSellBookOrderCount() == 0
//    assert.Equal(GetVolumes("BTC_EUR", "BUY") == []
//    assert.Equal(GetVolumes("BTC_EUR", "SELL") == []
//  }
//
//  test "create/1 stop limit buy order activate and settle opposite side", c {
//    # when: a stop limit order is created and then a crossing trade occurs
//    ProcessTradeOrder(c.tradingAccount, "BTC_EUR", models.StopLimit, "BUY", 10, 2, "GTC")
//    ProcessTradeOrder(c.tradingAccount2, "BTC_EUR", "LIMIT", "SELL", 10, 3, "GTC")
//    ProcessTradeOrder(c.tradingAccount, "BTC_EUR", "LIMIT", "BUY", 10, 1, "GTC")
//
//    # then: the order becomes activated and settled
//    assert.Equal(utils.GetCount("stop_order") == 0
//    assert.Equal(GetSellBookOrderCount() == 0
//    assert.Equal(GetVolumes("BTC_EUR", "BUY") == []
//    assert.Equal(GetVolumes("BTC_EUR", "SELL") == []
//  }
//
//  test "create/1 stop limit sell order activate and settle before worse price orders", c {
//    # when: a stop limit order is created and then a crossing trade occurs
//    ProcessTradeOrder(c.tradingAccount, "BTC_EUR", models.StopLimit, "SELL", 10, 1, "GTC")
//    ProcessTradeOrder(c.tradingAccount, "BTC_EUR", "LIMIT", "SELL", 11, 1, "GTC")
//    ProcessTradeOrder(c.tradingAccount, "BTC_EUR", "LIMIT", "SELL", 10, 1, "GTC")
//    ProcessTradeOrder(c.tradingAccount2, "BTC_EUR", "LIMIT", "BUY", 11, 3, "GTC")
//
//    # then: the order becomes activated and settled
//    assert.Equal(GetTradePrices() == [10, 10, 11]
//    assert.Equal(utils.GetCount("stop_order") == 0
//    assert.Equal(GetSellBookOrderCount() == 0
//    assert.Equal(GetVolumes("BTC_EUR", "BUY") == []
//    assert.Equal(GetVolumes("BTC_EUR", "SELL") == []
//  }
//
//  test "create/1 stop limit buy order activate and settle before worse price orders", c {
//    # when: a stop limit order is created and then a crossing trade occurs
//    ProcessTradeOrder(c.tradingAccount, "BTC_EUR", models.StopLimit, "BUY", 10, 1, "GTC")
//    ProcessTradeOrder(c.tradingAccount, "BTC_EUR", "LIMIT", "BUY", 9, 1, "GTC")
//    ProcessTradeOrder(c.tradingAccount, "BTC_EUR", "LIMIT", "BUY", 10, 1, "GTC")
//    ProcessTradeOrder(c.tradingAccount2, "BTC_EUR", "LIMIT", "SELL", 9, 3, "GTC")
//
//    # then: the order becomes activated and settled
//    assert.Equal(GetTradePrices() == [10, 10, 9]
//    assert.Equal(utils.GetCount("stop_order") == 0
//    assert.Equal(GetSellBookOrderCount() == 0
//    assert.Equal(GetVolumes("BTC_EUR", "BUY") == []
//    assert.Equal(GetVolumes("BTC_EUR", "SELL") == []
//  }
//
//  test "create/1 stop limit sell order activate by market and settle", c {
//    # when: a stop limit order is created and then a crossing trade occurs
//    ProcessTradeOrder(c.tradingAccount, "BTC_EUR", models.StopLimit, "SELL", 10, 2, "GTC")
//    ProcessTradeOrder(c.tradingAccount, "BTC_EUR", models.Market, "SELL", 1, "GTC")
//    ProcessTradeOrder(c.tradingAccount2, "BTC_EUR", "LIMIT", "BUY", 10, 3, "GTC")
//
//    # then: the order becomes activated and settled
//    assert.Equal(utils.GetCount("stop_order") == 0
//    assert.Equal(GetSellBookOrderCount() == 0
//    assert.Equal(GetVolumes("BTC_EUR", "BUY") == []
//    assert.Equal(GetVolumes("BTC_EUR", "SELL") == []
//
//    # when: a stop limit order is created and then a crossing trade occurs
//    ProcessTradeOrder(c.tradingAccount, "BTC_EUR", models.StopLimit, "SELL", 10, 2, "GTC")
//    ProcessTradeOrder(c.tradingAccount2, "BTC_EUR", "LIMIT", "BUY", 10, 3, "GTC")
//    ProcessTradeOrder(c.tradingAccount, "BTC_EUR", models.Market, "SELL", 1, "GTC")
//
//    # then: the order becomes activated and settled
//    assert.Equal(utils.GetCount("stop_order") == 0
//    assert.Equal(GetSellBookOrderCount() == 0
//    assert.Equal(GetVolumes("BTC_EUR", "BUY") == []
//    assert.Equal(GetVolumes("BTC_EUR", "SELL") == []
//  }
//
//  test "create/1 stop limit buy order activate by market and settle", c {
//    # when: a stop limit order is created and then a crossing trade occurs
//    ProcessTradeOrder(c.tradingAccount, "BTC_EUR", models.StopLimit, "BUY", 10, 2, "GTC")
//    ProcessTradeOrder(c.tradingAccount, "BTC_EUR", models.Market, "BUY", 10, "GTC")
//    ProcessTradeOrder(c.tradingAccount2, "BTC_EUR", "LIMIT", "SELL", 10, 3, "GTC")
//
//    # then: the order becomes activated and settled
//    assert.Equal(utils.GetCount("stop_order") == 0
//    assert.Equal(GetSellBookOrderCount() == 0
//    assert.Equal(GetVolumes("BTC_EUR", "BUY") == []
//    assert.Equal(GetVolumes("BTC_EUR", "SELL") == []
//
//    # when: a stop limit order is created and then a crossing trade occurs
//    ProcessTradeOrder(c.tradingAccount, "BTC_EUR", models.StopLimit, "BUY", 10, 2, "GTC")
//    ProcessTradeOrder(c.tradingAccount2, "BTC_EUR", "LIMIT", "SELL", 10, 3, "GTC")
//    ProcessTradeOrder(c.tradingAccount, "BTC_EUR", models.Market, "BUY", 10, "GTC")
//
//    # then: the order becomes activated and settled
//    assert.Equal(utils.GetCount("stop_order") == 0
//
//    assert.Equal(GetSellBookOrderCount() == 0
//    assert.Equal(GetVolumes("BTC_EUR", "BUY") == []
//    assert.Equal(GetVolumes("BTC_EUR", "SELL") == []
//  }
//}
