package main

import (
	"fmt"
	"os"
)

func Example() {
	// This example demonstrates how to use the OpenAPI v3 generator
	// with the test proto file in testdata/test.proto.

	err := flags.Parse([]string{
		"--allow_merge=true",
		"--include_package_in_tags=false",
		"--fqn_for_openapi_name=true",
		"--output=./testdata/test.yaml",
	})
	if err != nil {
		fmt.Println("Error parsing flags:", err)
		os.Exit(1)
	}

	main()

	// Output:
	// {
	//   "openapi": "3.0.0",
	//   "info": {
	//     "title": "Test API",
	//     "version": "1.0.0",
	//     "description": "API generated from test.proto"
	//   },
	//   "paths": {
	//     "/v1/users": {
	//       "get": {
	//         "operationId": "ListUsers",
	//         "summary": "List users with pagination",
	//         "parameters": [
	//           {
	//             "name": "page_size",
	//             "in": "query",
	//             "description": "Number of items per page",
	//             "schema": {
	//               "type": "integer"
	//             }
	//           },
	//           {
	//             "name": "page_token",
	//             "in": "query",
	//             "description": "Token for the next page",
	//             "schema": {
	//               "type": "string"
	//             }
	//           },
	//           {
	//             "name": "filter",
	//             "in": "query",
	//             "description": "Filter criteria",
	//             "schema": {
	//               "type": "string"
	//             }
	//           },
	//           {
	//             "name": "order_by",
	//             "in": "query",
	//             "description": "Order by criteria",
	//             "schema": {
	//               "type": "string"
	//             }
	//           }
	//         ],
	//         "responses": {
	//           "200": {
	//             "description": "List of users",
	//             "content": {
	//               "application/json": {
	//                 "schema": {
	//                   "$ref": "#/components/schemas/ListUsersResponse"
	//                 }
	//               }
	//             }
	//           }
	//         }
	//       },
	//       "post": {
	//         "operationId": "CreateUser",
	//         "summary": "Create a new user",
	//         "requestBody": {
	//           "required": true,
	//           "content": {
	//             "application/json": {
	//               "schema": {
	//                 "$ref": "#/components/schemas/CreateUserRequest"
	//               }
	//             }
	//           }
	//         },
	//         "responses": {
	//           "200": {
	//             "description": "Created user",
	//             "content": {
	//               "application/json": {
	//                 "schema": {
	//                   "$ref": "#/components/schemas/User"
	//                 }
	//               }
	//             }
	//           }
	//         }
	//       }
	//     },
	//     "/v1/users/{user_id}": {
	//       "get": {
	//         "operationId": "GetUser",
	//         "summary": "Get user by ID",
	//         "parameters": [
	//           {
	//             "name": "user_id",
	//             "in": "path",
	//             "required": true,
	//             "schema": {
	//               "type": "string"
	//             }
	//           }
	//         ],
	//         "responses": {
	//           "200": {
	//             "description": "User details",
	//             "content": {
	//               "application/json": {
	//                 "schema": {
	//                   "$ref": "#/components/schemas/User"
	//                 }
	//               }
	//             }
	//           }
	//         }
	//       },
	//       "patch": {
	//         "operationId": "UpdateUser",
	//         "summary": "Update user",
	//         "parameters": [
	//           {
	//             "name": "user_id",
	//             "in": "path",
	//             "required": true,
	//             "schema": {
	//               "type": "string"
	//             }
	//           }
	//         ],
	//         "requestBody": {
	//           "required": true,
	//           "content": {
	//             "application/json": {
	//               "schema": {
	//                 "$ref": "#/components/schemas/UpdateUserRequest"
	//               }
	//             }
	//           }
	//         },
	//         "responses": {
	//           "200": {
	//             "description": "Updated user",
	//             "content": {
	//               "application/json": {
	//                 "schema": {
	//                   "$ref": "#/components/schemas/User"
	//                 }
	//               }
	//             }
	//           }
	//         }
	//       },
	//       "delete": {
	//         "operationId": "DeleteUser",
	//         "summary": "Delete user",
	//         "parameters": [
	//           {
	//             "name": "user_id",
	//             "in": "path",
	//             "required": true,
	//             "schema": {
	//               "type": "string"
	//             }
	//           }
	//         ],
	//         "responses": {
	//           "204": {
	//             "description": "User deleted successfully"
	//           }
	//         }
	//       }
	//     }
	//   },
	//   "components": {
	//     "schemas": {
	//       "User": {
	//         "type": "object",
	//         "properties": {
	//           "user_id": {
	//             "type": "string"
	//           },
	//           "email": {
	//             "type": "string"
	//           },
	//           "name": {
	//             "type": "string"
	//           },
	//           "type": {
	//             "type": "string",
	//             "enum": ["USER_TYPE_UNSPECIFIED", "USER_TYPE_ADMIN", "USER_TYPE_USER", "USER_TYPE_SERVICE"]
	//           },
	//           "created_at": {
	//             "type": "string",
	//             "format": "date-time"
	//           },
	//           "updated_at": {
	//             "type": "string",
	//             "format": "date-time"
	//           },
	//           "metadata": {
	//             "type": "object",
	//             "additionalProperties": {
	//               "type": "string"
	//             }
	//           },
	//           "tags": {
	//             "type": "array",
	//             "items": {
	//               "type": "string"
	//             }
	//           }
	//         }
	//       },
	//       "ListUsersResponse": {
	//         "type": "object",
	//         "properties": {
	//           "users": {
	//             "type": "array",
	//             "items": {
	//               "$ref": "#/components/schemas/User"
	//             }
	//           },
	//           "next_page_token": {
	//             "type": "string"
	//           },
	//           "total_size": {
	//             "type": "integer"
	//           }
	//         }
	//       },
	//       "CreateUserRequest": {
	//         "type": "object",
	//         "properties": {
	//           "user": {
	//             "$ref": "#/components/schemas/User"
	//           }
	//         }
	//       },
	//       "UpdateUserRequest": {
	//         "type": "object",
	//         "properties": {
	//           "user_id": {
	//             "type": "string"
	//           },
	//           "user": {
	//             "$ref": "#/components/schemas/User"
	//           }
	//         }
	//       }
	//     }
	//   }
	// }
}
