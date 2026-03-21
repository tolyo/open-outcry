package services

import (
	"open-outcry/pkg/db"
	"open-outcry/pkg/models/currency_account"
)

func (assert *ServiceTestSuite) TestDepositTransfer() {
	db.DeleteAll("transfer")
	appEntity1, _ := Acc("test3")
	CreateTransferDeposit(appEntity1, 10.00, "EUR", "BANK", "REF123")
	acc := currencyaccount.FindCurrencyAccountByAppEntityIdAndCurrencyName(appEntity1, "EUR")
	assert.Equal(1010.00, acc.Amount)
	assert.Equal(1010.00, acc.AmountAvailable)
	assert.Equal(0.00, acc.AmountReserved)
	assert.Equal(3, db.GetCount("transfer"))
	CreateTransferDeposit(appEntity1, 10.00, "EUR", "BANK", "REF125")
	acc = currencyaccount.FindCurrencyAccountByAppEntityIdAndCurrencyName(appEntity1, "EUR")
	assert.Equal(1020.00, acc.Amount)
	assert.Equal(1020.00, acc.AmountAvailable)
	assert.Equal(0.00, acc.AmountReserved)
	assert.Equal(4, db.GetCount("transfer"))
}
