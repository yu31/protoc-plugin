package godefaults

import (
	"fmt"
	"os"
	"runtime"

	"github.com/yu31/protoc-plugin/xgo/internal/generator"
	"github.com/yu31/protoc-plugin/xgo/internal/utils"
	"google.golang.org/protobuf/compiler/protogen"
)

const version = "0.0.1"

type plugin struct {
	g    *protogen.GeneratedFile
	file *protogen.File

	// The valid message lists.
	messages []*protogen.Message

	// The message of currently being processed.
	message *protogen.Message

	// The fileds of currently being processed.
	fields []*protogen.Field
}

func New() generator.Plugin {
	return &plugin{}
}

// Name identifies the plugin.
func (p *plugin) Name() string {
	return "defaults"
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
	defer func() {
		if r := recover(); r != nil {
			println(fmt.Sprintf(
				"panic on -> file: %s, message: %s",
				p.file.Desc.FullName(), msg.Desc.Name(),
			))

			println(fmt.Sprintf("recover: %v", r))
			buf := make([]byte, 4096)
			_ = runtime.Stack(buf, true)
			println(string(buf))

			os.Exit(1)
		}
	}()

	p.message = msg
	p.fields = utils.LoadFieldList(msg)

	p.generateCode()
}