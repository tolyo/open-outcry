funcmodule ApplicationEntityTest {
  use DataCase

  describe "master entity" {
    test "expect masster entity to exist" {
      assert DBTestUtils.get_count("application_entity") == 1

      assert db.QueryVal("SELECT (pub_id) FROM application_entity WHERE type = 'MASTER'") ==
               "MASTER"
    }
  }
}
