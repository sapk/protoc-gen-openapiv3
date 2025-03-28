package options

//go:generate protoc --go_out=. --go_opt=paths=source_relative openapiv3.proto
//go:generate protoc --go_out=. --go_opt=paths=source_relative annotations.proto
