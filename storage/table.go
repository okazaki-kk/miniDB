package storage

import (
	"fmt"
	"io"

	"github.com/okazaki-kk/miniDB/internal/sql"
)

type Table struct {
	name       string
	rows       map[int64]sql.Row
	scheme     sql.Scheme
	keys       []int64
	primaryKey sql.Column
}

func NewTable(name string, scheme sql.Scheme) *Table {
	var pk sql.Column
	for col := range scheme {
		if scheme[col].PrimaryKey {
			pk = scheme[col]
			break
		}
	}

	return &Table{
		name:       name,
		scheme:     scheme,
		primaryKey: pk,
		rows:       make(map[int64]sql.Row),
	}
}

func (t *Table) Name() string {
	return t.name
}

func (t *Table) PrimaryKey() sql.Column {
	return t.primaryKey
}

func (t *Table) Scheme() sql.Scheme {
	return t.scheme
}

func (t *Table) Sequence() sql.Sequence {
	return nil
}

func (t *Table) Scan() (sql.RowIter, error) {
	rows := make([]sql.Row, 0, len(t.keys))

	for _, key := range t.keys {
		rows = append(rows, t.rows[key])
	}

	i := &iter{rows: rows}

	return i, nil
}

func (t *Table) Insert(key int64, row sql.Row) error {
	if _, ok := t.rows[key]; ok {
		return fmt.Errorf("duplicate primary key %d", key)
	}

	t.rows[key] = row
	t.keys = append(t.keys, key)

	return nil
}

func (t *Table) Delete(key int64) error {
	if _, ok := t.rows[key]; !ok {
		return fmt.Errorf("key %d not found", key)
	}

	for index := range t.keys {
		if t.keys[index] == key {
			t.keys = append(t.keys[:index], t.keys[index+1:]...)
			break
		}
	}

	delete(t.rows, key)

	return nil
}

func (t *Table) Update(key int64, row sql.Row) error {
	if _, ok := t.rows[key]; !ok {
		return fmt.Errorf("key %d not found", key)
	}

	t.rows[key] = row

	return nil
}

type iter struct {
	index int
	rows  []sql.Row
}

func (i *iter) Next() (sql.Row, error) {
	if i.index > len(i.rows)-1 {
		return nil, io.EOF
	}

	row := i.rows[i.index]
	i.index++

	return row, nil
}

func (i *iter) Close() error {
	i.rows = nil
	return nil
}
