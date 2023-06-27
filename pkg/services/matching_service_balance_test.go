package services

import (
	"open-outcry/pkg/models"
	"open-outcry/pkg/utils"
)

type testcase struct {
	currency       models.CurrencyName
	orders         []models.TradeOrder
	expectedStates []models.PaymentAccount
}

var testcases = []testcase{
	{
		currency: "BTC",
		orders: []models.TradeOrder{
			{Side: models.Sell, Price: 10, Amount: 100},
		},
		expectedStates: []models.PaymentAccount{
			{Amount: 1000.0, AmountAvailable: 1000.0, AmountReserved: 0},
			{Amount: 1000.0, AmountAvailable: 900.0, AmountReserved: 100.00},
			{Amount: 1000.0, AmountAvailable: 800.0, AmountReserved: 200.00},
		},
	},
	{
		currency: "EUR",
		orders: []models.TradeOrder{
			{Side: models.Buy, Price: 10, Amount: 10},
		},
		expectedStates: []models.PaymentAccount{
			{Amount: 1000.0, AmountAvailable: 1000.0, AmountReserved: 0},
			{Amount: 1000.0, AmountAvailable: 900.0, AmountReserved: 100.00},
			{Amount: 1000.0, AmountAvailable: 800.0, AmountReserved: 200.00},
		},
	},
}

func (assert *ServiceTestSuite) TestProcessLimitOrderReservedBalance() {
	for _, c := range testcases {
		// given:
		helper := func(expectedState models.PaymentAccount) {
			var paymentAccount = models.FindPaymentAccountByAppEntityIdAndCurrencyName(assert.appEntity1, c.currency)
			assert.Equal(expectedState.Amount, paymentAccount.Amount)
			assert.Equal(expectedState.AmountAvailable, paymentAccount.AmountAvailable)
			assert.Equal(expectedState.AmountReserved, paymentAccount.AmountReserved)
		}
		// then:
		helper(c.expectedStates[0])
		// when: a limit order is sent to an empty matching unit
		ProcessTradeOrder(assert.tradingAccount1, "BTC_EUR", "LIMIT", c.orders[0].Side, c.orders[0].Price, c.orders[0].Amount, "GTC")
		// then: reseved balance to the account should be increased
		helper(c.expectedStates[1])
		// when: another limit order is sent to an empty matching unit
		ProcessTradeOrder(assert.tradingAccount1, "BTC_EUR", "LIMIT", c.orders[0].Side, c.orders[0].Price, c.orders[0].Amount, "GTC")
		helper(c.expectedStates[2])
	}
}

var testCases2 = []testcase{
	{
		orders: []models.TradeOrder{
			{Side: models.Sell, Price: 10, Amount: 10},
			{Side: models.Buy, Price: 10, Amount: 10},
		},
		expectedStates: []models.PaymentAccount{
			// seller
			{Amount: 1000.0, AmountAvailable: 1000.0, AmountReserved: 0}, //debit
			{Amount: 1000.0, AmountAvailable: 1000.0, AmountReserved: 0}, //credit
			// buyer
			{Amount: 1000.0, AmountAvailable: 1000.0, AmountReserved: 0}, //debit
			{Amount: 1000.0, AmountAvailable: 1000.0, AmountReserved: 0}, //credit

			// seller
			{Amount: 1000.0, AmountAvailable: 990.0, AmountReserved: 10.00}, //debit
			{Amount: 1000.0, AmountAvailable: 1000.0, AmountReserved: 0},    //credit

			// buyer
			{Amount: 1000.0, AmountAvailable: 1000.0, AmountReserved: 0}, //debit
			{Amount: 1000.0, AmountAvailable: 1000.0, AmountReserved: 0}, //credit

			// seller
			{Amount: 990.0, AmountAvailable: 990.0, AmountReserved: 0},   //debit
			{Amount: 1100.0, AmountAvailable: 1100.0, AmountReserved: 0}, //credit
			// buyer
			{Amount: 900.0, AmountAvailable: 900.0, AmountReserved: 0},   //debit
			{Amount: 1010.0, AmountAvailable: 1010.0, AmountReserved: 0}, //credit
		},
	},
}

func (assert *ServiceTestSuite) TestProcessLimitSellOrderAgainstMatchingBuyOrderWithFundsTransfer() {
	for _, c := range testCases2 {
		helper := func(accounts []models.PaymentAccount) {
			var sellerDebitAccount = models.FindPaymentAccountByAppEntityIdAndCurrencyName(assert.appEntity1, "BTC")
			var sellerCreditAccount = models.FindPaymentAccountByAppEntityIdAndCurrencyName(assert.appEntity1, "EUR")
			var buyerDebitAccount = models.FindPaymentAccountByAppEntityIdAndCurrencyName(assert.appEntity2, "EUR")
			var buyerCreditAccount = models.FindPaymentAccountByAppEntityIdAndCurrencyName(assert.appEntity2, "BTC")
			assert.Equal(accounts[0].Amount, sellerDebitAccount.Amount)
			assert.Equal(accounts[0].AmountAvailable, sellerDebitAccount.AmountAvailable)
			assert.Equal(accounts[0].AmountReserved, sellerDebitAccount.AmountReserved)

			assert.Equal(accounts[1].Amount, sellerCreditAccount.Amount)
			assert.Equal(accounts[1].AmountAvailable, sellerCreditAccount.AmountAvailable)
			assert.Equal(accounts[1].AmountReserved, sellerCreditAccount.AmountReserved)

			assert.Equal(accounts[2].Amount, buyerDebitAccount.Amount)
			assert.Equal(accounts[2].AmountAvailable, buyerDebitAccount.AmountAvailable)
			assert.Equal(accounts[2].AmountReserved, buyerDebitAccount.AmountReserved)

			assert.Equal(accounts[3].Amount, buyerCreditAccount.Amount)
			assert.Equal(accounts[3].AmountAvailable, buyerCreditAccount.AmountAvailable)
			assert.Equal(accounts[3].AmountReserved, buyerCreditAccount.AmountReserved)

		}
		// given: -- a seller and a buyer
		helper(c.expectedStates[0:4])

		ProcessTradeOrder(assert.tradingAccount1, "BTC_EUR", "LIMIT", c.orders[0].Side, c.orders[0].Price, c.orders[0].Amount, "GTC")
		assert.Equal(0, GetTradeCount())
		helper(c.expectedStates[4:8])
		// when: a trade is executed
		ProcessTradeOrder(assert.tradingAccount2, "BTC_EUR", "LIMIT", c.orders[1].Side, c.orders[1].Price, c.orders[1].Amount, "GTC")
		assert.Equal(1, GetTradeCount())
		// then: 4 payments are executed in addition to 4 deposits
		assert.Equal(8, utils.GetCount("payment"))
		helper(c.expectedStates[8:12])

	}
}
