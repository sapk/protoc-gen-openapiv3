package options

//go:generate protoc --proto_path=../../ --go_out=../../ --go_opt=paths=source_relative ../../protoc-gen-openapiv3/options/openapiv3.proto
//go:generate protoc --proto_path=../../ --go_out=../../ --go_opt=paths=source_relative ../../protoc-gen-openapiv3/options/annotations.proto
