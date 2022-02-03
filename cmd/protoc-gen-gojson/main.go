package main

import (
	"github.com/yu31/protoc-plugin/cmd/internal/generator"
	"github.com/yu31/protoc-plugin/cmd/internal/plugins/gojson"
)

func main() {
	generator.Do(gojson.New())
}
