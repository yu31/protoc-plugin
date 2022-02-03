package gojson

import (
	"github.com/yu31/protoc-plugin/xgo/pb/pbjson"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
)

func (p *plugin) loadFileOptions(file *protogen.File) *pbjson.SerializeOptions {
	i := proto.GetExtension(file.Desc.Options(), pbjson.E_File)
	fileOptions := i.(*pbjson.SerializeOptions)
	if fileOptions == nil {
		fileOptions = &pbjson.SerializeOptions{}
	}
	return fileOptions
}

// The SerializeOptions priority from low to high is: file_options -> msg_options
func (p *plugin) loadMessageOptions(msg *protogen.Message) *pbjson.SerializeOptions {
	i := proto.GetExtension(msg.Desc.Options(), pbjson.E_Message)
	msgOptions := i.(*pbjson.SerializeOptions)
	if msgOptions == nil {
		msgOptions = &pbjson.SerializeOptions{}
	}

	fileOptions := p.fileOptions

	if msgOptions.NameStyle == nil {
		msgOptions.NameStyle = fileOptions.NameStyle
	}
	if msgOptions.Ignore == nil {
		msgOptions.Ignore = fileOptions.Ignore
	}
	if msgOptions.Omitempty == nil {
		msgOptions.Omitempty = fileOptions.Omitempty
	}
	if msgOptions.UseEnumString == nil {
		msgOptions.UseEnumString = fileOptions.UseEnumString
	}
	if msgOptions.HideOneofKey == nil {
		msgOptions.HideOneofKey = fileOptions.HideOneofKey
	}
	if msgOptions.DisallowUnknownFields == nil {
		msgOptions.DisallowUnknownFields = fileOptions.DisallowUnknownFields
	}

	// Set default value for message options.
	if msgOptions.NameStyle == nil {
		style := pbjson.NameStyle_TextName
		msgOptions.NameStyle = &style
	}
	if msgOptions.DisallowUnknownFields == nil {
		ok := false
		msgOptions.DisallowUnknownFields = &ok
	}

	return msgOptions
}

func (p *plugin) loadOneOfOptions(oneof *protogen.Oneof) *pbjson.OneofOptions {
	msgOptions := p.msgOptions
	i := proto.GetExtension(oneof.Desc.Options(), pbjson.E_Oneof)
	oneOfOptions := i.(*pbjson.OneofOptions)
	if oneOfOptions == nil {
		oneOfOptions = &pbjson.OneofOptions{}
	}
	if oneOfOptions.Ignore == nil {
		oneOfOptions.Ignore = msgOptions.Ignore
	}
	if oneOfOptions.Omitempty == nil {
		oneOfOptions.Omitempty = msgOptions.Omitempty
	}
	if oneOfOptions.HideOneofKey == nil {
		oneOfOptions.HideOneofKey = msgOptions.HideOneofKey
	}

	// Set default value for oneof options.
	ok1 := false
	if oneOfOptions.Ignore == nil {
		oneOfOptions.Ignore = &ok1
	}
	if oneOfOptions.Omitempty == nil {
		oneOfOptions.Omitempty = &ok1
	}
	if oneOfOptions.HideOneofKey == nil {
		oneOfOptions.HideOneofKey = &ok1
	}

	// Ignore field if json == "-"
	if oneOfOptions.Json != nil && *oneOfOptions.Json == "-" {
		ok2 := true
		oneOfOptions.Ignore = &ok2
	}
	return oneOfOptions
}

func (p *plugin) loadEnumOptions(enum *protogen.Enum) *pbjson.EnumOptions {
	msgOptions := p.msgOptions
	i := proto.GetExtension(enum.Desc.Options(), pbjson.E_Enum)
	enumOptions := i.(*pbjson.EnumOptions)
	if enumOptions == nil {
		enumOptions = &pbjson.EnumOptions{}
	}
	if enumOptions.UseEnumString == nil {
		enumOptions.UseEnumString = msgOptions.UseEnumString
	}

	return enumOptions
}

func (p *plugin) loadFieldOptions(field *protogen.Field) *pbjson.FieldOptions {
	msgOptions := p.msgOptions
	i := proto.GetExtension(field.Desc.Options(), pbjson.E_Field)
	fieldOptions := i.(*pbjson.FieldOptions)
	if fieldOptions == nil {
		fieldOptions = &pbjson.FieldOptions{}
	}
	if fieldOptions.Omitempty == nil {
		fieldOptions.Omitempty = msgOptions.Omitempty
	}
	if fieldOptions.UseEnumString == nil {
		fieldOptions.UseEnumString = msgOptions.UseEnumString
	}
	if fieldOptions.Ignore == nil {
		fieldOptions.Ignore = msgOptions.Ignore
	}

	if field.Enum != nil && fieldOptions.UseEnumString == nil {
		enumOptions := p.loadEnumOptions(field.Enum)
		fieldOptions.UseEnumString = enumOptions.UseEnumString
	}

	// Set default value for oneof options.
	ok1 := false
	if fieldOptions.Omitempty == nil {
		fieldOptions.Omitempty = &ok1
	}
	if fieldOptions.Ignore == nil {
		fieldOptions.Ignore = &ok1
	}
	if fieldOptions.UseEnumString == nil {
		fieldOptions.UseEnumString = &ok1
	}

	// Ignore field if json == "-"
	if fieldOptions.Json != nil && *fieldOptions.Json == "-" {
		ok2 := true
		fieldOptions.Ignore = &ok2
	}
	return fieldOptions
}
