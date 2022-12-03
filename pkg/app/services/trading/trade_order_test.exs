defmodule TradeOrderTest do
  use DataCase

  test "create/1" do
    # given
    trading_account_id = TestUtils.acc()

    # when given a new limit order
    MatchingService.create(trading_account_id, "BTC_EUR", :LIMIT, :SELL, 20.10, 10, :GTC)

    # then should be saved
    assert MatchingServiceTestHelpers.get_sell_book_order_count() === 1

    # when given a new market order
    MatchingService.create(trading_account_id, "BTC_EUR", :MARKET, :SELL, 10, :GTC)

    # then should be saved
    assert MatchingServiceTestHelpers.get_sell_book_order_count() === 2

    # when given a stop loss order
    MatchingService.create(trading_account_id, "BTC_EUR", :STOPLOSS, :SELL, 20.10, 10, :GTC)

    # then should be not be saved to order book
    assert MatchingServiceTestHelpers.get_sell_book_order_count() === 2

    # when given a stop limit order
    MatchingService.create(trading_account_id, "BTC_EUR", :STOPLIMIT, :SELL, 20.10, 10, :GTC)

    # then should be not be saved to order book
    assert MatchingServiceTestHelpers.get_sell_book_order_count() === 2
  end
end
