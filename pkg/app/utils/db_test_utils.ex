defmodule DBTestUtils do
  @spec get_count(String.t()) :: number()
  def get_count(table_name) do
    DB.query_val("SELECT COUNT(*) FROM #{table_name}")
  end
end
