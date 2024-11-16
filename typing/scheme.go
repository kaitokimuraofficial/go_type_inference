package typing

type Scheme struct {
	BoundVars []Variable
	Type      Type
}

func NewScheme(typ Type) *Scheme {
	return &Scheme{BoundVars: []Variable{}, Type: typ}
}

func FreeVariables(s Scheme) []Variable {
	vars := s.Type.Variables()
	return difference(vars, s.BoundVars)
}
