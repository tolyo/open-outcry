defmodule DB do
  @moduledoc """
    Helper functions to working with db
  """

  require Logger

  @spec query_exists(any, binary) :: boolean
  def query_exists(val, content), do: query_val(val, content) !== nil

  @spec query_val(any, binary) :: any
  def query_val([a | b], content), do: process([a | b], content) |> extract_single()
  def query_val(arg, content), do: query([arg], content) |> extract_single()

  @spec query_val(binary) :: any
  def query_val(content), do: query([], content) |> extract_single()

  defp extract_single(nil), do: nil
  defp extract_single([]), do: nil
  defp extract_single([[val]]), do: val
  defp extract_single([[id, json, date1, date2]]), do: [id, json, date1, date2]

  @spec query_list(binary) :: [any]
  def query_list(content), do: query([], content) |> List.flatten()
  @spec query_list(any, binary) :: [any]
  def query_list(val, content), do: query(val, content) |> List.flatten()

  @spec query(binary) :: nil | [binary | [any]]
  def query(content), do: query([], content)

  @spec query(any, binary) :: nil | [binary | [any]]
  def query([], content), do: process([], content)
  def query([a | b], content), do: process([a | b], content)
  def query(arg, content), do: query([arg], content)

  @spec execute(binary) :: :ok
  def execute(content), do: execute([], content)
  @spec execute(any, binary) :: :ok
  def execute([], content), do: execute_helper([], content)
  def execute([a | b], content), do: execute_helper([a | b], content)
  def execute(a, content), do: execute_helper([a], content)

  defp execute_helper(arg_list, content) do
    res = Ecto.Adapters.SQL.query!(Repo, content, arg_list)
    Logger.info("#{inspect(res)}")
    :ok
  end

  defp process([], content), do: Ecto.Adapters.SQL.query!(Repo, content, []).rows
  defp process([a | b], content), do: Ecto.Adapters.SQL.query!(Repo, content, [a | b]).rows

  @spec migrate_up(any) :: [any]
  def migrate_up(name) do
    Logger.info("migrate up => #{name}")
    migrate_helper("priv/repo/migrations/up/#{name}")
  end

  @spec migrate_up(binary, binary) :: :ok
  def migrate_up(args, name) do
    Logger.info("migrate up => #{name}")
    migrate_command(args, "priv/repo/migrations/up/#{name}")
  end

  @spec migrate_down(any) :: [any]
  def migrate_down(name) do
    Logger.info("migrate down => #{name}")
    migrate_helper("priv/repo/migrations/down/#{name}")
  end

  @doc """
    Convenience method to upgrading models. Name param takes the name on the model
    and will resolve to upgrade a file residing as lib/app/models/{name}/{name}.sql.
    Example DB.migrate_up_table("user") will migrate up a table is file
    lib/app/models/user/user.sql
  """
  @spec migrate_up_table(any) :: [any]
  def migrate_up_table(name) do
    Logger.info("migrate_up_table=> #{name}")

    if MainApplication.prod() == true do
      migrate_helper("priv/repo/migrations/up/lib/app/models/#{name}/#{name}.sql")
    else
      case File.read("lib/app/models/#{name}/#{name}.sql") do
        {:ok, content} ->
          Logger.info(content)

          Enum.map(String.split(content, ";"), fn x ->
            Logger.info(x)
            execute(x)
          end)
      end
    end
  end

  @spec migrate_up_function(String.t()) :: :ok
  def migrate_up_function(name) do
    Logger.info("migrate up function=> #{name}")

    path =
      case MainApplication.prod() do
        true ->
          Path.wildcard(Application.app_dir(:server, "") <> "/priv/**/#{name}.sql")
          |> List.first()

        false ->
          Path.wildcard("./lib/**/#{name}.sql") |> List.first()
      end

    case File.read(path) do
      {:ok, content} ->
        Logger.info(content)
        execute(content)
    end
  end

  @spec migrate_up_type(any) :: [any]
  def migrate_up_type(name) do
    Logger.info("migrate up type=> #{name}")

    if MainApplication.prod() == true do
      migrate_helper("priv/repo/migrations/up/lib/app/models/#{name}/#{name}.sql")
    else
      case File.read("lib/app/models/#{name}/#{name}.sql") do
        {:ok, content} ->
          Logger.info(content)
          Enum.map(String.split(content, ";"), fn x -> execute(x) end)
      end
    end
  end

  @spec migrate_down_function(any) :: :ok
  def migrate_down_function(name) do
    Logger.info("migrate down function=> #{name}")
    execute("DROP FUNCTION IF EXISTS #{name};")
  end

  @spec migrate_down_table(any) :: :ok
  def migrate_down_table(name) do
    Logger.info("migrate down table=> #{name}")
    execute("DROP TABLE IF EXISTS  #{name} CASCADE;")
  end

  @spec migrate_down_type(any) :: :ok
  def migrate_down_type(name) do
    Logger.info("migrate_down_type down type=> #{name}")
    execute("DROP TYPE IF EXISTS  #{name} CASCADE;")
  end

  @spec migrate_helper(binary | [binary]) :: [any]
  def migrate_helper(url) do
    path = Application.app_dir(:server, url)
    {_, content} = File.read(path)
    Logger.info(content)

    content
    |> String.split("--\n")
    |> Enum.map(fn x ->
      Logger.info("Executing SQL command =>")
      Logger.info(x)
      execute(x)
    end)
  end

  @spec migrate_command(binary, binary) :: :ok
  def migrate_command(args, url) do
    path = Application.app_dir(:server, url)
    {_, content} = File.read(path)
    Logger.info(content)
    execute(args, content)
  end
end
