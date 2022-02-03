package utils

import (
	"fmt"

	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
)

// FieldGoType returns the Go type used for a field.
// If it returns pointer=true, the struct field is a pointer to the type.
func FieldGoType(g *protogen.GeneratedFile, field *protogen.Field) (goType string) {
	if field.Desc.IsWeak() {
		//return "struct{}", false
		panic(fmt.Sprintf("unsupported case IsWeak; field: %s", field.Desc.FullName()))
	}

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
	case protoreflect.BytesKind:
		goType = "[]byte"
	case protoreflect.BoolKind:
		goType = "bool"
	case protoreflect.EnumKind:
		goType = g.QualifiedGoIdent(field.Enum.GoIdent)
	case protoreflect.MessageKind, protoreflect.GroupKind:
		goType = "*" + g.QualifiedGoIdent(field.Message.GoIdent)
	}

	switch {
	case field.Desc.IsList():
		return "[]" + goType
	case field.Desc.IsMap():
		keyType := FieldGoType(g, field.Message.Fields[0])
		valType := FieldGoType(g, field.Message.Fields[1])
		return fmt.Sprintf("map[%v]%v", keyType, valType)
	}
	return goType
}

// FieldIsOneOf check the field is oneof
func FieldIsOneOf(field *protogen.Field) bool {
	return field.Oneof != nil && !field.Oneof.Desc.IsSynthetic()
}

// FieldIsPointer check the field is pointer
func FieldIsPointer(field *protogen.Field) bool {
	if field.Desc.IsMap() {
		return false
	}
	if field.Desc.IsList() {
		return false
	}
	if FieldIsOneOf(field) {
		return false
	}
	switch field.Desc.Kind() {
	case protoreflect.BytesKind:
		// rely on nullability of slices for presence
		return false
	case protoreflect.MessageKind, protoreflect.GroupKind:
		// rely on nullability of slices for presence
		return false
	}
	return field.Desc.HasPresence()
}

// LoadFieldList return valid fields list in message.
func LoadFieldList(message *protogen.Message) []*protogen.Field {
	fields := make([]*protogen.Field, 0)
	for _, field := range message.Fields {
		if FieldIsOneOf(field) && field.Oneof.Fields[0] != field {
			continue // only generate for first appearance
		}
		fields = append(fields, field)
	}
	return fields
}
