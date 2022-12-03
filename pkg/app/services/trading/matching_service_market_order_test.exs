defmodule MatchingServiceMarketOrderTest do
  use DataCase

  import TestUtils
  use ExUnit.Case, async: false

  @moduledoc """

  Primary market order test cases:
    * An incoming market order hits an outstanding limit order in the opposite book
    * An incoming limit order hits an outstanding market order in the opposite book
    * An incoming market order hits another market order of the opposite side

  Additional use cases to test:
    * If the opposite book does not have outstanding limit orders,
      then the trade settles at the incoming order’s limit price
    * If the opposite book does have limit orders, the trade settles at the better of two prices –
      the incoming order’s limit or the best limit from the opposite book –
      the term “better of two prices” is from the point of view of the incoming limit order.
      In other words, if incoming limit order would have crossed with outstanding opposite
      “best limit” order in the absence of market order,
      then trade at that, potentially improved, “best limit” price.

    * when a market order matches with another market order in the opposite order book, there are three possibilites to test:
      - the opposite order book with the resting market order also contains one or more outstanding limit orders –
        in this case the opposite book has a “best limit” price and this price becomes the price for the trade
      - the opposte order book does not have outstanding limit orders, so the “best limit” price is not defined.
        In this case the trade occurs at the “reference price”. Most often reference price is the last traded price for a security.
      - if no trades have occured (upon security launch), market orders simply rest
  """

  test "process/1 market sell order save" do
    # when: a market order is sent to an empty matching unit
    MatchingService.create(acc(), "BTC_EUR", :MARKET, :SELL, 100, :GTC)

    # then: a matching unit should save the trade order but it should not be visible to the order book
    assert MatchingServiceTestHelpers.get_sell_book_order_count() == 1
    assert OrderBookService.get_volumes("BTC_EUR", :SELL) == []
  end

  test "process/1 market buy order save" do
    # when: a market order is sent to an empty matching unit
    MatchingService.create(acc(), "BTC_EUR", :MARKET, :BUY, 100, :GTC)

    # then: a matching unit should save the trade order but it should not be visible to the order book
    assert MatchingServiceTestHelpers.get_buy_book_order_count() == 1
    assert OrderBookService.get_volumes("BTC_EUR", :BUY) == []
  end

  test "process/1 market sell order with existing sell limit" do
    account = acc()

    # when: a market order is sent to an non empty matching unit
    MatchingService.create(account, "BTC_EUR", :LIMIT, :SELL, 10, 100, :GTC)
    MatchingService.create(account, "BTC_EUR", :MARKET, :SELL, 100, :GTC)

    # then: a matching unit should save the trade order but it should not be visible to the order book
    assert MatchingServiceTestHelpers.get_sell_book_order_count() == 2

    assert OrderBookService.get_volumes("BTC_EUR", :SELL) == [
             {10, 100}
           ]
  end

  test "process/1 market buy order with existing buy limit" do
    account = acc()

    # when: a market order is sent to an non empty matching unit
    MatchingService.create(account, "BTC_EUR", :LIMIT, :BUY, 10, 10, :GTC)
    MatchingService.create(account, "BTC_EUR", :MARKET, :BUY, 10, :GTC)

    # then: a matching unit should save the trade order but it should not be visible to the order book
    assert MatchingServiceTestHelpers.get_buy_book_order_count() == 2

    assert OrderBookService.get_volumes("BTC_EUR", :BUY) == [
             {10.0, 10}
           ]
  end

  test "process/1 market sell order with existing buy limit" do
    account = acc()
    account2 = acc2()

    # when:
    MatchingService.create(account, "BTC_EUR", :LIMIT, :BUY, 10, 100, :GTC)
    MatchingService.create(account2, "BTC_EUR", :MARKET, :SELL, 100, :GTC)

    # then:
    assert MatchingServiceTestHelpers.get_trade_count() == 1
    assert MatchingServiceTestHelpers.get_trade_prices() == [10]
    assert MatchingServiceTestHelpers.get_buy_book_order_count() == 0
    assert MatchingServiceTestHelpers.get_sell_book_order_count() == 0
    assert OrderBookService.get_volumes("BTC_EUR", :BUY) == []
    assert OrderBookService.get_volumes("BTC_EUR", :SELL) == []
  end

  test "process/1 market buy order with existing sell limit" do
    account = acc()
    account2 = acc2()

    # when:
    MatchingService.create(account, "BTC_EUR", :LIMIT, :SELL, 10, 100, :GTC)
    MatchingService.create(account2, "BTC_EUR", :MARKET, :BUY, 100, :GTC)

    # then:
    assert MatchingServiceTestHelpers.get_trade_count() == 1
    assert MatchingServiceTestHelpers.get_trade_prices() == [10]
    assert MatchingServiceTestHelpers.get_buy_book_order_count() == 0
    assert MatchingServiceTestHelpers.get_sell_book_order_count() == 1
    assert OrderBookService.get_volumes("BTC_EUR", :BUY) == []
    assert OrderBookService.get_volumes("BTC_EUR", :SELL) == [{10, 90}]
  end

  test "process/1 market sell order with multiple existing buy limits" do
    account = acc()
    account2 = acc2()

    # when:
    MatchingService.create(account, "BTC_EUR", :LIMIT, :BUY, 5, 10, :GTC)
    MatchingService.create(account, "BTC_EUR", :LIMIT, :BUY, 4, 10, :GTC)
    MatchingService.create(account, "BTC_EUR", :LIMIT, :BUY, 7, 10, :GTC)
    MatchingService.create(account2, "BTC_EUR", :MARKET, :SELL, 30, :GTC)

    # then:
    assert MatchingServiceTestHelpers.get_trade_count() == 3
    assert MatchingServiceTestHelpers.get_trade_prices() == [7, 5, 4]
    assert MatchingServiceTestHelpers.get_buy_book_order_count() == 0
    assert MatchingServiceTestHelpers.get_sell_book_order_count() == 0
    assert OrderBookService.get_volumes("BTC_EUR", :BUY) == []
    assert OrderBookService.get_volumes("BTC_EUR", :SELL) == []
  end

  test "process/1 market buy order with multiple existing sell limits" do
    account = acc()
    account2 = acc2()

    # when:
    MatchingService.create(account, "BTC_EUR", :LIMIT, :SELL, 5, 10, :GTC)
    MatchingService.create(account, "BTC_EUR", :LIMIT, :SELL, 4, 10, :GTC)
    MatchingService.create(account, "BTC_EUR", :LIMIT, :SELL, 7, 10, :GTC)
    MatchingService.create(account2, "BTC_EUR", :MARKET, :BUY, 160, :GTC)

    # then:
    assert MatchingServiceTestHelpers.get_trade_count() == 3
    assert MatchingServiceTestHelpers.get_trade_prices() == [4, 5, 7]
    assert MatchingServiceTestHelpers.get_buy_book_order_count() == 0
    assert MatchingServiceTestHelpers.get_sell_book_order_count() == 0
    assert OrderBookService.get_volumes("BTC_EUR", :BUY) == []
    assert OrderBookService.get_volumes("BTC_EUR", :SELL) == []
  end

  test "process/1 limit buy order with existing market sell" do
    account = acc()
    account2 = acc2()

    # when: the opposite book does not have outstanding limit orders
    MatchingService.create(account2, "BTC_EUR", :MARKET, :SELL, 10, :GTC)
    MatchingService.create(account, "BTC_EUR", :LIMIT, :BUY, 5, 10, :GTC)

    # then: the trade settles at the incoming order’s limit price
    assert MatchingServiceTestHelpers.get_trade_count() == 1
    assert MatchingServiceTestHelpers.get_trade_prices() == [5]
    assert MatchingServiceTestHelpers.get_buy_book_order_count() == 0
    assert MatchingServiceTestHelpers.get_sell_book_order_count() == 0
    assert OrderBookService.get_volumes("BTC_EUR", :BUY) == []
    assert OrderBookService.get_volumes("BTC_EUR", :SELL) == []
  end

  test "process/1 limit sell order with existing market buy" do
    account = acc()
    account2 = acc2()

    # when: the opposite book does not have outstanding limit orders
    MatchingService.create(account2, "BTC_EUR", :MARKET, :BUY, 10, :GTC)
    MatchingService.create(account, "BTC_EUR", :LIMIT, :SELL, 5, 10, :GTC)

    # then: the trade settles at the incoming order’s limit price
    assert MatchingServiceTestHelpers.get_trade_count() == 1
    assert MatchingServiceTestHelpers.get_trade_prices() == [5]
    assert MatchingServiceTestHelpers.get_buy_book_order_count() == 0
    assert MatchingServiceTestHelpers.get_sell_book_order_count() == 1
    assert OrderBookService.get_volumes("BTC_EUR", :BUY) == []
    assert OrderBookService.get_volumes("BTC_EUR", :SELL) == [{5.0, 8.0}]
  end

  test "process/1 limit buy order with existing market sell and better limit sell" do
    account = acc()
    account2 = acc2()

    # when: the opposite book has a market and a limit order
    MatchingService.create(account2, "BTC_EUR", :LIMIT, :SELL, 5, 10, :GTC)
    MatchingService.create(account2, "BTC_EUR", :MARKET, :SELL, 10, :GTC)
    MatchingService.create(account, "BTC_EUR", :LIMIT, :BUY, 6, 20, :GTC)

    # then: the trade settles at the book order’s limit price as if market order did not exist
    assert MatchingServiceTestHelpers.get_trade_count() == 2
    assert MatchingServiceTestHelpers.get_trade_prices() == [5, 5]
    assert MatchingServiceTestHelpers.get_buy_book_order_count() == 0
    assert MatchingServiceTestHelpers.get_sell_book_order_count() == 0
    assert OrderBookService.get_volumes("BTC_EUR", :BUY) == []
    assert OrderBookService.get_volumes("BTC_EUR", :SELL) == []
  end

  test "process/1 limit sell order with existing market buy and better limit buy" do
    account = acc()
    account2 = acc2()

    # when: the opposite book has a market and a limit order
    MatchingService.create(account2, "BTC_EUR", :LIMIT, :BUY, 5, 10, :GTC)
    MatchingService.create(account2, "BTC_EUR", :MARKET, :BUY, 50, :GTC)
    MatchingService.create(account, "BTC_EUR", :LIMIT, :SELL, 4, 20, :GTC)

    # then: the trade settles at the book order’s limit price as if market order did not exist
    assert MatchingServiceTestHelpers.get_trade_count() == 2
    assert MatchingServiceTestHelpers.get_trade_prices() == [5, 5]
    assert MatchingServiceTestHelpers.get_buy_book_order_count() == 0
    assert MatchingServiceTestHelpers.get_sell_book_order_count() == 0
    assert OrderBookService.get_volumes("BTC_EUR", :BUY) == []
    assert OrderBookService.get_volumes("BTC_EUR", :SELL) == []
  end

  test "process/1 limit buy order with existing market sell and worse limit sell" do
    account = acc()
    account2 = acc2()

    # when: the opposite book has a market and a limit order
    MatchingService.create(account2, "BTC_EUR", :LIMIT, :SELL, 7, 10, :GTC)
    MatchingService.create(account2, "BTC_EUR", :MARKET, :SELL, 10, :GTC)
    MatchingService.create(account, "BTC_EUR", :LIMIT, :BUY, 6, 10, :GTC)

    # then: the trade settles at the incoming order’s limit price
    assert MatchingServiceTestHelpers.get_trade_count() == 1
    assert MatchingServiceTestHelpers.get_trade_prices() == [6]
    assert MatchingServiceTestHelpers.get_buy_book_order_count() == 0
    assert MatchingServiceTestHelpers.get_sell_book_order_count() == 1
    assert OrderBookService.get_volumes("BTC_EUR", :BUY) == []
    assert OrderBookService.get_volumes("BTC_EUR", :SELL) == [{7, 10}]
  end

  test "process/1 limit sell order with existing market buy and worse limit buy" do
    account = acc()
    account2 = acc2()

    # when: the opposite book has a market and a limit order
    MatchingService.create(account2, "BTC_EUR", :LIMIT, :BUY, 4, 10, :GTC)
    MatchingService.create(account2, "BTC_EUR", :MARKET, :BUY, 100, :GTC)
    MatchingService.create(account, "BTC_EUR", :LIMIT, :SELL, 5, 20, :GTC)

    # then: the trade settles at the incoming order’s limit price
    assert MatchingServiceTestHelpers.get_trade_count() == 1
    assert MatchingServiceTestHelpers.get_trade_prices() == [5]
    assert MatchingServiceTestHelpers.get_buy_book_order_count() == 1
    assert MatchingServiceTestHelpers.get_sell_book_order_count() == 0
    assert OrderBookService.get_volumes("BTC_EUR", :BUY) == [{4, 10}]
    assert OrderBookService.get_volumes("BTC_EUR", :SELL) == []
  end

  test "process/1 market buy order with existing market sell and no reference price" do
    account = acc()
    account2 = acc2()

    # when: there is not a trade with a reference price
    MatchingService.create(account2, "BTC_EUR", :MARKET, :SELL, 10, :GTC)
    MatchingService.create(account, "BTC_EUR", :MARKET, :BUY, 50, :GTC)

    # then: the trade settles at a reference price
    assert MatchingServiceTestHelpers.get_trade_count() == 0
    assert MatchingServiceTestHelpers.get_trade_prices() == []
    assert MatchingServiceTestHelpers.get_buy_book_order_count() == 1
    assert MatchingServiceTestHelpers.get_sell_book_order_count() == 1
    assert OrderBookService.get_volumes("BTC_EUR", :BUY) == []
    assert OrderBookService.get_volumes("BTC_EUR", :SELL) == []
  end

  test "process/1 market sell order with existing market buy and no reference price" do
    account = acc()
    account2 = acc2()

    # when: there is not a trade with a reference price
    MatchingService.create(account2, "BTC_EUR", :MARKET, :BUY, 50, :GTC)
    MatchingService.create(account, "BTC_EUR", :MARKET, :SELL, 10, :GTC)

    # then: the trade settles at a reference price
    assert MatchingServiceTestHelpers.get_trade_count() == 0
    assert MatchingServiceTestHelpers.get_trade_prices() == []
    assert MatchingServiceTestHelpers.get_buy_book_order_count() == 1
    assert MatchingServiceTestHelpers.get_sell_book_order_count() == 1
    assert OrderBookService.get_volumes("BTC_EUR", :BUY) == []
    assert OrderBookService.get_volumes("BTC_EUR", :SELL) == []
  end

  test "process/1 market buy order with existing market sell and a limit sell" do
    account = acc()
    account2 = acc2()

    # when: the opposite book has a market and a limit order
    MatchingService.create(account2, "BTC_EUR", :LIMIT, :SELL, 5, 10, :GTC)
    MatchingService.create(account2, "BTC_EUR", :MARKET, :SELL, 10, :GTC)
    MatchingService.create(account, "BTC_EUR", :MARKET, :BUY, 50, :GTC)

    # then: the trade settles at the book order’s limit price
    assert MatchingServiceTestHelpers.get_trade_count() == 1
    assert MatchingServiceTestHelpers.get_trade_prices() == [5]
    assert MatchingServiceTestHelpers.get_buy_book_order_count() == 0
    assert MatchingServiceTestHelpers.get_sell_book_order_count() == 1
    assert OrderBookService.get_volumes("BTC_EUR", :BUY) == []
    assert OrderBookService.get_volumes("BTC_EUR", :SELL) == [{5, 10}]
  end

  test "process/1 market sell order with existing market buy and a limit buy" do
    account = acc()
    account2 = acc2()

    # when: the opposite book has a market and a limit order
    MatchingService.create(account2, "BTC_EUR", :LIMIT, :BUY, 5, 10, :GTC)
    MatchingService.create(account2, "BTC_EUR", :MARKET, :BUY, 50, :GTC)
    MatchingService.create(account, "BTC_EUR", :MARKET, :SELL, 10, :GTC)

    # then: the trade settles at the book order’s limit price
    assert MatchingServiceTestHelpers.get_trade_count() == 1
    assert MatchingServiceTestHelpers.get_trade_prices() == [5]
    assert MatchingServiceTestHelpers.get_buy_book_order_count() == 1
    assert MatchingServiceTestHelpers.get_sell_book_order_count() == 0
    assert OrderBookService.get_volumes("BTC_EUR", :BUY) == [{5, 10}]
    assert OrderBookService.get_volumes("BTC_EUR", :SELL) == []
  end

  test "process/1 market buy order with existing market sell and a reference price" do
    account = acc()
    account2 = acc2()

    # when: there is a trade with a reference price
    MatchingService.create(account2, "BTC_EUR", :LIMIT, :SELL, 5, 1, :GTC)
    MatchingService.create(account, "BTC_EUR", :LIMIT, :BUY, 5, 1, :GTC)
    MatchingService.create(account2, "BTC_EUR", :MARKET, :SELL, 10, :GTC)
    MatchingService.create(account, "BTC_EUR", :MARKET, :BUY, 50, :GTC)

    # then: the trade settles at a reference price
    assert MatchingServiceTestHelpers.get_trade_count() == 2
    assert MatchingServiceTestHelpers.get_trade_prices() == [5, 5]
    assert MatchingServiceTestHelpers.get_buy_book_order_count() == 0
    assert MatchingServiceTestHelpers.get_sell_book_order_count() == 0
    assert OrderBookService.get_volumes("BTC_EUR", :BUY) == []
    assert OrderBookService.get_volumes("BTC_EUR", :SELL) == []
  end

  test "process/1 market sell order with existing market buy and a reference price" do
    account = acc()
    account2 = acc2()

    # when: the opposite book has a market and a limit order
    MatchingService.create(account2, "BTC_EUR", :LIMIT, :SELL, 5, 1, :GTC)
    MatchingService.create(account, "BTC_EUR", :LIMIT, :BUY, 5, 1, :GTC)
    MatchingService.create(account2, "BTC_EUR", :MARKET, :BUY, 50, :GTC)
    MatchingService.create(account, "BTC_EUR", :MARKET, :SELL, 10, :GTC)

    # then: the trade settles at a reference price
    assert MatchingServiceTestHelpers.get_trade_count() == 2
    assert MatchingServiceTestHelpers.get_trade_prices() == [5, 5]
    assert MatchingServiceTestHelpers.get_buy_book_order_count() == 0
    assert MatchingServiceTestHelpers.get_sell_book_order_count() == 0
    assert OrderBookService.get_volumes("BTC_EUR", :BUY) == []
    assert OrderBookService.get_volumes("BTC_EUR", :SELL) == []
  end

  test "process/1 limit buy order with existing market sell and a reference price" do
    account = acc()
    account2 = acc2()

    # when: there is a trade with a reference price
    MatchingService.create(account2, "BTC_EUR", :LIMIT, :SELL, 5, 1, :GTC)
    MatchingService.create(account, "BTC_EUR", :LIMIT, :BUY, 5, 1, :GTC)
    MatchingService.create(account2, "BTC_EUR", :MARKET, :SELL, 10, :GTC)
    MatchingService.create(account, "BTC_EUR", :LIMIT, :BUY, 6, 10, :GTC)

    # then: the trade settles at a reference price
    assert MatchingServiceTestHelpers.get_trade_count() == 2
    assert MatchingServiceTestHelpers.get_trade_prices() == [5, 5]
    assert MatchingServiceTestHelpers.get_buy_book_order_count() == 0
    assert MatchingServiceTestHelpers.get_sell_book_order_count() == 0
    assert OrderBookService.get_volumes("BTC_EUR", :BUY) == []
    assert OrderBookService.get_volumes("BTC_EUR", :SELL) == []
  end

  # ---- TESTS BELOW REQUIRE VARIOUS HACKS TO SIMULATE TIMESTAMP INCREMENTS
  test "process/1 market buy order with existing market sell and multiple reference prices" do
    account = acc()
    account2 = acc2()

    # when: there is a trade with a reference price
    MatchingService.create(account2, "BTC_EUR", :LIMIT, :SELL, 4, 1, :GTC)
    MatchingService.create(account, "BTC_EUR", :LIMIT, :BUY, 4, 1, :GTC)

    DB.execute(
      "UPDATE trade SET created_at = current_timestamp + interval '1 second' WHERE price = 4"
    )

    MatchingService.create(account2, "BTC_EUR", :LIMIT, :SELL, 6, 1, :GTC)
    MatchingService.create(account, "BTC_EUR", :LIMIT, :BUY, 6, 1, :GTC)

    DB.execute(
      "UPDATE trade SET created_at = current_timestamp + interval '2 second' WHERE price = 6"
    )

    MatchingService.create(account2, "BTC_EUR", :LIMIT, :SELL, 5, 1, :GTC)
    MatchingService.create(account, "BTC_EUR", :LIMIT, :BUY, 5, 1, :GTC)

    DB.execute(
      "UPDATE trade SET created_at = current_timestamp + interval '3 second' WHERE price = 5"
    )

    MatchingService.create(account2, "BTC_EUR", :MARKET, :SELL, 10, :GTC)
    MatchingService.create(account, "BTC_EUR", :MARKET, :BUY, 50, :GTC)

    # then: the trade settles at last reference price
    assert MatchingServiceTestHelpers.get_trade_count() == 4
    assert MatchingServiceTestHelpers.get_trade_prices() |> List.first() == 5
    assert MatchingServiceTestHelpers.get_buy_book_order_count() == 0
    assert MatchingServiceTestHelpers.get_sell_book_order_count() == 0
    assert OrderBookService.get_volumes("BTC_EUR", :BUY) == []
    assert OrderBookService.get_volumes("BTC_EUR", :SELL) == []
  end

  test "process/1 market sell order with existing market buy and multiple reference prices" do
    account = acc()
    account2 = acc2()

    # when: the opposite book has a market and a limit order
    MatchingService.create(account2, "BTC_EUR", :LIMIT, :SELL, 4, 1, :GTC)
    MatchingService.create(account, "BTC_EUR", :LIMIT, :BUY, 4, 1, :GTC)

    DB.execute(
      "UPDATE trade SET created_at = current_timestamp + interval '1 second' WHERE price = 4"
    )

    MatchingService.create(account2, "BTC_EUR", :LIMIT, :SELL, 6, 1, :GTC)
    MatchingService.create(account, "BTC_EUR", :LIMIT, :BUY, 6, 1, :GTC)

    DB.execute(
      "UPDATE trade SET created_at = current_timestamp + interval '2 second' WHERE price = 6"
    )

    MatchingService.create(account2, "BTC_EUR", :LIMIT, :SELL, 5, 1, :GTC)
    MatchingService.create(account, "BTC_EUR", :LIMIT, :BUY, 5, 1, :GTC)

    DB.execute(
      "UPDATE trade SET created_at = current_timestamp + interval '3 second' WHERE price = 5"
    )

    MatchingService.create(account2, "BTC_EUR", :MARKET, :BUY, 50, :GTC)
    MatchingService.create(account, "BTC_EUR", :MARKET, :SELL, 10, :GTC)

    # then: the trade settles at last reference price
    assert MatchingServiceTestHelpers.get_trade_count() == 4
    assert MatchingServiceTestHelpers.get_trade_prices() |> List.first() == 5
    assert MatchingServiceTestHelpers.get_buy_book_order_count() == 0
    assert MatchingServiceTestHelpers.get_sell_book_order_count() == 0
    assert OrderBookService.get_volumes("BTC_EUR", :BUY) == []
    assert OrderBookService.get_volumes("BTC_EUR", :SELL) == []
  end
end
