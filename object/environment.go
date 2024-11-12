package object

import "go_type_inference/ast"

func NewEnvironment() *Environment {
	s := make(map[ast.Identifier]Value)
	return &Environment{Store: s}
}

func NewTypeEnvironment() *TypeEnvironment {
	s := make(map[ast.Identifier]InferredObject)
	return &TypeEnvironment{Store: s}
}

type Environment struct {
	Store map[ast.Identifier]Value
}

type TypeEnvironment struct {
	Store map[ast.Identifier]InferredObject
}

func (e *Environment) Get(k ast.Identifier) (Value, bool) {
	v, ok := e.Store[k]
	return v, ok
}

func (e *Environment) Set(k ast.Identifier, v Value) Value {
	e.Store[k] = v
	return v
}

func (e *TypeEnvironment) Get(k ast.Identifier) (InferredObject, bool) {
	obj, ok := e.Store[k]
	return obj, ok
}

func (e *TypeEnvironment) Set(k ast.Identifier, v InferredObject) InferredObject {
	e.Store[k] = v
	return v
}
