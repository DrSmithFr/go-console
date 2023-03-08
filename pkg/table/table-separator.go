package table

type TableSeparatorInterface interface {
	IsSeparator() bool
	IsDouble() bool
}

type TableSeparator struct {
	TableCell
	double bool
}

// TableSeparator constructors
func NewTableSeparator() *TableSeparator {
	return &TableSeparator{
		double: false,
	}
}

func NewTableSeparatorDouble() *TableSeparator {
	return &TableSeparator{
		double: true,
	}
}

// Implement TableSeparatorInterface

var _ TableSeparatorInterface = (*TableSeparator)(nil)

func (t *TableSeparator) IsSeparator() bool {
	return true
}

func (t *TableSeparator) IsDouble() bool {
	return t.double
}
