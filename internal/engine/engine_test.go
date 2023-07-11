package engine

import (
	"testing"

	"github.com/okazaki-kk/miniDB/storage"
	"github.com/stretchr/testify/assert"
)

func TestCreateDatabase(t *testing.T) {
	catalog := storage.NewCatalog()
	engine := New(*catalog)
	message, err := engine.CreateDatabase("test")
	assert.NoError(t, err)
	assert.Equal(t, "create database test\n", message)

	message, err = engine.CreateDatabase("test")
	assert.Error(t, err)
	assert.Equal(t, "database \"test\" already exist", err.Error())
	assert.Equal(t, "", message)
}
