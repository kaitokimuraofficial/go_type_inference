package typing

import (
	"github.com/google/go-cmp/cmp/cmpopts"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestReplace(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		desc string
		typ  Typ
		frm  TyVar
		to   Typ
		want Typ
	}{
		{
			desc: "TyInt does not change",
			typ:  &TyInt{},
			frm:  TyVar{Variable: 1},
			to:   &TyInt{},
			want: &TyInt{},
		},
		{
			desc: "TyBool does not change",
			typ:  &TyBool{},
			frm:  TyVar{Variable: 1},
			to:   &TyInt{},
			want: &TyBool{},
		},
		{
			desc: "TyIndent converts if it matches",
			typ:  &TyVar{Variable: 1},
			frm:  TyVar{Variable: 1},
			to:   &TyInt{},
			want: &TyInt{},
		},
		{
			desc: "TyIndent does not convert if it does not match",
			typ:  &TyVar{Variable: 1},
			frm:  TyVar{Variable: 2},
			to:   &TyInt{},
			want: &TyVar{Variable: 1},
		},
		{
			desc: "simple TyFun case",
			typ: &TyFun{
				Abs: &TyVar{
					Variable: 1,
				},
				App: &TyVar{
					Variable: 1,
				},
			},
			frm: TyVar{Variable: 1},
			to:  &TyBool{},
			want: &TyFun{
				Abs: &TyBool{},
				App: &TyBool{},
			},
		},
		{
			desc: "complicated TyFun case",
			typ: &TyFun{
				Abs: &TyFun{
					Abs: &TyFun{
						Abs: &TyVar{
							Variable: 3,
						},
						App: &TyVar{
							Variable: 1,
						},
					},
					App: &TyVar{
						Variable: 1,
					},
				},
				App: &TyVar{
					Variable: 2,
				},
			},
			frm: TyVar{Variable: 1},
			to:  &TyInt{},
			want: &TyFun{
				Abs: &TyFun{
					Abs: &TyFun{
						Abs: &TyVar{
							Variable: 3,
						},
						App: &TyInt{},
					},
					App: &TyInt{},
				},
				App: &TyVar{
					Variable: 2,
				},
			},
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.desc, func(t *testing.T) {
			t.Parallel()

			got := tc.typ.replace(tc.frm, tc.to)

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
		typ  Typ
		want []Variable
	}{
		{
			name: "integer",
			typ:  &TyInt{},
			want: []Variable{},
		},
		{
			name: "boolean",
			typ:  &TyBool{},
			want: []Variable{},
		},
		{
			name: "function bool to bool",
			typ: &TyFun{
				Abs: &TyBool{},
				App: &TyBool{},
			},
			want: []Variable{},
		},
		{
			name: "function ident to ident",
			typ: &TyFun{
				Abs: &TyVar{
					Variable: 1,
				},
				App: &TyVar{
					Variable: 1,
				},
			},
			want: []Variable{1},
		},
		{
			name: "nested function",
			typ: &TyFun{
				Abs: &TyFun{
					Abs: &TyFun{
						Abs: &TyVar{
							Variable: 3,
						},
						App: &TyVar{
							Variable: 1,
						},
					},
					App: &TyVar{
						Variable: 1,
					},
				},
				App: &TyVar{
					Variable: 3,
				},
			},
			want: []Variable{1, 3},
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			got := tc.typ.Variables()

			opt := cmpopts.SortSlices(func(i, j Variable) bool {
				return i < j
			})

			if diff := cmp.Diff(tc.want, got, opt); diff != "" {
				t.Errorf("returned unexpected difference (-want +got):\n%s", diff)
			}
		})
	}
}
