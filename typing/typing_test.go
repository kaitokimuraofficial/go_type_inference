package typing

import (
	"go_type_inference/ast"
	"go_type_inference/parser"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestInfer(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name  string
		input string
		want  Type
	}{
		{
			name:  "identifier",
			input: "i",
			want:  &TyInt{},
		},
		{
			name:  "integer",
			input: "2",
			want:  &TyInt{},
		},
		{
			name:  "boolean",
			input: "true",
			want:  &TyBool{},
		},
		{
			name:  "add two integers",
			input: "2 + 3",
			want:  &TyInt{},
		},
		{
			name:  "add two identifiers",
			input: "i + v",
			want:  &TyInt{},
		},
		{
			name:  "multiple two integers",
			input: "2 * 3",
			want:  &TyInt{},
		},
		{
			name:  "compare two integers",
			input: "2 < 3",
			want:  &TyBool{},
		},
		{
			name:  "parenthesized expression",
			input: "(2 + 3)",
			want:  &TyInt{},
		},
		{
			name:  "if true expression",
			input: "if true then 2 else 3",
			want:  &TyInt{},
		},
		{
			name:  "if else expression",
			input: "if false then 2 else 3",
			want:  &TyInt{},
		},
		{
			name:  "nested if expression",
			input: "if (if false then true else false) then 2 else 3",
			want:  &TyInt{},
		},
		{
			name:  "nested if expression with identifiers in environment",
			input: "if (if i < v then true else false) then 2 else i",
			want:  &TyInt{},
		},
		{
			name:  "let declaration",
			input: "let x = 2",
			want:  &TyInt{},
		},
		{
			name:  "let expression",
			input: "let x = 2 in x + 3",
			want:  &TyInt{},
		},
		{
			name:  "nested let expression",
			input: "let x = 2 in let y = 3 in x + y",
			want:  &TyInt{},
		},
		{
			name:  "fun abstraction",
			input: "fun y -> y + 3",
			want: &TyFun{
				Abs: &TyInt{},
				App: &TyInt{},
			},
		},
		{
			name:  "fun application-1",
			input: "(fun x -> x + 3 ) 2",
			want:  &TyInt{},
		},
		{
			name:  "fun application-2",
			input: "(fun param -> param + 3 ) 2",
			want:  &TyInt{},
		},
		{
			name:  "nested function application-1",
			input: "(fun x -> (fun y -> x + y)) 2",
			want: &TyFun{
				Abs: &TyInt{},
				App: &TyInt{},
			},
		},
		{
			name:  "nested function application-2",
			input: "(fun x -> (fun y -> x + y)) 2 3",
			want:  &TyInt{},
		},
		{
			name:  "fun abstraction",
			input: "fun y -> y + 3",
			want: &TyFun{
				Abs: &TyInt{},
				App: &TyInt{},
			},
		},
		{
			name:  "fun application-1",
			input: "(fun x -> x + 3 ) 2",
			want:  &TyInt{},
		},
		{
			name:  "fun application-2",
			input: "(fun param -> param + 3 ) 2",
			want:  &TyInt{},
		},
		{
			name:  "nested function application-1",
			input: "(fun x -> (fun y -> x + y)) 2",
			want: &TyFun{
				Abs: &TyInt{},
				App: &TyInt{},
			},
		},
		{
			name:  "nested function application-2",
			input: "(fun x -> (fun y -> x + y)) 2 3",
			want:  &TyInt{},
		},
		{
			name:  "recursive function declaration",
			input: "let rec f = fun n -> (if 10 < n then 1 else n * f (n + 1))",
			want: &TyFun{
				Abs: &TyInt{},
				App: &TyInt{},
			},
		},
		{
			name:  "recursive function expression",
			input: "let rec fact = fun n -> (if 9 < n then 1 else n * (fact (n+1))) in fact 8",
			want:  &TyInt{},
		},
	}

	for _, tc := range testCases {
		tc := tc

		env := &Environment{Store: make(map[ast.Identifier]Type)}
		env.Set(ast.Identifier{Value: "b"}, &TyBool{})
		env.Set(ast.Identifier{Value: "i"}, &TyInt{})
		env.Set(ast.Identifier{Value: "v"}, &TyInt{})

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			p := parser.Parse(tc.input)
			_, got := Infer(p, env)

			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("Infer(%q) returned unexpected difference (-want +got):\n%s", tc.input, diff)
			}
		})
	}
}
