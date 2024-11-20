package typing

import (
	"go_type_inference/ast"
	"go_type_inference/token"
	"log"
	"strconv"
)

// Infer receives term and type environment, returns substitution and type
func Infer(node ast.Node, env *Environment) ([]Substitution, Scheme) {
	switch n := node.(type) {
	case *ast.DeclStmt:
		return Infer(n.Decl, env)
	case *ast.ExprStmt:
		return Infer(n.Expr, env)
	case *ast.LetDecl:
		return inferLetDecl(*n, env)
	case *ast.Integer:
		return []Substitution{}, Scheme{BTV: []Variable{}, Type: &TyInt{}}
	case *ast.Boolean:
		return []Substitution{}, Scheme{BTV: []Variable{}, Type: &TyBool{}}
	case *ast.Identifier:
		return inferVar(*n, env)
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

	return nil, Scheme{BTV: []Variable{}, Type: nil}
}

func inferLetDecl(d ast.LetDecl, env *Environment) ([]Substitution, Scheme) {
	s, t := Infer(d.Expr, env)
	cs := ConvertTo(s)
	subst := Unify(cs)

	typ := Substitute(subst, t.Type)
	sch := NewScheme(typ)
	for _, v := range typ.Variables() {
		vs := strconv.Itoa(int(v))
		if _, ok := env.Get(ast.Identifier{Value: vs}); !ok {
			sch.BTV = append(sch.BTV, v)
		}
	}

	env.Set(d.Id, *sch)
	return []Substitution{}, *sch
}

func inferVar(i ast.Identifier, env *Environment) ([]Substitution, Scheme) {
	sch, ok := env.Get(i)
	if !ok {
		log.Fatalf("variable %q is not bound", i.Value)
	}

	subst := []Substitution{}
	for _, bv := range sch.BTV {
		subst = append(subst, Substitution{
			Var:  TyVar{Variable: bv},
			Type: FreshTyVar(),
		})
	}

	return []Substitution{}, Scheme{BTV: []Variable{}, Type: Substitute(subst, sch.Type)}
}

func inferBinOpExpr(e ast.BinOpExpr, env *Environment) ([]Substitution, Scheme) {
	ls, lt := Infer(e.Left, env)
	rs, rt := Infer(e.Right, env)

	var b Typ

	switch e.Op {
	case token.PLUS, token.ASTERISK:
		b = &TyInt{}
	case token.LT:
		b = &TyBool{}
	}

	newCS := Union(ConvertTo(ls), ConvertTo(rs), []Constraint{{
		Left:  lt.Type,
		Right: &TyInt{},
	},
		{
			Left:  rt.Type,
			Right: &TyInt{},
		}})

	s := Unify(newCS)

	return s, Scheme{BTV: []Variable{}, Type: Substitute(s, b)}
}

func inferIfExpr(e ast.IfExpr, env *Environment) ([]Substitution, Scheme) {
	s1, t1 := Infer(e.Condition, env)
	s2, t2 := Infer(e.Consequence, env)
	s3, t3 := Infer(e.Alternative, env)

	newCS := Union(ConvertTo(s1), ConvertTo(s2), ConvertTo(s3), []Constraint{
		{
			Left:  t1.Type,
			Right: &TyBool{},
		},
		{
			Left:  t2.Type,
			Right: t3.Type,
		}})

	s := Unify(newCS)

	return s, Scheme{BTV: []Variable{}, Type: Substitute(s, t2.Type)}
}

func inferLetExpr(e ast.LetExpr, env *Environment) ([]Substitution, Scheme) {
	s1, t1 := Infer(e.BindingExpr, env)
	cs := ConvertTo(s1)
	subst := Unify(cs)

	bindingTyp := Substitute(subst, t1.Type)
	sch := NewScheme(bindingTyp)
	for _, v := range bindingTyp.Variables() {
		vs := strconv.Itoa(int(v))
		if _, ok := env.Get(ast.Identifier{Value: vs}); !ok {
			sch.BTV = append(sch.BTV, v)
		}
	}

	env.Set(e.Id, *sch)
	s2, t2 := Infer(e.BodyExpr, env)

	newCS := Union(ConvertTo(s1), ConvertTo(s2))

	s := Unify(newCS)

	return s, Scheme{BTV: []Variable{}, Type: Substitute(s, t2.Type)}
}

func inferFunExpr(e ast.FunExpr, env *Environment) ([]Substitution, Scheme) {
	freshIdent := FreshTyVar()
	env.Set(e.Param, *NewScheme(freshIdent))

	s, t := Infer(e.BodyExpr, env)

	return s, Scheme{BTV: []Variable{}, Type: &TyFun{Abs: Substitute(s, freshIdent), App: t.Type}}
}

func inferAppExpr(e ast.AppExpr, env *Environment) ([]Substitution, Scheme) {
	s1, t1 := Infer(e.Function, env)
	s2, t2 := Infer(e.Argument, env)

	freshIdent := FreshTyVar()

	newCS := Union(ConvertTo(s1), ConvertTo(s2), []Constraint{
		{
			Left: t1.Type,
			Right: &TyFun{
				Abs: t2.Type,
				App: freshIdent,
			},
		},
	})

	s := Unify(newCS)

	return s, Scheme{BTV: []Variable{}, Type: Substitute(s, freshIdent)}
}
