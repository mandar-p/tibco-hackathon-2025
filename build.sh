#!/bin/bash

# TIBCOpilot Build Script

echo "Building TIBCOpilot modules..."

# Create bin directory if it doesn't exist
mkdir -p bin

# Build all modules
echo "Building API Server..."
go build -o bin/api-server ./cmd/api-server
if [ $? -ne 0 ]; then
    echo "Failed to build API server"
    exit 1
fi

echo "Building Executor..."
go build -o bin/executor ./cmd/executor
if [ $? -ne 0 ]; then
    echo "Failed to build executor"
    exit 1
fi

echo "Building Git Uploader..."
go build -o bin/git-uploader ./cmd/git-uploader
if [ $? -ne 0 ]; then
    echo "Failed to build git uploader"
    exit 1
fi

echo "Build complete! Binaries available in ./bin/"
echo ""
echo "To run the modules:"
echo "  ./bin/api-server     - Start REST API server"
echo "  ./bin/executor       - Start command executor"  
echo "  ./bin/git-uploader   - Run git uploader"
