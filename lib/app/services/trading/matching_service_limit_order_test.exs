defmodule MatchingServiceLimitOrderTest do
  use DataCase

  import TestUtils

  test "process/1 limit sell order save" do
    # when: a limit order is sent to an empty matching unit
    res = MatchingService.create(acc(), "BTC_EUR", :LIMIT, :SELL, 10, 100, :GTC)

    # then: a matching unit should save the trade order on save order to the order book
    assert res != nil
    assert res == DB.query_val("SELECT pub_id FROM trade_order WHERE pub_id = '#{res}'")
    assert MatchingServiceTestHelpers.get_sell_book_order_count() == 1

    assert OrderBookService.get_volumes("BTC_EUR", :SELL) == [
             {10, 100}
           ]
  end

  test "process/1 limit buy order save" do
    # when: a limit order is sent to an empty matching unit
    MatchingService.create(acc(), "BTC_EUR", :LIMIT, :BUY, 10, 100, :GTC)

    # then: a matching unit should save the orderSELL
    assert MatchingServiceTestHelpers.get_buy_book_order_count() == 1

    assert OrderBookService.get_volumes("BTC_EUR", :BUY) == [
             {10, 100}
           ]
  end

  test "process/1 limit no match case incoming buy" do
    # when: there is a SELL order in the book and a BUY limit order arrives that does not cross
    MatchingService.create(acc(), "BTC_EUR", :LIMIT, :SELL, 10, 100, :GTC)
    MatchingService.create(acc2(), "BTC_EUR", :LIMIT, :BUY, 9, 100, :GTC)

    # then: the book should have both orders and no trade should be generated
    assert MatchingServiceTestHelpers.get_sell_book_order_count() == 1
    assert MatchingServiceTestHelpers.get_buy_book_order_count() == 1
    assert MatchingServiceTestHelpers.get_trade_count() == 0

    assert OrderBookService.get_volumes("BTC_EUR", :SELL) == [
             {10, 100}
           ]

    assert OrderBookService.get_volumes("BTC_EUR", :BUY) == [
             {9, 100}
           ]
  end

  test "process/1 limit no match case incoming sell" do
    # when: there is a BUY order in the book and a SELL limit order arrives that does not cross
    MatchingService.create(acc(), "BTC_EUR", :LIMIT, :BUY, 9, 100, :GTC)
    MatchingService.create(acc2(), "BTC_EUR", :LIMIT, :SELL, 10, 100, :GTC)

    # then: the book should have both orders and no trade should be generated
    assert MatchingServiceTestHelpers.get_sell_book_order_count() == 1
    assert MatchingServiceTestHelpers.get_buy_book_order_count() == 1
    assert MatchingServiceTestHelpers.get_trade_count() == 0

    assert OrderBookService.get_volumes("BTC_EUR", :SELL) == [
             {10, 100}
           ]

    assert OrderBookService.get_volumes("BTC_EUR", :BUY) == [
             {9, 100}
           ]
  end

  test "process/1 limit exact match incoming buy" do
    # when: there is a SELL order in the book
    MatchingService.create(acc(), "BTC_EUR", :LIMIT, :SELL, 10, 100, :GTC)

    # then:
    assert MatchingServiceTestHelpers.get_sell_book_order_count() == 1
    assert MatchingServiceTestHelpers.get_buy_book_order_count() == 0
    assert MatchingServiceTestHelpers.get_trade_count() == 0

    assert MatchingServiceTestHelpers.get_available_limit_volume(:SELL, 10) ==
             100

    assert OrderBookService.get_volumes("BTC_EUR", :SELL) == [
             {10, 100}
           ]

    # when: a BUY limit order arrives that crossed
    MatchingService.create(acc2(), "BTC_EUR", :LIMIT, :BUY, 10, 100, :GTC)

    # then: the book should have no orders and a single trade should be generated
    assert MatchingServiceTestHelpers.get_sell_book_order_count() == 0
    assert MatchingServiceTestHelpers.get_buy_book_order_count() == 0
    assert MatchingServiceTestHelpers.get_trade_count() == 1
    assert MatchingServiceTestHelpers.get_available_limit_volume(:SELL, 10) == 0
    assert OrderBookService.get_volumes("BTC_EUR", :SELL) == []
    assert OrderBookService.get_volumes("BTC_EUR", :BUY) == []
  end

  test "process/1 limit partial match incoming buy single trade" do
    MatchingService.create(acc(), "BTC_EUR", :LIMIT, :SELL, 10.00, 100.00, :GTC)

    # then:
    assert MatchingServiceTestHelpers.get_sell_book_order_count() == 1
    assert MatchingServiceTestHelpers.get_buy_book_order_count() == 0
    assert MatchingServiceTestHelpers.get_trade_count() == 0

    assert OrderBookService.get_volumes("BTC_EUR", :SELL) == [
             {10.00, 100}
           ]

    # when: incoming buy order is only partially matched
    MatchingService.create(acc2(), "BTC_EUR", :LIMIT, :BUY, 10.00, 50.00, :GTC)

    # then: the book should have one sell order and a single trade should be generated
    assert MatchingServiceTestHelpers.get_sell_book_order_count() == 1
    assert MatchingServiceTestHelpers.get_buy_book_order_count() == 0
    assert MatchingServiceTestHelpers.get_trade_count() == 1

    assert OrderBookService.get_volumes("BTC_EUR", :SELL) == [
             {10.00, 50}
           ]

    assert OrderBookService.get_volumes("BTC_EUR", :BUY) == []
  end

  test "process/1 limit overflow match incoming buy single trade" do
    # when: there is a SELL order in the book
    MatchingService.create(acc(), "BTC_EUR", :LIMIT, :SELL, 10.00, 10.00, :GTC)

    # then:
    assert MatchingServiceTestHelpers.get_sell_book_order_count() == 1
    assert MatchingServiceTestHelpers.get_buy_book_order_count() == 0
    assert MatchingServiceTestHelpers.get_trade_count() == 0
    assert MatchingServiceTestHelpers.get_available_limit_volume(:SELL, 10) == 10

    # when: a BUY limit order arrives that crosses and is more that the book amount
    MatchingService.create(acc2(), "BTC_EUR", :LIMIT, :BUY, 10.00, 15, :GTC)

    # then: the book should be one buy order,  no sell orders and a one trade
    assert MatchingServiceTestHelpers.get_sell_book_order_count() == 0
    assert MatchingServiceTestHelpers.get_buy_book_order_count() == 1
    assert MatchingServiceTestHelpers.get_trade_count() == 1
    assert OrderBookService.get_volumes("BTC_EUR", :SELL) == []

    assert OrderBookService.get_volumes("BTC_EUR", :BUY) == [
             {10.00, 5}
           ]
  end

  test "process/1 limit partial match incoming sell single trade" do
    # when: there is a BUY order
    MatchingService.create(acc(), "BTC_EUR", :LIMIT, :BUY, 10.00, 100.00, :GTC)

    # then:
    assert MatchingServiceTestHelpers.get_sell_book_order_count() == 0
    assert MatchingServiceTestHelpers.get_buy_book_order_count() == 1
    assert MatchingServiceTestHelpers.get_trade_count() == 0

    assert MatchingServiceTestHelpers.get_available_limit_volume(:BUY, 10) == 100

    # when: incoming SELL order is only partially matched
    MatchingService.create(acc2(), "BTC_EUR", :LIMIT, :SELL, 10.00, 50.00, :GTC)

    # then: the book should have one BUY order and a 1 trade should be generated
    assert MatchingServiceTestHelpers.get_sell_book_order_count() == 0
    assert MatchingServiceTestHelpers.get_buy_book_order_count() == 1
    assert MatchingServiceTestHelpers.get_trade_count() == 1
    assert MatchingServiceTestHelpers.get_available_limit_volume(:BUY, 10) == 50
    assert OrderBookService.get_volumes("BTC_EUR", :SELL) == []

    assert OrderBookService.get_volumes("BTC_EUR", :BUY) == [
             {10.00, 50}
           ]
  end

  test "process/1 limit overflow match incoming sell single trade" do
    # when: there is a BUY order in the book
    MatchingService.create(acc(), "BTC_EUR", :LIMIT, :BUY, 10.00, 100.00, :GTC)

    # then:
    assert MatchingServiceTestHelpers.get_sell_book_order_count() == 0
    assert MatchingServiceTestHelpers.get_buy_book_order_count() == 1
    assert MatchingServiceTestHelpers.get_trade_count() == 0
    assert MatchingServiceTestHelpers.get_available_limit_volume(:BUY, 10) == 100

    # when: a SELL limit order arrives that crosses and is more that the book amount
    MatchingService.create(acc2(), "BTC_EUR", :LIMIT, :SELL, 10.00, 150.00, :GTC)

    # then: the book should be one SELL order,  no BUY orders and 1 trade
    assert MatchingServiceTestHelpers.get_sell_book_order_count() == 1
    assert MatchingServiceTestHelpers.get_buy_book_order_count() == 0
    assert MatchingServiceTestHelpers.get_trade_count() == 1
    assert MatchingServiceTestHelpers.get_available_limit_volume(:SELL, 10) == 50

    assert OrderBookService.get_volumes("BTC_EUR", :SELL) == [
             {10.00, 50}
           ]

    assert OrderBookService.get_volumes("BTC_EUR", :BUY) == []
  end

  test "process/1 limit exact match incoming buys multiple trades" do
    # when: there is a SELL order in the book and 2 BUY limit order arrive that cross
    MatchingService.create(acc(), "BTC_EUR", :LIMIT, :SELL, 10.00, 10.00, :GTC)

    # then:
    assert MatchingServiceTestHelpers.get_sell_book_order_count() == 1
    assert MatchingServiceTestHelpers.get_buy_book_order_count() == 0
    assert MatchingServiceTestHelpers.get_trade_count() == 0

    assert MatchingServiceTestHelpers.get_available_limit_volume(:SELL, 10) == 10

    # when: incoming buy order that are partially matched
    trading_account = acc2()
    MatchingService.create(trading_account, "BTC_EUR", :LIMIT, :BUY, 10.00, 5.00, :GTC)
    MatchingService.create(trading_account, "BTC_EUR", :LIMIT, :BUY, 10.00, 5.00, :GTC)

    # then: the book should have 2 trades should be generated
    assert MatchingServiceTestHelpers.get_sell_book_order_count() == 0
    assert MatchingServiceTestHelpers.get_buy_book_order_count() == 0
    assert MatchingServiceTestHelpers.get_trade_count() == 2
    assert MatchingServiceTestHelpers.get_available_limit_volume(:SELL, 10) == 0
    assert OrderBookService.get_volumes("BTC_EUR", :SELL) == []
    assert OrderBookService.get_volumes("BTC_EUR", :BUY) == []
  end

  test "process/1 limit exact match incoming sell multiple trades" do
    trading_account = acc()
    trading_account2 = acc2()

    # when: there is a BUY order in the book and 2 SELL limit order arrive that cross
    MatchingService.create(trading_account, "BTC_EUR", :LIMIT, :BUY, 10.00, 10.00, :GTC)

    # then:
    assert MatchingServiceTestHelpers.get_sell_book_order_count() == 0
    assert MatchingServiceTestHelpers.get_buy_book_order_count() == 1
    assert MatchingServiceTestHelpers.get_trade_count() == 0
    assert MatchingServiceTestHelpers.get_available_limit_volume(:BUY, 10) == 10

    # when: incoming buy order is only partially matched
    MatchingService.create(trading_account2, "BTC_EUR", :LIMIT, :SELL, 10.00, 5.00, :GTC)
    MatchingService.create(trading_account2, "BTC_EUR", :LIMIT, :SELL, 10.00, 5.00, :GTC)

    # then: the book should have 2 trades should be generated
    assert MatchingServiceTestHelpers.get_sell_book_order_count() == 0
    assert MatchingServiceTestHelpers.get_buy_book_order_count() == 0
    assert MatchingServiceTestHelpers.get_trade_count() == 2
    assert OrderBookService.get_volumes("BTC_EUR", :SELL) == []
    assert OrderBookService.get_volumes("BTC_EUR", :BUY) == []
  end

  test "process/1 limit partial match multiple book sells to multiple trades" do
    # given:
    trading_account = acc()
    trading_account2 = acc2()

    # when: there are multiple SELL orders in the book and BUY limit order arrive that cross but can fill only partially one of the orders
    MatchingService.create(trading_account, "BTC_EUR", :LIMIT, :SELL, 10.00, 50.00, :GTC)
    MatchingService.create(trading_account, "BTC_EUR", :LIMIT, :SELL, 10.00, 50.00, :GTC)

    # then:
    assert MatchingServiceTestHelpers.get_sell_book_order_count() == 2
    assert MatchingServiceTestHelpers.get_buy_book_order_count() == 0
    assert MatchingServiceTestHelpers.get_trade_count() == 0

    assert MatchingServiceTestHelpers.get_available_limit_volume(:SELL, 10) ==
             100

    # when: incoming buy order is only partially matched
    MatchingService.create(trading_account2, "BTC_EUR", :LIMIT, :BUY, 10.00, 75.00, :GTC)

    # then: the book should have 2 trades should be generated
    assert MatchingServiceTestHelpers.get_sell_book_order_count() == 1
    assert MatchingServiceTestHelpers.get_buy_book_order_count() == 0
    assert MatchingServiceTestHelpers.get_trade_count() == 2

    assert OrderBookService.get_volumes("BTC_EUR", :SELL) == [
             {10.00, 25}
           ]

    assert OrderBookService.get_volumes("BTC_EUR", :BUY) == []
  end

  test "process/1 limit partial match  multiple book buys to multiple trades" do
    # given:
    trading_account = acc()
    trading_account2 = acc2()

    # when: there are multiple BUY orders in the book and SELL limit order arrive that cross but can fill only partially one of the orders
    MatchingService.create(trading_account, "BTC_EUR", :LIMIT, :BUY, 10.00, 50.00, :GTC)
    MatchingService.create(trading_account, "BTC_EUR", :LIMIT, :BUY, 10.00, 50.00, :GTC)

    # then:
    assert MatchingServiceTestHelpers.get_sell_book_order_count() == 0
    assert MatchingServiceTestHelpers.get_buy_book_order_count() == 2
    assert MatchingServiceTestHelpers.get_trade_count() == 0
    assert MatchingServiceTestHelpers.get_available_limit_volume(:BUY, 10) == 100

    # when: incoming SELL order is only partially matched
    MatchingService.create(trading_account2, "BTC_EUR", :LIMIT, :SELL, 10.00, 75.00, :GTC)

    # then: the book should have 2 trades should be generated
    assert MatchingServiceTestHelpers.get_sell_book_order_count() == 0
    assert MatchingServiceTestHelpers.get_buy_book_order_count() == 1
    assert MatchingServiceTestHelpers.get_trade_count() == 2
    assert OrderBookService.get_volumes("BTC_EUR", :SELL) == []

    assert OrderBookService.get_volumes("BTC_EUR", :BUY) == [
             {10.00, 25}
           ]
  end

  test "process/1 limit exact match multiple book sells to multiple trades" do
    # given:
    trading_account = acc()
    trading_account2 = acc2()

    # when:
    MatchingService.create(trading_account, "BTC_EUR", :LIMIT, :SELL, 10.00, 10.00, :GTC)
    MatchingService.create(trading_account, "BTC_EUR", :LIMIT, :SELL, 10.00, 10.00, :GTC)
    MatchingService.create(trading_account, "BTC_EUR", :LIMIT, :SELL, 10.00, 10.00, :GTC)
    MatchingService.create(trading_account, "BTC_EUR", :LIMIT, :SELL, 10.00, 10.00, :GTC)

    # then:
    assert MatchingServiceTestHelpers.get_sell_book_order_count() == 4
    assert MatchingServiceTestHelpers.get_buy_book_order_count() == 0
    assert MatchingServiceTestHelpers.get_trade_count() == 0
    assert MatchingServiceTestHelpers.get_available_limit_volume(:SELL, 10) == 40

    # when: incoming buy order is only partially matched
    MatchingService.create(trading_account2, "BTC_EUR", :LIMIT, :BUY, 10.00, 40.00, :GTC)

    # then: the book should have 2 trades should be generated
    assert MatchingServiceTestHelpers.get_sell_book_order_count() == 0
    assert MatchingServiceTestHelpers.get_buy_book_order_count() == 0
    assert MatchingServiceTestHelpers.get_trade_count() == 4
    assert OrderBookService.get_volumes("BTC_EUR", :SELL) == []
    assert OrderBookService.get_volumes("BTC_EUR", :BUY) == []
  end

  test "process/1 limit exact match multiple book buy to multiple trades" do
    # given:
    trading_account = acc()
    trading_account2 = acc2()

    # when:
    MatchingService.create(trading_account, "BTC_EUR", :LIMIT, :BUY, 10.00, 5.00, :GTC)
    MatchingService.create(trading_account, "BTC_EUR", :LIMIT, :BUY, 10.00, 5.00, :GTC)
    MatchingService.create(trading_account, "BTC_EUR", :LIMIT, :BUY, 10.00, 5.00, :GTC)
    MatchingService.create(trading_account, "BTC_EUR", :LIMIT, :BUY, 10.00, 5.00, :GTC)

    # then:
    assert MatchingServiceTestHelpers.get_sell_book_order_count() == 0
    assert MatchingServiceTestHelpers.get_buy_book_order_count() == 4
    assert MatchingServiceTestHelpers.get_trade_count() == 0
    assert MatchingServiceTestHelpers.get_available_limit_volume(:BUY, 10) == 20

    # when: incoming buy order is only partially matched
    MatchingService.create(trading_account2, "BTC_EUR", :LIMIT, :SELL, 10.00, 20.00, :GTC)

    # then: the book should have 2 trades should be generated
    assert MatchingServiceTestHelpers.get_sell_book_order_count() == 0
    assert MatchingServiceTestHelpers.get_buy_book_order_count() == 0
    assert MatchingServiceTestHelpers.get_trade_count() == 4
    assert OrderBookService.get_volumes("BTC_EUR", :SELL) == []
    assert OrderBookService.get_volumes("BTC_EUR", :BUY) == []
  end

  test "process/1 limit incomplete match multiple book sells to multiple trades" do
    # given:
    trading_account = acc()
    trading_account2 = acc2()

    # when: there are multiple SELL orders in the book and BUY limit order arrive that cross but can be only partially filled
    MatchingService.create(trading_account, "BTC_EUR", :LIMIT, :SELL, 10.00, 5.00, :GTC)
    MatchingService.create(trading_account, "BTC_EUR", :LIMIT, :SELL, 10.00, 5.00, :GTC)

    # then:
    assert MatchingServiceTestHelpers.get_sell_book_order_count() == 2
    assert MatchingServiceTestHelpers.get_buy_book_order_count() == 0
    assert MatchingServiceTestHelpers.get_trade_count() == 0
    assert MatchingServiceTestHelpers.get_available_limit_volume(:SELL, 10) == 10

    # when: incoming buy order is only partially matched
    MatchingService.create(trading_account2, "BTC_EUR", :LIMIT, :BUY, 10.00, 17, :GTC)

    # then: the book should have 2 trades should be generated
    assert MatchingServiceTestHelpers.get_sell_book_order_count() == 0
    assert MatchingServiceTestHelpers.get_buy_book_order_count() == 1
    assert MatchingServiceTestHelpers.get_trade_count() == 2
    assert OrderBookService.get_volumes("BTC_EUR", :SELL) == []

    assert OrderBookService.get_volumes("BTC_EUR", :BUY) == [
             {10.00, 7}
           ]
  end

  test "process/1 limit incomplete match multiple book buys to multiple trades" do
    # given:
    trading_account = acc()
    trading_account2 = acc2()

    # when: there are multiple BUY orders in the book and SELL limit order arrive that cross but can be only partially filled
    MatchingService.create(trading_account, "BTC_EUR", :LIMIT, :BUY, 10.00, 50.00, :GTC)
    MatchingService.create(trading_account, "BTC_EUR", :LIMIT, :BUY, 10.00, 50.00, :GTC)

    # then:
    assert MatchingServiceTestHelpers.get_sell_book_order_count() == 0
    assert MatchingServiceTestHelpers.get_buy_book_order_count() == 2
    assert MatchingServiceTestHelpers.get_trade_count() == 0
    assert MatchingServiceTestHelpers.get_available_limit_volume(:BUY, 10) == 100

    # when: incoming SELL order is only partially matched
    MatchingService.create(trading_account2, "BTC_EUR", :LIMIT, :SELL, 10.00, 175.00, :GTC)

    # then: the book should have 2 trades should be generated
    assert MatchingServiceTestHelpers.get_sell_book_order_count() == 1
    assert MatchingServiceTestHelpers.get_buy_book_order_count() == 0
    assert MatchingServiceTestHelpers.get_trade_count() == 2

    assert OrderBookService.get_volumes("BTC_EUR", :SELL) == [
             {10.00, 75}
           ]

    assert OrderBookService.get_volumes("BTC_EUR", :BUY) == []
  end

  test "process/1 limit partial match multiple book sells to multiple trades multiple prices" do
    # given:
    trading_account = acc()
    trading_account2 = acc2()

    # when: there are multiple SELL orders in the book and BUY limit order arrive that cross but can fill only partially one of the orders
    MatchingService.create(trading_account, "BTC_EUR", :LIMIT, :SELL, 9.00, 50.00, :GTC)
    MatchingService.create(trading_account, "BTC_EUR", :LIMIT, :SELL, 10.00, 50.00, :GTC)

    # then:
    assert MatchingServiceTestHelpers.get_sell_book_order_count() == 2
    assert MatchingServiceTestHelpers.get_buy_book_order_count() == 0
    assert MatchingServiceTestHelpers.get_trade_count() == 0

    # when: incoming buy order is only partially matched
    MatchingService.create(trading_account2, "BTC_EUR", :LIMIT, :BUY, 11.00, 75.00, :GTC)

    # then: the book should have 2 trades should be generated
    assert MatchingServiceTestHelpers.get_sell_book_order_count() == 1
    assert MatchingServiceTestHelpers.get_buy_book_order_count() == 0
    assert MatchingServiceTestHelpers.get_trade_count() == 2

    assert OrderBookService.get_volumes("BTC_EUR", :SELL) == [
             {10.00, 25}
           ]

    assert OrderBookService.get_volumes("BTC_EUR", :BUY) == []
  end

  test "process/1 limit partial match  multiple book buys to multiple trades multiple prices" do
    # given:
    trading_account = acc()
    trading_account2 = acc2()

    # when: there are multiple BUY orders in the book and SELL limit order arrive that cross but can fill only partially one of the orders
    MatchingService.create(trading_account, "BTC_EUR", :LIMIT, :BUY, 9.00, 50.00, :GTC)
    MatchingService.create(trading_account, "BTC_EUR", :LIMIT, :BUY, 10.00, 50.00, :GTC)

    # then:
    assert MatchingServiceTestHelpers.get_sell_book_order_count() == 0
    assert MatchingServiceTestHelpers.get_buy_book_order_count() == 2
    assert MatchingServiceTestHelpers.get_trade_count() == 0
    assert MatchingServiceTestHelpers.get_available_limit_volume(:BUY, 10) == 50
    assert MatchingServiceTestHelpers.get_available_limit_volume(:BUY, 9) == 100

    # when: incoming SELL order is only partially matched
    MatchingService.create(trading_account2, "BTC_EUR", :LIMIT, :SELL, 8.00, 75.00, :GTC)

    # then: the book should have 2 trades should be generated
    assert MatchingServiceTestHelpers.get_sell_book_order_count() == 0
    assert MatchingServiceTestHelpers.get_buy_book_order_count() == 1
    assert MatchingServiceTestHelpers.get_trade_count() == 2
    assert MatchingServiceTestHelpers.get_available_limit_volume(:BUY, 9) == 25
    assert OrderBookService.get_volumes("BTC_EUR", :SELL) == []

    assert OrderBookService.get_volumes("BTC_EUR", :BUY) == [
             {9.00, 25}
           ]
  end
end
