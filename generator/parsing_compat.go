package generator

import (
	"fmt"

	v2options "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	v3options "github.com/sapk/protoc-gen-openapiv3/options"
)

// convertV2ToV3 converts OpenAPI v2 annotations to v3 format
func (g *OpenAPIGenerator) convertV2ToV3(parsed *ParsedFile) error {
	// Convert security schemes
	for _, scheme := range parsed.SecuritySchemes {
		// Convert v2 type enum to v3 string
		switch scheme.Type {
		case "basic":
			scheme.Type = "http"
			scheme.Scheme = "basic"
		case "apiKey":
			scheme.Type = "apiKey"
		case "oauth2":
			scheme.Type = "oauth2"
		}
	}

	// Convert operations
	for i := range parsed.Services {
		for j := range parsed.Services[i].Methods {
			method := &parsed.Services[i].Methods[j]
			if method.Operation == nil {
				continue
			}

			// Convert parameters
			if len(method.Parameters) > 0 {
				// Convert header parameters to security schemes if they are API keys
				for _, param := range method.Parameters {
					if param.In == "header" && param.Schema != nil && param.Schema.Type == "string" {
						// Check if it's an API key
						if param.Name == "X-API-Key" {
							// Add API key security scheme if not exists
							found := false
							for _, scheme := range parsed.SecuritySchemes {
								if scheme.Type == "apiKey" &&
									scheme.Name == param.Name &&
									scheme.In == "header" {
									found = true
									break
								}
							}
							if !found {
								parsed.SecuritySchemes = append(parsed.SecuritySchemes, &v3options.SecurityScheme{
									Type:        "apiKey",
									Name:        param.Name,
									In:          "header",
									Description: param.Description,
								})
							}
						}
					}
				}
			}

			// Convert responses
			if len(method.Responses) > 0 {
				for _, response := range method.Responses {
					// Convert schema references
					if response.Content != nil {
						for _, content := range response.Content {
							if content.Schema != nil && content.Schema.Type == "object" {
								ref := content.Schema.Ref
								if ref != "" {
									// Convert v2 reference format to v3
									content.Schema.Ref = convertV2RefToV3(ref)
								}
							}
						}
					}
				}
			}

			// Convert request body if present
			if method.RequestBody != nil && method.RequestBody.Content != nil {
				for _, content := range method.RequestBody.Content {
					if content.Schema != nil && content.Schema.Type == "object" {
						ref := content.Schema.Ref
						if ref != "" {
							// Convert v2 reference format to v3
							content.Schema.Ref = convertV2RefToV3(ref)
						}
					}
				}
			}
		}
	}

	// Convert message definitions
	for i := range parsed.Messages {
		message := &parsed.Messages[i]
		for j := range message.Fields {
			field := &message.Fields[j]
			// Convert field types if needed
			if field.Type == "google.protobuf.Timestamp" {
				field.Type = "string"
				// Add format annotation
				if field.Annotations == nil {
					field.Annotations = make(map[string]string)
				}
				field.Annotations["format"] = "date-time"
			}
		}
	}

	return nil
}

// convertV2RefToV3 converts OpenAPI v2 reference format to v3
func convertV2RefToV3(ref string) string {
	// Remove #/definitions/ prefix and replace with #/components/schemas/
	if len(ref) > 13 && ref[:13] == "#/definitions/" {
		return "#/components/schemas/" + ref[13:]
	}
	return ref
}

// convertV2SecurityToV3 converts OpenAPI v2 security requirements to v3 format
func convertV2SecurityToV3(security []*v2options.SecurityRequirement) []*v3options.SecurityRequirement {
	if len(security) == 0 {
		return nil
	}

	v3Security := make([]*v3options.SecurityRequirement, 0, len(security))
	for _, req := range security {
		// Convert each security requirement in the map
		for name, value := range req.GetSecurityRequirement() {
			v3Req := &v3options.SecurityRequirement{
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
func convertV2TagToV3(tag *v2options.Tag) *v3options.Tag {
	if tag == nil {
		return nil
	}

	return &v3options.Tag{
		Name:         tag.GetName(),
		Description:  tag.GetDescription(),
		ExternalDocs: convertV2ExternalDocsToV3(tag.GetExternalDocs()),
	}
}

// convertV2ExternalDocsToV3 converts OpenAPI v2 external documentation to v3 format
func convertV2ExternalDocsToV3(docs *v2options.ExternalDocumentation) *v3options.ExternalDocumentation {
	if docs == nil {
		return nil
	}

	return &v3options.ExternalDocumentation{
		Description: docs.GetDescription(),
		Url:         docs.GetUrl(),
	}
}

// convertV2InfoToV3 converts OpenAPI v2 info to v3 format
func convertV2InfoToV3(info *v2options.Info) *v3options.Info {
	if info == nil {
		return nil
	}

	return &v3options.Info{
		Title:          info.GetTitle(),
		Description:    info.GetDescription(),
		TermsOfService: info.GetTermsOfService(),
		Contact:        convertV2ContactToV3(info.GetContact()),
		License:        convertV2LicenseToV3(info.GetLicense()),
		Version:        info.GetVersion(),
	}
}

// convertV2ContactToV3 converts OpenAPI v2 contact to v3 format
func convertV2ContactToV3(contact *v2options.Contact) *v3options.Contact {
	if contact == nil {
		return nil
	}

	return &v3options.Contact{
		Name:  contact.GetName(),
		Url:   contact.GetUrl(),
		Email: contact.GetEmail(),
	}
}

// convertV2LicenseToV3 converts OpenAPI v2 license to v3 format
func convertV2LicenseToV3(license *v2options.License) *v3options.License {
	if license == nil {
		return nil
	}

	return &v3options.License{
		Name: license.GetName(),
		Url:  license.GetUrl(),
	}
}

// convertV2ServerToV3 converts OpenAPI v2 server information to v3 format
func convertV2ServerToV3(host, basePath string, schemes []v2options.Scheme) *v3options.Server {
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

	return &v3options.Server{
		Url:         url,
		Description: fmt.Sprintf("Server for %s", host),
	}
}
