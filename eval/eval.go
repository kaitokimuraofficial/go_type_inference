package eval

import (
	"go_type_inference/ast"
	"go_type_inference/token"
	"log"
)

func Eval(node ast.Node, env Env) Value {
	switch n := node.(type) {
	case ast.DeclStmt:
		return Eval(n.Decl, env)
	case ast.ExprStmt:
		return Eval(n.Expr, env)
	case ast.LetDecl:
		return evalLetDecl(n, env)
	case ast.RecDecl:
		return evalRecDecl(n, env)
	case ast.Integer:
		return Integer{Value: n.Value}
	case ast.Boolean:
		return Boolean{Value: n.Value}
	case ast.Ident:
		return evalIdent(n, env)
	case ast.BinOpExpr:
		return evalBinOpExpr(n, env)
	case ast.IfExpr:
		return evalIfExpr(n, env)
	case ast.LetExpr:
		return evalLetExpr(n, env)
	case ast.FunExpr:
		return Function{Param: n.Param, Body: n.Body, Env: env}
	case ast.AppExpr:
		return evalAppExpr(n, env)
	case ast.LetRecExpr:
		return evalLetRecExpr(n, env)
	default:
		log.Fatalf("unexpected node type: %T", n)
	}

	return nil
}

func evalLetDecl(d ast.LetDecl, env Env) Value {
	v := Eval(d.Expr, env)
	env.Set(d.Id, v)
	return v
}

func evalRecDecl(d ast.RecDecl, env Env) Value {
	emptyEnv := Env{Store: make(map[ast.Ident]Value)}
	f := Function{Param: d.Param, Body: d.Body, Env: emptyEnv}
	env.Set(d.Id, f)
	return f
}

func evalIdent(i ast.Ident, env Env) Value {
	val, ok := env.Get(i)
	if !ok {
		log.Fatalf("variable %q is not bound", i.Value)
	}
	return val
}

// While the binary operator operands are not strictly required to be integers,
// this program expects both operands to be integers.
func evalBinOpExpr(e ast.BinOpExpr, env Env) Value {
	lv, ok := Eval(e.Left, env).(Integer)
	if !ok {
		log.Fatal("left operand is not Integer")
	}

	rv, ok := Eval(e.Right, env).(Integer)
	if !ok {
		log.Fatal("right operand is not Integer")
	}

	switch e.Op {
	case token.PLUS:
		return Integer{Value: lv.Value + rv.Value}
	case token.ASTERISK:
		return Integer{Value: lv.Value * rv.Value}
	case token.LT:
		return Boolean{Value: lv.Value < rv.Value}
	default:
		log.Fatalf("%s is not supported operator", e.Op)
	}

	return nil
}

func evalIfExpr(e ast.IfExpr, env Env) Value {
	v, ok := Eval(e.Cond, env).(Boolean)
	if !ok {
		log.Fatal("condition is not Boolean")
	}

	if v.Value {
		return Eval(e.Cons, env)
	}

	return Eval(e.Alt, env)
}

func evalLetExpr(e ast.LetExpr, env Env) Value {
	v := Eval(e.Bind, env)
	env.Set(e.Id, v)
	return Eval(e.Body, env)
}

func evalAppExpr(e ast.AppExpr, env Env) Value {
	fn := Eval(e.Func, env)
	arg := Eval(e.Arg, env)

	v, ok := fn.(Function)
	if !ok {
		log.Fatal("left-hand side expression is not a function abstraction")
	}

	env.Set(v.Param, arg)
	return Eval(v.Body, env)
}

func evalLetRecExpr(e ast.LetRecExpr, env Env) Value {
	emptyEnv := Env{Store: make(map[ast.Ident]Value)}
	f := Function{Param: e.Param, Body: e.Bind, Env: emptyEnv}
	env.Set(e.Id, f)
	return Eval(e.Body, env)
}
