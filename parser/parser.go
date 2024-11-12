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
const LET = 57358
const ASSIGN = 57359
const IN = 57360
const RARROW = 57361
const FUN = 57362
const REC = 57363

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
	"LET",
	"ASSIGN",
	"IN",
	"RARROW",
	"FUN",
	"REC",
}

var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line parser/parser.go.y:172

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
	case token.ASSIGN:
		return ASSIGN
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
	case token.RARROW:
		return RARROW
	case token.ELSE:
		return ELSE
	case token.FALSE:
		return FALSE
	case token.FUN:
		return FUN
	case token.IF:
		return IF
	case token.IN:
		return IN
	case token.LET:
		return LET
	case token.THEN:
		return THEN
	case token.TRUE:
		return TRUE
	case token.REC:
		return REC
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

const yyLast = 75

var yyAct = [...]int8{
	2, 14, 13, 10, 54, 18, 15, 16, 17, 19,
	22, 9, 47, 58, 55, 28, 12, 23, 37, 34,
	29, 11, 59, 18, 15, 16, 17, 19, 35, 9,
	38, 40, 46, 42, 20, 3, 33, 49, 45, 11,
	28, 44, 36, 41, 30, 50, 27, 51, 25, 53,
	48, 21, 24, 25, 32, 39, 57, 56, 52, 60,
	61, 18, 15, 16, 17, 19, 43, 31, 26, 4,
	6, 5, 7, 8, 1,
}

var yyPact = [...]int16{
	19, -1000, -1000, 30, -1000, -1000, -1000, -1000, -1000, 1,
	39, 64, 31, 57, -1000, -1000, -1000, -1000, -1000, 1,
	27, 63, 43, 15, 57, 57, -1, 57, -1000, 46,
	1, 26, 1, 62, 24, 34, 31, 1, 57, -1000,
	14, -8, 38, 20, 1, -1000, 1, 54, 1, -16,
	14, -1000, -5, -1000, 53, 1, -6, 4, 1, 1,
	4, -1000,
}

var yyPgo = [...]int8{
	0, 74, 0, 73, 72, 71, 70, 3, 16, 2,
	1, 69,
}

var yyR1 = [...]int8{
	0, 1, 1, 1, 2, 2, 2, 2, 2, 3,
	4, 5, 6, 6, 7, 7, 8, 8, 9, 9,
	10, 10, 10, 10, 10, 11,
}

var yyR2 = [...]int8{
	0, 1, 4, 8, 1, 1, 1, 1, 1, 10,
	4, 6, 3, 1, 3, 1, 3, 1, 2, 1,
	1, 1, 1, 1, 3, 6,
}

var yyChk = [...]int16{
	-1000, -1, -2, 16, -11, -5, -6, -4, -3, 10,
	-7, 20, -8, -9, -10, 5, 6, 7, 4, 8,
	4, 21, -2, 16, 13, 14, 4, 15, -10, -2,
	17, 4, 11, 21, 4, -7, -8, 19, -9, 9,
	-2, 17, -2, 4, 17, -2, 18, 20, 12, 17,
	-2, -2, 4, -2, 20, 19, 4, -2, 19, 18,
	-2, -2,
}

var yyDef = [...]int8{
	0, -2, 1, 0, 4, 5, 6, 7, 8, 0,
	13, 0, 15, 17, 19, 20, 21, 22, 23, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 18, 0,
	0, 0, 0, 0, 0, 12, 14, 0, 16, 24,
	2, 0, 0, 0, 0, 10, 0, 0, 0, 0,
	0, 11, 0, 25, 0, 0, 0, 3, 0, 0,
	0, 9,
}

var yyTok1 = [...]int8{
	1,
}

var yyTok2 = [...]int8{
	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
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
//line parser/parser.go.y:44
		{
			yyVAL.statement = ast.Statement{Expr: yyDollar[1].expr}
			yylex.(*LexerWrapper).Result = yyVAL.statement
		}
	case 2:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser/parser.go.y:49
		{
			yyVAL.statement = ast.Declaration{Id: ast.Identifier{Token: yyDollar[2].token, Value: yyDollar[2].token.Literal}, Expr: yyDollar[4].expr}
			yylex.(*LexerWrapper).Result = yyVAL.statement
		}
	case 3:
		yyDollar = yyS[yypt-8 : yypt+1]
//line parser/parser.go.y:54
		{
			yyVAL.statement = ast.RecDeclaration{Id: ast.Identifier{Token: yyDollar[3].token, Value: yyDollar[3].token.Literal}, Param: ast.Identifier{Token: yyDollar[6].token, Value: yyDollar[6].token.Literal}, BodyExpr: yyDollar[8].expr}
			yylex.(*LexerWrapper).Result = yyVAL.statement
		}
	case 4:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser/parser.go.y:61
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 5:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser/parser.go.y:65
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 6:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser/parser.go.y:69
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 7:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser/parser.go.y:73
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 8:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser/parser.go.y:77
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 9:
		yyDollar = yyS[yypt-10 : yypt+1]
//line parser/parser.go.y:83
		{
			yyVAL.expr = ast.LetRecExpr{Token: yyDollar[2].token, Id: ast.Identifier{Token: yyDollar[3].token, Value: yyDollar[3].token.Literal}, Param: ast.Identifier{Token: yyDollar[6].token, Value: yyDollar[6].token.Literal}, BindingExpr: yyDollar[8].expr, BodyExpr: yyDollar[10].expr}
		}
	case 10:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser/parser.go.y:89
		{
			yyVAL.expr = ast.FunExpr{Token: yyDollar[1].token, Param: ast.Identifier{Token: yyDollar[2].token, Value: yyDollar[2].token.Literal}, BodyExpr: yyDollar[4].expr}
		}
	case 11:
		yyDollar = yyS[yypt-6 : yypt+1]
//line parser/parser.go.y:95
		{
			yyVAL.expr = ast.LetExpr{Token: yyDollar[1].token, Id: ast.Identifier{Token: yyDollar[2].token, Value: yyDollar[2].token.Literal}, BindingExpr: yyDollar[4].expr, BodyExpr: yyDollar[6].expr}
		}
	case 12:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser/parser.go.y:101
		{
			yyVAL.expr = ast.BinOpExpr{Token: yyDollar[2].token, Left: yyDollar[1].expr, Operator: token.LT, Right: yyDollar[3].expr}
		}
	case 13:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser/parser.go.y:105
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 14:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser/parser.go.y:111
		{
			yyVAL.expr = ast.BinOpExpr{Token: yyDollar[2].token, Left: yyDollar[1].expr, Operator: token.PLUS, Right: yyDollar[3].expr}
		}
	case 15:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser/parser.go.y:115
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 16:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser/parser.go.y:121
		{
			yyVAL.expr = ast.BinOpExpr{Token: yyDollar[2].token, Left: yyDollar[1].expr, Operator: token.ASTERISK, Right: yyDollar[3].expr}
		}
	case 17:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser/parser.go.y:125
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 18:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser/parser.go.y:131
		{
			yyVAL.expr = ast.AppExpr{Token: token.Token{Type: token.FUN, Literal: "("}, Function: yyDollar[1].expr, Argument: yyDollar[2].expr}
		}
	case 19:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser/parser.go.y:135
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 20:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser/parser.go.y:141
		{
			intValue, err := strconv.Atoi(yyDollar[1].token.Literal)
			if err != nil {
				yylex.(*LexerWrapper).Error(fmt.Sprintf("invalid integer value: %s", yyDollar[1].token.Literal))
				return 1
			}
			yyVAL.expr = ast.Integer{Token: yyDollar[1].token, Value: intValue}
		}
	case 21:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser/parser.go.y:150
		{
			yyVAL.expr = ast.Boolean{Token: yyDollar[1].token, Value: true}
		}
	case 22:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser/parser.go.y:154
		{
			yyVAL.expr = ast.Boolean{Token: yyDollar[1].token, Value: false}
		}
	case 23:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser/parser.go.y:158
		{
			yyVAL.expr = ast.Identifier{Token: yyDollar[1].token, Value: yyDollar[1].token.Literal}
		}
	case 24:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser/parser.go.y:162
		{
			yyVAL.expr = yyDollar[2].expr
		}
	case 25:
		yyDollar = yyS[yypt-6 : yypt+1]
//line parser/parser.go.y:168
		{
			yyVAL.expr = ast.IfExpr{Token: yyDollar[1].token, Condition: yyDollar[2].expr, Consequence: yyDollar[4].expr, Alternative: yyDollar[6].expr}
		}
	}
	goto yystack /* stack new state and value */
}
