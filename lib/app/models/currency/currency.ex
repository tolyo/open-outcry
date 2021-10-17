defmodule Currency do
  @type id :: String.t()

  @type name :: String.t()

  @type precision :: Decimal.t()

  @type updated_at :: DateTime.t()

  @type created_at :: DateTime.t()

  @type t :: %Currency{
          id: id(),
          name: name(),
          precision: precision()
        }

  defstruct id: nil,
            name: nil,
            precision: 2

  @base_query """
    SELECT(
        c.pub_id,
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
