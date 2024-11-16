package typing

import "log"

// ----------------------------------------------------------------------------
// Constraint

// Constraint represents a set of constraints that the type must satisfy.
type Constraint struct {
	Left  Type
	Right Type
}

func Unify(cs []Constraint) []Substitution {
	for i, c := range cs {
		left, right := c.Left, c.Right

		if left == right {
			newCS := append(cs[:i], cs[i+1:]...)
			return Unify(newCS)
		}

		lf, lok := left.(*TyFun)
		rf, rok := right.(*TyFun)
		if lok && rok {
			newCS := append(cs[:i], cs[i+1:]...)
			newCS = append(newCS,
				Constraint{Left: lf.Abs, Right: rf.Abs},
				Constraint{Left: lf.App, Right: rf.App},
			)
			return Unify(newCS)
		}

		li, lok := left.(*TyIdent)
		if lok && !ContainsIn(right.Variables(), li.Variable) {
			newCS := append(cs[:i], cs[i+1:]...)
			replaced := replace(newCS, *li, right)
			substitution := Unify(replaced)
			return append(substitution, struct {
				Variable Variable
				Type     Type
			}{
				Variable: li.Variable,
				Type:     right,
			})
		}

		ri, rok := right.(*TyIdent)
		if rok && !ContainsIn(left.Variables(), ri.Variable) {
			newCS := append(cs[:i], cs[i+1:]...)
			replaced := replace(newCS, *ri, left)
			substitution := Unify(replaced)
			return append(substitution, struct {
				Variable Variable
				Type     Type
			}{
				Variable: ri.Variable,
				Type:     left,
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

// replace replaces all occurrences of the 'frm' TyIdent in the Constraints to the 'to' type.
func replace(cs []Constraint, frm TyIdent, to Type) []Constraint {
	replaced := []Constraint{}

	for _, c := range cs {
		replaced = append(replaced, Constraint{
			Left:  c.Left.Convert(frm, to),
			Right: c.Right.Convert(frm, to),
		})
	}
	return replaced
}

// ----------------------------------------------------------------------------
// Substitution

// Substitution represents the mapping of type variables to their inferred results.
type Substitution struct {
	Variable Variable
	Type     Type
}

func ConvertTo(ss []Substitution) []Constraint {
	cs := []Constraint{}

	for _, s := range ss {
		tmp := Constraint{
			Left:  &TyIdent{Variable: s.Variable},
			Right: s.Type,
		}
		cs = append(cs, tmp)
	}

	return cs
}

func Substitute(ss []Substitution, typ Type) Type {
	switch t := typ.(type) {
	case *TyInt:
		return &TyInt{}
	case *TyBool:
		return &TyBool{}
	case *TyFun:
		return &TyFun{Abs: Substitute(ss, t.Abs), App: Substitute(ss, t.App)}
	case *TyIdent:
		for _, s := range ss {
			if s.Variable == t.Variable {
				return s.Type
			}
		}
		return t
	default:
		log.Fatalf("unexpected type: %T", t)
	}

	return nil
}
