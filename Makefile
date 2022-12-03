compile:
	mix do deps.get, deps.compile

# run: Start dev mode
run:
	@go run main.go

db-update:
	MIX_ENV=dev mix ecto.migrate
	MIX_ENV=test mix ecto.migrate

db-downgrade:
	MIX_ENV=dev mix ecto.rollback --all
	MIX_ENV=test mix ecto.rollback --all

# Helper for executing a hard reset on the database
db-rebuild: db-downgrade db-update

lint:
	mix format mix.exs 'lib/**/*.{ex,exs}' 'test/**/*.{ex,exs}'
	
check:
	mix dialyzer --format dialyzer

check-coverage:
	MIX_ENV=test mix coveralls ./lib/

run-test:
	MIX_ENV=test mix test test/ lib/

run-quality-check: lint check run-test check-coverage

style-check:
	mix credo --all 


build-api:
	node node_modules/swagger-cli/swagger-cli.js bundle -o static/docs/api/openapi.json -t json -r lib/web/openapi/openapi.yaml


