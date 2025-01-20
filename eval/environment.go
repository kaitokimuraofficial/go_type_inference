package eval

import "go_type_inference/ast"

type Environment struct {
	Store map[ast.Identifier]Value
}

func (e Environment) Get(k ast.Identifier) (Value, bool) {
	v, ok := e.Store[k]
	return v, ok
}

func (e Environment) Set(k ast.Identifier, v Value) Value {
	e.Store[k] = v
	return v
}
