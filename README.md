# protoc-gen-openapiv3

A Protocol Buffers plugin that generates OpenAPI v3 specifications from your proto files. This tool serves as a drop-in replacement for `protoc-gen-openapiv2` from grpc-gateway, providing support for OpenAPI v3 specifications.

## Overview

This tool generates OpenAPI v3 (formerly known as Swagger) specifications from your Protocol Buffer definitions. It's designed to be a direct replacement for the OpenAPI v2 generator in the grpc-gateway ecosystem, offering enhanced features and compatibility with the latest OpenAPI specification.

## Features

- Generates OpenAPI v3 specifications from Protocol Buffer files
- Compatible with existing grpc-gateway annotations
- Supports all OpenAPI v3 features
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

service YourService {
  rpc YourMethod(YourRequest) returns (YourResponse) {
    option (google.api.http) = {
      get: "/v1/your-method"
    };
  }
}
```

2. Generate the OpenAPI specification:

```bash
protoc -I . \
  --openapiv3_out=. \
  your.proto
```

## Configuration

The generator supports various options that can be passed through protoc:

- `allow_merge`: Enable merging of OpenAPI specifications
- `include_package_in_tags`: Include package name in operation tags
- `fqn_for_openapi_name`: Use fully qualified names for OpenAPI names
- `openapi_configuration`: Path to OpenAPI configuration file

Example with options:

```bash
protoc -I . \
  --openapiv3_out=allow_merge=true,include_package_in_tags=true:. \
  your.proto
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Acknowledgments

- Inspired by [grpc-gateway](https://github.com/grpc-ecosystem/grpc-gateway)
- Built on top of the Protocol Buffers ecosystem 