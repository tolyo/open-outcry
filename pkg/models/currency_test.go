package models

func (assert *ModelsTestSuite) TestGetCurrencies() {
	// expect currencies to be populated
	assert.GreaterOrEqual(3, len(GetCurrencies()))
}
