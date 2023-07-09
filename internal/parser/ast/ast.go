// Package ast declares the types used to represent syntax trees for the NanoDB's SQL dialect.
package ast

import (
	"github.com/okazaki-kk/miniDB/internal/parser/token"
)

// Node represents AST-node of the syntax tree for SQL query.
type Node interface{}

// Statement represents syntax tree node of SQL statement (like: SELECT).
type Statement interface {
	Node
	statementNode()
}

// Expression represents syntax tree node of SQL expression (like: id < 10 AND id > 5).
type Expression interface {
	Node
	expressionNode()
}

// SelectStatement node represents a SELECT statement.
type SelectStatement struct {
	Result  []ResultStatement
	From    *FromStatement
	Where   *WhereStatement
	OrderBy *OrderByStatement
	Limit   *LimitStatement
	Offset  *OffsetStatement
}

// ResultStatement node represents a returning expression in a SELECT statement.
type ResultStatement struct {
	Expr Expression
}

// FromStatement node represents a FROM statement.
type FromStatement struct {
	Table string
}

type WhereStatement struct {
	Expr Expression
}

type CreateTableStatement struct {
	Table   string
	Columns []Column
}

// Column node represents a table column definition.
type Column struct {
	Name       string
	Type       token.TokenType
	Default    Expression
	Nullable   bool
	PrimaryKey bool
}

type OrderByStatement struct {
	Column    string
	Direction token.TokenType
}

type LimitStatement struct {
	Value Expression
}

type OffsetStatement struct {
	Value Expression
}

func (s *SelectStatement) statementNode()      {}
func (s *ResultStatement) statementNode()      {}
func (s *FromStatement) statementNode()        {}
func (s *WhereStatement) statementNode()       {}
func (s *CreateTableStatement) statementNode() {}
func (s *OrderByStatement) statementNode()     {}
func (s *LimitStatement) statementNode()       {}
func (s *OffsetStatement) statementNode()      {}

// IdentExpr node represents an identifier.
type IdentExpr struct {
	Name string
}

// ScalarExpr node represents a literal of basic type.
type ScalarExpr struct {
	Type    token.TokenType
	Literal string
}

type ConditionExpr struct {
	Left     Expression
	Operator token.TokenType
	Right    Expression
}

// AsteriskExpr node represents asterisk at `SELECT *` expression.
type AsteriskExpr struct{}

func (e *IdentExpr) expressionNode()     {}
func (e *ScalarExpr) expressionNode()    {}
func (e *AsteriskExpr) expressionNode()  {}
func (e *ConditionExpr) expressionNode() {}
