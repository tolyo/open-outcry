package services

//    These tests apply to balance constraints of order matching as we never want to allow
//    order into the order book that do no have sufficient leverage.
//    For limit orders, we ensure that the seller has sufficient amount in base currency or instument
//    and that the buys has sufficient amount in quote currency which is limit price times base currency or instrument.
//
//    For market order we ensure that a seller instrument amount is valid or buyer quote currency amount is valid.
//  `
//
//  test "process/1 limit sell order save with insufficient funds" {
// given:
//    appEntityId := CreateClient()
//    PaymentAccount.create(appEntityId, "BTC")
//    PaymentService.deposit(appEntityId, 100, "BTC", "test", "Test")
//    tradingAccountId := TradingAccount.find_by_appEntityId(appEntityId).id
//    ProcessTradeOrder(tradingAccountId, "BTC_EUR", "LIMIT", "SELL", 10, 100, "GTC")
//    PaymentService.deposit(appEntityId, 100, "BTC", "test", "Test")
// when: a limit order is sent with insufficient funds
//
// then: exception is raised
//    assert_raise Postgrex.Error, fn ->
//   this should be causing a failure
//      ProcessTradeOrder(tradingAccountId, "BTC_EUR", "LIMIT", "SELL", 10, 101, "GTC")
//    }
//  }
//
//  test "process/1 limit buy order save with insufficient funds" {
// given:
//    appEntityId := create_client()
//    PaymentService.deposit(appEntityId, 100, "EUR", "test", "Test")
//    tradingAccountId := TradingAccount.find_by_appEntityId(appEntityId).id
//    ProcessTradeOrder(tradingAccountId, "BTC_EUR", "LIMIT", "BUY", 10, 10, "GTC")
//    PaymentService.deposit(appEntityId, 100, "EUR", "test", "Test")
// when: a limit order is sent with insufficient funds
//
// then: exception is raised
//    assert_raise Postgrex.Error, fn ->
//      ProcessTradeOrder(tradingAccountId, "BTC_EUR", "LIMIT", "BUY", 10, 11, "GTC")
//    }
//  }
//
//  test "process/1 market sell order save with insufficient funds" {
// given:
//    appEntityId := CreateCLient()
//    PaymentAccount.create(appEntityId, "BTC")
//    PaymentService.deposit(appEntityId, 100, "BTC", "test", "Test")
//    tradingAccountId := TradingAccount.find_by_appEntityId(appEntityId).id
//    ProcessTradeOrder(tradingAccountId, "BTC_EUR", models.Market, "SELL", 100, "GTC")
//    PaymentService.deposit(appEntityId, 100, "BTC", "test", "Test")
// when: a market order is sent with insufficient funds
//
// then: exception is raised
//    assert_raise Postgrex.Error, fn ->
//      ProcessTradeOrder(tradingAccountId, "BTC_EUR", models.Market, "SELL", 101, "GTC")
//    }
//  }
//
//  test "process/1 market buy order save with insufficient funds" {
// given:
//    appEntityId := create_client()
//    PaymentService.deposit(appEntityId, 100, "EUR", "test", "Test")
//    tradingAccountId := TradingAccount.find_by_appEntityId(appEntityId).id
//    ProcessTradeOrder(tradingAccountId, "BTC_EUR", models.Market, "BUY", 100, "GTC")
//    PaymentService.deposit(appEntityId, 100, "EUR", "test", "Test")
// when: a market order is sent with insufficient funds
//
// then: exception is raised
//    assert_raise Postgrex.Error, fn ->
//      ProcessTradeOrder(tradingAccountId, "BTC_EUR", models.Market, "BUY", 101, "GTC")
//    }
//  }
//}
