package main

import (
	"github.com/yu31/protoc-plugin/xgo/internal/generator"
	"github.com/yu31/protoc-plugin/xgo/internal/plugins/govalidator"
)

func main() {
	generator.Do(govalidator.New())
}
