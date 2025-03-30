package generator_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sapk/protoc-gen-openapiv3/generator"
	"github.com/sapk/protoc-gen-openapiv3/options"
)

func TestConvertToOpenAPI(t *testing.T) {
	// Create a test ParsedFile
	parsedFile := &generator.ParsedFile{
		Package: "test.package",
		Services: []generator.ParsedService{
			{
				Name: "TestService",
				Methods: []generator.ParsedMethod{
					{
						Name:       "GetUser",
						InputType:  "test.package.GetUserRequest",
						OutputType: "test.package.User",
						HTTPMethod: "GET",
						HTTPPath:   "/v1/users/{user_id}",
						Operation: &options.Operation{
							Summary: "GetUser operation",
						},
					},
					{
						Name:       "CreateUser",
						InputType:  "test.package.CreateUserRequest",
						OutputType: "test.package.User",
						HTTPMethod: "POST",
						HTTPPath:   "/v1/users",
						HTTPBody:   "user",
						Operation: &options.Operation{
							Summary: "CreateUser operation",
						},
					},
					{
						Name:       "UpdateUser",
						InputType:  "test.package.UpdateUserRequest",
						OutputType: "test.package.User",
						HTTPMethod: "PUT",
						HTTPPath:   "/v1/users/{user_id}",
						HTTPBody:   "user",
						Operation: &options.Operation{
							Summary: "UpdateUser operation",
						},
					},
					{
						Name:       "PatchUser",
						InputType:  "test.package.PatchUserRequest",
						OutputType: "test.package.User",
						HTTPMethod: "PATCH",
						HTTPPath:   "/v1/users/{user_id}",
						HTTPBody:   "user",
						Operation: &options.Operation{
							Summary: "PatchUser operation",
						},
					},
					{
						Name:       "DeleteUser",
						InputType:  "test.package.DeleteUserRequest",
						OutputType: "google.protobuf.Empty",
						HTTPMethod: "DELETE",
						HTTPPath:   "/v1/users/{user_id}",
						Operation: &options.Operation{
							Summary: "DeleteUser operation",
						},
					},
				},
			},
		},
		Messages: []generator.ParsedMessage{
			{
				Name: "GetUserRequest",
				Fields: []generator.ParsedField{
					{
						Name:   "user_id",
						Type:   "string",
						Number: 1,
					},
				},
			},
			{
				Name: "CreateUserRequest",
				Fields: []generator.ParsedField{
					{
						Name:   "user",
						Type:   "User",
						Number: 1,
					},
				},
			},
			{
				Name: "UpdateUserRequest",
				Fields: []generator.ParsedField{
					{
						Name:   "user_id",
						Type:   "string",
						Number: 1,
					},
					{
						Name:   "user",
						Type:   "User",
						Number: 2,
					},
				},
			},
			{
				Name: "PatchUserRequest",
				Fields: []generator.ParsedField{
					{
						Name:   "user_id",
						Type:   "string",
						Number: 1,
					},
					{
						Name:   "user",
						Type:   "User",
						Number: 2,
					},
				},
			},
			{
				Name: "DeleteUserRequest",
				Fields: []generator.ParsedField{
					{
						Name:   "user_id",
						Type:   "string",
						Number: 1,
					},
				},
			},
			{
				Name: "User",
				Fields: []generator.ParsedField{
					{
						Name:   "user_id",
						Type:   "string",
						Number: 1,
					},
					{
						Name:   "email",
						Type:   "string",
						Number: 2,
					},
				},
			},
		},
	}

	// Convert to OpenAPI
	doc, err := generator.ConvertToOpenAPI(parsedFile)
	assert.NoError(t, err)
	assert.NotNil(t, doc)

	// Test document version
	assert.Equal(t, "3.1.0", doc.Version)

	// Test info
	assert.Equal(t, "test.package", doc.Info.Title)
	assert.Equal(t, "1.0.0", doc.Info.Version)

	// Test paths
	assert.Equal(t, 2, doc.Paths.PathItems.Len())

	// Test /v1/users/{user_id} path
	pathItem, ok := doc.Paths.PathItems.Get("/v1/users/{user_id}")
	assert.True(t, ok)
	assert.NotNil(t, pathItem)

	// Test GET operation
	assert.NotNil(t, pathItem.Get)
	assert.Equal(t, "GetUser", pathItem.Get.OperationId)
	assert.Equal(t, "GetUser operation", pathItem.Get.Summary)
	assert.Nil(t, pathItem.Get.RequestBody)

	// Test response content for GET operation
	response, ok := pathItem.Get.Responses.Codes.Get("200")
	assert.True(t, ok)
	assert.NotNil(t, response)
	responseContent, ok := response.Content.Get("application/json")
	assert.True(t, ok)
	assert.NotNil(t, responseContent.Schema)

	// Test PUT operation
	assert.NotNil(t, pathItem.Put)
	assert.Equal(t, "UpdateUser", pathItem.Put.OperationId)
	assert.Equal(t, "UpdateUser operation", pathItem.Put.Summary)
	assert.NotNil(t, pathItem.Put.RequestBody)

	// Test PATCH operation
	assert.NotNil(t, pathItem.Patch)
	assert.Equal(t, "PatchUser", pathItem.Patch.OperationId)
	assert.Equal(t, "PatchUser operation", pathItem.Patch.Summary)
	assert.NotNil(t, pathItem.Patch.RequestBody)

	// Test DELETE operation
	assert.NotNil(t, pathItem.Delete)
	assert.Equal(t, "DeleteUser", pathItem.Delete.OperationId)
	assert.Equal(t, "DeleteUser operation", pathItem.Delete.Summary)
	assert.Nil(t, pathItem.Delete.RequestBody)

	// Test /v1/users path
	pathItem, ok = doc.Paths.PathItems.Get("/v1/users")
	assert.True(t, ok)
	assert.NotNil(t, pathItem)

	// Test POST operation
	assert.NotNil(t, pathItem.Post)
	assert.Equal(t, "CreateUser", pathItem.Post.OperationId)
	assert.Equal(t, "CreateUser operation", pathItem.Post.Summary)
	assert.NotNil(t, pathItem.Post.RequestBody)
	content, ok := pathItem.Post.RequestBody.Content.Get("application/json")
	assert.True(t, ok)
	assert.NotNil(t, content.Schema)

	// Test schemas
	assert.Equal(t, 4, doc.Components.Schemas.Len())
}

func TestConvertToOpenAPI_NilFile(t *testing.T) {
	doc, err := generator.ConvertToOpenAPI(nil)
	assert.Error(t, err)
	assert.Nil(t, doc)
	assert.Equal(t, "parsedFile is nil", err.Error())
}
