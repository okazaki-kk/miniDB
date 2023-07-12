package storage

import (
	"fmt"

	"github.com/okazaki-kk/miniDB/internal/parser/ast"
	"github.com/okazaki-kk/miniDB/internal/parser/token"
	"github.com/okazaki-kk/miniDB/internal/sql"
)

type Scheme map[string]Column

type Column struct {
	Position   uint8
	Name       string
	DataType   sql.DataType
	PrimaryKey bool
	Nullable   bool
}

func CreateTableScheme(columns []ast.Column) (Scheme, error) {
	primaryKeys := 0
	scheme := make(Scheme, len(columns))

	for i := range columns {
		column, err := createSchemeColumn(uint8(i), columns[i])
		if err != nil {
			return nil, err
		}

		if column.PrimaryKey {
			primaryKeys++
		}

		scheme[column.Name] = column
	}

	if primaryKeys == 0 {
		return nil, fmt.Errorf("primary key is required")
	}

	if primaryKeys > 1 {
		return nil, fmt.Errorf("multiple primary keys are not allowed")
	}

	return scheme, nil
}

func createSchemeColumn(position uint8, column ast.Column) (Column, error) {
	var dataType sql.DataType

	switch column.Type {
	case token.INT:
		dataType = sql.Integer
	case token.TEXT:
		dataType = sql.Text
	case token.TRUE, token.FALSE:
		dataType = sql.Boolean
	default:
		return Column{}, fmt.Errorf("unexpected column type: %q", column.Type)
	}

	return Column{
		Position:   position,
		Name:       column.Name,
		DataType:   dataType,
		PrimaryKey: column.PrimaryKey,
		Nullable:   column.Nullable,
	}, nil
}
