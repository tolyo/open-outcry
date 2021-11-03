defmodule Instrument do
  @typedoc """
    `instrument.pub_id` db reference
  """
  @type id :: String.t()

  @typedoc """
    Ticker-like name of the instrument. For monetary instruments, a currency pair is used.
  """
  @type name :: String.t()

  @typedoc """
    The underlying currency of the FX instrument
  """
  @type base_currency() :: String.t()

  @type t :: %Instrument{
          id: id(),
          name: name(),
          base_currency: base_currency(),
          quote_currency: String.t(),
          active: boolean(),
          created_at: String.t(),
          updated_at: String.t()
        }

  defstruct id: nil,
            name: nil,
            base_currency: nil,
            quote_currency: nil,
            active: nil,
            created_at: nil,
            updated_at: nil
end
