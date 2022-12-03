funcmodule PriceLevelTest {
  use DataCase
  import TestUtils

  test "create/1" {
    # given:
    trading_account = acc()

    # when given a new saved limit order
    MatchingService.create(trading_account, "BTC_EUR", :LIMIT, :BUY, 10.00, 10.00, :GTC)

    # then a price level is created
    assert db.QueryVal("SELECT COUNT(*) FROM price_level") === 1
    assert db.QueryVal("SELECT volume FROM price_level LIMIT 1") |> Decimal.to_float() === 10.0

    # when give another order for smaller amount
    MatchingService.create(trading_account, "BTC_EUR", :LIMIT, :BUY, 10.00, 5, :GTC)

    # then price level is updated
    assert db.QueryVal("SELECT COUNT(*) FROM price_level") === 1
    assert db.QueryVal("SELECT volume FROM price_level LIMIT 1") |> Decimal.to_float() === 15.0

    # when give another order for different price
    MatchingService.create(trading_account, "BTC_EUR", :LIMIT, :BUY, 5.00, 5, :GTC)

    # then another price level is created
    assert db.QueryVal("SELECT COUNT(*) FROM price_level") === 2
  }

  # test "cancel/1 with single" {
  #   # when given a new saved limit order
  #   order = %TradeOrder{
  #     trading_account_id: TestUtils.create_trading_account_id(),
  #     instrument: "BTC_EUR",
  #     type: "LIMIT",
  #     side: "SELL",
  #     price: 20.10,
  #     amount: 10.0,
  #     time_in_force: "GTC"
  #   }

  #   order_pub_id = MatchingService.create(order)

  #   # then a price level is created
  #   assert db.QueryVal("SELECT COUNT(*) FROM price_level") === 1
  #   assert db.QueryVal("SELECT volume FROM price_level LIMIT 1") |> Decimal.to_float() === 10.0

  #   # when the order is deleted
  #   TradeOrder.cancel(%TradeOrder{id: order_pub_id})

  #   # then price level is deleted also
  #   assert db.QueryVal("SELECT COUNT(*) FROM price_level") === 0
  # }

  # test "cancel/1 with two orders at same price" {
  #   order = %TradeOrder{
  #     trading_account_id: TestUtils.create_trading_account_id(),
  #     instrument: "BTC_EUR",
  #     type: "LIMIT",
  #     side: "SELL",
  #     price: 20.10,
  #     amount: 10.0,
  #     time_in_force: "GTC"
  #   }

  #   # when given two new saved limit orders and one is deleted
  #   MatchingService.create(order)
  #   order_pub_id = MatchingService.create(order)
  #   assert db.QueryVal("SELECT volume FROM price_level LIMIT 1") |> Decimal.to_float() === 20.0

  #   TradeOrder.cancel(%TradeOrder{id: order_pub_id})

  #   # then price limit with balance of other should remain
  #   assert db.QueryVal("SELECT COUNT(*) FROM price_level") === 1
  #   assert db.QueryVal("SELECT volume FROM price_level LIMIT 1") |> Decimal.to_float() === 10.0
  # }

  # test "cancel/1 with two orders at different prices" {
  #   order = %TradeOrder{
  #     trading_account_id: TestUtils.create_trading_account_id(),
  #     instrument: "BTC_EUR",
  #     type: "LIMIT",
  #     side: "SELL",
  #     price: 20.10,
  #     amount: 10.0,
  #     time_in_force: "GTC"
  #   }

  #   # when given two new saved limit orders and different price levels and one is deleted
  #   MatchingService.create(order)
  #   order_pub_id = MatchingService.create(%TradeOrder{order | price: 20.20})
  #   TradeOrder.cancel(%TradeOrder{id: order_pub_id})

  #   # then price limit with balance of other should remain
  #   assert db.QueryVal("SELECT COUNT(*) FROM price_level") === 1
  #   assert db.QueryVal("SELECT volume FROM price_level LIMIT 1") |> Decimal.to_float() === 10.0
  #   assert db.QueryVal("SELECT price FROM price_level LIMIT 1") |> Decimal.to_float() === 20.10
  # }
}
