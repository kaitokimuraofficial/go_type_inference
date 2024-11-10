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
	case ast.RecDeclaration:
		return evalRecDeclaration(node, env)
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
		case *object.Function:
			return obj
		}
	case ast.Integer:
		return &object.Integer{Value: node.Value}
	case ast.Boolean:
		return &object.Boolean{Value: node.Value}
	case ast.FunExpr:
		return &object.Function{Param: node.Param, Body: node.BodyExpr, Env: *env}
	case ast.BinOpExpr:
		return evalBinOpExpr(node, env)
	case ast.IfExpr:
		return evalIfExpr(node, env)
	case ast.LetExpr:
		return evalLetExpr(node, env)
	case ast.AppExpr:
		return evalAppExpr(node, env)
	case ast.LetRecExpr:
		return evalLetRecExpr(node, env)
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

func evalRecDeclaration(rd ast.RecDeclaration, env *object.Environment) object.Object {
	f := &object.Function{Param: rd.Param, Body: rd.BodyExpr, Env: *env}
	env.Set(rd.Id.Value, f)

	return f
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

func evalAppExpr(ae ast.AppExpr, env *object.Environment) object.Object {
	f := Eval(ae.Function, env)
	arg := Eval(ae.Argument, env)

	if f.Type() == object.FUNCTION_OBJ {
		env.Set(f.(*object.Function).Param.Value, arg)
		return Eval(f.(*object.Function).Body, env)
	}

	log.Fatal("Non function value is applied: App")

	return nil
}

func evalLetRecExpr(lr ast.LetRecExpr, env *object.Environment) object.Object {
	f := &object.Function{Param: lr.Param, Body: lr.BindingExpr, Env: *env}
	env.Set(lr.Id.Value, f)

	return Eval(lr.BodyExpr, env)
}
