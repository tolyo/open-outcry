package models

import (
	"open-outcry/pkg/db"
)

func (assert *ModelsTestSuite) TestGetAppEntities() {
	// The seed data creates one MASTER entity
	entities := GetAppEntities()
	assert.GreaterOrEqual(len(entities), 1)

	// Find the MASTER entity
	found := false
	for _, e := range entities {
		if e.Type == Master {
			found = true
			assert.Equal(AppEntityExternalId("MASTER"), e.ExternalId)
		}
	}
	assert.True(found, "MASTER entity should exist")
}

func (assert *ModelsTestSuite) TestGetAppEntity() {
	// Get the MASTER entity by pub_id
	masterPubId := FindAppEntityExternalId("MASTER")

	entity := GetAppEntity(masterPubId)
	assert.NotNil(entity)
	assert.Equal(masterPubId, entity.Id)
	assert.Equal(Master, entity.Type)
	assert.Equal(AppEntityExternalId("MASTER"), entity.ExternalId)
}

func (assert *ModelsTestSuite) TestGetAppEntityNotFound() {
	entity := GetAppEntity("nonexistent-id")
	assert.Nil(entity)
}

func (assert *ModelsTestSuite) TestGetPaymentAccountsByAppEntity() {
	// MASTER has payment accounts created in seeds
	masterPubId := FindAppEntityExternalId("MASTER")
	accounts := GetPaymentAccountsByAppEntity(masterPubId)
	assert.GreaterOrEqual(len(accounts), 2) // EUR + BTC
}

func (assert *ModelsTestSuite) TestGetPaymentAccount() {
	masterPubId := FindAppEntityExternalId("MASTER")
	account := FindPaymentAccountByAppEntityIdAndCurrencyName(masterPubId, "EUR")
	assert.NotNil(account)
	assert.Equal(CurrencyName("EUR"), account.Currency)

	fetched := GetPaymentAccount(account.Id)
	assert.NotNil(fetched)
	assert.Equal(account.Id, fetched.Id)
}

func (assert *ModelsTestSuite) TestGetTradingAccount() {
	// MASTER doesn't have a trading account, but it's tested in service tests
	// Verify the count is right
	assert.Equal(0, db.GetCount("trading_account"))
}
