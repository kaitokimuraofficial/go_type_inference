package typing

import "go_type_inference/ast"

type Environment struct {
	Store map[ast.Identifier]Scheme
}

func (e *Environment) Get(k ast.Identifier) (Scheme, bool) {
	obj, ok := e.Store[k]
	return obj, ok
}

func (e *Environment) Set(k ast.Identifier, v Scheme) Scheme {
	e.Store[k] = v
	return v
}
