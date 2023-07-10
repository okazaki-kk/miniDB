package storage

import (
	"testing"

	"github.com/okazaki-kk/miniDB/internal/sql"
	"github.com/stretchr/testify/assert"
)

func TestDatabase(t *testing.T) {
	t.Parallel()

	t.Run("empty table", func(t *testing.T) {
		dbName := "playground"
		database := NewDatabase(dbName)
		tables := database.ListTables()
		assert.Empty(t, tables)
	})

	t.Run("return tables", func(t *testing.T) {
		database := NewDatabase("playground")
		scheme := sql.Scheme{
			"id": sql.Column{
				Position:   0,
				Name:       "id",
				DataType:   sql.Integer,
				PrimaryKey: true,
				Nullable:   false,
				Default:    nil,
			},
		}

		users, err := database.CreateTable("users", scheme)
		assert.NoError(t, err)

		tickets, err := database.CreateTable("tickets", scheme)
		assert.NoError(t, err)

		expected := []Table{
			users,
			tickets,
		}

		tables := database.ListTables()

		assert.ElementsMatch(t, expected, tables)
	})

	t.Run("get table", func(t *testing.T) {
		scheme := sql.Scheme{
			"id": sql.Column{
				Position:   0,
				Name:       "id",
				DataType:   sql.Integer,
				PrimaryKey: true,
				Nullable:   false,
				Default:    nil,
			},
		}

		database := NewDatabase("playground")
		expected, err := database.CreateTable("users", scheme)
		assert.NoError(t, err)

		table, err := database.GetTable("users")
		assert.NoError(t, err)
		assert.Equal(t, expected, table)
	})

	t.Run("create table", func(t *testing.T) {
		scheme := sql.Scheme{
			"id": sql.Column{
				Position:   0,
				Name:       "id",
				DataType:   sql.Integer,
				PrimaryKey: true,
				Nullable:   false,
				Default:    nil,
			},
		}

		database := NewDatabase("playground")
		table, err := database.CreateTable("users", scheme)
		assert.NoError(t, err)
		assert.Equal(t, scheme, table.Scheme())
	})

	t.Run("drop table", func(t *testing.T) {
		scheme := sql.Scheme{
			"id": sql.Column{
				Position:   0,
				Name:       "id",
				DataType:   sql.Integer,
				PrimaryKey: true,
				Nullable:   false,
				Default:    nil,
			},
		}

		database := NewDatabase("playground")
		_, err := database.CreateTable("users", scheme)
		assert.NoError(t, err)

		err = database.DropTable("users")
		assert.NoError(t, err)

		tables := database.ListTables()
		assert.Empty(t, tables)
	})
}
