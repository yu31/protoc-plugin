package main

import (
	"github.com/yu31/protoc-plugin/xgo/internal/generator"
	"github.com/yu31/protoc-plugin/xgo/internal/plugins/gosql"
)

func main() {
	generator.Do(gosql.New())
}
