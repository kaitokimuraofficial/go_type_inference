package typing

import (
	"go_type_inference/ast"
	"go_type_inference/parser"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestInfer(t *testing.T) {
	testCases := []struct {
		name  string
		input string
		want  Scheme
	}{
		{
			name:  "identifier",
			input: "i",
			want: Scheme{
				BTV: []Variable{}, Type: &TyInt{},
			},
		},
		{
			name:  "integer",
			input: "2",
			want: Scheme{
				BTV: []Variable{}, Type: &TyInt{},
			},
		},
		{
			name:  "boolean",
			input: "true",
			want: Scheme{
				BTV: []Variable{}, Type: &TyBool{},
			},
		},
		{
			name:  "add two integers",
			input: "2 + 3",
			want: Scheme{
				BTV: []Variable{}, Type: &TyInt{},
			},
		},
		{
			name:  "add two identifiers",
			input: "i + v",
			want: Scheme{
				BTV: []Variable{}, Type: &TyInt{},
			},
		},
		{
			name:  "multiple two integers",
			input: "2 * 3",
			want: Scheme{
				BTV: []Variable{}, Type: &TyInt{},
			},
		},
		{
			name:  "compare two integers",
			input: "2 < 3",
			want: Scheme{
				BTV: []Variable{}, Type: &TyBool{},
			},
		},
		{
			name:  "parenthesized expression",
			input: "(2 + 3)",
			want: Scheme{
				BTV: []Variable{}, Type: &TyInt{},
			},
		},
		{
			name:  "if true expression",
			input: "if true then 2 else 3",
			want: Scheme{
				BTV: []Variable{}, Type: &TyInt{},
			},
		},
		{
			name:  "if else expression",
			input: "if false then 2 else 3",
			want: Scheme{
				BTV: []Variable{}, Type: &TyInt{},
			},
		},
		{
			name:  "nested if expression",
			input: "if (if false then true else false) then 2 else 3",
			want: Scheme{
				BTV: []Variable{}, Type: &TyInt{},
			},
		},
		{
			name:  "nested if expression with identifiers in environment",
			input: "if (if i < v then true else false) then 2 else i",
			want: Scheme{
				BTV: []Variable{}, Type: &TyInt{},
			},
		},
		{
			name:  "let declaration",
			input: "let x = 2",
			want: Scheme{
				BTV: []Variable{}, Type: &TyInt{},
			},
		},
		{
			name:  "let expression",
			input: "let x = 2 in x + 3",
			want: Scheme{
				BTV: []Variable{}, Type: &TyInt{},
			},
		},
		{
			name:  "nested let expression",
			input: "let x = 2 in let y = 3 in x + y",
			want: Scheme{
				BTV: []Variable{}, Type: &TyInt{},
			},
		},
		{
			name:  "let poly declaration",
			input: "let f = fun x -> x",
			want: Scheme{
				BTV: []Variable{0},
				Type: &TyFun{
					Abs: &TyVar{Variable: 0},
					App: &TyVar{Variable: 0},
				},
			},
		},
		{
			name:  "let poly expression",
			input: "let f = fun x -> x in if f true then f 2 else f 3",
			want: Scheme{
				BTV: []Variable{}, Type: &TyInt{},
			},
		},
		{
			name:  "fun abstraction",
			input: "fun y -> y + 3",
			want: Scheme{
				BTV: []Variable{},
				Type: &TyFun{
					Abs: &TyInt{},
					App: &TyInt{},
				},
			},
		},
		{
			name:  "fun application-1",
			input: "(fun x -> x + 3 ) 2",
			want: Scheme{
				BTV: []Variable{}, Type: &TyInt{},
			},
		},
		{
			name:  "fun application-2",
			input: "(fun param -> param + 3 ) 2",
			want: Scheme{
				BTV: []Variable{}, Type: &TyInt{},
			},
		},
		{
			name:  "nested function application-1",
			input: "(fun x -> (fun y -> x + y)) 2",
			want: Scheme{
				BTV: []Variable{},
				Type: &TyFun{
					Abs: &TyInt{},
					App: &TyInt{},
				},
			},
		},
		{
			name:  "nested function application-2",
			input: "(fun x -> (fun y -> x + y)) 2 3",
			want: Scheme{
				BTV: []Variable{}, Type: &TyInt{},
			},
		},
		{
			name:  "fun abstraction",
			input: "fun y -> y + 3",
			want: Scheme{
				BTV: []Variable{},
				Type: &TyFun{
					Abs: &TyInt{},
					App: &TyInt{},
				},
			},
		},
		{
			name:  "fun application-1",
			input: "(fun x -> x + 3 ) 2",
			want: Scheme{
				BTV: []Variable{}, Type: &TyInt{},
			},
		},
		{
			name:  "fun application-2",
			input: "(fun param -> param + 3 ) 2",
			want: Scheme{
				BTV: []Variable{}, Type: &TyInt{},
			},
		},
		{
			name:  "nested function application-1",
			input: "(fun x -> (fun y -> x + y)) 2",
			want: Scheme{
				BTV: []Variable{},
				Type: &TyFun{
					Abs: &TyInt{},
					App: &TyInt{},
				},
			},
		},
		{
			name:  "nested function application-2",
			input: "(fun x -> (fun y -> x + y)) 2 3",
			want: Scheme{
				BTV: []Variable{}, Type: &TyInt{},
			},
		},
	}

	for _, tc := range testCases {
		tc := tc

		env := &Environment{Store: make(map[ast.Identifier]Scheme)}
		env.Set(ast.Identifier{Value: "b"}, Scheme{
			BTV:  []Variable{},
			Type: &TyInt{},
		})
		env.Set(ast.Identifier{Value: "i"}, Scheme{
			BTV:  []Variable{},
			Type: &TyInt{},
		})
		env.Set(ast.Identifier{Value: "v"}, Scheme{
			BTV:  []Variable{},
			Type: &TyInt{},
		})

		t.Run(tc.name, func(t *testing.T) {
			p := parser.Parse(tc.input)
			_, got := Infer(p, env)

			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("Infer(%q) returned unexpected difference (-want +got):\n%s", tc.input, diff)
			}
		})
	}
}
