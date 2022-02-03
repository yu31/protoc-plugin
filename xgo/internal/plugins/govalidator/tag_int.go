package govalidator

import (
	"fmt"
	"reflect"

	"github.com/yu31/protoc-plugin/xgo/internal/utils"
	"github.com/yu31/protoc-plugin/xgo/pb/pbvalidator"
	"github.com/yu31/protoc-plugin/xgo/pkg/protovalidator"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func (p *plugin) loadIntTags(field *protogen.Field, tagOptions *pbvalidator.TagOptions) *pbvalidator.IntTags {
	if tagOptions == nil || tagOptions.Kind == nil {
		return nil
	}

	switch ot := tagOptions.Kind.(type) {
	case *pbvalidator.TagOptions_Int:
		return ot.Int
	default:
		p.exitWithMsg(
			"%s: types <int32/int64/sint32/sint64/sfixed32/sfixed64> only support the kind of TagOptions <int>; and you provided: <%s>",
			p.buildIdentifierWithField(field), reflect.TypeOf(ot).Elem().Name(),
		)
	}
	return nil
}

func (p *plugin) processIntTags(fieldInfo *FieldInfo) []*protovalidator.TagInfo {
	options := p.loadIntTags(fieldInfo.Field, fieldInfo.TagOptions)
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
			case protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Sfixed32Kind:
				convertMethod = p.g.QualifiedGoIdent(validatorPackage.Ident("Int32PointerToString"))
			case protoreflect.Int64Kind, protoreflect.Sint64Kind, protoreflect.Sfixed64Kind:
				convertMethod = p.g.QualifiedGoIdent(validatorPackage.Ident("Int64PointerToString"))
			}
		} else {
			switch fieldInfo.Field.Desc.Kind() {
			case protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Sfixed32Kind:
				convertMethod = p.g.QualifiedGoIdent(validatorPackage.Ident("Int32ToString"))
			case protoreflect.Int64Kind, protoreflect.Sint64Kind, protoreflect.Sfixed64Kind:
				convertMethod = p.g.QualifiedGoIdent(validatorPackage.Ident("Int64ToString"))
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
		tagInfos = append(tagInfos, &protovalidator.TagInfo{Tag: protovalidator.TagIntEq, Cond: cond, Value: value, FieldValue: getFieldValue()})
	}
	if options.Ne != nil {
		value := *options.Ne
		if isPointer {
			cond = fmt.Sprintf("%s != nil && *%s != %d", itemName, itemName, value)
		} else {
			cond = fmt.Sprintf("%s != %d", itemName, value)
		}
		tagInfos = append(tagInfos, &protovalidator.TagInfo{Tag: protovalidator.TagIntNe, Cond: cond, Value: value, FieldValue: getFieldValue()})
	}
	if options.Gt != nil {
		value := *options.Gt
		if isPointer {
			cond = fmt.Sprintf("%s != nil && *%s > %d", itemName, itemName, value)
		} else {
			cond = fmt.Sprintf("%s > %d", itemName, value)
		}
		tagInfos = append(tagInfos, &protovalidator.TagInfo{Tag: protovalidator.TagIntGt, Cond: cond, Value: value, FieldValue: getFieldValue()})
	}
	if options.Lt != nil {
		value := *options.Lt
		if isPointer {
			cond = fmt.Sprintf("%s != nil && *%s < %d", itemName, itemName, value)
		} else {
			cond = fmt.Sprintf("%s < %d", itemName, value)
		}
		tagInfos = append(tagInfos, &protovalidator.TagInfo{Tag: protovalidator.TagIntLt, Cond: cond, Value: value, FieldValue: getFieldValue()})
	}
	if options.Gte != nil {
		value := *options.Gte
		if isPointer {
			cond = fmt.Sprintf("%s != nil && *%s >= %d", itemName, itemName, value)
		} else {
			cond = fmt.Sprintf("%s >= %d", itemName, value)
		}
		tagInfos = append(tagInfos, &protovalidator.TagInfo{Tag: protovalidator.TagIntGte, Cond: cond, Value: value, FieldValue: getFieldValue()})
	}
	if options.Lte != nil {
		value := *options.Lte
		if isPointer {
			cond = fmt.Sprintf("%s != nil && *%s <= %d", itemName, itemName, value)
		} else {
			cond = fmt.Sprintf("%s <= %d", itemName, value)
		}
		tagInfos = append(tagInfos, &protovalidator.TagInfo{Tag: protovalidator.TagIntLte, Cond: cond, Value: value, FieldValue: getFieldValue()})
	}

	if len(options.In) != 0 {
		varName := p.buildVariableNameForTagIn(fieldInfo)
		if isPointer {
			cond = fmt.Sprintf("%s != nil && %s[*%s]", itemName, varName, itemName)
		} else {
			cond = fmt.Sprintf("%s[%s]", varName, itemName)
		}
		tagInfos = append(tagInfos, &protovalidator.TagInfo{Tag: protovalidator.TagIntIn, Cond: cond, Value: options.In, FieldValue: getFieldValue()})
	}
	if len(options.NotIn) != 0 {
		varName := p.buildVariableNameForTagNotIn(fieldInfo)
		if isPointer {
			cond = fmt.Sprintf("%s != nil && !%s[*%s]", itemName, varName, itemName)
		} else {
			cond = fmt.Sprintf("!%s[%s]", varName, itemName)
		}
		tagInfos = append(tagInfos, &protovalidator.TagInfo{Tag: protovalidator.TagIntNotIn, Cond: cond, Value: options.NotIn, FieldValue: getFieldValue()})
	}

	return tagInfos
}
