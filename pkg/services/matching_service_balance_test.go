package services

import (
	"open-outcry/pkg/db"
	"open-outcry/pkg/models"
)

func (assert *ServiceTestSuite) TestProcessLimitSellOrderSaveWithReservedBalance() {
	// given:
	tradingAccountId := models.FindTradingAccountByApplicationEntityId(assert.appEntity1).Id
	paymentAccount := models.FindPaymentAccountByAppEntityIdAndCurrencyName(assert.appEntity1, "BTC")

	assert.Equal(1000.0, paymentAccount.Amount)
	assert.Equal(1000.0, paymentAccount.AmountAvailable)
	assert.Equal(0.0, paymentAccount.AmountReserved)

	// when:

	// then:
	paymentAccount = models.GetPaymentAccount(paymentAccount.Id)
	assert.Equal(1000.0, paymentAccount.Amount)
	assert.Equal(1000.0, paymentAccount.AmountAvailable)
	assert.Equal(0.0, paymentAccount.AmountReserved)

	// when: a limit order is sent to an empty matching unit
	ProcessTradeOrder(tradingAccountId, "BTC_EUR", "LIMIT", models.Sell, 10, 100, "GTC")

	// then: reseved balance to the account should be increased
	paymentAccount = models.FindPaymentAccountByAppEntityIdAndCurrencyName(assert.appEntity1, "BTC")
	assert.Equal(1000.0, paymentAccount.Amount)
	assert.Equal(100.0, paymentAccount.AmountReserved)
	assert.Equal(900.0, paymentAccount.AmountAvailable)

	// when: another limit order is sent to an empty matching unit
	ProcessTradeOrder(tradingAccountId, "BTC_EUR", "LIMIT", models.Sell, 10, 100, "GTC")

	// then: reseved balance to the account should be increased
	paymentAccount = models.FindPaymentAccountByAppEntityIdAndCurrencyName(assert.appEntity1, "BTC")
	assert.Equal(1000.0, paymentAccount.Amount)
	assert.Equal(200.0, paymentAccount.AmountReserved)
	assert.Equal(800.0, paymentAccount.AmountAvailable)
}

func (assert *ServiceTestSuite) TestProcessLimitBuyOrderSaveWithReservedBalance() {
	// given:
	tradingAccountId := models.FindTradingAccountByApplicationEntityId(assert.appEntity1).Id
	paymentAccount := models.FindPaymentAccountByAppEntityIdAndCurrencyName(assert.appEntity1, "EUR")

	assert.Equal(1000.0, paymentAccount.Amount)
	assert.Equal(0.0, paymentAccount.AmountReserved)
	assert.Equal(1000.0, paymentAccount.AmountAvailable)
	// when:

	// when: a limit order is sent to an empty matching unit
	ProcessTradeOrder(tradingAccountId, "BTC_EUR", "LIMIT", models.Buy, 10, 10, "GTC")

	// then: reseved balance to the account should be increased
	paymentAccount = models.GetPaymentAccount(paymentAccount.Id)
	assert.Equal(1000.0, paymentAccount.Amount)
	assert.Equal(100.0, paymentAccount.AmountReserved)
	assert.Equal(900.0, paymentAccount.AmountAvailable)

	// when: another limit order is sent to an empty matching unit
	ProcessTradeOrder(tradingAccountId, "BTC_EUR", "LIMIT", models.Buy, 10, 10, "GTC")

	// then: reseved balance to the account should be increased
	paymentAccount = models.GetPaymentAccount(paymentAccount.Id)

	assert.Equal(1000.0, paymentAccount.Amount)
	assert.Equal(200.0, paymentAccount.AmountReserved)
	assert.Equal(800.0, paymentAccount.AmountAvailable)
}

func (assert *ServiceTestSuite) TestProcessLimitSellOrderAgainstMatchingBuyOrderWithFundsTransfer() {
	// given: -- a seller
	sellerDebitAccount := models.FindPaymentAccountByAppEntityIdAndCurrencyName(assert.appEntity1, "BTC")
	sellerCreditAccount := models.FindPaymentAccountByAppEntityIdAndCurrencyName(assert.appEntity1, "EUR")

	assert.Equal(1000.0, sellerDebitAccount.Amount)
	assert.Equal(1000.0, sellerCreditAccount.Amount)

	buyerDebitAccount := models.FindPaymentAccountByAppEntityIdAndCurrencyName(assert.appEntity2, "EUR")
	buyerCreditAccount := models.FindPaymentAccountByAppEntityIdAndCurrencyName(assert.appEntity2, "BTC")

	assert.Equal(1000.0, buyerDebitAccount.Amount)

	// when: a trade is executed
	ProcessTradeOrder(assert.tradingAccount1, "BTC_EUR", "LIMIT", models.Sell, 10, 10, "GTC")
	assert.Equal(0, GetTradeCount())
	assert.Equal(10.0, models.GetPaymentAccount(sellerDebitAccount.Id).AmountReserved)
	ProcessTradeOrder(assert.tradingAccount2, "BTC_EUR", "LIMIT", models.Buy, 10, 10, "GTC")
	assert.Equal(1, GetTradeCount())

	// then: 4 payments are executed in addition to 4 deposits
	assert.Equal(8, db.QueryVal[int]("SELECT COUNT(*) FROM payment"))

	// and: seller debit balance should decrease for by 10  BTC of seller but reserved balance should be released
	assert.Equal(990.00, models.GetPaymentAccount(sellerDebitAccount.Id).Amount)
	assert.Equal(0.0, models.GetPaymentAccount(sellerDebitAccount.Id).AmountReserved)

	// and: seller credit balance should increase by 100 EUR
	assert.Equal(1100.0, models.GetPaymentAccount(sellerCreditAccount.Id).Amount)
	assert.Equal(0.0, models.GetPaymentAccount(sellerCreditAccount.Id).AmountReserved)

	// and: buy debit balance should decrease by 100 EUR but reserve balance should be released
	assert.Equal(900.0, models.GetPaymentAccount(buyerDebitAccount.Id).Amount)
	assert.Equal(0.0, models.GetPaymentAccount(buyerDebitAccount.Id).AmountReserved)

	// and: buy debit balance should increase by 10 BTC
	assert.Equal(1010.0, models.GetPaymentAccount(buyerCreditAccount.Id).Amount)
	assert.Equal(0.0, models.GetPaymentAccount(buyerCreditAccount.Id).AmountReserved)
}
