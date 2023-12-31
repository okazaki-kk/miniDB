package engine

import (
	"fmt"

	"github.com/okazaki-kk/miniDB/internal/parser"
	"github.com/okazaki-kk/miniDB/internal/parser/ast"
	"github.com/okazaki-kk/miniDB/internal/parser/lexer"
	"github.com/okazaki-kk/miniDB/storage"
)

type Engine struct {
	parser  parser.Parser
	catalog storage.Catalog
}

func New(catalog storage.Catalog) *Engine {
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
	case *ast.CreateTableStatement:
		return e.CreateTable(database, stmt.Table, stmt.Columns)
	default:
		return "", nil
	}
}

func (e *Engine) CreateDatabase(name string) (string, error) {
	db, err := e.catalog.CreateDatabase(name)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("create database %s\n", db.Name()), err
}

func (e *Engine) CreateTable(database string, tableName string, columns []ast.Column) (string, error) {
	db, err := e.catalog.GetDatabase(database)
	if err != nil {
		return "", err
	}

	scheme, err := storage.CreateTableScheme(columns)
	if err != nil {
		return "", err
	}

	_, err = db.CreateTable(tableName, scheme)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("create table %s\n", tableName), err
}
