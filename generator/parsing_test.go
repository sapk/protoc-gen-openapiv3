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
	cmd := exec.Command("protoc", "--descriptor_set_out="+pbFile, "--include_imports", "--proto_path=..", "../testdata/test.proto")
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
	require.Len(t, fdSet.File, 1)

	// Create a test plugin
	gen, err := protogen.Options{}.New(&pluginpb.CodeGeneratorRequest{
		ProtoFile: fdSet.File,
	})
	require.NoError(t, err)

	// Create a generator instance
	oapiGenerator := generator.NewOpenAPIGenerator(gen, &generator.Options{})

	// Parse the proto file
	parsed, err := oapiGenerator.ParseProtoFile(gen.Files[0])
	assert.NoError(t, err)
	assert.NotNil(t, parsed)

	// Verify package name
	assert.Equal(t, "test.package", parsed.Package)

	// Verify service
	assert.Len(t, parsed.Services, 1)
	service := parsed.Services[0]
	assert.Equal(t, "TestService", service.Name)
	assert.Len(t, service.Methods, 1)

	// Verify method
	method := service.Methods[0]
	assert.Equal(t, "TestMethod", method.Name)
	assert.Equal(t, "test.package.TestRequest", method.InputType)
	assert.Equal(t, "test.package.TestResponse", method.OutputType)

	// Verify messages
	assert.Len(t, parsed.Messages, 2)
	request := parsed.Messages[0]
	assert.Equal(t, "TestRequest", request.Name)
	assert.Len(t, request.Fields, 1)
	response := parsed.Messages[1]
	assert.Equal(t, "TestResponse", response.Name)
	assert.Len(t, response.Fields, 1)

	// Verify fields
	field := request.Fields[0]
	assert.Equal(t, "test_field", field.Name)
	assert.Equal(t, "string", field.Type)
	assert.Equal(t, int32(1), field.Number)
}
