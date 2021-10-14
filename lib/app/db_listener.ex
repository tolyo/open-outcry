defmodule DBListener do
  @moduledoc """
    DB listener process for listening to incoming events, emitted by `Repo`.
  """

  use GenServer
  require Logger

  def child_spec(_) do
    %{
      id: __MODULE__,
      start: {__MODULE__, :start_link, []}
    }
  end

  @spec start_link(any) :: :ignore | {:error, any} | {:ok, pid}
  def start_link(opts \\ []), do: GenServer.start_link(__MODULE__, opts)

  @impl true
  @spec init(any) :: {:ok, any} | {:stop, {:error, any} | {:eventually, reference}}
  def init(opts) do
    case Repo.listen("db_event") do
      {:ok, _pid, _ref} ->
        {:ok, opts}

      error ->
        {:stop, error}
    end
  end

  @impl true
  def handle_info({:notification, _pid, _reference, "db_event", payload}, _state) do
    {:ok, message} = Jason.decode(payload)
    Logger.info(inspect(message))
    {:noreply, :event_handled}
  end
end
