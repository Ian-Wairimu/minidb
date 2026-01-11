package table

type DataType int

const (
	IntType DataType = iota
	StringType
)

type Column struct {
	Name    string
	Type    DataType
	Primary bool
	Unique  bool
}

type Schema struct {
	Columns []Column
	Index   map[string]int
	PK      string
}

func NewSchema(cols []Column) *Schema {
	idx := make(map[string]int)
	var pk string
	for i, c := range cols {
		idx[c.Name] = i
		if c.Primary {
			pk = c.Name
		}
	}
	return &Schema{Columns: cols, Index: idx, PK: pk}
}
