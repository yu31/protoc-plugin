package govalidator

import (
	"fmt"
	"reflect"

	"github.com/yu31/protoc-plugin/xgo/pb/pbvalidator"
	"github.com/yu31/protoc-plugin/xgo/pkg/protovalidator"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func (p *plugin) loadMapTags(field *protogen.Field, tagOptions *pbvalidator.TagOptions) *pbvalidator.MapTags {
	if tagOptions == nil || tagOptions.Kind == nil {
		return nil
	}

	switch ot := tagOptions.Kind.(type) {
	case *pbvalidator.TagOptions_Map:
		return ot.Map
	default:
		p.exitWithMsg(
			"%s: types <map> only support the kind of TagOptions <map>; and you provided: <%s>",
			p.buildIdentifierWithField(field), reflect.TypeOf(ot).Elem().Name(),
		)
	}
	return nil
}

func (p *plugin) processMapTags(fieldInfo *FieldInfo) {
	p.processMapValue(fieldInfo)
	p.processMapItems(fieldInfo)
}

func (p *plugin) processMapValue(fieldInfo *FieldInfo) {
	options := p.loadMapTags(fieldInfo.Field, fieldInfo.TagOptions)
	if options == nil {
		return
	}

	itemName := "this." + fieldInfo.Field.GoName

	var tagInfos []*protovalidator.TagInfo
	var cond string

	getFieldValue := func() string {
		convertMethod := p.g.QualifiedGoIdent(strconvPackage.Ident("Itoa"))
		filedValue := fmt.Sprintf("%s(len(%s))", convertMethod, itemName)
		return filedValue
	}

	if options.NotNull != nil && *options.NotNull {
		cond = fmt.Sprintf("%s != nil", itemName)
		tagInfos = append(tagInfos, &protovalidator.TagInfo{Tag: protovalidator.TagMapNotNull, Cond: cond, Value: nil, FieldValue: ""})
	}
	if options.LenEq != nil {
		value := *options.LenEq
		cond = fmt.Sprintf("len(%s) == %d", itemName, value)
		tagInfos = append(tagInfos, &protovalidator.TagInfo{Tag: protovalidator.TagMapLenEq, Cond: cond, Value: value, FieldValue: getFieldValue()})
	}
	if options.LenNe != nil {
		value := *options.LenNe
		cond = fmt.Sprintf("len(%s) != %d", itemName, value)
		tagInfos = append(tagInfos, &protovalidator.TagInfo{Tag: protovalidator.TagMapLenNe, Cond: cond, Value: value, FieldValue: getFieldValue()})
	}
	if options.LenGt != nil {
		value := *options.LenGt
		cond = fmt.Sprintf("len(%s) > %d", itemName, value)
		tagInfos = append(tagInfos, &protovalidator.TagInfo{Tag: protovalidator.TagMapLenGt, Cond: cond, Value: value, FieldValue: getFieldValue()})
	}
	if options.LenLt != nil {
		value := *options.LenLt
		cond = fmt.Sprintf("len(%s) < %d", itemName, value)
		tagInfos = append(tagInfos, &protovalidator.TagInfo{Tag: protovalidator.TagMapLenLt, Cond: cond, Value: value, FieldValue: getFieldValue()})
	}
	if options.LenGte != nil {
		value := *options.LenGte
		cond = fmt.Sprintf("len(%s) >= %d", itemName, value)
		tagInfos = append(tagInfos, &protovalidator.TagInfo{Tag: protovalidator.TagMapLenGte, Cond: cond, Value: value, FieldValue: getFieldValue()})
	}
	if options.LenLte != nil {
		value := *options.LenLte
		cond = fmt.Sprintf("len(%s) <= %d", itemName, value)
		tagInfos = append(tagInfos, &protovalidator.TagInfo{Tag: protovalidator.TagMapLenLte, Cond: cond, Value: value, FieldValue: getFieldValue()})
	}

	p.genCodeWithTagInfos(fieldInfo, tagInfos)
}

func (p *plugin) processMapItems(fieldInfo *FieldInfo) {
	options := p.loadMapTags(fieldInfo.Field, fieldInfo.TagOptions)

	itemName := "this." + fieldInfo.Field.GoName

	if options != nil && options.Key != nil && options.Key.Kind != nil {
		subField := fieldInfo.Field.Message.Fields[0]
		subFieldInfo := &FieldInfo{
			CheckIf: fieldInfo.CheckIf,
			//Name:    string(subField.Desc.Name()),
			Name:       fieldInfo.Name,
			Field:      subField,
			IsOneOf:    false,
			InOneOf:    false,
			TagOptions: options.Key,
			IsCheckIf:  fieldInfo.IsCheckIf,
			Parent:     fieldInfo.Parent,
			IsListItem: false,
			IsMapKey:   true,
			IsMapValue: false,
		}

		tagInfos := p.getTagInfos(subFieldInfo)
		if len(tagInfos) != 0 {
			p.g.P("for item := range ", itemName, " {")
			p.g.P("    _ = item // To avoid unused panics.")

			p.genCodeWithTagInfos(subFieldInfo, tagInfos)

			p.g.P("}")
		}
	}

	isMessage := fieldInfo.Field.Desc.MapValue().Kind() == protoreflect.MessageKind

	if (options != nil && options.Value != nil && options.Value.Kind != nil) || isMessage {
		var subTagOptions *pbvalidator.TagOptions
		if options != nil {
			subTagOptions = options.Value
		}

		subField := fieldInfo.Field.Message.Fields[1]
		subFieldInfo := &FieldInfo{
			CheckIf: fieldInfo.CheckIf,
			//Name:    string(subField.Desc.Name()),
			Name:       fieldInfo.Name,
			Field:      subField,
			IsOneOf:    false,
			InOneOf:    false,
			TagOptions: subTagOptions,
			IsCheckIf:  fieldInfo.IsCheckIf,
			Parent:     fieldInfo.Parent,
			IsListItem: false,
			IsMapKey:   false,
			IsMapValue: true,
		}

		// checkMessage indicates whether need to invoker the message's Validate method.
		checkMessage := isMessage
		if isMessage {
			checkMessage = p.isNeedCheckMessage(subFieldInfo)
		}

		tagInfos := p.getTagInfos(subFieldInfo)
		if len(tagInfos) != 0 || checkMessage {
			p.g.P("for _, item := range ", itemName, " {")
			p.g.P("    _ = item // To avoid unused panics.")

			p.genCodeWithTagInfos(subFieldInfo, tagInfos)

			p.g.P("}")
		}
	}
}
