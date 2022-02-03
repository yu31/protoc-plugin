package govalidator

import (
	"fmt"
	"reflect"

	"github.com/yu31/protoc-plugin/cmd/internal/generator/utils"
	"github.com/yu31/protoc-plugin/xgo/pkg/protovalidator"

	"github.com/yu31/protoc-plugin/xgo/pb/pbvalidator"
	"google.golang.org/protobuf/compiler/protogen"
)

func (p *plugin) loadBoolTags(field *protogen.Field, tagOptions *pbvalidator.TagOptions) *pbvalidator.BoolTags {
	if tagOptions == nil || tagOptions.Kind == nil {
		return nil
	}

	switch ot := tagOptions.Kind.(type) {
	case *pbvalidator.TagOptions_Bool:
		return ot.Bool
	default:
		p.exitWithMsg(
			"%s: types <bool> only support the kind of TagOptions <bool>; and you provided: <%s>",
			p.buildIdentifierWithField(field), reflect.TypeOf(ot).Elem().Name(),
		)
	}
	return nil
}

func (p *plugin) processBoolTags(fieldInfo *FieldInfo) []*protovalidator.TagInfo {
	options := p.loadBoolTags(fieldInfo.Field, fieldInfo.TagOptions)
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
			convertMethod = p.g.QualifiedGoIdent(validatorPackage.Ident("BoolPointerToString"))
		} else {
			convertMethod = p.g.QualifiedGoIdent(validatorPackage.Ident("BoolToString"))
		}

		filedValue := fmt.Sprintf("%s(%s)", convertMethod, itemName)
		return filedValue
	}
	if options.Eq != nil {
		value := *options.Eq
		if isPointer {
			cond = fmt.Sprintf("%s != nil && *%s == %v", itemName, itemName, value)
		} else {
			cond = fmt.Sprintf("%s == %v", itemName, value)
		}
		tagInfos = append(tagInfos, &protovalidator.TagInfo{Tag: protovalidator.TagBoolEq, Cond: cond, Value: value, FieldValue: getFieldValue()})
	}

	return tagInfos
}
