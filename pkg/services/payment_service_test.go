package services

import (
	"open-outcry/pkg/models"
	"open-outcry/pkg/utils"
)

func (assert *ServiceTestSuite) TestDepositPayment() {
	// given a customer
	pubId := CreateClient()
	assert.Equal(0, utils.GetCount("payment"))

	// when amount is deposited
	CreatePaymentDeposit(pubId, 10.00, "EUR", "BANK", "REF123")

	// then amount should increase and payment should be created
	acc := models.FindPaymentAccountByAppEntityIdAndCurrencyName(pubId, "EUR")
	assert.Equal(10.00, acc.Amount)
	assert.Equal(10.00, acc.AmountAvailable)
	assert.Equal(0.00, acc.AmountReserved)
	assert.Equal(1, utils.GetCount("payment"))

	// when amount is deposited
	CreatePaymentDeposit(pubId, 10.00, "EUR", "BANK", "REF125")

	// then amount should increase and payment should be created
	acc = models.FindPaymentAccountByAppEntityIdAndCurrencyName(pubId, "EUR")
	assert.Equal(20.00, acc.Amount)
	assert.Equal(20.00, acc.AmountAvailable)
	assert.Equal(0.00, acc.AmountReserved)
	assert.Equal(2, utils.GetCount("payment"))
}
