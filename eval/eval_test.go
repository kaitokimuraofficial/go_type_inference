package eval

import (
	"go_type_inference/ast"
	"go_type_inference/object"
	"go_type_inference/parser"
	"go_type_inference/token"
	"reflect"
	"testing"
)

func TestEval(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		input    string
		expected object.Object
	}{
		"identifier": {
			input: "x",
			expected: object.Integer{
				Value: 10,
			},
		},
		"integer": {
			input: "2",
			expected: object.Integer{
				Value: 2,
			},
		},
		"boolean": {
			input: "true",
			expected: object.Boolean{
				Value: true,
			},
		},
		"binary operator(PLUS)": {
			input: "2 + 3",
			expected: object.Integer{
				Value: 5,
			},
		},
		"binary operator(PLUS) identifier": {
			input: "i + x",
			expected: object.Integer{
				Value: 11,
			},
		},
		"binary operator(ASTERISK)": {
			input: "2 * 3",
			expected: object.Integer{
				Value: 6,
			},
		},
		"binary operator(LT)": {
			input: "2 < 3",
			expected: object.Boolean{
				Value: true,
			},
		},
		"if true": {
			input: "if true then 2 else 3",
			expected: object.Integer{
				Value: 2,
			},
		},
		"if else": {
			input: "if false then 2 else 3",
			expected: object.Integer{
				Value: 3,
			},
		},
		"parentheses integer": {
			input: "(2 + 3)",
			expected: object.Integer{
				Value: 5,
			},
		},
		"nested if true": {
			input: "if (if false then true else false) then 2 else 3",
			expected: object.Integer{
				Value: 3,
			},
		},
		"nested if identifier": {
			input: "if (if x < v then true else false) then 2 else i",
			expected: object.Integer{
				Value: 1,
			},
		},
		"let declaration": {
			input: "let x = 2",
			expected: object.Integer{
				Value: 2,
			},
		},
		"let in": {
			input: "let x = 2 in x + 3",
			expected: object.Integer{
				Value: 5,
			},
		},
		"nested let in": {
			input: "let x = 2 in let y = 3 in x + y",
			expected: object.Integer{
				Value: 5,
			},
		},
		"fun abstraction": {
			input: "fun y -> y + 3",
			expected: object.Function{
				Param: ast.Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "y"},
					Value: "y",
				},
				Body: ast.BinOpExpr{
					Token: token.Token{Type: token.PLUS, Literal: "+"},
					Left: ast.Identifier{
						Token: token.Token{Type: token.IDENT, Literal: "y"},
						Value: "y",
					},
					Operator: "+",
					Right: ast.Integer{
						Token: token.Token{Type: token.INT, Literal: "3"},
						Value: 3,
					},
				},
				Env: object.Environment{
					Store: map[string]object.Object{
						"i": object.Integer{Value: 1},
						"v": object.Integer{Value: 5},
						"x": object.Integer{Value: 10},
					},
				},
			},
		},
		"fun application-1": {
			input: "(fun x -> x + 3 ) 2",
			expected: object.Integer{
				Value: 5,
			},
		},
		"fun application-2": {
			input: "(fun param -> param + 3 ) 2",
			expected: object.Integer{
				Value: 5,
			},
		},
		"nested function application-1": {
			input: "(fun x -> (fun y -> x + y)) 2",
			expected: object.Function{
				Param: ast.Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "y"},
					Value: "y",
				},
				Body: ast.BinOpExpr{
					Token: token.Token{Type: token.PLUS, Literal: "+"},
					Left: ast.Identifier{
						Token: token.Token{Type: token.IDENT, Literal: "x"},
						Value: "x",
					},
					Operator: "+",
					Right: ast.Identifier{
						Token: token.Token{Type: token.IDENT, Literal: "y"},
						Value: "y",
					},
				},
				Env: object.Environment{
					Store: map[string]object.Object{
						"i": object.Integer{Value: 1},
						"v": object.Integer{Value: 5},
						"x": object.Integer{Value: 2},
					},
				},
			},
		},
		"nested function application-2": {
			input: "(fun x -> (fun y -> x + y)) 2 3",
			expected: object.Integer{
				Value: 5,
			},
		},
		"let rec expression": {
			input: "let rec fact = fun n -> (if 9 < n then 1 else n * (fact (n+1))) in fact 8",
			expected: object.Integer{
				Value: 72,
			},
		},
	}

	for name, tt := range tests {
		tt := tt

		env := object.NewEnvironment()
		env.Set("i", object.Integer{Value: 1})
		env.Set("v", object.Integer{Value: 5})
		env.Set("x", object.Integer{Value: 10})

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := Eval(parser.Parse(tt.input), env)

			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("expect: %q, but got: %q", tt.expected, got)
			}
		})
	}
}
