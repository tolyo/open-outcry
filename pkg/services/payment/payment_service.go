funcmodule PaymentService {
  @spec deposit(
          ApplicationEntity.id(),
          decimal,
          PaymentAccount.currency(),
          Payment.external_reference_number(),
          Payment.details()
        ) Payment.id()
  func deposit(customer_id, amount, currency, reference, details) {
    [customer_id, amount, currency, reference, details]
    |> db.QueryVal("
      SELECT create_payment('DEPOSIT', 'MASTER', $2, $3, $1, $4, $5)
    ")
  }

  @spec transfer(
          ApplicationEntity.id(),
          ApplicationEntity.id(),
          decimal,
          PaymentAccount.currency(),
          Payment.details()
        ) Payment.id() | {:error, String.t()}
  func transfer(customer_from_id, customer_to_id, amount, currency, details) {
    [customer_from_id, customer_to_id, amount, currency, details]
    |> db.QueryVal("
      SELECT create_payment('TRANSFER', $1, $3, $4, $2, 'N/A', $5)
    ")
  }
}
