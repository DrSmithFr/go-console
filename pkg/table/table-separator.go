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
