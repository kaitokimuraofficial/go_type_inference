package eval

import "go_type_inference/ast"

type Environment struct {
	Store map[ast.Ident]Value
}

func (e Environment) Get(k ast.Ident) (Value, bool) {
	v, ok := e.Store[k]
	return v, ok
}

func (e Environment) Set(k ast.Ident, v Value) Value {
	e.Store[k] = v
	return v
}
