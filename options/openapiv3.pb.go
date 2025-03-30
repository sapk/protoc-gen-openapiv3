// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: options/openapiv3.proto

package options

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Contact information for the exposed API.
type Contact struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The identifying name of the contact person/organization.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The URL pointing to the contact information. MUST be in the format of a URL.
	Url string `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	// The email address of the contact person/organization. MUST be in the format of an email address.
	Email string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *Contact) Reset() {
	*x = Contact{}
	if protoimpl.UnsafeEnabled {
		mi := &file_options_openapiv3_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Contact) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Contact) ProtoMessage() {}

func (x *Contact) ProtoReflect() protoreflect.Message {
	mi := &file_options_openapiv3_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Contact.ProtoReflect.Descriptor instead.
func (*Contact) Descriptor() ([]byte, []int) {
	return file_options_openapiv3_proto_rawDescGZIP(), []int{0}
}

func (x *Contact) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Contact) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *Contact) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

// License information for the exposed API.
type License struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The license name used for the API.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// A URL to the license used for the API. MUST be in the format of a URL.
	Url string `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *License) Reset() {
	*x = License{}
	if protoimpl.UnsafeEnabled {
		mi := &file_options_openapiv3_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *License) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*License) ProtoMessage() {}

func (x *License) ProtoReflect() protoreflect.Message {
	mi := &file_options_openapiv3_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use License.ProtoReflect.Descriptor instead.
func (*License) Descriptor() ([]byte, []int) {
	return file_options_openapiv3_proto_rawDescGZIP(), []int{1}
}

func (x *License) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *License) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

// Info provides metadata about the API.
type Info struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The title of the application.
	Title string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	// A short description of the application. CommonMark syntax MAY be used for rich text representation.
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	// A URL to the Terms of Service for the API. MUST be in the format of a URL.
	TermsOfService string `protobuf:"bytes,3,opt,name=terms_of_service,json=termsOfService,proto3" json:"terms_of_service,omitempty"`
	// The contact information for the exposed API.
	Contact *Contact `protobuf:"bytes,4,opt,name=contact,proto3" json:"contact,omitempty"`
	// The license information for the exposed API.
	License *License `protobuf:"bytes,5,opt,name=license,proto3" json:"license,omitempty"`
	// The version of the OpenAPI document.
	Version string `protobuf:"bytes,6,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *Info) Reset() {
	*x = Info{}
	if protoimpl.UnsafeEnabled {
		mi := &file_options_openapiv3_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Info) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Info) ProtoMessage() {}

func (x *Info) ProtoReflect() protoreflect.Message {
	mi := &file_options_openapiv3_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Info.ProtoReflect.Descriptor instead.
func (*Info) Descriptor() ([]byte, []int) {
	return file_options_openapiv3_proto_rawDescGZIP(), []int{2}
}

func (x *Info) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Info) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Info) GetTermsOfService() string {
	if x != nil {
		return x.TermsOfService
	}
	return ""
}

func (x *Info) GetContact() *Contact {
	if x != nil {
		return x.Contact
	}
	return nil
}

func (x *Info) GetLicense() *License {
	if x != nil {
		return x.License
	}
	return nil
}

func (x *Info) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

// OpenAPI Server Variable object
type ServerVariable struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Enum        []string `protobuf:"bytes,1,rep,name=enum,proto3" json:"enum,omitempty"`
	Default     string   `protobuf:"bytes,2,opt,name=default,proto3" json:"default,omitempty"`
	Description string   `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *ServerVariable) Reset() {
	*x = ServerVariable{}
	if protoimpl.UnsafeEnabled {
		mi := &file_options_openapiv3_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServerVariable) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerVariable) ProtoMessage() {}

func (x *ServerVariable) ProtoReflect() protoreflect.Message {
	mi := &file_options_openapiv3_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerVariable.ProtoReflect.Descriptor instead.
func (*ServerVariable) Descriptor() ([]byte, []int) {
	return file_options_openapiv3_proto_rawDescGZIP(), []int{3}
}

func (x *ServerVariable) GetEnum() []string {
	if x != nil {
		return x.Enum
	}
	return nil
}

func (x *ServerVariable) GetDefault() string {
	if x != nil {
		return x.Default
	}
	return ""
}

func (x *ServerVariable) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

// OpenAPI Server object
type Server struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url         string                     `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	Description string                     `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Variables   map[string]*ServerVariable `protobuf:"bytes,3,rep,name=variables,proto3" json:"variables,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Server) Reset() {
	*x = Server{}
	if protoimpl.UnsafeEnabled {
		mi := &file_options_openapiv3_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Server) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Server) ProtoMessage() {}

func (x *Server) ProtoReflect() protoreflect.Message {
	mi := &file_options_openapiv3_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Server.ProtoReflect.Descriptor instead.
func (*Server) Descriptor() ([]byte, []int) {
	return file_options_openapiv3_proto_rawDescGZIP(), []int{4}
}

func (x *Server) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *Server) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Server) GetVariables() map[string]*ServerVariable {
	if x != nil {
		return x.Variables
	}
	return nil
}

// OAuth2 Flow object
type OAuth2Flow struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The authorization URL to be used for this flow.
	AuthorizationUrl string `protobuf:"bytes,1,opt,name=authorization_url,json=authorizationUrl,proto3" json:"authorization_url,omitempty"`
	// The token URL to be used for this flow.
	TokenUrl string `protobuf:"bytes,2,opt,name=token_url,json=tokenUrl,proto3" json:"token_url,omitempty"`
	// The URL to be used for obtaining refresh tokens.
	RefreshUrl string `protobuf:"bytes,3,opt,name=refresh_url,json=refreshUrl,proto3" json:"refresh_url,omitempty"`
	// The available scopes for the OAuth2 security scheme.
	Scopes map[string]string `protobuf:"bytes,4,rep,name=scopes,proto3" json:"scopes,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *OAuth2Flow) Reset() {
	*x = OAuth2Flow{}
	if protoimpl.UnsafeEnabled {
		mi := &file_options_openapiv3_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OAuth2Flow) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OAuth2Flow) ProtoMessage() {}

func (x *OAuth2Flow) ProtoReflect() protoreflect.Message {
	mi := &file_options_openapiv3_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OAuth2Flow.ProtoReflect.Descriptor instead.
func (*OAuth2Flow) Descriptor() ([]byte, []int) {
	return file_options_openapiv3_proto_rawDescGZIP(), []int{5}
}

func (x *OAuth2Flow) GetAuthorizationUrl() string {
	if x != nil {
		return x.AuthorizationUrl
	}
	return ""
}

func (x *OAuth2Flow) GetTokenUrl() string {
	if x != nil {
		return x.TokenUrl
	}
	return ""
}

func (x *OAuth2Flow) GetRefreshUrl() string {
	if x != nil {
		return x.RefreshUrl
	}
	return ""
}

func (x *OAuth2Flow) GetScopes() map[string]string {
	if x != nil {
		return x.Scopes
	}
	return nil
}

// OAuth2 Flows object
type OAuth2Flows struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Configuration for the OAuth2 Implicit flow
	Implicit *OAuth2Flow `protobuf:"bytes,1,opt,name=implicit,proto3" json:"implicit,omitempty"`
	// Configuration for the OAuth2 Authorization Code flow
	AuthorizationCode *OAuth2Flow `protobuf:"bytes,2,opt,name=authorization_code,json=authorizationCode,proto3" json:"authorization_code,omitempty"`
	// Configuration for the OAuth2 Client Credentials flow
	ClientCredentials *OAuth2Flow `protobuf:"bytes,3,opt,name=client_credentials,json=clientCredentials,proto3" json:"client_credentials,omitempty"`
	// Configuration for the OAuth2 Password flow
	Password *OAuth2Flow `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *OAuth2Flows) Reset() {
	*x = OAuth2Flows{}
	if protoimpl.UnsafeEnabled {
		mi := &file_options_openapiv3_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OAuth2Flows) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OAuth2Flows) ProtoMessage() {}

func (x *OAuth2Flows) ProtoReflect() protoreflect.Message {
	mi := &file_options_openapiv3_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OAuth2Flows.ProtoReflect.Descriptor instead.
func (*OAuth2Flows) Descriptor() ([]byte, []int) {
	return file_options_openapiv3_proto_rawDescGZIP(), []int{6}
}

func (x *OAuth2Flows) GetImplicit() *OAuth2Flow {
	if x != nil {
		return x.Implicit
	}
	return nil
}

func (x *OAuth2Flows) GetAuthorizationCode() *OAuth2Flow {
	if x != nil {
		return x.AuthorizationCode
	}
	return nil
}

func (x *OAuth2Flows) GetClientCredentials() *OAuth2Flow {
	if x != nil {
		return x.ClientCredentials
	}
	return nil
}

func (x *OAuth2Flows) GetPassword() *OAuth2Flow {
	if x != nil {
		return x.Password
	}
	return nil
}

// Security Scheme object
type SecurityScheme struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The type of the security scheme.
	Type string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	// A description for security scheme.
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	// The name of the header, query or cookie parameter to be used.
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// The location of the API key.
	In string `protobuf:"bytes,4,opt,name=in,proto3" json:"in,omitempty"`
	// The name of the HTTP Authentication scheme.
	Scheme string `protobuf:"bytes,5,opt,name=scheme,proto3" json:"scheme,omitempty"`
	// A hint to the client to identify how the bearer token is formatted.
	BearerFormat string `protobuf:"bytes,6,opt,name=bearer_format,json=bearerFormat,proto3" json:"bearer_format,omitempty"`
	// An object containing configuration information for the flow types supported.
	Flows *OAuth2Flows `protobuf:"bytes,7,opt,name=flows,proto3" json:"flows,omitempty"`
	// Well-known URL to discover the OpenID Connect provider metadata.
	OpenIdConnectUrl string `protobuf:"bytes,8,opt,name=open_id_connect_url,json=openIdConnectUrl,proto3" json:"open_id_connect_url,omitempty"`
}

func (x *SecurityScheme) Reset() {
	*x = SecurityScheme{}
	if protoimpl.UnsafeEnabled {
		mi := &file_options_openapiv3_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SecurityScheme) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SecurityScheme) ProtoMessage() {}

func (x *SecurityScheme) ProtoReflect() protoreflect.Message {
	mi := &file_options_openapiv3_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SecurityScheme.ProtoReflect.Descriptor instead.
func (*SecurityScheme) Descriptor() ([]byte, []int) {
	return file_options_openapiv3_proto_rawDescGZIP(), []int{7}
}

func (x *SecurityScheme) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *SecurityScheme) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *SecurityScheme) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SecurityScheme) GetIn() string {
	if x != nil {
		return x.In
	}
	return ""
}

func (x *SecurityScheme) GetScheme() string {
	if x != nil {
		return x.Scheme
	}
	return ""
}

func (x *SecurityScheme) GetBearerFormat() string {
	if x != nil {
		return x.BearerFormat
	}
	return ""
}

func (x *SecurityScheme) GetFlows() *OAuth2Flows {
	if x != nil {
		return x.Flows
	}
	return nil
}

func (x *SecurityScheme) GetOpenIdConnectUrl() string {
	if x != nil {
		return x.OpenIdConnectUrl
	}
	return ""
}

// Security Requirement object
type SecurityRequirement struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the security scheme.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The scopes required for the security scheme.
	Scopes []string `protobuf:"bytes,2,rep,name=scopes,proto3" json:"scopes,omitempty"`
}

func (x *SecurityRequirement) Reset() {
	*x = SecurityRequirement{}
	if protoimpl.UnsafeEnabled {
		mi := &file_options_openapiv3_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SecurityRequirement) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SecurityRequirement) ProtoMessage() {}

func (x *SecurityRequirement) ProtoReflect() protoreflect.Message {
	mi := &file_options_openapiv3_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SecurityRequirement.ProtoReflect.Descriptor instead.
func (*SecurityRequirement) Descriptor() ([]byte, []int) {
	return file_options_openapiv3_proto_rawDescGZIP(), []int{8}
}

func (x *SecurityRequirement) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SecurityRequirement) GetScopes() []string {
	if x != nil {
		return x.Scopes
	}
	return nil
}

// OpenAPI Operation object
type Operation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Summary     string   `protobuf:"bytes,1,opt,name=summary,proto3" json:"summary,omitempty"`
	Description string   `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Tags        []string `protobuf:"bytes,3,rep,name=tags,proto3" json:"tags,omitempty"`
	Produces    []string `protobuf:"bytes,4,rep,name=produces,proto3" json:"produces,omitempty"`
	Consumes    []string `protobuf:"bytes,5,rep,name=consumes,proto3" json:"consumes,omitempty"`
	// TODO repeated Parameter parameters = 6;
	// TODO repeated Response responses = 7;
	Deprecated bool                   `protobuf:"varint,8,opt,name=deprecated,proto3" json:"deprecated,omitempty"`
	Security   []*SecurityRequirement `protobuf:"bytes,9,rep,name=security,proto3" json:"security,omitempty"` // TODO RequestBody request_body = 10;
}

func (x *Operation) Reset() {
	*x = Operation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_options_openapiv3_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Operation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Operation) ProtoMessage() {}

func (x *Operation) ProtoReflect() protoreflect.Message {
	mi := &file_options_openapiv3_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Operation.ProtoReflect.Descriptor instead.
func (*Operation) Descriptor() ([]byte, []int) {
	return file_options_openapiv3_proto_rawDescGZIP(), []int{9}
}

func (x *Operation) GetSummary() string {
	if x != nil {
		return x.Summary
	}
	return ""
}

func (x *Operation) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Operation) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *Operation) GetProduces() []string {
	if x != nil {
		return x.Produces
	}
	return nil
}

func (x *Operation) GetConsumes() []string {
	if x != nil {
		return x.Consumes
	}
	return nil
}

func (x *Operation) GetDeprecated() bool {
	if x != nil {
		return x.Deprecated
	}
	return false
}

func (x *Operation) GetSecurity() []*SecurityRequirement {
	if x != nil {
		return x.Security
	}
	return nil
}

var File_options_openapiv3_proto protoreflect.FileDescriptor

var file_options_openapiv3_proto_rawDesc = []byte{
	0x0a, 0x17, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70,
	0x69, 0x76, 0x33, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1c, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x5f, 0x67, 0x65, 0x6e, 0x5f, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x33, 0x2e,
	0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x45, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x61,
	0x63, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x2f,
	0x0a, 0x07, 0x4c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a,
	0x03, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x22,
	0x84, 0x02, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20,
	0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x28, 0x0a, 0x10, 0x74, 0x65, 0x72, 0x6d, 0x73, 0x5f, 0x6f, 0x66, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x74, 0x65, 0x72, 0x6d,
	0x73, 0x4f, 0x66, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3f, 0x0a, 0x07, 0x63, 0x6f,
	0x6e, 0x74, 0x61, 0x63, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x5f, 0x67, 0x65, 0x6e, 0x5f, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69,
	0x76, 0x33, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x61,
	0x63, 0x74, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x12, 0x3f, 0x0a, 0x07, 0x6c,
	0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x5f, 0x67, 0x65, 0x6e, 0x5f, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70,
	0x69, 0x76, 0x33, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x4c, 0x69, 0x63, 0x65,
	0x6e, 0x73, 0x65, 0x52, 0x07, 0x6c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x60, 0x0a, 0x0e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x56, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x65, 0x6e, 0x75, 0x6d,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x65, 0x6e, 0x75, 0x6d, 0x12, 0x18, 0x0a, 0x07,
	0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x64,
	0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xfb, 0x01, 0x0a, 0x06, 0x53, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x51, 0x0a, 0x09, 0x76, 0x61, 0x72, 0x69, 0x61,
	0x62, 0x6c, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x5f, 0x67, 0x65, 0x6e, 0x5f, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76,
	0x33, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2e, 0x56, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x09, 0x76, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x1a, 0x6a, 0x0a, 0x0e, 0x56, 0x61,
	0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x42,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x5f, 0x67, 0x65, 0x6e, 0x5f, 0x6f, 0x70, 0x65, 0x6e, 0x61,
	0x70, 0x69, 0x76, 0x33, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x53, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x56, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x80, 0x02, 0x0a, 0x0a, 0x4f, 0x41, 0x75, 0x74, 0x68,
	0x32, 0x46, 0x6c, 0x6f, 0x77, 0x12, 0x2b, 0x0a, 0x11, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69,
	0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x10, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x55,
	0x72, 0x6c, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x75, 0x72, 0x6c, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x55, 0x72, 0x6c, 0x12,
	0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x55, 0x72, 0x6c,
	0x12, 0x4c, 0x0a, 0x06, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x34, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x5f, 0x67, 0x65, 0x6e, 0x5f, 0x6f, 0x70,
	0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x33, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x4f, 0x41, 0x75, 0x74, 0x68, 0x32, 0x46, 0x6c, 0x6f, 0x77, 0x2e, 0x53, 0x63, 0x6f, 0x70, 0x65,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x73, 0x1a, 0x39,
	0x0a, 0x0b, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xcb, 0x02, 0x0a, 0x0b, 0x4f, 0x41,
	0x75, 0x74, 0x68, 0x32, 0x46, 0x6c, 0x6f, 0x77, 0x73, 0x12, 0x44, 0x0a, 0x08, 0x69, 0x6d, 0x70,
	0x6c, 0x69, 0x63, 0x69, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x5f, 0x67, 0x65, 0x6e, 0x5f, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69,
	0x76, 0x33, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x4f, 0x41, 0x75, 0x74, 0x68,
	0x32, 0x46, 0x6c, 0x6f, 0x77, 0x52, 0x08, 0x69, 0x6d, 0x70, 0x6c, 0x69, 0x63, 0x69, 0x74, 0x12,
	0x57, 0x0a, 0x12, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x5f, 0x67, 0x65, 0x6e, 0x5f, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69,
	0x76, 0x33, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x4f, 0x41, 0x75, 0x74, 0x68,
	0x32, 0x46, 0x6c, 0x6f, 0x77, 0x52, 0x11, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x57, 0x0a, 0x12, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x5f, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x5f, 0x67, 0x65,
	0x6e, 0x5f, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x33, 0x2e, 0x6f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x4f, 0x41, 0x75, 0x74, 0x68, 0x32, 0x46, 0x6c, 0x6f, 0x77, 0x52, 0x11,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c,
	0x73, 0x12, 0x44, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x5f, 0x67, 0x65, 0x6e,
	0x5f, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x33, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x4f, 0x41, 0x75, 0x74, 0x68, 0x32, 0x46, 0x6c, 0x6f, 0x77, 0x52, 0x08, 0x70,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x97, 0x02, 0x0a, 0x0e, 0x53, 0x65, 0x63, 0x75,
	0x72, 0x69, 0x74, 0x79, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x20,
	0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d,
	0x62, 0x65, 0x61, 0x72, 0x65, 0x72, 0x5f, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x62, 0x65, 0x61, 0x72, 0x65, 0x72, 0x46, 0x6f, 0x72, 0x6d, 0x61,
	0x74, 0x12, 0x3f, 0x0a, 0x05, 0x66, 0x6c, 0x6f, 0x77, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x29, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x5f, 0x67, 0x65, 0x6e, 0x5f, 0x6f, 0x70,
	0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x33, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x4f, 0x41, 0x75, 0x74, 0x68, 0x32, 0x46, 0x6c, 0x6f, 0x77, 0x73, 0x52, 0x05, 0x66, 0x6c, 0x6f,
	0x77, 0x73, 0x12, 0x2d, 0x0a, 0x13, 0x6f, 0x70, 0x65, 0x6e, 0x5f, 0x69, 0x64, 0x5f, 0x63, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x10, 0x6f, 0x70, 0x65, 0x6e, 0x49, 0x64, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x55, 0x72,
	0x6c, 0x22, 0x41, 0x0a, 0x13, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x52, 0x65, 0x71,
	0x75, 0x69, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x63, 0x6f, 0x70, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x73, 0x63,
	0x6f, 0x70, 0x65, 0x73, 0x22, 0x82, 0x02, 0x0a, 0x09, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x12, 0x20, 0x0a, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12,
	0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x74, 0x61,
	0x67, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x65, 0x73, 0x18, 0x04,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x65, 0x73, 0x12, 0x1a,
	0x0a, 0x08, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x08, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x64, 0x65,
	0x70, 0x72, 0x65, 0x63, 0x61, 0x74, 0x65, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a,
	0x64, 0x65, 0x70, 0x72, 0x65, 0x63, 0x61, 0x74, 0x65, 0x64, 0x12, 0x4d, 0x0a, 0x08, 0x73, 0x65,
	0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x5f, 0x67, 0x65, 0x6e, 0x5f, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70,
	0x69, 0x76, 0x33, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x53, 0x65, 0x63, 0x75,
	0x72, 0x69, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52,
	0x08, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x42, 0x2e, 0x5a, 0x2c, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x61, 0x70, 0x6b, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76,
	0x33, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_options_openapiv3_proto_rawDescOnce sync.Once
	file_options_openapiv3_proto_rawDescData = file_options_openapiv3_proto_rawDesc
)

func file_options_openapiv3_proto_rawDescGZIP() []byte {
	file_options_openapiv3_proto_rawDescOnce.Do(func() {
		file_options_openapiv3_proto_rawDescData = protoimpl.X.CompressGZIP(file_options_openapiv3_proto_rawDescData)
	})
	return file_options_openapiv3_proto_rawDescData
}

var file_options_openapiv3_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_options_openapiv3_proto_goTypes = []interface{}{
	(*Contact)(nil),             // 0: protoc_gen_openapiv3.options.Contact
	(*License)(nil),             // 1: protoc_gen_openapiv3.options.License
	(*Info)(nil),                // 2: protoc_gen_openapiv3.options.Info
	(*ServerVariable)(nil),      // 3: protoc_gen_openapiv3.options.ServerVariable
	(*Server)(nil),              // 4: protoc_gen_openapiv3.options.Server
	(*OAuth2Flow)(nil),          // 5: protoc_gen_openapiv3.options.OAuth2Flow
	(*OAuth2Flows)(nil),         // 6: protoc_gen_openapiv3.options.OAuth2Flows
	(*SecurityScheme)(nil),      // 7: protoc_gen_openapiv3.options.SecurityScheme
	(*SecurityRequirement)(nil), // 8: protoc_gen_openapiv3.options.SecurityRequirement
	(*Operation)(nil),           // 9: protoc_gen_openapiv3.options.Operation
	nil,                         // 10: protoc_gen_openapiv3.options.Server.VariablesEntry
	nil,                         // 11: protoc_gen_openapiv3.options.OAuth2Flow.ScopesEntry
}
var file_options_openapiv3_proto_depIdxs = []int32{
	0,  // 0: protoc_gen_openapiv3.options.Info.contact:type_name -> protoc_gen_openapiv3.options.Contact
	1,  // 1: protoc_gen_openapiv3.options.Info.license:type_name -> protoc_gen_openapiv3.options.License
	10, // 2: protoc_gen_openapiv3.options.Server.variables:type_name -> protoc_gen_openapiv3.options.Server.VariablesEntry
	11, // 3: protoc_gen_openapiv3.options.OAuth2Flow.scopes:type_name -> protoc_gen_openapiv3.options.OAuth2Flow.ScopesEntry
	5,  // 4: protoc_gen_openapiv3.options.OAuth2Flows.implicit:type_name -> protoc_gen_openapiv3.options.OAuth2Flow
	5,  // 5: protoc_gen_openapiv3.options.OAuth2Flows.authorization_code:type_name -> protoc_gen_openapiv3.options.OAuth2Flow
	5,  // 6: protoc_gen_openapiv3.options.OAuth2Flows.client_credentials:type_name -> protoc_gen_openapiv3.options.OAuth2Flow
	5,  // 7: protoc_gen_openapiv3.options.OAuth2Flows.password:type_name -> protoc_gen_openapiv3.options.OAuth2Flow
	6,  // 8: protoc_gen_openapiv3.options.SecurityScheme.flows:type_name -> protoc_gen_openapiv3.options.OAuth2Flows
	8,  // 9: protoc_gen_openapiv3.options.Operation.security:type_name -> protoc_gen_openapiv3.options.SecurityRequirement
	3,  // 10: protoc_gen_openapiv3.options.Server.VariablesEntry.value:type_name -> protoc_gen_openapiv3.options.ServerVariable
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_options_openapiv3_proto_init() }
func file_options_openapiv3_proto_init() {
	if File_options_openapiv3_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_options_openapiv3_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Contact); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_options_openapiv3_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*License); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_options_openapiv3_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Info); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_options_openapiv3_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServerVariable); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_options_openapiv3_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Server); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_options_openapiv3_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OAuth2Flow); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_options_openapiv3_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OAuth2Flows); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_options_openapiv3_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SecurityScheme); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_options_openapiv3_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SecurityRequirement); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_options_openapiv3_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Operation); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_options_openapiv3_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_options_openapiv3_proto_goTypes,
		DependencyIndexes: file_options_openapiv3_proto_depIdxs,
		MessageInfos:      file_options_openapiv3_proto_msgTypes,
	}.Build()
	File_options_openapiv3_proto = out.File
	file_options_openapiv3_proto_rawDesc = nil
	file_options_openapiv3_proto_goTypes = nil
	file_options_openapiv3_proto_depIdxs = nil
}
