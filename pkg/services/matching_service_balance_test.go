package services

//funcmodule MatchingServiceBalanceTest {
//  use DataCase
//
//  test "process/1 limit sell order save with reserved balance" {
// given:
//    appEntityId := TestUtils.create_client()
//    PaymentAccount.create(appEntityId, "BTC")
//    tradingAccountId := TradingAccount.find_by_appEntityId(appEntityId).id
//
//    payment_account =
//      FindPaymentAccountByAppEntityIdAndCurrencyName(appEntityId, "BTC")
//
//    assert Decimal.equal?(payment_account.amount, 0)
//    assert Decimal.equal?(payment_account.AmountReserved, 0)
//    assert Decimal.equal?(payment_account.amount_available, 0)
//
// when:
//    PaymentService.deposit(appEntityId, 1000, "BTC", "Test", "Test")
//
// then:
//    payment_account =
//      FindPaymentAccountByAppEntityIdAndCurrencyName(appEntityId, "BTC")
//
//    assert Decimal.equal?(payment_account.amount, 1000)
//    assert Decimal.equal?(payment_account.AmountReserved, 0)
//    assert Decimal.equal?(payment_account.amount_available, 1000)
//
// when: a limit order is sent to an empty matching unit
//    ProcessTradeOrder(tradingAccountId, "BTC_EUR", "LIMIT", "SELL", 10, 100, "GTC")
//
// then: reseved balance to the account should be increased
//    payment_account =
//      FindPaymentAccountByAppEntityIdAndCurrencyName(appEntityId, "BTC")
//
//    assert Decimal.equal?(payment_account.amount, 1000)
//    assert Decimal.equal?(payment_account.AmountReserved, 100)
//    assert Decimal.equal?(payment_account.amount_available, 900)
//
// when: another limit order is sent to an empty matching unit
//    ProcessTradeOrder(tradingAccountId, "BTC_EUR", "LIMIT", "SELL", 10, 100, "GTC")
//
// then: reseved balance to the account should be increased
//    payment_account =
//      FindPaymentAccountByAppEntityIdAndCurrencyName(appEntityId, "BTC")
//
//    assert Decimal.equal?(payment_account.amount, 1000)
//    assert Decimal.equal?(payment_account.AmountReserved, 200)
//    assert Decimal.equal?(payment_account.amount_available, 800)
//  }
//
//  test "process/1 limit buy order save with reserved balance" {
// given:
//    appEntityId := TestUtils.create_client()
//
//    payment_account_id =
//      FindPaymentAccountByAppEntityIdAndCurrencyName(appEntityId, "EUR").id
//
//    tradingAccountId := TradingAccount.find_by_appEntityId(appEntityId).id
//
//    payment_account = PaymentAccount.get(payment_account_id)
//
//    assert Decimal.equal?(payment_account.amount, 0)
//    assert Decimal.equal?(payment_account.AmountReserved, 0)
//    assert Decimal.equal?(payment_account.amount_available, 0)
//
// when:
//    PaymentService.deposit(appEntityId, 1000, "EUR", "Test", "Test")
//
// then:
//    payment_account = PaymentAccount.get(payment_account_id)
//
//    assert Decimal.equal?(payment_account.amount, 1000)
//    assert Decimal.equal?(payment_account.AmountReserved, 0)
//    assert Decimal.equal?(payment_account.amount_available, 1000)
//
// when: a limit order is sent to an empty matching unit
//    ProcessTradeOrder(tradingAccountId, "BTC_EUR", "LIMIT", "BUY", 10, 10, "GTC")
//
// then: reseved balance to the account should be increased
//    payment_account = PaymentAccount.get(payment_account_id)
//
//    assert Decimal.equal?(payment_account.amount, 1000)
//    assert Decimal.equal?(payment_account.AmountReserved, 100)
//    assert Decimal.equal?(payment_account.amount_available, 900)
//
// when: another limit order is sent to an empty matching unit
//    ProcessTradeOrder(tradingAccountId, "BTC_EUR", "LIMIT", "BUY", 10, 10, "GTC")
//
// then: reseved balance to the account should be increased
//    payment_account = PaymentAccount.get(payment_account_id)
//
//    assert Decimal.equal?(payment_account.amount, 1000)
//    assert Decimal.equal?(payment_account.AmountReserved, 200)
//    assert Decimal.equal?(payment_account.amount_available, 800)
//  }
//
//  test "process/1 limit sell order against matching buy order with funds transfer" {
// given:
// -- a seller
//    tradingAccount := Acc()
//    appEntityId := TradingAccount.get(tradingAccount).appEntityId
//
//    seller_debit_account =
//      FindPaymentAccountByAppEntityIdAndCurrencyName(appEntityId, "BTC")
//
//    seller_credit_account =
//      FindPaymentAccountByAppEntityIdAndCurrencyName(appEntityId, "EUR")
//
//    assert seller_debit_account.amount |> Decimal.to_float() == 1000
//    assert seller_credit_account.amount |> Decimal.to_float() == 1000
//
// -- buyer
//    tradingAccount2 := Acc2()
//    appEntityId2 = TradingAccount.get(tradingAccount2).appEntityId
//
//    buyer_debit_account =
//      FindPaymentAccountByAppEntityIdAndCurrencyName(appEntityId2, "EUR")
//
//    buyer_credit_account =
//      FindPaymentAccountByAppEntityIdAndCurrencyName(appEntityId2, "BTC")
//
//    assert buyer_debit_account.amount |> Decimal.to_float() == 1000
//
// when: a trade is executed
//    ProcessTradeOrder(tradingAccount, "BTC_EUR", "LIMIT", "SELL", 10, 10, "GTC")
//    assert MatchingServiceTestHelpers.get_trade_count() == 0
//    assert PaymentAccount.get(seller_debit_account.id).AmountReserved |> Decimal.to_float() == 10
//    ProcessTradeOrder(tradingAccount2, "BTC_EUR", "LIMIT", "BUY", 10, 10, "GTC")
//    assert MatchingServiceTestHelpers.get_trade_count() == 1
//
// then: 4 payments are executed in addition to 4 deposits
//    assert db.QueryVal("SELECT COUNT(*) FROM payment") == 8
//
// and: seller debit balance should decrease for by 10  BTC of seller but reserved balance should be released
//    assert PaymentAccount.get(seller_debit_account.id).amount |> Decimal.to_float() == 990
//    assert PaymentAccount.get(seller_debit_account.id).AmountReserved |> Decimal.to_float() == 0
//
// and: seller credit balance should increase by 100 EUR
//    assert PaymentAccount.get(seller_credit_account.id).amount |> Decimal.to_float() == 1100
//    assert PaymentAccount.get(seller_credit_account.id).AmountReserved |> Decimal.to_float() == 0
//
// and: buy debit balance should decrease by 100 EUR but reserve balance should be released
//    assert PaymentAccount.get(buyer_debit_account.id).amount |> Decimal.to_float() == 900
//    assert PaymentAccount.get(buyer_debit_account.id).AmountReserved |> Decimal.to_float() == 0
//
// and: buy debit balance should increase by 10 BTC
//    assert PaymentAccount.get(buyer_credit_account.id).amount |> Decimal.to_float() == 1010
//    assert PaymentAccount.get(buyer_credit_account.id).AmountReserved |> Decimal.to_float() == 0
//  }
//}
