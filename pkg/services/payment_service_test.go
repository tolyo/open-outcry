package services

import (
	"open-outcry/pkg/models"
	"open-outcry/pkg/utils"
)

func (assert *ServiceTestSuite) TestDepositPayment() {
	// given a customer
	utils.DeleteAll("payment")
	appEntity1, _ := Acc("test3")

	// when amount is deposited
	CreatePaymentDeposit(appEntity1, 10.00, "EUR", "BANK", "REF123")

	// then amount should increase and payment should be created
	acc := models.FindPaymentAccountByAppEntityIdAndCurrencyName(appEntity1, "EUR")
	assert.Equal(1010.00, acc.Amount)
	assert.Equal(1010.00, acc.AmountAvailable)
	assert.Equal(0.00, acc.AmountReserved)
	assert.Equal(3, utils.GetCount("payment"))

	// when amount is deposited
	CreatePaymentDeposit(appEntity1, 10.00, "EUR", "BANK", "REF125")

	// then amount should increase and payment should be created
	acc = models.FindPaymentAccountByAppEntityIdAndCurrencyName(appEntity1, "EUR")
	assert.Equal(1020.00, acc.Amount)
	assert.Equal(1020.00, acc.AmountAvailable)
	assert.Equal(0.00, acc.AmountReserved)
	assert.Equal(4, utils.GetCount("payment"))
}
