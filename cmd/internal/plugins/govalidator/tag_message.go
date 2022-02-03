package govalidator

import (
	"fmt"
	"reflect"

	"github.com/yu31/protoc-plugin/xgo/pb/pbvalidator"
	"github.com/yu31/protoc-plugin/xgo/pkg/protovalidator"
	"google.golang.org/protobuf/compiler/protogen"
)

func (p *plugin) loadMessageTags(field *protogen.Field, tagOptions *pbvalidator.TagOptions) *pbvalidator.MessageTags {
	if tagOptions == nil || tagOptions.Kind == nil {
		return nil
	}

	switch ot := tagOptions.Kind.(type) {
	case *pbvalidator.TagOptions_Message:
		return ot.Message
	default:
		p.exitWithMsg(
			"%s: types <message> only support the kind of TagOptions <message>; and you provided: <%s>",
			p.buildIdentifierWithField(field), reflect.TypeOf(ot).Elem().Name(),
		)
	}
	return nil
}

func (p *plugin) processMessageTags(fieldInfo *FieldInfo) []*protovalidator.TagInfo {
	options := p.loadMessageTags(fieldInfo.Field, fieldInfo.TagOptions)
	if options == nil {
		return nil
	}

	itemName := p.getGoItemName(fieldInfo.Field)

	var tagInfos []*protovalidator.TagInfo
	var cond string

	if options.NotNull != nil && *options.NotNull {
		cond = fmt.Sprintf("%s != nil", itemName)
		tagInfos = append(tagInfos, &protovalidator.TagInfo{Tag: protovalidator.TagMessageNotNull, Cond: cond, Value: nil, FieldValue: ""})
	}

	return tagInfos
}

func (p *plugin) isNeedCheckMessage(fieldInfo *FieldInfo) bool {
	options := p.loadMessageTags(fieldInfo.Field, fieldInfo.TagOptions)
	if options != nil && options.Skip != nil && *options.Skip {
		return false
	}
	return true
}
