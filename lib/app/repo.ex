defmodule Repo do
  @moduledoc """
    DB repository with even listenert hook, allowing subscribers to be notified of database events
  """
  use Ecto.Repo,
    otp_app: :exchange,
    adapter: Ecto.Adapters.Postgres

  @spec listen(String.t()) :: {:error, any} | {:eventually, reference} | {:ok, pid, reference}
  def listen(event_name) do
    with {:ok, pid} <- Postgrex.Notifications.start_link(__MODULE__.config()),
         {:ok, ref} <- Postgrex.Notifications.listen(pid, event_name) do
      {:ok, pid, ref}
    end
  end
end
