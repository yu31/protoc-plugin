package govalidator

import (
	"fmt"
	"os"

	"github.com/yu31/protoc-plugin/xgo/pb/pbvalidator"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func (p *plugin) getValidateMethodName() string {
	// TODO: Supported user-defined method name by command arguments.
	return "Validate"
}

func (p *plugin) buildIdentifierWithField(field *protogen.Field) string {
	name := string(field.Desc.Name())
	if field.Parent.Desc.IsMapEntry() {
		name = string(field.Parent.Desc.Name()) + "_" + name
	}

	return fmt.Sprintf(
		"govalidator: <file(%s) message(%s) field(%s)>",
		string(p.file.GoImportPath), p.message.GoIdent.GoName, name,
	)
}

func (p *plugin) buildIdentifierWithOneOf(field *protogen.Field) string {
	return fmt.Sprintf(
		"govalidator: <file(%s) message(%s) field(%s)>",
		string(p.file.GoImportPath), p.message.GoIdent.GoName, field.Oneof.Desc.Name(),
	)
}

func (p *plugin) buildIdentifierWithName(name string) string {
	return fmt.Sprintf(
		"govalidator: <file(%s) message(%s) field(%s)>",
		string(p.file.GoImportPath), p.message.GoIdent.GoName, name,
	)
}

func (p *plugin) exitWithMsg(format string, a ...interface{}) {
	println(fmt.Sprintf(format, a...))
	os.Exit(1)
}

func (p *plugin) loadValidOptionsFromField(field *protogen.Field) *pbvalidator.ValidOptions {
	i := proto.GetExtension(field.Desc.Options(), pbvalidator.E_Field)
	options := i.(*pbvalidator.ValidOptions)
	if options == nil {
		options = &pbvalidator.ValidOptions{}
	}
	return options
}

func (p *plugin) loadValidOptionsFromOneOf(field *protogen.Field) *pbvalidator.ValidOptions {
	i := proto.GetExtension(field.Oneof.Desc.Options(), pbvalidator.E_Oneof)
	options := i.(*pbvalidator.ValidOptions)
	if options == nil {
		options = &pbvalidator.ValidOptions{}
	}
	return options
}

func (p *plugin) fieldToGoType(field *protogen.Field) string {
	if field.Desc.IsMap() {
		panic("govalidator: unsupported map type in this method.")
	}

	var goType string
	switch field.Desc.Kind() {
	case protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Sfixed32Kind:
		goType = "int32"
	case protoreflect.Uint32Kind, protoreflect.Fixed32Kind:
		goType = "uint32"
	case protoreflect.Int64Kind, protoreflect.Sint64Kind, protoreflect.Sfixed64Kind:
		goType = "int64"
	case protoreflect.Uint64Kind, protoreflect.Fixed64Kind:
		goType = "uint64"
	case protoreflect.FloatKind:
		goType = "float32"
	case protoreflect.DoubleKind:
		goType = "float64"
	case protoreflect.StringKind:
		goType = "string"
	case protoreflect.BoolKind:
		goType = "bool"
	case protoreflect.EnumKind:
		goType = p.g.QualifiedGoIdent(field.Enum.GoIdent)
	case protoreflect.MessageKind, protoreflect.GroupKind:
		goType = "*" + p.g.QualifiedGoIdent(field.Message.GoIdent)
	default:
		panic(fmt.Sprintf("unsupported case, field: %s, kind: %s", field.Desc.Name(), field.Desc.Kind()))
	}

	return goType
}
