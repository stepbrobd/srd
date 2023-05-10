package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"

	"github.com/charmbracelet/log"
)

type State struct {
	Goroutines map[string]bool
	Memory     map[string]int
	Channels   map[string][]bool
}

type Visitor struct {
	State *State
}

func (v *Visitor) Visit(node ast.Node) ast.Visitor {
	switch n := node.(type) {
	case *ast.GoStmt:
		// Goroutine Creation Rule
		v.State.Goroutines[fmt.Sprint(n.Call.Fun)] = true
	case *ast.IncDecStmt:
		// Memory Modification Rule
		// (Assuming this occurs inside a goroutine for simplicity)
		for goroutine := range v.State.Goroutines {
			v.State.Memory[goroutine]++
		}
	case *ast.SendStmt:
		// Channel Send Rule
		v.State.Channels[n.Chan.(*ast.Ident).Name] = append(v.State.Channels[n.Chan.(*ast.Ident).Name], true)
	case *ast.UnaryExpr:
		if n.Op == token.ARROW {
			// Channel Receive Rule
			v.State.Channels[n.X.(*ast.Ident).Name] = v.State.Channels[n.X.(*ast.Ident).Name][1:]
		}
	case *ast.CallExpr:
		// Print Rule (Assuming all calls are print statements for simplicity)
		log.Info(n.Fun)
	}
	return v
}

func main() {
	filename := os.Args[len(os.Args)-1]

	// Create a FileSet to work with
	fset := token.NewFileSet()

	// Parse the file and create an AST
	file, err := parser.ParseFile(fset, filename, nil, parser.ParseComments)
	if err != nil {
		panic(err)
	}

	// Initialize the state and visitor
	state := &State{
		Goroutines: make(map[string]bool),
		Memory:     make(map[string]int),
		Channels:   make(map[string][]bool),
	}
	visitor := &Visitor{State: state}

	// Use the visitor to traverse the AST
	ast.Walk(visitor, file)

	// Print out the state after traversal
	log.Info("Goroutines:", visitor.State.Goroutines)
	log.Info("Memory:", visitor.State.Memory)
	log.Info("Channels:", visitor.State.Channels)

	// Check for potential data races
	dataRace := false
	for location, accesses := range visitor.State.Memory {
		if accesses > 1 {
			log.Info("Potential data race at memory location:", location)
			dataRace = true
		}
	}
	if !dataRace {
		log.Info("No potential data races detected.")
	}
}
