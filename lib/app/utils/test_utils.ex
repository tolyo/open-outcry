defmodule TestUtils do
  @spec create_client() :: ApplicationEntity.id()
  def create_client() do
    RegistrationService.create_client("test")
  end

  @spec create_client2() :: ApplicationEntity.id()
  def create_client2() do
    RegistrationService.create_client("test2")
  end

  def create_trading_account_id() do
    application_entity_id = create_client()
    PaymentAccount.create(application_entity_id, "BTC")
    PaymentService.deposit(application_entity_id, 1000, "BTC", "Test", "Test")
    PaymentService.deposit(application_entity_id, 1000, "EUR", "Test", "Test")
    TradingAccount.find_by_application_entity_id(application_entity_id).id
  end

  def create_trading_account_id2() do
    application_entity_id = create_client2()
    PaymentAccount.create(application_entity_id, "BTC")
    PaymentService.deposit(application_entity_id, 1000, "BTC", "Test", "Test")
    PaymentService.deposit(application_entity_id, 1000, "EUR", "Test", "Test")
    TradingAccount.find_by_application_entity_id(application_entity_id).id
  end

  # shorthand methods
  def acc() do
    TestUtils.create_trading_account_id()
  end

  def acc2() do
    TestUtils.create_trading_account_id2()
  end

  def get_application_entity_id() do
    DB.query_val("
      SELECT pub_id FROM application_entity
      WHERE external_id = 'test';
    ")
  end

  def get_application_entity_id2() do
    DB.query_val("
      SELECT pub_id FROM application_entity
      WHERE external_id = 'test2';
    ")
  end
end
