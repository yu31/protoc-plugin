package govalidator

import (
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"

	"github.com/yu31/protoc-plugin/xgo/pb/pbvalidator"
	"github.com/yu31/protoc-plugin/xgo/pkg/protovalidator"
)

func (p *plugin) generateVariableForField(fieldInfo *FieldInfo) {
	if fieldInfo.TagOptions == nil {
		return
	}
	switch {
	case fieldInfo.Field.Desc.IsMap():
		p.generateVariableForMap(fieldInfo)
	case fieldInfo.Field.Desc.IsList():
		p.generateVariableForList(fieldInfo)
	default:
		p.generateVariableForBasic(fieldInfo)
	}
}

func (p *plugin) generateVariableForMap(fieldInfo *FieldInfo) {
	options := p.loadMapTags(fieldInfo.Field, fieldInfo.TagOptions)
	if options == nil {
		return
	}

	if options.Key != nil && options.Key.Kind != nil {
		subField := fieldInfo.Field.Message.Fields[0]
		subFieldInfo := &FieldInfo{
			CheckIf:    fieldInfo.CheckIf,
			Name:       fieldInfo.Name,
			Field:      subField,
			IsOneOf:    false,
			InOneOf:    false,
			TagOptions: options.Key,
			IsCheckIf:  fieldInfo.IsCheckIf,
			Parent:     fieldInfo.Parent,
		}
		p.generateVariableForBasic(subFieldInfo)
	}
	if options.Value != nil && options.Value.Kind != nil {
		subField := fieldInfo.Field.Message.Fields[1]
		subFieldInfo := &FieldInfo{
			CheckIf:    fieldInfo.CheckIf,
			Name:       fieldInfo.Name,
			Field:      subField,
			IsOneOf:    false,
			InOneOf:    false,
			TagOptions: options.Value,
			IsCheckIf:  fieldInfo.IsCheckIf,
			Parent:     fieldInfo.Parent,
		}
		p.generateVariableForBasic(subFieldInfo)
	}
}

func (p *plugin) generateVariableForList(fieldInfo *FieldInfo) {
	options := p.loadRepeatedTags(fieldInfo.Field, fieldInfo.TagOptions)
	if options == nil {
		return
	}
	if options.Item != nil && options.Item.Kind != nil {
		subFieldInfo := &FieldInfo{
			CheckIf:    fieldInfo.CheckIf,
			Name:       fieldInfo.Name,
			Field:      fieldInfo.Field,
			IsOneOf:    false,
			InOneOf:    false,
			TagOptions: options.Item,
			IsCheckIf:  fieldInfo.IsCheckIf,
			Parent:     fieldInfo.Parent,
		}
		p.generateVariableForBasic(subFieldInfo)
	}
}

func (p *plugin) generateVariableForBasic(fieldInfo *FieldInfo) {
	if fieldInfo.TagOptions == nil {
		return
	}
	switch v := fieldInfo.TagOptions.Kind.(type) {
	case *pbvalidator.TagOptions_Float:
		if len(v.Float.In) != 0 {
			varName := p.buildVariableNameForTagIn(fieldInfo)
			p.encodeInAndNotInVariableValue(fieldInfo, protovalidator.TagFloatIn, varName, v.Float.In)
		}
		if len(v.Float.NotIn) != 0 {
			varName := p.buildVariableNameForTagNotIn(fieldInfo)
			p.encodeInAndNotInVariableValue(fieldInfo, protovalidator.TagFloatNotIn, varName, v.Float.NotIn)
		}
	case *pbvalidator.TagOptions_Int:
		if len(v.Int.In) != 0 {
			varName := p.buildVariableNameForTagIn(fieldInfo)
			p.encodeInAndNotInVariableValue(fieldInfo, protovalidator.TagIntIn, varName, v.Int.In)
		}
		if len(v.Int.NotIn) != 0 {
			varName := p.buildVariableNameForTagNotIn(fieldInfo)
			p.encodeInAndNotInVariableValue(fieldInfo, protovalidator.TagIntNotIn, varName, v.Int.NotIn)
		}
	case *pbvalidator.TagOptions_Uint:
		if len(v.Uint.In) != 0 {
			varName := p.buildVariableNameForTagIn(fieldInfo)
			p.encodeInAndNotInVariableValue(fieldInfo, protovalidator.TagUintIn, varName, v.Uint.In)
		}
		if len(v.Uint.NotIn) != 0 {
			varName := p.buildVariableNameForTagNotIn(fieldInfo)
			p.encodeInAndNotInVariableValue(fieldInfo, protovalidator.TagUintNotIn, varName, v.Uint.NotIn)
		}
	case *pbvalidator.TagOptions_String_:
		if len(v.String_.In) != 0 {
			varName := p.buildVariableNameForTagIn(fieldInfo)
			p.encodeInAndNotInVariableValue(fieldInfo, protovalidator.TagStringIn, varName, v.String_.In)

		}
		if len(v.String_.NotIn) != 0 {
			varName := p.buildVariableNameForTagNotIn(fieldInfo)
			p.encodeInAndNotInVariableValue(fieldInfo, protovalidator.TagStringNotIn, varName, v.String_.NotIn)
		}

		if v.String_.Regex != nil && *v.String_.Regex != "" {
			// Check the regex expression.
			_, err := regexp.Compile(*v.String_.Regex)
			if err != nil {
				p.exitWithMsg(
					"%s, option: <regex>; cannot compile regex expr, error: %v",
					p.buildIdentifierWithName(fieldInfo.Name), err,
				)
			}
			p.g.P("var ",
				p.buildVariableNameForTagRegex(fieldInfo),
				" = ",
				regexpPackage.Ident("MustCompile"),
				"(`", *v.String_.Regex, "`)")
		}
	case *pbvalidator.TagOptions_Enum:
		if len(v.Enum.In) != 0 {
			varName := p.buildVariableNameForTagIn(fieldInfo)
			p.encodeInAndNotInVariableValue(fieldInfo, protovalidator.TagEnumIn, varName, v.Enum.In)
		}
		if len(v.Enum.NotIn) != 0 {
			varName := p.buildVariableNameForTagNotIn(fieldInfo)
			p.encodeInAndNotInVariableValue(fieldInfo, protovalidator.TagEnumNotIn, varName, v.Enum.NotIn)
		}
		if v.Enum.InEnums != nil && *v.Enum.InEnums {
			varName := p.buildVariableNameForTagInEnums(fieldInfo)

			var s strings.Builder
			s.WriteString("map[")
			s.WriteString(p.fieldToGoType(fieldInfo.Field))
			s.WriteString("]bool {")

			l := len(fieldInfo.Field.Enum.Values)
			for i, vv := range fieldInfo.Field.Enum.Values {
				s.WriteString(strconv.FormatInt(int64(vv.Desc.Number()), 10))
				s.WriteString(": true")
				if i != l-1 {
					s.WriteString(",")
				}
			}

			s.WriteString("}")
			p.g.P("var ", varName, " = ", s.String())
		}
	}
}

// encode variable value for option In Or NotIn
func (p *plugin) encodeInAndNotInVariableValue(fieldInfo *FieldInfo, option string, varName string, values interface{}) {
	var validEnums map[int32]bool

	if fieldInfo.Field.Enum != nil {
		validEnums = make(map[int32]bool)
		for _, vv := range fieldInfo.Field.Enum.Values {
			validEnums[int32(vv.Desc.Number())] = true
		}
	}

	cacheKey := make(map[interface{}]struct{})

	var s strings.Builder

	s.WriteString("map[")
	s.WriteString(p.fieldToGoType(fieldInfo.Field))
	s.WriteString("]bool {")

	valueOf := reflect.ValueOf(values)
	if valueOf.Kind() != reflect.Slice {
		panic("invalid values")
	}

	valueLen := valueOf.Len()
	for i := 0; i < valueLen; i++ {
		var key interface{}

		item := valueOf.Index(i)
		switch item.Kind() {
		case reflect.Float64:
			key = item.Float()
			s.WriteString(strconv.FormatFloat(item.Float(), 'f', -1, 64))
		case reflect.Int64:
			key = item.Int()
			s.WriteString(strconv.FormatInt(item.Int(), 10))
		case reflect.Uint64:
			key = item.Uint()
			s.WriteString(strconv.FormatUint(item.Uint(), 10))
		case reflect.String:
			key = item.String()
			s.WriteString(strconv.Quote(item.String()))
		case reflect.Int32:
			if fieldInfo.Field.Enum != nil {
				if _, ok := validEnums[int32(item.Int())]; !ok {
					p.exitWithMsg(
						"%s: [%d] not a valid enum number",
						p.buildIdentifierWithName(fieldInfo.Name), item.Int(),
					)
				}
			}
			key = item.Int()
			s.WriteString(strconv.FormatInt(item.Int(), 10))
		default:
			panic(fmt.Errorf("unsupported reflect kind: %s", item.Kind().String()))
		}

		if _, ok := cacheKey[key]; ok {
			p.exitWithMsg(
				"%s, option: <%s>; found duplicated item in %v",
				p.buildIdentifierWithName(fieldInfo.Name), option, values,
			)
		}
		cacheKey[key] = struct{}{}

		s.WriteString(": true")
		if i != valueLen-1 {
			s.WriteString(",")
		}
	}

	s.WriteString("}")
	p.g.P("var ", varName, " = ", s.String())
}
