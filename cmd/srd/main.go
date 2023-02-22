package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"time"

	"github.com/charmbracelet/log"
)

func main() {
	for i := 1; i <= 100; i++ {
		log.Info(fmt.Sprintf("Running %d/100...", i))
		time.Sleep(10 * time.Millisecond)
	}

	//Create a FileSet to work with
	fset := token.NewFileSet()
	//Parse the file and create an AST
	file, err := parser.ParseFile(fset, "./cmd/srd/main.go", nil, parser.ParseComments)
	if err != nil {
		panic(err)
	}
	ast.Inspect(file, func(n ast.Node) bool {
		// Find Function Call Statements
		funcCall, ok := n.(*ast.CallExpr)
		if ok {
			fmt.Println(funcCall.Fun)
		}
		return true
	})
}
