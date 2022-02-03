package gojson

import (
	"fmt"
	"os"

	"google.golang.org/protobuf/compiler/protogen"
)

func (p *plugin) checkJSONKey() {
	msg := p.message
	fields := p.fields

	cacheFields := make(map[string]*protogen.Field)

	dupOneOfs := make([]map[string][]string, 0)
	emptyOneOfs := make([][]string, 0)

	emptyFields := make([]string, 0)
	dupFields := make(map[string][]string)

	checkFieldDup := func(field *protogen.Field) {
		options := p.loadFieldOptions(field)
		if *options.Ignore {
			return
		}
		jsonKey := p.getFieldKey(options, field)
		if jsonKey == "" {
			emptyFields = append(emptyFields, string(field.Desc.Name()))
			return
		}
		if _, ok := cacheFields[jsonKey]; ok {
			dupFields[jsonKey] = append(dupFields[jsonKey], string(field.Desc.Name()))
			return
		}
		cacheFields[jsonKey] = field
	}

LOOP:
	for _, field := range fields {
		// Check general field.
		if field.Oneof == nil || field.Oneof.Desc.IsSynthetic() {
			checkFieldDup(field)
			continue LOOP
		}

		oneOfOptions := p.loadOneOfOptions(field.Oneof)
		if *oneOfOptions.Ignore {
			continue LOOP
		}

		if *oneOfOptions.HideOneofKey {
			for _, f := range field.Oneof.Fields {
				checkFieldDup(f)
			}
			continue LOOP
		}

		// oneOf key not hide in json. check it.
		oneOfKey := p.getOneOfKey(oneOfOptions, field.Oneof)
		if oneOfKey == "" {
			emptyFields = append(emptyFields, field.Oneof.GoName)
		} else {
			if _, ok := cacheFields[oneOfKey]; ok {
				// find duplicate key
				dupFields[oneOfKey] = append(dupFields[oneOfKey], field.Oneof.GoName)
			} else {
				cacheFields[oneOfKey] = field
			}
		}

		// Check oneof's fields
		cacheOneOf := make(map[string]*protogen.Field)
		dupOneOf := make(map[string][]string)
		emptyOneOf := make([]string, 0)

	ONEOF:
		for _, f := range field.Oneof.Fields {
			fieldOptions := p.loadFieldOptions(f)
			if *fieldOptions.Ignore {
				continue ONEOF
			}
			jsonKey := p.getFieldKey(fieldOptions, f)
			if jsonKey == "" {
				emptyOneOf = append(emptyOneOf, string(f.Desc.Name()))
			} else {
				if _, ok := cacheOneOf[jsonKey]; ok {
					dupOneOf[jsonKey] = append(dupOneOf[jsonKey], string(f.Desc.Name()))
					continue ONEOF
				}
				cacheOneOf[jsonKey] = f
			}
		}

		for jsonKey := range dupOneOf {
			if x, ok := cacheOneOf[jsonKey]; ok {
				dupOneOf[jsonKey] = append(dupOneOf[jsonKey], string(x.Desc.Name()))
			}
		}
		if len(dupOneOf) != 0 {
			dupOneOfs = append(dupOneOfs, dupOneOf)
		}
		if len(emptyOneOf) != 0 {
			emptyOneOfs = append(emptyOneOfs, emptyOneOf)
		}
	}

	if len(dupFields) == 0 && len(dupOneOfs) == 0 && len(emptyFields) == 0 && len(emptyOneOfs) == 0 {
		return
	}

	if len(emptyFields) != 0 {
		println(fmt.Sprintf(
			"gojson: <file(%s) message(%s)>: (general) the json key is empty in fields %v",
			string(p.file.GoImportPath), msg.GoIdent.GoName, emptyFields,
		))
	}

	for jsonKey, value := range dupFields {
		if x, ok := cacheFields[jsonKey]; ok {
			value = append(value, string(x.Desc.Name()))
		}
		println(fmt.Sprintf(
			"gojson: <file(%s) message(%s)>: (general) Found duplicate json key [%s] both in fields %v",
			string(p.file.GoImportPath), msg.GoIdent.GoName, jsonKey, value,
		))
	}

	for _, emptyOneOf := range emptyOneOfs {
		println(fmt.Sprintf(
			"gojson: <file(%s) message(%s)>: (oneof) the json key is empty in fields %v",
			string(p.file.GoImportPath), msg.GoIdent.GoName, emptyOneOf,
		))
	}

	for _, dupOneOf := range dupOneOfs {
		for jsonKey, value := range dupOneOf {
			println(fmt.Sprintf(
				"gojson: <file(%s) message(%s)>: (oneof) Found duplicate json key [%s] both in fields %v",
				string(p.file.GoImportPath), msg.GoIdent.GoName, jsonKey, value,
			))
		}
	}
	os.Exit(1)
}
