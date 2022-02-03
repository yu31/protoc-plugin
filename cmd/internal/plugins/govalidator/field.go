package govalidator

import (
	"github.com/yu31/protoc-plugin/cmd/internal/generator/utils"
	"github.com/yu31/protoc-plugin/xgo/pb/pbvalidator"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type CheckIfInfo struct {
	Field *FieldInfo
}

type FieldInfo struct {
	CheckIf    *CheckIfInfo
	Name       string
	Field      *protogen.Field
	IsOneOf    bool
	InOneOf    bool
	TagOptions *pbvalidator.TagOptions
	IsCheckIf  bool
	Parent     *FieldInfo
	IsListItem bool
	IsMapKey   bool
	IsMapValue bool
}

func (p *plugin) loadFieldList() {
	plainFieldMap := make(map[string]*protogen.Field)
	oneOfFieldMap := make(map[string]*protogen.Field)
	fields := make([]*protogen.Field, 0)

	for _, field := range p.message.Fields {
		if utils.FieldIsOneOf(field) {
			if field.Oneof.Fields[0] != field {
				continue // only generate for first appearance
			}
			plainFieldMap[string(field.Oneof.Desc.Name())] = field
			for _, oneOfField := range field.Oneof.Fields {
				oneOfFieldMap[string(oneOfField.Desc.Name())] = oneOfField
			}
		} else {
			plainFieldMap[string(field.Desc.Name())] = field
		}
		fields = append(fields, field)
	}

	// build field infos.
	var filedInfos []*FieldInfo

	processField := func(field *protogen.Field, inOneOf bool) {
		var validOptions *pbvalidator.ValidOptions
		var name string

		isOneOf := utils.FieldIsOneOf(field)

		if inOneOf || !isOneOf {
			name = string(field.Desc.Name())
			validOptions = p.loadValidOptionsFromField(field)
		} else {
			name = string(field.Oneof.Desc.Name())
			validOptions = p.loadValidOptionsFromOneOf(field)
		}

		if validOptions.Tags == nil || validOptions.Tags.Kind == nil {
			if field.Desc.IsMap() && field.Desc.MapValue().Kind() != protoreflect.MessageKind {
				return
			}
			if field.Desc.Kind() != protoreflect.MessageKind {
				return
			}
		}

		fieldInfo := &FieldInfo{
			CheckIf:    nil,
			Name:       name,
			Field:      field,
			IsOneOf:    isOneOf,
			InOneOf:    inOneOf,
			TagOptions: validOptions.Tags,
			IsCheckIf:  false,
			Parent:     nil,
			IsListItem: false,
			IsMapKey:   false,
			IsMapValue: false,
		}

		checkIf := validOptions.CheckIf
		if checkIf != nil && checkIf.Field != "" {
			//if checkIf.Field == fieldInfo.Name {
			//	p.exitWithMsg("%s; cannot point to itself in check_if option",
			//		p.buildIdentifierWithName(fieldInfo.Name),
			//	)
			//}

			checkField := &FieldInfo{
				CheckIf:    nil,
				Name:       "",
				Field:      nil,
				IsOneOf:    false,
				InOneOf:    false,
				TagOptions: checkIf.Tags,
				IsCheckIf:  true,
				Parent:     fieldInfo,
				IsListItem: false,
				IsMapKey:   false,
				IsMapValue: false,
			}

			if f1, ok := plainFieldMap[checkIf.Field]; ok {
				checkField.Field = f1
				checkField.IsOneOf = utils.FieldIsOneOf(f1)
				if checkField.IsOneOf {
					checkField.Name = string(f1.Oneof.Desc.Name())
				} else {
					checkField.Name = string(f1.Desc.Name())
				}
			}
			if f2, ok := oneOfFieldMap[checkIf.Field]; ok {
				checkField.Field = f2
				checkField.Name = string(f2.Desc.Name())
				checkField.IsOneOf = true
				checkField.InOneOf = true
			}

			if checkField.Field == nil {
				p.exitWithMsg(
					"%s; not found check_if field: %s",
					p.buildIdentifierWithName(fieldInfo.Name), validOptions.CheckIf.Field,
				)
			}

			fieldInfo.CheckIf = &CheckIfInfo{Field: checkField}
		}

		filedInfos = append(filedInfos, fieldInfo)
	}

	for _, field := range fields {
		processField(field, false)
		if utils.FieldIsOneOf(field) {
			for _, oneOfField := range field.Oneof.Fields {
				processField(oneOfField, true)
			}
		}
	}

	p.filedInfos = filedInfos
}
