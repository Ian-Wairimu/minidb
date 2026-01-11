package engine

import (
	"minidb/internal/table"

	"vitess.io/vitess/go/vt/sqlparser"
)

func (e *Engine) selectRows(s *sqlparser.Select) ([]table.Row, error) {
	t := e.tables[s.From[0].(*sqlparser.AliasedTableExpr).Expr.(sqlparser.TableName).Name.String()]
	return t.Rows, nil
}
