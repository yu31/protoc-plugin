package govalidator

import (
	"github.com/yu31/protoc-plugin/cmd/internal/generator"
	"github.com/yu31/protoc-plugin/cmd/internal/generator/utils"
	"google.golang.org/protobuf/compiler/protogen"
)

const version = "0.0.1"

var (
	validatorPackage = protogen.GoImportPath("github.com/yu31/protoc-plugin/xgo/pkg/protovalidator")
	regexpPackage    = protogen.GoImportPath("regexp")
	stringsPackage   = protogen.GoImportPath("strings")
	utf8Package      = protogen.GoImportPath("unicode/utf8")
	strconvPackage   = protogen.GoImportPath("strconv")
)

type plugin struct {
	g    *protogen.GeneratedFile
	file *protogen.File

	messages []*protogen.Message

	// The message of currently being processed.
	message *protogen.Message

	// The fileds of currently being processed.
	filedInfos []*FieldInfo
}

func New() generator.Plugin {
	return &plugin{}
}

// Name identifies the plugin.
func (p *plugin) Name() string {
	return "validator"
}

// Version identifies the plugin version.
func (p *plugin) Version() string {
	return version
}

func (p *plugin) Init(file *protogen.File) bool {
	if len(file.Messages) == 0 {
		return false
	}
	p.file = file
	p.messages = utils.LoadValidMessages(file.Messages)
	return true
}

// Generate produces the code generated by the plugin for this file,
// except for the imports, by calling the generator's methods P, In, and Out.
func (p *plugin) Generate(g *protogen.GeneratedFile) {
	p.g = g
	for _, msg := range p.messages {
		p.generateMessage(msg)
	}
}

func (p *plugin) generateMessage(msg *protogen.Message) {
	p.message = msg

	p.loadFieldList()

	// TODO: check cycle message and skip.

	p.generateCode()
}

//func (p *plugin) processMessageTags(msg *protogen.Message)  {
//	for _, field := range msg.Fields {
//		switch field.Desc.Kind() {
//		case protoreflect.MessageKind:
//			println(field.Message.GoIdent.GoName, field.Desc.TextName())
//			if field.Message.GoIdent.GoName != p.message.GoIdent.GoName {
//				p.processMessageTags(field.Message)
//			}
//		}
//	}
//}