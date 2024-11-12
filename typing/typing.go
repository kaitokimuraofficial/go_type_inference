package typing

import (
	"go_type_inference/ast"
	"go_type_inference/object"
	"go_type_inference/token"
	"log"
)

func Infer(node ast.Node, env *object.TypeEnvironment) object.InferredObject {
	switch node := node.(type) {
	case ast.Statement:
		return inferStatement(node, env)
	case ast.Declaration:
		return inferDeclaration(node, env)
	case ast.Identifier:
		return inferIdentifier(node, env)
	case ast.Integer:
		return object.TyInt{}
	case ast.Boolean:
		return object.TyBool{}
	case ast.BinOpExpr:
		return inferBinOpExpr(node, env)
	case ast.IfExpr:
		return inferIfExpr(node, env)
	case ast.LetExpr:
		return inferLetExpr(node, env)
	default:
		log.Fatalf("Type inference not implemented for node type: %s", node.String())
	}

	return nil
}

func inferStatement(s ast.Statement, env *object.TypeEnvironment) object.InferredObject {
	return Infer(s.Expr, env)
}

func inferDeclaration(d ast.Declaration, env *object.TypeEnvironment) object.InferredObject {
	t := Infer(d.Expr, env)
	env.Set(d.Id, t)
	return t
}

func inferIdentifier(i ast.Identifier, env *object.TypeEnvironment) object.InferredObject {
	t, ok := env.Get(i)
	if !ok {
		log.Fatal("Variable not bound")
	}
	return t
}

// While the binary operator operands are not strictly required to be integers,
// this program expects both operands to be integers.
func inferBinOpExpr(be ast.BinOpExpr, env *object.TypeEnvironment) object.InferredObject {
	lt := Infer(be.Left, env)
	rt := Infer(be.Right, env)

	if lt.Type() != object.INTEGER_TYPE || rt.Type() != object.INTEGER_TYPE {
		log.Fatalf("Both arguments must be Integer for operator %s", be.Token.Type)
	}

	switch be.Token.Type {
	case token.PLUS:
		return object.TyInt{}
	case token.ASTERISK:
		return object.TyInt{}
	case token.LT:
		return object.TyBool{}
	default:
		log.Fatal("The combination of binary operator and argument is incorrect: BinOp")
	}

	return nil
}

func inferIfExpr(ie ast.IfExpr, env *object.TypeEnvironment) object.InferredObject {
	cndType := Infer(ie.Condition, env)
	consType := Infer(ie.Consequence, env)
	altType := Infer(ie.Alternative, env)

	if cndType.Type() != object.BOOLEAN_TYPE {
		log.Fatal("Not Implemented!")
	}

	if consType.Type() == altType.Type() {
		return consType
	}

	log.Fatalf("consequence and alternative types do not match: If")
	return nil
}

func inferLetExpr(le ast.LetExpr, env *object.TypeEnvironment) object.InferredObject {
	t := Infer(le.BindingExpr, env)
	env.Set(le.Id, t)
	return Infer(le.BodyExpr, env)
}
