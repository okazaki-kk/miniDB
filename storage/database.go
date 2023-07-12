package storage

import "fmt"

type Database struct {
	name   string
	tables map[string]Table
}

func NewDatabase(name string) *Database {
	return &Database{name: name, tables: make(map[string]Table)}
}

func (d *Database) Name() string {
	return d.name
}

func (d *Database) ListTables() []Table {
	tables := make([]Table, 0, len(d.tables))

	for _, t := range d.tables {
		tables = append(tables, t)
	}

	return tables
}

func (d Database) GetTable(name string) (Table, error) {
	if table, ok := d.tables[name]; ok {
		return table, nil
	}

	return Table{}, fmt.Errorf("table %q not found", name)
}

func (d *Database) CreateTable(name string, scheme Scheme) (Table, error) {
	if _, ok := d.tables[name]; ok {
		return Table{}, fmt.Errorf("table %q already exist", name)
	}

	table := NewTable(name, scheme)
	d.tables[name] = *table

	return *table, nil
}

func (d *Database) DropTable(name string) error {
	if _, ok := d.tables[name]; !ok {
		return fmt.Errorf("table %s not found", name)
	}

	delete(d.tables, name)

	return nil
}
