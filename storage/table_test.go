package storage

import (
	"io"
	"testing"

	"github.com/okazaki-kk/miniDB/internal/sql"
	"github.com/okazaki-kk/miniDB/internal/sql/datatype"
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

func TestTable_Insert(t *testing.T) {
	t.Run("insert without errors", func(t *testing.T) {
		scheme := sql.Scheme{
			"name": sql.Column{
				Position:   0,
				Name:       "name",
				DataType:   sql.Text,
				PrimaryKey: false,
				Nullable:   false,
				Default:    nil,
			},
			"id": sql.Column{
				Position:   1,
				Name:       "id",
				DataType:   sql.Integer,
				PrimaryKey: true,
				Nullable:   false,
				Default:    nil,
			},
		}

		key := int64(1)
		expected := sql.Row{
			datatype.NewText("Max"),
			datatype.NewInteger(key),
		}

		database := NewDatabase("playground")
		table, err := database.CreateTable("users", scheme)
		assert.NoError(t, err)

		err = table.Insert(key, expected)
		assert.NoError(t, err)

		iter, err := table.Scan()
		assert.NoError(t, err)

		row, err := iter.Next()
		assert.NoError(t, err)
		assert.Equal(t, expected, row)

		row, err = iter.Next()
		assert.ErrorIs(t, io.EOF, err)
		assert.Nil(t, row)
	})
}

func TestTable_Update(t *testing.T) {
	t.Run("update without errors", func(t *testing.T) {

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
				Position:   1,
				Name:       "name",
				DataType:   sql.Text,
				PrimaryKey: false,
				Nullable:   false,
				Default:    nil,
			},
		}

		row := sql.Row{
			datatype.NewInteger(1),
			datatype.NewText("Max"),
		}

		updated := sql.Row{
			datatype.NewInteger(1),
			datatype.NewText("Tom"),
		}

		database := NewDatabase("playground")
		table, err := database.CreateTable("users", scheme)
		assert.NoError(t, err)

		err = table.Insert(1, row)
		assert.NoError(t, err)

		err = table.Update(1, updated)
		assert.NoError(t, err)

		iter, err := table.Scan()
		assert.NoError(t, err)

		actual, err := iter.Next()
		assert.NoError(t, err)
		assert.Equal(t, updated, actual)

		row, err = iter.Next()
		assert.ErrorIs(t, io.EOF, err)
		assert.Nil(t, row)
	})
}

func TestTable_Delete(t *testing.T) {
	t.Parallel()

	t.Run("deletes one row", func(t *testing.T) {
		t.Parallel()

		scheme := sql.Scheme{
			"id": sql.Column{
				Position:   1,
				Name:       "id",
				DataType:   sql.Integer,
				PrimaryKey: true,
				Nullable:   false,
				Default:    nil,
			},
		}

		key := int64(1)
		expected := sql.Row{
			datatype.NewInteger(key),
		}

		database := NewDatabase("playground")
		table, err := database.CreateTable("users", scheme)
		assert.NoError(t, err)

		err = table.Insert(key, expected)
		assert.NoError(t, err)

		err = table.Delete(key)
		assert.NoError(t, err)

		iter, err := table.Scan()
		assert.NoError(t, err)

		row, err := iter.Next()
		assert.ErrorIs(t, io.EOF, err)
		assert.Nil(t, row)
	})

	t.Run("deletes all rows", func(t *testing.T) {
		t.Parallel()

		scheme := sql.Scheme{
			"id": sql.Column{
				Position:   1,
				Name:       "id",
				DataType:   sql.Integer,
				PrimaryKey: true,
				Nullable:   false,
				Default:    nil,
			},
		}

		rows := []struct {
			key int64
			row sql.Row
		}{
			{
				key: 1,
				row: sql.Row{datatype.NewInteger(1)},
			},
			{
				key: 2,
				row: sql.Row{datatype.NewInteger(2)},
			},
			{
				key: 3,
				row: sql.Row{datatype.NewInteger(3)},
			},
		}

		database := NewDatabase("playground")
		table, err := database.CreateTable("users", scheme)
		assert.NoError(t, err)

		for _, r := range rows {
			err = table.Insert(r.key, r.row)
			assert.NoError(t, err)
		}

		for _, r := range rows {
			err = table.Delete(r.key)
			assert.NoError(t, err)
		}

		iter, err := table.Scan()
		assert.NoError(t, err)

		row, err := iter.Next()
		assert.ErrorIs(t, io.EOF, err)
		assert.Nil(t, row)
	})
}
