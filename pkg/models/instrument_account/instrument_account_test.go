package instrumentaccount

import (
	"testing"

	"open-outcry/pkg/db"
	"open-outcry/pkg/models/testutil"

	"github.com/stretchr/testify/assert"
)

func TestInstrumentAccountsStartEmpty(t *testing.T) {
	teardown := testutil.Setup(t)
	defer teardown()

	assert.Equal(t, 0, db.GetCount("instrument_account"))
}
