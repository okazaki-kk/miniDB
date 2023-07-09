package datatype

import (
	"testing"

	"github.com/okazaki-kk/miniDB/internal/sql"
	"github.com/stretchr/testify/assert"
)

func TestInteger_Raw(t *testing.T) {
	expected := int64(10)
	b := NewInteger(expected)

	switch value := b.Raw().(type) {
	case int64:
		assert.Equal(t, expected, value)
	default:
		assert.Failf(t, "fail", "unexpected type %T", value)
	}
}

func TestInteger_String(t *testing.T) {
	i := NewInteger(10)
	assert.Equal(t, "10", i.String())
}

func TestInteger_DataType(t *testing.T) {
	i := NewInteger(10)
	assert.Equal(t, sql.Integer, i.DataType())
}
