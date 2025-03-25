package generator

import (
	"strings"

	"google.golang.org/protobuf/compiler/protogen"
)

// Options contains all the configuration options for the OpenAPI generator
type Options struct {
	AllowMerge           bool
	IncludePackageInTags bool
	FQNForOpenAPIName    bool
	OpenAPIConfiguration string
}

// OpenAPIGenerator handles the generation of OpenAPI specifications
type OpenAPIGenerator struct {
	gen     *protogen.Plugin
	options *Options
}

// NewOpenAPIGenerator creates a new OpenAPI generator with the given options
func NewOpenAPIGenerator(gen *protogen.Plugin, options *Options) *OpenAPIGenerator {
	return &OpenAPIGenerator{
		gen:     gen,
		options: options,
	}
}

// Generate processes a single proto file and generates its OpenAPI specification
func (g *OpenAPIGenerator) Generate(file *protogen.File) error {
	// TODO: Implement the actual OpenAPI generation logic
	// This will include:
	// 1. Parsing the proto file
	// 2. Extracting service definitions
	// 3. Processing HTTP annotations
	// 4. Generating OpenAPI paths and components
	// 5. Writing the output file

	return nil
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
