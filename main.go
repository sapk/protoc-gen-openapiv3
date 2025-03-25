package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/sapk/protoc-gen-openapiv3/generator"

	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/types/pluginpb"
)

var (
	flags             flag.FlagSet
	allowMerge        = flags.Bool("allow_merge", false, "if true, merge generation_opt into a single file")
	includePkgInTags  = flags.Bool("include_package_in_tags", false, "if true, include the package name in the operation tags")
	fqnForOpenAPIName = flags.Bool("fqn_for_openapi_name", false, "if true, use the full qualified name for OpenAPI names")
	outputFile        = flags.String("output", "openapi.yaml", "path to OpenAPI configuration file")
	outputFormat      = flags.String("output-format", "yaml", "format of OpenAPI configuration file")
)

func main() {
	// Check if we're running in test mode
	if len(os.Args) > 0 && strings.HasSuffix(os.Args[0], ".test") {
		// Filter out test-related flags
		var filteredArgs []string
		for _, arg := range os.Args[1:] {
			if !strings.HasPrefix(arg, "-test.") {
				filteredArgs = append(filteredArgs, arg)
			}
		}
		os.Args = append([]string{os.Args[0]}, filteredArgs...)
	}

	// Debug log the arguments
	log.Printf("Program arguments: %v", os.Args)

	protogen.Options{
		ParamFunc: flags.Set,
	}.Run(func(gen *protogen.Plugin) error {
		gen.SupportedFeatures = uint64(pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL)

		// Parse the command line flags
		if err := flags.Parse(os.Args[1:]); err != nil {
			return fmt.Errorf("failed to parse flags: %v", err)
		}

		// Debug log the parsed flags
		log.Printf("Parsed flags - output: %s, format: %s", *outputFile, *outputFormat)

		// Create a new OpenAPI generator
		generator := generator.NewOpenAPIGenerator(gen, &generator.Options{
			AllowMerge:           *allowMerge,
			IncludePackageInTags: *includePkgInTags,
			FQNForOpenAPIName:    *fqnForOpenAPIName,
			OutputFile:           *outputFile,
			OutputFormat:         generator.OutputFormat(*outputFormat),
		})

		// Process each proto file
		for _, f := range gen.Files {
			if f.Generate {
				if err := generator.Generate(f); err != nil {
					return fmt.Errorf("failed to generate OpenAPI spec for %s: %v", f.Desc.Path(), err)
				}
			}
		}

		return nil
	})
}
