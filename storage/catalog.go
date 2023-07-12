package storage

import "fmt"

type Catalog struct {
	databases map[string]Database
}

func NewCatalog() *Catalog {
	return &Catalog{
		databases: make(map[string]Database),
	}
}

func (c *Catalog) GetDatabase(name string) (Database, error) {
	if database, ok := c.databases[name]; ok {
		return database, nil
	}

	return Database{}, fmt.Errorf("database %q not found", name)
}

func (c *Catalog) ListDatabases() ([]Database, error) {
	databases := make([]Database, 0, len(c.databases))

	for name := range c.databases {
		databases = append(databases, c.databases[name])
	}

	return databases, nil
}

func (c *Catalog) CreateDatabase(name string) (Database, error) {
	if _, ok := c.databases[name]; ok {
		return Database{}, fmt.Errorf("database %s already exist", name)
	}

	database := NewDatabase(name)
	c.databases[name] = *database

	return *database, nil
}

func (c *Catalog) DropDatabase(name string) error {
	if _, ok := c.databases[name]; ok {
		delete(c.databases, name)
		return nil
	}

	return fmt.Errorf("database %q not found", name)
}
