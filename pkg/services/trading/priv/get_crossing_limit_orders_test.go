funcmodule GetCrossingLimitOrdersTest {
  use DataCase

  import TestUtils

  test "get_crossing_limit_orders/3 test price SELL side" {
    # given:
    trading_account = acc()

    # should return none
    assert MatchingServiceTestHelpers.get_crossing_limit_orders(1, :SELL, 10.00) == []

    # when given a new order
    MatchingService.create(trading_account, "BTC_EUR", :LIMIT, :SELL, 10.00, 1, :GTC)

    # then count should increase
    assert MatchingServiceTestHelpers.get_crossing_limit_orders(1, :SELL, 10.00) |> length == 1

    # # when given another new order
    MatchingService.create(trading_account, "BTC_EUR", :LIMIT, :SELL, 10.00, 1, :GTC)

    # then count should increase
    assert MatchingServiceTestHelpers.get_crossing_limit_orders(1, :SELL, 10.00) |> length == 2

    # when given another new order with crossing price
    MatchingService.create(trading_account, "BTC_EUR", :LIMIT, :SELL, 9.00, 1, :GTC)

    # then count should increase
    assert MatchingServiceTestHelpers.get_crossing_limit_orders(1, :SELL, 10.00) |> length == 3

    # when given another new order non crossing price
    MatchingService.create(trading_account, "BTC_EUR", :LIMIT, :SELL, 19.00, 1, :GTC)

    # then count should not change
    assert MatchingServiceTestHelpers.get_crossing_limit_orders(1, :SELL, 10.00) |> length == 3

    # when given another new order with crossing price for buy side
    trading_account2 = TestUtils.acc2()
    MatchingService.create(trading_account2, "BTC_EUR", :LIMIT, :BUY, 10.00, 1, :GTC)

    # then count should decrease
    assert MatchingServiceTestHelpers.get_crossing_limit_orders(1, :SELL, 10.00) |> length == 2

    # when given another new order non crossing price
    MatchingService.create(trading_account, "BTC_EUR", :LIMIT, :SELL, 10.01, 1, :GTC)

    # then count should increase
    assert MatchingServiceTestHelpers.get_crossing_limit_orders(1, :SELL, 10.00) |> length == 2

    # when given another new order with crossing price
    MatchingService.create(
      trading_account,
      "BTC_EUR",
      :LIMIT,
      :SELL,
      10.000 - 0.000001,
      10.00,
      :GTC
    )

    # then count should increase because buy side is emtpy
    assert MatchingServiceTestHelpers.get_crossing_limit_orders(1, :SELL, 10.00) |> length == 3
  }

  # test "get_crossing_limit_orders/3 test price BUY side" {
  #   # should return none
  #   assert MatchingServiceTestHelpers.get_crossing_limit_orders(1, :BUY, 10.00) == []

  #   # when given a new order
  #   order = %TradeOrder{
  #     trading_account_id: TestUtils.create_trading_account_id(),
  #     instrument: "BTC_EUR",
  #     type: "LIMIT",
  #     side: :BUY,
  #     price: 10.00,
  #     amount: 10.0,
  #     time_in_force: "GTC"
  #   }

  #   MatchingService.create(order)

  #   # then count should increase
  #   assert MatchingServiceTestHelpers.get_crossing_limit_orders(1, :BUY, 10.00) |> length == 1

  #   # when given another new order
  #   MatchingService.create(order)

  #   # then count should increase
  #   assert MatchingServiceTestHelpers.get_crossing_limit_orders(1, :BUY, 10.00) |> length == 2

  #   # when given another new order with crossing price
  #   MatchingService.create(%TradeOrder{order | price: 11.00})

  #   # then count should increase
  #   assert MatchingServiceTestHelpers.get_crossing_limit_orders(1, :BUY, 10.00) |> length == 3

  #   # when given another new order non crossing price
  #   MatchingService.create(%TradeOrder{order | price: 9.00})

  #   # then count should not change
  #   assert MatchingServiceTestHelpers.get_crossing_limit_orders(1, :BUY, 10.00) |> length == 3

  #   # when given another new order with crossing price for buy side
  #   MatchingService.create(%TradeOrder{order | price: 10.00, side: :SELL})

  #   # then count should not change
  #   assert MatchingServiceTestHelpers.get_crossing_limit_orders(1, :BUY, 10.00) |> length == 3

  #   # when given another new order non crossing price
  #   MatchingService.create(%TradeOrder{order | price: 9.99999})

  #   # then count should not change
  #   assert MatchingServiceTestHelpers.get_crossing_limit_orders(1, :BUY, 10.00) |> length == 3

  #   # when given another new order with crossing price
  #   MatchingService.create(%TradeOrder{order | price: 10.000001})

  #   # then count should increase
  #   assert MatchingServiceTestHelpers.get_crossing_limit_orders(1, :BUY, 10.00) |> length == 4
  # }

  # @{c `
  #   Test for ordering and limit of orders on the sell side. TradeOrders should be ordered by date (FIFO)
  #   and by price where where lowest prices are selected first and up to but not above the limit.
  # `
  # test "get_crossing_limit_orders/3 test ordering SELL side" {
  #   # should return none
  #   assert MatchingServiceTestHelpers.get_crossing_limit_orders(1, :SELL, 10.00) == []

  #   # when given a new order
  #   order = %TradeOrder{
  #     trading_account_id: TestUtils.create_trading_account_id(),
  #     instrument: "BTC_EUR",
  #     type: "LIMIT",
  #     side: :SELL,
  #     price: 10.00,
  #     amount: 10.0,
  #     time_in_force: "GTC"
  #   }

  #   pub_id1 = MatchingService.create(order)

  #   # then should first item in list
  #   %TradeOrder{id: pub_id} =
  #     MatchingServiceTestHelpers.get_crossing_limit_orders(1, :SELL, 10.00)
  #     |> Enum.fetch!(0)
  #     |> from_db_tuple

  #   assert pub_id1 === pub_id

  #   # when given another new order
  #   pub_id2 = MatchingService.create(order)
  #   order_list = MatchingServiceTestHelpers.get_crossing_limit_orders(1, :SELL, 10.00)

  #   # then it should be second in list
  #   %TradeOrder{id: pub_id} = order_list |> Enum.fetch!(0) |> from_db_tuple
  #   assert pub_id1 === pub_id
  #   %TradeOrder{id: pub_id} = order_list |> Enum.fetch!(1) |> from_db_tuple
  #   assert pub_id2 === pub_id

  #   # when given another new order with crossing price
  #   pub_id3 = MatchingService.create(%TradeOrder{order | price: 9.00})
  #   order_list = MatchingServiceTestHelpers.get_crossing_limit_orders(1, :SELL, 10.00)

  #   # then it should be first in list
  #   %TradeOrder{id: pub_id} = order_list |> Enum.fetch!(0) |> from_db_tuple
  #   assert pub_id3 === pub_id

  #   # when given another new order with crossing price
  #   pub_id4 = MatchingService.create(%TradeOrder{order | price: 9.50})
  #   order_list = MatchingServiceTestHelpers.get_crossing_limit_orders(1, :SELL, 10.00)

  #   # then it should be 2rd in list
  #   %TradeOrder{id: pub_id} = order_list |> Enum.fetch!(1) |> from_db_tuple
  #   assert pub_id4 === pub_id
  #   %TradeOrder{id: pub_id} = order_list |> Enum.fetch!(0) |> from_db_tuple
  #   assert pub_id3 === pub_id

  #   # when given another new order with crossing price
  #   pub_id5 = MatchingService.create(%TradeOrder{order | price: 9.25})
  #   order_list = MatchingServiceTestHelpers.get_crossing_limit_orders(1, :SELL, 10.00)

  #   # then it should be 2rd in list
  #   %TradeOrder{id: pub_id} = order_list |> Enum.fetch!(1) |> from_db_tuple
  #   assert pub_id5 === pub_id
  #   %TradeOrder{id: pub_id} = order_list |> Enum.fetch!(2) |> from_db_tuple
  #   assert pub_id4 === pub_id
  #   %TradeOrder{id: pub_id} = order_list |> Enum.fetch!(0) |> from_db_tuple
  #   assert pub_id3 === pub_id
  # }

  # @{c `
  #   Test for ordering of orders on the buy side. TradeOrders should be ordered by date (FIFO)
  #   and by price where where highest prices are selected first {wn to but now below limit price.
  # `
  # test "get_crossing_limit_orders/3 test ordering BUY side" {
  #   # should return none
  #   assert MatchingServiceTestHelpers.get_crossing_limit_orders(1, :BUY, 10.00) == []

  #   # when given a new order
  #   order = %TradeOrder{
  #     trading_account_id: TestUtils.create_trading_account_id(),
  #     instrument: "BTC_EUR",
  #     type: "LIMIT",
  #     side: :BUY,
  #     price: 10.00,
  #     amount: 10.0,
  #     time_in_force: "GTC"
  #   }

  #   pub_id1 = MatchingService.create(order)

  #   # then should first item in list
  #   %TradeOrder{id: pub_id} =
  #     MatchingServiceTestHelpers.get_crossing_limit_orders(1, :BUY, 10.00)
  #     |> Enum.fetch!(0)
  #     |> from_db_tuple

  #   assert pub_id1 === pub_id
  #   assert MatchingServiceTestHelpers.get_crossing_limit_orders(1, :BUY, 11.00) |> length == 0
  #   assert MatchingServiceTestHelpers.get_crossing_limit_orders(1, :BUY, 9.00) |> length == 1

  #   # when given another limit order
  #   pub_id2 = MatchingService.create(order)

  #   assert MatchingServiceTestHelpers.get_crossing_limit_orders(1, :BUY, 10.00) |> length == 2
  #   assert MatchingServiceTestHelpers.get_crossing_limit_orders(1, :BUY, 11.00) |> length == 0
  #   assert MatchingServiceTestHelpers.get_crossing_limit_orders(1, :BUY, 8.00) |> length == 2

  #   order_list = MatchingServiceTestHelpers.get_crossing_limit_orders(1, :BUY, 10.00)

  #   # then it should be second in list
  #   %TradeOrder{id: pub_id} = order_list |> Enum.fetch!(0) |> from_db_tuple
  #   assert pub_id1 === pub_id
  #   %TradeOrder{id: pub_id} = order_list |> Enum.fetch!(1) |> from_db_tuple
  #   assert pub_id2 === pub_id

  #   # when given another new order with a higher crossing price
  #   pub_id3 = MatchingService.create(%TradeOrder{order | price: 11.00})

  #   assert MatchingServiceTestHelpers.get_crossing_limit_orders(1, :BUY, 10.00) |> length == 3
  #   assert MatchingServiceTestHelpers.get_crossing_limit_orders(1, :BUY, 12.00) |> length == 0
  #   assert MatchingServiceTestHelpers.get_crossing_limit_orders(1, :BUY, 8.00) |> length == 3

  #   order_list = MatchingServiceTestHelpers.get_crossing_limit_orders(1, :BUY, 11.00)

  #   # then it should be first in list
  #   %TradeOrder{id: pub_id} = order_list |> Enum.fetch!(0) |> from_db_tuple
  #   assert pub_id3 === pub_id

  #   # when given another new order with crossing price
  #   pub_id4 = MatchingService.create(%TradeOrder{order | price: 10.50})
  #   order_list = MatchingServiceTestHelpers.get_crossing_limit_orders(1, :BUY, 10.00)

  #   # then it should be 2rd in list
  #   %TradeOrder{id: pub_id} = order_list |> Enum.fetch!(1) |> from_db_tuple
  #   assert pub_id4 === pub_id
  #   %TradeOrder{id: pub_id} = order_list |> Enum.fetch!(0) |> from_db_tuple
  #   assert pub_id3 === pub_id

  #   # when given another new order with crossing price
  #   pub_id5 = MatchingService.create(%TradeOrder{order | price: 10.75})
  #   order_list = MatchingServiceTestHelpers.get_crossing_limit_orders(1, :BUY, 10.00)

  #   # then it should be 2rd in list
  #   %TradeOrder{id: pub_id} = order_list |> Enum.fetch!(1) |> from_db_tuple
  #   assert pub_id5 === pub_id
  #   %TradeOrder{id: pub_id} = order_list |> Enum.fetch!(2) |> from_db_tuple
  #   assert pub_id4 === pub_id
  #   %TradeOrder{id: pub_id} = order_list |> Enum.fetch!(0) |> from_db_tuple
  #   assert pub_id3 === pub_id
  # }
}
