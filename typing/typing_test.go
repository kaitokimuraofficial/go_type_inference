package typing

import (
	"go_type_inference/ast"
	"go_type_inference/object"
	"go_type_inference/parser"
	"go_type_inference/token"
	"reflect"
	"testing"
)

func TestInfer(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		input    string
		expected object.InferredObject
	}{
		"identifier": {
			input:    "i",
			expected: object.TyInt{},
		},
		"integer": {
			input:    "2",
			expected: object.TyInt{},
		},
		"boolean": {
			input:    "true",
			expected: object.TyBool{},
		},
		"binary operator(PLUS)": {
			input:    "2 + 3",
			expected: object.TyInt{},
		},
		"binary operator(PLUS) identifier": {
			input:    "i + v",
			expected: object.TyInt{},
		},
		"binary operator(ASTERISK)": {
			input:    "2 * 3",
			expected: object.TyInt{},
		},
		"binary operator(LT)": {
			input:    "2 < 3",
			expected: object.TyBool{},
		},
		"if true": {
			input:    "if true then 2 else 3",
			expected: object.TyInt{},
		},
		"if else": {
			input:    "if false then 2 else 3",
			expected: object.TyInt{},
		},
		"parentheses integer": {
			input:    "(2 + 3)",
			expected: object.TyInt{},
		},
		"nested if true": {
			input:    "if (if false then true else false) then 2 else 3",
			expected: object.TyInt{},
		},
		"nested if identifier": {
			input:    "if (if i < v then true else false) then 2 else i",
			expected: object.TyInt{},
		},
		"let declaration": {
			input:    "let x = 2",
			expected: object.TyInt{},
		},
		"let in": {
			input:    "let x = 2 in x + 3",
			expected: object.TyInt{},
		},
		"nested let in": {
			input:    "let x = 2 in let y = 3 in x + y",
			expected: object.TyInt{},
		},
	}

	for name, tt := range tests {
		tt := tt

		env := object.NewTypeEnvironment()
		env.Set(ast.Identifier{Token: token.Token{Type: token.IDENT, Literal: "b"}, Value: "b"}, object.TyBool{})
		env.Set(ast.Identifier{Token: token.Token{Type: token.IDENT, Literal: "i"}, Value: "i"}, object.TyInt{})
		env.Set(ast.Identifier{Token: token.Token{Type: token.IDENT, Literal: "v"}, Value: "v"}, object.TyInt{})

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := Infer(parser.Parse(tt.input), env)

			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("expect: %s, but got: %s", tt.expected, got)
			}
		})
	}
}
