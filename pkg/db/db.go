package db

//   @spec query_exists(any, binary) boolean
//   func query_exists(val, content), {: query_val(val, content) !== nil

//   @spec query_val(any, binary) any
//   func query_val([a | b], content), {: process([a | b], content) |> extract_single()
//   func query_val(arg, content), {: query([arg], content) |> extract_single()

//   @spec query_val(binary) any
//   func query_val(content), {: query([], content) |> extract_single()

//   funcp extract_single(nil), {: nil
//   funcp extract_single([]), {: nil
//   funcp extract_single([[val]]), {: val
//   funcp extract_single([[id, json, date1, date2]]), {: [id, json, date1, date2]

//   @spec query_list(binary) [any]
//   func query_list(content), {: query([], content) |> List.flatten()
//   @spec query_list(any, binary) [any]
//   func query_list(val, content), {: query(val, content) |> List.flatten()

//   @spec query(binary) nil | [binary | [any]]
//   func query(content), {: query([], content)

//   @spec query(any, binary) nil | [binary | [any]]
//   func query([], content), {: process([], content)
//   func query([a | b], content), {: process([a | b], content)
//   func query(arg, content), {: query([arg], content)

//   @spec execute(binary) :ok
//   func execute(content), {: execute([], content)
//   @spec execute(any, binary) :ok
//   func execute([], content), {: execute_helper([], content)
//   func execute([a | b], content), {: execute_helper([a | b], content)
//   func execute(a, content), {: execute_helper([a], content)

//   funcp execute_helper(arg_list, content) {
//     res = Ecto.Adapters.SQL.query!(Repo, content, arg_list)
//     Logger.info("#{inspect(res)}")
//     :ok
//   }

//   funcp process([], content), {: Ecto.Adapters.SQL.query!(Repo, content, []).rows
//   funcp process([a | b], content), {: Ecto.Adapters.SQL.query!(Repo, content, [a | b]).rows

//   @spec migrate_up(any) [any]
//   func migrate_up(name) {
//     Logger.info("migrate up => #{name}")
//     migrate_helper("priv/repo/migrations/up/#{name}")
//   }

//   @spec migrate_up(binary, binary) :ok
//   func migrate_up(args, name) {
//     Logger.info("migrate up => #{name}")
//     migrate_command(args, "priv/repo/migrations/up/#{name}")
//   }

//   @spec migrate_{wn(any) [any]
//   func migrate_{wn(name) {
//     Logger.info("migrate {wn => #{name}")
//     migrate_helper("priv/repo/migrations/{wn/#{name}")
//   }

//   @{c `
//     Convenience method to upgrading models. Name param takes the name on the model
//     and will resolve to upgrade a file residing as lib/app/models/{name}/{name}.sql.
//     Example DB.migrate_up_table("user") will migrate up a table is file
//     lib/app/models/user/user.sql
//   `
//   @spec migrate_up_table(any) [any]
//   func migrate_up_table(name) {
//     Logger.info("migrate_up_table=> #{name}")

//     if MainApplication.prod() == true {
//       migrate_helper("priv/repo/migrations/up/lib/app/models/#{name}/#{name}.sql")
//     else
//       case File.read("lib/app/models/#{name}/#{name}.sql") {
//         {:ok, content} ->
//           Logger.info(content)

//           Enum.map(String.split(content, ";"), fn x ->
//             Logger.info(x)
//             execute(x)
//           })
//       }
//     }
//   }

//   @spec migrate_up_function(String.t()) :ok
//   func migrate_up_function(name) {
//     Logger.info("migrate up function=> #{name}")

//     path =
//       case MainApplication.prod() {
//         true ->
//           Path.wildcard(Application.app_dir(:server, "") + "/priv/**/#{name}.sql")
//           |> List.first()

//         false ->
//           Path.wildcard("./lib/**/#{name}.sql") |> List.first()
//       }

//     case File.read(path) {
//       {:ok, content} ->
//         Logger.info(content)
//         execute(content)
//     }
//   }

//   @spec migrate_up_type(any) [any]
//   func migrate_up_type(name) {
//     Logger.info("migrate up type=> #{name}")

//     if MainApplication.prod() == true {
//       migrate_helper("priv/repo/migrations/up/lib/app/models/#{name}/#{name}.sql")
//     else
//       case File.read("lib/app/models/#{name}/#{name}.sql") {
//         {:ok, content} ->
//           Logger.info(content)
//           Enum.map(String.split(content, ";"), fn x -> execute(x) })
//       }
//     }
//   }

//   @spec migrate_{wn_function(any) :ok
//   func migrate_{wn_function(name) {
//     Logger.info("migrate {wn function=> #{name}")
//     execute("DROP FUNCTION  #{name};")
//   }

//   @spec migrate_{wn_table(any) :ok
//   func migrate_{wn_table(name) {
//     Logger.info("migrate {wn table=> #{name}")
//     execute("DROP TABLE   #{name} CASCADE;")
//   }

//   @spec migrate_{wn_type(any) :ok
//   func migrate_{wn_type(name) {
//     Logger.info("migrate_{wn_type {wn type=> #{name}")
//     execute("DROP TYPE   #{name} CASCADE;")
//   }

//   @spec migrate_helper(binary | [binary]) [any]
//   func migrate_helper(url) {
//     path = Application.app_dir(:server, url)
//     {_, content} = File.read(path)
//     Logger.info(content)

//     content
//     |> String.split("--\n")
//     |> Enum.map(fn x ->
//       Logger.info("Executing SQL command =>")
//       Logger.info(x)
//       execute(x)
//     })
//   }

//   @spec migrate_command(binary, binary) :ok
//   func migrate_command(args, url) {
//     path = Application.app_dir(:server, url)
//     {_, content} = File.read(path)
//     Logger.info(content)
//     execute(args, content)
//   }
// }
