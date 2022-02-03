package govalidator

import (
	"fmt"
	"reflect"

	"github.com/yu31/protoc-plugin/xgo/pb/pbvalidator"
	"github.com/yu31/protoc-plugin/xgo/pkg/protovalidator"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func (p *plugin) loadRepeatedTags(field *protogen.Field, tagOptions *pbvalidator.TagOptions) *pbvalidator.RepeatedTags {
	if tagOptions == nil || tagOptions.Kind == nil {
		return nil
	}

	switch ot := tagOptions.Kind.(type) {
	case *pbvalidator.TagOptions_Repeated:
		return ot.Repeated
	default:
		p.exitWithMsg(
			"%s: types <repeated> only support the kind of TagOptions <repeated>; and you provided: <%s>",
			p.buildIdentifierWithField(field), reflect.TypeOf(ot).Elem().Name(),
		)
	}
	return nil
}

func (p *plugin) processListTags(fieldInfo *FieldInfo) {
	p.processListValue(fieldInfo)
	p.processListItems(fieldInfo)
}

func (p *plugin) processListValue(fieldInfo *FieldInfo) {
	options := p.loadRepeatedTags(fieldInfo.Field, fieldInfo.TagOptions)
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
		tagInfos = append(tagInfos, &protovalidator.TagInfo{Tag: protovalidator.TagRepeatedNotNull, Cond: cond, Value: nil, FieldValue: ""})
	}
	if options.LenEq != nil {
		value := *options.LenEq
		cond = fmt.Sprintf("len(%s) == %d", itemName, value)
		tagInfos = append(tagInfos, &protovalidator.TagInfo{Tag: protovalidator.TagRepeatedLenEq, Cond: cond, Value: value, FieldValue: getFieldValue()})
	}
	if options.LenNe != nil {
		value := *options.LenNe
		cond = fmt.Sprintf("len(%s) != %d", itemName, value)
		tagInfos = append(tagInfos, &protovalidator.TagInfo{Tag: protovalidator.TagRepeatedLenNe, Cond: cond, Value: value, FieldValue: getFieldValue()})
	}
	if options.LenGt != nil {
		value := *options.LenGt
		cond = fmt.Sprintf("len(%s) > %d", itemName, value)
		tagInfos = append(tagInfos, &protovalidator.TagInfo{Tag: protovalidator.TagRepeatedLenGt, Cond: cond, Value: value, FieldValue: getFieldValue()})
	}
	if options.LenLt != nil {
		value := *options.LenLt
		cond = fmt.Sprintf("len(%s) < %d", itemName, value)
		tagInfos = append(tagInfos, &protovalidator.TagInfo{Tag: protovalidator.TagRepeatedLenLt, Cond: cond, Value: value, FieldValue: getFieldValue()})
	}
	if options.LenGte != nil {
		value := *options.LenGte
		cond = fmt.Sprintf("len(%s) >= %d", itemName, value)
		tagInfos = append(tagInfos, &protovalidator.TagInfo{Tag: protovalidator.TagRepeatedLenGte, Cond: cond, Value: value, FieldValue: getFieldValue()})
	}
	if options.LenLte != nil {
		value := *options.LenLte
		cond = fmt.Sprintf("len(%s) <= %d", itemName, value)
		tagInfos = append(tagInfos, &protovalidator.TagInfo{Tag: protovalidator.TagRepeatedLenLte, Cond: cond, Value: value, FieldValue: getFieldValue()})
	}

	if options.Unique != nil && *options.Unique {
		var uniqueMethod string
		switch fieldInfo.Field.Desc.Kind() {
		case protoreflect.StringKind:
			uniqueMethod = "SliceIsUniqueString"
		case protoreflect.BytesKind:
			uniqueMethod = "SliceIsUniqueBytes"
		case protoreflect.BoolKind:
			uniqueMethod = "SliceIsUniqueBool"
		case protoreflect.FloatKind:
			uniqueMethod = "SliceIsUniqueFloat32"
		case protoreflect.DoubleKind:
			uniqueMethod = "SliceIsUniqueFloat64"
		case protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Sfixed32Kind:
			uniqueMethod = "SliceIsUniqueInt32"
		case protoreflect.Int64Kind, protoreflect.Sint64Kind, protoreflect.Sfixed64Kind:
			uniqueMethod = "SliceIsUniqueInt64"
		case protoreflect.Uint32Kind, protoreflect.Fixed32Kind:
			uniqueMethod = "SliceIsUniqueUint32"
		case protoreflect.Uint64Kind, protoreflect.Fixed64Kind:
			uniqueMethod = "SliceIsUniqueUint64"
		case protoreflect.EnumKind:
			uniqueMethod = "SliceIsUniqueEnum"
		case protoreflect.MessageKind:
			uniqueMethod = "SliceIsUniqueMessage"
		default:
			p.exitWithMsg("%s: unsupported option tag <unique> for kind of %s",
				p.buildIdentifierWithField(fieldInfo.Field), fieldInfo.Field.Desc.Kind())
		}

		cond = fmt.Sprintf("%s(%s)", p.g.QualifiedGoIdent(validatorPackage.Ident(uniqueMethod)), itemName)
		tagInfos = append(tagInfos, &protovalidator.TagInfo{Tag: protovalidator.TagRepeatedUnique, Cond: cond, Value: nil, FieldValue: ""})
	}

	p.genCodeWithTagInfos(fieldInfo, tagInfos)
}

func (p *plugin) processListItems(fieldInfo *FieldInfo) {
	options := p.loadRepeatedTags(fieldInfo.Field, fieldInfo.TagOptions)
	itemName := "this." + fieldInfo.Field.GoName

	isMessage := fieldInfo.Field.Desc.Kind() == protoreflect.MessageKind

	if (options != nil && options.Item != nil && options.Item.Kind != nil) || isMessage {
		var subTagOptions *pbvalidator.TagOptions
		if options != nil {
			subTagOptions = options.Item
		}

		subFieldInfo := &FieldInfo{
			CheckIf:    fieldInfo.CheckIf,
			Name:       fieldInfo.Name,
			Field:      fieldInfo.Field,
			IsOneOf:    false,
			InOneOf:    false,
			TagOptions: subTagOptions,
			IsCheckIf:  fieldInfo.IsCheckIf,
			Parent:     fieldInfo.Parent,
			IsListItem: true,
			IsMapKey:   false,
			IsMapValue: false,
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
