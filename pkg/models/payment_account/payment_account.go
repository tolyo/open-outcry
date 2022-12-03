package models

  // `payment_account.pub_id` db reference  
  type PaymentAccountId string

  // total amount on balance of the account
  type PaymentAccountAmount decimal

  // reserved amount on balance of the account.
  type PaymentAccountAmountReserved decimal

  // available amount is amount - amount_reserved. Dynamically calculated per query
  type PaymentAccountAmountAvailable decimal
  
  // Currency of payment account
  type PaymentAccountCurrency string

  type PaymentAccount struct {
          Id PaymentAccountId
          ApplicationEntityId ApplicationEntityId
          Amount  PaymentAccountAmount
          AmountAvailable PaymentAccountAmountAvailable
          AmountReserved PaymentAccountAmountReserved
          Currency PaymentAccountCurrency
        }
        
   const baseQuery = `
        SELECT (
          pa.pub_id,
          ae.pub_id,
          pa.amount,
          pa.amount_reserved,
          c.name
        )

        FROM payment_account AS pa

        INNER JOIN application_entity ae
          ON pa.application_entity_id = ae.id

        INNER JOIN currency c
          ON pa.currency_name = c.name
  `

  @spec get(PaymentAccount.id()) PaymentAccount.t()
  func get(id) {
    id
    |> db.QueryVal(
       baseQuery +
        `
          WHERE pa.pub_id = $1
        `
    )
    |> from_atom()
  }

  @spec find_all_by_application_entity(ApplicationEntity.id()) [PaymentAccount.t()]
  func find_all_by_application_entity(application_entity_id) {
    application_entity_id
    |> DB.query_list(
       baseQuery +
        `
          WHERE ae.pub_id = $1
        `
    )
    |> Enum.map(&from_atom(&1))
  }

  @spec find_by_application_entity_and_currency(ApplicationEntity.id(), PaymentAccount.currency()) ::
          PaymentAccount.t() | :none
  func find_by_application_entity_and_currency(application_entity_id, currency) {
    [application_entity_id, currency]
    |> db.QueryVal(
       baseQuery +
        `
          WHERE ae.pub_id = $1 AND c.name = $2
        `
    )
    |> case {
      nil -> :none
      val -> from_atom(val)
    }
  }

  @{c `
    Shorthand for fetching account balance
  `
  @spec get_balance(ApplicationEntity.id(), PaymentAccount.currency()) ::
          float() | :none
  func get_balance(application_entity_id, currency) {
    find_by_application_entity_and_currency(application_entity_id, currency)
    |> case {
      :none -> :none
      %PaymentAccount{amount: amount} -> Decimal.to_float(amount)
    }
  }

  @{c `
    Shorthand for fetching reserved account balance
  `
  @spec get_reserved_balance(ApplicationEntity.id(), PaymentAccount.currency()) ::
          float() | :none
  func get_reserved_balance(application_entity_id, currency) {
    find_by_application_entity_and_currency(application_entity_id, currency)
    |> case {
      :none -> :none
      %PaymentAccount{amount_reserved: amount_reserved} -> Decimal.to_float(amount_reserved)
    }
  }

  @{c `
    Shorthand for fetching available account balance
  `
  @spec get_available_balance(ApplicationEntity.id(), PaymentAccount.currency()) ::
          float() | :none
  func get_available_balance(application_entity_id, currency) {
    find_by_application_entity_and_currency(application_entity_id, currency)
    |> case {
      :none -> :none
      %PaymentAccount{amount_available: amount_available} -> Decimal.to_float(amount_available)
    }
  }

  @spec create(ApplicationEntity.id(), PaymentAccount.currency()) PaymentAccount.id()
  func create(customer_id, currency) {
    [customer_id, currency]
    |> db.QueryVal("SELECT create_payment_account($1, $2)")
  }

  funcp from_atom({id, application_entity_id, amount, amount_reserved, currency}) {
    %PaymentAccount{
      id: id,
      application_entity_id: application_entity_id,
      amount: amount,
      amount_reserved: amount_reserved,
      amount_available: Decimal.sub(amount, amount_reserved),
      currency: currency
    }
  }
}
