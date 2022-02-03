package main

import (
	"github.com/yu31/protoc-plugin/cmd/internal/generator"
	"github.com/yu31/protoc-plugin/cmd/internal/plugins/govalidator"
)

func main() {
	generator.Do(govalidator.New())
}
