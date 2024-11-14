package typing

import (
	"github.com/google/go-cmp/cmp/cmpopts"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestConvert(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		desc string
		typ  Type
		frm  TyIdent
		to   Type
		want Type
	}{
		{
			desc: "TyInt does not change",
			typ:  &TyInt{},
			frm:  TyIdent{Variable: 1},
			to:   &TyInt{},
			want: &TyInt{},
		},
		{
			desc: "TyBool does not change",
			typ:  &TyBool{},
			frm:  TyIdent{Variable: 1},
			to:   &TyInt{},
			want: &TyBool{},
		},
		{
			desc: "TyIndent converts if it matches",
			typ:  &TyIdent{Variable: 1},
			frm:  TyIdent{Variable: 1},
			to:   &TyInt{},
			want: &TyInt{},
		},
		{
			desc: "TyIndent does not convert if it does not match",
			typ:  &TyIdent{Variable: 1},
			frm:  TyIdent{Variable: 2},
			to:   &TyInt{},
			want: &TyIdent{Variable: 1},
		},
		{
			desc: "simple TyFun case",
			typ: &TyFun{
				Abs: &TyIdent{
					Variable: 1,
				},
				App: &TyIdent{
					Variable: 1,
				},
			},
			frm: TyIdent{Variable: 1},
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
						Abs: &TyIdent{
							Variable: 3,
						},
						App: &TyIdent{
							Variable: 1,
						},
					},
					App: &TyIdent{
						Variable: 1,
					},
				},
				App: &TyIdent{
					Variable: 2,
				},
			},
			frm: TyIdent{Variable: 1},
			to:  &TyInt{},
			want: &TyFun{
				Abs: &TyFun{
					Abs: &TyFun{
						Abs: &TyIdent{
							Variable: 3,
						},
						App: &TyInt{},
					},
					App: &TyInt{},
				},
				App: &TyIdent{
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
		typ  Type
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
				Abs: &TyIdent{
					Variable: 1,
				},
				App: &TyIdent{
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
						Abs: &TyIdent{
							Variable: 3,
						},
						App: &TyIdent{
							Variable: 1,
						},
					},
					App: &TyIdent{
						Variable: 1,
					},
				},
				App: &TyIdent{
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
