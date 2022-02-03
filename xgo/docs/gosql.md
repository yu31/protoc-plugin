# gosql

Generated code to implement sql.Scanner and driver.Valuer for you protobuf message.

## Dependency
```bash
github.com/golang/protobuf v1.5.2
```

## Installation
```bash
go get -u -d github.com/yu31/proto-go-plugin

go install github.com/yu31/proto-go-plugin/cmd/protoc-gen-gosql
```

## Example

The proto file see [gosql_test.proto](../tests/gosqltest/gosql_test.proto)

And generate by
```bash
protoc -I=. -I="${GOPATH}"/src --go_opt=paths=source_relative --gosql_opt=paths=source_relative --go_out=. --gosql_out=. ./tests/gosqltest/gosql_test.proto

```

The code generated see [gosql_test.sql.pb.go](../tests/gosqltest/gosql_test.sql.pb.go)
