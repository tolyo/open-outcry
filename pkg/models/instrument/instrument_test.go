package instrument

import (
	"testing"

	"open-outcry/pkg/models/testutil"

	"github.com/stretchr/testify/assert"
)

func TestGetInstruments(t *testing.T) {
	teardown := testutil.Setup(t)
	defer teardown()

	instruments := GetInstruments()
	assert.GreaterOrEqual(t, len(instruments), 1, "should return at least one non-FX instrument")

	found := false
	for _, inst := range instruments {
		if inst.Name == "SPX" {
			found = true
			assert.NotEmpty(t, inst.Id, "pub_id should be populated")
			assert.Equal(t, InstrumentQuoteCurrency("EUR"), inst.QuoteCurrency)
			assert.Empty(t, inst.BaseCurrency)
		}
	}
	assert.True(t, found, "SPX instrument should be in the non-FX instrument list")
}

func TestGetInstrumentsDoesNotIncludeFx(t *testing.T) {
	teardown := testutil.Setup(t)
	defer teardown()

	instruments := GetInstruments()
	for _, inst := range instruments {
		assert.NotEqual(t, InstrumentName("BTC_EUR"), inst.Name,
			"FX instrument BTC_EUR should not appear in non-FX list")
	}
}

func TestGetFxInstruments(t *testing.T) {
	teardown := testutil.Setup(t)
	defer teardown()

	instruments := GetFxInstruments()
	assert.GreaterOrEqual(t, len(instruments), 1, "should return at least one FX instrument")

	found := false
	for _, inst := range instruments {
		if inst.Name == "BTC_EUR" {
			found = true
			assert.NotEmpty(t, inst.Id, "pub_id should be populated")
			assert.Equal(t, InstrumentBaseCurrency("BTC"), inst.BaseCurrency)
			assert.Equal(t, InstrumentQuoteCurrency("EUR"), inst.QuoteCurrency)
		}
	}
	assert.True(t, found, "BTC_EUR instrument should be in the FX instrument list")
}

func TestGetFxInstrumentsDoesNotIncludeNonFx(t *testing.T) {
	teardown := testutil.Setup(t)
	defer teardown()

	instruments := GetFxInstruments()
	for _, inst := range instruments {
		assert.NotEqual(t, InstrumentName("SPX"), inst.Name,
			"Non-FX instrument SPX should not appear in FX instrument list")
	}
}
