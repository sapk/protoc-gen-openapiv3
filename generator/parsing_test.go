package generator_test

import (
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"
	"google.golang.org/protobuf/types/pluginpb"

	"github.com/sapk/protoc-gen-openapiv3/generator"
)

func TestParseProtoFile(t *testing.T) {
	// Create a temporary directory for the test
	tmpDir := t.TempDir()
	pbFile := filepath.Join(tmpDir, "test.pb")

	// Generate descriptor set
	cmd := exec.Command("protoc",
		"--descriptor_set_out="+pbFile,
		"--include_imports",
		"--proto_path=../",
		"--proto_path=../testdata",
		"../testdata/test.proto")
	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("protoc failed: %v\nOutput: %s", err, output)
	}

	// Read the descriptor set
	data, err := os.ReadFile(pbFile)
	if err != nil {
		t.Fatalf("Failed to read descriptor set: %v", err)
	}

	// Parse the file descriptor set
	fdSet := &descriptorpb.FileDescriptorSet{}
	err = proto.Unmarshal(data, fdSet)
	require.NoError(t, err)

	// Find the main proto file and keep track of its index
	mainFileIndex := -1
	for i, file := range fdSet.File {
		if file.GetName() == "testdata/test.proto" {
			mainFileIndex = i
			break
		}
	}
	require.NotEqual(t, -1, mainFileIndex, "Main proto file not found")

	// Create a test plugin with all files but only process the main one
	gen, err := protogen.Options{}.New(&pluginpb.CodeGeneratorRequest{
		ProtoFile:      fdSet.File,
		FileToGenerate: []string{"testdata/test.proto"},
	})
	require.NoError(t, err)

	// Create a generator instance
	oapiGenerator := generator.NewOpenAPIGenerator(gen, &generator.Options{})

	// Parse the proto file
	parsed, err := oapiGenerator.ParseProtoFile(gen.Files[mainFileIndex])
	assert.NoError(t, err)
	assert.NotNil(t, parsed)

	// Verify package name
	assert.Equal(t, "test.package", parsed.Package)

	// Verify service
	assert.Len(t, parsed.Services, 1)
	service := parsed.Services[0]
	assert.Equal(t, "UserService", service.Name)
	assert.Len(t, service.Methods, 6)

	// Verify methods
	methods := make(map[string]generator.ParsedMethod)
	for _, method := range service.Methods {
		methods[method.Name] = method
	}

	// Verify GetUser method
	getUser := methods["GetUser"]
	assert.Equal(t, "test.package.GetUserRequest", getUser.InputType)
	assert.Equal(t, "test.package.User", getUser.OutputType)
	assert.Equal(t, "GET", getUser.HTTPMethod)
	assert.Equal(t, "/v1/users/{user_id}", getUser.HTTPPath)

	// Verify ListUsers method
	listUsers := methods["ListUsers"]
	assert.Equal(t, "test.package.ListUsersRequest", listUsers.InputType)
	assert.Equal(t, "test.package.ListUsersResponse", listUsers.OutputType)
	assert.Equal(t, "GET", listUsers.HTTPMethod)
	assert.Equal(t, "/v1/users", listUsers.HTTPPath)

	// Verify CreateUser method
	createUser := methods["CreateUser"]
	assert.Equal(t, "test.package.CreateUserRequest", createUser.InputType)
	assert.Equal(t, "test.package.User", createUser.OutputType)
	assert.Equal(t, "POST", createUser.HTTPMethod)
	assert.Equal(t, "/v1/users", createUser.HTTPPath)
	assert.Equal(t, "user", createUser.HTTPBody)

	// Verify UpdateUser method
	updateUser := methods["UpdateUser"]
	assert.Equal(t, "test.package.UpdateUserRequest", updateUser.InputType)
	assert.Equal(t, "test.package.User", updateUser.OutputType)
	assert.Equal(t, "PUT", updateUser.HTTPMethod)
	assert.Equal(t, "/v1/users/{user_id}", updateUser.HTTPPath)
	assert.Equal(t, "user", updateUser.HTTPBody)

	// Verify PatchUser method
	patchUser := methods["PatchUser"]
	assert.Equal(t, "test.package.PatchUserRequest", patchUser.InputType)
	assert.Equal(t, "test.package.User", patchUser.OutputType)
	assert.Equal(t, "PATCH", patchUser.HTTPMethod)
	assert.Equal(t, "/v1/users/{user_id}", patchUser.HTTPPath)
	assert.Equal(t, "user", patchUser.HTTPBody)

	// Verify DeleteUser method
	deleteUser := methods["DeleteUser"]
	assert.Equal(t, "test.package.DeleteUserRequest", deleteUser.InputType)
	assert.Equal(t, "google.protobuf.Empty", deleteUser.OutputType)
	assert.Equal(t, "DELETE", deleteUser.HTTPMethod)
	assert.Equal(t, "/v1/users/{user_id}", deleteUser.HTTPPath)

	// Verify messages
	messages := make(map[string]generator.ParsedMessage)
	for _, msg := range parsed.Messages {
		messages[msg.Name] = msg
	}

	// Verify User message
	user := messages["User"]
	assert.Len(t, user.Fields, 9)
	userFields := make(map[string]generator.ParsedField)
	for _, field := range user.Fields {
		userFields[field.Name] = field
	}

	assert.Equal(t, "string", userFields["user_id"].Type)
	assert.Equal(t, "string", userFields["email"].Type)
	assert.Equal(t, "string", userFields["full_name"].Type)
	assert.Equal(t, "test.package.UserStatus", userFields["status"].Type)
	assert.Equal(t, "repeated string", userFields["roles"].Type)
	assert.Equal(t, "optional test.package.Address", userFields["address"].Type)
	assert.Equal(t, "map<string, string>", userFields["metadata"].Type)
	assert.Equal(t, "google.protobuf.Timestamp", userFields["created_at"].Type)
	assert.Equal(t, "google.protobuf.Timestamp", userFields["updated_at"].Type)

	// Verify Address message
	address := messages["Address"]
	assert.Len(t, address.Fields, 5)
	addressFields := make(map[string]generator.ParsedField)
	for _, field := range address.Fields {
		addressFields[field.Name] = field
	}

	assert.Equal(t, "string", addressFields["street"].Type)
	assert.Equal(t, "string", addressFields["city"].Type)
	assert.Equal(t, "string", addressFields["state"].Type)
	assert.Equal(t, "string", addressFields["country"].Type)
	assert.Equal(t, "string", addressFields["postal_code"].Type)

	// Verify enums
	assert.Len(t, parsed.Enums, 1)
	userStatus := parsed.Enums[0]
	assert.Equal(t, "UserStatus", userStatus.Name)
	assert.Len(t, userStatus.Values, 5)

	// Verify enum values
	enumValues := make(map[string]generator.ParsedEnumValue)
	for _, value := range userStatus.Values {
		enumValues[value.Name] = value
	}

	assert.Equal(t, int32(0), enumValues["USER_STATUS_UNSPECIFIED"].Number)
	assert.Equal(t, int32(1), enumValues["USER_STATUS_ACTIVE"].Number)
	assert.Equal(t, int32(2), enumValues["USER_STATUS_INACTIVE"].Number)
	assert.Equal(t, int32(3), enumValues["USER_STATUS_SUSPENDED"].Number)
	assert.Equal(t, int32(4), enumValues["USER_STATUS_DELETED"].Number)
}
