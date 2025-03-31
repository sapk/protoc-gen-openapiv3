package generator

import (
	"fmt"

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
			v3Req := &options.SecurityRequirement{
				Name:   "",
				Scopes: make([]string, 0),
			}

			for name, value := range req.SecurityRequirement {
				v3Req.Name = name
				v3Req.Scopes = append(v3Req.Scopes, value.Scope...)
			}

			parsed.Security = append(parsed.Security, v3Req)
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

// convertV2ServerToV3 converts OpenAPI v2 server information to v3 format
func convertV2ServerToV3(host, basePath string, schemes []v2options.Scheme) *options.Server {
	if host == "" {
		return nil
	}

	// Build the server URL
	url := ""
	if len(schemes) > 0 {
		url = string(schemes[0]) + "://"
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
