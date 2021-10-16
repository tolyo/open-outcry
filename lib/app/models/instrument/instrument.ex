defmodule Instrument do
  @typedoc """
    `instrument.pub_id` db reference
  """
  @type id :: String.t()

  @typedoc """
    Currency pair of the instument
  """
  @type name :: String.t()

  defstruct pub_id: nil,
            name: nil,
            base_currency: nil,
            quote_currency: nil,
            active: nil,
            created_at: nil,
            updated_at: nil
end
