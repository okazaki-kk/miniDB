package storage

import (
	"io"
	"testing"

	"github.com/okazaki-kk/miniDB/internal/sql"
	"github.com/stretchr/testify/assert"
)

func TestTable(t *testing.T) {
	t.Parallel()

	tableName := "users"
	scheme := sql.Scheme{
		"id": sql.Column{
			Position:   0,
			Name:       "id",
			DataType:   sql.Integer,
			PrimaryKey: true,
			Nullable:   false,
			Default:    nil,
		},
		"name": sql.Column{
			Position:   0,
			Name:       "name",
			DataType:   sql.Text,
			PrimaryKey: false,
			Nullable:   false,
			Default:    nil,
		},
	}

	database := NewDatabase("playground")
	table, err := database.CreateTable(tableName, scheme)

	assert.NoError(t, err)
	assert.Equal(t, tableName, table.Name())
	assert.Equal(t, scheme, table.Scheme())
	assert.Equal(t, scheme["id"], table.PrimaryKey())

	iter, err := table.Scan()
	assert.NoError(t, err)

	row, err := iter.Next()
	assert.ErrorIs(t, io.EOF, err)
	assert.Nil(t, row)

	err = iter.Close()
	assert.NoError(t, err)
}
