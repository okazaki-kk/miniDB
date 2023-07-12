package storage

import "github.com/okazaki-kk/miniDB/internal/sql"

type Scheme map[string]Column

type Column struct {
	Position   uint8
	Name       string
	DataType   sql.DataType
	PrimaryKey bool
	Nullable   bool
}
