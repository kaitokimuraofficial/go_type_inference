package typing

import (
	"go_type_inference/ast"
	"go_type_inference/token"
	"log"
	"strconv"
)

// Infer receives term and type environment, returns substitution and type
func Infer(node ast.Node, env *Environment) ([]Substitution, Type) {
	switch n := node.(type) {
	case *ast.DeclStmt:
		return Infer(n.Decl, env)
	case *ast.ExprStmt:
		return Infer(n.Expr, env)
	case *ast.LetDecl:
		return inferLetDecl(*n, env)
	case *ast.Integer:
		return []Substitution{}, &TyInt{}
	case *ast.Boolean:
		return []Substitution{}, &TyBool{}
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

func inferLetDecl(d ast.LetDecl, env *Environment) ([]Substitution, Type) {
	s, t := Infer(d.Expr, env)
	cs := ConvertTo(s)
	subst := Unify(cs)

	typ := Substitute(subst, t)
	sch := NewScheme(typ)
	for _, v := range typ.Variables() {
		vs := strconv.Itoa(int(v))
		if _, ok := env.Get(ast.Identifier{Value: vs}); !ok {
			sch.BoundVars = append(sch.BoundVars, v)
		}
	}

	env.Set(d.Id, *sch)
	return []Substitution{}, sch.Type
}

func inferIdentifier(i ast.Identifier, env *Environment) ([]Substitution, Type) {
	sch, ok := env.Get(i)
	if !ok {
		log.Fatalf("variable %q is not bound", i.Value)
	}

	subst := []Substitution{}
	for _, boundVar := range sch.BoundVars {
		subst = append(subst, struct {
			Variable Variable
			Type     Type
		}{
			Variable: boundVar,
			Type:     NewFreshTyIdent(),
		})
	}

	return []Substitution{}, Substitute(subst, sch.Type)
}

func inferBinOpExpr(e ast.BinOpExpr, env *Environment) ([]Substitution, Type) {
	ls, lt := Infer(e.Left, env)
	rs, rt := Infer(e.Right, env)

	c, t := inferPrimitive(e.Op, lt, rt)

	newCS := Union(ConvertTo(ls), ConvertTo(rs), c)

	s := Unify(newCS)

	return s, Substitute(s, t)
}

// inferPrimitive receives token.Type and two Type, returns Constraints and Type
func inferPrimitive(op token.Type, left Type, right Type) ([]Constraint, Type) {
	switch op {
	case token.PLUS:
		c := []Constraint{
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
		c := []Constraint{
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
		c := []Constraint{
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

func inferIfExpr(e ast.IfExpr, env *Environment) ([]Substitution, Type) {
	s1, t1 := Infer(e.Condition, env)
	s2, t2 := Infer(e.Consequence, env)
	s3, t3 := Infer(e.Alternative, env)

	cs1 := []Constraint{
		{
			Left:  t1,
			Right: &TyBool{},
		},
	}

	cs2 := []Constraint{
		{
			Left:  t2,
			Right: t3,
		},
	}

	newCS := Union(ConvertTo(s1), ConvertTo(s2), ConvertTo(s3), cs1, cs2)

	s := Unify(newCS)

	return s, Substitute(s, t2)
}

func inferLetExpr(e ast.LetExpr, env *Environment) ([]Substitution, Type) {
	s1, t1 := Infer(e.BindingExpr, env)
	cs := ConvertTo(s1)
	subst := Unify(cs)

	bindingTyp := Substitute(subst, t1)
	sch := NewScheme(bindingTyp)
	for _, v := range bindingTyp.Variables() {
		vs := strconv.Itoa(int(v))
		if _, ok := env.Get(ast.Identifier{Value: vs}); !ok {
			sch.BoundVars = append(sch.BoundVars, v)
		}
	}

	env.Set(e.Id, *sch)
	s2, t2 := Infer(e.BodyExpr, env)

	newCS := Union(ConvertTo(s1), ConvertTo(s2))

	s := Unify(newCS)

	return s, Substitute(s, t2)
}

func inferFunExpr(e ast.FunExpr, env *Environment) ([]Substitution, Type) {
	freshIdent := NewFreshTyIdent()
	env.Set(e.Param, *NewScheme(freshIdent))

	s, t := Infer(e.BodyExpr, env)

	return s, &TyFun{Abs: Substitute(s, freshIdent), App: t}
}

func inferAppExpr(e ast.AppExpr, env *Environment) ([]Substitution, Type) {
	s1, t1 := Infer(e.Function, env)
	s2, t2 := Infer(e.Argument, env)

	freshIdent := NewFreshTyIdent()

	cs := []Constraint{
		{
			Left: t1,
			Right: &TyFun{
				Abs: t2,
				App: freshIdent,
			},
		},
	}
	newCS := Union(ConvertTo(s1), ConvertTo(s2), cs)

	s := Unify(newCS)

	return s, Substitute(s, freshIdent)
}
