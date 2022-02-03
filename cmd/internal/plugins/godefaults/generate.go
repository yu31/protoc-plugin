package godefaults

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/yu31/protoc-plugin/cmd/internal/generator/utils"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func (p *plugin) generateCode() {
	msg := p.message

	p.g.P("// Set default value for message ", msg.Desc.FullName())
	p.g.P("func (this *", msg.GoIdent.GoName, ") ", p.getMethodName(), "() {")
	p.g.P("    if this == nil {")
	p.g.P(`        return `)
	p.g.P("    }")

	for _, field := range p.fields {
		isOneOf := utils.FieldIsOneOf(field)
		if !isOneOf {
			p.checkFieldOptions(field)
		}

		switch {
		case field.Desc.IsMap():
			p.setMap(field)
		case field.Desc.IsList():
			p.setList(field)
		case isOneOf:
			p.setOneOf(field)
		default:
			p.setBasic(field)
		}
	}

	p.g.P("    return")
	p.g.P("}")
}

func (p *plugin) convertToString(field *protogen.Field, s string) string {
	if field.Desc.Kind() != protoreflect.StringKind {
		return s
	}
	return strconv.Quote(s)
	//return `"` + strings.ReplaceAll(s, `"`, `\"`) + `"`
}

func (p *plugin) setMap(field *protogen.Field) {
	options := p.loadFieldOptions(field)
	if options == nil || options.Map == nil {
		return
	}

	valueSet := options.Map
	goType := utils.FieldGoType(p.g, field)

	// Sorted key to keep the results consistent.
	keys := make([]string, 0, len(valueSet))
	for k := range valueSet {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var s strings.Builder
	s.WriteString(goType)
	s.WriteString("{")

	for i, k := range keys {
		v := valueSet[k]

		p.checkValue(field.Message.Fields[0], p.buildIdentifierWithField(field), goType, k)
		k = p.convertToString(field.Message.Fields[0], k)

		p.checkValue(field.Message.Fields[1], p.buildIdentifierWithField(field), goType, v)
		v = p.convertToString(field.Message.Fields[1], v)

		s.WriteString(k + ":" + v)
		if i < len(valueSet)-1 {
			s.WriteString(",")
		}
	}
	s.WriteString("}")

	p.g.P("if this.", field.GoName, " == nil", " {")
	p.g.P("    this.", field.GoName, " = ", s.String())
	p.g.P("}")
}

func (p *plugin) setList(field *protogen.Field) {
	options := p.loadFieldOptions(field)
	if options == nil || options.Array == nil {
		return
	}

	valueSet := options.Array
	goType := utils.FieldGoType(p.g, field)

	var s strings.Builder
	s.WriteString(goType)
	s.WriteString("{")

	for i, v := range valueSet {
		p.checkValue(field, p.buildIdentifierWithField(field), goType, v)
		v = p.convertToString(field, v)
		s.WriteString(v)
		if i < len(valueSet)-1 {
			s.WriteString(",")
		}
	}
	s.WriteString("}")

	p.g.P("if this.", field.GoName, " == nil", " {")
	p.g.P("    this.", field.GoName, " = ", s.String())
	p.g.P("}")
}

func (p *plugin) setOneOf(field *protogen.Field) {
	oneOfName := field.Oneof.GoName

	processOneOfFields := func() {
		// process field of oneof.
		p.g.P("switch v := this.", oneOfName, ".(type) {")
		for _, oneOfField := range field.Oneof.Fields {
			p.checkFieldOptions(oneOfField)
			p.g.P("case *", p.g.QualifiedGoIdent(oneOfField.GoIdent), ":")
			p.setBasic(oneOfField)
		}
		p.g.P("default:")
		p.g.P("    _ = v // to avoid unused panic")
		p.g.P("}")
	}

	options := p.loadOneOfOptions(field.Oneof)
	if options == nil || options.Field == nil {
		processOneOfFields()
		return
	}

	valueSet := *options.Field

	var defaultField *protogen.Field
	for _, f := range field.Oneof.Fields {
		if string(f.Desc.Name()) == valueSet {
			defaultField = f
			break
		}
	}
	if defaultField == nil {
		println(
			fmt.Sprintf(
				"%s; cannot found the default field [%s] in oneof fields.",
				p.buildIdentifierWithOneOf(field), *options.Field,
			))
		os.Exit(1)
	}
	oneOfType := p.g.QualifiedGoIdent(defaultField.GoIdent)
	p.g.P("if this.", oneOfName, " == nil", " {")
	p.g.P("    this.", oneOfName, " = ", "new(", oneOfType, ")")
	p.g.P("}")

	processOneOfFields()
}

func (p *plugin) setBasic(field *protogen.Field) {
	isOneOf := utils.FieldIsOneOf(field)
	var itemName string
	if isOneOf {
		itemName = "v" + "." + field.GoName
	} else {
		itemName = "this" + "." + field.GoName
	}

	processMessage := func() {
		if field.Desc.Kind() == protoreflect.MessageKind {
			p.g.P("if ", itemName, " != nil {")
			p.g.P("    if dt, ok := interface{}(", itemName, ").(interface {", p.getMethodName(), "()}); ok {")
			p.g.P("        dt.", p.getMethodName(), "()")
			p.g.P("    }")
			p.g.P("}")
		}
	}

	options := p.loadFieldOptions(field)
	if options == nil || options.Basic == nil {
		processMessage()
		return
	}

	valueSet := *options.Basic
	goType := utils.FieldGoType(p.g, field)
	isPointer := utils.FieldIsPointer(field)

	p.checkValue(field, p.buildIdentifierWithField(field), goType, valueSet)
	valueSet = p.convertToString(field, valueSet)

	var emptyCond string
	switch field.Desc.Kind() {
	case protoreflect.FloatKind, protoreflect.DoubleKind,
		protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Sfixed32Kind,
		protoreflect.Int64Kind, protoreflect.Sint64Kind, protoreflect.Sfixed64Kind,
		protoreflect.Uint32Kind, protoreflect.Fixed32Kind,
		protoreflect.Uint64Kind, protoreflect.Fixed64Kind,
		protoreflect.EnumKind:
		emptyCond = itemName + "== 0"
	case protoreflect.BoolKind:
		emptyCond = "!" + itemName
	case protoreflect.StringKind:
		emptyCond = itemName + `== ""`
	case protoreflect.MessageKind:
		emptyCond = itemName + `== nil`
	}

	if isPointer {
		emptyCond = itemName + "== nil"
	}

	p.g.P("if ", emptyCond, " {")
	if isPointer {
		p.g.P("x := ", goType, "(", valueSet, ")")
		p.g.P(itemName, " = &x")
	} else if field.Desc.Kind() == protoreflect.MessageKind {
		p.g.P(itemName, " = new(", p.g.QualifiedGoIdent(field.Message.GoIdent), ")")
	} else {
		p.g.P(itemName, " = ", valueSet)
	}
	p.g.P("}")

	processMessage()
}
