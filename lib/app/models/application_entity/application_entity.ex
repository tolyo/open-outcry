defmodule ApplicationEntity do
  @moduledoc """
    Application entity is any generic enity capable of being an actor in financial transaction
  """

  @typedoc """
    `application_entity.pub_id` db reference
  """
  @type id :: String.t()

  @typedoc """
    `application_entity.external_id field
  """
  @type external_id :: String.t()

  @typedoc """
    Type of application entity
  """
  @type type :: :CLIENT | :MASTER

  @type t :: %ApplicationEntity{
          id: id(),
          type: type(),
          external_id: external_id()
        }

  defstruct id: nil,
            type: nil,
            external_id: nil
end
