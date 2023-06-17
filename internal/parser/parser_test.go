package parser

import (
	"testing"

	"github.com/okazaki-kk/miniDB/internal/parser/ast"
	"github.com/okazaki-kk/miniDB/internal/parser/lexer"
	"github.com/stretchr/testify/assert"
)

func TestParser_Select(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input string
		stmt  ast.Statement
	}{
		{
			input: "SELECT * FROM users;",
			stmt: &ast.SelectStatement{
				Result: []ast.ResultStatement{
					{
						Expr: &ast.AsteriskExpr{},
					},
				},
				From: &ast.FromStatement{
					Table: "users",
				},
			},
		},
		{
			input: "SELECT id, name FROM users;",
			stmt: &ast.SelectStatement{
				Result: []ast.ResultStatement{
					{
						Expr: &ast.IdentExpr{
							Name: "id",
						},
					},
					{
						Expr: &ast.IdentExpr{
							Name: "name",
						},
					},
				},
				From: &ast.FromStatement{
					Table: "users",
				},
			},
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.input, func(t *testing.T) {
			t.Parallel()

			p := New(lexer.New(test.input))
			stmts, err := p.Parse()
			assert.NoError(t, err)
			assert.Equal(t, test.stmt, stmts)
		})
	}
}
