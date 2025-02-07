package eval_test

import (
	"go_type_inference/ast"
	"go_type_inference/eval"
	"go_type_inference/token"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestEval(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name  string
		input ast.Node
		want  eval.Value
	}{
		{
			name:  "x",
			input: ast.Ident{Value: "x"},
			want:  eval.Integer{Value: 10},
		},
		{
			name:  "2",
			input: ast.Integer{Value: 2},
			want:  eval.Integer{Value: 2},
		},
		{
			name:  "true",
			input: ast.Boolean{Value: true},
			want:  eval.Boolean{Value: true},
		},
		{
			name: "2 + 3",
			input: ast.BinOpExpr{
				Op:    token.PLUS,
				Left:  ast.Integer{Value: 2},
				Right: ast.Integer{Value: 3},
			},
			want: eval.Integer{Value: 5},
		},
		{
			name: "2 * 3",
			input: ast.BinOpExpr{
				Op:    token.ASTERISK,
				Left:  ast.Integer{Value: 2},
				Right: ast.Integer{Value: 3},
			},
			want: eval.Integer{Value: 6},
		},
		{
			name: "2 < 3",
			input: ast.BinOpExpr{
				Op:    token.LT,
				Left:  ast.Integer{Value: 2},
				Right: ast.Integer{Value: 3},
			},
			want: eval.Boolean{Value: true},
		},
		{
			name: "i + x",
			input: ast.BinOpExpr{
				Op:    token.PLUS,
				Left:  ast.Ident{Value: "i"},
				Right: ast.Ident{Value: "x"},
			},
			want: eval.Integer{Value: 11},
		},
		{
			name: "(2 + 3)",
			input: ast.ExprStmt{
				Expr: ast.BinOpExpr{
					Op:    token.PLUS,
					Left:  ast.Integer{Value: 2},
					Right: ast.Integer{Value: 3},
				},
			},
			want: eval.Integer{Value: 5},
		},
		{
			name: "if true then 2 else 3",
			input: ast.ExprStmt{
				Expr: ast.IfExpr{
					Cond: ast.Boolean{Value: true},
					Cons: ast.Integer{Value: 2},
					Alt:  ast.Integer{Value: 3},
				},
			},
			want: eval.Integer{Value: 2},
		},
		{
			name: "if false then 2 else 3",
			input: ast.ExprStmt{
				Expr: ast.IfExpr{
					Cond: ast.Boolean{Value: false},
					Cons: ast.Integer{Value: 2},
					Alt:  ast.Integer{Value: 3},
				},
			},
			want: eval.Integer{Value: 3},
		},
		{
			name: "if (if false then true else false) then 2 else 3",
			input: ast.ExprStmt{
				Expr: ast.IfExpr{
					Cond: ast.IfExpr{
						Cond: ast.Boolean{Value: false},
						Cons: ast.Boolean{Value: true},
						Alt:  ast.Boolean{Value: false},
					},
					Cons: ast.Integer{Value: 2},
					Alt:  ast.Integer{Value: 3},
				},
			},
			want: eval.Integer{Value: 3},
		},
		{
			name: "if (if x < v then true else false) then 2 else i",
			input: ast.ExprStmt{
				Expr: ast.IfExpr{
					Cond: ast.IfExpr{
						Cond: ast.BinOpExpr{
							Op:    token.LT,
							Left:  ast.Ident{Value: "x"},
							Right: ast.Ident{Value: "v"},
						},
						Cons: ast.Boolean{Value: true},
						Alt:  ast.Boolean{Value: false},
					},
					Cons: ast.Integer{Value: 2},
					Alt:  ast.Ident{Value: "i"},
				},
			},
			want: eval.Integer{Value: 1},
		},
		{
			name: "let x = 2",
			input: ast.DeclStmt{
				Decl: ast.LetDecl{
					Id:   ast.Ident{Value: "x"},
					Expr: ast.Integer{Value: 2},
				},
			},
			want: eval.Integer{Value: 2},
		},
		{
			name: "let x = 2 in x + 3",
			input: ast.ExprStmt{
				Expr: ast.LetExpr{
					Id:   ast.Ident{Value: "x"},
					Bind: ast.Integer{Value: 2},
					Body: ast.BinOpExpr{
						Op:    token.PLUS,
						Left:  ast.Ident{Value: "x"},
						Right: ast.Integer{Value: 3},
					},
				},
			},
			want: eval.Integer{Value: 5},
		},
		{
			name: "let x = 2 in let y = 3 in x + y",
			input: ast.ExprStmt{
				Expr: ast.LetExpr{
					Id:   ast.Ident{Value: "x"},
					Bind: ast.Integer{Value: 2},
					Body: ast.LetExpr{
						Id:   ast.Ident{Value: "y"},
						Bind: ast.Integer{Value: 3},
						Body: ast.BinOpExpr{
							Op:    token.PLUS,
							Left:  ast.Ident{Value: "x"},
							Right: ast.Integer{Value: 3},
						},
					},
				},
			},
			want: eval.Integer{Value: 5},
		},
		{
			name: "(fun x -> x + 3 ) 2",
			input: ast.ExprStmt{
				Expr: ast.AppExpr{
					Func: ast.FunExpr{
						Param: ast.Ident{Value: "x"},
						Body: ast.BinOpExpr{
							Op:    token.PLUS,
							Left:  ast.Ident{Value: "x"},
							Right: ast.Integer{Value: 3},
						},
					},
					Arg: ast.Integer{Value: 2},
				},
			},
			want: eval.Integer{Value: 5},
		},
		{
			name: "(fun x -> (fun y -> x + y)) 2",
			input: ast.ExprStmt{
				Expr: ast.AppExpr{
					Func: ast.FunExpr{
						Param: ast.Ident{Value: "x"},
						Body: ast.FunExpr{
							Param: ast.Ident{Value: "y"},
							Body: ast.BinOpExpr{
								Op:    token.PLUS,
								Left:  ast.Ident{Value: "x"},
								Right: ast.Ident{Value: "y"},
							},
						},
					},
					Arg: ast.Integer{Value: 2},
				},
			},
			want: eval.Function{
				Param: ast.Ident{Value: "y"},
				Body: ast.BinOpExpr{
					Op:    token.PLUS,
					Left:  ast.Ident{Value: "x"},
					Right: ast.Ident{Value: "y"},
				},
				Env: eval.Env{
					Store: map[ast.Ident]eval.Value{
						{Value: "i"}: eval.Integer{Value: 1},
						{Value: "v"}: eval.Integer{Value: 5},
						{Value: "x"}: eval.Integer{Value: 2},
					},
				},
			},
		},
		{
			name: "(fun x -> (fun y -> x + y)) 2 3",
			input: ast.ExprStmt{
				Expr: ast.AppExpr{
					Func: ast.AppExpr{
						Func: ast.FunExpr{
							Param: ast.Ident{Value: "x"},
							Body: ast.FunExpr{
								Param: ast.Ident{Value: "y"},
								Body: ast.BinOpExpr{
									Op:    token.PLUS,
									Left:  ast.Ident{Value: "x"},
									Right: ast.Ident{Value: "y"},
								},
							},
						},
						Arg: ast.Integer{Value: 2},
					},
					Arg: ast.Integer{Value: 3},
				},
			},
			want: eval.Integer{Value: 5},
		},
		{
			name: "let rec f = fun n -> (if 10 < n then 1 else n * f (n + 1))",
			input: ast.DeclStmt{
				Decl: ast.RecDecl{
					Id:    ast.Ident{Value: "f"},
					Param: ast.Ident{Value: "n"},
					Body: ast.IfExpr{
						Cond: ast.BinOpExpr{
							Op:    token.LT,
							Left:  ast.Integer{Value: 10},
							Right: ast.Ident{Value: "n"},
						},
						Cons: ast.Integer{Value: 1},
						Alt: ast.BinOpExpr{
							Op:   token.ASTERISK,
							Left: ast.Ident{Value: "n"},
							Right: ast.AppExpr{
								Func: ast.Ident{Value: "f"},
								Arg: ast.BinOpExpr{
									Op:    token.PLUS,
									Left:  ast.Ident{Value: "n"},
									Right: ast.Integer{Value: 1},
								},
							},
						},
					},
				},
			},
			want: eval.Function{
				Param: ast.Ident{Value: "n"},
				Body: ast.IfExpr{
					Cond: ast.BinOpExpr{
						Op:    token.LT,
						Left:  ast.Integer{Value: 10},
						Right: ast.Ident{Value: "n"},
					},
					Cons: ast.Integer{Value: 1},
					Alt: ast.BinOpExpr{
						Op:   token.ASTERISK,
						Left: ast.Ident{Value: "n"},
						Right: ast.AppExpr{
							Func: ast.Ident{Value: "f"},
							Arg: ast.BinOpExpr{
								Op:    token.PLUS,
								Left:  ast.Ident{Value: "n"},
								Right: ast.Integer{Value: 1},
							},
						},
					},
				},
				Env: eval.Env{Store: map[ast.Ident]eval.Value{}},
			},
		},
		{
			name: "let rec fact = fun n -> (if 9 < n then 1 else n * (fact (n+1))) in fact 8",
			input: ast.ExprStmt{
				Expr: ast.LetRecExpr{
					Id:    ast.Ident{Value: "fact"},
					Param: ast.Ident{Value: "n"},
					Bind: ast.IfExpr{
						Cond: ast.BinOpExpr{
							Op:    token.LT,
							Left:  ast.Integer{Value: 9},
							Right: ast.Ident{Value: "n"},
						},
						Cons: ast.Integer{Value: 1},
						Alt: ast.BinOpExpr{
							Op:   token.ASTERISK,
							Left: ast.Ident{Value: "n"},
							Right: ast.AppExpr{
								Func: ast.Ident{Value: "fact"},
								Arg: ast.BinOpExpr{
									Op:    token.PLUS,
									Left:  ast.Ident{Value: "n"},
									Right: ast.Integer{Value: 1},
								},
							},
						},
					},
					Body: ast.AppExpr{
						Func: ast.Ident{Value: "fact"},
						Arg:  ast.Integer{Value: 8},
					},
				},
			},
			want: eval.Integer{Value: 72},
		},
	}

	for _, tc := range testCases {
		tc := tc

		env := eval.Env{Store: make(map[ast.Ident]eval.Value)}
		env.Set(ast.Ident{Value: "i"}, eval.Integer{Value: 1})
		env.Set(ast.Ident{Value: "v"}, eval.Integer{Value: 5})
		env.Set(ast.Ident{Value: "x"}, eval.Integer{Value: 10})

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			got := eval.Eval(tc.input, env)

			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("Eval(%q) returned unexpected difference (-want +got):\n%s", tc.input, diff)
			}
		})
	}
}
