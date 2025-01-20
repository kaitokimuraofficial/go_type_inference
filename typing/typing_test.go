package typing_test

import (
	"go_type_inference/ast"
	"go_type_inference/parser"
	"go_type_inference/typing"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestInfer(t *testing.T) {
	testCases := []struct {
		name  string
		input string
		want  typing.Type
	}{
		{
			name:  "identifier",
			input: "i",
			want:  typing.TyInt{},
		},
		{
			name:  "integer",
			input: "2",
			want:  typing.TyInt{},
		},
		{
			name:  "boolean",
			input: "true",
			want:  typing.TyBool{},
		},
		{
			name:  "add two integers",
			input: "2 + 3",
			want:  typing.TyInt{},
		},
		{
			name:  "add two identifiers",
			input: "i + v",
			want:  typing.TyInt{},
		},
		{
			name:  "multiple two integers",
			input: "2 * 3",
			want:  typing.TyInt{},
		},
		{
			name:  "compare two integers",
			input: "2 < 3",
			want:  typing.TyBool{},
		},
		{
			name:  "parenthesized expression",
			input: "(2 + 3)",
			want:  typing.TyInt{},
		},
		{
			name:  "if true expression",
			input: "if true then 2 else 3",
			want:  typing.TyInt{},
		},
		{
			name:  "if else expression",
			input: "if false then 2 else 3",
			want:  typing.TyInt{},
		},
		{
			name:  "nested if expression",
			input: "if (if false then true else false) then 2 else 3",
			want:  typing.TyInt{},
		},
		{
			name:  "nested if expression with identifiers in environment",
			input: "if (if i < v then true else false) then 2 else i",
			want:  typing.TyInt{},
		},
		{
			name:  "let declaration",
			input: "let x = 2",
			want:  typing.TyInt{},
		},
		{
			name:  "let expression",
			input: "let x = 2 in x + 3",
			want:  typing.TyInt{},
		},
		{
			name:  "nested let expression",
			input: "let x = 2 in let y = 3 in x + y",
			want:  typing.TyInt{},
		},
		{
			name:  "let poly declaration",
			input: "let f = fun x -> x",
			want: typing.TyFun{
				Abs: typing.TyVar{Variable: 0},
				App: typing.TyVar{Variable: 0},
			},
		},
		{
			name:  "let poly expression",
			input: "let f = fun x -> x in if f true then f 2 else f 3",
			want:  typing.TyInt{},
		},
		{
			name:  "fun abstraction",
			input: "fun y -> y + 3",
			want: typing.TyFun{
				Abs: typing.TyInt{},
				App: typing.TyInt{},
			},
		},
		{
			name:  "fun application-1",
			input: "(fun x -> x + 3 ) 2",
			want:  typing.TyInt{},
		},
		{
			name:  "fun application-2",
			input: "(fun param -> param + 3 ) 2",
			want:  typing.TyInt{},
		},
		{
			name:  "nested function application-1",
			input: "(fun x -> (fun y -> x + y)) 2",
			want: typing.TyFun{
				Abs: typing.TyInt{},
				App: typing.TyInt{},
			},
		},
		{
			name:  "nested function application-2",
			input: "(fun x -> (fun y -> x + y)) 2 3",
			want:  typing.TyInt{},
		},
		{
			name:  "fun abstraction",
			input: "fun y -> y + 3",
			want: typing.TyFun{
				Abs: typing.TyInt{},
				App: typing.TyInt{},
			},
		},
		{
			name:  "fun application-1",
			input: "(fun x -> x + 3 ) 2",
			want:  typing.TyInt{},
		},
		{
			name:  "fun application-2",
			input: "(fun param -> param + 3 ) 2",
			want:  typing.TyInt{},
		},
		{
			name:  "nested function application-1",
			input: "(fun x -> (fun y -> x + y)) 2",
			want: typing.TyFun{
				Abs: typing.TyInt{},
				App: typing.TyInt{},
			},
		},
		{
			name:  "nested function application-2",
			input: "(fun x -> (fun y -> x + y)) 2 3",
			want:  typing.TyInt{},
		},
		// {
		// 	name:  "recursive function declaration",
		// 	input: "let rec double = fun f -> fun x -> f f x",
		// 	want: &TyFun{
		// 		Abs: &TyFun{
		// 			Abs: &TyVar{Variable: 1},
		// 			App: &TyVar{Variable: 1},
		// 		},
		// 		App: &TyFun{
		// 			Abs: &TyVar{Variable: 1},
		// 			App: &TyVar{Variable: 1},
		// 		},
		// 	},
		// },
		// {
		// 	name:  "recursive function expression",
		// 	input: "let rec double = fun f -> fun x -> f f x in double (fun x -> x + 1 ) 4",
		// 	want:  &TyInt{},
		// },
	}

	for _, tc := range testCases {
		tc := tc

		env := typing.Environment{Store: make(map[typing.Variable]typing.Type)}
		env.Set(ast.Identifier{Value: "b"}, typing.TyScheme{
			BoundVars: []typing.Variable{},
			Type:      typing.TyInt{},
		})
		env.Set(ast.Identifier{Value: "i"}, typing.TyScheme{
			BoundVars: []typing.Variable{},
			Type:      typing.TyInt{},
		})
		env.Set(ast.Identifier{Value: "v"}, typing.TyScheme{
			BoundVars: []typing.Variable{},
			Type:      typing.TyInt{},
		})

		t.Run(tc.name, func(t *testing.T) {
			p := parser.Parse(tc.input)
			_, got := typing.Infer(p, env)

			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("Infer(%q) returned unexpected difference (-want +got):\n%s", tc.input, diff)
			}
		})
	}
}
