funcmodule MatchingServiceBalanceTest {
  use DataCase

  test "process/1 limit sell order save with reserved balance" {
    # given:
    application_entity_id = TestUtils.create_client()
    PaymentAccount.create(application_entity_id, "BTC")
    trading_account_id = TradingAccount.find_by_application_entity_id(application_entity_id).id

    payment_account =
      PaymentAccount.find_by_application_entity_and_currency(application_entity_id, "BTC")

    assert Decimal.equal?(payment_account.amount, 0)
    assert Decimal.equal?(payment_account.amount_reserved, 0)
    assert Decimal.equal?(payment_account.amount_available, 0)

    # when:
    PaymentService.deposit(application_entity_id, 1000, "BTC", "Test", "Test")

    # then:
    payment_account =
      PaymentAccount.find_by_application_entity_and_currency(application_entity_id, "BTC")

    assert Decimal.equal?(payment_account.amount, 1000)
    assert Decimal.equal?(payment_account.amount_reserved, 0)
    assert Decimal.equal?(payment_account.amount_available, 1000)

    # when: a limit order is sent to an empty matching unit
    MatchingService.create(trading_account_id, "BTC_EUR", :LIMIT, :SELL, 10, 100, :GTC)

    # then: reseved balance to the account should be increased
    payment_account =
      PaymentAccount.find_by_application_entity_and_currency(application_entity_id, "BTC")

    assert Decimal.equal?(payment_account.amount, 1000)
    assert Decimal.equal?(payment_account.amount_reserved, 100)
    assert Decimal.equal?(payment_account.amount_available, 900)

    # when: another limit order is sent to an empty matching unit
    MatchingService.create(trading_account_id, "BTC_EUR", :LIMIT, :SELL, 10, 100, :GTC)

    # then: reseved balance to the account should be increased
    payment_account =
      PaymentAccount.find_by_application_entity_and_currency(application_entity_id, "BTC")

    assert Decimal.equal?(payment_account.amount, 1000)
    assert Decimal.equal?(payment_account.amount_reserved, 200)
    assert Decimal.equal?(payment_account.amount_available, 800)
  }

  test "process/1 limit buy order save with reserved balance" {
    # given:
    application_entity_id = TestUtils.create_client()

    payment_account_id =
      PaymentAccount.find_by_application_entity_and_currency(application_entity_id, "EUR").id

    trading_account_id = TradingAccount.find_by_application_entity_id(application_entity_id).id

    payment_account = PaymentAccount.get(payment_account_id)

    assert Decimal.equal?(payment_account.amount, 0)
    assert Decimal.equal?(payment_account.amount_reserved, 0)
    assert Decimal.equal?(payment_account.amount_available, 0)

    # when:
    PaymentService.deposit(application_entity_id, 1000, "EUR", "Test", "Test")

    # then:
    payment_account = PaymentAccount.get(payment_account_id)

    assert Decimal.equal?(payment_account.amount, 1000)
    assert Decimal.equal?(payment_account.amount_reserved, 0)
    assert Decimal.equal?(payment_account.amount_available, 1000)

    # when: a limit order is sent to an empty matching unit
    MatchingService.create(trading_account_id, "BTC_EUR", :LIMIT, :BUY, 10, 10, :GTC)

    # then: reseved balance to the account should be increased
    payment_account = PaymentAccount.get(payment_account_id)

    assert Decimal.equal?(payment_account.amount, 1000)
    assert Decimal.equal?(payment_account.amount_reserved, 100)
    assert Decimal.equal?(payment_account.amount_available, 900)

    # when: another limit order is sent to an empty matching unit
    MatchingService.create(trading_account_id, "BTC_EUR", :LIMIT, :BUY, 10, 10, :GTC)

    # then: reseved balance to the account should be increased
    payment_account = PaymentAccount.get(payment_account_id)

    assert Decimal.equal?(payment_account.amount, 1000)
    assert Decimal.equal?(payment_account.amount_reserved, 200)
    assert Decimal.equal?(payment_account.amount_available, 800)
  }

  test "process/1 limit sell order against matching buy order with funds transfer" {
    # given:
    # -- a seller
    trading_account = TestUtils.acc()
    application_entity_id = TradingAccount.get(trading_account).application_entity_id

    seller_debit_account =
      PaymentAccount.find_by_application_entity_and_currency(application_entity_id, "BTC")

    seller_credit_account =
      PaymentAccount.find_by_application_entity_and_currency(application_entity_id, "EUR")

    assert seller_debit_account.amount |> Decimal.to_float() == 1000
    assert seller_credit_account.amount |> Decimal.to_float() == 1000

    # -- buyer
    trading_account2 = TestUtils.acc2()
    application_entity_id2 = TradingAccount.get(trading_account2).application_entity_id

    buyer_debit_account =
      PaymentAccount.find_by_application_entity_and_currency(application_entity_id2, "EUR")

    buyer_credit_account =
      PaymentAccount.find_by_application_entity_and_currency(application_entity_id2, "BTC")

    assert buyer_debit_account.amount |> Decimal.to_float() == 1000

    # when: a trade is executed
    MatchingService.create(trading_account, "BTC_EUR", :LIMIT, :SELL, 10, 10, :GTC)
    assert MatchingServiceTestHelpers.get_trade_count() == 0
    assert PaymentAccount.get(seller_debit_account.id).amount_reserved |> Decimal.to_float() == 10
    MatchingService.create(trading_account2, "BTC_EUR", :LIMIT, :BUY, 10, 10, :GTC)
    assert MatchingServiceTestHelpers.get_trade_count() == 1

    # then: 4 payments are executed in addition to 4 deposits
    assert db.QueryVal("SELECT COUNT(*) FROM payment") == 8

    # and: seller debit balance should decrease for by 10  BTC of seller but reserved balance should be released
    assert PaymentAccount.get(seller_debit_account.id).amount |> Decimal.to_float() == 990
    assert PaymentAccount.get(seller_debit_account.id).amount_reserved |> Decimal.to_float() == 0

    # and: seller credit balance should increase by 100 EUR
    assert PaymentAccount.get(seller_credit_account.id).amount |> Decimal.to_float() == 1100
    assert PaymentAccount.get(seller_credit_account.id).amount_reserved |> Decimal.to_float() == 0

    # and: buy debit balance should decrease by 100 EUR but reserve balance should be released
    assert PaymentAccount.get(buyer_debit_account.id).amount |> Decimal.to_float() == 900
    assert PaymentAccount.get(buyer_debit_account.id).amount_reserved |> Decimal.to_float() == 0

    # and: buy debit balance should increase by 10 BTC
    assert PaymentAccount.get(buyer_credit_account.id).amount |> Decimal.to_float() == 1010
    assert PaymentAccount.get(buyer_credit_account.id).amount_reserved |> Decimal.to_float() == 0
  }
}
