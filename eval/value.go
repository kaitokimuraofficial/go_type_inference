package eval

import "go_type_inference/ast"

type Value interface {
	value()
}

type (
	Integer struct {
		Value int
	}

	Boolean struct {
		Value bool
	}

	Function struct {
		Param ast.Ident
		Body  ast.Expr
		Env   Environment
	}
)

func (Integer) value()  {}
func (Boolean) value()  {}
func (Function) value() {}
