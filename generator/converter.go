package generator

import (
	"fmt"
	"strings"

	"github.com/pb33f/libopenapi/datamodel/high/base"
	high "github.com/pb33f/libopenapi/datamodel/high/v3"
	"github.com/pb33f/libopenapi/orderedmap"
)

// ConvertToOpenAPI converts a ParsedFile to a libopenapi Document
func ConvertToOpenAPI(parsedFile *ParsedFile) (*high.Document, error) {
	if parsedFile == nil {
		return nil, fmt.Errorf("parsedFile is nil")
	}

	// Create the root document
	doc := &high.Document{
		Version: "3.1.0",
		Info: &base.Info{
			Title:   parsedFile.Package,
			Version: "1.0.0",
		},
		Paths: &high.Paths{
			PathItems: orderedmap.New[string, *high.PathItem](),
		},
		Components: &high.Components{
			Schemas: orderedmap.New[string, *base.SchemaProxy](),
		},
	}

	// Convert services to paths
	for _, service := range parsedFile.Services {
		for _, method := range service.Methods {
			path := convertMethodToPath(method)
			pathItem := &high.PathItem{
				Post: &high.Operation{
					OperationId: method.Name,
					Summary:     fmt.Sprintf("%s operation", method.Name),
					RequestBody: &high.RequestBody{
						Content: orderedmap.New[string, *high.MediaType](),
					},
					Responses: &high.Responses{
						Codes: orderedmap.New[string, *high.Response](),
					},
				},
			}

			// Add request body content
			pathItem.Post.RequestBody.Content.Set("application/json", &high.MediaType{
				Schema: convertMessageToSchema(parsedFile, method.InputType),
			})

			// Add response content
			response := &high.Response{
				Content: orderedmap.New[string, *high.MediaType](),
			}
			response.Content.Set("application/json", &high.MediaType{
				Schema: convertMessageToSchema(parsedFile, method.OutputType),
			})
			pathItem.Post.Responses.Codes.Set("200", response)

			// Add path item
			doc.Paths.PathItems.Set(path, pathItem)
		}
	}

	// Convert messages to schemas
	for _, message := range parsedFile.Messages {
		schema := convertMessageToSchema(parsedFile, message.Name)
		doc.Components.Schemas.Set(message.Name, schema)
	}

	return doc, nil
}

// convertMethodToPath converts a method name to a path
func convertMethodToPath(method ParsedMethod) string {
	// Convert camelCase to kebab-case
	path := method.Name
	for i := 1; i < len(path); i++ {
		if path[i] >= 'A' && path[i] <= 'Z' {
			path = path[:i] + "-" + strings.ToLower(string(path[i])) + path[i+1:]
		}
	}
	return "/" + strings.ToLower(path)
}

// convertMessageToSchema converts a message to a schema
func convertMessageToSchema(parsedFile *ParsedFile, messageName string) *base.SchemaProxy {
	// Find the message in the parsed file
	var message *ParsedMessage
	for i := range parsedFile.Messages {
		if parsedFile.Messages[i].Name == messageName {
			message = &parsedFile.Messages[i]
			break
		}
	}

	if message == nil {
		return base.CreateSchemaProxy(&base.Schema{
			Type: []string{"object"},
		})
	}

	// Create the schema
	schema := &base.Schema{
		Type:       []string{"object"},
		Properties: orderedmap.New[string, *base.SchemaProxy](),
		Required:   make([]string, 0),
	}

	// Convert fields to properties
	for _, field := range message.Fields {
		property := convertFieldToSchema(&field)
		schema.Properties.Set(field.Name, property)
		if !strings.HasPrefix(field.Type, "optional") {
			schema.Required = append(schema.Required, field.Name)
		}
	}

	return base.CreateSchemaProxy(schema)
}

// convertFieldToSchema converts a field to a schema
func convertFieldToSchema(field *ParsedField) *base.SchemaProxy {
	return base.CreateSchemaProxy(&base.Schema{
		Type: []string{convertProtoTypeToOpenAPIType(field.Type)},
	})
}

// convertProtoTypeToOpenAPIType converts a proto type to an OpenAPI type
func convertProtoTypeToOpenAPIType(protoType string) string {
	switch protoType {
	case "string":
		return "string"
	case "int32", "int64", "uint32", "uint64":
		return "integer"
	case "float", "double":
		return "number"
	case "bool":
		return "boolean"
	case "bytes":
		return "string"
	default:
		return "object"
	}
}
