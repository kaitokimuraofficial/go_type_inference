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
		obj, ok := env.Get(node)
		if !ok {
			log.Fatal("Variable not bound")
		}
		return obj
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
		log.Fatal("Not Implemented!")
	}

	return nil
}

func inferStatement(s ast.Statement, env *object.TypeEnvironment) object.InferredObject {
	return Infer(s.Expr, env)
}

func inferDeclaration(d ast.Declaration, env *object.TypeEnvironment) object.InferredObject {
	v := Infer(d.Expr, env)
	env.Set(d.Id, v)

	return v
}

func inferBinOpExpr(be ast.BinOpExpr, env *object.TypeEnvironment) object.InferredObject {
	lt := Infer(be.Left, env)
	rt := Infer(be.Right, env)

	switch be.Token.Type {
	case token.PLUS:
		if lt.Type() == object.INTEGER_TYPE && rt.Type() == object.INTEGER_TYPE {
			return object.TyInt{}
		} else {
			log.Fatal("Both arguments must be integer: +")
		}
	case token.ASTERISK:
		if lt.Type() == object.INTEGER_TYPE && rt.Type() == object.INTEGER_TYPE {
			return object.TyInt{}
		} else {
			log.Fatal("Both arguments must be integer: *")
		}
	case token.LT:
		if lt.Type() == object.INTEGER_TYPE && rt.Type() == object.INTEGER_TYPE {
			return object.TyBool{}
		} else {
			log.Fatal("Both arguments must be integer: <")
		}
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
		return consType
	}

	if consType.Type() == altType.Type() {
		return consType
	}

	return nil
}

func inferLetExpr(le ast.LetExpr, env *object.TypeEnvironment) object.InferredObject {
	bindType := Infer(le.BindingExpr, env)
	env.Set(le.Id, bindType)

	return Infer(le.BodyExpr, env)
}
