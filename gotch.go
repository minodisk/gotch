package main

import (
	"fmt"
	"go/parser"
	"go/token"
)

func main() {
	fset := token.NewFileSet()
	// file, err := parser.ParseFile(fset, "../learn-hue-api/main.go", nil, parser.ImportsOnly)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// for key, val := range file.Imports {
	// 	fmt.Println(key, val)
	// 	fmt.Printf("%#", val)
	// 	fmt.Printf("%#", val.Path.Value)
	// 	imports := make(map[string]*ast.Object)
	// 	ast.Importer
	// }

	path := "../learn-hue-api"
	pkgs, err := parser.ParseDir(fset, path, nil, parser.ImportsOnly)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(len(pkgs))
	for name, pkg := range pkgs {
		fmt.Println(name, pkg)
		fmt.Println(len(pkg.Imports))
		fmt.Println(pkg.Imports)
		for id, p := range pkg.Imports {
			fmt.Println(id, p)
		}
	}
}
