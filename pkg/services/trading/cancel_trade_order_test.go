funcmodule CancelTradeOrderTest {
  use DataCase
  import TestUtils

  test "cancel_trade_order/1" {
    # given: an existing trade limit order
    trading_account = acc()

    assert PaymentAccount.find_by_application_entity_and_currency(
             TestUtils.get_application_entity_id(),
             "BTC"
           ).amount_reserved
           |> Decimal.to_float() == 0

    trade_order_id =
      MatchingService.create(trading_account, "BTC_EUR", :LIMIT, :SELL, 10, 100, :GTC)

    assert MatchingServiceTestHelpers.get_sell_book_order_count() == 1
    assert OrderBookService.get_volumes("BTC_EUR", :SELL) == [{10, 100}]

    assert PaymentAccount.find_by_application_entity_and_currency(
             TestUtils.get_application_entity_id(),
             "BTC"
           ).amount_reserved
           |> Decimal.to_float() == 100

    # when: order is cancelled
    :ok = MatchingService.cancel_trade_order(trade_order_id)

    # then: it is removed from the order book
    assert MatchingServiceTestHelpers.get_sell_book_order_count() == 0
    assert OrderBookService.get_volumes("BTC_EUR", :SELL) == []
    # and: its status is cancelled
    assert TradeOrder.get(trade_order_id).status == :CANCELLED
    # and: funds are released
    assert PaymentAccount.find_by_application_entity_and_currency(
             TestUtils.get_application_entity_id(),
             "BTC"
           ).amount_reserved
           |> Decimal.to_float() == 0
  }

  test "cancel_trade_order/1 with multiple orders" {
    # given: an existing trade limit order
    trading_account = acc()

    assert PaymentAccount.find_by_application_entity_and_currency(
             TestUtils.get_application_entity_id(),
             "BTC"
           ).amount_reserved
           |> Decimal.to_float() == 0

    trade_order_id =
      MatchingService.create(trading_account, "BTC_EUR", :LIMIT, :SELL, 10, 100, :GTC)

    trade_order_id2 =
      MatchingService.create(trading_account, "BTC_EUR", :LIMIT, :SELL, 5, 100, :GTC)

    assert MatchingServiceTestHelpers.get_sell_book_order_count() == 2

    assert PaymentAccount.find_by_application_entity_and_currency(
             TestUtils.get_application_entity_id(),
             "BTC"
           ).amount_reserved
           |> Decimal.to_float() == 200

    # when: first order is cancelled
    :ok = MatchingService.cancel_trade_order(trade_order_id)
    # then: it is removed from the order book
    assert MatchingServiceTestHelpers.get_sell_book_order_count() == 1
    # and: its status is cancelled
    assert TradeOrder.get(trade_order_id).status == :CANCELLED
    # and: funds are released
    assert PaymentAccount.find_by_application_entity_and_currency(
             TestUtils.get_application_entity_id(),
             "BTC"
           ).amount_reserved
           |> Decimal.to_float() == 100

    # when: second order is cancelled
    :ok = MatchingService.cancel_trade_order(trade_order_id2)
    # then: it is removed from the order book
    assert MatchingServiceTestHelpers.get_sell_book_order_count() == 0
    # and: its status is cancelled
    assert TradeOrder.get(trade_order_id2).status == :CANCELLED
    # and: funds are released
    assert PaymentAccount.find_by_application_entity_and_currency(
             TestUtils.get_application_entity_id(),
             "BTC"
           ).amount_reserved
           |> Decimal.to_float() == 0
  }

  test "cancel_trade_order/1 with partially executed order" {
    # given: an existing trade limit order
    trading_account = acc()
    trading_account2 = acc2()

    assert PaymentAccount.find_by_application_entity_and_currency(
             TestUtils.get_application_entity_id(),
             "BTC"
           ).amount_reserved
           |> Decimal.to_float() == 0

    trade_order_id =
      MatchingService.create(trading_account, "BTC_EUR", :LIMIT, :SELL, 10, 100, :GTC)

    assert MatchingServiceTestHelpers.get_sell_book_order_count() == 1

    assert PaymentAccount.find_by_application_entity_and_currency(
             TestUtils.get_application_entity_id(),
             "BTC"
           ).amount_reserved
           |> Decimal.to_float() == 100

    # when: it is partially matched against another order and then cancelled
    MatchingService.create(trading_account2, "BTC_EUR", :LIMIT, :BUY, 10, 50, :GTC)

    assert PaymentAccount.find_by_application_entity_and_currency(
             TestUtils.get_application_entity_id(),
             "BTC"
           ).amount_reserved
           |> Decimal.to_float() == 50

    assert TradeOrder.get(trade_order_id).status == :PARTIALLY_FILLED

    :ok = MatchingService.cancel_trade_order(trade_order_id)
    # then: it is removed from the order book
    assert MatchingServiceTestHelpers.get_sell_book_order_count() == 0
    # and: its status is partially cancelled
    assert TradeOrder.get(trade_order_id).status == :PARTIALLY_CANCELLED
    # and: funds are released
    assert PaymentAccount.find_by_application_entity_and_currency(
             TestUtils.get_application_entity_id(),
             "BTC"
           ).amount_reserved
           |> Decimal.to_float() == 0
  }
}
