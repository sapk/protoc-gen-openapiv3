package generator

import (
	"fmt"
	"strings"

	"github.com/pb33f/libopenapi/datamodel/high/base"
	high "github.com/pb33f/libopenapi/datamodel/high/v3"
	"github.com/pb33f/libopenapi/orderedmap"
	"gopkg.in/yaml.v3"
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
			path := method.HTTPPath
			if path == "" {
				path = convertMethodToPath(method)
			}

			// Get or create path item
			pathItem, exists := doc.Paths.PathItems.Get(path)
			if !exists {
				pathItem = &high.PathItem{}
			}

			// Create operation based on HTTP method
			operation := &high.Operation{
				OperationId: method.Name,
				Summary:     fmt.Sprintf("%s operation", method.Name),
				Responses: &high.Responses{
					Codes: orderedmap.New[string, *high.Response](),
				},
				Parameters: make([]*high.Parameter, 0),
			}

			// Extract path parameters
			pathParams := extractPathParameters(path)
			for _, param := range pathParams {
				required := true
				operation.Parameters = append(operation.Parameters, &high.Parameter{
					Name:        param,
					In:          "path",
					Required:    &required,
					Schema:      createPrimitiveSchema("string"),
					Description: fmt.Sprintf("Path parameter %s", param),
				})
			}

			// Find the input message type
			var inputMessage *ParsedMessage
			for _, msg := range parsedFile.Messages {
				if msg.Name == method.InputType {
					inputMessage = &msg
					break
				}
			}

			// Add query parameters from input message fields
			if inputMessage != nil {
				bodyField := method.HTTPBody
				pathParamMap := make(map[string]bool)
				for _, param := range pathParams {
					pathParamMap[param] = true
				}

				for _, field := range inputMessage.Fields {
					// Skip fields that are in the path or body
					if pathParamMap[field.Name] || field.Name == bodyField {
						continue
					}

					// Create query parameter
					required := !strings.HasPrefix(field.Type, "optional")
					param := &high.Parameter{
						Name:        field.Name,
						In:          "query",
						Required:    &required,
						Description: fmt.Sprintf("Query parameter %s", field.Name),
					}

					// Handle array type for query parameters
					if strings.HasPrefix(field.Type, "repeated ") {
						itemType := strings.TrimPrefix(field.Type, "repeated ")
						if schema := createPrimitiveSchema(itemType); schema != nil {
							explode := true
							param.Schema = base.CreateSchemaProxy(&base.Schema{
								Type: []string{"array"},
								Items: &base.DynamicValue[*base.SchemaProxy, bool]{
									A: schema,
								},
							})
							param.Style = "form"
							param.Explode = &explode
						}
					} else {
						param.Schema = convertFieldToSchema(&field, parsedFile, doc)
					}

					operation.Parameters = append(operation.Parameters, param)
				}
			}

			// Set operation based on HTTP method
			switch method.HTTPMethod {
			case "GET":
				pathItem.Get = operation
			case "POST":
				pathItem.Post = operation
			case "PUT":
				pathItem.Put = operation
			case "PATCH":
				pathItem.Patch = operation
			case "DELETE":
				pathItem.Delete = operation
			default:
				// Default to POST if no HTTP method is specified
				pathItem.Post = operation
			}

			// Add request body if specified
			if method.HTTPBody != "" {
				operation.RequestBody = &high.RequestBody{
					Content: orderedmap.New[string, *high.MediaType](),
				}
				operation.RequestBody.Content.Set("application/json", &high.MediaType{
					Schema: convertMessageToSchema(parsedFile, method.InputType, doc),
				})
			}

			// Add response content
			response := &high.Response{
				Description: fmt.Sprintf("Response for %s operation", method.Name),
			}

			// Only add content if the response type is not Empty
			if method.OutputType != "google.protobuf.Empty" {
				response.Content = orderedmap.New[string, *high.MediaType]()
				response.Content.Set("application/json", &high.MediaType{
					Schema: convertMessageToSchema(parsedFile, method.OutputType, doc),
				})
			}

			operation.Responses.Codes.Set("200", response)

			// Update path item
			doc.Paths.PathItems.Set(path, pathItem)
		}
	}

	return doc, nil
}

// convertMethodToPath converts a method name to a path (fallback when no HTTP path is specified)
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

// extractPathParameters extracts parameter names from a path template
func extractPathParameters(path string) []string {
	var params []string
	start := 0
	for i := 0; i < len(path); i++ {
		if path[i] == '{' {
			start = i + 1
		} else if path[i] == '}' && start > 0 {
			param := path[start:i]
			// Handle field path notation (e.g., {user.name})
			if dotIndex := strings.LastIndex(param, "."); dotIndex != -1 {
				param = param[dotIndex+1:]
			}
			params = append(params, param)
			start = 0
		}
	}
	return params
}

// convertMessageToSchema converts a message to a schema
func convertMessageToSchema(parsedFile *ParsedFile, messageName string, doc *high.Document) *base.SchemaProxy {
	// Create a map to track schemas being processed
	processingSchemas := make(map[string]bool)

	// Inner function to handle the actual conversion
	var convert func(string) *base.SchemaProxy
	convert = func(name string) *base.SchemaProxy {
		// Check if we're already processing this schema
		if processingSchemas[name] {
			// If we are, create a reference to avoid infinite recursion
			refName := name
			if strings.Contains(name, ".") {
				parts := strings.Split(name, ".")
				refName = parts[len(parts)-1]
			}
			return base.CreateSchemaProxyRef(fmt.Sprintf("#/components/schemas/%s", refName))
		}

		// Mark this schema as being processed
		processingSchemas[name] = true
		defer func() {
			processingSchemas[name] = false
		}()

		// Try to convert the schema using different handlers
		if schema := handleMessage(parsedFile, name, doc); schema != nil {
			return schema
		}

		if schema := handleEnum(parsedFile, name, doc); schema != nil {
			return schema
		}

		// Handle as reference
		return handleReference(name, doc, convert)
	}

	return convert(messageName)
}

// handleMessage handles conversion of a message type to a schema
func handleMessage(parsedFile *ParsedFile, name string, doc *high.Document) *base.SchemaProxy {
	if parsedFile == nil {
		return nil
	}

	var message *ParsedMessage
	for i := range parsedFile.Messages {
		if parsedFile.Messages[i].Name == name {
			message = &parsedFile.Messages[i]
			break
		}
	}

	if message == nil {
		return nil
	}

	// Create the schema
	schema := &base.Schema{
		Type:       []string{"object"},
		Properties: orderedmap.New[string, *base.SchemaProxy](),
		Required:   make([]string, 0),
	}

	// Convert fields to properties
	for _, field := range message.Fields {
		property := convertFieldToSchema(&field, parsedFile, doc)
		schema.Properties.Set(field.Name, property)
		if !strings.HasPrefix(field.Type, "optional") {
			schema.Required = append(schema.Required, field.Name)
		}
	}

	// Create the schema proxy and store it in components
	schemaProxy := base.CreateSchemaProxy(schema)
	return schemaProxy
}

// handleEnum handles conversion of an enum type to a schema
func handleEnum(parsedFile *ParsedFile, name string, doc *high.Document) *base.SchemaProxy {
	if parsedFile == nil {
		return nil
	}

	var enum *ParsedEnum
	for i := range parsedFile.Enums {
		if parsedFile.Enums[i].Name == name {
			enum = &parsedFile.Enums[i]
			break
		}
	}

	if enum == nil {
		return nil
	}

	// Create enum schema
	schema := &base.Schema{
		Type: []string{"string"},
		Enum: make([]*yaml.Node, len(enum.Values)),
	}
	for i, value := range enum.Values {
		schema.Enum[i] = &yaml.Node{
			Kind:  yaml.ScalarNode,
			Value: value.Name,
		}
	}

	if len(schema.Enum) > 0 {
		schema.Default = schema.Enum[0]
	}

	schemaProxy := base.CreateSchemaProxy(schema)
	return schemaProxy
}

// handleReference handles conversion of a reference type to a schema
func handleReference(name string, doc *high.Document, convert func(string) *base.SchemaProxy) *base.SchemaProxy {
	// Strip package name from the reference
	refName := name
	if strings.Contains(name, ".") {
		parts := strings.Split(name, ".")
		refName = parts[len(parts)-1]
	}

	// Check if the schema already exists in components
	if _, exists := doc.Components.Schemas.Get(refName); !exists {
		schema := convert(refName)
		doc.Components.Schemas.Set(refName, schema)
	}

	return base.CreateSchemaProxyRef(fmt.Sprintf("#/components/schemas/%s", refName))
}

// createPrimitiveSchema creates a schema for a primitive type
func createPrimitiveSchema(primitiveType string) *base.SchemaProxy {
	switch primitiveType {
	case "string":
		return base.CreateSchemaProxy(&base.Schema{
			Type: []string{"string"},
		})
	case "bytes":
		return base.CreateSchemaProxy(&base.Schema{
			Type: []string{"string"},
		})
	case "int32", "int64", "uint32", "uint64":
		return base.CreateSchemaProxy(&base.Schema{
			Type: []string{"integer"},
		})
	case "float", "double":
		return base.CreateSchemaProxy(&base.Schema{
			Type: []string{"number"},
		})
	case "bool":
		return base.CreateSchemaProxy(&base.Schema{
			Type: []string{"boolean"},
		})
	case "google.protobuf.Timestamp":
		return base.CreateSchemaProxy(&base.Schema{
			Type:   []string{"string"},
			Format: "date-time",
		})
	case "google.protobuf.Empty":
		return base.CreateSchemaProxy(nil)
	default:
		return nil
	}
}

// convertFieldToSchema converts a field to a schema
func convertFieldToSchema(field *ParsedField, parsedFile *ParsedFile, doc *high.Document) *base.SchemaProxy {
	// Handle special types
	if strings.HasPrefix(field.Type, "repeated ") {
		itemType := strings.TrimPrefix(field.Type, "repeated ")
		// For primitive types, create the schema directly
		if schema := createPrimitiveSchema(itemType); schema != nil {
			return base.CreateSchemaProxy(&base.Schema{
				Type: []string{"array"},
				Items: &base.DynamicValue[*base.SchemaProxy, bool]{
					A: schema,
				},
			})
		}
		// For message types, create a reference
		return base.CreateSchemaProxy(&base.Schema{
			Type: []string{"array"},
			Items: &base.DynamicValue[*base.SchemaProxy, bool]{
				A: convertMessageToSchema(parsedFile, itemType, doc),
			},
		})
	}

	if strings.HasPrefix(field.Type, "optional ") {
		itemType := strings.TrimPrefix(field.Type, "optional ")
		return convertMessageToSchema(parsedFile, itemType, doc)
	}

	if strings.HasPrefix(field.Type, "map<") {
		// Extract key and value types from map<key, value>
		mapType := strings.TrimPrefix(field.Type, "map<")
		mapType = strings.TrimSuffix(mapType, ">")
		parts := strings.Split(mapType, ", ")
		if len(parts) != 2 {
			return base.CreateSchemaProxy(&base.Schema{
				Type: []string{"object"},
			})
		}
		valueType := parts[1]

		// Handle primitive value types directly
		var valueSchema *base.SchemaProxy
		if schema := createPrimitiveSchema(valueType); schema != nil {
			valueSchema = schema
		} else {
			// For message types, create a reference
			valueSchema = convertMessageToSchema(parsedFile, valueType, doc)
		}

		return base.CreateSchemaProxy(&base.Schema{
			Type: []string{"object"},
			AdditionalProperties: &base.DynamicValue[*base.SchemaProxy, bool]{
				A: valueSchema,
			},
		})
	}

	// Handle primitive types
	if schema := createPrimitiveSchema(field.Type); schema != nil {
		return schema
	}

	// For message types, create a reference
	return convertMessageToSchema(parsedFile, field.Type, doc)
}
