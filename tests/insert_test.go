package tests

import (
	"testing"

	"minidb/internal/table"
)

func TestInsert(t *testing.T) {
	s := table.NewSchema([]table.Column{
		{Name: "id", Primary: true},
	})
	tbl := table.New("users", s)

	if err := tbl.Insert(table.Row{Values: []any{1}}); err != nil {
		t.Fatal(err)
	}
}
