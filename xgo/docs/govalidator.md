# govalidator

Generated code for validate field for protobuf message.

## Dependency
```bash
github.com/golang/protobuf v1.5.2
```

## Installation
```bash
go get -u -d github.com/yu31/proto-go-plugin

go install github.com/yu31/proto-go-plugin/cmd/protoc-gen-govalidator
```

## Example

The proto file see [govalidator_test.proto](../tests/govalidatortest/govalidator_test.proto)

And generated by
```bash
protoc -I=. -I="${GOPATH}"/src --go_opt=paths=source_relative --govalidator_opt=paths=source_relative --go_out=. --govalidator_out=. ./tests/govalidatortest/govalidator_test.proto

```

The code generated see [govalidator_test.validator.pb.go](../tests/govalidatortest/govalidator_test.validator.pb.go)
