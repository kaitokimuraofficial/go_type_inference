package main

import (
	"fmt"
	"go_type_inference/eval"
	"go_type_inference/object"
	"go_type_inference/parser"
)

func main() {
	input := `if true then (if false then 3 else 2) else 3`

	// Parse the input string
	p := parser.Parse(input)

	// Make Environment
	env := object.NewEnvironment()

	// Evaluate AST made from parser.Parse
	e := eval.Eval(p, env)

	fmt.Println(e.Inspect())
}
