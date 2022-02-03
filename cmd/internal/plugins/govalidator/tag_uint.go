package govalidator

import (
	"fmt"

	"github.com/yu31/protoc-plugin/cmd/internal/generator/utils"
	"github.com/yu31/protoc-plugin/xgo/pb/pbvalidator"
	"github.com/yu31/protoc-plugin/xgo/pkg/protovalidator"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func (p *plugin) loadUintTags(field *protogen.Field, tagOptions *pbvalidator.TagOptions) *pbvalidator.UintTags {
	if tagOptions == nil || tagOptions.Kind == nil {
		return nil
	}
	ot, ok := tagOptions.Kind.(*pbvalidator.TagOptions_Uint)
	if !ok {
		p.exitWithMsg(
			"%s: types uint32/uint64/fixed32/fixed64 only support tags kind <uint>",
			p.buildIdentifierWithField(field),
		)
	}
	return ot.Uint
}

func (p *plugin) processUintTags(fieldInfo *FieldInfo) []*protovalidator.TagInfo {
	options := p.loadUintTags(fieldInfo.Field, fieldInfo.TagOptions)
	if options == nil {
		return nil
	}

	isPointer := utils.FieldIsPointer(fieldInfo.Field)
	itemName := p.getGoItemName(fieldInfo.Field)

	var tagInfos []*protovalidator.TagInfo
	var cond string

	getFieldValue := func() string {
		var convertMethod string

		if isPointer {
			switch fieldInfo.Field.Desc.Kind() {
			case protoreflect.Uint32Kind, protoreflect.Fixed32Kind:
				convertMethod = p.g.QualifiedGoIdent(validatorPackage.Ident("Uint32PointerToString"))
			case protoreflect.Uint64Kind, protoreflect.Fixed64Kind:
				convertMethod = p.g.QualifiedGoIdent(validatorPackage.Ident("Uint64PointerToString"))
			}
		} else {
			switch fieldInfo.Field.Desc.Kind() {
			case protoreflect.Uint32Kind, protoreflect.Fixed32Kind:
				convertMethod = p.g.QualifiedGoIdent(validatorPackage.Ident("Uint32ToString"))
			case protoreflect.Uint64Kind, protoreflect.Fixed64Kind:
				convertMethod = p.g.QualifiedGoIdent(validatorPackage.Ident("Uint64ToString"))
			}
		}

		filedValue := fmt.Sprintf("%s(%s)", convertMethod, itemName)
		return filedValue
	}

	if options.Eq != nil {
		value := *options.Eq
		if isPointer {
			cond = fmt.Sprintf("%s != nil && *%s == %d", itemName, itemName, value)
		} else {
			cond = fmt.Sprintf("%s == %d", itemName, value)
		}
		tagInfos = append(tagInfos, &protovalidator.TagInfo{Tag: protovalidator.TagUintEq, Cond: cond, Value: value, FieldValue: getFieldValue()})
	}
	if options.Ne != nil {
		value := *options.Ne
		if isPointer {
			cond = fmt.Sprintf("%s != nil && *%s != %d", itemName, itemName, value)
		} else {
			cond = fmt.Sprintf("%s != %d", itemName, value)
		}
		tagInfos = append(tagInfos, &protovalidator.TagInfo{Tag: protovalidator.TagUintNe, Cond: cond, Value: value, FieldValue: getFieldValue()})
	}
	if options.Gt != nil {
		value := *options.Gt
		if isPointer {
			cond = fmt.Sprintf("%s != nil && *%s > %d", itemName, itemName, value)
		} else {
			cond = fmt.Sprintf("%s > %d", itemName, value)
		}
		tagInfos = append(tagInfos, &protovalidator.TagInfo{Tag: protovalidator.TagUintGt, Cond: cond, Value: value, FieldValue: getFieldValue()})
	}
	if options.Lt != nil {
		value := *options.Lt
		if isPointer {
			cond = fmt.Sprintf("%s != nil && *%s < %d", itemName, itemName, value)
		} else {
			cond = fmt.Sprintf("%s < %d", itemName, value)
		}
		tagInfos = append(tagInfos, &protovalidator.TagInfo{Tag: protovalidator.TagUintLt, Cond: cond, Value: value, FieldValue: getFieldValue()})
	}
	if options.Gte != nil {
		value := *options.Gte
		if isPointer {
			cond = fmt.Sprintf("%s != nil && *%s >= %d", itemName, itemName, value)
		} else {
			cond = fmt.Sprintf("%s >= %d", itemName, value)
		}
		tagInfos = append(tagInfos, &protovalidator.TagInfo{Tag: protovalidator.TagUintGte, Cond: cond, Value: value, FieldValue: getFieldValue()})
	}
	if options.Lte != nil {
		value := *options.Lte
		if isPointer {
			cond = fmt.Sprintf("%s != nil && *%s <= %d", itemName, itemName, value)
		} else {
			cond = fmt.Sprintf("%s <= %d", itemName, value)
		}
		tagInfos = append(tagInfos, &protovalidator.TagInfo{Tag: protovalidator.TagUintLte, Cond: cond, Value: value, FieldValue: getFieldValue()})
	}

	if len(options.In) != 0 {
		varName := p.buildVariableNameForTagIn(fieldInfo)
		if isPointer {
			cond = fmt.Sprintf("%s != nil && %s[*%s]", itemName, varName, itemName)
		} else {
			cond = fmt.Sprintf("%s[%s]", varName, itemName)
		}
		tagInfos = append(tagInfos, &protovalidator.TagInfo{Tag: protovalidator.TagUintIn, Cond: cond, Value: options.In, FieldValue: getFieldValue()})
	}
	if len(options.NotIn) != 0 {
		varName := p.buildVariableNameForTagNotIn(fieldInfo)
		if isPointer {
			cond = fmt.Sprintf("%s != nil && !%s[*%s]", itemName, varName, itemName)
		} else {
			cond = fmt.Sprintf("!%s[%s]", varName, itemName)
		}
		tagInfos = append(tagInfos, &protovalidator.TagInfo{Tag: protovalidator.TagUintNotIn, Cond: cond, Value: options.NotIn, FieldValue: getFieldValue()})
	}

	return tagInfos
}
