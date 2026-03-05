package models

import (
	"open-outcry/pkg/db"
)

func (assert *ModelsTestSuite) TestGetAppEntities() {
	entities := GetAppEntities()
	assert.GreaterOrEqual(len(entities), 1)

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

func (assert *ModelsTestSuite) TestGetCurrencyAccountsByAppEntity() {
	masterPubId := FindAppEntityExternalId("MASTER")
	accounts := GetCurrencyAccountsByAppEntity(masterPubId)
	assert.GreaterOrEqual(len(accounts), 2) // EUR + BTC
}

func (assert *ModelsTestSuite) TestGetCurrencyAccount() {
	masterPubId := FindAppEntityExternalId("MASTER")
	account := FindCurrencyAccountByAppEntityIdAndCurrencyName(masterPubId, "EUR")
	assert.NotNil(account)
	assert.Equal(CurrencyName("EUR"), account.Currency)

	fetched := GetCurrencyAccount(account.Id)
	assert.NotNil(fetched)
	assert.Equal(account.Id, fetched.Id)
}

func (assert *ModelsTestSuite) TestGetInstrumentAccount() {
	assert.Equal(0, db.GetCount("instrument_account"))
}
