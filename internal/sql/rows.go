package sql

import "io"

type Value interface {
	Raw() any
	String() string
	DataType() DataType
}

type Row []Value

type RowIter interface {
	Next() (Row, error)
	Close() error
}

type SliceRowsIter struct {
	rows  []Row
	index int
}

func (i *SliceRowsIter) Next() (Row, error) {
	if i.index > len(i.rows)-1 {
		return nil, io.EOF
	}

	row := i.rows[i.index]
	i.index++

	return row, nil
}

func (i *SliceRowsIter) Close() error {
	i.rows = nil
	return nil
}
