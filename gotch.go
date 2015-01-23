package main

import (
	"fmt"
	"go/build"
	"go/parser"
	"go/token"
	"path"
	"strconv"
)

func main() {
	files, err := findFiles("../learn-hue-api/main.go")
	fmt.Println("==============")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(files)
	watch(files)
}

func findFiles(filename string) ([]string, error) {
	// fmt.Println(filename)
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, filename, nil, parser.ImportsOnly)
	if err != nil {
		return nil, err
	}
	// fmt.Printf("%+v\n", file)
	// fmt.Println("--------------")
	pkgs := make(map[string]*build.Package)
	for _, val := range file.Imports {
		pkgName, err := strconv.Unquote(val.Path.Value)
		if err != nil {
			return nil, err
		}
		findPackages(&pkgs, pkgName)
		if err != nil {
			return nil, err
		}
	}
	fmt.Println("=====")
	fmt.Println(pkgs)

	var files []string
	for _, pkg := range pkgs {
		// fmt.Println(pkg)
		for _, file := range pkg.GoFiles {
			files = append(files, path.Join(pkg.Dir, file))
		}
		// for _, pkgName := range pkg.Imports {
		// 	files, err = imports(pkgName, files)
		// 	if err != nil {
		// 		return nil, err
		// 	}
		// }
	}
	return files, nil
}

func findPackages(pkgs *map[string]*build.Package, pkgName string) error {
	if (*pkgs)[pkgName] != nil {
		return nil
	}
	pkg, err := build.Import(pkgName, "", 0)
	if err != nil {
		return err
	}
	// fmt.Printf("%+v\n", pkg)
	fmt.Println("--------------")
	fmt.Println("Package:", pkgName)
	fmt.Println("Dir    :", pkg.Dir)
	fmt.Println("Files  :", pkg.GoFiles)
	fmt.Println("Imports:", pkg.Imports)
	(*pkgs)[pkgName] = pkg
	for _, pn := range pkg.Imports {
		findPackages(pkgs, pn)
	}
	return nil
}
