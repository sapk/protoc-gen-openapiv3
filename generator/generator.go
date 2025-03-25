package generator

import (
	"fmt"
	"log"

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
	// Parse the proto file
	parsedFile, err := g.ParseProtoFile(file)
	if err != nil {
		return fmt.Errorf("failed to parse proto file: %w", err)
	}

	// Temporary debug log of parsedFile until we implement the remaining steps
	log.Printf("Parsed file details - Package: %s, Services: %d, Messages: %d, Enums: %d",
		parsedFile.Package,
		len(parsedFile.Services),
		len(parsedFile.Messages),
		len(parsedFile.Enums))

	// TODO: Implement remaining steps using parsedFile
	// 2. Extracting service definitions from parsedFile.Services
	// 3. Processing HTTP annotations from parsedFile.Annotations
	// 4. Generating OpenAPI paths and components using parsedFile.Messages and parsedFile.Enums
	// 5. Writing the output file

	return nil
}
