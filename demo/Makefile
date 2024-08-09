# Define the build directory
BUILD_DIR = dist

# Run server in dev mode
serve:
	@npm run serve

# Run prettier source
pretty:
	@npx prettier . --write

# Build for production
build: clean_build
	@npm run build

# Clean build directory if it exists
clean_build:
	@if [ -d "$(BUILD_DIR)" ]; then \
		echo "Removing $(BUILD_DIR)..."; \
		rm -r "$(BUILD_DIR)"; \
	fi