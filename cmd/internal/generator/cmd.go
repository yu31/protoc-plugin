package generator

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/types/pluginpb"
)

func Do(plugin Plugin) {
	if plugin.Name() == "" {
		panic("plugin name not sets.")
	}
	if plugin.Version() == "" {
		panic("plugin version not sets.")
	}

	if len(os.Args) == 2 && os.Args[1] == "version" {
		_, _ = fmt.Fprintf(os.Stdout, "%v %v\n", filepath.Base(os.Args[0]), plugin.Version())
		os.Exit(0)
	}

	var flags flag.FlagSet

	var options = protogen.Options{
		ParamFunc: flags.Set,
		ImportRewriteFunc: func(importPath protogen.GoImportPath) protogen.GoImportPath {
			return importPath
		},
	}

	options.Run(func(pp *protogen.Plugin) error {
		for _, file := range pp.Files {
			if !file.Generate {
				continue
			}
			if !plugin.Init(file) {
				//println("proto plugin:", "go"+plugin.Name(), "- ignore file:", file.Desc.Path())
				continue
			}

			_ = generateFile(plugin, pp, file)
		}
		pp.SupportedFeatures = uint64(pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL)
		return nil
	})
}
