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

	return nil
}

// convertV2RefToV3 converts a v2 reference to v3 format
func convertV2RefToV3(ref string) string {
	if strings.HasPrefix(ref, "#/definitions/") {
		return "#/components/schemas/" + strings.TrimPrefix(ref, "#/definitions/")
	}
	return ref
}

// convertV2SecurityToV3 converts OpenAPI v2 security requirements to v3 format
func convertV2SecurityToV3(security []*v2options.SecurityRequirement) []*options.SecurityRequirement {
	if len(security) == 0 {
		return nil
	}

	v3Security := make([]*options.SecurityRequirement, 0, len(security))
	for _, req := range security {
		// Convert each security requirement in the map
		for name, value := range req.GetSecurityRequirement() {
			v3Req := &options.SecurityRequirement{
				Name:   name,
				Scopes: make([]string, 0),
			}
			if value != nil {
				v3Req.Scopes = append(v3Req.Scopes, value.GetScope()...)
			}
			v3Security = append(v3Security, v3Req)
		}
	}
	return v3Security
}

// convertV2TagToV3 converts OpenAPI v2 tag to v3 format
func convertV2TagToV3(tag *v2options.Tag) *options.Tag {
	if tag == nil {
		return nil
	}

	return &options.Tag{
		Name:         tag.GetName(),
		Description:  tag.GetDescription(),
		ExternalDocs: convertV2ExternalDocsToV3(tag.GetExternalDocs()),
	}
}

// convertV2ExternalDocsToV3 converts OpenAPI v2 external documentation to v3 format
func convertV2ExternalDocsToV3(docs *v2options.ExternalDocumentation) *options.ExternalDocumentation {
	if docs == nil {
		return nil
	}

	return &options.ExternalDocumentation{
		Description: docs.GetDescription(),
		Url:         docs.GetUrl(),
	}
}

// convertV2InfoToV3 converts OpenAPI v2 info to v3 format
func convertV2InfoToV3(info *v2options.Info) *options.Info {
	if info == nil {
		return nil
	}

	return &options.Info{
		Title:          info.GetTitle(),
		Description:    info.GetDescription(),
		TermsOfService: info.GetTermsOfService(),
		Contact:        convertV2ContactToV3(info.GetContact()),
		License:        convertV2LicenseToV3(info.GetLicense()),
		Version:        info.GetVersion(),
	}
}

// convertV2ContactToV3 converts OpenAPI v2 contact to v3 format
func convertV2ContactToV3(contact *v2options.Contact) *options.Contact {
	if contact == nil {
		return nil
	}

	return &options.Contact{
		Name:  contact.GetName(),
		Url:   contact.GetUrl(),
		Email: contact.GetEmail(),
	}
}

// convertV2LicenseToV3 converts OpenAPI v2 license to v3 format
func convertV2LicenseToV3(license *v2options.License) *options.License {
	if license == nil {
		return nil
	}

	return &options.License{
		Name: license.GetName(),
		Url:  license.GetUrl(),
	}
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
