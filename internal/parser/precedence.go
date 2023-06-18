package parser

import "github.com/okazaki-kk/miniDB/internal/parser/token"

const (
	_ int = iota
	LOWEST
	EOF
	Ident

	// Special chars
	Comma      // ,
	Semicolon  // ;
	OpenParen  // (
	CloseParen // )

	// Comparison operators
	Equal              // =
	LessThan           // <
	GreaterThan        // >
	NotEqual           // !=
	LessThanOrEqual    // <=
	GreaterThanOrEqual // >=

	// Logical operators
	And
	Or
	Not

	// Mathematical operators
	Add // +
	Sub // -
	Mul // *
	Div // /
	Mod // %
	Pow // ^

	// Types
	Integer
	Float
	Text
	Boolean
	Null

	// Keywords
	Create
	Table
	Database
	Drop
	Select
	As
	From
	Where
	Order
	By
	Asc
	Desc
	Limit
	Offset
	Insert
	Into
	Values
	Update
	Set
	Delete
	Default
	Primary
	Key
)

var precedences = map[token.TokenType]int{
	token.OR:       1,
	token.AND:      2,
	token.EQ:       3,
	token.NOT_EQ:   3,
	token.LT:       3,
	token.GT:       3,
	token.PLUS:     4,
	token.MINUS:    4,
	token.ASTERISK: 5,
	token.SLASH:    5,
}
