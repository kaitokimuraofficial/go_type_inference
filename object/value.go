package object

import (
	"bytes"
	"fmt"
	"go_type_inference/ast"
)

type ValueType int

const (
	INTEGER_VAL ValueType = iota
	BOOLEAN_VAL
	FUNCTION_VAL
)

type Value interface {
	Type() ValueType
	Inspect() string
}

type Integer struct {
	Value int
}

func (i Integer) Type() ValueType {
	return INTEGER_VAL
}

func (i Integer) Inspect() string {
	return fmt.Sprintf("%d", i.Value)
}

type Boolean struct {
	Value bool
}

func (b Boolean) Type() ValueType {
	return BOOLEAN_VAL
}

func (b Boolean) Inspect() string {
	return fmt.Sprintf("%t", b.Value)
}

type Function struct {
	Param ast.Identifier
	Body  ast.Expr
	Env   Environment
}

func (f Function) Type() ValueType {
	return FUNCTION_VAL
}

func (f Function) Inspect() string {
	var out bytes.Buffer

	out.WriteString("fun")
	out.WriteString("(")
	out.WriteString(f.Param.String())
	out.WriteString(", ")
	out.WriteString(f.Body.String())
	out.WriteString(")")

	return out.String()
}
