package main

import (
	"github.com/yu31/protoc-plugin/cmd/internal/generator"
	"github.com/yu31/protoc-plugin/cmd/internal/plugins/gosql"
)

func main() {
	generator.Do(gosql.New())
}
