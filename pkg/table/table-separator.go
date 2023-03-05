package table

type TableSeparatorInterface interface {
	IsSeparator() bool
}

type TableSeparator struct {
	TableCell
}

// Implement TableSeparatorInterface

func (t *TableSeparator) IsSeparator() bool {
	return true
}

// Implement TableCell fluent setters

var _ TableSeparatorInterface = (*TableSeparator)(nil)
var _ TableSeparatorInterface = (*TableSeparator)(nil)

func (t *TableSeparator) SetValue(value string) *TableSeparator {
	t.value = value
	return t
}

func (t *TableSeparator) SetRowspan(rowspan int) *TableSeparator {
	t.rowspan = rowspan
	return t
}

func (t *TableSeparator) SetColspan(colspan int) *TableSeparator {
	t.colspan = colspan
	return t
}
