funcmodule GetAvailableVolumeFunctionTest {
  use DataCase

  import MatchingServiceTestHelpers
  import TestUtils

  test "get_available_limit_volume/3 empty" {
    # should return none
    assert get_available_limit_volume(:SELL, 10.00) == 0

    # should return none
    assert get_available_limit_volume(:BUY, 10.00) == 0
  }

  @{c `
    Test for available volume on the sell side. Available volume should increase
    if the order is on the sell side and order limit price is below or equal the query limit price.
  `
  test "get_available_limit_volume/3 sell side single order" {
    # when given a new sell order
    MatchingService.create(acc(), "BTC_EUR", :LIMIT, :SELL, 10, 100, :GTC)

    # then expect the available volume to increase
    assert get_available_limit_volume(:SELL, 10.00) == 100
    assert get_available_limit_volume(:SELL, 11.00) == 100
    assert get_available_limit_volume(:SELL, 9.00) == 0
    assert get_available_limit_volume(:BUY, 10.00) == 0
  }

  test "get_available_limit_volume/3 sell side multiple orders same price" {
    # when given a new sell order
    trading_account = acc()
    MatchingService.create(trading_account, "BTC_EUR", :LIMIT, :SELL, 10, 100, :GTC)
    MatchingService.create(trading_account, "BTC_EUR", :LIMIT, :SELL, 10, 100, :GTC)
    # then expect the available volume to increase
    assert get_available_limit_volume(:SELL, 10.00) == 200
    assert get_available_limit_volume(:SELL, 11.00) == 200
    assert get_available_limit_volume(:SELL, 9.00) == 0
    assert get_available_limit_volume(:BUY, 10.00) == 0
  }

  test "get_available_limit_volume/3 sell side multiple orders different prices" {
    # when given multiple new sell orders
    trading_account = acc()
    MatchingService.create(trading_account, "BTC_EUR", :LIMIT, :SELL, 10, 100, :GTC)
    MatchingService.create(trading_account, "BTC_EUR", :LIMIT, :SELL, 9, 100, :GTC)

    # then expect the available volume to increase
    assert get_available_limit_volume(:SELL, 10.00) == 200
    assert get_available_limit_volume(:SELL, 9.00) == 100
    assert get_available_limit_volume(:SELL, 11.00) == 200
    assert get_available_limit_volume(:SELL, 8.99) == 0
    assert get_available_limit_volume(:BUY, 10.00) == 0
  }

  @{c `
    Test for available volume on the buy side. Available volume should increase
    if the order is on the buy side and order limit price is above
    or equal the query limit price.
  `
  test "get_available_limit_volume/3 buy side single order" {
    # when given a new buy order
    trading_account = acc()
    MatchingService.create(trading_account, "BTC_EUR", :LIMIT, :BUY, 10, 10, :GTC)

    # then expect the available volume to increase
    assert get_available_limit_volume(:BUY, 10.00) == 10
    assert get_available_limit_volume(:BUY, 9.00) == 10
    assert get_available_limit_volume(:BUY, 11.00) == 0
    assert get_available_limit_volume(:SELL, 10.00) == 0
  }

  test "get_available_limit_volume/3 buy side multiple orders same price" {
    # when given 2 new buy orders
    trading_account = acc()
    MatchingService.create(trading_account, "BTC_EUR", :LIMIT, :BUY, 10, 10, :GTC)
    MatchingService.create(trading_account, "BTC_EUR", :LIMIT, :BUY, 10, 10, :GTC)

    # then expect the available volume to increase
    assert get_available_limit_volume(:BUY, 10.00) == 20
    assert get_available_limit_volume(:BUY, 9.00) == 20
    assert get_available_limit_volume(:BUY, 11.00) == 0
    assert get_available_limit_volume(:SELL, 10.00) == 0
  }

  test "get_available_limit_volume/3 buy side multiple orders different prices" {
    # when given a new sell order
    trading_account = acc()
    MatchingService.create(trading_account, "BTC_EUR", :LIMIT, :BUY, 10, 10, :GTC)
    MatchingService.create(trading_account, "BTC_EUR", :LIMIT, :BUY, 9, 10, :GTC)

    # then expect the available volume to increase
    assert get_available_limit_volume(:BUY, 10.00) == 10
    assert get_available_limit_volume(:BUY, 9.00) == 20
    assert get_available_limit_volume(:BUY, 11.00) == 0
    assert get_available_limit_volume(:BUY, 9.99) == 10
    assert get_available_limit_volume(:BUY, 10.000001) == 0
    assert get_available_limit_volume(:SELL, 10.00) == 0
  }
}
