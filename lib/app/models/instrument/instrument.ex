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

  @typedoc """
    The default currency for market quotes of the instrument
  """
  @type quote_currency() :: Currency.name()

  @type t :: %Instrument{
          id: id(),
          name: name(),
          base_currency: base_currency(),
          quote_currency: quote_currency(),
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

  @spec create_instrument(Instrument.name(), Currency.name()) :: Instrument.id()
  def create_instrument(name, quote_currency) do
    [name, quote_currency]
    |> DB.query_val("""
      INSERT INTO instrument(
        name,
        quote_currency
      ) VALUES (
        $1,
        $2
      )
      RETURNING pub_id;
    """)
  end

  @spec create_fx_instrument(Instrument.name(), Currency.name(), Currency.name()) ::
          Instrument.id()
  def create_fx_instrument(name, base_currency, quote_currency) do
    [name, base_currency, quote_currency]
    |> DB.query_val("""
    INSERT INTO instrument(
      name,
      base_currency,
      quote_currency,
      fx_instrument
    ) VALUES (
      $1,
      $2,
      $3,
      TRUE
    )
    RETURNING pub_id;
    """)
  end
end
