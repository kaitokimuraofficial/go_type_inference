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

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
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

	for _, tt := range tests {
		got := l.NextToken()

		if got.Type != tt.expectedType {
			t.Errorf("[TOKEN TYPE ERROR] expect: %q, but got: %q", tt.expectedType, got.Type)
		}

		if got.Literal != tt.expectedLiteral {
			t.Errorf("[TOKEN LITERAL ERROR] expect: %q, but got: %q", tt.expectedLiteral, got.Literal)
		}
	}

}
