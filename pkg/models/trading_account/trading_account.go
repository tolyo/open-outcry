package models

// `trading_account.pub_id` db reference
  type TradingAccountId string

  type TradingAccount struct {
      id TradingAccountId
      ApplicationEntityId ApplicationEntityExternalId
  }

  const baseQuery =  `
    SELECT (
      t.pub_id,
      ae.pub_id
    )

    FROM trading_account AS t

    INNER JOIN application_entity ae
          ON ae.id = t.application_entity_id

  `

  func Get(id TradingAccountId) TradingAccount {
    db.QueryVal(baseQuery +"WHERE t.pub_id = $1",id)
  }

  func FindByApplicationEntity(application_entity_id ApplicationEntityId) TradingAccount {

    db.QueryVal(
      baseQuery + "WHERE ae.pub_id = $1", application_entity_id
    )
    
  }