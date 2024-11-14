package typing

type Type interface {
	typ()
}

type (
	TyInt struct{}

	TyBool struct{}

	TyFun struct {
		Abs Type
		App Type
	}
	TyIdent struct {
		Value Type
	}
)

func (*TyInt) typ()   {}
func (*TyBool) typ()  {}
func (*TyFun) typ()   {}
func (*TyIdent) typ() {}
