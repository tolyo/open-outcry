package currency

import (
	"testing"

	"open-outcry/pkg/models/testutil"

	"github.com/stretchr/testify/assert"
)

func TestGetCurrencies(t *testing.T) {
	teardown := testutil.Setup(t)
	defer teardown()

	assert.GreaterOrEqual(t, len(GetCurrencies()), 3)
}
