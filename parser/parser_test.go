package parser

import (
	"go_type_inference/ast"
	"go_type_inference/token"
	"testing"
)

func TestParse(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name  string
		input string
		want  ast.Stmt
	}{
		{
			name:  "identifier",
			input: "x",
			want: ast.Statement{
				Expr: ast.Identifier{
					Value: "x",
				},
			},
		},
		{
			name:  "integer",
			input: "4",
			want: ast.Statement{
				Expr: ast.Integer{
					Value: 4,
				},
			},
		},
		{
			name:  "boolean",
			input: "true",
			want: ast.Statement{
				Expr: ast.Boolean{
					Value: true,
				},
			},
		},
		{
			name:  "binary operator",
			input: "2 + 3",
			want: ast.Statement{
				Expr: ast.BinOpExpr{
					Type: token.PLUS,
					Left: ast.Integer{
						Value: 2,
					},
					Right: ast.Integer{
						Value: 3,
					},
				},
			},
		},
		{
			name:  "if",
			input: "if true then true else false",
			want: ast.Statement{
				Expr: ast.IfExpr{
					Condition: ast.Boolean{
						Value: true,
					},
					Consequence: ast.Boolean{
						Value: true,
					},
					Alternative: ast.Boolean{
						Value: false,
					},
				},
			},
		},
		{
			name:  "parentheses",
			input: "(true)",
			want: ast.Statement{
				Expr: ast.Boolean{
					Value: true,
				},
			},
		},
		{
			name:  "let declaration",
			input: "let x = 5",
			want: ast.Declaration{
				Expr: ast.Integer{
					Value: 5,
				},
				Id: ast.Identifier{
					Value: "x",
				},
			},
		},
		{
			name:  "let in",
			input: "let x = 2 in x + 3",
			want: ast.Statement{
				Expr: ast.LetExpr{
					Id: ast.Identifier{
						Value: "x",
					},
					BindingExpr: ast.Integer{
						Value: 2,
					},
					BodyExpr: ast.BinOpExpr{
						Type: token.PLUS,
						Left: ast.Identifier{
							Value: "x",
						},
						Right: ast.Integer{

							Value: 3,
						},
					},
				},
			},
		},
		{
			name:  "fun abstraction",
			input: "fun x -> x + 3",
			want: ast.Statement{
				Expr: ast.FunExpr{
					Param: ast.Identifier{
						Value: "x",
					},
					BodyExpr: ast.BinOpExpr{
						Type: token.PLUS,
						Left: ast.Identifier{
							Value: "x",
						},
						Right: ast.Integer{
							Value: 3,
						},
					},
				},
			},
		},
		{
			name:  "function application",
			input: "(fun x -> x + 3 ) 2",
			want: ast.Statement{
				Expr: ast.AppExpr{
					Function: ast.FunExpr{
						Param: ast.Identifier{
							Value: "x",
						},
						BodyExpr: ast.BinOpExpr{
							Type: token.PLUS,
							Left: ast.Identifier{
								Value: "x",
							},
							Right: ast.Integer{
								Value: 3,
							},
						},
					},
					Argument: ast.Integer{
						Value: 2,
					},
				},
			},
		},
		{
			name:  "nested function abstraction",
			input: "fun x -> (fun y -> x + y)",
			want: ast.Statement{
				Expr: ast.FunExpr{
					Param: ast.Identifier{
						Value: "x",
					},
					BodyExpr: ast.FunExpr{
						Param: ast.Identifier{
							Value: "y",
						},
						BodyExpr: ast.BinOpExpr{
							Type: token.PLUS,
							Left: ast.Identifier{
								Value: "x",
							},
							Right: ast.Identifier{
								Value: "y",
							},
						},
					},
				},
			},
		},
		{
			name:  "nested function application",
			input: "(fun x -> (fun y -> x + y)) 2",
			want: ast.Statement{
				Expr: ast.AppExpr{
					Function: ast.FunExpr{
						Param: ast.Identifier{
							Value: "x",
						},
						BodyExpr: ast.FunExpr{
							Param: ast.Identifier{
								Value: "y",
							},
							BodyExpr: ast.BinOpExpr{
								Type: token.PLUS,
								Left: ast.Identifier{
									Value: "x",
								},
								Right: ast.Identifier{
									Value: "y",
								},
							},
						},
					},
					Argument: ast.Integer{
						Value: 2,
					},
				},
			},
		},
		{
			name:  "let rec declaration",
			input: "let rec f = fun n -> if n < 10 then 1 else n * f (n + 1)",
			want: ast.RecDeclaration{
				Id: ast.Identifier{
					Value: "f",
				},
				Param: ast.Identifier{
					Value: "n",
				},
				BodyExpr: ast.IfExpr{
					Condition: ast.BinOpExpr{
						Type: token.LT,
						Left: ast.Identifier{
							Value: "n",
						},
						Right: ast.Integer{
							Value: 10,
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
			},
		},
		{
			name:  "let rec in expression",
			input: "let rec fact = fun n -> if n < 1 then 1 else n * (fact (n + 1)) in fact 5",
			want: ast.Statement{
				Expr: ast.LetRecExpr{
					Id: ast.Identifier{
						Value: "fact",
					},
					Param: ast.Identifier{
						Value: "n",
					},
					BindingExpr: ast.IfExpr{
						Condition: ast.BinOpExpr{
							Type: token.LT,
							Left: ast.Identifier{
								Value: "n",
							},
							Right: ast.Integer{
								Value: 1,
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
									Value: "fact",
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
					BodyExpr: ast.AppExpr{
						Function: ast.Identifier{
							Value: "fact",
						},
						Argument: ast.Integer{
							Value: 5,
						},
					},
				},
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			got := Parse(tc.input)

			if got != tc.want {
				t.Errorf("expect: %s, but got: %s", tc.want, got)
			}
		})
	}
}
