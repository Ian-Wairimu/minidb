package engine

import (
	"context"
	"errors"
	"minidb/internal/table"

	"vitess.io/vitess/go/vt/sqlparser"
)

func (e *Engine) insert(ctx context.Context, s *sqlparser.Insert) error {
	tableName := sqlparser.String(s.Table)
	t := e.tables[tableName]
	if t == nil {
		return errors.New("table not found")
	}

	row := table.Row{}

	if values, ok := s.Rows.(sqlparser.Values); ok && len(values) > 0 {
		for _, v := range values[0] {
			switch expr := v.(type) {
			case *sqlparser.Literal:
				// Literal is the correct type in newer Vitess versions
				row.Values = append(row.Values, expr.Val)
			default:
				// Fallback: convert to string representation
				row.Values = append(row.Values, sqlparser.String(expr))
			}
		}
	}

	e.wal.Append(sqlparser.String(s))
	return t.Insert(row)
}
