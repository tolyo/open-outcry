# API management make

INPUT_YAML:="./api/openapi.yaml"
OUTPUT_YAML:="../pkg/static/api.yaml"

install-api:
	npm i

bundle-api:
	npx @redocly/cli bundle \
		$(INPUT_YAML) \
		--output $(OUTPUT_YAML)

validate-api:
	make bundle-api
	npx @redocly/cli lint \
		$(OUTPUT_YAML) \
		--format=checkstyle

generate-api: ## Generate server bindings, move model files, fix imports
	make validate-api
	npx @openapitools/openapi-generator-cli generate \
		-i $(OUTPUT_YAML) \
		-g go-server \
		-o pkg/rest \
		--additional-properties=packageName=api \
		--additional-properties=sourceFolder=api \
		--additional-properties=outputAsLibrary=true
	make lint