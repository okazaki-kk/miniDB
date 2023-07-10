package sql

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
