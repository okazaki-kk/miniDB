package token

import "strings"

type TokenType string

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers + literals
	IDENT = "IDENT" // add, foobar, x, y, ...
	INT   = "INT"   // 1343456

	// Operators
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"

	LT = "<"
	GT = ">"

	EQ     = "="
	NOT_EQ = "!="

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"

	SELECT   = "SELECT"
	FROM     = "FROM"
	AND      = "AND"
	OR       = "OR"
	NOT      = "NOT"
	TEXT     = "TEXT"
	CREATE   = "CREATE"
	TABLE    = "TABLE"
	INSERT   = "INSERT"
	VALUES   = "VALUES"
	UPDATE   = "UPDATE"
	SET      = "SET"
	DELETE   = "DELETE"
	WHERE    = "WHERE"
	DATABASE = "DATABASE"
	DROP     = "DROP"
	ORDER    = "ORDER"
	BY       = "BY"
	DESC     = "DESC"
	ASC      = "ASC"
	LIMIT    = "LIMIT"
	OFFSET   = "OFFSET"
	INTO     = "INTO"
	DEFAULT  = "DEFAULT"
	NULL     = "NULL"
)

type Token struct {
	Type    TokenType
	Literal string
}

var keywords = map[string]TokenType{
	"TRUE":     TRUE,
	"FALSE":    FALSE,
	"SELECT":   SELECT,
	"FROM":     FROM,
	"AND":      AND,
	"OR":       OR,
	"NOT":      NOT,
	"TEXT":     TEXT,
	"CREATE":   CREATE,
	"TABLE":    TABLE,
	"INSERT":   INSERT,
	"VALUES":   VALUES,
	"UPDATE":   UPDATE,
	"SET":      SET,
	"DELETE":   DELETE,
	"WHERE":    WHERE,
	"DATABASE": DATABASE,
	"DROP":     DROP,
	"ORDER":    ORDER,
	"BY":       BY,
	"DESC":     DESC,
	"ASC":      ASC,
	"LIMIT":    LIMIT,
	"OFFSET":   OFFSET,
	"INTO":     INTO,
	"DEFAULT":  DEFAULT,
	"NULL":     NULL,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	if tok, ok := keywords[strings.ToUpper(ident)]; ok {
		return tok
	}
	if tok, ok := keywords[strings.ToLower(ident)]; ok {
		return tok
	}
	return IDENT
}
