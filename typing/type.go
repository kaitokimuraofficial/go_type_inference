package typing

import (
	"slices"
	"sync"
)

// Type

type (
	Type interface {
		Convert(TyVar, Type) Type
		Variables() []Variable
		typ()
	}

	Mono interface {
		Type
		monoType()
	}

	Poly interface {
		Type
		polyType()
	}
)

// Polymorphic Type

type TyScheme struct {
	BoundVars []Variable
	Type      Type
}

func NewScheme(typ Type) TyScheme {
	return TyScheme{BoundVars: []Variable{}, Type: typ}
}

func FreeVariables(s TyScheme) []Variable {
	vars := s.Type.Variables()
	return slices.DeleteFunc(vars, func(v Variable) bool {
		return slices.Contains(s.BoundVars, v)
	})
}

func (t TyScheme) Convert(TyVar, Type) Type {
	return TyBool{}
}
func (t TyScheme) Variables() []Variable {
	return t.BoundVars
}

func (TyScheme) polyType() {}

func (TyScheme) typ() {}

// Monomorphic Type

type (
	TyInt struct{}

	TyBool struct{}

	TyFun struct {
		Abs Type
		App Type
	}

	TyVar struct {
		Variable Variable
	}
)

// Variable

type Variable int

var (
	counter     Variable = 0
	counterLock sync.Mutex
)

func FreshTyVar() TyVar {
	counterLock.Lock()
	defer counterLock.Unlock()

	v := counter
	counter++
	return TyVar{Variable: v}
}

func (t TyInt) Convert(TyVar, Type) Type {
	return t
}
func (t TyBool) Convert(TyVar, Type) Type {
	return t
}
func (t TyFun) Convert(ident TyVar, to Type) Type {
	abs := t.Abs.Convert(ident, to)
	app := t.App.Convert(ident, to)
	return TyFun{Abs: abs, App: app}
}
func (t TyVar) Convert(ident TyVar, to Type) Type {
	if t.Variable == ident.Variable {
		return to
	}
	return t
}

func (TyInt) Variables() []Variable {
	return []Variable{}
}
func (TyBool) Variables() []Variable {
	return []Variable{}
}
func (t TyFun) Variables() []Variable {
	absVars := t.Abs.Variables()
	appVars := t.App.Variables()
	return slices.Compact(slices.Concat(absVars, appVars))
}
func (t TyVar) Variables() []Variable {
	return []Variable{t.Variable}
}

func (TyInt) monoType()  {}
func (TyBool) monoType() {}
func (TyFun) monoType()  {}
func (TyVar) monoType()  {}

func (TyInt) typ()  {}
func (TyBool) typ() {}
func (TyFun) typ()  {}
func (TyVar) typ()  {}
