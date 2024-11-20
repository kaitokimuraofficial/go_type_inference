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
// TyScheme

type TyScheme interface {
	tyScheme()
	Variables() []Variable
}

type Sch interface {
	TyScheme
	sch()
}

type Typ interface {
	TyScheme
	typ()
	replace(TyVar, Typ) Typ
}

// ----------------------------------------------------------------------------
// Typ

type (
	TyInt struct{}

	TyBool struct{}

	TyFun struct {
		Abs Typ
		App Typ
	}

	TyVar struct {
		Variable Variable
	}
)

func FreshTyVar() *TyVar {
	return &TyVar{Variable: fresh()}
}

func (t *TyInt) replace(frm TyVar, to Typ) Typ {
	return t
}
func (t *TyBool) replace(frm TyVar, to Typ) Typ {
	return t
}
func (t *TyFun) replace(frm TyVar, to Typ) Typ {
	abs := t.Abs.replace(frm, to)
	app := t.App.replace(frm, to)
	return &TyFun{Abs: abs, App: app}
}
func (t *TyVar) replace(frm TyVar, to Typ) Typ {
	if frm == *t {
		return to
	}
	return t
}

func (t *TyInt) Variables() []Variable {
	return []Variable{}
}
func (t *TyBool) Variables() []Variable {
	return []Variable{}
}
func (t *TyFun) Variables() []Variable {
	v1 := t.Abs.Variables()
	v2 := t.App.Variables()
	return union(v1, v2)
}
func (t *TyVar) Variables() []Variable {
	return []Variable{t.Variable}
}

func (*TyInt) tyScheme()  {}
func (*TyBool) tyScheme() {}
func (*TyFun) tyScheme()  {}
func (*TyVar) tyScheme()  {}

func (*TyInt) typ()  {}
func (*TyBool) typ() {}
func (*TyFun) typ()  {}
func (*TyVar) typ()  {}

// ----------------------------------------------------------------------------
// Sch

type Scheme struct {
	BTV  []Variable
	Type Typ
}

func (*Scheme) tyScheme() {}

func (*Scheme) sch() {}

func NewScheme(typ Typ) *Scheme {
	return &Scheme{BTV: []Variable{}, Type: typ}
}

// func FreeVariables(s Scheme) []Variable {
// 	vars := s.Type.Variables()
// 	return difference(vars, s.BoundVars)
// }
