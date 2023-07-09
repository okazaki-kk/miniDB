package parser

import (
	"testing"

	"github.com/okazaki-kk/miniDB/internal/parser/ast"
	"github.com/okazaki-kk/miniDB/internal/parser/lexer"
	"github.com/okazaki-kk/miniDB/internal/parser/token"
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
		{
			input: "SELECT age FROM customers WHERE id = 10;",
			stmt: &ast.SelectStatement{
				Result: []ast.ResultStatement{
					{
						Expr: &ast.IdentExpr{
							Name: "age",
						},
					},
				},
				From: &ast.FromStatement{
					Table: "customers",
				},
				Where: &ast.WhereStatement{
					Expr: &ast.ConditionExpr{
						Left:     &ast.IdentExpr{Name: "id"},
						Operator: token.EQ,
						Right: &ast.ScalarExpr{
							Type:    token.INT,
							Literal: "10",
						},
					},
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

func TestParser_CreateTable(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input string
		stmt  ast.Statement
	}{
		{
			input: "CREATE TABLE users (id INT, name TEXT);",
			stmt: &ast.CreateTableStatement{
				Table: "users",
				Columns: []ast.Column{
					{
						Name: "id",
						Type: token.INT,
					},
					{
						Name: "name",
						Type: token.TEXT,
					},
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
