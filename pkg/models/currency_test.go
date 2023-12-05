package models

func (assert *ModelsTestSuite) TestGetCurrencies() {
	// expect currenceis to be popullated
	assert.GreaterOrEqual(3, len(GetCurrencies()))
}
