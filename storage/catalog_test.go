package storage

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCatalog(t *testing.T) {
	c := NewCatalog()
	db, err := c.CreateDatabase("test")
	assert.NoError(t, err)
	assert.Equal(t, "test", db.Name())

	db, err = c.CreateDatabase("test")
	assert.Error(t, err)
	assert.Equal(t, "database \"test\" already exist", err.Error())

	db, err = c.GetDatabase("test")
	assert.NoError(t, err)
	assert.Equal(t, "test", db.Name())

	_, err = c.CreateDatabase("test2")
	assert.NoError(t, err)

	dbs, err := c.ListDatabases()
	assert.NoError(t, err)
	assert.Equal(t, 2, len(dbs))

	err = c.DropDatabase("test")
	assert.NoError(t, err)

	dbs, err = c.ListDatabases()
	assert.NoError(t, err)
	assert.Equal(t, 1, len(dbs))
	assert.Equal(t, "test2", dbs[0].Name())
}
