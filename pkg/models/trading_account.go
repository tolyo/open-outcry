package models

import "open-outcry/pkg/db"

// `trading_account.pub_id` db reference
type TradingAccountId string

type TradingAccount struct {
	Id          TradingAccountId
	AppEntityId AppEntityExternalId
}

const tradingAccountBaseQuery = `
    SELECT t.pub_id, ae.pub_id
    FROM trading_account AS t
    INNER JOIN app_entity ae
          ON ae.id = t.app_entity_id
`

func GetTradingAccount(id TradingAccountId) *TradingAccount {
	return helper(tradingAccountBaseQuery+"WHERE t.pub_id = $1", id)
}

func FindTradingAccountByApplicationEntityId(appEntityId AppEntityId) *TradingAccount {
	return helper(tradingAccountBaseQuery+"WHERE ae.pub_id = $1", appEntityId)
}

func helper(query string, arg any) *TradingAccount {
	var tradingAccount TradingAccount
	db.Instance().QueryRow(
		query, arg,
	).Scan(&tradingAccount.Id, &tradingAccount.AppEntityId)

	return &tradingAccount
}
