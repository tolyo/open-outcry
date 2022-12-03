funcmodule RegistrationService {
  require Logger

  @spec create_client(ApplicationEntity.external_id()) ApplicationEntity.id()
  func create_client(external_id) {
    external_id
    |> db.QueryVal("SELECT create_client($1)")
  }
}
