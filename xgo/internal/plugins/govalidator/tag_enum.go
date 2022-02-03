package govalidator

import (
	"fmt"
	"reflect"

	"github.com/yu31/protoc-plugin/xgo/internal/utils"
	"github.com/yu31/protoc-plugin/xgo/pb/pbvalidator"
	"github.com/yu31/protoc-plugin/xgo/pkg/protovalidator"
	"google.golang.org/protobuf/compiler/protogen"
)

func (p *plugin) loadEnumTags(field *protogen.Field, tagOptions *pbvalidator.TagOptions) *pbvalidator.EnumTags {
	if tagOptions == nil || tagOptions.Kind == nil {
		return nil
	}

	switch ot := tagOptions.Kind.(type) {
	case *pbvalidator.TagOptions_Enum:
		return ot.Enum
	default:
		p.exitWithMsg(
			"%s: types <enum> only support the kind of TagOptions <enum>; and you provided: <%s>",
			p.buildIdentifierWithField(field), reflect.TypeOf(ot).Elem().Name(),
		)
	}
	return nil
}

func (p *plugin) processEnumTags(fieldInfo *FieldInfo) []*protovalidator.TagInfo {
	options := p.loadEnumTags(fieldInfo.Field, fieldInfo.TagOptions)
	if options == nil {
		return nil
	}

	isPointer := utils.FieldIsPointer(fieldInfo.Field)
	itemName := p.getGoItemName(fieldInfo.Field)

	var tagInfos []*protovalidator.TagInfo
	var cond string

	getFieldValue := func() string {
		var filedValue string

		if isPointer {
			convertMethod := p.g.QualifiedGoIdent(validatorPackage.Ident("EnumPointerToString"))
			filedValue = fmt.Sprintf("%s(%s)", convertMethod, itemName)
		} else {
			convertMethod := p.g.QualifiedGoIdent(validatorPackage.Ident("Int32ToString"))
			filedValue = fmt.Sprintf("%s(int32(%s))", convertMethod, itemName)
		}
		return filedValue
	}

	if options.Eq != nil {
		value := *options.Eq
		if isPointer {
			cond = fmt.Sprintf("%s != nil && *%s == %d", itemName, itemName, value)
		} else {
			cond = fmt.Sprintf("%s == %d", itemName, value)
		}
		tagInfos = append(tagInfos, &protovalidator.TagInfo{Tag: protovalidator.TagEnumEq, Cond: cond, Value: value, FieldValue: getFieldValue()})
	}
	if options.Ne != nil {
		value := *options.Ne
		if isPointer {
			cond = fmt.Sprintf("%s != nil && *%s != %d", itemName, itemName, value)
		} else {
			cond = fmt.Sprintf("%s != %d", itemName, value)
		}
		tagInfos = append(tagInfos, &protovalidator.TagInfo{Tag: protovalidator.TagEnumNe, Cond: cond, Value: value, FieldValue: getFieldValue()})
	}
	if options.Gt != nil {
		value := *options.Gt
		if isPointer {
			cond = fmt.Sprintf("%s != nil && *%s > %d", itemName, itemName, value)
		} else {
			cond = fmt.Sprintf("%s > %d", itemName, value)
		}
		tagInfos = append(tagInfos, &protovalidator.TagInfo{Tag: protovalidator.TagEnumGt, Cond: cond, Value: value, FieldValue: getFieldValue()})
	}
	if options.Lt != nil {
		value := *options.Lt
		if isPointer {
			cond = fmt.Sprintf("%s != nil && *%s < %d", itemName, itemName, value)
		} else {
			cond = fmt.Sprintf("%s < %d", itemName, value)
		}
		tagInfos = append(tagInfos, &protovalidator.TagInfo{Tag: protovalidator.TagEnumLt, Cond: cond, Value: value, FieldValue: getFieldValue()})
	}
	if options.Gte != nil {
		value := *options.Gte
		if isPointer {
			cond = fmt.Sprintf("%s != nil && *%s >= %d", itemName, itemName, value)
		} else {
			cond = fmt.Sprintf("%s >= %d", itemName, value)
		}
		tagInfos = append(tagInfos, &protovalidator.TagInfo{Tag: protovalidator.TagEnumGte, Cond: cond, Value: value, FieldValue: getFieldValue()})
	}
	if options.Lte != nil {
		value := *options.Lte
		if isPointer {
			cond = fmt.Sprintf("%s != nil && *%s <= %d", itemName, itemName, value)
		} else {
			cond = fmt.Sprintf("%s <= %d", itemName, value)
		}
		tagInfos = append(tagInfos, &protovalidator.TagInfo{Tag: protovalidator.TagEnumLte, Cond: cond, Value: value, FieldValue: getFieldValue()})
	}

	if len(options.In) != 0 {
		varName := p.buildVariableNameForTagIn(fieldInfo)
		if isPointer {
			cond = fmt.Sprintf("%s != nil && %s[*%s]", itemName, varName, itemName)
		} else {
			cond = fmt.Sprintf("%s[%s]", varName, itemName)
		}
		tagInfos = append(tagInfos, &protovalidator.TagInfo{Tag: protovalidator.TagEnumIn, Cond: cond, Value: options.In, FieldValue: getFieldValue()})
	}
	if len(options.NotIn) != 0 {
		varName := p.buildVariableNameForTagNotIn(fieldInfo)
		if isPointer {
			cond = fmt.Sprintf("%s != nil && !%s[*%s]", itemName, varName, itemName)
		} else {
			cond = fmt.Sprintf("!%s[%s]", varName, itemName)
		}
		tagInfos = append(tagInfos, &protovalidator.TagInfo{Tag: protovalidator.TagEnumNotIn, Cond: cond, Value: options.NotIn, FieldValue: getFieldValue()})
	}

	if options.InEnums != nil && *options.InEnums {
		varName := p.buildVariableNameForTagInEnums(fieldInfo)
		if isPointer {
			cond = fmt.Sprintf("%s != nil && %s[*%s]", itemName, varName, itemName)
		} else {
			cond = fmt.Sprintf("%s[%s]", varName, itemName)
		}
		value := make([]int32, 0, len(fieldInfo.Field.Enum.Values))
		for _, vv := range fieldInfo.Field.Enum.Values {
			value = append(value, int32(vv.Desc.Number()))
		}

		tagInfos = append(tagInfos, &protovalidator.TagInfo{Tag: protovalidator.TagEnumInEnums, Cond: cond, Value: value, FieldValue: getFieldValue()})
	}

	return tagInfos
}
