package eval

import (
	"go_type_inference/ast"
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
		want  Value
	}{
		{
			name:  "identifier",
			input: "x",
			want: &Integer{
				Value: 10,
			},
		},
		{
			name:  "integer",
			input: "2",
			want: &Integer{
				Value: 2,
			},
		},
		{
			name:  "boolean",
			input: "true",
			want: &Boolean{
				Value: true,
			},
		},
		{
			name:  "add two integers",
			input: "2 + 3",
			want: &Integer{
				Value: 5,
			},
		},
		{
			name:  "multiple two integers",
			input: "2 * 3",
			want: &Integer{
				Value: 6,
			},
		},
		{
			name:  "compare two integers",
			input: "2 < 3",
			want: &Boolean{
				Value: true,
			},
		},
		{
			name:  "add two identifiers",
			input: "i + x",
			want: &Integer{
				Value: 11,
			},
		},
		{
			name:  "parenthesized expression",
			input: "(2 + 3)",
			want: &Integer{
				Value: 5,
			},
		},
		{
			name:  "if true expression",
			input: "if true then 2 else 3",
			want: &Integer{
				Value: 2,
			},
		},
		{
			name:  "if else expression",
			input: "if false then 2 else 3",
			want: &Integer{
				Value: 3,
			},
		},
		{
			name:  "nested if true expression",
			input: "if (if false then true else false) then 2 else 3",
			want: &Integer{
				Value: 3,
			},
		},
		{
			name:  "nested if expression with identifiers",
			input: "if (if x < v then true else false) then 2 else i",
			want: &Integer{
				Value: 1,
			},
		},
		{
			name:  "let declaration",
			input: "let x = 2",
			want: &Integer{
				Value: 2,
			},
		},
		{
			name:  "let expression",
			input: "let x = 2 in x + 3",
			want: &Integer{
				Value: 5,
			},
		},
		{
			name:  "nested let expression",
			input: "let x = 2 in let y = 3 in x + y",
			want: &Integer{
				Value: 5,
			},
		},
		{
			name:  "function declaration",
			input: "fun y -> y + 3",
			want: &Function{
				Param: ast.Identifier{
					Value: "y",
				},
				Body: &ast.BinOpExpr{
					Op: token.PLUS,
					Left: &ast.Identifier{
						Value: "y",
					},
					Right: &ast.Integer{
						Value: 3,
					},
				},
				Env: Environment{
					Store: map[ast.Identifier]Value{
						{Value: "i"}: &Integer{Value: 1},
						{Value: "v"}: &Integer{Value: 5},
						{Value: "x"}: &Integer{Value: 10},
					},
				},
			},
		},
		{
			name:  "function application with identifier in environment",
			input: "(fun x -> x + 3 ) 2",
			want: &Integer{
				Value: 5,
			},
		},
		{
			name:  "function application with identifier not in environment",
			input: "(fun param -> param + 3 ) 2",
			want: &Integer{
				Value: 5,
			},
		},
		{
			name:  "nested function application given one argument",
			input: "(fun x -> (fun y -> x + y)) 2",
			want: &Function{
				Param: ast.Identifier{
					Value: "y",
				},
				Body: &ast.BinOpExpr{
					Op: token.PLUS,
					Left: &ast.Identifier{
						Value: "x",
					},
					Right: &ast.Identifier{
						Value: "y",
					},
				},
				Env: Environment{
					Store: map[ast.Identifier]Value{
						{Value: "i"}: &Integer{Value: 1},
						{Value: "v"}: &Integer{Value: 5},
						{Value: "x"}: &Integer{Value: 2},
					},
				},
			},
		},
		{
			name:  "nested function application given two arguments",
			input: "(fun x -> (fun y -> x + y)) 2 3",
			want: &Integer{
				Value: 5,
			},
		},
		{
			name:  "recursive function declaration",
			input: "let rec f = fun n -> (if 10 < n then 1 else n * f (n + 1))",
			want: &Function{
				Param: ast.Identifier{
					Value: "n",
				},
				Body: &ast.IfExpr{
					Condition: &ast.BinOpExpr{
						Op: token.LT,
						Left: &ast.Integer{
							Value: 10,
						},
						Right: &ast.Identifier{
							Value: "n",
						},
					},
					Consequence: &ast.Integer{
						Value: 1,
					},
					Alternative: &ast.BinOpExpr{
						Op: token.ASTERISK,
						Left: &ast.Identifier{
							Value: "n",
						},
						Right: &ast.AppExpr{
							Function: &ast.Identifier{
								Value: "f",
							},
							Argument: &ast.BinOpExpr{
								Op: token.PLUS,
								Left: &ast.Identifier{
									Value: "n",
								},
								Right: &ast.Integer{
									Value: 1,
								},
							},
						},
					},
				},
				Env: Environment{
					Store: map[ast.Identifier]Value{},
				},
			},
		},
		{
			name:  "recursive function expression",
			input: "let rec fact = fun n -> (if 9 < n then 1 else n * (fact (n+1))) in fact 8",
			want: &Integer{
				Value: 72,
			},
		},
	}

	for _, tc := range testCases {
		tc := tc

		env := &Environment{Store: make(map[ast.Identifier]Value)}
		env.Set(ast.Identifier{Value: "i"}, &Integer{Value: 1})
		env.Set(ast.Identifier{Value: "v"}, &Integer{Value: 5})
		env.Set(ast.Identifier{Value: "x"}, &Integer{Value: 10})

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			p := parser.Parse(tc.input)
			got := Eval(p, env)

			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("Eval(%q) returned unexpected difference (-want +got):\n%s", tc.input, diff)
			}
		})
	}
}
