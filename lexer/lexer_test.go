package lexer_test

import (
	"go_type_inference/lexer"
	"go_type_inference/token"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestNextToken(t *testing.T) {
	testCases := []struct {
		name  string
		input string
		want  []token.Token
	}{
		{
			name: "example sentence",
			input: `=*+()
			=*4 true in let inlet fun -> rec
			;;
			`,
			want: []token.Token{
				{Type: token.ASSIGN, Literal: "="},
				{Type: token.ASTERISK, Literal: "*"},
				{Type: token.PLUS, Literal: "+"},
				{Type: token.LPAREN, Literal: "("},
				{Type: token.RPAREN, Literal: ")"},
				{Type: token.ASSIGN, Literal: "="},
				{Type: token.ASTERISK, Literal: "*"},
				{Type: token.INT, Literal: "4"},
				{Type: token.TRUE, Literal: "true"},
				{Type: token.IN, Literal: "in"},
				{Type: token.LET, Literal: "let"},
				{Type: token.IDENT, Literal: "inlet"},
				{Type: token.FUN, Literal: "fun"},
				{Type: token.RARROW, Literal: "->"},
				{Type: token.REC, Literal: "rec"},
				{Type: token.SEMISEMI, Literal: ";;"},
				{Type: token.EOF, Literal: ""},
			},
		},
		{
			name:  "illegal sentence with -",
			input: `push-back`,
			want: []token.Token{
				{Type: token.IDENT, Literal: "push"},
				{Type: token.ILLEGAL, Literal: "-"},
				{Type: token.IDENT, Literal: "back"},
				{Type: token.EOF, Literal: ""},
			},
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			l := lexer.New(tc.input)

			var got []token.Token

			for {
				tok := l.NextToken()
				got = append(got, tok)

				if tok.Type == token.EOF {
					break
				}
			}

			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("lexer.New(%q) (-want +got):\n%s", tc.input, diff)
			}
		})
	}
}
