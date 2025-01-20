package parser_test

import (
	"go_type_inference/ast"
	"go_type_inference/parser"
	"go_type_inference/token"
	"testing"

	"github.com/google/go-cmp/cmp"
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
			want: ast.ExprStmt{
				Expr: ast.Identifier{
					Value: "x",
				},
			},
		},
		{
			name:  "integer",
			input: "4",
			want: ast.ExprStmt{
				Expr: ast.Integer{
					Value: 4,
				},
			},
		},
		{
			name:  "boolean",
			input: "true",
			want: ast.ExprStmt{
				Expr: ast.Boolean{
					Value: true,
				},
			},
		},
		{
			name:  "primitive operator",
			input: "2 + 3",
			want: ast.ExprStmt{
				Expr: ast.BinOpExpr{
					Op: token.PLUS,
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
			name:  "if expression",
			input: "if true then true else false",
			want: ast.ExprStmt{
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
			name:  "parenthesized expression",
			input: "(true)",
			want: ast.ExprStmt{
				Expr: ast.Boolean{
					Value: true,
				},
			},
		},
		{
			name:  "let declaration",
			input: "let x = 5",
			want: ast.DeclStmt{
				Decl: ast.LetDecl{
					Id: ast.Identifier{
						Value: "x",
					},
					Expr: ast.Integer{
						Value: 5,
					},
				},
			},
		},
		{
			name:  "let expression",
			input: "let x = 2 in x + 3",
			want: ast.ExprStmt{
				Expr: ast.LetExpr{
					Id: ast.Identifier{
						Value: "x",
					},
					BindingExpr: ast.Integer{
						Value: 2,
					},
					BodyExpr: ast.BinOpExpr{
						Op: token.PLUS,
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
			name:  "function declaration",
			input: "fun x -> x + 3",
			want: ast.ExprStmt{
				Expr: ast.FunExpr{
					Param: ast.Identifier{
						Value: "x",
					},
					BodyExpr: ast.BinOpExpr{
						Op: token.PLUS,
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
			want: ast.ExprStmt{
				Expr: ast.AppExpr{
					Function: ast.FunExpr{
						Param: ast.Identifier{
							Value: "x",
						},
						BodyExpr: ast.BinOpExpr{
							Op: token.PLUS,
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
			name:  "nested function declaration",
			input: "fun x -> (fun y -> x + y)",
			want: ast.ExprStmt{
				Expr: ast.FunExpr{
					Param: ast.Identifier{
						Value: "x",
					},
					BodyExpr: ast.FunExpr{
						Param: ast.Identifier{
							Value: "y",
						},
						BodyExpr: ast.BinOpExpr{
							Op: token.PLUS,
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
			want: ast.ExprStmt{
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
								Op: token.PLUS,
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
			name:  "recursive function declaration",
			input: "let rec f = fun n -> if n < 10 then 1 else n * f (n + 1)",
			want: ast.DeclStmt{
				Decl: ast.RecDecl{
					Id: ast.Identifier{
						Value: "f",
					},
					Param: ast.Identifier{
						Value: "n",
					},
					BodyExpr: ast.IfExpr{
						Condition: ast.BinOpExpr{
							Op: token.LT,
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
							Op: token.ASTERISK,
							Left: ast.Identifier{
								Value: "n",
							},
							Right: ast.AppExpr{
								Function: ast.Identifier{
									Value: "f",
								},
								Argument: ast.BinOpExpr{
									Op: token.PLUS,
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
		},
		{
			name:  "recursive function expression",
			input: "let rec fact = fun n -> if n < 1 then 1 else n * (fact (n + 1)) in fact 5",
			want: ast.ExprStmt{
				Expr: ast.LetRecExpr{
					Id: ast.Identifier{
						Value: "fact",
					},
					Param: ast.Identifier{
						Value: "n",
					},
					BindingExpr: ast.IfExpr{
						Condition: ast.BinOpExpr{
							Op: token.LT,
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
							Op: token.ASTERISK,
							Left: ast.Identifier{
								Value: "n",
							},
							Right: ast.AppExpr{
								Function: ast.Identifier{
									Value: "fact",
								},
								Argument: ast.BinOpExpr{
									Op: token.PLUS,
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

			got := parser.Parse(tc.input)

			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("Parse(%q) returned unexpected difference (-want +got):\n%s", tc.input, diff)
			}
		})
	}
}
