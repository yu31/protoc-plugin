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

func (p *plugin) loadFloatTags(field *protogen.Field, tagOptions *pbvalidator.TagOptions) *pbvalidator.FloatTags {
	if tagOptions == nil || tagOptions.Kind == nil {
		return nil
	}

	switch ot := tagOptions.Kind.(type) {
	case *pbvalidator.TagOptions_Float:
		return ot.Float
	default:
		p.exitWithMsg(
			"%s: types <float/double> only support the kind of TagOptions <float>; and you provided: <%s>",
			p.buildIdentifierWithField(field), reflect.TypeOf(ot).Elem().Name(),
		)
	}
	return nil
}

func (p *plugin) processFloatTags(fieldInfo *FieldInfo) []*protovalidator.TagInfo {
	options := p.loadFloatTags(fieldInfo.Field, fieldInfo.TagOptions)
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
			if fieldInfo.Field.Desc.Kind() == protoreflect.FloatKind {
				convertMethod = p.g.QualifiedGoIdent(validatorPackage.Ident("Float32PointerToString"))
			} else {
				convertMethod = p.g.QualifiedGoIdent(validatorPackage.Ident("Float64PointerToString"))
			}
		} else {
			if fieldInfo.Field.Desc.Kind() == protoreflect.FloatKind {
				convertMethod = p.g.QualifiedGoIdent(validatorPackage.Ident("Float32ToString"))
			} else {
				convertMethod = p.g.QualifiedGoIdent(validatorPackage.Ident("Float64ToString"))
			}
		}
		filedValue := fmt.Sprintf("%s(%s)", convertMethod, itemName)
		return filedValue
	}

	if options.Eq != nil {
		value := *options.Eq
		if isPointer {
			cond = fmt.Sprintf("%s != nil && *%s == %f", itemName, itemName, value)
		} else {
			cond = fmt.Sprintf("%s == %f", itemName, value)
		}
		tagInfos = append(tagInfos, &protovalidator.TagInfo{Tag: protovalidator.TagFloatEq, Cond: cond, Value: value, FieldValue: getFieldValue()})
	}
	if options.Ne != nil {
		value := *options.Ne
		if isPointer {
			cond = fmt.Sprintf("%s != nil && *%s != %f", itemName, itemName, value)
		} else {
			cond = fmt.Sprintf("%s != %f", itemName, value)
		}
		tagInfos = append(tagInfos, &protovalidator.TagInfo{Tag: protovalidator.TagFloatNe, Cond: cond, Value: value, FieldValue: getFieldValue()})
	}
	if options.Gt != nil {
		value := *options.Gt
		if isPointer {
			cond = fmt.Sprintf("%s != nil && *%s > %f", itemName, itemName, value)
		} else {
			cond = fmt.Sprintf("%s > %f", itemName, value)
		}
		tagInfos = append(tagInfos, &protovalidator.TagInfo{Tag: protovalidator.TagFloatGt, Cond: cond, Value: value, FieldValue: getFieldValue()})
	}
	if options.Lt != nil {
		value := *options.Lt
		if isPointer {
			cond = fmt.Sprintf("%s != nil && *%s < %f", itemName, itemName, value)
		} else {
			cond = fmt.Sprintf("%s < %f", itemName, value)
		}
		tagInfos = append(tagInfos, &protovalidator.TagInfo{Tag: protovalidator.TagFloatLt, Cond: cond, Value: value, FieldValue: getFieldValue()})
	}
	if options.Gte != nil {
		value := *options.Gte
		if isPointer {
			cond = fmt.Sprintf("%s != nil && *%s >= %f", itemName, itemName, value)
		} else {
			cond = fmt.Sprintf("%s >= %f", itemName, value)
		}
		tagInfos = append(tagInfos, &protovalidator.TagInfo{Tag: protovalidator.TagFloatGte, Cond: cond, Value: value, FieldValue: getFieldValue()})
	}
	if options.Lte != nil {
		value := *options.Lte
		if isPointer {
			cond = fmt.Sprintf("%s != nil && *%s <= %f", itemName, itemName, value)
		} else {
			cond = fmt.Sprintf("%s <= %f", itemName, value)
		}
		tagInfos = append(tagInfos, &protovalidator.TagInfo{Tag: protovalidator.TagFloatLte, Cond: cond, Value: value, FieldValue: getFieldValue()})
	}

	if len(options.In) != 0 {
		varName := p.buildVariableNameForTagIn(fieldInfo)
		if isPointer {
			cond = fmt.Sprintf("%s != nil && %s[*%s]", itemName, varName, itemName)
		} else {
			cond = fmt.Sprintf("%s[%s]", varName, itemName)
		}
		tagInfos = append(tagInfos, &protovalidator.TagInfo{Tag: protovalidator.TagFloatIn, Cond: cond, Value: options.In, FieldValue: getFieldValue()})
	}
	if len(options.NotIn) != 0 {
		varName := p.buildVariableNameForTagNotIn(fieldInfo)
		if isPointer {
			cond = fmt.Sprintf("%s != nil && !%s[*%s]", itemName, varName, itemName)
		} else {
			cond = fmt.Sprintf("!%s[%s]", varName, itemName)
		}
		tagInfos = append(tagInfos, &protovalidator.TagInfo{Tag: protovalidator.TagFloatNotIn, Cond: cond, Value: options.NotIn, FieldValue: getFieldValue()})
	}

	return tagInfos
}
