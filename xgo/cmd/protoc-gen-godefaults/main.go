package main

import (
	"github.com/yu31/protoc-plugin/xgo/internal/generator"
	"github.com/yu31/protoc-plugin/xgo/internal/plugins/godefaults"
)

func main() {
	generator.Do(godefaults.New())
}
