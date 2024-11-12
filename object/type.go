package object

type InferredType int

const (
	INTEGER_TYPE InferredType = iota
	BOOLEAN_TYPE
	FUNCTION_TYPE
	IDENTIFIER_TYPE
)

type InferredObject interface {
	Type() InferredType
}

type TyInt struct{}

func (t TyInt) Type() InferredType {
	return INTEGER_TYPE
}

type TyBool struct{}

func (t TyBool) Type() InferredType {
	return BOOLEAN_TYPE
}

type TyFun struct {
	Abs InferredType
	App InferredType
}

func (t TyFun) Type() InferredType {
	return FUNCTION_TYPE
}

type TyIdent struct {
	Value int
}

func (t TyIdent) Type() InferredType {
	return IDENTIFIER_TYPE
}
