# proto-go-plugin

Implements some helpful plugin for protobuf golang.

Support Plugins:
- [gosql](xgo/docs/gosql.md): Generated code for implements interface sql.Scanner and driver.Valuer. 
- [gojson](xgo/docs/gojson.md): Generated code implements interface json.Marshaler and json.Unmarshaler, Supported oneof.  
- [godefaults](xgo/docs/godefaults.md): Generated code to set default value for message.
- [govalidator](xgo/docs/govalidator.md): Generated code for validate field for message.

References:
 - [protoc-gen-go](google.golang.org/protobuf/cmd/protoc-gen-go/internal_gengo)
