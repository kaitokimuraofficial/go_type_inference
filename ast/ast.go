package ast

import (
	"fmt"
	"go_type_inference/token"
	"strconv"
)

type Node interface {
	String() string
}

type Stmt interface {
	Node
	statementNode()
}

type Expr interface {
	Node
	expressionNode()
}

type Statement struct {
	Expr Expr
}

type Declaration struct {
	Expr Expr
	Id   Identifier
}

type RecDeclaration struct {
	Id       Identifier
	Param    Identifier
	BodyExpr Expr
}

func (s Statement) statementNode() {}
func (s Statement) String() string {
	return s.Expr.String()
}

func (d Declaration) statementNode() {}
func (d Declaration) String() string {
	return fmt.Sprintf(
		"let %s = %s",
		d.Id.String(),
		d.Expr.String(),
	)
}

func (rd RecDeclaration) statementNode() {}
func (rd RecDeclaration) String() string {
	return fmt.Sprintf(
		"let rec %s = %s",
		rd.Id.String(),
		rd.BodyExpr.String(),
	)
}

type Integer struct {
	Value int
}

type Boolean struct {
	Value bool
}

type Identifier struct {
	Value string
}

type BinOpExpr struct {
	Type  token.Type
	Left  Expr
	Right Expr
}

type IfExpr struct {
	Condition   Expr
	Consequence Expr
	Alternative Expr
}

type LetExpr struct {
	Id          Identifier
	BindingExpr Expr
	BodyExpr    Expr
}

type FunExpr struct {
	Param    Identifier
	BodyExpr Expr
}

type AppExpr struct {
	Function Expr
	Argument Expr
}

type LetRecExpr struct {
	Id          Identifier
	Param       Identifier
	BindingExpr Expr
	BodyExpr    Expr
}

func (b Boolean) expressionNode() {}
func (b Boolean) String() string {
	return strconv.FormatBool(b.Value)
}

func (i Integer) expressionNode() {}
func (i Integer) String() string {
	return strconv.Itoa(i.Value)
}

func (be BinOpExpr) expressionNode() {}
func (be BinOpExpr) String() string {
	return fmt.Sprintf(
		"%s %s %s",
		be.Left.String(),
		be.Type,
		be.Left.String(),
	)
}

func (ie IfExpr) expressionNode() {}
func (ie IfExpr) String() string {
	return fmt.Sprintf(
		"if %s then ( %s ) else ( %s )",
		ie.Condition.String(),
		ie.Consequence.String(),
		ie.Alternative.String(),
	)
}

func (le LetExpr) expressionNode() {}
func (le LetExpr) String() string {
	return fmt.Sprintf(
		"let %s = %s in %s",
		le.Id.String(),
		le.BindingExpr.String(),
		le.BodyExpr.String(),
	)
}

func (fe FunExpr) expressionNode() {}
func (fe FunExpr) String() string {
	return fmt.Sprintf(
		"fun %s -> %s",
		fe.Param.String(),
		fe.BodyExpr.String(),
	)
}

func (ae AppExpr) expressionNode() {}
func (ae AppExpr) String() string {
	return fmt.Sprintf(
		"(%s, %s)",
		ae.Function.String(),
		ae.Argument.String(),
	)
}

func (lr LetRecExpr) expressionNode() {}
func (lr LetRecExpr) String() string {
	return fmt.Sprintf(
		"let rec %s = fun %s -> %s in %s",
		lr.Id.String(),
		lr.Param.String(),
		lr.BindingExpr.String(),
		lr.BodyExpr.String(),
	)
}

func (i Identifier) expressionNode() {}
func (i Identifier) String() string {
	return i.Value
}
