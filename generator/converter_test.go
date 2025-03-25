package generator_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sapk/protoc-gen-openapiv3/generator"
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
						Name:       "TestMethod",
						InputType:  "TestRequest",
						OutputType: "TestResponse",
					},
				},
			},
		},
		Messages: []generator.ParsedMessage{
			{
				Name: "TestRequest",
				Fields: []generator.ParsedField{
					{
						Name:   "test_field",
						Type:   "string",
						Number: 1,
					},
				},
			},
			{
				Name: "TestResponse",
				Fields: []generator.ParsedField{
					{
						Name:   "result",
						Type:   "string",
						Number: 1,
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
	assert.Equal(t, "3.0.0", doc.Version)

	// Test info
	assert.Equal(t, "test.package", doc.Info.Title)
	assert.Equal(t, "1.0.0", doc.Info.Version)

	// Test paths
	assert.Equal(t, 1, doc.Paths.PathItems.Len())
	pathItem, ok := doc.Paths.PathItems.Get("/test-method")
	assert.True(t, ok)
	assert.NotNil(t, pathItem)

	// Test operation
	assert.NotNil(t, pathItem.Post)
	assert.Equal(t, "TestMethod", pathItem.Post.OperationId)
	assert.Equal(t, "TestMethod operation", pathItem.Post.Summary)

	// Test request body
	requestBody := pathItem.Post.RequestBody
	assert.NotNil(t, requestBody)
	content, ok := requestBody.Content.Get("application/json")
	assert.True(t, ok)
	assert.NotNil(t, content.Schema)

	// Test response
	responses := pathItem.Post.Responses
	assert.NotNil(t, responses)
	response, ok := responses.Codes.Get("200")
	assert.True(t, ok)
	assert.NotNil(t, response)
	responseContent, ok := response.Content.Get("application/json")
	assert.True(t, ok)
	assert.NotNil(t, responseContent.Schema)

	// Test schemas
	assert.Equal(t, 2, doc.Components.Schemas.Len())

	// Test request schema
	requestSchema, ok := doc.Components.Schemas.Get("TestRequest")
	assert.True(t, ok)
	assert.NotNil(t, requestSchema)

	// Test response schema
	responseSchema, ok := doc.Components.Schemas.Get("TestResponse")
	assert.True(t, ok)
	assert.NotNil(t, responseSchema)
}

func TestConvertToOpenAPI_NilFile(t *testing.T) {
	doc, err := generator.ConvertToOpenAPI(nil)
	assert.Error(t, err)
	assert.Nil(t, doc)
	assert.Equal(t, "parsedFile is nil", err.Error())
}
