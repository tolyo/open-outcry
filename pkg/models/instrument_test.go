package models

func (assert *ModelsTestSuite) TestGetInstruments() {
	assert.GreaterOrEqual(1, len(GetInstruments()))
}

func (assert *ModelsTestSuite) TestGetFxInstruments() {
	assert.GreaterOrEqual(1, len(GetFxInstruments()))
}
