defmodule OrderTimeInForceTest do
  use DataCase
  import TestUtils

  test "GTC" do
    # given
    trading_account = acc()
    entity = TestUtils.get_application_entity_id()

    # when: given a new order
    trade_order =
      MatchingService.create(trading_account, "BTC_EUR", :LIMIT, :SELL, 1.0, 1.0, :GTC)

    # then: it remains in the book until cancelled
    assert DB.query_val("SELECT COUNT(*) FROM book_order") == 1
    assert TradeOrder.get(trade_order).status == :OPEN

    assert PaymentAccount.find_by_application_entity_and_currency(entity, "BTC").amount_reserved
           |> Decimal.to_float() == 1
  end

  test "FOK" do
    # given
    trading_account = acc()
    trading_account2 = acc2()
    entity = TestUtils.get_application_entity_id()

    # when: given a new order that cannot be filled
    trade_order =
      MatchingService.create(trading_account, "BTC_EUR", :LIMIT, :SELL, 1.0, 1.0, :FOK)

    # then: it is rejected
    assert DB.query_val("SELECT COUNT(*) FROM book_order") == 0
    assert TradeOrder.get(trade_order).status == :REJECTED
    assert OrderBookService.get_volume_at_price("BTC_EUR", :SELL, 1.0) == 0

    assert PaymentAccount.find_by_application_entity_and_currency(entity, "BTC").amount_reserved
           |> Decimal.to_float() == 0

    # when: given a new order that cannot be filled even when other orders present
    MatchingService.create(trading_account2, "BTC_EUR", :LIMIT, :BUY, 1.0, 1.0, :GTC)

    trade_order =
      MatchingService.create(trading_account, "BTC_EUR", :LIMIT, :SELL, 1.0, 2.0, :FOK)

    # then: it is rejected
    assert TradeOrder.get(trade_order).status == :REJECTED

    assert PaymentAccount.find_by_application_entity_and_currency(entity, "BTC").amount_reserved
           |> Decimal.to_float() == 0

    # when: added another market order that can fill
    MatchingService.create(trading_account2, "BTC_EUR", :MARKET, :BUY, 2.0, :GTC)

    trade_order =
      MatchingService.create(trading_account, "BTC_EUR", :LIMIT, :SELL, 1.0, 2.0, :FOK)

    # then: it is not reject
    assert TradeOrder.get(trade_order).status == :FILLED
    assert OrderBookService.get_volume_at_price("BTC_EUR", :SELL, 1.0) == 0

    # assert PaymentAccount.find_by_application_entity_and_currency(entity, "BTC").amount_reserved
    #        |> Decimal.to_float() == 0
  end

  test "IOC" do
    # given
    trading_account = acc()
    trading_account2 = acc2()
    entity = TestUtils.get_application_entity_id()

    # when: given a new order that cannot be filled
    trade_order =
      MatchingService.create(trading_account, "BTC_EUR", :LIMIT, :SELL, 1.0, 1.0, :IOC)

    # then: it is rejected
    assert DB.query_val("SELECT COUNT(*) FROM book_order") == 0
    assert TradeOrder.get(trade_order).status == :REJECTED
    assert OrderBookService.get_volume_at_price("BTC_EUR", :SELL, 1.0) == 0

    assert PaymentAccount.find_by_application_entity_and_currency(entity, "BTC").amount_reserved
           |> Decimal.to_float() == 0

    # when: given a new order that can only be partially filled by a standing order in the order book
    MatchingService.create(trading_account2, "BTC_EUR", :LIMIT, :BUY, 1.0, 1, :GTC)
    trade_order = MatchingService.create(trading_account, "BTC_EUR", :LIMIT, :SELL, 1.0, 2, :IOC)

    # then: it is partially rejected
    assert MatchingServiceTestHelpers.get_trade_count() == 1
    assert DB.query_val("SELECT COUNT(*) FROM book_order") == 0
    assert TradeOrder.get(trade_order).status == :PARTIALLY_REJECTED
    assert OrderBookService.get_volume_at_price("BTC_EUR", :SELL, 1.0) == 0

    assert PaymentAccount.find_by_application_entity_and_currency(entity, "BTC").amount_reserved
           |> Decimal.to_float() == 0
  end
end
