defmodule Exchange.Repo do
  use Ecto.Repo,
    otp_app: :exchange,
    adapter: Ecto.Adapters.Postgres
end
