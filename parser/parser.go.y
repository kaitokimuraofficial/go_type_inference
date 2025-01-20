%{
package parser

import (
	"fmt"
	"go_type_inference/ast"
	"go_type_inference/lexer"
	"go_type_inference/token"
	"strconv"
)
%}

%union{
    expr ast.Expr
    statement ast.Stmt
    token token.Token
}

%type<statement> statement
%type<expr> expr
%type<expr> ltexpr
%type<expr> ifexpr
%type<expr> letexpr
%type<expr> funexpr
%type<expr> letrecexpr
%type<expr> pexpr
%type<expr> mexpr
%type<expr> appexpr
%type<expr> aexpr

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
        $$ = ast.ExprStmt{Expr: $1}
        yylex.(*LexerWrapper).Result = $$
    }
    | LET IDENT ASSIGN expr
    {
        $$ = ast.DeclStmt{Decl: ast.LetDecl{Id: ast.Identifier{Value: $2.Literal}, Expr: $4}}
        yylex.(*LexerWrapper).Result = $$
    }
    | LET REC IDENT ASSIGN FUN IDENT RARROW expr
    {
        $$ = ast.DeclStmt{Decl: ast.RecDecl{Id: ast.Identifier{Value: $3.Literal}, Param: ast.Identifier{Value: $6.Literal}, BodyExpr: $8}}
        yylex.(*LexerWrapper).Result = $$
    }

expr
    : ltexpr
    {
        $$ = $1
    }
    | ifexpr
    {
        $$ = $1
    }
    | funexpr
    {
        $$ = $1
    }
    | letexpr
    {
        $$ = $1
    }
    | letrecexpr
    {
        $$ = $1
    }

ltexpr
    : pexpr
    {
        $$ = $1
    }
    | pexpr LT pexpr
    {
        $$ = ast.BinOpExpr{Op: $2.Type, Left: $1, Right: $3}
    }

ifexpr
    : IF expr THEN expr ELSE expr
    {
        $$ = ast.IfExpr{Condition: $2, Consequence: $4, Alternative: $6}
    }

letexpr
    : LET IDENT ASSIGN expr IN expr
    {
        $$ = ast.LetExpr{Id: ast.Identifier{Value: $2.Literal}, BindingExpr: $4, BodyExpr: $6}
    }

funexpr
    : FUN IDENT RARROW expr
    {
        $$ = ast.FunExpr{Param: ast.Identifier{Value: $2.Literal}, BodyExpr: $4}
    }

letrecexpr
    : LET REC IDENT ASSIGN FUN IDENT RARROW expr IN expr
    {
        $$ = ast.LetRecExpr{Id: ast.Identifier{Value: $3.Literal}, Param: ast.Identifier{Value: $6.Literal}, BindingExpr: $8, BodyExpr: $10}
    }

pexpr
    : mexpr
    {
        $$ = $1
    }
    | pexpr PLUS mexpr
    {
        $$ = ast.BinOpExpr{Op: $2.Type, Left: $1, Right: $3}
    }

mexpr
    : appexpr
    {
        $$ = $1
    }
    | mexpr ASTERISK appexpr
    {
        $$ = ast.BinOpExpr{Op: $2.Type, Left: $1, Right: $3}
    }

appexpr
    : aexpr
    {
        $$ = $1
    }
    | appexpr aexpr
    {
        $$ = ast.AppExpr{Function: $1, Argument: $2}
    }

aexpr
    : INT
    {
        intValue, err := strconv.Atoi($1.Literal)
        if err != nil {
            yylex.(*LexerWrapper).Error(fmt.Sprintf("invalid integer value: %s", $1.Literal))
            return -1
        }
        $$ = ast.Integer{Value: intValue}
    }
    | TRUE
    {
        $$ = ast.Boolean{Value: true}
    }
    | FALSE
    {
        $$ = ast.Boolean{Value: false}
    }
    | IDENT
    {
        $$ = ast.Identifier{Value: $1.Literal}
    }
    | LPAREN expr RPAREN
    {
        $$ = $2
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
