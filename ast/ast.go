package ast

import "go_type_inference/token"

type Node interface {
	node()
}

type Expr interface {
	Node
	exprNode()
}

type Stmt interface {
	Node
	stmtNode()
}

type Decl interface {
	Node
	declNode()
}

// ----------------------------------------------------------------------------
// Expressions

type (
	Integer struct {
		Value int
	}

	Boolean struct {
		Value bool
	}

	Identifier struct {
		Value string
	}

	BinOpExpr struct {
		Op    token.Type
		Left  Expr
		Right Expr
	}

	IfExpr struct {
		Condition   Expr
		Consequence Expr
		Alternative Expr
	}

	LetExpr struct {
		Id          Identifier
		BindingExpr Expr
		BodyExpr    Expr
	}

	FunExpr struct {
		Param    Identifier
		BodyExpr Expr
	}

	AppExpr struct {
		Function Expr
		Argument Expr
	}

	LetRecExpr struct {
		Id          Identifier
		Param       Identifier
		BindingExpr Expr
		BodyExpr    Expr
	}
)

func (*Integer) node()    {}
func (*Boolean) node()    {}
func (*Identifier) node() {}
func (*BinOpExpr) node()  {}
func (*IfExpr) node()     {}
func (*LetExpr) node()    {}
func (*FunExpr) node()    {}
func (*AppExpr) node()    {}
func (*LetRecExpr) node() {}

// exprNode() ensures that only expression nodes can be
// assigned to an Expr
func (*Integer) exprNode()    {}
func (*Boolean) exprNode()    {}
func (*Identifier) exprNode() {}
func (*BinOpExpr) exprNode()  {}
func (*IfExpr) exprNode()     {}
func (*LetExpr) exprNode()    {}
func (*FunExpr) exprNode()    {}
func (*AppExpr) exprNode()    {}
func (*LetRecExpr) exprNode() {}

// ----------------------------------------------------------------------------
// Statements

type (
	DeclStmt struct {
		Decl Decl
	}

	ExprStmt struct {
		Expr Expr
	}
)

func (*DeclStmt) node() {}
func (*ExprStmt) node() {}

// stmtNode() ensures that only statement nodes can be
// assigned to a Stmt.
func (*DeclStmt) stmtNode() {}
func (*ExprStmt) stmtNode() {}

// ----------------------------------------------------------------------------
// Declarations

type (
	LetDecl struct {
		Id   Identifier
		Expr Expr
	}

	RecDecl struct {
		Id       Identifier
		Param    Identifier
		BodyExpr Expr
	}
)

func (*LetDecl) node() {}
func (*RecDecl) node() {}

// declNode() ensures that only declaration nodes can be
// assigned to a Decl.
func (*LetDecl) declNode() {}
func (*RecDecl) declNode() {}
