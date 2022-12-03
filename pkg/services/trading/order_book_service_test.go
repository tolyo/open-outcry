funcmodule OrderBookServiceTest {
  use DataCase

  test "get_volume_at_price/3" {
    # when: a single sell limit order is added to the order book
    MatchingService.create(TestUtils.acc(), "BTC_EUR", :LIMIT, :SELL, 10.6, 100, :GTC)

    # then:
    assert OrderBookService.get_volume_at_price("BTC_EUR", :SELL, 10.6) |> Decimal.to_float() ==
             100

    assert db.QueryVal(`
           SELECT SUM(volume)
             FROM price_level
             WHERE side = 'SELL'
               AND instrument_id = (SELECT id FROM instrument WHERE name = 'BTC_EUR')
               AND price =  10.6
           `)
           |> Decimal.to_float() == 100

    # when: a single buy limit order is added to the order book
    MatchingService.create(TestUtils.acc2(), "BTC_EUR", :LIMIT, :BUY, 9.5, 100, :GTC)

    # then:
    assert OrderBookService.get_volume_at_price("BTC_EUR", :BUY, 9.5) |> Decimal.to_float() ==
             100

    assert db.QueryVal(`
             SELECT SUM(volume)
               FROM price_level
               WHERE side = 'BUY'
                 AND instrument_id = (SELECT id FROM instrument WHERE name = 'BTC_EUR')
                 AND price =  9.5
           `)
           |> Decimal.to_float() == 100
  }

  test "get_volumes/2 sell side" {
    trading_account = TestUtils.acc()

    # when
    MatchingService.create(trading_account, "BTC_EUR", :LIMIT, :SELL, 10.7, 100, :GTC)
    MatchingService.create(trading_account, "BTC_EUR", :LIMIT, :SELL, 10.6, 100, :GTC)
    MatchingService.create(trading_account, "BTC_EUR", :LIMIT, :SELL, 10.7, 100, :GTC)
    MatchingService.create(trading_account, "BTC_EUR", :LIMIT, :SELL, 10.4, 100, :GTC)

    # then should be sorted with cheapest orders first
    assert OrderBookService.get_volumes("BTC_EUR", :SELL) == [
             {10.4, 100},
             {10.6, 100},
             {10.7, 200}
           ]
  }

  test "get_volumes/2 buy side" {
    trading_account = TestUtils.acc()

    # when
    MatchingService.create(trading_account, "BTC_EUR", :LIMIT, :BUY, 1.7, 10, :GTC)
    MatchingService.create(trading_account, "BTC_EUR", :LIMIT, :BUY, 1.6, 10, :GTC)
    MatchingService.create(trading_account, "BTC_EUR", :LIMIT, :BUY, 1.7, 10, :GTC)
    MatchingService.create(trading_account, "BTC_EUR", :LIMIT, :BUY, 1.4, 10, :GTC)

    # then should be sorted with most expensive orders first
    assert OrderBookService.get_volumes("BTC_EUR", :BUY) == [
             {1.7, 20},
             {1.6, 10},
             {1.4, 10}
           ]
  }
}
