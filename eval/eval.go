package eval

import (
	"go_type_inference/ast"
	"go_type_inference/object"
	"go_type_inference/token"
	"log"
)

func Eval(node ast.Node, env *object.Environment) object.Value {
	switch node := node.(type) {
	case ast.Statement:
		return evalStatement(node, env)
	case ast.Declaration:
		return evalDeclaration(node, env)
	case ast.RecDeclaration:
		return evalRecDeclaration(node, env)
	case ast.Identifier:
		return evalIdentifier(node, env)
	case ast.Integer:
		return object.Integer{Value: node.Value}
	case ast.Boolean:
		return object.Boolean{Value: node.Value}
	case ast.FunExpr:
		return object.Function{Param: node.Param, Body: node.BodyExpr, Env: *env}
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
	default:
		log.Fatalf("Evaluator not implemented for node: %s", node.String())
	}

	return nil
}

func evalStatement(s ast.Statement, env *object.Environment) object.Value {
	return Eval(s.Expr, env)
}

func evalDeclaration(d ast.Declaration, env *object.Environment) object.Value {
	v := Eval(d.Expr, env)
	env.Set(d.Id, v)
	return v
}

func evalRecDeclaration(rd ast.RecDeclaration, env *object.Environment) object.Value {
	dummyEnv := object.NewEnvironment()
	f := object.Function{Param: rd.Param, Body: rd.BodyExpr, Env: *dummyEnv}
	env.Set(rd.Id, f)
	return f
}

func evalIdentifier(i ast.Identifier, env *object.Environment) object.Value {
	val, ok := env.Get(i)
	if !ok {
		log.Fatalf("Variable '%s' not bound", i.Value)
	}
	return val
}

// While the binary operator operands are not strictly required to be integers,
// this program expects both operands to be integers.
func evalBinOpExpr(be ast.BinOpExpr, env *object.Environment) object.Value {
	lv := Eval(be.Left, env)
	rv := Eval(be.Right, env)

	if lv.Type() != object.INTEGER_VAL || rv.Type() != object.INTEGER_VAL {
		log.Fatalf("Both arguments must be integers for operator %d", be.Type)
	}

	switch be.Type {
	case token.PLUS:
		return object.Integer{Value: lv.(object.Integer).Value + rv.(object.Integer).Value}
	case token.ASTERISK:
		return object.Integer{Value: lv.(object.Integer).Value * rv.(object.Integer).Value}
	case token.LT:
		return object.Boolean{Value: lv.(object.Integer).Value < rv.(object.Integer).Value}
	default:
		log.Fatal("The combination of binary operator and argument is incorrect: BinOp")
	}

	return nil
}

func evalIfExpr(ie ast.IfExpr, env *object.Environment) object.Value {
	cnd := Eval(ie.Condition, env)

	if cnd.Type() != object.BOOLEAN_VAL {
		log.Fatal("Condition must be boolean: If")
	}

	if cnd.(object.Boolean).Value {
		return Eval(ie.Consequence, env)
	}

	return Eval(ie.Alternative, env)
}

func evalLetExpr(le ast.LetExpr, env *object.Environment) object.Value {
	v := Eval(le.BindingExpr, env)
	env.Set(le.Id, v)
	return Eval(le.BodyExpr, env)
}

func evalAppExpr(ae ast.AppExpr, env *object.Environment) object.Value {
	fn := Eval(ae.Function, env)
	arg := Eval(ae.Argument, env)

	if fn, ok := fn.(object.Function); ok {
		env.Set(fn.Param, arg)
		return Eval(fn.Body, env)
	}

	log.Fatalf("Non-function value applied: %v", fn)
	return nil

}

func evalLetRecExpr(lr ast.LetRecExpr, env *object.Environment) object.Value {
	f := object.Function{Param: lr.Param, Body: lr.BindingExpr, Env: *env}
	env.Set(lr.Id, f)
	return Eval(lr.BodyExpr, env)
}
