defmodule Currency do
  @type name :: String.t()

  @type precision :: Decimal.t()

  @type t :: %Currency{
          name: name(),
          precision: precision()
        }

  defstruct name: nil,
            precision: 2

  @base_query """
    SELECT(
        c.name,
        c.precision
    )

    FROM currency AS c
  """

  @spec exists(Currency.name()) :: boolean()
  def exists(name) do
    name
    |> DB.query_exists(@base_query <> "WHERE c.name = $1")
  end
end
