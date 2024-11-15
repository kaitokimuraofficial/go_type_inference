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

// ContainsIn checks if a given list of Variable contains a specific Variable
func ContainsIn(vars []Variable, target Variable) bool {
	for _, v := range vars {
		if v == target {
			return true
		}
	}
	return false
}

// ----------------------------------------------------------------------------
// Type

type Type interface {
	Convert(TyIdent, Type) Type
	Variables() []Variable
}

type (
	TyInt struct{}

	TyBool struct{}

	TyFun struct {
		Abs Type
		App Type
	}

	TyIdent struct {
		Variable Variable
	}
)

func NewFreshTyIdent() *TyIdent {
	return &TyIdent{Variable: fresh()}
}

func (t *TyInt) Convert(TyIdent, Type) Type {
	return t
}
func (t *TyBool) Convert(TyIdent, Type) Type {
	return t
}
func (t *TyFun) Convert(ident TyIdent, to Type) Type {
	abs := t.Abs.Convert(ident, to)
	app := t.App.Convert(ident, to)
	return &TyFun{Abs: abs, App: app}
}
func (t *TyIdent) Convert(ident TyIdent, to Type) Type {
	if t.Variable == ident.Variable {
		return to
	}
	return t
}

func (*TyInt) Variables() []Variable {
	return []Variable{}
}
func (*TyBool) Variables() []Variable {
	return []Variable{}
}
func (t *TyFun) Variables() []Variable {
	absVars := t.Abs.Variables()
	appVars := t.App.Variables()
	return union(absVars, appVars)
}
func (t *TyIdent) Variables() []Variable {
	return []Variable{t.Variable}
}
