package godefaults

import (
	"fmt"
	"os"
	"strconv"

	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func (p *plugin) checkFieldOptions(field *protogen.Field) {
	{ // check supported kind.
		if field.Desc.IsWeak() {
			println(fmt.Sprintf("%s: [warring] unsupported case IsWeak", p.buildIdentifierWithField(field)))
			os.Exit(1)
		}
		if field.Desc.IsExtension() {
			println(fmt.Sprintf("%s: [warring] unsupported case IsExtension", p.buildIdentifierWithField(field)))
			os.Exit(1)
		}
		if field.Desc.IsPlaceholder() {
			println(fmt.Sprintf("%s: [warring] unsupported case IsPlaceholder", p.buildIdentifierWithField(field)))
			os.Exit(1)
		}

		switch field.Desc.Kind() {
		case protoreflect.FloatKind, protoreflect.DoubleKind,
			protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Sfixed32Kind,
			protoreflect.Int64Kind, protoreflect.Sint64Kind, protoreflect.Sfixed64Kind,
			protoreflect.Uint32Kind, protoreflect.Fixed32Kind,
			protoreflect.Uint64Kind, protoreflect.Fixed64Kind,
			protoreflect.BytesKind, protoreflect.StringKind,
			protoreflect.EnumKind, protoreflect.MessageKind,
			protoreflect.BoolKind:
		default:
			println(fmt.Sprintf(
				"%s: [warring] unsupported kind of %s",
				p.buildIdentifierWithField(field), field.Desc.Kind().String(),
			))
			os.Exit(1)
		}
	}

	{ // check options.
		options := p.loadFieldOptions(field)
		if options == nil {
			return
		}
		switch {
		case field.Desc.IsMap():
			if options.Array != nil {
				println(
					fmt.Sprintf("%s: cannot set option [array] in map field.", p.buildIdentifierWithField(field)),
				)
				os.Exit(1)
			}
			if options.Basic != nil {
				println(
					fmt.Sprintf("%s: cannot set option [basic] in map field.", p.buildIdentifierWithField(field)),
				)
				os.Exit(1)
			}
		case field.Desc.IsList():
			if options.Map != nil {
				println(
					fmt.Sprintf("%s: cannot set option [map] in repeated field.", p.buildIdentifierWithField(field)),
				)
				os.Exit(1)
			}
			if options.Basic != nil {
				println(
					fmt.Sprintf("%s: cannot set option [basic] in repeated field.", p.buildIdentifierWithField(field)),
				)
				os.Exit(1)
			}
		default:
			if options.Map != nil {
				println(
					fmt.Sprintf("%s: cannot set option [map] in lieteral field.", p.buildIdentifierWithField(field)),
				)
				os.Exit(1)
			}
			if options.Array != nil {
				println(
					fmt.Sprintf("%s: cannot set option [array] in lieteral field.", p.buildIdentifierWithField(field)),
				)
				os.Exit(1)
			}
		}
	}
}

func (p *plugin) checkValue(field *protogen.Field, id string, goType string, value string) {
	var err error
	switch field.Desc.Kind() {
	case protoreflect.FloatKind:
		_, err = strconv.ParseFloat(value, 32)
	case protoreflect.DoubleKind:
		_, err = strconv.ParseFloat(value, 64)
	case protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Sfixed32Kind:
		_, err = strconv.ParseInt(value, 10, 32)
	case protoreflect.Int64Kind, protoreflect.Sint64Kind, protoreflect.Sfixed64Kind:
		_, err = strconv.ParseInt(value, 10, 64)
	case protoreflect.Uint32Kind, protoreflect.Fixed32Kind:
		_, err = strconv.ParseUint(value, 10, 32)
	case protoreflect.Uint64Kind, protoreflect.Fixed64Kind:
		_, err = strconv.ParseUint(value, 10, 64)
	case protoreflect.BoolKind:
		_, err = strconv.ParseBool(value)
	case protoreflect.StringKind:
		//if len(value) < 2 || value[0] != '"' || value[len(value)-1] != '"' {
		//	err = errors.New("invalid string format")
		//}
	case protoreflect.BytesKind:
		println(
			fmt.Sprintf("%s: gotype: <%s>, unsupported kind <%s>", id, goType, field.Desc.Kind()),
		)
		os.Exit(1)
	case protoreflect.MessageKind:
		if field.Desc.IsList() || field.Desc.IsMap() {
			println(
				fmt.Sprintf("%s: gotype: <%s>, unsupported kind <%s>", id, goType, field.Desc.Kind()))
			os.Exit(1)
		}
	case protoreflect.EnumKind:
		var v int64
		v, err = strconv.ParseInt(value, 10, 32)
		if err == nil {
			var isValid bool
			for _, vv := range field.Enum.Values {
				if v == int64(vv.Desc.Number()) {
					isValid = true
					break
				}
			}
			if !isValid {
				println(
					fmt.Sprintf("%s: gotype: <%s>; [%s] not a valid enum number", id, goType, value),
				)
				os.Exit(1)
			}
		}
	}

	if err != nil {
		println(
			fmt.Sprintf("%s: gotype: <%s>; cannot convert [%s] to kind of <%s>.", id, goType, value, field.Desc.Kind()))
		os.Exit(1)
	}
}
