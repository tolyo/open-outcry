defmodule PaymentAccount do
  require Logger

  @typedoc """
    `payment_account.pub_id` db reference
  """
  @type id :: String.t()

  @typedoc """
    total amount on balance of the account
  """
  @type amount :: Decimal.t()

  @typedoc """
    reserved amount on balance of the account.
  """
  @type amount_reserved :: Decimal.t()

  @typedoc """
    available amount is amount - amount_reserved. Dynamically calculated per query
  """
  @type amount_available :: Decimal.t()

  @typedoc """
    Currency of payment account
  """
  @type currency :: String.t()

  @type t :: %PaymentAccount{
          id: id(),
          application_entity_id: ApplicationEntity.id(),
          amount: amount(),
          amount_available: amount_available(),
          amount_reserved: amount_reserved(),
          currency: currency()
        }
  @derive Jason.Encoder
  defstruct id: nil,
            application_entity_id: nil,
            amount: 0.00,
            amount_reserved: 0.00,
            amount_available: 0.00,
            currency: nil

  @base_query """
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
  """

  @spec get(PaymentAccount.id()) :: PaymentAccount.t()
  def get(id) do
    id
    |> DB.query_val(
      @base_query <>
        """
          WHERE pa.pub_id = $1
        """
    )
    |> from_atom()
  end

  @spec find_all_by_application_entity(ApplicationEntity.id()) :: [PaymentAccount.t()]
  def find_all_by_application_entity(application_entity_id) do
    application_entity_id
    |> DB.query_list(
      @base_query <>
        """
          WHERE ae.pub_id = $1
        """
    )
    |> Enum.map(&from_atom(&1))
  end

  @spec find_by_application_entity_and_currency(ApplicationEntity.id(), PaymentAccount.currency()) ::
          PaymentAccount.t() | :none
  def find_by_application_entity_and_currency(application_entity_id, currency) do
    [application_entity_id, currency]
    |> DB.query_val(
      @base_query <>
        """
          WHERE ae.pub_id = $1 AND c.name = $2
        """
    )
    |> case do
      nil -> :none
      val -> from_atom(val)
    end
  end

  @doc """
    Shorthand for fetching account balance
  """
  @spec get_balance(ApplicationEntity.id(), PaymentAccount.currency()) ::
          float() | :none
  def get_balance(application_entity_id, currency) do
    find_by_application_entity_and_currency(application_entity_id, currency)
    |> case do
      :none -> :none
      %PaymentAccount{amount: amount} -> Decimal.to_float(amount)
    end
  end

  @doc """
    Shorthand for fetching reserved account balance
  """
  @spec get_reserved_balance(ApplicationEntity.id(), PaymentAccount.currency()) ::
          float() | :none
  def get_reserved_balance(application_entity_id, currency) do
    find_by_application_entity_and_currency(application_entity_id, currency)
    |> case do
      :none -> :none
      %PaymentAccount{amount_reserved: amount_reserved} -> Decimal.to_float(amount_reserved)
    end
  end

  @doc """
    Shorthand for fetching available account balance
  """
  @spec get_available_balance(ApplicationEntity.id(), PaymentAccount.currency()) ::
          float() | :none
  def get_available_balance(application_entity_id, currency) do
    find_by_application_entity_and_currency(application_entity_id, currency)
    |> case do
      :none -> :none
      %PaymentAccount{amount_available: amount_available} -> Decimal.to_float(amount_available)
    end
  end

  @spec create(ApplicationEntity.id(), PaymentAccount.currency()) :: PaymentAccount.id()
  def create(customer_id, currency) do
    [customer_id, currency]
    |> DB.query_val("SELECT create_payment_account($1, $2)")
  end

  defp from_atom({id, application_entity_id, amount, amount_reserved, currency}) do
    %PaymentAccount{
      id: id,
      application_entity_id: application_entity_id,
      amount: amount,
      amount_reserved: amount_reserved,
      amount_available: Decimal.sub(amount, amount_reserved),
      currency: currency
    }
  end
end
