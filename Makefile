# TIBCOpilot Makefile

# Build all modules
build:
	@echo "Building all modules..."
	@go build -o bin/api-server ./cmd/api-server
	@go build -o bin/executor ./cmd/executor
	@go build -o bin/git-uploader ./cmd/git-uploader
	@echo "Build complete. Binaries are in ./bin/"

# Build individual modules
build-api:
	@go build -o bin/api-server ./cmd/api-server

build-executor:
	@go build -o bin/executor ./cmd/executor

build-git:
	@go build -o bin/git-uploader ./cmd/git-uploader

# Run modules
run-api:
	@go run ./cmd/api-server

run-executor:
	@go run ./cmd/executor

run-git:
	@go run ./cmd/git-uploader

# Setup directories
setup:
	@mkdir -p bin data
	@touch data/commands.txt data/execution.log

# Clean build artifacts
clean:
	@rm -rf bin/
	@echo "Clean complete."

# Install dependencies
deps:
	@go mod tidy
	@go mod download

# Test the API
test-api:
	@curl -X POST http://localhost:8080/api/v1/generate-commands \
		-H "Content-Type: application/json" \
		-d '{"user_prompt": "Give me commands to create a new bwdesign workspace", "api_url": "https://api.anthropic.com/v1/messages", "api_key": "your-api-key", "model_name": "claude-sonnet-4-20250514"}'

# Show help
help:
	@echo "Available targets:"
	@echo "  build         - Build all modules"
	@echo "  build-api     - Build API server only"
	@echo "  build-executor - Build executor only" 
	@echo "  build-git     - Build git uploader only"
	@echo "  run-api       - Run API server"
	@echo "  run-executor  - Run command executor"
	@echo "  run-git       - Run git uploader"
	@echo "  setup         - Create required directories"
	@echo "  clean         - Remove build artifacts"
	@echo "  deps          - Install dependencies"
	@echo "  test-api      - Test API endpoint"
	@echo "  help          - Show this help"

.PHONY: build build-api build-executor build-git run-api run-executor run-git setup clean deps test-api help