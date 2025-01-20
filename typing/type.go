package typing

import (
	"sync"
)

// ----------------------------------------------------------------------------
// Variable

type Variable int

var (
	counter     Variable = 0
	counterLock sync.Mutex
)

func fresh() Variable {
	counterLock.Lock()
	defer counterLock.Unlock()

	v := counter
	counter++
	return v
}

// union combines two slices of Variables into a single slice with no duplicates.
func union(vars1, vars2 []Variable) []Variable {
	m := make(map[Variable]bool)
	for _, v := range append(vars1, vars2...) {
		m[v] = true
	}

	keys := make([]Variable, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}

	return keys
}

// difference returns the set of all things that belong to vars1 but not vars2.
func difference(vars1, vars2 []Variable) []Variable {
	m := make(map[Variable]bool)
	for _, v := range vars1 {
		m[v] = true
	}

	for _, v := range vars2 {
		m[v] = false
	}

	keys := make([]Variable, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}

	return keys
}

// ----------------------------------------------------------------------------
// Type

type Type interface {
	Convert(TyVar, Type) Type
	Variables() []Variable
	typ()
}

type Mono interface {
	Type
	monoType()
}

type Poly interface {
	Type
	polyType()
}

// ----------------------------------------------------------------------------
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
	return difference(vars, s.BoundVars)
}

func (t TyScheme) Convert(TyVar, Type) Type {
	return TyBool{}
}
func (t TyScheme) Variables() []Variable {
	return t.BoundVars
}

func (TyScheme) polyType() {}

func (TyScheme) typ() {}

// ----------------------------------------------------------------------------
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

func NewFreshTyVar() TyVar {
	return TyVar{Variable: fresh()}
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
	return union(absVars, appVars)
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
