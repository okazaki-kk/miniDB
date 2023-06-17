package lexer

import (
	"testing"

	"github.com/okazaki-kk/miniDB/internal/parser/token"
	"github.com/stretchr/testify/assert"
)

func TestLexer_NextToken(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input     string
		tokenType token.TokenType
		literal   string
	}{
		{
			input:     "id",
			tokenType: token.IDENT,
			literal:   "id",
		},
		{
			input:     "_id",
			tokenType: token.IDENT,
			literal:   "_id",
		},
		{
			input:     "i9",
			tokenType: token.IDENT,
			literal:   "i9",
		},
		{
			input:     "    i9  ",
			tokenType: token.IDENT,
			literal:   "i9",
		},
		{
			input:     ",",
			tokenType: token.COMMA,
			literal:   ",",
		},
		{
			input:     ";",
			tokenType: token.SEMICOLON,
			literal:   ";",
		},
		{
			input:     "(",
			tokenType: token.LPAREN,
			literal:   "(",
		},
		{
			input:     ")",
			tokenType: token.RPAREN,
			literal:   ")",
		},
		{
			input:     "=",
			tokenType: token.EQ,
			literal:   "=",
		},
		{
			input:     "<",
			tokenType: token.LT,
			literal:   "<",
		},
		{
			input:     ">",
			tokenType: token.GT,
			literal:   ">",
		},
		{
			input:     "!=",
			tokenType: token.NOT_EQ,
			literal:   "!=",
		},
		{
			input:     "!",
			tokenType: token.BANG,
			literal:   "!",
		},
		{
			input:     "AND",
			tokenType: token.AND,
			literal:   "AND",
		},
		{
			input:     "OR",
			tokenType: token.OR,
			literal:   "OR",
		},
		{
			input:     "NOT",
			tokenType: token.NOT,
			literal:   "NOT",
		},
		{
			input:     "+",
			tokenType: token.PLUS,
			literal:   "+",
		},
		{
			input:     "-",
			tokenType: token.MINUS,
			literal:   "-",
		},
		{
			input:     "*",
			tokenType: token.ASTERISK,
			literal:   "*",
		},
		{
			input:     "/",
			tokenType: token.SLASH,
			literal:   "/",
		},
		{
			input:     "10",
			tokenType: token.INT,
			literal:   "10",
		},
		{
			input:     "'value'",
			tokenType: token.TEXT,
			literal:   "value",
		},
		{
			input:     "true",
			tokenType: token.TRUE,
			literal:   "true",
		},
		{
			input:     "false",
			tokenType: token.FALSE,
			literal:   "false",
		},
		{
			input:     "CREATE",
			tokenType: token.CREATE,
			literal:   "CREATE",
		},
		{
			input:     "TABLE",
			tokenType: token.TABLE,
			literal:   "TABLE",
		},
		{
			input:     "DATABASE",
			tokenType: token.DATABASE,
			literal:   "DATABASE",
		},
		{
			input:     "DROP",
			tokenType: token.DROP,
			literal:   "DROP",
		},
		{
			input:     "SELECT",
			tokenType: token.SELECT,
			literal:   "SELECT",
		},
		{
			input:     "FROM",
			tokenType: token.FROM,
			literal:   "FROM",
		},
		{
			input:     "WHERE",
			tokenType: token.WHERE,
			literal:   token.WHERE,
		},
		{
			input:     "ORDER",
			tokenType: token.ORDER,
			literal:   "ORDER",
		},
		{
			input:     "BY",
			tokenType: token.BY,
			literal:   "BY",
		},
		{
			input:     "ASC",
			tokenType: token.ASC,
			literal:   "ASC",
		},
		{
			input:     "DESC",
			tokenType: token.DESC,
			literal:   "DESC",
		},
		{
			input:     "LIMIT",
			tokenType: token.LIMIT,
			literal:   "LIMIT",
		},
		{
			input:     "OFFSET",
			tokenType: token.OFFSET,
			literal:   "OFFSET",
		},
		{
			input:     "INSERT",
			tokenType: token.INSERT,
			literal:   "INSERT",
		},
		{
			input:     "INTO",
			tokenType: token.INTO,
			literal:   "INTO",
		},
		{
			input:     "VALUES",
			tokenType: token.VALUES,
			literal:   "VALUES",
		},
		{
			input:     "UPDATE",
			tokenType: token.UPDATE,
			literal:   "UPDATE",
		},
		{
			input:     "SET",
			tokenType: token.SET,
			literal:   "SET",
		},
		{
			input:     "DELETE",
			tokenType: token.DELETE,
			literal:   "DELETE",
		},
		{
			input:     "DEFAULT",
			tokenType: token.DEFAULT,
			literal:   "DEFAULT",
		},
		{
			input:     "NULL",
			tokenType: token.NULL,
			literal:   "NULL",
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.input, func(t *testing.T) {
			t.Parallel()

			lx := New(test.input)
			nextToken := lx.NextToken()
			assert.Equal(t, test.tokenType, nextToken.Type)
			assert.Equal(t, test.literal, nextToken.Literal)
		})
	}
}
