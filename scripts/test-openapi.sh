#!/bin/bash

set -e

# Build the plugin
echo "Building protoc-gen-openapiv3..."
go build -o protoc-gen-openapiv3

# Generate OpenAPI spec
echo "Generating OpenAPI spec..."
protoc --plugin=protoc-gen-openapiv3=./protoc-gen-openapiv3 \
  --openapiv3_out=output=./testdata/test.openapi.yaml,output-format=yaml:. \
  --proto_path=./testdata \
  ./testdata/test.proto

# Install swagger-cli if not already installed
if ! command -v swagger-cli &> /dev/null; then
  echo "Installing swagger-cli..."
  sudo npm install -g @apidevtools/swagger-cli
fi

# Validate OpenAPI spec
echo "Validating OpenAPI spec..."
swagger-cli validate ./testdata/test.openapi.yaml

echo "Test completed successfully!" 