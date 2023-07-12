package engine

import (
	"testing"

	"github.com/okazaki-kk/miniDB/internal/parser/ast"
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

func TestCreateTable(t *testing.T) {
	catalog := storage.NewCatalog()
	engine := New(*catalog)
	_, err := engine.CreateDatabase("test")
	assert.NoError(t, err)

	message, err := engine.CreateTable("test", "test", []ast.Column{
		{Name: "id", Type: "INT", PrimaryKey: true},
		{Name: "name", Type: "TEXT", Nullable: true},
	})
	assert.NoError(t, err)
	assert.Equal(t, "create table test\n", message)

	message, err = engine.CreateTable("test", "test", []ast.Column{
		{Name: "id", Type: "INT", PrimaryKey: true},
		{Name: "name", Type: "TEXT", Nullable: true},
	})
	assert.Error(t, err)
	assert.Equal(t, "table \"test\" already exist", err.Error())
	assert.Equal(t, "", message)

	message, err = engine.CreateTable("test", "test2", []ast.Column{
		{Name: "id", Type: "INT", PrimaryKey: true},
		{Name: "name", Type: "TEXT", Nullable: true},
	})
	assert.NoError(t, err)
	assert.Equal(t, "create table test2\n", message)

	message, err = engine.CreateTable("test", "test3", []ast.Column{
		{Name: "id", Type: "INT", PrimaryKey: true},
		{Name: "name", Type: "TEXT", Nullable: true},
	})
	assert.NoError(t, err)
	assert.Equal(t, "create table test3\n", message)
}
