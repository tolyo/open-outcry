package currencyaccount

import (
	"testing"

	appentity "open-outcry/pkg/models/app_entity"
	"open-outcry/pkg/models/testutil"

	"github.com/stretchr/testify/assert"
)

func TestGetCurrencyAccountsByAppEntity(t *testing.T) {
	teardown := testutil.Setup(t)
	defer teardown()

	masterPubID := appentity.FindAppEntityExternalId("MASTER")
	accounts := GetCurrencyAccountsByAppEntity(masterPubID)
	assert.GreaterOrEqual(t, len(accounts), 2)
}

func TestGetCurrencyAccount(t *testing.T) {
	teardown := testutil.Setup(t)
	defer teardown()

	masterPubID := appentity.FindAppEntityExternalId("MASTER")
	account := FindCurrencyAccountByAppEntityIdAndCurrencyName(masterPubID, "EUR")
	assert.NotNil(t, account)
	assert.Equal(t, CurrencyAccountId(account.Id), account.Id)

	fetched := GetCurrencyAccount(account.Id)
	assert.NotNil(t, fetched)
	assert.Equal(t, account.Id, fetched.Id)
}
