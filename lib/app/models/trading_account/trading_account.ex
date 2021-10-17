defmodule TradingAccount do
  @typedoc """
    `trading_account.pub_id` db reference
  """
  @type id :: String.t()

  @type t :: %TradingAccount{
          id: id(),
          application_entity_id: ApplicationEntity.id()
        }

  defstruct id: nil,
            application_entity_id: nil

  @base_query """
    SELECT (
      t.pub_id,
      ae.pub_id
    )

    FROM trading_account AS t

    INNER JOIN application_entity ae
          ON ae.id = t.application_entity_id

  """

  @spec get(TradingAccount.id()) :: TradingAccount.t()
  def get(id) do
    id
    |> DB.query_val(
      @base_query <>
        """
          WHERE t.pub_id = $1
        """
    )
    |> case do
      x -> from_atom(x)
    end
  end

  @spec find_by_application_entity_id(ApplicationEntity.id()) :: TradingAccount.t()
  def find_by_application_entity_id(application_entity_id) do
    application_entity_id
    |> DB.query_val(
      @base_query <>
        """
          WHERE ae.pub_id = $1
        """
    )
    |> case do
      x -> from_atom(x)
    end
  end

  @spec from_atom({String.t(), String.t()}) :: TradingAccount.t()
  def from_atom({id, application_entity_id}) do
    %TradingAccount{
      id: id,
      application_entity_id: application_entity_id
    }
  end
end
