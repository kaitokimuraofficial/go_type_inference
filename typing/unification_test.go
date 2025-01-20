package typing_test

import (
	"go_type_inference/typing"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestUnify(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name  string
		input []typing.Constraint
		want  []typing.Substitution
	}{
		{
			name: "x↦Bool y↦(x→Int)",
			input: []typing.Constraint{
				{
					Left:  typing.TyVar{Variable: 1},
					Right: typing.TyBool{},
				},
				{
					Left: typing.TyVar{Variable: 2},
					Right: typing.TyFun{
						Abs: typing.TyVar{Variable: 1},
						App: typing.TyInt{},
					},
				},
			},
			want: []typing.Substitution{
				{
					Variable: 2,
					Type: typing.TyFun{
						Abs: typing.TyBool{},
						App: typing.TyInt{},
					},
				},
				{
					Variable: 1,
					Type:     typing.TyBool{},
				},
			},
		},
		{
			name: "y↦(x→Int), x↦Bool (reversed from the previous case)",
			input: []typing.Constraint{
				{
					Left: typing.TyVar{Variable: 2},
					Right: typing.TyFun{
						Abs: typing.TyVar{Variable: 1},
						App: typing.TyInt{},
					},
				},
				{
					Left:  typing.TyVar{Variable: 1},
					Right: typing.TyBool{},
				},
			},
			want: []typing.Substitution{
				{
					Variable: 1,
					Type:     typing.TyBool{},
				},
				{
					Variable: 2,
					Type: typing.TyFun{
						Abs: typing.TyVar{Variable: 1},
						App: typing.TyInt{},
					},
				},
			},
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			got := typing.Unify(tc.input)

			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("returned unexpected difference (-want +got):\n%s", diff)
			}
		})
	}
}
