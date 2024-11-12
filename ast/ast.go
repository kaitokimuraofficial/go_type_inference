package ast

import (
	"bytes"
	"go_type_inference/token"
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
	var out bytes.Buffer

	out.WriteString(s.Expr.String())

	return out.String()
}

func (d Declaration) statementNode() {}
func (d Declaration) TokenLiteral() string {
	return d.Expr.TokenLiteral()
}
func (d Declaration) String() string {
	var out bytes.Buffer

	out.WriteString(d.Expr.String())

	return out.String()
}

func (r RecDeclaration) statementNode() {}
func (r RecDeclaration) TokenLiteral() string {
	return r.BodyExpr.TokenLiteral()
}
func (r RecDeclaration) String() string {
	var out bytes.Buffer

	out.WriteString("let rec ")
	out.WriteString(r.Id.Value)
	out.WriteString(" = ")
	out.WriteString(r.BodyExpr.String())

	return out.String()
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
	return b.Token.Literal
}

func (i Integer) expressionNode() {}
func (i Integer) TokenLiteral() string {
	return i.Token.Literal
}
func (i Integer) String() string {
	return i.Token.Literal
}

func (be BinOpExpr) expressionNode() {}
func (be BinOpExpr) TokenLiteral() string {
	return be.Token.Literal
}
func (be BinOpExpr) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(be.Left.String())
	out.WriteString(" " + be.Operator + " ")
	out.WriteString(be.Right.String())
	out.WriteString(")")

	return out.String()
}

func (ie IfExpr) expressionNode() {}
func (ie IfExpr) TokenLiteral() string {
	return ie.Token.Literal
}
func (ie IfExpr) String() string {
	var out bytes.Buffer

	out.WriteString("if ")
	out.WriteString(ie.Condition.String())
	out.WriteString(" then ( ")
	out.WriteString(ie.Consequence.String())
	out.WriteString(" ) else ( ")
	out.WriteString(ie.Alternative.String())
	out.WriteString(" ) ")

	return out.String()
}

func (le LetExpr) expressionNode() {}
func (le LetExpr) TokenLiteral() string {
	return le.Token.Literal
}
func (le LetExpr) String() string {
	var out bytes.Buffer

	out.WriteString("let ")
	out.WriteString(le.Id.Value)
	out.WriteString(" = ")
	out.WriteString(le.BindingExpr.String())
	out.WriteString(" in ")
	out.WriteString(le.BodyExpr.String())
	out.WriteString(" ")

	return out.String()
}

func (fe FunExpr) expressionNode() {}
func (fe FunExpr) TokenLiteral() string {
	// return fe.Token.Literal
	return fe.BodyExpr.TokenLiteral()
}
func (fe FunExpr) String() string {
	var out bytes.Buffer

	out.WriteString("fun ")
	out.WriteString(fe.Param.Value)
	out.WriteString(" -> ")
	out.WriteString(fe.BodyExpr.String())
	out.WriteString(" ")

	return out.String()
}

func (ae AppExpr) expressionNode() {}
func (ae AppExpr) TokenLiteral() string {
	return ae.Token.Literal
}
func (ae AppExpr) String() string {
	var out bytes.Buffer

	out.WriteString("( ")
	out.WriteString(ae.Function.String())
	out.WriteString(", ")
	out.WriteString(ae.Argument.String())
	out.WriteString(" ) ")

	return out.String()
}

func (lr LetRecExpr) expressionNode() {}
func (lr LetRecExpr) TokenLiteral() string {
	return lr.Token.Literal
}
func (lr LetRecExpr) String() string {
	var out bytes.Buffer

	out.WriteString("let rec ")
	out.WriteString(lr.Id.Value)
	out.WriteString(" = fun ")
	out.WriteString(lr.Param.Value)
	out.WriteString(" -> ")
	out.WriteString(lr.BindingExpr.String())
	out.WriteString(" in ")
	out.WriteString(lr.BodyExpr.String())

	return out.String()
}

func (i Identifier) expressionNode() {}
func (i Identifier) TokenLiteral() string {
	return i.Token.Literal
}
func (i Identifier) String() string {
	return i.Token.Literal
}
