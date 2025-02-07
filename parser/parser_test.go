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
			input: "x ;;",
			want:  ast.ExprStmt{Expr: ast.Ident{Value: "x"}},
		},
		{
			name:  "integer",
			input: "4 ;;",
			want:  ast.ExprStmt{Expr: ast.Integer{Value: 4}},
		},
		{
			name:  "boolean",
			input: "true ;;",
			want:  ast.ExprStmt{Expr: ast.Boolean{Value: true}},
		},
		{
			name:  "primitive operator",
			input: "2 + 3 ;;",
			want: ast.ExprStmt{
				Expr: ast.BinOpExpr{
					Op:    token.PLUS,
					Left:  ast.Integer{Value: 2},
					Right: ast.Integer{Value: 3},
				},
			},
		},
		{
			name:  "if expression",
			input: "if true then true else false ;;",
			want: ast.ExprStmt{
				Expr: ast.IfExpr{
					Cond: ast.Boolean{Value: true},
					Cons: ast.Boolean{Value: true},
					Alt:  ast.Boolean{Value: false},
				},
			},
		},
		{
			name:  "parenthesized expression",
			input: "(true) ;;",
			want: ast.ExprStmt{
				Expr: ast.Boolean{Value: true},
			},
		},
		{
			name:  "let declaration",
			input: "let x = 5 ;;",
			want: ast.DeclStmt{
				Decl: ast.LetDecl{
					Id:   ast.Ident{Value: "x"},
					Expr: ast.Integer{Value: 5},
				},
			},
		},
		{
			name:  "let expression",
			input: "let x = 2 in x + 3 ;;",
			want: ast.ExprStmt{
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
		},
		{
			name:  "function declaration",
			input: "fun x -> x + 3 ;;",
			want: ast.ExprStmt{
				Expr: ast.FunExpr{
					Param: ast.Ident{Value: "x"},
					Body: ast.BinOpExpr{
						Op:    token.PLUS,
						Left:  ast.Ident{Value: "x"},
						Right: ast.Integer{Value: 3},
					},
				},
			},
		},
		{
			name:  "function application",
			input: "(fun x -> x + 3 ) 2 ;;",
			want: ast.ExprStmt{
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
		},
		{
			name:  "nested function declaration",
			input: "fun x -> (fun y -> x + y) ;;",
			want: ast.ExprStmt{
				Expr: ast.FunExpr{
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
			},
		},
		{
			name:  "nested function application",
			input: "(fun x -> (fun y -> x + y)) 2 ;;",
			want: ast.ExprStmt{
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
					Arg: ast.Integer{
						Value: 2,
					},
				},
			},
		},
		{
			name:  "recursive function declaration",
			input: "let rec f = fun n -> if n < 10 then 1 else n * f (n + 1) ;;",
			want: ast.DeclStmt{
				Decl: ast.RecDecl{
					Id:    ast.Ident{Value: "f"},
					Param: ast.Ident{Value: "n"},
					Body: ast.IfExpr{
						Cond: ast.BinOpExpr{
							Op:    token.LT,
							Left:  ast.Ident{Value: "n"},
							Right: ast.Integer{Value: 10},
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
		},
		{
			name:  "recursive function expression",
			input: "let rec fact = fun n -> if n < 1 then 1 else n * (fact (n + 1)) in fact 5 ;;",
			want: ast.ExprStmt{
				Expr: ast.LetRecExpr{
					Id:    ast.Ident{Value: "fact"},
					Param: ast.Ident{Value: "n"},
					Bind: ast.IfExpr{
						Cond: ast.BinOpExpr{
							Op:    token.LT,
							Left:  ast.Ident{Value: "n"},
							Right: ast.Integer{Value: 1},
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
						Arg:  ast.Integer{Value: 5},
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
