package lexer

import (
	"go_type_inference/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `=*+()
=*4 ;true
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
		{token.SEMISEMI, ";;"},
		{token.EOF, ""},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}

}
