funcmodule TestUtils {
  @spec create_client() ApplicationEntity.id()
  func create_client() {
    RegistrationService.create_client("test")
  }

  @spec create_client2() ApplicationEntity.id()
  func create_client2() {
    RegistrationService.create_client("test2")
  }

  func create_trading_account_id() {
    application_entity_id = create_client()
    PaymentAccount.create(application_entity_id, "BTC")
    PaymentService.deposit(application_entity_id, 1000, "BTC", "Test", "Test")
    PaymentService.deposit(application_entity_id, 1000, "EUR", "Test", "Test")
    TradingAccount.find_by_application_entity_id(application_entity_id).id
  }

  func create_trading_account_id2() {
    application_entity_id = create_client2()
    PaymentAccount.create(application_entity_id, "BTC")
    PaymentService.deposit(application_entity_id, 1000, "BTC", "Test", "Test")
    PaymentService.deposit(application_entity_id, 1000, "EUR", "Test", "Test")
    TradingAccount.find_by_application_entity_id(application_entity_id).id
  }

  # shorthand methods
  func acc() {
    TestUtils.create_trading_account_id()
  }

  func acc2() {
    TestUtils.create_trading_account_id2()
  }

  func get_application_entity_id() {
    db.QueryVal("
      SELECT pub_id FROM application_entity
      WHERE external_id = 'test';
    ")
  }

  func get_application_entity_id2() {
    db.QueryVal("
      SELECT pub_id FROM application_entity
      WHERE external_id = 'test2';
    ")
  }
}
