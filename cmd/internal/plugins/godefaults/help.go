package godefaults

import (
	"fmt"

	"github.com/yu31/protoc-plugin/xgo/pb/pbdefaults"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
)

func (p *plugin) getMethodName() string {
	// TODO: Supported user-defined method name by command arguments.
	return "SetDefaults"
}

func (p *plugin) buildIdentifierWithField(field *protogen.Field) string {
	return fmt.Sprintf(
		"godefaults: <file(%s) message(%s) field(%s)>",
		string(p.file.GoImportPath), p.message.GoIdent.GoName, field.Desc.Name(),
	)
}

func (p *plugin) buildIdentifierWithOneOf(field *protogen.Field) string {
	return fmt.Sprintf(
		"godefaults: <file(%s) message(%s) field(%s)>",
		string(p.file.GoImportPath), p.message.GoIdent.GoName, field.Oneof.Desc.Name(),
	)
}

func (p *plugin) loadFieldOptions(field *protogen.Field) *pbdefaults.FieldOptions {
	i := proto.GetExtension(field.Desc.Options(), pbdefaults.E_Field)
	fieldOptions := i.(*pbdefaults.FieldOptions)
	return fieldOptions
}

func (p *plugin) loadOneOfOptions(oneOf *protogen.Oneof) *pbdefaults.OneOfOptions {
	i := proto.GetExtension(oneOf.Desc.Options(), pbdefaults.E_Oneof)
	oneOfOptions := i.(*pbdefaults.OneOfOptions)
	return oneOfOptions
}
