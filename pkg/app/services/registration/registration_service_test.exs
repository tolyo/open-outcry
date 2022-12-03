defmodule RegistrationServiceTest do
  use DataCase
  import DBTestUtils

  test "create_client/1" do
    # given
    application_entity_count = get_count("application_entity")
    payment_account_count = get_count("payment_account")

    # when
    pub_id = RegistrationService.create_client("test")

    # then pub id is returned
    assert pub_id != nil
    assert get_count("application_entity") == application_entity_count + 1
    assert get_count("payment_account") == payment_account_count + 1
  end
end
