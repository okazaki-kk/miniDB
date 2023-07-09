package datatype

import (
	"testing"

	"github.com/okazaki-kk/miniDB/internal/sql"
	"github.com/stretchr/testify/assert"
)

func TestBoolean_Raw(t *testing.T) {
	fn := func(expected bool) {
		b := NewBoolean(expected)

		switch value := b.Raw().(type) {
		case bool:
			assert.Equal(t, expected, value)
		default:
			assert.Failf(t, "fail", "unexpected type %T", value)
		}
	}

	fn(true)
	fn(false)
}

func TestBoolean_String(t *testing.T) {
	b := NewBoolean(true)
	assert.Equal(t, "true", b.String())

	b = NewBoolean(false)
	assert.Equal(t, "false", b.String())
}

func TestBoolean_DataType(t *testing.T) {
	b := NewBoolean(true)
	assert.Equal(t, sql.Boolean, b.DataType())
}
