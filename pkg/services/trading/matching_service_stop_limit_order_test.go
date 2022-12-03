funcmodule MatchingServiceStopLimitOrderTest {
  use DataCase

  import TestUtils

  setup {
    [
      trading_account: acc(),
      entity: get_application_entity_id(),
      trading_account2: acc2()
    ]
  }

  test "create/1 stop limit sell order save", c {
    # when: a stop limit order is created
    res = MatchingService.create(c.trading_account, "BTC_EUR", :STOPLIMIT, :SELL, 10, 100, :GTC)

    # then:
    assert res != nil
    assert DBTestUtils.get_count("stop_order") == 1
    assert MatchingServiceTestHelpers.get_sell_book_order_count() == 0
    assert OrderBookService.get_volumes("BTC_EUR", :SELL) == []

    assert PaymentAccount.find_by_application_entity_and_currency(c.entity, "BTC").amount_reserved
           |> Decimal.to_float() == 100
  }

  test "create/1 stop limit buy order buy", c {
    # when: a stop limit order is created
    res = MatchingService.create(c.trading_account, "BTC_EUR", :STOPLIMIT, :BUY, 10, 100, :GTC)

    # then:
    assert res != nil
    assert DBTestUtils.get_count("stop_order") == 1
    assert MatchingServiceTestHelpers.get_buy_book_order_count() == 0
    assert OrderBookService.get_volumes("BTC_EUR", :BUY) == []

    assert PaymentAccount.find_by_application_entity_and_currency(c.entity, "EUR").amount_reserved
           |> Decimal.to_float() == 1000
  }

  test "create/1 stop limit sell order activate", c {
    # when: a stop limit order is created and then a crossing trade occurs
    MatchingService.create(c.trading_account, "BTC_EUR", :STOPLIMIT, :SELL, 10, 2, :GTC)
    MatchingService.create(c.trading_account, "BTC_EUR", :LIMIT, :SELL, 10, 1, :GTC)
    MatchingService.create(c.trading_account2, "BTC_EUR", :LIMIT, :BUY, 10, 1, :GTC)

    # then: the order becomes activated
    assert DBTestUtils.get_count("stop_order") == 0
    assert MatchingServiceTestHelpers.get_sell_book_order_count() == 1
    assert OrderBookService.get_volumes("BTC_EUR", :SELL) == [{10, 2}]
  }

  test "create/1 stop limit buy order activate", c {
    # when: a stop limit order is created and then a crossing trade occurs
    MatchingService.create(c.trading_account, "BTC_EUR", :STOPLIMIT, :BUY, 10, 2, :GTC)
    MatchingService.create(c.trading_account2, "BTC_EUR", :LIMIT, :BUY, 10, 1, :GTC)
    MatchingService.create(c.trading_account, "BTC_EUR", :LIMIT, :SELL, 10, 1, :GTC)

    # then: the order becomes activated
    assert DBTestUtils.get_count("stop_order") == 0
    assert MatchingServiceTestHelpers.get_buy_book_order_count() == 1
    assert OrderBookService.get_volumes("BTC_EUR", :BUY) == [{10, 2}]
  }

  test "create/1 stop limit sell order non crossing ", c {
    # when: a stop limit order is created and then a non-crossing trade occurs
    MatchingService.create(c.trading_account, "BTC_EUR", :STOPLIMIT, :SELL, 10, 2, :GTC)
    MatchingService.create(c.trading_account, "BTC_EUR", :LIMIT, :SELL, 11, 1, :GTC)
    MatchingService.create(c.trading_account2, "BTC_EUR", :LIMIT, :BUY, 11, 1, :GTC)

    # then: the order remains deactivated
    assert DBTestUtils.get_count("stop_order") == 1
    assert MatchingServiceTestHelpers.get_sell_book_order_count() == 0
    assert OrderBookService.get_volumes("BTC_EUR", :SELL) == []

    # when: a crossing trade occurs
    MatchingService.create(c.trading_account, "BTC_EUR", :LIMIT, :SELL, 9, 1, :GTC)
    MatchingService.create(c.trading_account2, "BTC_EUR", :LIMIT, :BUY, 9, 1, :GTC)

    # then: it becomes activated
    assert DBTestUtils.get_count("stop_order") == 0
    assert MatchingServiceTestHelpers.get_sell_book_order_count() == 1
    assert OrderBookService.get_volumes("BTC_EUR", :SELL) == [{10, 2}]
  }

  test "create/1 stop limit buy order non crossing ", c {
    # when: a stop limit order is created and then a non-crossing trade occurs
    MatchingService.create(c.trading_account, "BTC_EUR", :STOPLIMIT, :BUY, 10, 2, :GTC)
    MatchingService.create(c.trading_account2, "BTC_EUR", :LIMIT, :BUY, 11, 1, :GTC)
    MatchingService.create(c.trading_account, "BTC_EUR", :LIMIT, :SELL, 11, 1, :GTC)

    # then: the order remains deactivated
    assert DBTestUtils.get_count("stop_order") == 1
    assert MatchingServiceTestHelpers.get_sell_book_order_count() == 0
    assert OrderBookService.get_volumes("BTC_EUR", :SELL) == []

    # when: a crossing trade occurs
    MatchingService.create(c.trading_account2, "BTC_EUR", :LIMIT, :BUY, 9, 1, :GTC)
    MatchingService.create(c.trading_account, "BTC_EUR", :LIMIT, :SELL, 9, 1, :GTC)

    # then: it becomes activated
    assert DBTestUtils.get_count("stop_order") == 0
    assert MatchingServiceTestHelpers.get_buy_book_order_count() == 1
    assert OrderBookService.get_volumes("BTC_EUR", :BUY) == [{10, 2}]
  }

  test "create/1 stop limit sell order activate and settle", c {
    # when: a stop limit order is created and then a crossing trade occurs
    MatchingService.create(c.trading_account, "BTC_EUR", :STOPLIMIT, :SELL, 10, 2, :GTC)
    MatchingService.create(c.trading_account, "BTC_EUR", :LIMIT, :SELL, 10, 1, :GTC)
    MatchingService.create(c.trading_account2, "BTC_EUR", :LIMIT, :BUY, 10, 3, :GTC)

    # then: the order becomes activated and settled
    assert DBTestUtils.get_count("stop_order") == 0
    assert MatchingServiceTestHelpers.get_sell_book_order_count() == 0
    assert OrderBookService.get_volumes("BTC_EUR", :BUY) == []
    assert OrderBookService.get_volumes("BTC_EUR", :SELL) == []
  }

  test "create/1 stop limit buy order activate and settle", c {
    # when: a stop limit order is created and then a crossing trade occurs
    MatchingService.create(c.trading_account, "BTC_EUR", :STOPLIMIT, :BUY, 10, 2, :GTC)
    MatchingService.create(c.trading_account, "BTC_EUR", :LIMIT, :BUY, 10, 1, :GTC)
    MatchingService.create(c.trading_account2, "BTC_EUR", :LIMIT, :SELL, 10, 3, :GTC)

    # then: the order becomes activated and settled
    assert DBTestUtils.get_count("stop_order") == 0
    assert MatchingServiceTestHelpers.get_sell_book_order_count() == 0
    assert OrderBookService.get_volumes("BTC_EUR", :BUY) == []
    assert OrderBookService.get_volumes("BTC_EUR", :SELL) == []
  }

  test "create/1 stop limit sell order activate and settle opposite side", c {
    # when: a stop limit order is created and then a crossing trade occurs
    MatchingService.create(c.trading_account, "BTC_EUR", :STOPLIMIT, :SELL, 10, 2, :GTC)
    MatchingService.create(c.trading_account2, "BTC_EUR", :LIMIT, :BUY, 10, 3, :GTC)
    MatchingService.create(c.trading_account, "BTC_EUR", :LIMIT, :SELL, 10, 1, :GTC)

    # then: the order becomes activated and settled
    assert DBTestUtils.get_count("stop_order") == 0
    assert MatchingServiceTestHelpers.get_sell_book_order_count() == 0
    assert OrderBookService.get_volumes("BTC_EUR", :BUY) == []
    assert OrderBookService.get_volumes("BTC_EUR", :SELL) == []
  }

  test "create/1 stop limit buy order activate and settle opposite side", c {
    # when: a stop limit order is created and then a crossing trade occurs
    MatchingService.create(c.trading_account, "BTC_EUR", :STOPLIMIT, :BUY, 10, 2, :GTC)
    MatchingService.create(c.trading_account2, "BTC_EUR", :LIMIT, :SELL, 10, 3, :GTC)
    MatchingService.create(c.trading_account, "BTC_EUR", :LIMIT, :BUY, 10, 1, :GTC)

    # then: the order becomes activated and settled
    assert DBTestUtils.get_count("stop_order") == 0
    assert MatchingServiceTestHelpers.get_sell_book_order_count() == 0
    assert OrderBookService.get_volumes("BTC_EUR", :BUY) == []
    assert OrderBookService.get_volumes("BTC_EUR", :SELL) == []
  }

  test "create/1 stop limit sell order activate and settle before worse price orders", c {
    # when: a stop limit order is created and then a crossing trade occurs
    MatchingService.create(c.trading_account, "BTC_EUR", :STOPLIMIT, :SELL, 10, 1, :GTC)
    MatchingService.create(c.trading_account, "BTC_EUR", :LIMIT, :SELL, 11, 1, :GTC)
    MatchingService.create(c.trading_account, "BTC_EUR", :LIMIT, :SELL, 10, 1, :GTC)
    MatchingService.create(c.trading_account2, "BTC_EUR", :LIMIT, :BUY, 11, 3, :GTC)

    # then: the order becomes activated and settled
    assert MatchingServiceTestHelpers.get_trade_prices() == [10, 10, 11]
    assert DBTestUtils.get_count("stop_order") == 0
    assert MatchingServiceTestHelpers.get_sell_book_order_count() == 0
    assert OrderBookService.get_volumes("BTC_EUR", :BUY) == []
    assert OrderBookService.get_volumes("BTC_EUR", :SELL) == []
  }

  test "create/1 stop limit buy order activate and settle before worse price orders", c {
    # when: a stop limit order is created and then a crossing trade occurs
    MatchingService.create(c.trading_account, "BTC_EUR", :STOPLIMIT, :BUY, 10, 1, :GTC)
    MatchingService.create(c.trading_account, "BTC_EUR", :LIMIT, :BUY, 9, 1, :GTC)
    MatchingService.create(c.trading_account, "BTC_EUR", :LIMIT, :BUY, 10, 1, :GTC)
    MatchingService.create(c.trading_account2, "BTC_EUR", :LIMIT, :SELL, 9, 3, :GTC)

    # then: the order becomes activated and settled
    assert MatchingServiceTestHelpers.get_trade_prices() == [10, 10, 9]
    assert DBTestUtils.get_count("stop_order") == 0
    assert MatchingServiceTestHelpers.get_sell_book_order_count() == 0
    assert OrderBookService.get_volumes("BTC_EUR", :BUY) == []
    assert OrderBookService.get_volumes("BTC_EUR", :SELL) == []
  }

  test "create/1 stop limit sell order activate by market and settle", c {
    # when: a stop limit order is created and then a crossing trade occurs
    MatchingService.create(c.trading_account, "BTC_EUR", :STOPLIMIT, :SELL, 10, 2, :GTC)
    MatchingService.create(c.trading_account, "BTC_EUR", :MARKET, :SELL, 1, :GTC)
    MatchingService.create(c.trading_account2, "BTC_EUR", :LIMIT, :BUY, 10, 3, :GTC)

    # then: the order becomes activated and settled
    assert DBTestUtils.get_count("stop_order") == 0
    assert MatchingServiceTestHelpers.get_sell_book_order_count() == 0
    assert OrderBookService.get_volumes("BTC_EUR", :BUY) == []
    assert OrderBookService.get_volumes("BTC_EUR", :SELL) == []

    # when: a stop limit order is created and then a crossing trade occurs
    MatchingService.create(c.trading_account, "BTC_EUR", :STOPLIMIT, :SELL, 10, 2, :GTC)
    MatchingService.create(c.trading_account2, "BTC_EUR", :LIMIT, :BUY, 10, 3, :GTC)
    MatchingService.create(c.trading_account, "BTC_EUR", :MARKET, :SELL, 1, :GTC)

    # then: the order becomes activated and settled
    assert DBTestUtils.get_count("stop_order") == 0
    assert MatchingServiceTestHelpers.get_sell_book_order_count() == 0
    assert OrderBookService.get_volumes("BTC_EUR", :BUY) == []
    assert OrderBookService.get_volumes("BTC_EUR", :SELL) == []
  }

  test "create/1 stop limit buy order activate by market and settle", c {
    # when: a stop limit order is created and then a crossing trade occurs
    MatchingService.create(c.trading_account, "BTC_EUR", :STOPLIMIT, :BUY, 10, 2, :GTC)
    MatchingService.create(c.trading_account, "BTC_EUR", :MARKET, :BUY, 10, :GTC)
    MatchingService.create(c.trading_account2, "BTC_EUR", :LIMIT, :SELL, 10, 3, :GTC)

    # then: the order becomes activated and settled
    assert DBTestUtils.get_count("stop_order") == 0
    assert MatchingServiceTestHelpers.get_sell_book_order_count() == 0
    assert OrderBookService.get_volumes("BTC_EUR", :BUY) == []
    assert OrderBookService.get_volumes("BTC_EUR", :SELL) == []

    # when: a stop limit order is created and then a crossing trade occurs
    MatchingService.create(c.trading_account, "BTC_EUR", :STOPLIMIT, :BUY, 10, 2, :GTC)
    MatchingService.create(c.trading_account2, "BTC_EUR", :LIMIT, :SELL, 10, 3, :GTC)
    MatchingService.create(c.trading_account, "BTC_EUR", :MARKET, :BUY, 10, :GTC)

    # then: the order becomes activated and settled
    assert DBTestUtils.get_count("stop_order") == 0

    assert MatchingServiceTestHelpers.get_sell_book_order_count() == 0
    assert OrderBookService.get_volumes("BTC_EUR", :BUY) == []
    assert OrderBookService.get_volumes("BTC_EUR", :SELL) == []
  }
}
