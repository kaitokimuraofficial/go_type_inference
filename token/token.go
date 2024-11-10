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

	RARROW = "->"

	// Keyword
	ELSE  = "ELSE"
	FALSE = "FALSE"
	FUN   = "FUN"
	IF    = "IF"
	IN    = "IN"
	LET   = "LET"
	THEN  = "THEN"
	TRUE  = "TRUE"
	REC   = "REC"
)

var keywords = map[string]TokenType{
	"else":  ELSE,
	"false": FALSE,
	"fun":   FUN,
	"if":    IF,
	"in":    IN,
	"let":   LET,
	"then":  THEN,
	"true":  TRUE,
	"rec":   REC,
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
