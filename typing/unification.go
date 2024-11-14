package typing

import "log"

// ----------------------------------------------------------------------------
// ConstraintSet

// ConstraintSet represents a set of constraints that the type must satisfy.
type ConstraintSet []struct {
	Left  Type
	Right Type
}

func (cs ConstraintSet) Unify() Substitution {
	for i, c := range cs {
		left, right := c.Left, c.Right

		if left == right {
			newCS := append(cs[:i], cs[i+1:]...)
			return newCS.Unify()
		}

		lf, lok := left.(*TyFun)
		rf, rok := right.(*TyFun)
		if lok && rok {
			newCS := append(cs[:i], cs[i+1:]...)
			newCS = append(newCS,
				struct {
					Left  Type
					Right Type
				}{
					Left:  lf.Abs,
					Right: rf.Abs,
				},
				struct {
					Left  Type
					Right Type
				}{
					Left:  lf.App,
					Right: rf.App,
				})
			return newCS.Unify()
		}

		li, lok := left.(*TyIdent)
		if lok && !ContainsIn(right.Variables(), li.Variable) {
			newCS := append(cs[:i], cs[i+1:]...)
			replacedConstSet := newCS.replace(*li, right)
			res := replacedConstSet.Unify()
			return append(res, struct {
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
			replacedConstSet := newCS.replace(*ri, left)
			res := replacedConstSet.Unify()
			return append(res, struct {
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

// Union combines two slices of ConstraintSets into a single slice with no duplicates.
func Union(css ...ConstraintSet) ConstraintSet {
	m := make(map[struct {
		Left  Type
		Right Type
	}]bool)
	for _, cs := range css {
		for _, v := range cs {
			m[v] = true
		}
	}

	keys := make([]struct {
		Left  Type
		Right Type
	}, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}

	return keys
}

// replace replaces all occurrences of the 'frm' TyIdent in the ConstraintSet to the 'to' type.
func (cs ConstraintSet) replace(frm TyIdent, to Type) ConstraintSet {
	newConstSet := make(ConstraintSet, len(cs))

	for i, pair := range cs {
		newConstSet[i] = struct {
			Left  Type
			Right Type
		}{
			Left:  pair.Left.Convert(frm, to),
			Right: pair.Right.Convert(frm, to),
		}
	}
	return newConstSet
}

// ----------------------------------------------------------------------------
// Substitution

// Substitution represents the mapping of type variables to their inferred results.
type Substitution []struct {
	Variable Variable
	Type     Type
}

func (s *Substitution) ConvertTo() ConstraintSet {
	cs := ConstraintSet{}

	for _, subst := range *s {
		tmp := struct {
			Left  Type
			Right Type
		}{
			Left:  &TyIdent{Variable: subst.Variable},
			Right: subst.Type,
		}
		cs = append(cs, tmp)
	}

	return cs
}

func (s *Substitution) Substitute(typ Type) Type {
	switch t := typ.(type) {
	case *TyInt:
		return &TyInt{}
	case *TyBool:
		return &TyBool{}
	case *TyFun:
		return &TyFun{Abs: s.Substitute(t.Abs), App: s.Substitute(t.App)}
	case *TyIdent:
		for _, subst := range *s {
			if subst.Variable == t.Variable {
				return subst.Type
			}
		}
		return t
	default:
		log.Fatalf("unexpected type: %T", t)
	}

	return nil
}
