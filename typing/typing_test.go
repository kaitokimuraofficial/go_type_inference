package typing

import (
	"go_type_inference/ast"
	"go_type_inference/object"
	"go_type_inference/parser"
	"reflect"
	"testing"
)

func TestInfer(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name  string
		input string
		want  object.InferredObject
	}{
		{
			name:  "identifier",
			input: "i",
			want:  object.TyInt{},
		},
		{
			name:  "integer",
			input: "2",
			want:  object.TyInt{},
		},
		{
			name:  "boolean",
			input: "true",
			want:  object.TyBool{},
		},
		{
			name:  "binary operator(PLUS)",
			input: "2 + 3",
			want:  object.TyInt{},
		},
		{
			name:  "binary operator(PLUS) identifier",
			input: "i + v",
			want:  object.TyInt{},
		},
		{
			name:  "binary operator(ASTERISK)",
			input: "2 * 3",
			want:  object.TyInt{},
		},
		{
			name:  "binary operator(LT)",
			input: "2 < 3",
			want:  object.TyBool{},
		},
		{
			name:  "if true",
			input: "if true then 2 else 3",
			want:  object.TyInt{},
		},
		{
			name:  "if else",
			input: "if false then 2 else 3",
			want:  object.TyInt{},
		},
		{
			name:  "parentheses integer",
			input: "(2 + 3)",
			want:  object.TyInt{},
		},
		{
			name:  "nested if true",
			input: "if (if false then true else false) then 2 else 3",
			want:  object.TyInt{},
		},
		{
			name:  "nested if identifier",
			input: "if (if i < v then true else false) then 2 else i",
			want:  object.TyInt{},
		},
		{
			name:  "let declaration",
			input: "let x = 2",
			want:  object.TyInt{},
		},
		{
			name:  "let in",
			input: "let x = 2 in x + 3",
			want:  object.TyInt{},
		},
		{
			name:  "nested let in",
			input: "let x = 2 in let y = 3 in x + y",
			want:  object.TyInt{},
		},
	}

	for _, tc := range testCases {
		tc := tc

		env := object.NewTypeEnvironment()
		env.Set(ast.Identifier{Value: "b"}, object.TyBool{})
		env.Set(ast.Identifier{Value: "i"}, object.TyInt{})
		env.Set(ast.Identifier{Value: "v"}, object.TyInt{})

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			got := Infer(parser.Parse(tc.input), env)

			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("expect: %s, but got: %s", tc.want, got)
			}
		})
	}
}
