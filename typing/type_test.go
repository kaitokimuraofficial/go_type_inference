package typing_test

import (
	"go_type_inference/typing"
	"testing"

	"github.com/google/go-cmp/cmp/cmpopts"

	"github.com/google/go-cmp/cmp"
)

func TestConvert(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		desc string
		typ  typing.Type
		frm  typing.TyVar
		to   typing.Type
		want typing.Type
	}{
		{
			desc: "TyInt does not change",
			typ:  typing.TyInt{},
			frm:  typing.TyVar{Variable: 1},
			to:   typing.TyInt{},
			want: typing.TyInt{},
		},
		{
			desc: "TyBool does not change",
			typ:  typing.TyBool{},
			frm:  typing.TyVar{Variable: 1},
			to:   typing.TyInt{},
			want: typing.TyBool{},
		},
		{
			desc: "TyIndent converts if it matches",
			typ:  typing.TyVar{Variable: 1},
			frm:  typing.TyVar{Variable: 1},
			to:   typing.TyInt{},
			want: typing.TyInt{},
		},
		{
			desc: "TyIndent does not convert if it does not match",
			typ:  typing.TyVar{Variable: 1},
			frm:  typing.TyVar{Variable: 2},
			to:   typing.TyInt{},
			want: typing.TyVar{Variable: 1},
		},
		{
			desc: "simple TyFun case",
			typ: typing.TyFun{
				Abs: typing.TyVar{
					Variable: 1,
				},
				App: typing.TyVar{
					Variable: 1,
				},
			},
			frm: typing.TyVar{Variable: 1},
			to:  typing.TyBool{},
			want: typing.TyFun{
				Abs: typing.TyBool{},
				App: typing.TyBool{},
			},
		},
		{
			desc: "complicated TyFun case",
			typ: typing.TyFun{
				Abs: typing.TyFun{
					Abs: typing.TyFun{
						Abs: typing.TyVar{
							Variable: 3,
						},
						App: typing.TyVar{
							Variable: 1,
						},
					},
					App: typing.TyVar{
						Variable: 1,
					},
				},
				App: typing.TyVar{
					Variable: 2,
				},
			},
			frm: typing.TyVar{Variable: 1},
			to:  typing.TyInt{},
			want: typing.TyFun{
				Abs: typing.TyFun{
					Abs: typing.TyFun{
						Abs: typing.TyVar{
							Variable: 3,
						},
						App: typing.TyInt{},
					},
					App: typing.TyInt{},
				},
				App: typing.TyVar{
					Variable: 2,
				},
			},
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.desc, func(t *testing.T) {
			t.Parallel()

			got := tc.typ.Convert(tc.frm, tc.to)

			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("returned unexpected difference (-want +got):\n%s", diff)
			}
		})
	}
}

func TestVariables(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name string
		typ  typing.Type
		want []typing.Variable
	}{
		{
			name: "integer",
			typ:  typing.TyInt{},
			want: []typing.Variable{},
		},
		{
			name: "boolean",
			typ:  typing.TyBool{},
			want: []typing.Variable{},
		},
		{
			name: "function bool to bool",
			typ: typing.TyFun{
				Abs: typing.TyBool{},
				App: typing.TyBool{},
			},
			want: []typing.Variable{},
		},
		{
			name: "function ident to ident",
			typ: typing.TyFun{
				Abs: typing.TyVar{
					Variable: 1,
				},
				App: typing.TyVar{
					Variable: 1,
				},
			},
			want: []typing.Variable{1},
		},
		{
			name: "nested function",
			typ: typing.TyFun{
				Abs: typing.TyFun{
					Abs: typing.TyFun{
						Abs: typing.TyVar{
							Variable: 3,
						},
						App: typing.TyVar{
							Variable: 1,
						},
					},
					App: typing.TyVar{
						Variable: 1,
					},
				},
				App: typing.TyVar{
					Variable: 3,
				},
			},
			want: []typing.Variable{1, 3},
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			got := tc.typ.Variables()

			opt := cmpopts.SortSlices(func(i, j typing.Variable) bool {
				return i < j
			})

			if diff := cmp.Diff(tc.want, got, opt); diff != "" {
				t.Errorf("returned unexpected difference (-want +got):\n%s", diff)
			}
		})
	}
}
