package govalidator

import (
	"fmt"
	"reflect"

	"github.com/yu31/protoc-plugin/xgo/pb/pbvalidator"
	"github.com/yu31/protoc-plugin/xgo/pkg/protovalidator"
	"google.golang.org/protobuf/compiler/protogen"
)

func (p *plugin) loadOneOfTags(field *protogen.Field, tagOptions *pbvalidator.TagOptions) *pbvalidator.OneOfTags {
	if tagOptions == nil || tagOptions.Kind == nil {
		return nil
	}
	switch ot := tagOptions.Kind.(type) {
	case *pbvalidator.TagOptions_Oneof:
		return ot.Oneof
	default:
		p.exitWithMsg(
			"%s: types <oneof> only support the kind of TagOptions <oneof>; and you provided: <%s>",
			p.buildIdentifierWithOneOf(field), reflect.TypeOf(ot).Elem().Name(),
		)
	}
	return nil
}

func (p *plugin) processOnoOfTags(fieldInfo *FieldInfo) {
	options := p.loadOneOfTags(fieldInfo.Field, fieldInfo.TagOptions)
	if options == nil {
		return
	}

	itemName := "this." + fieldInfo.Field.Oneof.GoName

	var optionInfos []*protovalidator.TagInfo
	var cond string

	if options.NotNull != nil && *options.NotNull {
		cond = fmt.Sprintf("%s != nil", itemName)
		optionInfos = append(optionInfos, &protovalidator.TagInfo{Tag: protovalidator.TagOneOfNotNull, Cond: cond, Value: nil, FieldValue: ""})
	}

	p.genCodeWithTagInfos(fieldInfo, optionInfos)
}
