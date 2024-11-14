package typing

import "go_type_inference/ast"

type Environment struct {
	Store map[ast.Identifier]Type
}

func (e *Environment) Get(k ast.Identifier) (Type, bool) {
	obj, ok := e.Store[k]
	return obj, ok
}

func (e *Environment) Set(k ast.Identifier, v Type) Type {
	e.Store[k] = v
	return v
}
