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

# Validate OpenAPI spec
echo "Validating OpenAPI spec..."
redocly lint ./testdata/test.openapi.yaml

echo "Test completed successfully!" 