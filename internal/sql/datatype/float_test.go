package datatype

import (
	"testing"

	"github.com/okazaki-kk/miniDB/internal/sql"
	"github.com/stretchr/testify/assert"
)

func TestFloat_Raw(t *testing.T) {
	expected := float64(10)
	b := NewFloat(expected)

	switch value := b.Raw().(type) {
	case float64:
		assert.Equal(t, expected, value)
	default:
		assert.Failf(t, "fail", "unexpected type %T", value)
	}
}

func TestFloat_String(t *testing.T) {
	f := NewFloat(10.0006)
	assert.Equal(t, "1.00006E+01", f.String())
}

func TestFloat_DataType(t *testing.T) {
	f := NewFloat(10)
	assert.Equal(t, sql.Float, f.DataType())
}
