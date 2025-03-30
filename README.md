# protoc-gen-openapiv3

A Protocol Buffers plugin that generates OpenAPI v3 specifications from your proto files. This tool serves as a drop-in replacement for `protoc-gen-openapiv2` from grpc-gateway, providing support for OpenAPI v3 specifications.

## Overview

This tool generates OpenAPI v3 (formerly known as Swagger) specifications from your Protocol Buffer definitions. It's designed to be a direct replacement for the OpenAPI v2 generator in the grpc-gateway ecosystem, offering enhanced features and compatibility with the latest OpenAPI specification.

## Features

- Generates OpenAPI v3 specifications from Protocol Buffer files
- Compatible with existing grpc-gateway annotations
- Supports OpenAPI v3 features including:
  - Response schemas and references
  - Security schemes (OAuth2, API Key, HTTP)
  - Server configurations
  - Request/Response content types
  - Schema components and references
- Drop-in replacement for protoc-gen-openapiv2
- Maintains backward compatibility with existing proto files

## Installation

```bash
go install github.com/sapk/protoc-gen-openapiv3@latest
```

## Usage

1. Add the following to your proto files:

```protobuf
syntax = "proto3";

package your.package;

import "google/api/annotations.proto";
import "options/annotations.proto";

// Define API information
option (protoc_gen_openapiv3.options.info) = {
  title: "Your API"
  description: "API description"
  version: "1.0.0"
};

// Define security schemes
option (protoc_gen_openapiv3.options.securityScheme) = {
  type: "oauth2"
  description: "OAuth2 authentication"
  flows: {
    authorization_code: {
      authorization_url: "https://auth.example.com/oauth/authorize"
      token_url: "https://auth.example.com/oauth/token"
      scopes: {
        name: "read"
        description: "Read access"
      }
    }
  }
};

service YourService {
  rpc YourMethod(YourRequest) returns (YourResponse) {
    option (google.api.http) = {
      get: "/v1/your-method"
    };
    option (protoc_gen_openapiv3.options.operation) = {
      summary: "Get something"
      description: "Detailed description"
      responses: {
        code: "200"
        description: "Success"
        content: {
          key: "application/json"
          value: {
            schema: {
              ref: "#/components/schemas/YourResponse"
            }
          }
        }
      }
      responses: {
        code: "400"
        description: "Bad Request"
        content: {
          key: "application/json"
          value: {
            schema: {
              ref: "#/components/schemas/Error"
            }
          }
        }
      }
    };
  }
}
```

2. Generate the OpenAPI specification:

```bash
go build -o protoc-gen-openapiv3 && protoc --plugin=protoc-gen-openapiv3=./protoc-gen-openapiv3   --openapiv3_out=output=./testdata/test.openapi.yaml,output-format=yaml:. --proto_path=./testdata --proto_path=./ ./testdata/test.proto
```

## Configuration

The generator supports various options that can be passed through protoc:

- `allow_merge`: Enable merging of OpenAPI specifications
- `include_package_in_tags`: Include package name in operation tags
- `fqn_for_openapi_name`: Use fully qualified names for OpenAPI names
- `openapi_configuration`: Path to OpenAPI configuration file

Example with options:

```bash
go build -o protoc-gen-openapiv3 && protoc --openapiv3_out=output=./testdata/test.openapi.json,output-format=json,allow_merge=true,include_package_in_tags=true:. --plugin=protoc-gen-openapiv3=./protoc-gen-openapiv3 --proto_path=./testdata --proto_path=./ ./testdata/test.proto 
```

## Schema Handling

The generator automatically handles schema references and components:

- All message types used in requests and responses are automatically added to the components section
- Response schemas that reference message types (like error responses) are properly included
- Schema references are resolved and the corresponding components are generated
- Support for primitive types, arrays, maps, and nested objects

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Acknowledgments

- Inspired by [grpc-gateway](https://github.com/grpc-ecosystem/grpc-gateway)
- Built on top of the Protocol Buffers ecosystem 