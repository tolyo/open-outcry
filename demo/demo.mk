# Define the build directory
BUILD_DIR = api/dist
NPM_PREFIX = --prefix demo

setup-demo:
	@npm $(NPM_PREFIX) i
	@go get ./...

# Run server in dev mode
serve-demo:
	@go run main.go

# Run prettier source
pretty-demo:
	@npx $(NPM_PREFIX) prettier . --write

# Build for production
build-demo: clean-demo
	@npm $(NPM_PREFIX) run build

# Clean build directory if it exists
clean-demo:
	@if [ -d "$(BUILD_DIR)" ]; then \
		echo "Removing $(BUILD_DIR)..."; \
		rm -r "$(BUILD_DIR)"; \
	fi