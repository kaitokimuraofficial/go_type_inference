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

func (s Statement) statementNode() {}
func (s Statement) TokenLiteral() string {
	return s.Expr.TokenLiteral()
}
func (s Statement) String() string {
	var out bytes.Buffer

	out.WriteString(s.Expr.String())

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

func (i Identifier) expressionNode() {}
func (i Identifier) TokenLiteral() string {
	return i.Token.Literal
}
func (i Identifier) String() string {
	return i.Token.Literal
}
