package lexer

import (
	"go_type_inference/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `=*+()
=*4 ;true in let inlet fun -> rec
;;
`

	wants := []struct {
		Type    token.Type
		Literal string
	}{
		{token.ASSIGN, "="},
		{token.ASTERISK, "*"},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.ASSIGN, "="},
		{token.ASTERISK, "*"},
		{token.INT, "4"},
		{token.SEMI, ";"},
		{token.TRUE, "true"},
		{token.IN, "in"},
		{token.LET, "let"},
		{token.IDENT, "inlet"},
		{token.FUN, "fun"},
		{token.RARROW, "->"},
		{token.REC, "rec"},
		{token.SEMISEMI, ";;"},
		{token.EOF, ""},
	}

	l := New(input)

	for _, want := range wants {
		got := l.NextToken()

		if got.Type != want.Type {
			t.Errorf("different token types: got %q, but got: %q", got.Type, want.Type)
		}

		if got.Literal != want.Literal {
			t.Errorf("different token literals: got %q, but got: %q", got.Literal, want.Literal)
		}
	}

}
