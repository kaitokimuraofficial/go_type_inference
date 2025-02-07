package eval_test

import (
	"go_type_inference/ast"
	"go_type_inference/eval"
	"go_type_inference/parser"
	"go_type_inference/token"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestEval(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name  string
		input string
		want  eval.Value
	}{
		{
			name:  "identifier",
			input: "x",
			want: eval.Integer{
				Value: 10,
			},
		},
		{
			name:  "integer",
			input: "2",
			want: eval.Integer{
				Value: 2,
			},
		},
		{
			name:  "boolean",
			input: "true",
			want: eval.Boolean{
				Value: true,
			},
		},
		{
			name:  "add two integers",
			input: "2 + 3",
			want: eval.Integer{
				Value: 5,
			},
		},
		{
			name:  "multiple two integers",
			input: "2 * 3",
			want: eval.Integer{
				Value: 6,
			},
		},
		{
			name:  "compare two integers",
			input: "2 < 3",
			want: eval.Boolean{
				Value: true,
			},
		},
		{
			name:  "add two identifiers",
			input: "i + x",
			want: eval.Integer{
				Value: 11,
			},
		},
		{
			name:  "parenthesized expression",
			input: "(2 + 3)",
			want: eval.Integer{
				Value: 5,
			},
		},
		{
			name:  "if true expression",
			input: "if true then 2 else 3",
			want: eval.Integer{
				Value: 2,
			},
		},
		{
			name:  "if else expression",
			input: "if false then 2 else 3",
			want: eval.Integer{
				Value: 3,
			},
		},
		{
			name:  "nested if true expression",
			input: "if (if false then true else false) then 2 else 3",
			want: eval.Integer{
				Value: 3,
			},
		},
		{
			name:  "nested if expression with identifiers",
			input: "if (if x < v then true else false) then 2 else i",
			want: eval.Integer{
				Value: 1,
			},
		},
		{
			name:  "let declaration",
			input: "let x = 2",
			want: eval.Integer{
				Value: 2,
			},
		},
		{
			name:  "let expression",
			input: "let x = 2 in x + 3",
			want: eval.Integer{
				Value: 5,
			},
		},
		{
			name:  "nested let expression",
			input: "let x = 2 in let y = 3 in x + y",
			want: eval.Integer{
				Value: 5,
			},
		},
		{
			name:  "function declaration",
			input: "fun y -> y + 3",
			want: eval.Function{
				Param: ast.Ident{
					Value: "y",
				},
				Body: ast.BinOpExpr{
					Op: token.PLUS,
					Left: ast.Ident{
						Value: "y",
					},
					Right: ast.Integer{
						Value: 3,
					},
				},
				Env: eval.Environment{
					Store: map[ast.Ident]eval.Value{
						{Value: "i"}: eval.Integer{Value: 1},
						{Value: "v"}: eval.Integer{Value: 5},
						{Value: "x"}: eval.Integer{Value: 10},
					},
				},
			},
		},
		{
			name:  "function application with identifier in environment",
			input: "(fun x -> x + 3 ) 2",
			want: eval.Integer{
				Value: 5,
			},
		},
		{
			name:  "function application with identifier not in environment",
			input: "(fun param -> param + 3 ) 2",
			want: eval.Integer{
				Value: 5,
			},
		},
		{
			name:  "nested function application given one argument",
			input: "(fun x -> (fun y -> x + y)) 2",
			want: eval.Function{
				Param: ast.Ident{
					Value: "y",
				},
				Body: ast.BinOpExpr{
					Op: token.PLUS,
					Left: ast.Ident{
						Value: "x",
					},
					Right: ast.Ident{
						Value: "y",
					},
				},
				Env: eval.Environment{
					Store: map[ast.Ident]eval.Value{
						{Value: "i"}: eval.Integer{Value: 1},
						{Value: "v"}: eval.Integer{Value: 5},
						{Value: "x"}: eval.Integer{Value: 2},
					},
				},
			},
		},
		{
			name:  "nested function application given two arguments",
			input: "(fun x -> (fun y -> x + y)) 2 3",
			want: eval.Integer{
				Value: 5,
			},
		},
		{
			name:  "recursive function declaration",
			input: "let rec f = fun n -> (if 10 < n then 1 else n * f (n + 1))",
			want: eval.Function{
				Param: ast.Ident{
					Value: "n",
				},
				Body: ast.IfExpr{
					Cond: ast.BinOpExpr{
						Op: token.LT,
						Left: ast.Integer{
							Value: 10,
						},
						Right: ast.Ident{
							Value: "n",
						},
					},
					Cons: ast.Integer{
						Value: 1,
					},
					Alt: ast.BinOpExpr{
						Op: token.ASTERISK,
						Left: ast.Ident{
							Value: "n",
						},
						Right: ast.AppExpr{
							Func: ast.Ident{
								Value: "f",
							},
							Arg: ast.BinOpExpr{
								Op: token.PLUS,
								Left: ast.Ident{
									Value: "n",
								},
								Right: ast.Integer{
									Value: 1,
								},
							},
						},
					},
				},
				Env: eval.Environment{
					Store: map[ast.Ident]eval.Value{},
				},
			},
		},
		{
			name:  "recursive function expression",
			input: "let rec fact = fun n -> (if 9 < n then 1 else n * (fact (n+1))) in fact 8",
			want: eval.Integer{
				Value: 72,
			},
		},
	}

	for _, tc := range testCases {
		tc := tc

		env := eval.Environment{Store: make(map[ast.Ident]eval.Value)}
		env.Set(ast.Ident{Value: "i"}, eval.Integer{Value: 1})
		env.Set(ast.Ident{Value: "v"}, eval.Integer{Value: 5})
		env.Set(ast.Ident{Value: "x"}, eval.Integer{Value: 10})

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			p := parser.Parse(tc.input)
			got := eval.Eval(p, env)

			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("Eval(%q) returned unexpected difference (-want +got):\n%s", tc.input, diff)
			}
		})
	}
}
