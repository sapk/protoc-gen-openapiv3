package generator

import (
	"fmt"
	"strings"

	"google.golang.org/protobuf/compiler/protogen"
)

// ParsedFile represents the parsed proto file with all necessary information
type ParsedFile struct {
	Package     string
	Services    []*ParsedService
	Messages    []*ParsedMessage
	Enums       []*ParsedEnum
	Imports     []string
	Annotations map[string]string
}

// ParsedService represents a parsed service definition
type ParsedService struct {
	Name        string
	Methods     []*ParsedMethod
	Annotations map[string]string
}

// ParsedMethod represents a parsed method definition
type ParsedMethod struct {
	Name        string
	InputType   string
	OutputType  string
	Annotations map[string]string
}

// ParsedMessage represents a parsed message definition
type ParsedMessage struct {
	Name        string
	Fields      []*ParsedField
	Annotations map[string]string
}

// ParsedField represents a parsed field definition
type ParsedField struct {
	Name        string
	Type        string
	Number      int32
	Annotations map[string]string
}

// ParsedEnum represents a parsed enum definition
type ParsedEnum struct {
	Name        string
	Values      []*ParsedEnumValue
	Annotations map[string]string
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
		Package:     string(file.Desc.Package()),
		Annotations: make(map[string]string),
		Services:    make([]*ParsedService, 0),
		Messages:    make([]*ParsedMessage, 0),
		Enums:       make([]*ParsedEnum, 0),
		Imports:     make([]string, 0),
	}

	// Parse imports
	for i := 0; i < file.Desc.Imports().Len(); i++ {
		imp := file.Desc.Imports().Get(i)
		parsed.Imports = append(parsed.Imports, imp.Path())
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

	return parsed, nil
}

// parseService parses a service definition
func (g *OpenAPIGenerator) parseService(service *protogen.Service) (*ParsedService, error) {
	parsed := &ParsedService{
		Name:        string(service.Desc.Name()),
		Methods:     make([]*ParsedMethod, 0),
		Annotations: make(map[string]string),
	}

	// Parse methods
	for _, method := range service.Methods {
		parsedMethod, err := g.parseMethod(method)
		if err != nil {
			return nil, fmt.Errorf("failed to parse method %s: %w", method.Desc.Name(), err)
		}
		parsed.Methods = append(parsed.Methods, parsedMethod)
	}

	return parsed, nil
}

// parseMethod parses a method definition
func (g *OpenAPIGenerator) parseMethod(method *protogen.Method) (*ParsedMethod, error) {
	parsed := &ParsedMethod{
		Name:        string(method.Desc.Name()),
		InputType:   string(method.Input.Desc.FullName()),
		OutputType:  string(method.Output.Desc.FullName()),
		Annotations: make(map[string]string),
	}

	return parsed, nil
}

// parseMessage parses a message definition
func (g *OpenAPIGenerator) parseMessage(message *protogen.Message) (*ParsedMessage, error) {
	parsed := &ParsedMessage{
		Name:        string(message.Desc.Name()),
		Fields:      make([]*ParsedField, 0),
		Annotations: make(map[string]string),
	}

	// Parse fields
	for _, field := range message.Fields {
		parsedField, err := g.parseField(field)
		if err != nil {
			return nil, fmt.Errorf("failed to parse field %s: %w", field.Desc.Name(), err)
		}
		parsed.Fields = append(parsed.Fields, parsedField)
	}

	return parsed, nil
}

// parseField parses a field definition
func (g *OpenAPIGenerator) parseField(field *protogen.Field) (*ParsedField, error) {
	parsed := &ParsedField{
		Name:        string(field.Desc.Name()),
		Type:        string(field.Desc.Kind().String()),
		Number:      int32(field.Desc.Number()),
		Annotations: make(map[string]string),
	}

	return parsed, nil
}

// parseEnum parses an enum definition
func (g *OpenAPIGenerator) parseEnum(enum *protogen.Enum) (*ParsedEnum, error) {
	parsed := &ParsedEnum{
		Name:        string(enum.Desc.Name()),
		Values:      make([]*ParsedEnumValue, 0),
		Annotations: make(map[string]string),
	}

	// Parse enum values
	for _, value := range enum.Values {
		parsedValue, err := g.parseEnumValue(value)
		if err != nil {
			return nil, fmt.Errorf("failed to parse enum value %s: %w", value.Desc.Name(), err)
		}
		parsed.Values = append(parsed.Values, parsedValue)
	}

	return parsed, nil
}

// parseEnumValue parses an enum value
func (g *OpenAPIGenerator) parseEnumValue(value *protogen.EnumValue) (*ParsedEnumValue, error) {
	parsed := &ParsedEnumValue{
		Name:        string(value.Desc.Name()),
		Number:      int32(value.Desc.Number()),
		Annotations: make(map[string]string),
	}

	return parsed, nil
}

// toOpenAPIName converts proto field names to OpenAPI names
func toOpenAPIName(name string) string {
	// Convert snake_case to camelCase
	parts := strings.Split(name, "_")
	for i := 1; i < len(parts); i++ {
		parts[i] = strings.Title(parts[i])
	}
	return strings.Join(parts, "")
}
