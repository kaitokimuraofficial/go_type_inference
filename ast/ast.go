package ast

import (
	"fmt"
	"go_type_inference/token"
	"strconv"
)

type Node interface {
	TokenLiteral() string
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
func (s Statement) TokenLiteral() string {
	return s.Expr.TokenLiteral()
}
func (s Statement) String() string {
	return s.Expr.String()
}

func (d Declaration) statementNode() {}
func (d Declaration) TokenLiteral() string {
	return d.Expr.TokenLiteral()
}
func (d Declaration) String() string {
	return fmt.Sprintf(
		"let %s = %s",
		d.Id.String(),
		d.Expr.String(),
	)
}

func (rd RecDeclaration) statementNode() {}
func (rd RecDeclaration) TokenLiteral() string {
	return rd.BodyExpr.TokenLiteral()
}
func (rd RecDeclaration) String() string {
	return fmt.Sprintf(
		"let rec %s = %s",
		rd.Id.String(),
		rd.BodyExpr.String(),
	)
}

type Integer struct {
	Token token.Token
	Value int
}

type Boolean struct {
	Token token.Token
	Value bool
}

type Identifier struct {
	Token token.Token
	Value string
}

type BinOpExpr struct {
	Token    token.Token
	Left     Expr
	Operator string
	Right    Expr
}

type IfExpr struct {
	Token       token.Token
	Condition   Expr
	Consequence Expr
	Alternative Expr
}

type LetExpr struct {
	Token       token.Token
	Id          Identifier
	BindingExpr Expr
	BodyExpr    Expr
}

type FunExpr struct {
	Token    token.Token
	Param    Identifier
	BodyExpr Expr
}

type AppExpr struct {
	Token    token.Token
	Function Expr
	Argument Expr
}

type LetRecExpr struct {
	Token       token.Token
	Id          Identifier
	Param       Identifier
	BindingExpr Expr
	BodyExpr    Expr
}

func (b Boolean) expressionNode() {}
func (b Boolean) TokenLiteral() string {
	return b.Token.Literal
}
func (b Boolean) String() string {
	return strconv.FormatBool(b.Value)
}

func (i Integer) expressionNode() {}
func (i Integer) TokenLiteral() string {
	return i.Token.Literal
}
func (i Integer) String() string {
	return strconv.Itoa(i.Value)
}

func (be BinOpExpr) expressionNode() {}
func (be BinOpExpr) TokenLiteral() string {
	return be.Token.Literal
}
func (be BinOpExpr) String() string {
	return fmt.Sprintf(
		"%s %s %s",
		be.Left.String(),
		be.Operator,
		be.Left.String(),
	)
}

func (ie IfExpr) expressionNode() {}
func (ie IfExpr) TokenLiteral() string {
	return ie.Token.Literal
}
func (ie IfExpr) String() string {
	return fmt.Sprintf(
		"if %s then ( %s ) else ( %s )",
		ie.Condition.String(),
		ie.Consequence.String(),
		ie.Alternative.String(),
	)
}

func (le LetExpr) expressionNode() {}
func (le LetExpr) TokenLiteral() string {
	return le.Token.Literal
}
func (le LetExpr) String() string {
	return fmt.Sprintf(
		"let %s = %s in %s",
		le.Id.String(),
		le.BindingExpr.String(),
		le.BodyExpr.String(),
	)
}

func (fe FunExpr) expressionNode() {}
func (fe FunExpr) TokenLiteral() string {
	return fe.Token.Literal
}
func (fe FunExpr) String() string {
	return fmt.Sprintf(
		"fun %s -> %s",
		fe.Param.String(),
		fe.BodyExpr.String(),
	)
}

func (ae AppExpr) expressionNode() {}
func (ae AppExpr) TokenLiteral() string {
	return ae.Token.Literal
}
func (ae AppExpr) String() string {
	return fmt.Sprintf(
		"(%s, %s)",
		ae.Function.String(),
		ae.Argument.String(),
	)
}

func (lr LetRecExpr) expressionNode() {}
func (lr LetRecExpr) TokenLiteral() string {
	return lr.Token.Literal
}
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
func (i Identifier) TokenLiteral() string {
	return i.Token.Literal
}
func (i Identifier) String() string {
	return i.Value
}
