package object

func NewEnvironment() *Environment {
	s := make(map[string]Object)
	return &Environment{Store: s}
}

func NewTypeEnvironment() *TypeEnvironment {
	s := make(map[string]InferObject)
	return &TypeEnvironment{Store: s}
}

type Environment struct {
	Store map[string]Object
}

type TypeEnvironment struct {
	Store map[string]InferObject
}

func (e *Environment) Get(name string) (Object, bool) {
	obj, ok := e.Store[name]
	return obj, ok
}

func (e *Environment) Set(name string, val Object) Object {
	e.Store[name] = val
	return val
}

func (e *TypeEnvironment) Get(name string) (InferObject, bool) {
	obj, ok := e.Store[name]
	return obj, ok
}

func (e *TypeEnvironment) Set(name string, val InferObject) InferObject {
	e.Store[name] = val
	return val
}
