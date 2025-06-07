Compile protobuf schema
> protoc --proto_path=. --go_out=. schema.proto


> go build -o bin/