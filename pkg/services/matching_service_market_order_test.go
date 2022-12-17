package services

//funcmodule MatchingServiceMarketOrderTest {
//  use DataCase
//
//  import TestUtils
//  use ExUnit.Case, async: false
//
//  @module{c `
//
//  Primary market order test cases:
//    * An incoming market order hits an outstanding limit order in the opposite book
//    * An incoming limit order hits an outstanding market order in the opposite book
//    * An incoming market order hits another market order of the opposite side
//
//  Additional use cases to test:
//    * If the opposite book {es not have outstanding limit orders,
//      then the trade settles at the incoming order’s limit price
//    * If the opposite book {es have limit orders, the trade settles at the better of two prices –
//      the incoming order’s limit or the best limit from the opposite book –
//      the term “better of two prices” is from the point of view of the incoming limit order.
//      In other words, if incoming limit order would have crossed with outstanding opposite
//      “best limit” order in the absence of market order,
//      then trade at that, potentially improved, “best limit” price.
//
//    * when a market order matches with another market order in the opposite order book, there are three possibilites to test:
//      - the opposite order book with the resting market order also contains one or more outstanding limit orders –
//        in this case the opposite book has a “best limit” price and this price becomes the price for the trade
//      - the opposte order book {es not have outstanding limit orders, so the “best limit” price is not funcined.
//        In this case the trade occurs at the “reference price”. Most often reference price is the last traded price for a security.
//      - if no trades have occured (upon security launch), market orders simply rest
//  `
//
//  test "process/1 market sell order save" {
// when: a market order is sent to an empty matching unit
//    ProcessTradeOrder(Acc(), "BTC_EUR", models.Market, "SELL", 100, "GTC")
//
// then: a matching unit should save the trade order but it should not be visible to the order book
//    assert.Equal(GetSellBookOrderCount() == 1
//    assert.Equal(GetVolumes("BTC_EUR", "SELL") == []
//  }
//
//  test "process/1 market buy order save" {
// when: a market order is sent to an empty matching unit
//    ProcessTradeOrder(Acc(), "BTC_EUR", models.Market, "BUY", 100, "GTC")
//
// then: a matching unit should save the trade order but it should not be visible to the order book
//    assert.Equal(GetBuyBookOrderCount() == 1
//    assert.Equal(GetVolumes("BTC_EUR", "BUY") == []
//  }
//
//  test "process/1 market sell order with existing sell limit" {
//    account = Acc()
//
// when: a market order is sent to an non empty matching unit
//    ProcessTradeOrder(account, "BTC_EUR", "LIMIT", "SELL", 10, 100, "GTC")
//    ProcessTradeOrder(account, "BTC_EUR", models.Market, "SELL", 100, "GTC")
//
// then: a matching unit should save the trade order but it should not be visible to the order book
//    assert.Equal(GetSellBookOrderCount() == 2
//
//    assert.Equal(GetVolumes("BTC_EUR", "SELL") == [
//             {10, 100}
//           ]
//  }
//
//  test "process/1 market buy order with existing buy limit" {
//    account = Acc()
//
// when: a market order is sent to an non empty matching unit
//    ProcessTradeOrder(account, "BTC_EUR", "LIMIT", "BUY", 10, 10, "GTC")
//    ProcessTradeOrder(account, "BTC_EUR", models.Market, "BUY", 10, "GTC")
//
// then: a matching unit should save the trade order but it should not be visible to the order book
//    assert.Equal(GetBuyBookOrderCount() == 2
//
//    assert.Equal(GetVolumes("BTC_EUR", "BUY") == [
//             {10.0, 10}
//           ]
//  }
//
//  test "process/1 market sell order with existing buy limit" {
//    account = Acc()
//    account2 = Acc2()
//
// when:
//    ProcessTradeOrder(account, "BTC_EUR", "LIMIT", "BUY", 10, 100, "GTC")
//    ProcessTradeOrder(account2, "BTC_EUR", models.Market, "SELL", 100, "GTC")
//
// then:
//    assert.Equal(get_trade_count() == 1
//    assert.Equal(GetTradePrices() == [10]
//    assert.Equal(GetBuyBookOrderCount() == 0
//    assert.Equal(GetSellBookOrderCount() == 0
//    assert.Equal(GetVolumes("BTC_EUR", "BUY") == []
//    assert.Equal(GetVolumes("BTC_EUR", "SELL") == []
//  }
//
//  test "process/1 market buy order with existing sell limit" {
//    account = Acc()
//    account2 = Acc2()
//
// when:
//    ProcessTradeOrder(account, "BTC_EUR", "LIMIT", "SELL", 10, 100, "GTC")
//    ProcessTradeOrder(account2, "BTC_EUR", models.Market, "BUY", 100, "GTC")
//
// then:
//    assert.Equal(get_trade_count() == 1
//    assert.Equal(GetTradePrices() == [10]
//    assert.Equal(GetBuyBookOrderCount() == 0
//    assert.Equal(GetSellBookOrderCount() == 1
//    assert.Equal(GetVolumes("BTC_EUR", "BUY") == []
//    assert.Equal(GetVolumes("BTC_EUR", "SELL") == [{10, 90}]
//  }
//
//  test "process/1 market sell order with multiple existing buy limits" {
//    account = Acc()
//    account2 = Acc2()
//
// when:
//    ProcessTradeOrder(account, "BTC_EUR", "LIMIT", "BUY", 5, 10, "GTC")
//    ProcessTradeOrder(account, "BTC_EUR", "LIMIT", "BUY", 4, 10, "GTC")
//    ProcessTradeOrder(account, "BTC_EUR", "LIMIT", "BUY", 7, 10, "GTC")
//    ProcessTradeOrder(account2, "BTC_EUR", models.Market, "SELL", 30, "GTC")
//
// then:
//    assert.Equal(get_trade_count() == 3
//    assert.Equal(GetTradePrices() == [7, 5, 4]
//    assert.Equal(GetBuyBookOrderCount() == 0
//    assert.Equal(GetSellBookOrderCount() == 0
//    assert.Equal(GetVolumes("BTC_EUR", "BUY") == []
//    assert.Equal(GetVolumes("BTC_EUR", "SELL") == []
//  }
//
//  test "process/1 market buy order with multiple existing sell limits" {
//    account = Acc()
//    account2 = Acc2()
//
// when:
//    ProcessTradeOrder(account, "BTC_EUR", "LIMIT", "SELL", 5, 10, "GTC")
//    ProcessTradeOrder(account, "BTC_EUR", "LIMIT", "SELL", 4, 10, "GTC")
//    ProcessTradeOrder(account, "BTC_EUR", "LIMIT", "SELL", 7, 10, "GTC")
//    ProcessTradeOrder(account2, "BTC_EUR", models.Market, "BUY", 160, "GTC")
//
// then:
//    assert.Equal(get_trade_count() == 3
//    assert.Equal(GetTradePrices() == [4, 5, 7]
//    assert.Equal(GetBuyBookOrderCount() == 0
//    assert.Equal(GetSellBookOrderCount() == 0
//    assert.Equal(GetVolumes("BTC_EUR", "BUY") == []
//    assert.Equal(GetVolumes("BTC_EUR", "SELL") == []
//  }
//
//  test "process/1 limit buy order with existing market sell" {
//    account = Acc()
//    account2 = Acc2()
//
// when: the opposite book {es not have outstanding limit orders
//    ProcessTradeOrder(account2, "BTC_EUR", models.Market, "SELL", 10, "GTC")
//    ProcessTradeOrder(account, "BTC_EUR", "LIMIT", "BUY", 5, 10, "GTC")
//
// then: the trade settles at the incoming order’s limit price
//    assert.Equal(get_trade_count() == 1
//    assert.Equal(GetTradePrices() == [5]
//    assert.Equal(GetBuyBookOrderCount() == 0
//    assert.Equal(GetSellBookOrderCount() == 0
//    assert.Equal(GetVolumes("BTC_EUR", "BUY") == []
//    assert.Equal(GetVolumes("BTC_EUR", "SELL") == []
//  }
//
//  test "process/1 limit sell order with existing market buy" {
//    account = Acc()
//    account2 = Acc2()
//
// when: the opposite book {es not have outstanding limit orders
//    ProcessTradeOrder(account2, "BTC_EUR", models.Market, "BUY", 10, "GTC")
//    ProcessTradeOrder(account, "BTC_EUR", "LIMIT", "SELL", 5, 10, "GTC")
//
// then: the trade settles at the incoming order’s limit price
//    assert.Equal(get_trade_count() == 1
//    assert.Equal(GetTradePrices() == [5]
//    assert.Equal(GetBuyBookOrderCount() == 0
//    assert.Equal(GetSellBookOrderCount() == 1
//    assert.Equal(GetVolumes("BTC_EUR", "BUY") == []
//    assert.Equal(GetVolumes("BTC_EUR", "SELL") == [{5.0, 8.0}]
//  }
//
//  test "process/1 limit buy order with existing market sell and better limit sell" {
//    account = Acc()
//    account2 = Acc2()
//
// when: the opposite book has a market and a limit order
//    ProcessTradeOrder(account2, "BTC_EUR", "LIMIT", "SELL", 5, 10, "GTC")
//    ProcessTradeOrder(account2, "BTC_EUR", models.Market, "SELL", 10, "GTC")
//    ProcessTradeOrder(account, "BTC_EUR", "LIMIT", "BUY", 6, 20, "GTC")
//
// then: the trade settles at the book order’s limit price as if market order did not exist
//    assert.Equal(get_trade_count() == 2
//    assert.Equal(GetTradePrices() == [5, 5]
//    assert.Equal(GetBuyBookOrderCount() == 0
//    assert.Equal(GetSellBookOrderCount() == 0
//    assert.Equal(GetVolumes("BTC_EUR", "BUY") == []
//    assert.Equal(GetVolumes("BTC_EUR", "SELL") == []
//  }
//
//  test "process/1 limit sell order with existing market buy and better limit buy" {
//    account = Acc()
//    account2 = Acc2()
//
// when: the opposite book has a market and a limit order
//    ProcessTradeOrder(account2, "BTC_EUR", "LIMIT", "BUY", 5, 10, "GTC")
//    ProcessTradeOrder(account2, "BTC_EUR", models.Market, "BUY", 50, "GTC")
//    ProcessTradeOrder(account, "BTC_EUR", "LIMIT", "SELL", 4, 20, "GTC")
//
// then: the trade settles at the book order’s limit price as if market order did not exist
//    assert.Equal(get_trade_count() == 2
//    assert.Equal(GetTradePrices() == [5, 5]
//    assert.Equal(GetBuyBookOrderCount() == 0
//    assert.Equal(GetSellBookOrderCount() == 0
//    assert.Equal(GetVolumes("BTC_EUR", "BUY") == []
//    assert.Equal(GetVolumes("BTC_EUR", "SELL") == []
//  }
//
//  test "process/1 limit buy order with existing market sell and worse limit sell" {
//    account = Acc()
//    account2 = Acc2()
//
//    # when: the opposite book has a market and a limit order
//    ProcessTradeOrder(account2, "BTC_EUR", "LIMIT", "SELL", 7, 10, "GTC")
//    ProcessTradeOrder(account2, "BTC_EUR", models.Market, "SELL", 10, "GTC")
//    ProcessTradeOrder(account, "BTC_EUR", "LIMIT", "BUY", 6, 10, "GTC")
//
//    # then: the trade settles at the incoming order’s limit price
//    assert.Equal(get_trade_count() == 1
//    assert.Equal(GetTradePrices() == [6]
//    assert.Equal(GetBuyBookOrderCount() == 0
//    assert.Equal(GetSellBookOrderCount() == 1
//    assert.Equal(GetVolumes("BTC_EUR", "BUY") == []
//    assert.Equal(GetVolumes("BTC_EUR", "SELL") == [{7, 10}]
//  }
//
//  test "process/1 limit sell order with existing market buy and worse limit buy" {
//    account = Acc()
//    account2 = Acc2()
//
//    # when: the opposite book has a market and a limit order
//    ProcessTradeOrder(account2, "BTC_EUR", "LIMIT", "BUY", 4, 10, "GTC")
//    ProcessTradeOrder(account2, "BTC_EUR", models.Market, "BUY", 100, "GTC")
//    ProcessTradeOrder(account, "BTC_EUR", "LIMIT", "SELL", 5, 20, "GTC")
//
//    # then: the trade settles at the incoming order’s limit price
//    assert.Equal(get_trade_count() == 1
//    assert.Equal(GetTradePrices() == [5]
//    assert.Equal(GetBuyBookOrderCount() == 1
//    assert.Equal(GetSellBookOrderCount() == 0
//    assert.Equal(GetVolumes("BTC_EUR", "BUY") == [{4, 10}]
//    assert.Equal(GetVolumes("BTC_EUR", "SELL") == []
//  }
//
//  test "process/1 market buy order with existing market sell and no reference price" {
//    account = Acc()
//    account2 = Acc2()
//
//    # when: there is not a trade with a reference price
//    ProcessTradeOrder(account2, "BTC_EUR", models.Market, "SELL", 10, "GTC")
//    ProcessTradeOrder(account, "BTC_EUR", models.Market, "BUY", 50, "GTC")
//
//    # then: the trade settles at a reference price
//    assert.Equal(get_trade_count() == 0
//    assert.Equal(GetTradePrices() == []
//    assert.Equal(GetBuyBookOrderCount() == 1
//    assert.Equal(GetSellBookOrderCount() == 1
//    assert.Equal(GetVolumes("BTC_EUR", "BUY") == []
//    assert.Equal(GetVolumes("BTC_EUR", "SELL") == []
//  }
//
//  test "process/1 market sell order with existing market buy and no reference price" {
//    account = Acc()
//    account2 = Acc2()
//
//    # when: there is not a trade with a reference price
//    ProcessTradeOrder(account2, "BTC_EUR", models.Market, "BUY", 50, "GTC")
//    ProcessTradeOrder(account, "BTC_EUR", models.Market, "SELL", 10, "GTC")
//
//    # then: the trade settles at a reference price
//    assert.Equal(get_trade_count() == 0
//    assert.Equal(GetTradePrices() == []
//    assert.Equal(GetBuyBookOrderCount() == 1
//    assert.Equal(GetSellBookOrderCount() == 1
//    assert.Equal(GetVolumes("BTC_EUR", "BUY") == []
//    assert.Equal(GetVolumes("BTC_EUR", "SELL") == []
//  }
//
//  test "process/1 market buy order with existing market sell and a limit sell" {
//    account = Acc()
//    account2 = Acc2()
//
//    # when: the opposite book has a market and a limit order
//    ProcessTradeOrder(account2, "BTC_EUR", "LIMIT", "SELL", 5, 10, "GTC")
//    ProcessTradeOrder(account2, "BTC_EUR", models.Market, "SELL", 10, "GTC")
//    ProcessTradeOrder(account, "BTC_EUR", models.Market, "BUY", 50, "GTC")
//
//    # then: the trade settles at the book order’s limit price
//    assert.Equal(get_trade_count() == 1
//    assert.Equal(GetTradePrices() == [5]
//    assert.Equal(GetBuyBookOrderCount() == 0
//    assert.Equal(GetSellBookOrderCount() == 1
//    assert.Equal(GetVolumes("BTC_EUR", "BUY") == []
//    assert.Equal(GetVolumes("BTC_EUR", "SELL") == [{5, 10}]
//  }
//
//  test "process/1 market sell order with existing market buy and a limit buy" {
//    account = Acc()
//    account2 = Acc2()
//
//    # when: the opposite book has a market and a limit order
//    ProcessTradeOrder(account2, "BTC_EUR", "LIMIT", "BUY", 5, 10, "GTC")
//    ProcessTradeOrder(account2, "BTC_EUR", models.Market, "BUY", 50, "GTC")
//    ProcessTradeOrder(account, "BTC_EUR", models.Market, "SELL", 10, "GTC")
//
//    # then: the trade settles at the book order’s limit price
//    assert.Equal(get_trade_count() == 1
//    assert.Equal(GetTradePrices() == [5]
//    assert.Equal(GetBuyBookOrderCount() == 1
//    assert.Equal(GetSellBookOrderCount() == 0
//    assert.Equal(GetVolumes("BTC_EUR", "BUY") == [{5, 10}]
//    assert.Equal(GetVolumes("BTC_EUR", "SELL") == []
//  }
//
//  test "process/1 market buy order with existing market sell and a reference price" {
//    account = Acc()
//    account2 = Acc2()
//
//    # when: there is a trade with a reference price
//    ProcessTradeOrder(account2, "BTC_EUR", "LIMIT", "SELL", 5, 1, "GTC")
//    ProcessTradeOrder(account, "BTC_EUR", "LIMIT", "BUY", 5, 1, "GTC")
//    ProcessTradeOrder(account2, "BTC_EUR", models.Market, "SELL", 10, "GTC")
//    ProcessTradeOrder(account, "BTC_EUR", models.Market, "BUY", 50, "GTC")
//
//    # then: the trade settles at a reference price
//    assert.Equal(get_trade_count() == 2
//    assert.Equal(GetTradePrices() == [5, 5]
//    assert.Equal(GetBuyBookOrderCount() == 0
//    assert.Equal(GetSellBookOrderCount() == 0
//    assert.Equal(GetVolumes("BTC_EUR", "BUY") == []
//    assert.Equal(GetVolumes("BTC_EUR", "SELL") == []
//  }
//
//  test "process/1 market sell order with existing market buy and a reference price" {
//    account = Acc()
//    account2 = Acc2()
//
//    # when: the opposite book has a market and a limit order
//    ProcessTradeOrder(account2, "BTC_EUR", "LIMIT", "SELL", 5, 1, "GTC")
//    ProcessTradeOrder(account, "BTC_EUR", "LIMIT", "BUY", 5, 1, "GTC")
//    ProcessTradeOrder(account2, "BTC_EUR", models.Market, "BUY", 50, "GTC")
//    ProcessTradeOrder(account, "BTC_EUR", models.Market, "SELL", 10, "GTC")
//
//    # then: the trade settles at a reference price
//    assert.Equal(get_trade_count() == 2
//    assert.Equal(GetTradePrices() == [5, 5]
//    assert.Equal(GetBuyBookOrderCount() == 0
//    assert.Equal(GetSellBookOrderCount() == 0
//    assert.Equal(GetVolumes("BTC_EUR", "BUY") == []
//    assert.Equal(GetVolumes("BTC_EUR", "SELL") == []
//  }
//
//  test "process/1 limit buy order with existing market sell and a reference price" {
//    account = Acc()
//    account2 = Acc2()
//
//    # when: there is a trade with a reference price
//    ProcessTradeOrder(account2, "BTC_EUR", "LIMIT", "SELL", 5, 1, "GTC")
//    ProcessTradeOrder(account, "BTC_EUR", "LIMIT", "BUY", 5, 1, "GTC")
//    ProcessTradeOrder(account2, "BTC_EUR", models.Market, "SELL", 10, "GTC")
//    ProcessTradeOrder(account, "BTC_EUR", "LIMIT", "BUY", 6, 10, "GTC")
//
//    # then: the trade settles at a reference price
//    assert.Equal(get_trade_count() == 2
//    assert.Equal(GetTradePrices() == [5, 5]
//    assert.Equal(GetBuyBookOrderCount() == 0
//    assert.Equal(GetSellBookOrderCount() == 0
//    assert.Equal(GetVolumes("BTC_EUR", "BUY") == []
//    assert.Equal(GetVolumes("BTC_EUR", "SELL") == []
//  }
//
//  # ---- TESTS BELOW REQUIRE VARIOUS HACKS TO SIMULATE TIMESTAMP INCREMENTS
//  test "process/1 market buy order with existing market sell and multiple reference prices" {
//    account = Acc()
//    account2 = Acc2()
//
//    # when: there is a trade with a reference price
//    ProcessTradeOrder(account2, "BTC_EUR", "LIMIT", "SELL", 4, 1, "GTC")
//    ProcessTradeOrder(account, "BTC_EUR", "LIMIT", "BUY", 4, 1, "GTC")
//
//    DB.execute(
//      "UPDATE trade SET created_at = current_timestamp + interval '1 second' WHERE price = 4"
//    )
//
//    ProcessTradeOrder(account2, "BTC_EUR", "LIMIT", "SELL", 6, 1, "GTC")
//    ProcessTradeOrder(account, "BTC_EUR", "LIMIT", "BUY", 6, 1, "GTC")
//
//    DB.execute(
//      "UPDATE trade SET created_at = current_timestamp + interval '2 second' WHERE price = 6"
//    )
//
//    ProcessTradeOrder(account2, "BTC_EUR", "LIMIT", "SELL", 5, 1, "GTC")
//    ProcessTradeOrder(account, "BTC_EUR", "LIMIT", "BUY", 5, 1, "GTC")
//
//    DB.execute(
//      "UPDATE trade SET created_at = current_timestamp + interval '3 second' WHERE price = 5"
//    )
//
//    ProcessTradeOrder(account2, "BTC_EUR", models.Market, "SELL", 10, "GTC")
//    ProcessTradeOrder(account, "BTC_EUR", models.Market, "BUY", 50, "GTC")
//
//    # then: the trade settles at last reference price
//    assert.Equal(get_trade_count() == 4
//    assert.Equal(GetTradePrices() |> List.first() == 5
//    assert.Equal(GetBuyBookOrderCount() == 0
//    assert.Equal(GetSellBookOrderCount() == 0
//    assert.Equal(GetVolumes("BTC_EUR", "BUY") == []
//    assert.Equal(GetVolumes("BTC_EUR", "SELL") == []
//  }
//
//  test "process/1 market sell order with existing market buy and multiple reference prices" {
//    account = Acc()
//    account2 = Acc2()
//
//    # when: the opposite book has a market and a limit order
//    ProcessTradeOrder(account2, "BTC_EUR", "LIMIT", "SELL", 4, 1, "GTC")
//    ProcessTradeOrder(account, "BTC_EUR", "LIMIT", "BUY", 4, 1, "GTC")
//
//    DB.execute(
//      "UPDATE trade SET created_at = current_timestamp + interval '1 second' WHERE price = 4"
//    )
//
//    ProcessTradeOrder(account2, "BTC_EUR", "LIMIT", "SELL", 6, 1, "GTC")
//    ProcessTradeOrder(account, "BTC_EUR", "LIMIT", "BUY", 6, 1, "GTC")
//
//    DB.execute(
//      "UPDATE trade SET created_at = current_timestamp + interval '2 second' WHERE price = 6"
//    )
//
//    ProcessTradeOrder(account2, "BTC_EUR", "LIMIT", "SELL", 5, 1, "GTC")
//    ProcessTradeOrder(account, "BTC_EUR", "LIMIT", "BUY", 5, 1, "GTC")
//
//    DB.execute(
//      "UPDATE trade SET created_at = current_timestamp + interval '3 second' WHERE price = 5"
//    )
//
//    ProcessTradeOrder(account2, "BTC_EUR", models.Market, "BUY", 50, "GTC")
//    ProcessTradeOrder(account, "BTC_EUR", models.Market, "SELL", 10, "GTC")
//
//    # then: the trade settles at last reference price
//    assert.Equal(get_trade_count() == 4
//    assert.Equal(GetTradePrices() |> List.first() == 5
//    assert.Equal(GetBuyBookOrderCount() == 0
//    assert.Equal(GetSellBookOrderCount() == 0
//    assert.Equal(GetVolumes("BTC_EUR", "BUY") == []
//    assert.Equal(GetVolumes("BTC_EUR", "SELL") == []
//  }
//}
