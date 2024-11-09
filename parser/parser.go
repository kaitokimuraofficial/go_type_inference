// Code generated by goyacc -o parser/parser.go parser/parser.go.y. DO NOT EDIT.

//line parser/parser.go.y:2
package parser

import __yyfmt__ "fmt"

//line parser/parser.go.y:2

import (
	"fmt"
	"go_type_inference/ast"
	"go_type_inference/lexer"
	"go_type_inference/token"
	"strconv"
)

//line parser/parser.go.y:13
type yySymType struct {
	yys       int
	statement ast.Stmt
	expr      ast.Expr
	token     token.Token
}

const IDENT = 57346
const INT = 57347
const TRUE = 57348
const FALSE = 57349
const LPAREN = 57350
const RPAREN = 57351
const IF = 57352
const THEN = 57353
const ELSE = 57354
const LT = 57355
const PLUS = 57356
const ASTERISK = 57357

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"IDENT",
	"INT",
	"TRUE",
	"FALSE",
	"LPAREN",
	"RPAREN",
	"IF",
	"THEN",
	"ELSE",
	"LT",
	"PLUS",
	"ASTERISK",
}

var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line parser/parser.go.y:115

type LexerWrapper struct {
	Lexer  *lexer.Lexer
	Result ast.Stmt
}

func (lw *LexerWrapper) Lex(lval *yySymType) int {
	tok := lw.Lexer.NextToken()
	lval.token = tok

	switch tok.Type {
	case token.ILLEGAL:
		return -1
	case token.EOF:
		return 0
	case token.IDENT:
		return IDENT
	case token.INT:
		return INT
	case token.ASTERISK:
		return ASTERISK
	case token.PLUS:
		return PLUS
	case token.LT:
		return LT
	case token.LPAREN:
		return LPAREN
	case token.RPAREN:
		return RPAREN
	case token.ELSE:
		return ELSE
	case token.FALSE:
		return FALSE
	case token.IF:
		return IF
	case token.THEN:
		return THEN
	case token.TRUE:
		return TRUE
	default:
		return -1
	}
}

func (lw *LexerWrapper) Error(e string) {
	fmt.Println("[error] " + e)
}

func Parse(input string) ast.Stmt {
	l := lexer.New(input)
	lw := &LexerWrapper{Lexer: l}
	yyParse(lw)
	return lw.Result
}

//line yacctab:1
var yyExca = [...]int8{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyPrivate = 57344

const yyLast = 32

var yyAct = [...]int8{
	2, 8, 6, 17, 7, 16, 14, 12, 9, 10,
	11, 13, 25, 5, 18, 15, 16, 19, 20, 22,
	24, 21, 23, 3, 4, 1, 26, 12, 9, 10,
	11, 13,
}

var yyPact = [...]int16{
	3, -1000, -1000, -1000, -1000, 3, 2, -12, -1000, -1000,
	-1000, -1000, -1000, 3, 6, 23, 23, 23, 13, 3,
	-9, -12, -1000, -1000, 0, 3, -1000,
}

var yyPgo = [...]int8{
	0, 25, 0, 24, 2, 4, 1, 23,
}

var yyR1 = [...]int8{
	0, 1, 2, 2, 3, 3, 4, 4, 5, 5,
	6, 6, 6, 6, 6, 7,
}

var yyR2 = [...]int8{
	0, 1, 1, 1, 3, 1, 3, 1, 3, 1,
	1, 1, 1, 1, 3, 6,
}

var yyChk = [...]int16{
	-1000, -1, -2, -7, -3, 10, -4, -5, -6, 5,
	6, 7, 4, 8, -2, 13, 14, 15, -2, 11,
	-4, -5, -6, 9, -2, 12, -2,
}

var yyDef = [...]int8{
	0, -2, 1, 2, 3, 0, 5, 7, 9, 10,
	11, 12, 13, 0, 0, 0, 0, 0, 0, 0,
	4, 6, 8, 14, 0, 0, 15,
}

var yyTok1 = [...]int8{
	1,
}

var yyTok2 = [...]int8{
	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15,
}

var yyTok3 = [...]int8{
	0,
}

var yyErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

//line yaccpar:1

/*	parser for yacc output	*/

var (
	yyDebug        = 0
	yyErrorVerbose = false
)

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

type yyParser interface {
	Parse(yyLexer) int
	Lookahead() int
}

type yyParserImpl struct {
	lval  yySymType
	stack [yyInitialStackSize]yySymType
	char  int
}

func (p *yyParserImpl) Lookahead() int {
	return p.char
}

func yyNewParser() yyParser {
	return &yyParserImpl{}
}

const yyFlag = -1000

func yyTokname(c int) string {
	if c >= 1 && c-1 < len(yyToknames) {
		if yyToknames[c-1] != "" {
			return yyToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func yyStatname(s int) string {
	if s >= 0 && s < len(yyStatenames) {
		if yyStatenames[s] != "" {
			return yyStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func yyErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !yyErrorVerbose {
		return "syntax error"
	}

	for _, e := range yyErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + yyTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := int(yyPact[state])
	for tok := TOKSTART; tok-1 < len(yyToknames); tok++ {
		if n := base + tok; n >= 0 && n < yyLast && int(yyChk[int(yyAct[n])]) == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if yyDef[state] == -2 {
		i := 0
		for yyExca[i] != -1 || int(yyExca[i+1]) != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; yyExca[i] >= 0; i += 2 {
			tok := int(yyExca[i])
			if tok < TOKSTART || yyExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if yyExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += yyTokname(tok)
	}
	return res
}

func yylex1(lex yyLexer, lval *yySymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = int(yyTok1[0])
		goto out
	}
	if char < len(yyTok1) {
		token = int(yyTok1[char])
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			token = int(yyTok2[char-yyPrivate])
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		token = int(yyTok3[i+0])
		if token == char {
			token = int(yyTok3[i+1])
			goto out
		}
	}

out:
	if token == 0 {
		token = int(yyTok2[1]) /* unknown char */
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", yyTokname(token), uint(char))
	}
	return char, token
}

func yyParse(yylex yyLexer) int {
	return yyNewParser().Parse(yylex)
}

func (yyrcvr *yyParserImpl) Parse(yylex yyLexer) int {
	var yyn int
	var yyVAL yySymType
	var yyDollar []yySymType
	_ = yyDollar // silence set and not used
	yyS := yyrcvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yyrcvr.char = -1
	yytoken := -1 // yyrcvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		yystate = -1
		yyrcvr.char = -1
		yytoken = -1
	}()
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", yyTokname(yytoken), yyStatname(yystate))
	}

	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	yyn = int(yyPact[yystate])
	if yyn <= yyFlag {
		goto yydefault /* simple state */
	}
	if yyrcvr.char < 0 {
		yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
	}
	yyn += yytoken
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = int(yyAct[yyn])
	if int(yyChk[yyn]) == yytoken { /* valid shift */
		yyrcvr.char = -1
		yytoken = -1
		yyVAL = yyrcvr.lval
		yystate = yyn
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	}

yydefault:
	/* default state action */
	yyn = int(yyDef[yystate])
	if yyn == -2 {
		if yyrcvr.char < 0 {
			yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if yyExca[xi+0] == -1 && int(yyExca[xi+1]) == yystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			yyn = int(yyExca[xi+0])
			if yyn < 0 || yyn == yytoken {
				break
			}
		}
		yyn = int(yyExca[xi+1])
		if yyn < 0 {
			goto ret0
		}
	}
	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			yylex.Error(yyErrorMessage(yystate, yytoken))
			Nerrs++
			if yyDebug >= 1 {
				__yyfmt__.Printf("%s", yyStatname(yystate))
				__yyfmt__.Printf(" saw %s\n", yyTokname(yytoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				yyn = int(yyPact[yyS[yyp].yys]) + yyErrCode
				if yyn >= 0 && yyn < yyLast {
					yystate = int(yyAct[yyn]) /* simulate a shift of "error" */
					if int(yyChk[yystate]) == yyErrCode {
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yytoken))
			}
			if yytoken == yyEofCode {
				goto ret1
			}
			yyrcvr.char = -1
			yytoken = -1
			goto yynewstate /* try again in the same state */
		}
	}

	/* reduction by production yyn */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", yyn, yyStatname(yystate))
	}

	yynt := yyn
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= int(yyR2[yyn])
	// yyp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if yyp+1 >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	yyn = int(yyR1[yyn])
	yyg := int(yyPgo[yyn])
	yyj := yyg + yyS[yyp].yys + 1

	if yyj >= yyLast {
		yystate = int(yyAct[yyg])
	} else {
		yystate = int(yyAct[yyj])
		if int(yyChk[yystate]) != -yyn {
			yystate = int(yyAct[yyg])
		}
	}
	// dummy call; replaced with literal code
	switch yynt {

	case 1:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser/parser.go.y:37
		{
			yyVAL.statement = ast.Statement{Expr: yyDollar[1].expr}
			yylex.(*LexerWrapper).Result = yyVAL.statement
		}
	case 2:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser/parser.go.y:44
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 3:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser/parser.go.y:48
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 4:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser/parser.go.y:54
		{
			yyVAL.expr = ast.BinOpExpr{Token: yyDollar[2].token, Left: yyDollar[1].expr, Operator: token.LT, Right: yyDollar[3].expr}
		}
	case 5:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser/parser.go.y:58
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 6:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser/parser.go.y:64
		{
			yyVAL.expr = ast.BinOpExpr{Token: yyDollar[2].token, Left: yyDollar[1].expr, Operator: token.PLUS, Right: yyDollar[3].expr}
		}
	case 7:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser/parser.go.y:68
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 8:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser/parser.go.y:74
		{
			yyVAL.expr = ast.BinOpExpr{Token: yyDollar[2].token, Left: yyDollar[1].expr, Operator: token.ASTERISK, Right: yyDollar[3].expr}
		}
	case 9:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser/parser.go.y:78
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 10:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser/parser.go.y:84
		{
			intValue, err := strconv.Atoi(yyDollar[1].token.Literal)
			if err != nil {
				yylex.(*LexerWrapper).Error(fmt.Sprintf("invalid integer value: %s", yyDollar[1].token.Literal))
				return 1
			}
			yyVAL.expr = ast.Integer{Token: yyDollar[1].token, Value: intValue}
		}
	case 11:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser/parser.go.y:93
		{
			yyVAL.expr = ast.Boolean{Token: yyDollar[1].token, Value: true}
		}
	case 12:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser/parser.go.y:97
		{
			yyVAL.expr = ast.Boolean{Token: yyDollar[1].token, Value: false}
		}
	case 13:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser/parser.go.y:101
		{
			yyVAL.expr = ast.Identifier{Token: yyDollar[1].token, Value: yyDollar[1].token.Literal}
		}
	case 14:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser/parser.go.y:105
		{
			yyVAL.expr = yyDollar[2].expr
		}
	case 15:
		yyDollar = yyS[yypt-6 : yypt+1]
//line parser/parser.go.y:111
		{
			yyVAL.expr = ast.IfExpr{Token: yyDollar[1].token, Condition: yyDollar[2].expr, Consequence: yyDollar[4].expr, Alternative: yyDollar[6].expr}
		}
	}
	goto yystack /* stack new state and value */
}