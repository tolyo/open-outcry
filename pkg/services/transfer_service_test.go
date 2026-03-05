package services

import (
	"open-outcry/pkg/db"
	"open-outcry/pkg/models"
)

func (assert *ServiceTestSuite) TestDepositTransfer() {
	// given a customer
	db.DeleteAll("transfer")
	appEntity1, _ := Acc("test3")

	// when amount is deposited
	CreateTransferDeposit(appEntity1, 10.00, "EUR", "BANK", "REF123")

	// then amount should increase and transfer should be created
	acc := models.FindCurrencyAccountByAppEntityIdAndCurrencyName(appEntity1, "EUR")
	assert.Equal(1010.00, acc.Amount)
	assert.Equal(1010.00, acc.AmountAvailable)
	assert.Equal(0.00, acc.AmountReserved)
	assert.Equal(3, db.GetCount("transfer"))

	// when amount is deposited
	CreateTransferDeposit(appEntity1, 10.00, "EUR", "BANK", "REF125")

	// then amount should increase and transfer should be created
	acc = models.FindCurrencyAccountByAppEntityIdAndCurrencyName(appEntity1, "EUR")
	assert.Equal(1020.00, acc.Amount)
	assert.Equal(1020.00, acc.AmountAvailable)
	assert.Equal(0.00, acc.AmountReserved)
	assert.Equal(4, db.GetCount("transfer"))
}
