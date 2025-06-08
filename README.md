Re-compile Protobuf schema
> go generate ./internal/generate

Re-compile IATA codes:
> go generate ./common/iata/generate

Both
> go generate ./internal ./common/iata/generate

Build
> go build -o bin/

User
```Go
package main
func main(){}
```

TODO:
1. Add API call to return JSON data
2. Add price graph functionality