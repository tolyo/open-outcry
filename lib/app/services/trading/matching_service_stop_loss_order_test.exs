defmodule MatchingServiceStopLossOrderTest do
  use DataCase

  import TestUtils

  test "create/1 stop loss sell order save" do
    # given
    account = acc()
    entity = TestUtils.get_application_entity_id()

    # when: a stop loss order is created
    res = MatchingService.create(account, "BTC_EUR", :STOPLOSS, :SELL, 10, 100, :GTC)

    # then:
    assert res != nil
    assert DBTestUtils.get_count("stop_order") == 1
    assert MatchingServiceTestHelpers.get_sell_book_order_count() == 0
    assert OrderBookService.get_volumes("BTC_EUR", :SELL) == []

    assert PaymentAccount.find_by_application_entity_and_currency(entity, "BTC").amount_reserved
           |> Decimal.to_float() == 100
  end

  test "create/1 stop loss buy order buy" do
    # given
    account = acc()
    entity = TestUtils.get_application_entity_id()

    # when: a stop loss order is created
    res = MatchingService.create(account, "BTC_EUR", :STOPLOSS, :BUY, 10, 100, :GTC)

    # then:
    assert res != nil
    assert DBTestUtils.get_count("stop_order") == 1
    assert MatchingServiceTestHelpers.get_buy_book_order_count() == 0
    assert OrderBookService.get_volumes("BTC_EUR", :BUY) == []

    assert PaymentAccount.find_by_application_entity_and_currency(entity, "EUR").amount_reserved
           |> Decimal.to_float() == 1000
  end

  test "create/1 stop loss sell order activate" do
    # given
    account = acc()
    account2 = acc2()

    # when: a stop loss order is created and then a crossing trade occurs
    MatchingService.create(account, "BTC_EUR", :STOPLOSS, :SELL, 10, 2, :GTC)
    MatchingService.create(account, "BTC_EUR", :LIMIT, :SELL, 10, 1, :GTC)
    MatchingService.create(account2, "BTC_EUR", :LIMIT, :BUY, 10, 1, :GTC)

    # then: the order becomes activated as a market order which is invvisible to the order book
    assert DBTestUtils.get_count("stop_order") == 0
    assert MatchingServiceTestHelpers.get_sell_book_order_count() == 1
    assert OrderBookService.get_volumes("BTC_EUR", :SELL) == []
  end

  test "create/1 stop loss buy order activate" do
    # given
    account = acc()
    account2 = acc2()

    # when: a stop loss order is created and then a crossing trade occurs
    MatchingService.create(account, "BTC_EUR", :STOPLOSS, :BUY, 10, 2, :GTC)
    MatchingService.create(account2, "BTC_EUR", :LIMIT, :BUY, 10, 1, :GTC)
    MatchingService.create(account, "BTC_EUR", :LIMIT, :SELL, 10, 1, :GTC)

    # then: the order becomes activated as a market order which is invvisible to the order book
    assert DBTestUtils.get_count("stop_order") == 0
    assert MatchingServiceTestHelpers.get_buy_book_order_count() == 1
    assert OrderBookService.get_volumes("BTC_EUR", :BUY) == []
  end

  test "create/1 stop loss sell order non crossing " do
    # given
    account = acc()
    account2 = acc2()

    # when: a stop loss order is created and then a non-crossing trade occurs
    MatchingService.create(account, "BTC_EUR", :STOPLOSS, :SELL, 10, 2, :GTC)
    MatchingService.create(account, "BTC_EUR", :LIMIT, :SELL, 11, 1, :GTC)
    MatchingService.create(account2, "BTC_EUR", :LIMIT, :BUY, 11, 1, :GTC)

    # then: the order remains deactivated
    assert DBTestUtils.get_count("stop_order") == 1
    assert MatchingServiceTestHelpers.get_sell_book_order_count() == 0
    assert OrderBookService.get_volumes("BTC_EUR", :SELL) == []

    # when: a crossing trade occurs
    MatchingService.create(account, "BTC_EUR", :LIMIT, :SELL, 9, 1, :GTC)
    MatchingService.create(account2, "BTC_EUR", :LIMIT, :BUY, 9, 1, :GTC)

    # then: it becomes activated
    assert DBTestUtils.get_count("stop_order") == 0
    assert MatchingServiceTestHelpers.get_sell_book_order_count() == 1
    assert OrderBookService.get_volumes("BTC_EUR", :SELL) == []
  end

  test "create/1 stop loss buy order non crossing " do
    # given
    account = acc()
    account2 = acc2()

    # when: a stop loss order is created and then a non-crossing trade occurs
    MatchingService.create(account, "BTC_EUR", :STOPLOSS, :BUY, 10, 2, :GTC)
    MatchingService.create(account2, "BTC_EUR", :LIMIT, :BUY, 11, 1, :GTC)
    MatchingService.create(account, "BTC_EUR", :LIMIT, :SELL, 11, 1, :GTC)

    # then: the order remains deactivated
    assert DBTestUtils.get_count("stop_order") == 1
    assert MatchingServiceTestHelpers.get_sell_book_order_count() == 0
    assert OrderBookService.get_volumes("BTC_EUR", :SELL) == []

    # when: a crossing trade occurs
    MatchingService.create(account2, "BTC_EUR", :LIMIT, :BUY, 9, 1, :GTC)
    MatchingService.create(account, "BTC_EUR", :LIMIT, :SELL, 9, 1, :GTC)

    # then: it becomes activated
    assert DBTestUtils.get_count("stop_order") == 0
    assert MatchingServiceTestHelpers.get_buy_book_order_count() == 1
    assert OrderBookService.get_volumes("BTC_EUR", :BUY) == []
  end

  test "create/1 stop loss sell order activate and settle" do
    # given
    account = acc()
    account2 = acc2()

    # when: a stop loss order is created and then a crossing trade occurs
    MatchingService.create(account, "BTC_EUR", :STOPLOSS, :SELL, 10, 2, :GTC)
    MatchingService.create(account, "BTC_EUR", :LIMIT, :SELL, 10, 1, :GTC)
    MatchingService.create(account2, "BTC_EUR", :LIMIT, :BUY, 10, 3, :GTC)

    # then: the order becomes activated and settled
    assert DBTestUtils.get_count("stop_order") == 0
    assert MatchingServiceTestHelpers.get_sell_book_order_count() == 0
    assert OrderBookService.get_volumes("BTC_EUR", :BUY) == []
    assert OrderBookService.get_volumes("BTC_EUR", :SELL) == []
  end

  test "create/1 stop loss buy order activate and settle" do
    # given
    account = acc()
    account2 = acc2()

    # when: a stop loss order is created and then a crossing trade occurs
    MatchingService.create(account, "BTC_EUR", :STOPLOSS, :BUY, 10, 20, :GTC)
    MatchingService.create(account, "BTC_EUR", :LIMIT, :BUY, 10, 1, :GTC)
    MatchingService.create(account2, "BTC_EUR", :LIMIT, :SELL, 10, 3, :GTC)

    # then: the order becomes activated and settled
    assert DBTestUtils.get_count("stop_order") == 0
    assert MatchingServiceTestHelpers.get_sell_book_order_count() == 0
    assert OrderBookService.get_volumes("BTC_EUR", :BUY) == []
    assert OrderBookService.get_volumes("BTC_EUR", :SELL) == []
  end

  # test "create/1 stop loss sell order activate and settle opposite side" do
  #   # given
  #   account = acc()
  #   account2 = acc2()

  #   # when: a stop loss order is created and then a crossing trade occurs
  #   MatchingService.create(account, "BTC_EUR", :STOPLOSS, :SELL, 10, 2, :GTC)
  #   MatchingService.create(account2, "BTC_EUR", :LIMIT, :BUY, 10, 3, :GTC)
  #   MatchingService.create(account, "BTC_EUR", :LIMIT, :SELL, 10, 1, :GTC)

  #   # then: the order becomes activated and settled
  #   assert DBTestUtils.get_count("stop_order") == 0
  #   assert MatchingServiceTestHelpers.get_sell_book_order_count() == 0
  #   assert OrderBookService.get_volumes("BTC_EUR", :BUY) == []
  #   assert OrderBookService.get_volumes("BTC_EUR", :SELL) == []
  # end

  # test "create/1 stop loss buy order activate and settle opposite side" do
  #   # given
  #   account = acc()
  #   account2 = acc2()

  #   # when: a stop loss order is created and then a crossing trade occurs
  #   MatchingService.create(account, "BTC_EUR", :STOPLOSS, :BUY, 10, 20, :GTC)
  #   MatchingService.create(account2, "BTC_EUR", :LIMIT, :SELL, 10, 3, :GTC)
  #   MatchingService.create(account, "BTC_EUR", :LIMIT, :BUY, 10, 1, :GTC)

  #   # then: the order becomes activated and settled
  #   assert DBTestUtils.get_count("stop_order") == 0
  #   assert MatchingServiceTestHelpers.get_sell_book_order_count() == 0
  #   assert OrderBookService.get_volumes("BTC_EUR", :BUY) == []
  #   assert OrderBookService.get_volumes("BTC_EUR", :SELL) == []
  # end

  test "create/1 stop loss sell order activate and settle before worse price orders" do
    # given
    account = acc()
    account2 = acc2()

    # when: a stop loss order is created and then a crossing trade occurs
    MatchingService.create(account, "BTC_EUR", :STOPLOSS, :SELL, 10, 1, :GTC)
    MatchingService.create(account, "BTC_EUR", :LIMIT, :SELL, 11, 1, :GTC)
    MatchingService.create(account, "BTC_EUR", :LIMIT, :SELL, 10, 1, :GTC)
    MatchingService.create(account2, "BTC_EUR", :LIMIT, :BUY, 11, 3, :GTC)

    # then: the order becomes activated and settled at last trade price which is 10
    assert MatchingServiceTestHelpers.get_trade_prices() == [10, 10, 11]
    assert DBTestUtils.get_count("stop_order") == 0
    assert MatchingServiceTestHelpers.get_sell_book_order_count() == 0
    assert OrderBookService.get_volumes("BTC_EUR", :BUY) == []
    assert OrderBookService.get_volumes("BTC_EUR", :SELL) == []
  end

  test "create/1 stop loss buy order activate and settle before worse price orders" do
    # given
    account = acc()
    account2 = acc2()

    # when: a stop loss order is created and then a crossing trade occurs
    MatchingService.create(account, "BTC_EUR", :STOPLOSS, :BUY, 10, 10, :GTC)
    MatchingService.create(account, "BTC_EUR", :LIMIT, :BUY, 9, 1, :GTC)
    MatchingService.create(account, "BTC_EUR", :LIMIT, :BUY, 10, 1, :GTC)
    MatchingService.create(account2, "BTC_EUR", :LIMIT, :SELL, 9, 3, :GTC)

    # then: the order becomes activated and settled
    assert MatchingServiceTestHelpers.get_trade_prices() == [10, 10, 9]
    # TODO fix rounding
    assert DBTestUtils.get_count("stop_order") == 0
    assert MatchingServiceTestHelpers.get_sell_book_order_count() == 0
    assert OrderBookService.get_volumes("BTC_EUR", :BUY) == []
    assert OrderBookService.get_volumes("BTC_EUR", :SELL) == []
  end

  # test "create/1 stop loss sell order activate by market and settle" do
  #   # given
  #   account = acc()
  #   account2 = acc2()

  #   # when: a stop loss order is created and then a crossing trade occurs
  #   MatchingService.create(account, "BTC_EUR", :STOPLOSS, :SELL, 10, 2, :GTC)
  #   MatchingService.create(account, "BTC_EUR", :MARKET, :SELL, 1, :GTC)
  #   MatchingService.create(account2, "BTC_EUR", :LIMIT, :BUY, 10, 3, :GTC)

  #   # then: the order becomes activated and settled
  #   assert DBTestUtils.get_count("stop_order") == 0
  #   assert MatchingServiceTestHelpers.get_sell_book_order_count() == 0
  #   assert OrderBookService.get_volumes("BTC_EUR", :BUY) == []
  #   assert OrderBookService.get_volumes("BTC_EUR", :SELL) == []

  #   # when: a stop loss order is created and then a crossing trade occurs
  #   MatchingService.create(account, "BTC_EUR", :STOPLOSS, :SELL, 10, 2, :GTC)
  #   MatchingService.create(account2, "BTC_EUR", :LIMIT, :BUY, 10, 3, :GTC)
  #   MatchingService.create(account, "BTC_EUR", :MARKET, :SELL, 1, :GTC)

  #   # then: the order becomes activated and settled
  #   assert DBTestUtils.get_count("stop_order") == 0
  #   assert MatchingServiceTestHelpers.get_sell_book_order_count() == 0
  #   assert OrderBookService.get_volumes("BTC_EUR", :BUY) == []
  #   assert OrderBookService.get_volumes("BTC_EUR", :SELL) == []
  # end

  # test "create/1 stop loss buy order activate by market and settle" do
  #   # given
  #   account = acc()
  #   account2 = acc2()

  #   # when: a stop loss order is created and then a crossing trade occurs
  #   MatchingService.create(account, "BTC_EUR", :STOPLOSS, :BUY, 10, 20, :GTC)
  #   MatchingService.create(account, "BTC_EUR", :MARKET, :BUY, 10, :GTC)
  #   MatchingService.create(account2, "BTC_EUR", :LIMIT, :SELL, 10, 3, :GTC)

  #   # then: the order becomes activated and settled
  #   assert DBTestUtils.get_count("stop_order") == 0
  #   assert MatchingServiceTestHelpers.get_sell_book_order_count() == 0
  #   assert OrderBookService.get_volumes("BTC_EUR", :BUY) == []
  #   assert OrderBookService.get_volumes("BTC_EUR", :SELL) == []

  #   # when: a stop loss order is created and then a crossing trade occurs
  #   MatchingService.create(account, "BTC_EUR", :STOPLOSS, :BUY, 10, 20, :GTC)
  #   MatchingService.create(account2, "BTC_EUR", :LIMIT, :SELL, 10, 3, :GTC)
  #   MatchingService.create(account, "BTC_EUR", :MARKET, :BUY, 10, :GTC)

  #   # then: the order becomes activated and settled
  #   assert DBTestUtils.get_count("stop_order") == 0

  #   assert MatchingServiceTestHelpers.get_sell_book_order_count() == 0
  #   assert OrderBookService.get_volumes("BTC_EUR", :BUY) == []
  #   assert OrderBookService.get_volumes("BTC_EUR", :SELL) == []
  # end
end
