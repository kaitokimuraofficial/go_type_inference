%{
package parser

import (
    "fmt"
    "strconv"
    "go_type_inference/ast"
    "go_type_inference/lexer"
    "go_type_inference/token"
)
%}

%union{
    statement ast.Stmt
    expr ast.Expr
    token token.Token
}

%type<statement> statement
%type<expr> expr
%type<expr> letrecexpr
%type<expr> funexpr
%type<expr> letexpr
%type<expr> ltexpr
%type<expr> pexpr
%type<expr> mexpr
%type<expr> appexpr
%type<expr> aexpr
%type<expr> ifexpr

%token<token> IDENT
%token<token> INT TRUE FALSE
%token<token> LPAREN RPAREN
%token<token> IF THEN ELSE
%token<token> LT PLUS ASTERISK
%token<token> LET ASSIGN IN
%token<token> RARROW FUN
%token<token> REC

%%

statement
    : expr
    {
        $$ = ast.Statement{Expr: $1}
        yylex.(*LexerWrapper).Result = $$
    }
    | LET IDENT ASSIGN expr
    {
        $$ = ast.Declaration{Id: ast.Identifier{Token: $2, Value: $2.Literal}, Expr: $4}
        yylex.(*LexerWrapper).Result = $$
    }
    | LET REC IDENT ASSIGN FUN IDENT RARROW expr
    {
        $$ = ast.RecDeclaration{Id: ast.Identifier{Token: $3, Value: $3.Literal}, Param: ast.Identifier{Token: $6, Value: $6.Literal}, BodyExpr: $8}
        yylex.(*LexerWrapper).Result = $$
    }

expr
    : ifexpr
    {
        $$ = $1
    }
    | letexpr
    {
        $$ = $1
    }
    | ltexpr
    {
        $$ = $1
    }
    | funexpr
    {
        $$ = $1
    }
    | letrecexpr
    {
        $$ = $1
    }

letrecexpr
    : LET REC IDENT ASSIGN FUN IDENT RARROW expr IN expr
    {
        $$ = ast.LetRecExpr{Token: $2, Id: ast.Identifier{Token: $3, Value: $3.Literal}, Param: ast.Identifier{Token: $6, Value: $6.Literal}, BindingExpr: $8, BodyExpr: $10}
    }

funexpr
    : FUN IDENT RARROW expr
    {
        $$ = ast.FunExpr{Token: $1, Param: ast.Identifier{Token: $2, Value: $2.Literal}, BodyExpr: $4}
    }

letexpr
    : LET IDENT ASSIGN expr IN expr
    {
        $$ = ast.LetExpr{Token: $1, Identifier: ast.Identifier{Token: $2, Value: $2.Literal}, BindingExpr: $4, BodyExpr: $6}
    }

ltexpr
    : pexpr LT pexpr
    {
        $$ = ast.BinOpExpr{Token: $2, Left: $1, Operator: token.LT, Right: $3}
    }
    | pexpr
    {
        $$ = $1
    }

pexpr
    : pexpr PLUS mexpr
    {
        $$ = ast.BinOpExpr{Token: $2, Left: $1, Operator: token.PLUS, Right: $3}
    }
    | mexpr
    {
        $$ = $1
    }

mexpr
    : mexpr ASTERISK appexpr
    {
        $$ = ast.BinOpExpr{Token: $2, Left: $1, Operator: token.ASTERISK, Right: $3}
    }
    | appexpr
    {
        $$ = $1
    }

appexpr
    : appexpr aexpr
    {
        $$ = ast.AppExpr{Token: token.Token{Type: token.FUN, Literal: "("}, Function: $1, Argument: $2}
    }
    | aexpr
    {
        $$ = $1
    }

aexpr
    : INT
    {
        intValue, err := strconv.Atoi($1.Literal)
        if err != nil {
            yylex.(*LexerWrapper).Error(fmt.Sprintf("invalid integer value: %s", $1.Literal))
            return 1
        }
        $$ = ast.Integer{Token: $1, Value: intValue}
    }
    | TRUE
    {
        $$ = ast.Boolean{Token: $1, Value: true}
    }
    | FALSE
    {
        $$ = ast.Boolean{Token: $1, Value: false}
    }
    | IDENT
    {
        $$ = ast.Identifier{Token: $1, Value: $1.Literal}
    }
    | LPAREN expr RPAREN
    {
        $$ = $2
    }

ifexpr
    : IF expr THEN expr ELSE expr
    {
        $$ = ast.IfExpr{Token: $1, Condition: $2, Consequence: $4, Alternative: $6}
    }

%%

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
