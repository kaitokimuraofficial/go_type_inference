package eval

import (
	"go_type_inference/ast"
	"go_type_inference/object"
	"go_type_inference/token"
	"log"
)

func Eval(node ast.Node, env *object.Environment) object.Object {
	switch node := node.(type) {
	case ast.Statement:
		return evalStatement(node, env)
	case ast.Declaration:
		return evalDeclaration(node, env)
	case ast.Identifier:
		obj, ok := env.Get(node.Value)
		if !ok {
			log.Fatal("Variable not bound")
		}
		switch obj := obj.(type) {
		case *object.Integer:
			return &object.Integer{Value: obj.Value}
		case *object.Boolean:
			return &object.Boolean{Value: obj.Value}
		}
	case ast.Integer:
		return &object.Integer{Value: node.Value}
	case ast.Boolean:
		return &object.Boolean{Value: node.Value}
	case ast.BinOpExpr:
		return evalBinOpExpr(node, env)
	case ast.IfExpr:
		return evalIfExpr(node, env)
	case ast.LetExpr:
		return evalLetExpr(node, env)
	}

	return nil
}

func evalStatement(s ast.Statement, env *object.Environment) object.Object {
	return Eval(s.Expr, env)
}

func evalDeclaration(d ast.Declaration, env *object.Environment) object.Object {
	v := Eval(d.Expr, env)
	env.Set(d.Id.Value, v)
	return v
}

func evalBinOpExpr(be ast.BinOpExpr, env *object.Environment) object.Object {
	lv := Eval(be.Left, env)
	rv := Eval(be.Right, env)

	switch be.Token.Type {
	case token.PLUS:
		if lv.Type() == object.INTEGER_OBJ && rv.Type() == object.INTEGER_OBJ {
			return &object.Integer{Value: lv.(*object.Integer).Value + rv.(*object.Integer).Value}
		} else {
			log.Fatal("Both arguments must be integer: +")
		}
	case token.ASTERISK:
		if lv.Type() == object.INTEGER_OBJ && rv.Type() == object.INTEGER_OBJ {
			return &object.Integer{Value: lv.(*object.Integer).Value * rv.(*object.Integer).Value}
		} else {
			log.Fatal("Both arguments must be integer: *")
		}
	case token.LT:
		if lv.Type() == object.INTEGER_OBJ && rv.Type() == object.INTEGER_OBJ {
			return &object.Boolean{Value: lv.(*object.Integer).Value < rv.(*object.Integer).Value}
		} else {
			log.Fatal("Both arguments must be integer: <")
		}
	default:
		log.Fatal("The combination of binary operator and argument is incorrect: BinOp")
	}

	return nil
}

func evalIfExpr(ie ast.IfExpr, env *object.Environment) object.Object {
	cnd := Eval(ie.Condition, env)

	if cnd.Type() == object.BOOLEAN_OBJ {
		if cnd.(*object.Boolean).Value {
			return Eval(ie.Consequence, env)
		} else {
			return Eval(ie.Alternative, env)
		}
	}

	log.Fatal("Condition must be boolean: If")

	return nil
}

func evalLetExpr(le ast.LetExpr, env *object.Environment) object.Object {
	v := Eval(le.BindingExpr, env)
	env.Set(le.Identifier.Value, v)
	return Eval(le.BodyExpr, env)
}
