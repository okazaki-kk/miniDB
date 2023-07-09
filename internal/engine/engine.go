package engine

import (
	"fmt"

	"github.com/okazaki-kk/miniDB/internal/parser"
	"github.com/okazaki-kk/miniDB/internal/parser/ast"
	"github.com/okazaki-kk/miniDB/internal/parser/lexer"
	"github.com/okazaki-kk/miniDB/internal/sql"
)

type Engine struct {
	parser  parser.Parser
	catalog sql.Catalog
}

func New(catalog sql.Catalog) *Engine {
	return &Engine{catalog: catalog}
}

func (e *Engine) Exec(database, input string) (string, error) {
	e.parser = *parser.New(lexer.New(input))
	stmt, err := e.parser.Parse()
	if err != nil {
		return "", err
	}

	switch stmt := stmt.(type) {
	case *ast.CreateDatabaseStatement:
		return e.CreateDatabase(stmt.Database)
	default:
		return "", nil
	}
}

func (e *Engine) CreateDatabase(name string) (string, error) {
	db, err := e.catalog.CreateDatabase(name)
	return fmt.Sprintf("create database %s\n", db.Name()), err
}
