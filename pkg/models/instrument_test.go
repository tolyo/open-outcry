package models

func (assert *ModelsTestSuite) TestGetInstruments() {
	instruments := GetInstruments()
	assert.GreaterOrEqual(len(instruments), 1, "should return at least one non-FX instrument")

	// Verify the seeded SPX instrument is present
	found := false
	for _, inst := range instruments {
		if inst.Name == "SPX" {
			found = true
			assert.NotEmpty(inst.Id, "pub_id should be populated")
			assert.Equal(InstrumentQuoteCurrency("EUR"), inst.QuoteCurrency)
			// Non-FX instruments should not have a base currency populated by the query
			assert.Empty(inst.BaseCurrency)
		}
	}
	assert.True(found, "SPX instrument should be in the non-FX instrument list")
}

func (assert *ModelsTestSuite) TestGetInstrumentsDoesNotIncludeFx() {
	instruments := GetInstruments()
	for _, inst := range instruments {
		assert.NotEqual(InstrumentName("BTC_EUR"), inst.Name,
			"FX instrument BTC_EUR should not appear in non-FX list")
	}
}

func (assert *ModelsTestSuite) TestGetFxInstruments() {
	instruments := GetFxInstruments()
	assert.GreaterOrEqual(len(instruments), 1, "should return at least one FX instrument")

	// Verify the seeded BTC_EUR instrument is present
	found := false
	for _, inst := range instruments {
		if inst.Name == "BTC_EUR" {
			found = true
			assert.NotEmpty(inst.Id, "pub_id should be populated")
			assert.Equal(InstrumentBaseCurrency("BTC"), inst.BaseCurrency)
			assert.Equal(InstrumentQuoteCurrency("EUR"), inst.QuoteCurrency)
		}
	}
	assert.True(found, "BTC_EUR instrument should be in the FX instrument list")
}

func (assert *ModelsTestSuite) TestGetFxInstrumentsDoesNotIncludeNonFx() {
	instruments := GetFxInstruments()
	for _, inst := range instruments {
		assert.NotEqual(InstrumentName("SPX"), inst.Name,
			"Non-FX instrument SPX should not appear in FX list")
	}
}
