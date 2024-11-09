package parser

import (
	"go_type_inference/ast"
	"go_type_inference/token"
	"testing"
)

func TestParse(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		input    string
		expected ast.Stmt
	}{
		"identifier": {
			input: "x",
			expected: ast.Statement{
				Expr: ast.Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "x"},
					Value: "x",
				},
			},
		},
		"integer": {
			input: "4",
			expected: ast.Statement{
				Expr: ast.Integer{
					Token: token.Token{Type: token.INT, Literal: "4"},
					Value: 4,
				},
			},
		},
		"boolean": {
			input: "true",
			expected: ast.Statement{
				Expr: ast.Boolean{
					Token: token.Token{Type: token.TRUE, Literal: "true"},
					Value: true,
				},
			},
		},
		"binary operator": {
			input: "2 + 3",
			expected: ast.Statement{
				Expr: ast.BinOpExpr{
					Token: token.Token{Type: token.PLUS, Literal: "+"},
					Left: ast.Integer{
						Token: token.Token{Type: token.INT, Literal: "2"},
						Value: 2,
					},
					Operator: "+",
					Right: ast.Integer{
						Token: token.Token{Type: token.INT, Literal: "3"},
						Value: 3,
					},
				},
			},
		},
		"if": {
			input: "if true then true else false",
			expected: ast.Statement{
				Expr: ast.IfExpr{
					Token: token.Token{Type: token.IF, Literal: "if"},
					Condition:   ast.Boolean{
						Token: token.Token{Type: token.TRUE, Literal: "true"},
						Value: true,
					},
					Consequence: ast.Boolean{
						Token: token.Token{Type: token.TRUE, Literal: "true"},
						Value: true,
					},
					Alternative: ast.Boolean{
						Token: token.Token{Type: token.FALSE, Literal: "false"},
						Value: false,
					},
				},
			},
		},
		"parentheses": {
			input: "(true)",
			expected: ast.Statement{
				Expr: ast.Boolean{
					Token: token.Token{Type: token.TRUE, Literal: "true"},
					Value: true,
				},
			},
		},
	}

	for name, tt := range tests {
		tt := tt
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := Parse(tt.input)

			if got != tt.expected {
				t.Errorf("expect: %s, but got: %s", tt.expected, got)
			}
		})
	}
}
