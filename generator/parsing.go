package generator

import (
	"fmt"

	"google.golang.org/genproto/googleapis/api/annotations"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"

	v2options "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	"github.com/sapk/protoc-gen-openapiv3/options"
)

// ParsedFile represents the parsed proto file with all necessary information
type ParsedFile struct {
	Package         string
	Services        []ParsedService
	Messages        []ParsedMessage
	Enums           []ParsedEnum
	Imports         []string
	Annotations     map[string]string
	Info            *options.Info
	Servers         []*options.Server
	SecuritySchemes []*options.SecurityScheme
	Security        []*options.SecurityRequirement
	Tags            []*options.Tag
	ExternalDocs    *options.ExternalDocumentation
	V2Swagger       *v2options.Swagger
}

// ParsedService represents a parsed service definition
type ParsedService struct {
	Name        string
	Methods     []ParsedMethod
	Annotations map[string]string
	Comment     string
}

// ParsedMethod represents a parsed method definition
type ParsedMethod struct {
	Name        string
	InputType   string
	OutputType  string
	HTTPMethod  string
	HTTPPath    string
	HTTPBody    string
	Annotations map[string]string
	Comment     string
	Operation   *options.Operation
	Security    []*options.SecurityRequirement
	Responses   []*options.Response
	RequestBody *options.RequestBody
	Parameters  []*options.Parameter
}

// ParsedMessage represents a parsed message definition
type ParsedMessage struct {
	Name        string
	Fields      []ParsedField
	Annotations map[string]string
	Comment     string
}

// ParsedField represents a parsed field definition
type ParsedField struct {
	Name        string
	Type        string
	Number      int32
	Annotations map[string]string
	Comment     string
}

// ParsedEnum represents a parsed enum definition
type ParsedEnum struct {
	Name        string
	Values      []ParsedEnumValue
	Annotations map[string]string
	Comment     string
}

// ParsedEnumValue represents a parsed enum value
type ParsedEnumValue struct {
	Name        string
	Number      int32
	Annotations map[string]string
}

// ParseProtoFile parses a proto file and extracts all necessary information
func (g *OpenAPIGenerator) ParseProtoFile(file *protogen.File) (*ParsedFile, error) {
	parsed := &ParsedFile{
		Package:         string(file.Desc.Package()),
		Annotations:     make(map[string]string),
		Services:        make([]ParsedService, 0),
		Messages:        make([]ParsedMessage, 0),
		Enums:           make([]ParsedEnum, 0),
		Imports:         make([]string, 0),
		Servers:         make([]*options.Server, 0),
		SecuritySchemes: make([]*options.SecurityScheme, 0),
		Security:        make([]*options.SecurityRequirement, 0),
		Tags:            make([]*options.Tag, 0),
	}

	// Parse imports
	for i := 0; i < file.Desc.Imports().Len(); i++ {
		imp := file.Desc.Imports().Get(i)
		parsed.Imports = append(parsed.Imports, imp.Path())
	}

	// Parse options
	if file.Desc.Options() != nil {
		// Parse OpenAPI Info options
		infoExt := proto.GetExtension(file.Desc.Options(), options.E_Info)
		if infoExt != nil {
			info, ok := infoExt.(*options.Info)
			if ok {
				parsed.Info = info
			}
		}

		// Parse OpenAPI Server options
		serversExt := proto.GetExtension(file.Desc.Options(), options.E_Server)
		if serversExt != nil {
			servers, ok := serversExt.([]*options.Server)
			if ok {
				parsed.Servers = servers
			}
		}

		// Parse OpenAPI SecurityScheme options
		securitySchemeExt := proto.GetExtension(file.Desc.Options(), options.E_SecurityScheme)
		if securitySchemeExt != nil {
			schemes, ok := securitySchemeExt.([]*options.SecurityScheme)
			if ok {
				parsed.SecuritySchemes = schemes
			}
		}

		// Parse OpenAPI Security options
		securityExt := proto.GetExtension(file.Desc.Options(), options.E_Security)
		if securityExt != nil {
			security, ok := securityExt.([]*options.SecurityRequirement)
			if ok {
				parsed.Security = security
			}
		}

		// Parse OpenAPI Tag options
		tagsExt := proto.GetExtension(file.Desc.Options(), options.E_Tag)
		if tagsExt != nil {
			tags, ok := tagsExt.([]*options.Tag)
			if ok {
				parsed.Tags = tags
			}
		}

		// Parse v2 Swagger options
		v2SwaggerExt := proto.GetExtension(file.Desc.Options(), v2options.E_Openapiv2Swagger)
		if v2SwaggerExt != nil {
			swagger, ok := v2SwaggerExt.(*v2options.Swagger)
			if ok {
				parsed.V2Swagger = swagger
			}
		}
	}

	// Parse services
	for _, service := range file.Services {
		parsedService, err := g.parseService(service)
		if err != nil {
			return nil, fmt.Errorf("failed to parse service %s: %w", service.Desc.Name(), err)
		}
		parsed.Services = append(parsed.Services, parsedService)
	}

	// Parse messages
	for _, message := range file.Messages {
		parsedMessage, err := g.parseMessage(message)
		if err != nil {
			return nil, fmt.Errorf("failed to parse message %s: %w", message.Desc.Name(), err)
		}
		parsed.Messages = append(parsed.Messages, parsedMessage)
	}

	// Parse enums
	for _, enum := range file.Enums {
		parsedEnum, err := g.parseEnum(enum)
		if err != nil {
			return nil, fmt.Errorf("failed to parse enum %s: %w", enum.Desc.Name(), err)
		}
		parsed.Enums = append(parsed.Enums, parsedEnum)
	}

	// Convert v2 annotations to v3 format
	if err := g.convertV2ToV3(parsed); err != nil {
		return nil, fmt.Errorf("failed to convert v2 to v3: %w", err)
	}

	return parsed, nil
}

// parseService parses a service definition
func (g *OpenAPIGenerator) parseService(service *protogen.Service) (ParsedService, error) {
	parsed := ParsedService{
		Name:        string(service.Desc.Name()),
		Methods:     make([]ParsedMethod, 0),
		Annotations: make(map[string]string),
		Comment:     string(service.Comments.Leading),
	}

	// Parse methods
	for _, method := range service.Methods {
		parsedMethod, err := g.parseMethod(method)
		if err != nil {
			return parsed, fmt.Errorf("failed to parse method %s: %w", method.Desc.Name(), err)
		}
		parsed.Methods = append(parsed.Methods, parsedMethod)
	}

	return parsed, nil
}

// parseMethod parses a method definition
func (g *OpenAPIGenerator) parseMethod(method *protogen.Method) (ParsedMethod, error) {
	parsed := ParsedMethod{
		Name:        string(method.Desc.Name()),
		InputType:   string(method.Input.Desc.FullName()),
		OutputType:  string(method.Output.Desc.FullName()),
		Annotations: make(map[string]string),
		Comment:     string(method.Comments.Leading),
		Security:    make([]*options.SecurityRequirement, 0),
		Responses:   make([]*options.Response, 0),
		Parameters:  make([]*options.Parameter, 0),
	}

	// Parse HTTP annotations
	if method.Desc.Options() != nil {
		httpRule := proto.GetExtension(method.Desc.Options(), annotations.E_Http).(*annotations.HttpRule)
		if httpRule != nil {
			// Parse HTTP method and path
			switch {
			case httpRule.GetGet() != "":
				parsed.HTTPMethod = "GET"
				parsed.HTTPPath = httpRule.GetGet()
			case httpRule.GetPost() != "":
				parsed.HTTPMethod = "POST"
				parsed.HTTPPath = httpRule.GetPost()
			case httpRule.GetPut() != "":
				parsed.HTTPMethod = "PUT"
				parsed.HTTPPath = httpRule.GetPut()
			case httpRule.GetDelete() != "":
				parsed.HTTPMethod = "DELETE"
				parsed.HTTPPath = httpRule.GetDelete()
			case httpRule.GetPatch() != "":
				parsed.HTTPMethod = "PATCH"
				parsed.HTTPPath = httpRule.GetPatch()
			}

			// Parse body field
			if httpRule.Body != "" {
				parsed.HTTPBody = httpRule.Body
			}
		}

		// Parse OpenAPI Operation annotation
		operationExt := proto.GetExtension(method.Desc.Options(), options.E_Operation)
		if operationExt != nil {
			operation, ok := operationExt.(*options.Operation)
			if ok {
				parsed.Operation = operation
				// Parse security requirements if present
				parsed.Security = operation.GetSecurity()
				// Parse responses
				parsed.Responses = operation.GetResponses()
				// Parse request body if present
				parsed.RequestBody = operation.GetRequestBody()
				// Parse parameters if present
				parsed.Parameters = operation.GetParameters()
			}
		}
	}

	return parsed, nil
}

// parseMessage parses a message definition
func (g *OpenAPIGenerator) parseMessage(message *protogen.Message) (ParsedMessage, error) {
	parsed := ParsedMessage{
		Name:        string(message.Desc.Name()),
		Fields:      make([]ParsedField, 0),
		Annotations: make(map[string]string),
		Comment:     string(message.Comments.Leading),
	}

	// Parse fields
	for _, field := range message.Fields {
		parsedField, err := g.parseField(field)
		if err != nil {
			return parsed, fmt.Errorf("failed to parse field %s: %w", field.Desc.Name(), err)
		}
		parsed.Fields = append(parsed.Fields, parsedField)
	}

	return parsed, nil
}

// parseField parses a field definition
func (g *OpenAPIGenerator) parseField(field *protogen.Field) (ParsedField, error) {
	parsed := ParsedField{
		Name:        string(field.Desc.Name()),
		Number:      int32(field.Desc.Number()),
		Annotations: make(map[string]string),
		Comment:     string(field.Comments.Leading),
	}

	// Handle field type
	switch {
	case field.Desc.IsMap():
		keyType := field.Message.Fields[0].Desc.Kind().String()
		valueType := field.Message.Fields[1].Desc.Kind().String()
		parsed.Type = fmt.Sprintf("map<%s, %s>", keyType, valueType)
	case field.Desc.IsList():
		parsed.Type = fmt.Sprintf("repeated %s", getFieldType(field))
	case field.Desc.HasOptionalKeyword():
		parsed.Type = fmt.Sprintf("optional %s", getFieldType(field))
	default:
		parsed.Type = getFieldType(field)
	}

	return parsed, nil
}

// getFieldType returns the field type as a string
func getFieldType(field *protogen.Field) string {
	if field.Enum != nil {
		return string(field.Enum.Desc.FullName())
	}
	if field.Message != nil {
		return string(field.Message.Desc.FullName())
	}
	return field.Desc.Kind().String()
}

// parseEnum parses an enum definition
func (g *OpenAPIGenerator) parseEnum(enum *protogen.Enum) (ParsedEnum, error) {
	parsed := ParsedEnum{
		Name:        string(enum.Desc.Name()),
		Values:      make([]ParsedEnumValue, 0),
		Annotations: make(map[string]string),
		Comment:     string(enum.Comments.Leading),
	}

	// Parse enum values
	for _, value := range enum.Values {
		parsedValue, err := g.parseEnumValue(value)
		if err != nil {
			return parsed, fmt.Errorf("failed to parse enum value %s: %w", value.Desc.Name(), err)
		}
		parsed.Values = append(parsed.Values, parsedValue)
	}

	return parsed, nil
}

// parseEnumValue parses an enum value
func (g *OpenAPIGenerator) parseEnumValue(value *protogen.EnumValue) (ParsedEnumValue, error) {
	parsed := ParsedEnumValue{
		Name:        string(value.Desc.Name()),
		Number:      int32(value.Desc.Number()),
		Annotations: make(map[string]string),
	}

	return parsed, nil
}
