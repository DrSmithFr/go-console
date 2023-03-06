package table

type TableSeparatorInterface interface {
	IsSeparator() bool
}

type TableSeparator struct {
	TableCell
}

// TableSeparator constructors
func NewTableSeparator() *TableSeparator {
	return new(TableSeparator)
}

// Implement TableSeparatorInterface

func (t *TableSeparator) IsSeparator() bool {
	return true
}

// Implement TableCell fluent setters

var _ TableSeparatorInterface = (*TableSeparator)(nil)
var _ TableSeparatorInterface = (*TableSeparator)(nil)

func (t *TableSeparator) SetValue(value string) *TableSeparator {
	t.Value = value
	return t
}

func (t *TableSeparator) SetRowspan(rowspan int) *TableSeparator {
	t.Rowspan = rowspan
	return t
}

func (t *TableSeparator) SetColspan(colspan int) *TableSeparator {
	t.Colspan = colspan
	return t
}
