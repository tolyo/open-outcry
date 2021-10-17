defmodule PaymentService do
  @spec deposit(
          ApplicationEntity.id(),
          number(),
          PaymentAccount.currency(),
          Payment.external_reference_number(),
          Payment.details()
        ) :: Payment.id()
  def deposit(customer_id, amount, currency, reference, details) do
    [customer_id, amount, currency, reference, details]
    |> DB.query_val("
      SELECT create_payment('DEPOSIT', 'MASTER', $2, $3, $1, $4, $5)
    ")
  end

  @spec transfer(
          ApplicationEntity.id(),
          ApplicationEntity.id(),
          number(),
          PaymentAccount.currency(),
          Payment.details()
        ) :: Payment.id() | {:error, String.t()}
  def transfer(customer_from_id, customer_to_id, amount, currency, details) do
    [customer_from_id, customer_to_id, amount, currency, details]
    |> DB.query_val("
      SELECT create_payment('TRANSFER', $1, $3, $4, $2, 'N/A', $5)
    ")
  end
end
