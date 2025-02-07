package main

import (
	"fmt"
	"go_type_inference/ast"
	"go_type_inference/eval"
	"go_type_inference/parser"
)

func main() {
	input := `if true then (if false then 3 else 2) else 3`

	// Parse the input string
	p := parser.Parse(input)

	// Make Environment
	env := eval.Environment{Store: make(map[ast.Ident]eval.Value)}

	// Evaluate AST made from parser.Parse
	e := eval.Eval(p, env)

	switch e := e.(type) {
	case *eval.Integer:
		fmt.Printf("%d", e.Value)
	case *eval.Boolean:
		fmt.Printf("%t", e.Value)
	case *eval.Function:
		fmt.Printf("function")
	}
}
