import Config

# Configure your database
#
# The MIX_TEST_PARTITION environment variable can be used
# to provide built-in test partitioning in CI environment.
# Run `mix help test` for more information.
config :exchange, Exchange.Repo,
  username: (System.get_env("POSTGRES_USER") || "postgres"),
  password: (System.get_env("POSTGRES_PASSWORD") || "postgres"),
  database: (System.get_env("POSTGRES_DB") || "exchange_test_db"),
  hostname: (System.get_env("POSTGRES_HOST") || "localhost"),
  pool: Ecto.Adapters.SQL.Sandbox,
  pool_size: 10

# We don't run a server during test. If one is required,
# you can enable the server option below.
config :exchange, ExchangeWeb.Endpoint,
  http: [ip: {127, 0, 0, 1}, port: 4002],
  secret_key_base: "ihnbR0w3oGs88eCtuOFo/5oQ69J3hoDRoy2lvm/3Pq/tnr29PJdCc3n0ZH/vBF/B",
  server: false

# In test we don't send emails.
config :exchange, Exchange.Mailer, adapter: Swoosh.Adapters.Test

# Print only warnings and errors during test
config :logger, level: :warn

# Initialize plugs at runtime for faster test compilation
config :phoenix, :plug_init_mode, :runtime
