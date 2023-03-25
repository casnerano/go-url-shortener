// Package analyzer search call os.Exit in main packages and report position.
// Implements analysis.Analyzer type interface for multi-check.
package exitcheck

import (
	"fmt"
	"go/ast"

	"golang.org/x/tools/go/analysis"
)

// Analyzer describes an analysis
var Analyzer = &analysis.Analyzer{
	Name: "exitcheck",
	Doc:  "check direct call to os.Exit in the main function of the main package",
	Run:  run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	exitDetectInCallExpr := func(x *ast.CallExpr) {
		s, ok := x.Fun.(*ast.SelectorExpr)
		if !ok {
			return
		}

		ident, ok := s.X.(*ast.Ident)
		if !ok {
			return
		}

		if ident.Name == "os" && s.Sel.Name == "Exit" {
			pass.Reportf(ident.NamePos, "direct call to os.Exit in main package main function")
		}
	}

	for _, file := range pass.Files {
		fmt.Println(file.Name.Name)
		if file.Name.Name != "main" {
			continue
		}

		for _, decl := range file.Decls {
			if fnDecl, ok := decl.(*ast.FuncDecl); ok && fnDecl.Name.Name == "main" {
				ast.Inspect(fnDecl.Body, func(node ast.Node) bool {
					if callExpr, callExprOk := node.(*ast.CallExpr); callExprOk {
						exitDetectInCallExpr(callExpr)
					}
					return true
				})
			}
		}
	}

	return nil, nil
}
