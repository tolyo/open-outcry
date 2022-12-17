package services

//funcmodule MatchingServiceBalanceTest {
//  use DataCase
//
func (assert *ServiceTestSuite) TestProcessLimitSellOrderSaveWithReservedBalance() {
// given:
//    appEntityId := CreateClient()
//    models.CreatePaymentAccount(appEntityId, "BTC")
//    tradingAccountId := model.FindTradingAccountByAppEntityId(appEntityId).id
//
//    payment_account =
//      models.FindPaymentAccountByAppEntityIdAndCurrencyName(appEntityId, "BTC")
//
//    assert.Equal(Decimal.equal?(payment_account.amount, 0)
//    assert.Equal(Decimal.equal?(payment_account.AmountReserved, 0)
//    assert.Equal(Decimal.equal?(payment_account.amount_available, 0)
//
// when:
//    CreatePaymentDeposit(appEntityId, 1000, "BTC", "Test", "Test")
//
// then:
//    payment_account =
//      models.FindPaymentAccountByAppEntityIdAndCurrencyName(appEntityId, "BTC")
//
//    assert.Equal(Decimal.equal?(payment_account.amount, 1000)
//    assert.Equal(Decimal.equal?(payment_account.AmountReserved, 0)
//    assert.Equal(Decimal.equal?(payment_account.amount_available, 1000)
//
// when: a limit order is sent to an empty matching unit
//    ProcessTradeOrder(tradingAccountId, "BTC_EUR", "LIMIT", "SELL", 10, 100, "GTC")
//
// then: reseved balance to the account should be increased
//    payment_account =
//      models.FindPaymentAccountByAppEntityIdAndCurrencyName(appEntityId, "BTC")
//
//    assert.Equal(Decimal.equal?(payment_account.amount, 1000)
//    assert.Equal(Decimal.equal?(payment_account.AmountReserved, 100)
//    assert.Equal(Decimal.equal?(payment_account.amount_available, 900)
//
// when: another limit order is sent to an empty matching unit
//    ProcessTradeOrder(tradingAccountId, "BTC_EUR", "LIMIT", "SELL", 10, 100, "GTC")
//
// then: reseved balance to the account should be increased
//    payment_account =
//      models.FindPaymentAccountByAppEntityIdAndCurrencyName(appEntityId, "BTC")
//
//    assert.Equal(Decimal.equal?(payment_account.amount, 1000)
//    assert.Equal(Decimal.equal?(payment_account.AmountReserved, 200)
//    assert.Equal(Decimal.equal?(payment_account.amount_available, 800)
}

func (assert *ServiceTestSuite) TestProcessLimitBuyOrderSaveWithReservedBalance() {
// given:
//    appEntityId := CreateClient()
//
//    payment_account_id =
//      models.FindPaymentAccountByAppEntityIdAndCurrencyName(appEntityId, "EUR").id
//
//    tradingAccountId := model.FindTradingAccountByAppEntityId(appEntityId).id
//
//    payment_account = PaymentAccount.get(payment_account_id)
//
//    assert.Equal(Decimal.equal?(payment_account.amount, 0)
//    assert.Equal(Decimal.equal?(payment_account.AmountReserved, 0)
//    assert.Equal(Decimal.equal?(payment_account.amount_available, 0)
//
// when:
//    CreatePaymentDeposit(appEntityId, 1000, "EUR", "Test", "Test")
//
// then:
//    payment_account = PaymentAccount.get(payment_account_id)
//
//    assert.Equal(Decimal.equal?(payment_account.amount, 1000)
//    assert.Equal(Decimal.equal?(payment_account.AmountReserved, 0)
//    assert.Equal(Decimal.equal?(payment_account.amount_available, 1000)
//
// when: a limit order is sent to an empty matching unit
//    ProcessTradeOrder(tradingAccountId, "BTC_EUR", "LIMIT", "BUY", 10, 10, "GTC")
//
// then: reseved balance to the account should be increased
//    payment_account = PaymentAccount.get(payment_account_id)
//
//    assert.Equal(Decimal.equal?(payment_account.amount, 1000)
//    assert.Equal(Decimal.equal?(payment_account.AmountReserved, 100)
//    assert.Equal(Decimal.equal?(payment_account.amount_available, 900)
//
// when: another limit order is sent to an empty matching unit
//    ProcessTradeOrder(tradingAccountId, "BTC_EUR", "LIMIT", "BUY", 10, 10, "GTC")
//
// then: reseved balance to the account should be increased
//    payment_account = PaymentAccount.get(payment_account_id)
//
//    assert.Equal(Decimal.equal?(payment_account.amount, 1000)
//    assert.Equal(Decimal.equal?(payment_account.AmountReserved, 200)
//    assert.Equal(Decimal.equal?(payment_account.amount_available, 800)
}

func (assert *ServiceTestSuite) TestProcessLimitSellOrderAgainstMatchingBuyOrderWithFundsTransfer() {
// given:
// -- a seller
//    tradingAccount := Acc()
//    appEntityId := TradingAccount.get(tradingAccount).appEntityId
//
//    seller_debit_account =
//      models.FindPaymentAccountByAppEntityIdAndCurrencyName(appEntityId, "BTC")
//
//    seller_credit_account =
//      models.FindPaymentAccountByAppEntityIdAndCurrencyName(appEntityId, "EUR")
//
//    assert.Equal(seller_debit_account.amount |> Decimal.to_float() == 1000
//    assert.Equal(seller_credit_account.amount |> Decimal.to_float() == 1000
//
// -- buyer
//    tradingAccount2 := Acc2()
//    appEntityId2 = TradingAccount.get(tradingAccount2).appEntityId
//
//    buyer_debit_account =
//      models.FindPaymentAccountByAppEntityIdAndCurrencyName(appEntityId2, "EUR")
//
//    buyer_credit_account =
//      models.FindPaymentAccountByAppEntityIdAndCurrencyName(appEntityId2, "BTC")
//
//    assert.Equal(buyer_debit_account.amount |> Decimal.to_float() == 1000
//
// when: a trade is executed
//    ProcessTradeOrder(tradingAccount, "BTC_EUR", "LIMIT", "SELL", 10, 10, "GTC")
//    assert.Equal(get_trade_count() == 0
//    assert.Equal(PaymentAccount.get(seller_debit_account.id).AmountReserved |> Decimal.to_float() == 10
//    ProcessTradeOrder(tradingAccount2, "BTC_EUR", "LIMIT", "BUY", 10, 10, "GTC")
//    assert.Equal(get_trade_count() == 1
//
// then: 4 payments are executed in addition to 4 deposits
//    assert.Equal(db.QueryVal("SELECT COUNT(*) FROM payment") == 8
//
// and: seller debit balance should decrease for by 10  BTC of seller but reserved balance should be released
//    assert.Equal(PaymentAccount.get(seller_debit_account.id).amount |> Decimal.to_float() == 990
//    assert.Equal(PaymentAccount.get(seller_debit_account.id).AmountReserved |> Decimal.to_float() == 0
//
// and: seller credit balance should increase by 100 EUR
//    assert.Equal(PaymentAccount.get(seller_credit_account.id).amount |> Decimal.to_float() == 1100
//    assert.Equal(PaymentAccount.get(seller_credit_account.id).AmountReserved |> Decimal.to_float() == 0
//
// and: buy debit balance should decrease by 100 EUR but reserve balance should be released
//    assert.Equal(PaymentAccount.get(buyer_debit_account.id).amount |> Decimal.to_float() == 900
//    assert.Equal(PaymentAccount.get(buyer_debit_account.id).AmountReserved |> Decimal.to_float() == 0
//
// and: buy debit balance should increase by 10 BTC
//    assert.Equal(PaymentAccount.get(buyer_credit_account.id).amount |> Decimal.to_float() == 1010
//    assert.Equal(PaymentAccount.get(buyer_credit_account.id).AmountReserved |> Decimal.to_float() == 0
//  }
}
