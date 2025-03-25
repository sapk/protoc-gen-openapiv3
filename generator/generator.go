package generator

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"

	high "github.com/pb33f/libopenapi/datamodel/high/v3"
	"google.golang.org/protobuf/compiler/protogen"
)

// Options contains all the configuration options for the OpenAPI generator
type Options struct {
	AllowMerge           bool
	IncludePackageInTags bool
	FQNForOpenAPIName    bool
	OpenAPIConfiguration string
	OutputFile           string // Path to output file, empty means stdout
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
	oapiDoc, err := ConvertToOpenAPI(parsedFile)
	if err != nil {
		return fmt.Errorf("failed to convert to OpenAPI: %w", err)
	}

	// Debug log of generated OpenAPI document
	log.Printf("Generated OpenAPI doc - Version: %s, Title: %s, Paths: %#v, Schemas: %#v",
		oapiDoc.Version,
		oapiDoc.Info.Title,
		oapiDoc.Paths.PathItems,
		oapiDoc.Components.Schemas)

	// TODO: Implement remaining steps using parsedFile
	// 5. Writing the output file

	// Write the output
	if err := g.writeOutput(oapiDoc); err != nil {
		return fmt.Errorf("failed to write output: %w", err)
	}

	return nil
}

// writeOutput writes the OpenAPI document to either stdout
func (g *OpenAPIGenerator) writeOutput(doc *high.Document) error {
	var writer io.Writer

	if g.options.OutputFile == "" {
		// Write to stdout
		writer = os.Stdout
	} else {
		/*
			// Create output directory if it doesn't exist
			outputDir := filepath.Dir(g.options.OutputFile)
			if err := os.MkdirAll(outputDir, 0755); err != nil {
				return fmt.Errorf("failed to create output directory: %w", err)
			}
		*/

		// Create or truncate the output file
		f, err := os.Create(g.options.OutputFile)
		if err != nil {
			return fmt.Errorf("failed to create output file: %w", err)
		}
		defer f.Close()
		writer = f
	}

	// Convert to JSON
	encoder := json.NewEncoder(writer)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(doc); err != nil {
		return fmt.Errorf("failed to encode OpenAPI document: %w", err)
	}

	return nil
}
