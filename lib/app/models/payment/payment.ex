defmodule Payment do
  require Logger

  @type id :: String.t()

  @type amount :: Decimal.t()

  @type currency :: PaymentAccount.currency()

  @type details :: String.t()

  @type external_reference_number :: String.t()

  defstruct pub_id: nil,
            number: nil,
            # type of payment
            type: nil,
            amount: nil,
            currency: nil,
            sender_account_id: nil,
            beneficiary_account_id: nil,
            details: nil,
            external_reference_number: nil,
            fee_sender: nil,
            # for deposits
            fee_beneficiary: nil,
            status: nil,
            total_amount: nil,
            debit_balance_amount: nil,
            credit_balance_amount: nil,
            created_at: nil,
            updated_at: nil
end
