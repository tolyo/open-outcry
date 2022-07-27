defmodule TradingAccountService do
  @spec create_trading_account_transfer(
          TradingAccount.id(),
          TradingAccount.id(),
          Instrument.name(),
          non_neg_integer()
        ) :: nil
  def create_trading_account_transfer(
        trading_account_from,
        trading_account_to,
        instrument_name,
        amount
      ) do
    [trading_account_from, trading_account_to, instrument_name, amount]
    |> DB.query("SELECT create_trading_account_transfer($1, $2, $3, $4)")
  end
end
