package typing

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestUnify(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name    string
		handler ConstraintSet
		want    Substitution
	}{
		{
			name: "x↦Bool y↦(x→Int)",
			handler: ConstraintSet{
				{
					Left:  &TyIdent{Variable: 1},
					Right: &TyBool{},
				},
				{
					Left: &TyIdent{Variable: 2},
					Right: &TyFun{
						Abs: &TyIdent{Variable: 1},
						App: &TyInt{},
					},
				},
			},
			want: Substitution{
				{
					Variable: 2,
					Type: &TyFun{
						Abs: &TyBool{},
						App: &TyInt{},
					},
				},
				{
					Variable: 1,
					Type:     &TyBool{},
				},
			},
		},
		{
			name: "y↦(x→Int), x↦Bool (reversed from the previous case)",
			handler: ConstraintSet{
				{
					Left: &TyIdent{Variable: 2},
					Right: &TyFun{
						Abs: &TyIdent{Variable: 1},
						App: &TyInt{},
					},
				},
				{
					Left:  &TyIdent{Variable: 1},
					Right: &TyBool{},
				},
			},
			want: Substitution{
				{
					Variable: 1,
					Type:     &TyBool{},
				},
				{
					Variable: 2,
					Type: &TyFun{
						Abs: &TyIdent{Variable: 1},
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

			got := tc.handler.Unify()

			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("returned unexpected difference (-want +got):\n%s", diff)
			}
		})
	}
}
