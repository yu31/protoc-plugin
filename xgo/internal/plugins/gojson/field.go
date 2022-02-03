package gojson

import (
	"fmt"

	"github.com/yu31/protoc-plugin/xgo/pb/pbjson"

	"google.golang.org/protobuf/compiler/protogen"
)

func (p *plugin) getFieldKey(fieldOptions *pbjson.FieldOptions, field *protogen.Field) string {
	msgOptions := p.msgOptions

	key := fieldOptions.Json
	if *fieldOptions.Ignore || (key != nil && *key == "-") {
		panic("this field should be ignore.")
	}

	if key != nil {
		return *key
	}

	var k string
	switch *msgOptions.NameStyle {
	case pbjson.NameStyle_TextName:
		k = field.Desc.TextName()
	case pbjson.NameStyle_GoName:
		k = field.GoName
	case pbjson.NameStyle_JSONName:
		k = field.Desc.JSONName()
	default:
		panic(fmt.Sprintf("gojson: unsupported NameStyle type [%s] in field options", msgOptions.NameStyle.String()))
	}
	return k
}

func (p *plugin) getOneOfKey(oneofOptions *pbjson.OneofOptions, oneof *protogen.Oneof) string {
	msgOptions := p.msgOptions
	key := oneofOptions.Json
	if *oneofOptions.Ignore || (key != nil && *key == "-") {
		panic("this oneof field should be ignore.")
	}

	if key != nil {
		return *key
	}

	var k string
	switch *msgOptions.NameStyle {
	case pbjson.NameStyle_TextName:
		k = string(oneof.Desc.Name())
	case pbjson.NameStyle_GoName:
		k = oneof.GoName
	case pbjson.NameStyle_JSONName:
		k = string(oneof.Desc.Name())
	default:
		panic(fmt.Sprintf("gojson: unsupported NameStyle type [%s] in oneof options", msgOptions.NameStyle.String()))
	}
	return k
}

func (p *plugin) guessBufLength(fields []*protogen.Field) int {
	n := 0

	for _, field := range fields {
		var jsonKey string
		if field.Oneof != nil {
			options := p.loadOneOfOptions(field.Oneof)
			if *options.Ignore {
				continue
			}
			jsonKey = p.getOneOfKey(options, field.Oneof)
		} else {
			options := p.loadFieldOptions(field)
			if *options.Ignore {
				continue
			}
			jsonKey = p.getFieldKey(options, field)
		}

		// Sum key length.
		n += len(jsonKey)
		// Sum key/value separator(':')
		n += 1
		// Sum field separator(',')
		n += 1

		// Sum value of length.
	}

	n *= 2

	// Sum object begin and end marker('{', '}').
	n += 2
	return n
}

//func (p *plugin) guessValueLength(field *protogen.Field) int {
//	n := 0
//
//	if field.Desc.IsMap() {
//		switch field.Desc.MapKey().Kind() {
//		case protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Sfixed32Kind,
//			protoreflect.Uint32Kind, protoreflect.Fixed32Kind:
//			n += 10
//		case protoreflect.Int64Kind, protoreflect.Sint64Kind, protoreflect.Sfixed64Kind,
//			protoreflect.Uint64Kind, protoreflect.Fixed64Kind:
//			n += 20
//		case protoreflect.StringKind:
//			n += 8
//		}
//	}
//
//	var kind protoreflect.Kind
//	if field.Desc.IsMap() {
//		kind = field.Desc.MapValue().Kind()
//	} else {
//		kind = field.Desc.Kind()
//	}
//
//	switch kind {
//	case protoreflect.BoolKind:
//		n += 4
//	case protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Sfixed32Kind,
//		protoreflect.Uint32Kind, protoreflect.Fixed32Kind,
//		protoreflect.FloatKind:
//		n += 10
//	case protoreflect.Int64Kind, protoreflect.Sint64Kind, protoreflect.Sfixed64Kind,
//		protoreflect.Uint64Kind, protoreflect.Fixed64Kind,
//		protoreflect.DoubleKind:
//		n += 20
//	case protoreflect.StringKind, protoreflect.BytesKind:
//		n += 8
//	case protoreflect.EnumKind:
//		n += 10
//	case protoreflect.MessageKind:
//		n += p.guessBufLength(field.Message.Fields)
//	}
//
//	return n
//}
