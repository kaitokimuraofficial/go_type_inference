package typing

import (
	"go_type_inference/ast"
)

type Environment struct {
	Store map[Variable]Type

	nextKey Variable
}

func (e Environment) Get(k ast.Ident) (Type, bool) {

	obj, ok := e.Store[k]
	return obj, ok
}

func (e Environment) Set(k ast.Ident, t Type) Type {
	e.Store[k] = t
	return t
}
