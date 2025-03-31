#!/bin/bash

(cd ../ && go build -o protoc-gen-openapiv3)

protoc --plugin=protoc-gen-openapiv3=../protoc-gen-openapiv3 --openapiv3_out=output=./test.openapi.yaml,output-format=yaml:. --proto_path=./ ./test.proto
protoc --plugin=protoc-gen-openapiv3=../protoc-gen-openapiv3 --openapiv3_out=output=./test.v2.openapi.yaml,output-format=yaml:. --proto_path=./ ./test.v2.proto

# Sort and format YAML files consistently
# go install github.com/mikefarah/yq/v4@latest
yq eval 'sortKeys(..)' -i ./test.openapi.yaml
yq eval 'sortKeys(..)' -i ./test.v2.openapi.yaml
