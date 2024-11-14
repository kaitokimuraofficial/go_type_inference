package typing

import (
	"go_type_inference/ast"
	"go_type_inference/token"
	"log"
)

func Infer(node ast.Node, env *Environment) Type {
	switch n := node.(type) {
	case *ast.DeclStmt:
		return Infer(n.Decl, env)
	case *ast.ExprStmt:
		return Infer(n.Expr, env)
	case *ast.LetDecl:
		return inferLetDecl(*n, env)
	case *ast.Integer:
		return &TyInt{}
	case *ast.Boolean:
		return &TyBool{}
	case *ast.Identifier:
		return inferIdentifier(*n, env)
	case *ast.BinOpExpr:
		return inferBinOpExpr(*n, env)
	case *ast.IfExpr:
		return inferIfExpr(*n, env)
	case *ast.LetExpr:
		return inferLetExpr(*n, env)
	default:
		log.Fatalf("unexpected node type: %T", n)
	}

	return nil
}

func inferLetDecl(d ast.LetDecl, env *Environment) Type {
	v := Infer(d.Expr, env)
	env.Set(d.Id, v)
	return v
}

func inferIdentifier(i ast.Identifier, env *Environment) Type {
	t, ok := env.Get(i)
	if !ok {
		log.Fatalf("variable %q is not bound", i.Value)
	}
	return t
}

// While the binary operator operands are not strictly required to be integers,
// this program expects both operands to be integers.
func inferBinOpExpr(e ast.BinOpExpr, env *Environment) Type {
	_, ok := Infer(e.Left, env).(*TyInt)
	if !ok {
		log.Fatalf("left operand type is not TyInt")
	}

	_, ok = Infer(e.Right, env).(*TyInt)
	if !ok {
		log.Fatalf("right operand type is not TyInt")
	}

	switch e.Op {
	case token.PLUS:
		return &TyInt{}
	case token.ASTERISK:
		return &TyInt{}
	case token.LT:
		return &TyBool{}
	default:
		log.Fatalf("%s is not supported operator", e.Op)
	}

	return nil
}

func inferIfExpr(e ast.IfExpr, env *Environment) Type {
	_, ok := Infer(e.Condition, env).(*TyBool)
	if !ok {
		log.Fatalf("condition is not Boolean")
	}

	cons := Infer(e.Consequence, env)
	alt := Infer(e.Alternative, env)

	if cons == alt {
		return cons
	}
	return alt
}

func inferLetExpr(e ast.LetExpr, env *Environment) Type {
	t := Infer(e.BindingExpr, env)
	env.Set(e.Id, t)
	return Infer(e.BodyExpr, env)
}
