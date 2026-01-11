package engine

import (
	"context"
	"errors"
	"minidb/internal/storage"
	"minidb/internal/table"
	"strings"

	"vitess.io/vitess/go/vt/sqlparser"
)

type Engine struct {
	tables map[string]*table.Table
	wal    *storage.WAL
}

func NewEngine(walPath string) (*Engine, error) {
	wal, err := storage.Open(walPath)
	if err != nil {
		return nil, err
	}
	return &Engine{
		tables: map[string]*table.Table{},
		wal:    wal,
	}, nil
}

func (e *Engine) Execute(ctx context.Context, sql string) (any, error) {
	parser := sqlparser.NewTestParser()
	stmt, err := parser.Parse(sql)
	if err != nil {
		return nil, err
	}

	switch s := stmt.(type) {
	case *sqlparser.CreateTable:
		return nil, e.createTable(s)
	case *sqlparser.Insert:
		return nil, e.insert(s)
	case *sqlparser.Select:
		return e.selectRows(s)
	default:
		return nil, errors.New("unsupported SQL")
	}
}

func (e *Engine) createTable(s *sqlparser.CreateTable) error {
	var cols []table.Column
	for _, c := range s.TableSpec.Columns {
		col := table.Column{
			Name: c.Name.String(),
			Type: table.StringType,
		}
		if strings.Contains(strings.ToLower(c.Type.Type), "int") {
			col.Type = table.IntType
		}
		cols = append(cols, col)
	}
	schema := table.NewSchema(cols)
	e.tables[s.Table.Name.String()] = table.New(s.Table.Name.String(), schema)
	return nil
}
