package main

import (
	"github.com/yu31/protoc-plugin/cmd/internal/generator"
	"github.com/yu31/protoc-plugin/cmd/internal/plugins/godefaults"
)

func main() {
	generator.Do(godefaults.New())
}
