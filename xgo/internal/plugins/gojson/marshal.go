package gojson

import (
	"fmt"

	"github.com/yu31/protoc-plugin/xgo/internal/utils"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func (p *plugin) generateMarshalCode() {
	msg := p.message
	fields := p.fields

	bufLen := p.guessBufLength(fields)

	p.g.P("// MarshalJSON for implements interface json.Marshaler. ")
	p.g.P("func (this *", msg.GoIdent.GoName, ") MarshalJSON() ([]byte, error) {")
	p.g.P("    if this == nil {")
	p.g.P(`        return []byte("null"), nil`)
	p.g.P("    }")
	p.g.P("    var err error")
	p.g.P("")
	// create a new encoder object.
	p.g.P("    encoder := ", encoderPackage.Ident("New"), "(", bufLen, ")")
	p.g.P("")
	p.g.P("    // Add JSON end identifier")
	p.g.P("    encoder.AppendObjectBegin()")
	p.g.P("")

	for _, field := range fields {
		switch {
		case utils.FieldIsOneOf(field):
			p.marshalOneOf(field.Oneof)
		case field.Desc.IsMap():
			p.marshalMap(field)
		case field.Desc.IsList():
			p.marshalList(field)
		default:
			p.marshalBasic(field)
		}
	}

	p.g.P("")
	p.g.P("    // Add JSON end identifier")
	p.g.P("    encoder.AppendObjectEnd()")
	p.g.P("    return encoder.Bytes(), err")

	// End of MarshalJSON.
	p.g.P("}")
}

func (p *plugin) marshalEncodeKey(key string) {
	p.g.P(`encoder.AppendObjectKey("`, key, `")`)
}

func (p *plugin) marshalOneOf(oneof *protogen.Oneof) {
	oneOfOptions := p.loadOneOfOptions(oneof)

	p.g.P("// Encode field type of oneof;",
		" | field: ", oneof.Desc.FullName(),
		" | GoName: ", oneof.GoName,
		" | omitempty: ", *oneOfOptions.Omitempty,
		" | ignore: ", *oneOfOptions.Ignore,
	)

	if *oneOfOptions.Ignore {
		return
	}

	oneOfKey := p.getOneOfKey(oneOfOptions, oneof)

	p.g.P("if this.", oneof.GoName, "!= nil {")
	p.g.P("    switch v := this.", oneof.GoName, ".(type) {")
	for _, field := range oneof.Fields {
		p.g.P("case *", p.g.QualifiedGoIdent(field.GoIdent), ":")
		p.marshalBasic(field)
	}
	p.g.P("    default:")
	p.g.P("        return nil, ", fmtPackage.Ident("Errorf"), `("invalid oneof field type: %v, jsonKey: `, oneOfKey, ", goName: ", oneof.GoName, ", field: ", oneof.Desc.FullName(), `"`, ", v)")
	// end switch
	p.g.P("   }")
	if !(*oneOfOptions.HideOneofKey) && !(*oneOfOptions.Omitempty) {
		p.g.P("} else {")
		p.marshalEncodeKey(oneOfKey)
		p.g.P("    encoder.AppendNil()")
	}
	// end if
	p.g.P("}")
}

func (p *plugin) marshalMap(field *protogen.Field) {
	fieldOptions := p.loadFieldOptions(field)
	p.g.P("// encode field type of map;",
		" | field: ", field.Desc.FullName(),
		" | keyKind: ", field.Desc.MapKey().Kind(),
		" | valueKind: ", field.Desc.MapValue().Kind(),
		" | goName: ", field.GoName,
		" | omitempty: ", *fieldOptions.Omitempty,
		" | ignore: ", *fieldOptions.Ignore,
	)

	if *fieldOptions.Ignore {
		return
	}

	key := p.getFieldKey(fieldOptions, field)

	encodeValue := func() {
		p.g.P("encoder.AppendObjectBegin()")
		p.g.P("for k, v := range ", "this.", field.GoName, " {")
		switch field.Desc.MapKey().Kind() {
		case protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Sfixed32Kind:
			p.g.P("encoder.AppendObjectKey(", strconvPackage.Ident("FormatInt"), "(int64(k), 10)", ")")
		case protoreflect.Int64Kind, protoreflect.Sint64Kind, protoreflect.Sfixed64Kind:
			p.g.P("encoder.AppendObjectKey(", strconvPackage.Ident("FormatInt"), "(k, 10)", ")")
		case protoreflect.Uint32Kind, protoreflect.Fixed32Kind:
			p.g.P("encoder.AppendObjectKey(", strconvPackage.Ident("FormatUint"), "(uint64(k), 10)", ")")
		case protoreflect.Uint64Kind, protoreflect.Fixed64Kind:
			p.g.P("encoder.AppendObjectKey(", strconvPackage.Ident("FormatUint"), "(k, 10)", ")")
		case protoreflect.StringKind:
			p.g.P("encoder.AppendObjectKey(k)")
		default:
			panic(fmt.Sprintf(
				"gojson: marshal: unsupported type of map key, field: %s, kind: %s",
				field.Desc.FullName(), field.Desc.MapKey().Kind().String(),
			))
		}

		p.marshalEncodeValue(field)

		p.g.P("}")
		p.g.P("encoder.AppendObjectEnd()")
	}

	if *fieldOptions.Omitempty {
		p.g.P("if len(this.", field.GoName, ") != 0 {")
		p.marshalEncodeKey(key)
		encodeValue()
		p.g.P("}")
	} else {
		p.marshalEncodeKey(key)
		p.g.P("if this.", field.GoName, "!= nil {")
		encodeValue()
		p.g.P("} else {")
		p.g.P("    encoder.AppendNil()")
		p.g.P("}")
	}
}

func (p *plugin) marshalList(field *protogen.Field) {
	fieldOptions := p.loadFieldOptions(field)
	p.g.P("// encode field type of list;",
		" | field: ", field.Desc.FullName(),
		" | kind:", field.Desc.Kind().GoString(),
		" | goName: ", field.GoName,
		" | omitempty: ", *fieldOptions.Omitempty,
		" | ignore: ", *fieldOptions.Ignore,
	)

	if *fieldOptions.Ignore {
		return
	}

	key := p.getFieldKey(fieldOptions, field)

	encodeValue := func() {
		p.g.P("encoder.AppendListBegin()")
		p.g.P("for i := range ", "this.", field.GoName, " {")
		p.marshalEncodeValue(field)
		p.g.P("}")
		p.g.P("encoder.AppendListEnd()")
	}

	if *fieldOptions.Omitempty {
		p.g.P("if len(this.", field.GoName, ") != 0 {")
		p.marshalEncodeKey(key)
		encodeValue()
		p.g.P("}")
	} else {
		p.marshalEncodeKey(key)
		p.g.P("if this.", field.GoName, "!= nil {")
		encodeValue()
		p.g.P("} else {")
		p.g.P("    encoder.AppendNil()")
		p.g.P("}")
	}
}

func (p *plugin) marshalBasic(field *protogen.Field) {
	// Check support
	{
		if field.Desc.IsMap() {
			panic(fmt.Sprintf("gojson: marshal: unsupported case IsMap; field: %s", field.Desc.FullName()))
		}
		if field.Desc.IsList() {
			panic(fmt.Sprintf("gojson: marshal: unsupported case IsList; field: %s", field.Desc.FullName()))
		}
		if field.Desc.IsWeak() {
			panic(fmt.Sprintf("gojson: marshal: unsupported case IsWeak; field: %s", field.Desc.FullName()))
		}
		if field.Desc.IsExtension() {
			panic(fmt.Sprintf("gojson: marshal: unsupported case IsExtension; filed: %s", field.Desc.FullName()))
		}
		if field.Desc.IsPlaceholder() {
			panic(fmt.Sprintf("gojson: marshal: unsupported case IsPlaceholder; filed: %s", field.Desc.FullName()))
		}
		if field.Desc.IsPacked() {
			panic(fmt.Sprintf("gojson: marshal: unsupported case IsPacked; filed: %s", field.Desc.FullName()))
		}
	}

	options := p.loadFieldOptions(field)
	p.g.P("// encode filed type of basic;",
		" | field: ", field.Desc.FullName(),
		" | kind: ", field.Desc.Kind().GoString(),
		" | GoName: ", field.GoName,
		" | omitempty: ", *options.Omitempty,
		" | ignore: ", *options.Ignore)

	if *options.Ignore {
		return
	}

	key := p.getFieldKey(options, field)

	isPointer := utils.FieldIsPointer(field)
	isOnoOf := utils.FieldIsOneOf(field)

	var notEmptyCond string
	var itemName string

	if isOnoOf {
		itemName = "v." + field.GoName
	} else {
		itemName = "this." + field.GoName
	}

	if isPointer {
		notEmptyCond = itemName + " != nil "
	} else {
		switch field.Desc.Kind() {
		case protoreflect.DoubleKind, protoreflect.FloatKind,
			protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Sfixed32Kind,
			protoreflect.Int64Kind, protoreflect.Sint64Kind, protoreflect.Sfixed64Kind,
			protoreflect.Uint32Kind, protoreflect.Fixed32Kind,
			protoreflect.Uint64Kind, protoreflect.Fixed64Kind:
			notEmptyCond = itemName + " != 0 "
		case protoreflect.BoolKind:
			notEmptyCond = itemName
		case protoreflect.StringKind:
			notEmptyCond = itemName + ` != "" `
		case protoreflect.BytesKind:
			notEmptyCond = "len(" + itemName + ")" + " != 0 "
		case protoreflect.MessageKind:
			notEmptyCond = itemName + " != nil "
		case protoreflect.EnumKind:
			if *options.UseEnumString {
				// Can't omit empty value for type enum field if use enum string.
				ok := false
				options.Omitempty = &ok
				notEmptyCond = ""
			} else {
				notEmptyCond = itemName + " != 0 "
			}
		default:
			panic(fmt.Sprintf("gojson: marshal: unsupported kind of %s, field: %s", field.Desc.Kind().String(), field.Desc.FullName()))
		}
	}

	encodeKeyValue := func() {
		p.marshalEncodeKey(key)
		p.marshalEncodeValue(field)
	}

	encodeField := func() {
		if *options.Omitempty {
			p.g.P("if ", notEmptyCond, " {")
			encodeKeyValue()
			p.g.P("}")
		} else {
			encodeKeyValue()
		}
	}

	if !utils.FieldIsOneOf(field) {
		encodeField()
		return
	}

	oneOfOptions := p.loadOneOfOptions(field.Oneof)
	if *oneOfOptions.HideOneofKey {
		encodeField()
		return
	}

	oneOfKey := p.getOneOfKey(oneOfOptions, field.Oneof)
	encodeOneOf := func() {
		// encode oneof.
		p.marshalEncodeKey(oneOfKey)
		p.g.P("encoder.AppendObjectBegin()")
		// encode field.
		encodeKeyValue()
		// encode oneof..
		p.g.P("encoder.AppendObjectEnd()")
	}

	if *options.Omitempty {
		p.g.P("if ", notEmptyCond, " {")
		encodeOneOf()
		p.g.P("}")
	} else {
		encodeOneOf()
	}
}

func (p *plugin) marshalEncodeValue(field *protogen.Field) {
	options := p.loadFieldOptions(field)

	var itemName string
	switch {
	case field.Desc.IsMap():
		itemName = "v"
	case field.Desc.IsList():
		itemName = "this." + field.GoName + "[i]"
	case utils.FieldIsOneOf(field):
		itemName = "v." + field.GoName
	default:
		itemName = "this." + field.GoName
	}

	if utils.FieldIsPointer(field) && field.Desc.Kind() != protoreflect.EnumKind {
		itemName = "*" + itemName
	}

	if field.Desc.IsMap() {
		field = field.Message.Fields[1]
	}

	switch field.Desc.Kind() {
	case protoreflect.DoubleKind:
		p.g.P("encoder.AppendFloat64(", itemName, ")")
	case protoreflect.FloatKind:
		p.g.P("encoder.AppendFloat32(", itemName, ")")
	case protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Sfixed32Kind:
		p.g.P("encoder.AppendInt32(", itemName, ")")
	case protoreflect.Int64Kind, protoreflect.Sint64Kind, protoreflect.Sfixed64Kind:
		p.g.P("encoder.AppendInt64(", itemName, ")")
	case protoreflect.Uint32Kind, protoreflect.Fixed32Kind:
		p.g.P("encoder.AppendUint32(", itemName, ")")
	case protoreflect.Uint64Kind, protoreflect.Fixed64Kind:
		p.g.P("encoder.AppendUint64(", itemName, ")")
	case protoreflect.BoolKind:
		p.g.P("encoder.AppendBool(", itemName, ")")
	case protoreflect.StringKind:
		p.g.P("encoder.AppendString(", itemName, ")")
	case protoreflect.BytesKind:
		p.g.P("encoder.AppendBytes(", itemName, ")")
	case protoreflect.MessageKind:
		p.g.P("err = encoder.AppendInterface(", itemName, ")")
		p.g.P("if err != nil {")
		p.g.P("    return nil, err")
		p.g.P("}")
	case protoreflect.EnumKind:
		if *options.UseEnumString {
			p.g.P("encoder.AppendString(", itemName, ".String()", ")")
		} else {
			p.g.P("encoder.AppendInt32(int32(", itemName, ".Number()", "))")
		}
	default:
		panic(fmt.Sprintf("gojson: marshal: unsupported kind of %s, field: %s", field.Desc.Kind().String(), field.Desc.FullName()))
	}
}
