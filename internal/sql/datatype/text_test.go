package datatype

import (
	"testing"

	"github.com/okazaki-kk/miniDB/internal/sql"
	"github.com/stretchr/testify/assert"
)

func TestText_Raw(t *testing.T) {
	t.Parallel()

	expected := "xyz"
	s := NewText(expected)

	switch value := s.Raw().(type) {
	case string:
		assert.Equal(t, expected, value)
	default:
		assert.Failf(t, "fail", "unexpected type %T", value)
	}
}

func TestText_String(t *testing.T) {
	t.Parallel()

	value := "xyz"
	text := NewText(value)
	assert.Equal(t, value, text.String())
}

func TestText_DataType(t *testing.T) {
	t.Parallel()

	n := NewText("xyz")
	assert.Equal(t, sql.Text, n.DataType())
}
