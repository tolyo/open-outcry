defmodule MatchingServiceAmountConstraintTest do
  use DataCase

  @moduledoc """
    These tests apply to balance constraints of order matching as we never want to allow
    order into the order book that do no have sufficient leverage.
    For limit orders, we ensure that the seller has sufficient amount in base currency or instument
    and that the buys has sufficient amount in quote currency which is limit price times base currency or instrument.

    For martker order we ensure that a seller instrument amount is valid or buyer quote currency amount is valid.
  """

  test "process/1 limit sell order save with insufficient funds" do
    # given:
    application_entity_id = TestUtils.create_client()
    PaymentAccount.create(application_entity_id, "BTC")
    PaymentService.deposit(application_entity_id, 100, "BTC", "test", "Test")
    trading_account_id = TradingAccount.find_by_application_entity_id(application_entity_id).id
    MatchingService.create(trading_account_id, "BTC_EUR", :LIMIT, :SELL, 10, 100, :GTC)
    PaymentService.deposit(application_entity_id, 100, "BTC", "test", "Test")
    # when: a limit order is sent with insufficient funds

    # then: exception is raised
    assert_raise Postgrex.Error, fn ->
      # this should be causing a failure
      MatchingService.create(trading_account_id, "BTC_EUR", :LIMIT, :SELL, 10, 101, :GTC)
    end
  end

  test "process/1 limit buy order save with insufficient funds" do
    # given:
    application_entity_id = TestUtils.create_client()
    PaymentService.deposit(application_entity_id, 100, "EUR", "test", "Test")
    trading_account_id = TradingAccount.find_by_application_entity_id(application_entity_id).id
    MatchingService.create(trading_account_id, "BTC_EUR", :LIMIT, :BUY, 10, 10, :GTC)
    PaymentService.deposit(application_entity_id, 100, "EUR", "test", "Test")
    # when: a limit order is sent with insufficient funds

    # then: exception is raised
    assert_raise Postgrex.Error, fn ->
      MatchingService.create(trading_account_id, "BTC_EUR", :LIMIT, :BUY, 10, 11, :GTC)
    end
  end

  test "process/1 market sell order save with insufficient funds" do
    # given:
    application_entity_id = TestUtils.create_client()
    PaymentAccount.create(application_entity_id, "BTC")
    PaymentService.deposit(application_entity_id, 100, "BTC", "test", "Test")
    trading_account_id = TradingAccount.find_by_application_entity_id(application_entity_id).id
    MatchingService.create(trading_account_id, "BTC_EUR", :MARKET, :SELL, 100, :GTC)
    PaymentService.deposit(application_entity_id, 100, "BTC", "test", "Test")
    # when: a market order is sent with insufficient funds

    # then: exception is raised
    assert_raise Postgrex.Error, fn ->
      MatchingService.create(trading_account_id, "BTC_EUR", :MARKET, :SELL, 101, :GTC)
    end
  end

  test "process/1 market buy order save with insufficient funds" do
    # given:
    application_entity_id = TestUtils.create_client()
    PaymentService.deposit(application_entity_id, 100, "EUR", "test", "Test")
    trading_account_id = TradingAccount.find_by_application_entity_id(application_entity_id).id
    MatchingService.create(trading_account_id, "BTC_EUR", :MARKET, :BUY, 100, :GTC)
    PaymentService.deposit(application_entity_id, 100, "EUR", "test", "Test")
    # when: a market order is sent with insufficient funds

    # then: exception is raised
    assert_raise Postgrex.Error, fn ->
      MatchingService.create(trading_account_id, "BTC_EUR", :MARKET, :BUY, 101, :GTC)
    end
  end
end
