package typing

import (
	"go_type_inference/ast"
	"go_type_inference/token"
	"log"
)

// Infer receives term and type environment, returns substitution and type
func Infer(node ast.Node, env *Environment) (Substitution, Type) {
	switch n := node.(type) {
	case *ast.DeclStmt:
		return Infer(n.Decl, env)
	case *ast.ExprStmt:
		return Infer(n.Expr, env)
	case *ast.LetDecl:
		return inferLetDecl(*n, env)
	// case *ast.RecDecl:
	// 	return inferRecDecl(*n, env)
	case *ast.Integer:
		return Substitution{}, &TyInt{}
	case *ast.Boolean:
		return Substitution{}, &TyBool{}
	case *ast.Identifier:
		return inferIdentifier(*n, env)
	case *ast.BinOpExpr:
		return inferBinOpExpr(*n, env)
	case *ast.IfExpr:
		return inferIfExpr(*n, env)
	case *ast.LetExpr:
		return inferLetExpr(*n, env)
	case *ast.FunExpr:
		return inferFunExpr(*n, env)
	case *ast.AppExpr:
		return inferAppExpr(*n, env)
	default:
		log.Fatalf("unexpected node type: %T", n)
	}

	return nil, nil
}

func inferLetDecl(d ast.LetDecl, env *Environment) (Substitution, Type) {
	_, t := Infer(d.Expr, env)
	env.Set(d.Id, t)
	return Substitution{}, t
}

// func inferRecDecl(d ast.RecDecl, env *Environment) (Substitution, Type) {
// 	return nil, nil
// }

func inferIdentifier(i ast.Identifier, env *Environment) (Substitution, Type) {
	t, ok := env.Get(i)
	if !ok {
		log.Fatalf("variable %q is not bound", i.Value)
	}
	return Substitution{}, t
}

func inferBinOpExpr(e ast.BinOpExpr, env *Environment) (Substitution, Type) {
	ls, lt := Infer(e.Left, env)
	rs, rt := Infer(e.Right, env)

	c, t := inferPrimitive(e.Op, lt, rt)

	newCS := Union(ls.ConvertTo(), rs.ConvertTo(), c)

	s := newCS.Unify()

	return s, s.Substitute(t)
}

// inferPrimitive receives token.Type and two Type, returns ConstraintSet and Type
func inferPrimitive(op token.Type, left Type, right Type) (ConstraintSet, Type) {
	switch op {
	case token.PLUS:
		c := ConstraintSet{
			{
				Left:  left,
				Right: &TyInt{},
			},
			{
				Left:  right,
				Right: &TyInt{},
			},
		}
		return c, &TyInt{}
	case token.ASTERISK:
		c := ConstraintSet{
			{
				Left:  left,
				Right: &TyInt{},
			},
			{
				Left:  right,
				Right: &TyInt{},
			},
		}
		return c, &TyInt{}
	case token.LT:
		c := ConstraintSet{
			{
				Left:  left,
				Right: &TyInt{},
			},
			{
				Left:  right,
				Right: &TyInt{},
			},
		}

		return c, &TyBool{}
	default:
		log.Fatalf("%s is not supported operator type", op)
	}

	return nil, nil
}

func inferIfExpr(e ast.IfExpr, env *Environment) (Substitution, Type) {
	s1, t1 := Infer(e.Condition, env)
	s2, t2 := Infer(e.Consequence, env)
	s3, t3 := Infer(e.Alternative, env)

	cs1 := ConstraintSet{
		{
			Left:  t1,
			Right: &TyBool{},
		},
	}

	cs2 := ConstraintSet{
		{
			Left:  t2,
			Right: t3,
		},
	}

	newCS := Union(s1.ConvertTo(), s2.ConvertTo(), s3.ConvertTo(), cs1, cs2)

	s := newCS.Unify()

	return s, s.Substitute(t2)
}

func inferLetExpr(e ast.LetExpr, env *Environment) (Substitution, Type) {
	s1, t1 := Infer(e.BindingExpr, env)
	env.Set(e.Id, t1)
	s2, t2 := Infer(e.BodyExpr, env)

	newCS := Union(s1.ConvertTo(), s2.ConvertTo())

	s := newCS.Unify()

	return s, s.Substitute(t2)
}

func inferFunExpr(e ast.FunExpr, env *Environment) (Substitution, Type) {
	freshVar := Fresh()
	freshIdent := &TyIdent{Variable: freshVar}
	env.Set(e.Param, freshIdent)

	s, t := Infer(e.BodyExpr, env)

	return s, &TyFun{Abs: s.Substitute(freshIdent), App: t}
}

func inferAppExpr(e ast.AppExpr, env *Environment) (Substitution, Type) {
	s1, t1 := Infer(e.Function, env)
	s2, t2 := Infer(e.Argument, env)

	freshVar := Fresh()
	freshIdent := &TyIdent{Variable: freshVar}

	cs1 := ConstraintSet{
		{
			Left: t1,
			Right: &TyFun{
				Abs: t2,
				App: freshIdent,
			},
		},
	}
	newCS := Union(s1.ConvertTo(), s2.ConvertTo(), cs1)

	s := newCS.Unify()

	return s, s.Substitute(freshIdent)
}
