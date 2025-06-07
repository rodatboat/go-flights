Compile protobuf schema
> protoc --proto_path=. --go_out=. ./internal/schema.proto


> go build -o bin/