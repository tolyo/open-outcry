defmodule RegistrationService do
  require Logger

  @spec create_client(ApplicationEntity.external_id()) :: ApplicationEntity.id()
  def create_client(external_id) do
    external_id
    |> DB.query_val("SELECT create_client($1)")
  end
end
