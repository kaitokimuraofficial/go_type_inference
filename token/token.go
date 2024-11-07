package token

type TokenType string

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifier
	IDENT = "IDENT"
	INT   = "INT"

	// Operator
	ASSIGN   = "="
	ASTERISK = "*"
	PLUS     = "+"

	LT = "<"

	// Delimiter
	SEMI     = ";"
	SEMISEMI = ";;"

	LPAREN = "("
	RPAREN = ")"

	// Keyword
	ELSE  = "ELSE"
	FALSE = "FALSE"
	IF    = "IF"
	THEN  = "THEN"
	TRUE  = "TRUE"
)

var keywords = map[string]TokenType{
	"else":  ELSE,
	"false": FALSE,
	"if":    IF,
	"then":  THEN,
	"true":  TRUE,
}

type Token struct {
	Type    TokenType
	Literal string
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
