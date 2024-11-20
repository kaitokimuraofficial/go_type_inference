package typing

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestUnify(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name  string
		input []Constraint
		want  []Substitution
	}{
		{
			name: "x↦Bool y↦(x→Int)",
			input: []Constraint{
				{
					Left:  &TyVar{Variable: 1},
					Right: &TyBool{},
				},
				{
					Left: &TyVar{Variable: 2},
					Right: &TyFun{
						Abs: &TyVar{Variable: 1},
						App: &TyInt{},
					},
				},
			},
			want: []Substitution{
				{
					Var: TyVar{Variable: 2},
					Type: &TyFun{
						Abs: &TyBool{},
						App: &TyInt{},
					},
				},
				{
					Var:  TyVar{Variable: 1},
					Type: &TyBool{},
				},
			},
		},
		{
			name: "y↦(x→Int), x↦Bool (reversed from the previous case)",
			input: []Constraint{
				{
					Left: &TyVar{Variable: 2},
					Right: &TyFun{
						Abs: &TyVar{Variable: 1},
						App: &TyInt{},
					},
				},
				{
					Left:  &TyVar{Variable: 1},
					Right: &TyBool{},
				},
			},
			want: []Substitution{
				{
					Var:  TyVar{Variable: 1},
					Type: &TyBool{},
				},
				{
					Var: TyVar{Variable: 2},
					Type: &TyFun{
						Abs: &TyVar{Variable: 1},
						App: &TyInt{},
					},
				},
			},
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			got := Unify(tc.input)

			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("returned unexpected difference (-want +got):\n%s", diff)
			}
		})
	}
}
