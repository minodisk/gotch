package main

import (
	"fmt"
	"go/build"
	"go/parser"
	"go/token"
	"regexp"
)

func main() {
	imports("../learn-hue-api/main.go")
}

func imports(path string) {
	fmt.Println(path)
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, path, nil, parser.ImportsOnly)
	if err != nil {
		fmt.Println(err)
	}
	for _, val := range file.Imports {
		re := regexp.MustCompile("^\"(.*)\"$")
		path := re.ReplaceAllString(val.Path.Value, "$1")
		fmt.Println(path)
		pkg, err := build.Import(path, "../", 0)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(pkg.Imports)
	}
}
