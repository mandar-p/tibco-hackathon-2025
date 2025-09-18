#!/bin/bash

# TIBCOpilot Test Script

echo "Testing TIBCOpilot setup..."

# Check if Go is installed
if ! command -v go &> /dev/null; then
    echo "❌ Go is not installed or not in PATH"
    exit 1
else
    echo "✅ Go is installed: $(go version)"
fi

# Check if config file exists
if [ ! -f "config/config.json" ]; then
    echo "❌ Configuration file not found: config/config.json"
    echo "   Please copy and configure config/config.json from the template"
    exit 1
else
    echo "✅ Configuration file found"
fi

# Check if required directories exist
if [ ! -d "data" ]; then
    echo "❌ Data directory not found"
    exit 1
else
    echo "✅ Data directory exists"
fi

# Try to build the modules
echo "Building modules..."
go build -o /tmp/test-api ./cmd/api-server
if [ $? -eq 0 ]; then
    echo "✅ API server builds successfully"
    rm -f /tmp/test-api
else
    echo "❌ API server build failed"
    exit 1
fi

go build -o /tmp/test-executor ./cmd/executor
if [ $? -eq 0 ]; then
    echo "✅ Executor builds successfully"
    rm -f /tmp/test-executor
else
    echo "❌ Executor build failed"
    exit 1
fi

go build -o /tmp/test-git ./cmd/git-uploader
if [ $? -eq 0 ]; then
    echo "✅ Git uploader builds successfully"
    rm -f /tmp/test-git
else
    echo "❌ Git uploader build failed"
    exit 1
fi

echo ""
echo "🎉 All tests passed! TIBCOpilot is ready to use."
echo ""
echo "Next steps:"
echo "1. Update your config/config.json with correct values"
echo "2. Run 'make build' or './build.sh' to build binaries"
echo "3. Start the API server: ./bin/api-server"
echo "4. Start the executor: ./bin/executor"
echo "5. Use the REST API to generate commands"
