package token

import "strconv"

type Type int

const (
	ILLEGAL Type = iota
	EOF

	// Identifier
	IDENT // x
	INT   // 123

	// Operator
	ASSIGN   // =
	ASTERISK // *
	PLUS     // +

	LT // <

	// Delimiter
	SEMI     // ;
	SEMISEMI // ;;

	LPAREN // (
	RPAREN // )

	RARROW // ->

	// Keyword
	ELSE
	FALSE
	FUN
	IF
	IN
	LET
	THEN
	TRUE
	REC
)

var tokens = [...]string{
	ILLEGAL: "ILLEGAL",
	EOF:     "EOF",

	IDENT: "IDENT",
	INT:   "INT",

	ASSIGN:   "=",
	ASTERISK: "*",
	PLUS:     "+",

	LT: "<",

	SEMI:     ";",
	SEMISEMI: ";;",

	LPAREN: "(",
	RPAREN: ")",

	RARROW: "->",

	ELSE:  "else",
	FALSE: "false",
	FUN:   "fun",
	IF:    "if",
	IN:    "in",
	LET:   "let",
	THEN:  "then",
	TRUE:  "true",
	REC:   "rec",
}

var keywords = map[string]Type{
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
	Type    Type
	Literal string
}

func New(t Type, ch byte) *Token {
	return &Token{Type: t, Literal: string(ch)}
}

func Lookup(ident string) Type {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}

func (t Type) String() string {
	s := ""
	if 0 <= t && t < Type(len(tokens)) {
		s = tokens[t]
	}
	if s == "" {
		s = "token(" + strconv.Itoa(int(t)) + ")"
	}
	return s
}
