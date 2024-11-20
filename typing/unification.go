package typing

import "log"

// ----------------------------------------------------------------------------
// Constraint

// Constraint represents a set of constraints that the type must satisfy.
type Constraint struct {
	Left  Typ
	Right Typ
}

func Unify(cs []Constraint) []Substitution {
	for i, c := range cs {
		left, right := c.Left, c.Right

		if left == right {
			return Unify(append(cs[:i], cs[i+1:]...))
		}

		lf, lok := left.(*TyFun)
		rf, rok := right.(*TyFun)
		if lok && rok {
			return Unify(append(
				cs[:i],
				append(
					cs[i+1:],
					Constraint{Left: lf.Abs, Right: rf.Abs},
					Constraint{Left: lf.App, Right: rf.App},
				)...,
			))
		}

		li, lok := left.(*TyVar)
		if lok && !ContainsIn(right.Variables(), li.Variable) {
			newCS := append(cs[:i], cs[i+1:]...)
			for i, c := range newCS {
				l := c.Left.replace(*li, right)
				r := c.Right.replace(*li, right)
				newCS[i] = Constraint{Left: l, Right: r}
			}

			substitution := Unify(newCS)
			return append(substitution, Substitution{
				Var:  TyVar{Variable: li.Variable},
				Type: right,
			})
		}

		ri, rok := right.(*TyVar)
		if rok && !ContainsIn(left.Variables(), ri.Variable) {
			newCS := append(cs[:i], cs[i+1:]...)
			for i, c := range newCS {
				l := c.Left.replace(*ri, left)
				r := c.Right.replace(*ri, left)
				newCS[i] = Constraint{Left: l, Right: r}
			}

			substitution := Unify(newCS)
			return append(substitution, Substitution{
				Var:  TyVar{Variable: ri.Variable},
				Type: left,
			})
		}
	}
	return nil
}

// Union combines two slices of Constraints into a single slice with no duplicates.
func Union(lists ...[]Constraint) []Constraint {
	m := make(map[Constraint]bool)
	for _, list := range lists {
		for _, c := range list {
			m[c] = true
		}
	}

	combined := []Constraint{}
	for k := range m {
		combined = append(combined, k)
	}

	return combined
}

// ----------------------------------------------------------------------------
// Substitution

// Substitution represents the mapping of type variables to their inferred results.
type Substitution struct {
	Var  TyVar
	Type Typ
}

func ConvertTo(ss []Substitution) []Constraint {
	cs := []Constraint{}

	for _, s := range ss {
		tmp := Constraint{
			Left:  &s.Var,
			Right: s.Type,
		}
		cs = append(cs, tmp)
	}

	return cs
}

func Substitute(ss []Substitution, typ Typ) Typ {
	switch t := typ.(type) {
	case *TyInt:
		return &TyInt{}
	case *TyBool:
		return &TyBool{}
	case *TyFun:
		return &TyFun{Abs: Substitute(ss, t.Abs), App: Substitute(ss, t.App)}
	case *TyVar:
		for _, s := range ss {
			if s.Var == *t {
				return s.Type
			}
		}
		return t
	default:
		log.Fatalf("unexpected type: %T", t)
	}

	return nil
}
