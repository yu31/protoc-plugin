package govalidator

import (
	"fmt"
	"reflect"

	"github.com/yu31/protoc-plugin/xgo/pb/pbvalidator"
	"github.com/yu31/protoc-plugin/xgo/pkg/protovalidator"
	"google.golang.org/protobuf/compiler/protogen"
)

func (p *plugin) loadBytesTags(field *protogen.Field, tagOptions *pbvalidator.TagOptions) *pbvalidator.BytesTags {
	if tagOptions == nil || tagOptions.Kind == nil {
		return nil
	}

	switch ot := tagOptions.Kind.(type) {
	case *pbvalidator.TagOptions_Bytes:
		return ot.Bytes
	default:
		p.exitWithMsg(
			"%s: types <bytes> only support the kind of TagOptions <bytes>; and you provided: <%s>",
			p.buildIdentifierWithField(field), reflect.TypeOf(ot).Elem().Name(),
		)
	}
	return nil
}

func (p *plugin) processBytesTags(fieldInfo *FieldInfo) []*protovalidator.TagInfo {
	options := p.loadBytesTags(fieldInfo.Field, fieldInfo.TagOptions)
	if options == nil {
		return nil
	}

	itemName := p.getGoItemName(fieldInfo.Field)

	var tagInfos []*protovalidator.TagInfo
	var cond string

	getFieldValue := func() string {
		convertMethod := p.g.QualifiedGoIdent(strconvPackage.Ident("Itoa"))
		filedValue := fmt.Sprintf("%s(len(%s))", convertMethod, itemName)
		return filedValue
	}

	if options.LenEq != nil {
		value := *options.LenEq
		cond = fmt.Sprintf("len(%s) == %d", itemName, value)
		tagInfos = append(tagInfos, &protovalidator.TagInfo{Tag: protovalidator.TagBytesLenEq, Cond: cond, Value: value, FieldValue: getFieldValue()})
	}
	if options.LenNe != nil {
		value := *options.LenNe
		cond = fmt.Sprintf("len(%s) != %d", itemName, value)
		tagInfos = append(tagInfos, &protovalidator.TagInfo{Tag: protovalidator.TagBytesLenNe, Cond: cond, Value: value, FieldValue: getFieldValue()})
	}
	if options.LenGt != nil {
		value := *options.LenGt
		cond = fmt.Sprintf("len(%s) > %d", itemName, value)
		tagInfos = append(tagInfos, &protovalidator.TagInfo{Tag: protovalidator.TagBytesLenGt, Cond: cond, Value: value, FieldValue: getFieldValue()})
	}
	if options.LenLt != nil {
		value := *options.LenLt
		cond = fmt.Sprintf("len(%s) < %d", itemName, value)
		tagInfos = append(tagInfos, &protovalidator.TagInfo{Tag: protovalidator.TagBytesLenLt, Cond: cond, Value: value, FieldValue: getFieldValue()})
	}
	if options.LenGte != nil {
		value := *options.LenGte
		cond = fmt.Sprintf("len(%s) >= %d", itemName, value)
		tagInfos = append(tagInfos, &protovalidator.TagInfo{Tag: protovalidator.TagBytesLenGte, Cond: cond, Value: value, FieldValue: getFieldValue()})
	}
	if options.LenLte != nil {
		value := *options.LenLte
		cond = fmt.Sprintf("len(%s) <= %d", itemName, value)
		tagInfos = append(tagInfos, &protovalidator.TagInfo{Tag: protovalidator.TagBytesLenLte, Cond: cond, Value: value, FieldValue: getFieldValue()})
	}

	return tagInfos
}
