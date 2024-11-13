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

	testCases := []struct {
		name  string
		input string
		want  object.Value
	}{
		{
			name:  "identifier",
			input: "x",
			want: object.Integer{
				Value: 10,
			},
		},
		{
			name:  "integer",
			input: "2",
			want: object.Integer{
				Value: 2,
			},
		},
		{
			name:  "boolean",
			input: "true",
			want: object.Boolean{
				Value: true,
			},
		},
		{
			name:  "binary operator(PLUS)",
			input: "2 + 3",
			want: object.Integer{
				Value: 5,
			},
		},
		{
			name:  "binary operator(PLUS) identifier",
			input: "i + x",
			want: object.Integer{
				Value: 11,
			},
		},
		{
			name:  "binary operator(ASTERISK)",
			input: "2 * 3",
			want: object.Integer{
				Value: 6,
			},
		},
		{
			name:  "binary operator(LT)",
			input: "2 < 3",
			want: object.Boolean{
				Value: true,
			},
		},
		{
			name:  "if true",
			input: "if true then 2 else 3",
			want: object.Integer{
				Value: 2,
			},
		},
		{
			name:  "if else",
			input: "if false then 2 else 3",
			want: object.Integer{
				Value: 3,
			},
		},
		{
			name:  "parentheses integer",
			input: "(2 + 3)",
			want: object.Integer{
				Value: 5,
			},
		},
		{
			name:  "nested if true",
			input: "if (if false then true else false) then 2 else 3",
			want: object.Integer{
				Value: 3,
			},
		},
		{
			name:  "nested if identifier",
			input: "if (if x < v then true else false) then 2 else i",
			want: object.Integer{
				Value: 1,
			},
		},
		{
			name:  "let declaration",
			input: "let x = 2",
			want: object.Integer{
				Value: 2,
			},
		},
		{
			name:  "let in",
			input: "let x = 2 in x + 3",
			want: object.Integer{
				Value: 5,
			},
		},
		{
			name:  "nested let in",
			input: "let x = 2 in let y = 3 in x + y",
			want: object.Integer{
				Value: 5,
			},
		},
		{
			name:  "fun abstraction",
			input: "fun y -> y + 3",
			want: object.Function{
				Param: ast.Identifier{
					Value: "y",
				},
				Body: ast.BinOpExpr{
					Type: token.PLUS,
					Left: ast.Identifier{
						Value: "y",
					},
					Right: ast.Integer{
						Value: 3,
					},
				},
				Env: object.Environment{
					Store: map[ast.Identifier]object.Value{
						{
							Value: "i",
						}: object.Integer{Value: 1},
						{
							Value: "v",
						}: object.Integer{Value: 5},
						{
							Value: "x",
						}: object.Integer{Value: 10},
					},
				},
			},
		},
		{
			name:  "fun application-1",
			input: "(fun x -> x + 3 ) 2",
			want: object.Integer{
				Value: 5,
			},
		},
		{
			name:  "fun application-2",
			input: "(fun param -> param + 3 ) 2",
			want: object.Integer{
				Value: 5,
			},
		},
		{
			name:  "nested function application-1",
			input: "(fun x -> (fun y -> x + y)) 2",
			want: object.Function{
				Param: ast.Identifier{
					Value: "y",
				},
				Body: ast.BinOpExpr{
					Type: token.PLUS,
					Left: ast.Identifier{
						Value: "x",
					},
					Right: ast.Identifier{
						Value: "y",
					},
				},
				Env: object.Environment{
					Store: map[ast.Identifier]object.Value{
						{
							Value: "i",
						}: object.Integer{Value: 1},
						{
							Value: "v",
						}: object.Integer{Value: 5},
						{
							Value: "x",
						}: object.Integer{Value: 2},
					},
				},
			},
		},
		{
			name:  "nested function application-2",
			input: "(fun x -> (fun y -> x + y)) 2 3",
			want: object.Integer{
				Value: 5,
			},
		},
		{
			name:  "let rec declaration",
			input: "let rec f = fun n -> (if 10 < n then 1 else n * f (n + 1))",
			want: object.Function{
				Param: ast.Identifier{
					Value: "n",
				},
				Body: ast.IfExpr{
					Condition: ast.BinOpExpr{
						Type: token.LT,
						Left: ast.Integer{
							Value: 10,
						},
						Right: ast.Identifier{
							Value: "n",
						},
					},
					Consequence: ast.Integer{
						Value: 1,
					},
					Alternative: ast.BinOpExpr{
						Type: token.ASTERISK,
						Left: ast.Identifier{
							Value: "n",
						},
						Right: ast.AppExpr{
							Function: ast.Identifier{
								Value: "f",
							},
							Argument: ast.BinOpExpr{
								Type: token.PLUS,
								Left: ast.Identifier{
									Value: "n",
								},
								Right: ast.Integer{
									Value: 1,
								},
							},
						},
					},
				},
				Env: object.Environment{
					Store: map[ast.Identifier]object.Value{},
				},
			},
		},
		{
			name:  "let rec expression",
			input: "let rec fact = fun n -> (if 9 < n then 1 else n * (fact (n+1))) in fact 8",
			want: object.Integer{
				Value: 72,
			},
		},
	}

	for _, tc := range testCases {
		tc := tc

		env := object.NewEnvironment()
		env.Set(ast.Identifier{Value: "i"}, object.Integer{Value: 1})
		env.Set(ast.Identifier{Value: "v"}, object.Integer{Value: 5})
		env.Set(ast.Identifier{Value: "x"}, object.Integer{Value: 10})

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			got := Eval(parser.Parse(tc.input), env)

			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("expect: %q, but got: %q", tc.want, got)
			}
		})
	}
}
