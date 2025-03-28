package generator

import (
	"fmt"

	"google.golang.org/genproto/googleapis/api/annotations"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"

	"github.com/sapk/protoc-gen-openapiv3/options"
)

// ParsedFile represents the parsed proto file with all necessary information
type ParsedFile struct {
	Package     string
	Services    []ParsedService
	Messages    []ParsedMessage
	Enums       []ParsedEnum
	Imports     []string
	Annotations map[string]string
	Info        *ParsedInfo
}

// ParsedService represents a parsed service definition
type ParsedService struct {
	Name        string
	Methods     []ParsedMethod
	Annotations map[string]string
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
}

// ParsedMessage represents a parsed message definition
type ParsedMessage struct {
	Name        string
	Fields      []ParsedField
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
	Values      []ParsedEnumValue
	Annotations map[string]string
}

// ParsedEnumValue represents a parsed enum value
type ParsedEnumValue struct {
	Name        string
	Number      int32
	Annotations map[string]string
}

// ParsedInfo represents the OpenAPI Info object
type ParsedInfo struct {
	Title          string
	Description    string
	TermsOfService string
	Contact        *ParsedContact
	License        *ParsedLicense
	Version        string
}

// ParsedContact represents the OpenAPI Contact object
type ParsedContact struct {
	Name  string
	URL   string
	Email string
}

// ParsedLicense represents the OpenAPI License object
type ParsedLicense struct {
	Name string
	URL  string
}

// ParseProtoFile parses a proto file and extracts all necessary information
func (g *OpenAPIGenerator) ParseProtoFile(file *protogen.File) (*ParsedFile, error) {
	parsed := &ParsedFile{
		Package:     string(file.Desc.Package()),
		Annotations: make(map[string]string),
		Services:    make([]ParsedService, 0),
		Messages:    make([]ParsedMessage, 0),
		Enums:       make([]ParsedEnum, 0),
		Imports:     make([]string, 0),
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
			info := infoExt.(*options.Info)
			parsed.Info = &ParsedInfo{
				Title:          info.GetTitle(),
				Description:    info.GetDescription(),
				TermsOfService: info.GetTermsOfService(),
				Version:        info.GetVersion(),
			}

			// Parse Contact
			if info.Contact != nil {
				parsed.Info.Contact = &ParsedContact{
					Name:  info.Contact.GetName(),
					URL:   info.Contact.GetUrl(),
					Email: info.Contact.GetEmail(),
				}
			}

			// Parse License
			if info.License != nil {
				parsed.Info.License = &ParsedLicense{
					Name: info.License.GetName(),
					URL:  info.License.GetUrl(),
				}
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

	return parsed, nil
}

// parseService parses a service definition
func (g *OpenAPIGenerator) parseService(service *protogen.Service) (ParsedService, error) {
	parsed := ParsedService{
		Name:        string(service.Desc.Name()),
		Methods:     make([]ParsedMethod, 0),
		Annotations: make(map[string]string),
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
	}

	return parsed, nil
}

// parseMessage parses a message definition
func (g *OpenAPIGenerator) parseMessage(message *protogen.Message) (ParsedMessage, error) {
	parsed := ParsedMessage{
		Name:        string(message.Desc.Name()),
		Fields:      make([]ParsedField, 0),
		Annotations: make(map[string]string),
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
