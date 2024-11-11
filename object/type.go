package object

type InferType int

const (
	INTEGER_TYPE InferType = iota
	BOOLEAN_TYPE
	FUNCTION_TYPE
	IDENTIFER_TYPE
)

type InferObject interface {
	Type() InferType
}

type TyInt struct{}

func (t TyInt) Type() InferType {
	return INTEGER_TYPE
}

type TyBool struct{}

func (t TyBool) Type() InferType {
	return BOOLEAN_TYPE
}

type TyFun struct {
	Abs InferType
	App InferType
}

func (t TyFun) Type() InferType {
	return FUNCTION_TYPE
}

type TyIdent struct {
	Value int
}

func (t TyIdent) Type() InferType {
	return IDENTIFER_TYPE
}
