package generator

import (
	"fmt"
	"strings"

	v2options "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	"github.com/sapk/protoc-gen-openapiv3/options"
)

// convertV2ToV3 converts OpenAPI v2 annotations to v3 format
func (g *OpenAPIGenerator) convertV2ToV3(parsed *ParsedFile) error {
	if parsed.V2Swagger == nil {
		return nil
	}

	// Convert info
	if parsed.V2Swagger.Info != nil {
		parsed.Info = &options.Info{
			Title:          parsed.V2Swagger.Info.Title,
			Description:    parsed.V2Swagger.Info.Description,
			TermsOfService: parsed.V2Swagger.Info.TermsOfService,
			Version:        parsed.V2Swagger.Info.Version,
		}

		// Convert contact
		if parsed.V2Swagger.Info.Contact != nil {
			parsed.Info.Contact = &options.Contact{
				Name:  parsed.V2Swagger.Info.Contact.Name,
				Url:   parsed.V2Swagger.Info.Contact.Url,
				Email: parsed.V2Swagger.Info.Contact.Email,
			}
		}

		// Convert license
		if parsed.V2Swagger.Info.License != nil {
			parsed.Info.License = &options.License{
				Name: parsed.V2Swagger.Info.License.Name,
				Url:  parsed.V2Swagger.Info.License.Url,
			}
		}
	}

	// Convert servers
	if parsed.V2Swagger.Host != "" || parsed.V2Swagger.BasePath != "" || len(parsed.V2Swagger.Schemes) > 0 {
		server := convertV2ServerToV3(parsed.V2Swagger.Host, parsed.V2Swagger.BasePath, parsed.V2Swagger.Schemes)
		if server != nil {
			parsed.Servers = append(parsed.Servers, server)
		}
	}

	// Convert security schemes
	if parsed.V2Swagger.SecurityDefinitions != nil {
		for name, scheme := range parsed.V2Swagger.SecurityDefinitions.Security {
			v3Scheme := &options.SecurityScheme{
				Description: scheme.Description,
				Name:        name,
			}

			switch scheme.Type {
			case v2options.SecurityScheme_TYPE_BASIC:
				v3Scheme.Type = "http"
				v3Scheme.Scheme = "basic"
			case v2options.SecurityScheme_TYPE_API_KEY:
				v3Scheme.Type = "apiKey"
				v3Scheme.In = "header"
				v3Scheme.Name = scheme.Name
			case v2options.SecurityScheme_TYPE_OAUTH2:
				v3Scheme.Type = "oauth2"
				if scheme.Flow == v2options.SecurityScheme_FLOW_ACCESS_CODE {
					v3Scheme.Flows = &options.OAuth2Flows{
						AuthorizationCode: &options.OAuth2Flow{
							AuthorizationUrl: scheme.AuthorizationUrl,
							TokenUrl:         scheme.TokenUrl,
							Scopes:           make([]*options.OAuth2Scope, 0),
						},
					}
					for scopeName, scopeDesc := range scheme.Scopes.Scope {
						v3Scheme.Flows.AuthorizationCode.Scopes = append(
							v3Scheme.Flows.AuthorizationCode.Scopes,
							&options.OAuth2Scope{
								Name:        scopeName,
								Description: scopeDesc,
							},
						)
					}
				}
			}

			parsed.SecuritySchemes = append(parsed.SecuritySchemes, v3Scheme)
		}
	}

	// Convert security requirements
	if parsed.V2Swagger.Security != nil {
		for _, req := range parsed.V2Swagger.Security {
			for name, value := range req.SecurityRequirement {
				v3Req := &options.SecurityRequirement{
					Name:   name,
					Scopes: value.Scope,
				}
				parsed.Security = append(parsed.Security, v3Req)
			}
		}
	}

	// Convert tags
	if parsed.V2Swagger.Tags != nil {
		for _, tag := range parsed.V2Swagger.Tags {
			v3Tag := &options.Tag{
				Name:        tag.Name,
				Description: tag.Description,
			}

			if tag.ExternalDocs != nil {
				v3Tag.ExternalDocs = &options.ExternalDocumentation{
					Description: tag.ExternalDocs.Description,
					Url:         tag.ExternalDocs.Url,
				}
			}

			parsed.Tags = append(parsed.Tags, v3Tag)
		}
	}

	// Convert external documentation
	if parsed.V2Swagger.ExternalDocs != nil {
		parsed.ExternalDocs = &options.ExternalDocumentation{
			Description: parsed.V2Swagger.ExternalDocs.Description,
			Url:         parsed.V2Swagger.ExternalDocs.Url,
		}
	}

	// Convert responses
	if parsed.V2Swagger.Responses != nil {
		// Note: Response conversion should be handled in the response-specific conversion logic
		// as it requires handling of different response types and content types
	}

	return nil
}

// convertV2OperationToV3 converts OpenAPI v2 operation options to v3 format
func convertV2OperationToV3(v2Op *v2options.Operation) *options.Operation {
	if v2Op == nil {
		return nil
	}

	v3Op := &options.Operation{
		Summary:     v2Op.Summary,
		Description: v2Op.Description,
		Tags:        v2Op.Tags,
		Deprecated:  v2Op.Deprecated,
	}

	// Convert parameters
	if v2Op.Parameters != nil && len(v2Op.Parameters.Headers) > 0 {
		v3Op.Parameters = make([]*options.Parameter, len(v2Op.Parameters.Headers))
		for i, param := range v2Op.Parameters.Headers {
			v3Param := &options.Parameter{
				Name:        param.Name,
				Description: param.Description,
				Required:    &param.Required,
			}

			// Convert parameter type
			switch param.Type {
			case v2options.HeaderParameter_STRING:
				v3Param.Schema = &options.Schema{
					Type: "string",
				}
			case v2options.HeaderParameter_NUMBER:
				v3Param.Schema = &options.Schema{
					Type: "number",
				}
			case v2options.HeaderParameter_INTEGER:
				v3Param.Schema = &options.Schema{
					Type: "integer",
				}
			case v2options.HeaderParameter_BOOLEAN:
				v3Param.Schema = &options.Schema{
					Type: "boolean",
				}
			}

			// Set parameter location
			v3Param.In = "header"

			v3Op.Parameters[i] = v3Param
		}
	}

	// Convert responses
	if len(v2Op.Responses) > 0 {
		v3Op.Responses = make([]*options.Response, 0)
		for code, resp := range v2Op.Responses {
			v3Resp := &options.Response{
				Code:        code,
				Description: resp.Description,
			}

			// Convert response schema if present
			if resp.Schema != nil && resp.Schema.JsonSchema != nil {
				v3Resp.Content = map[string]*options.MediaType{
					"application/json": {
						Schema: convertV2SchemaToV3(resp.Schema.JsonSchema),
					},
				}
			}

			v3Op.Responses = append(v3Op.Responses, v3Resp)
		}
	}

	// Convert security requirements
	if len(v2Op.Security) > 0 {
		v3Op.Security = make([]*options.SecurityRequirement, len(v2Op.Security))
		for i, sec := range v2Op.Security {
			for name, value := range sec.SecurityRequirement {
				v3Op.Security[i] = &options.SecurityRequirement{
					Name:   name,
					Scopes: value.Scope,
				}
			}
		}
	}

	return v3Op
}

// convertV2SchemaToV3 converts OpenAPI v2 schema to v3 format
func convertV2SchemaToV3(v2Schema *v2options.JSONSchema) *options.Schema {
	if v2Schema == nil {
		return nil
	}

	v3Schema := &options.Schema{
		Description: v2Schema.Description,
		Title:       v2Schema.Title,
		Default:     v2Schema.Default,
		Format:      v2Schema.Format,
	}

	// Convert type
	if len(v2Schema.Type) > 0 {
		switch v2Schema.Type[0] {
		case v2options.JSONSchema_STRING:
			v3Schema.Type = "string"
		case v2options.JSONSchema_NUMBER:
			v3Schema.Type = "number"
		case v2options.JSONSchema_INTEGER:
			v3Schema.Type = "integer"
		case v2options.JSONSchema_BOOLEAN:
			v3Schema.Type = "boolean"
		case v2options.JSONSchema_OBJECT:
			v3Schema.Type = "object"
		case v2options.JSONSchema_ARRAY:
			v3Schema.Type = "array"
		}
	}

	// Convert reference if present
	if v2Schema.Ref != "" {
		v3Schema.Ref = convertV2RefToV3(v2Schema.Ref)
	}

	// Convert required fields
	if len(v2Schema.Required) > 0 {
		v3Schema.Required = v2Schema.Required
	}

	// Convert enum values
	if len(v2Schema.Enum) > 0 {
		v3Schema.Enum = v2Schema.Enum
	}

	// Convert array items
	if v2Schema.Type != nil && len(v2Schema.Type) > 0 && v2Schema.Type[0] == v2options.JSONSchema_ARRAY {
		v3Schema.Items = &options.Schema{
			Type: "string", // Default to string array, adjust if needed
		}
	}

	return v3Schema
}

// convertV2RefToV3 converts OpenAPI v2 reference to v3 format
func convertV2RefToV3(ref string) string {
	// Convert from #/definitions/ to #/components/schemas/
	return strings.Replace(ref, "#/definitions/", "#/components/schemas/", 1)
}

// convertV2ServerToV3 converts OpenAPI v2 server information to v3 format
func convertV2ServerToV3(host, basePath string, schemes []v2options.Scheme) *options.Server {
	if host == "" {
		return nil
	}

	// Build the server URL
	url := ""
	if len(schemes) > 0 {
		switch schemes[0] {
		case v2options.Scheme_HTTP:
			url = "http://"
		case v2options.Scheme_HTTPS:
			url = "https://"
		case v2options.Scheme_WS:
			url = "ws://"
		case v2options.Scheme_WSS:
			url = "wss://"
		default:
			url = "https://"
		}
	} else {
		url = "https://"
	}
	url += host
	if basePath != "" {
		url += basePath
	}

	return &options.Server{
		Url:         url,
		Description: fmt.Sprintf("Server for %s", host),
	}
}
