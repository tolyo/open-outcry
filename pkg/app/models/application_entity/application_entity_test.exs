defmodule ApplicationEntityTest do
  use DataCase

  describe "master entity" do
    test "expect masster entity to exist" do
      assert DBTestUtils.get_count("application_entity") == 1

      assert DB.query_val("SELECT (pub_id) FROM application_entity WHERE type = 'MASTER'") ==
               "MASTER"
    end
  end
end
