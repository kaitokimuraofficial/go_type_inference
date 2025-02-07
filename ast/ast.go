package ast

import "go_type_inference/token"

type (
	Node interface {
		node()
	}

	Expr interface {
		Node
		exprNode()
	}

	Stmt interface {
		Node
		stmtNode()
	}

	Decl interface {
		Node
		declNode()
	}
)

// Expressions
type (
	Integer struct {
		Value int
	}

	Boolean struct {
		Value bool
	}

	Ident struct {
		Value string
	}

	BinOpExpr struct {
		Op    token.Type
		Left  Expr
		Right Expr
	}

	IfExpr struct {
		Cond Expr
		Cons Expr
		Alt  Expr
	}

	LetExpr struct {
		Id   Ident
		Bind Expr
		Body Expr
	}

	FunExpr struct {
		Param Ident
		Body  Expr
	}

	AppExpr struct {
		Func Expr
		Arg  Expr
	}

	LetRecExpr struct {
		Id    Ident
		Param Ident
		Bind  Expr
		Body  Expr
	}
)

func (Integer) node()    {}
func (Boolean) node()    {}
func (Ident) node()      {}
func (BinOpExpr) node()  {}
func (IfExpr) node()     {}
func (LetExpr) node()    {}
func (FunExpr) node()    {}
func (AppExpr) node()    {}
func (LetRecExpr) node() {}

func (Integer) exprNode()    {}
func (Boolean) exprNode()    {}
func (Ident) exprNode()      {}
func (BinOpExpr) exprNode()  {}
func (IfExpr) exprNode()     {}
func (LetExpr) exprNode()    {}
func (FunExpr) exprNode()    {}
func (AppExpr) exprNode()    {}
func (LetRecExpr) exprNode() {}

// Statements
type (
	DeclStmt struct {
		Decl Decl
	}

	ExprStmt struct {
		Expr Expr
	}
)

func (DeclStmt) node() {}
func (ExprStmt) node() {}

func (DeclStmt) stmtNode() {}
func (ExprStmt) stmtNode() {}

// Declarations
type (
	LetDecl struct {
		Id   Ident
		Expr Expr
	}

	RecDecl struct {
		Id    Ident
		Param Ident
		Body  Expr
	}
)

func (LetDecl) node() {}
func (RecDecl) node() {}

func (LetDecl) declNode() {}
func (RecDecl) declNode() {}
