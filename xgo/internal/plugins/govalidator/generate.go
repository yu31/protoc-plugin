package govalidator

import (
	"fmt"
	"os"
	"runtime"
	"strings"

	"github.com/yu31/protoc-plugin/xgo/internal/utils"
	"github.com/yu31/protoc-plugin/xgo/pb/pbvalidator"
	"github.com/yu31/protoc-plugin/xgo/pkg/protovalidator"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func (p *plugin) generateCode() {
	for _, fieldInfo := range p.filedInfos {
		p.generateCodeForField(fieldInfo)
	}

	p.generateMethodValidate()
}

func (p *plugin) generateCodeForField(fieldInfo *FieldInfo) {
	defer func() {
		if r := recover(); r != nil {
			println(fmt.Sprintf(
				"panic on -> file: %s, message: %s, field: %s",
				p.file.Desc.FullName(), p.message.Desc.Name(), fieldInfo.Field.Desc.Name(),
			))

			println(fmt.Sprintf("recover: %v", r))
			buf := make([]byte, 4096)
			_ = runtime.Stack(buf, true)
			println(string(buf))

			os.Exit(1)
		}
	}()

	p.checkTagOptions(fieldInfo)

	p.generateMethodCheckIf(fieldInfo)
	p.generateMethodCheckError(fieldInfo)
}

func (p *plugin) checkTagOptions(fieldInfo *FieldInfo) {
	checkBasic := func(field *protogen.Field, tagOptions *pbvalidator.TagOptions) {
		switch field.Desc.Kind() {
		case protoreflect.FloatKind, protoreflect.DoubleKind:
			p.loadFloatTags(field, tagOptions)
		case protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Sfixed32Kind,
			protoreflect.Int64Kind, protoreflect.Sint64Kind, protoreflect.Sfixed64Kind:
			p.loadIntTags(field, tagOptions)
		case protoreflect.Uint32Kind, protoreflect.Fixed32Kind,
			protoreflect.Uint64Kind, protoreflect.Fixed64Kind:
			p.loadUintTags(field, tagOptions)
		case protoreflect.StringKind:
			p.loadStringTags(field, tagOptions)
		case protoreflect.BytesKind:
			p.loadBytesTags(field, tagOptions)
		case protoreflect.BoolKind:
			p.loadBoolTags(field, tagOptions)
		case protoreflect.EnumKind:
			p.loadEnumTags(field, tagOptions)
		case protoreflect.MessageKind:
			p.loadMessageTags(field, tagOptions)
		}
	}

	doCheck := func(info *FieldInfo) {
		switch {
		case info.IsOneOf && !info.InOneOf:
			p.loadOneOfTags(info.Field, info.TagOptions)
		case info.Field.Desc.IsList():
			options := p.loadRepeatedTags(info.Field, info.TagOptions)
			if options != nil {
				checkBasic(info.Field, options.Item)
			}
		case info.Field.Desc.IsMap():
			options := p.loadMapTags(info.Field, info.TagOptions)
			if options != nil {
				checkBasic(info.Field.Message.Fields[0], options.Key)
				checkBasic(info.Field.Message.Fields[1], options.Value)
			}
		default:
			checkBasic(info.Field, info.TagOptions)
		}
	}

	doCheck(fieldInfo)

	if fieldInfo.CheckIf == nil {
		return
	}
	doCheck(fieldInfo.CheckIf.Field)
}

func (p *plugin) generateMethodCheckIf(fieldInfo *FieldInfo) {
	if fieldInfo.CheckIf == nil {
		return
	}

	p.generateVariableForField(fieldInfo.CheckIf.Field)

	p.g.P("func (this *", p.message.GoIdent.GoName, ") ", p.buildMethodNameForFieldCheckIf(fieldInfo), "() bool {")

	checkField := fieldInfo.CheckIf.Field

	switch {
	case checkField.InOneOf:
		p.g.P("v, ok := this.", checkField.Field.Oneof.GoName, ".(*", p.g.QualifiedGoIdent(checkField.Field.GoIdent), ")")
		p.g.P("_ = v // To avoid unused panics")
		p.g.P("if !ok {")
		p.g.P("    return false")
		p.g.P("}")
		p.processFieldTags(checkField)
	case checkField.IsOneOf:
		p.processOnoOfTags(checkField)
	default:
		p.processFieldTags(checkField)
	}

	p.g.P("    return true")
	p.g.P("}")
	p.g.P("")
}

func (p *plugin) generateMethodCheckError(fieldInfo *FieldInfo) {
	p.generateVariableForField(fieldInfo)

	p.g.P("func (this *", p.message.GoIdent.GoName, ") ", p.buildMethodNameForFieldValidate(fieldInfo), "() error {")

	if fieldInfo.CheckIf != nil {
		p.g.P("if !this.", p.buildMethodNameForFieldCheckIf(fieldInfo), "() {")
		p.g.P("    return nil")
		p.g.P("}")
	}

	switch {
	case fieldInfo.InOneOf:
		p.g.P("v, ok := this.", fieldInfo.Field.Oneof.GoName, ".(*", p.g.QualifiedGoIdent(fieldInfo.Field.GoIdent), ")")
		p.g.P("_ = v // To avoid unused panics")
		p.g.P("if !ok {")
		p.g.P("    return nil")
		p.g.P("}")
		p.processFieldTags(fieldInfo)
	case fieldInfo.IsOneOf:
		p.processOnoOfTags(fieldInfo)
	default:
		p.processFieldTags(fieldInfo)
	}

	p.g.P("    return nil")
	p.g.P("}")
	p.g.P("")
}

func (p *plugin) generateMethodValidate() {
	msg := p.message

	// Generated Validate Method.
	p.g.P("// Set default value for message ", msg.Desc.FullName())
	p.g.P("func (this *", msg.GoIdent.GoName, ") ", p.getValidateMethodName(), "() error {")
	p.g.P("    if this == nil {")
	p.g.P(`        return nil`)
	p.g.P("    }")

	for _, fieldInfo := range p.filedInfos {
		p.g.P("if err := this.", p.buildMethodNameForFieldValidate(fieldInfo), "(); err != nil {")
		p.g.P("	return err")
		p.g.P("}")
	}

	p.g.P("    return nil")
	p.g.P("}")
	p.g.P("")
}

func (p *plugin) processFieldTags(fieldInfo *FieldInfo) {
	field := fieldInfo.Field
	switch {
	case field.Desc.IsMap():
		p.processMapTags(fieldInfo)
	case field.Desc.IsList():
		p.processListTags(fieldInfo)
	default:
		p.processBasic(fieldInfo)
	}
}

func (p *plugin) getGoItemName(field *protogen.Field) string {
	var itemName string
	switch {
	case field.Parent.Desc.IsMapEntry(), field.Desc.IsList():
		itemName = "item"
	case utils.FieldIsOneOf(field):
		itemName = "v." + field.GoName
	default:
		itemName = "this." + field.GoName
	}
	return itemName
}

func (p *plugin) processBasic(fieldInfo *FieldInfo) {
	tagInfos := p.getTagInfos(fieldInfo)

	p.genCodeWithTagInfos(fieldInfo, tagInfos)
}

func (p *plugin) getTagInfos(fieldInfo *FieldInfo) []*protovalidator.TagInfo {
	var tagInfos []*protovalidator.TagInfo

	switch fieldInfo.Field.Desc.Kind() {
	case protoreflect.FloatKind, protoreflect.DoubleKind:
		tagInfos = p.processFloatTags(fieldInfo)
	case protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Sfixed32Kind,
		protoreflect.Int64Kind, protoreflect.Sint64Kind, protoreflect.Sfixed64Kind:
		tagInfos = p.processIntTags(fieldInfo)
	case protoreflect.Uint32Kind, protoreflect.Fixed32Kind,
		protoreflect.Uint64Kind, protoreflect.Fixed64Kind:
		tagInfos = p.processUintTags(fieldInfo)
	case protoreflect.StringKind:
		tagInfos = p.processStringTags(fieldInfo)
	case protoreflect.BytesKind:
		tagInfos = p.processBytesTags(fieldInfo)
	case protoreflect.BoolKind:
		tagInfos = p.processBoolTags(fieldInfo)
	case protoreflect.EnumKind:
		tagInfos = p.processEnumTags(fieldInfo)
	case protoreflect.MessageKind:
		tagInfos = p.processMessageTags(fieldInfo)
	default:
		panic(fmt.Sprintf("%s: unsupported case: %s", p.buildIdentifierWithField(fieldInfo.Field), fieldInfo.Field.Desc.Kind()))
	}
	return tagInfos
}

func (p *plugin) genCodeWithTagInfos(fieldInfo *FieldInfo, tagInfos []*protovalidator.TagInfo) {
	var fieldDescX string
	switch {
	case fieldInfo.IsListItem:
		fieldDescX = fmt.Sprintf("array item where in field '%s'", fieldInfo.Name)
	case fieldInfo.IsMapKey:
		fieldDescX = fmt.Sprintf("map key where in field '%s'", fieldInfo.Name)
	case fieldInfo.IsMapValue:
		fieldDescX = fmt.Sprintf("map value where in field '%s'", fieldInfo.Name)
	default:
		fieldDescX = fmt.Sprintf("field '%s'", fieldInfo.Name)
	}

	for _, tagInfo := range tagInfos {
		var cond string
		if strings.HasPrefix(tagInfo.Cond, "!") {
			cond = strings.TrimPrefix(tagInfo.Cond, "!")
		} else {
			cond = fmt.Sprintf("!(%s)", tagInfo.Cond)
		}

		var retValue string

		if fieldInfo.IsCheckIf {
			retValue = "false"
		} else {
			reason := protovalidator.BuildErrorReason(tagInfo, fieldDescX)

			var errFunc string

			fieldValueX := tagInfo.FieldValue
			if fieldValueX != "" {
				errFunc = p.g.QualifiedGoIdent(validatorPackage.Ident("FieldError1"))
				retValue = fmt.Sprintf(`%s("%s","%s", %s)`, errFunc, p.message.GoIdent.GoName, reason, fieldValueX)
			} else {
				errFunc = p.g.QualifiedGoIdent(validatorPackage.Ident("FieldError2"))
				retValue = fmt.Sprintf(`%s("%s","%s")`, errFunc, p.message.GoIdent.GoName, reason)
			}

		}

		p.g.P("if ", cond, " {")
		p.g.P("    return ", retValue)
		p.g.P("}")
	}

	if fieldInfo.IsCheckIf {
		return
	}
	if fieldInfo.Field.Desc.Kind() != protoreflect.MessageKind {
		return
	}
	if fieldInfo.Field.Desc.IsList() && !fieldInfo.IsListItem {
		return
	}
	if fieldInfo.Field.Desc.IsMap() && !fieldInfo.IsMapValue && !fieldInfo.IsMapKey {
		return
	}
	if !p.isNeedCheckMessage(fieldInfo) {
		return
	}

	//p.g.P("if err := ", p.getGoItemName(fieldInfo.Field), ".", p.getValidateMethodName(), "; err != nil {")
	//p.g.P("    return err")
	//p.g.P("}")

	p.g.P("if dt, ok := interface{}(", p.getGoItemName(fieldInfo.Field), ").(interface {", p.getValidateMethodName(), "() error }); ok {")
	p.g.P("   if err := dt.", p.getValidateMethodName(), "(); err != nil {")
	p.g.P("       return err")
	p.g.P("   }")
	p.g.P("}")
}
