package sql

type Catalog interface {
	GetDatabase(name string) (Database, error)
	ListDatabases() ([]Database, error)
	CreateDatabase(name string) (Database, error)
	DropDatabase(name string) error
}

type Database interface {
	Name() string
	GetTable(name string) (Table, error)
	ListTables() []Table
	CreateTable(name string, scheme Scheme) (Table, error)
	DropTable(name string) error
}

type Table interface {
	Name() string
	Scheme() Scheme
	PrimaryKey() Column
	Sequence() Sequence
	Scan() (RowIter, error)
	Insert(key int64, row Row) error
	Delete(key int64) error
	Update(key int64, row Row) error
}

type Sequence interface {
	Next() int64
}

type Scheme map[string]Column

type Column struct {
	Position   uint8
	Name       string
	DataType   DataType
	PrimaryKey bool
	Nullable   bool
	Default    Value
}

type CompareType int

const (
	Less    CompareType = -1
	Equal   CompareType = 0
	Greater CompareType = 1
)

type Value interface {
	Raw() any
	String() string
	DataType() DataType
}
