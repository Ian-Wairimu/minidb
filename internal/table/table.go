package table

import (
	"sync"

	"minidb/internal/index"
)

type Table struct {
	Name    string
	Schema  *Schema
	Rows    []Row
	Indexes map[string]index.Index
	mu      sync.RWMutex
}

func New(name string, schema *Schema) *Table {
	t := &Table{
		Name:    name,
		Schema:  schema,
		Indexes: map[string]index.Index{},
	}
	for _, c := range schema.Columns {
		if c.Primary || c.Unique {
			t.Indexes[c.Name] = index.NewHashIndex(true)
		}
	}
	return t
}

func (t *Table) Insert(r Row) error {
	t.mu.Lock()
	defer t.mu.Unlock()

	id := len(t.Rows)
	for col, idx := range t.Indexes {
		v := r.Values[t.Schema.Index[col]]
		if err := idx.Insert(v, id); err != nil {
			return err
		}
	}
	t.Rows = append(t.Rows, r)
	return nil
}
