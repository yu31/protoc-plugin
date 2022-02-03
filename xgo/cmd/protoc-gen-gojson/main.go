package main

import (
	"github.com/yu31/protoc-plugin/xgo/internal/generator"
	"github.com/yu31/protoc-plugin/xgo/internal/plugins/gojson"
)

func main() {
	generator.Do(gojson.New())
}
