package table

type TableColumnInterface interface {
	SetCell(cell TableCellInterface) *TableColumn
	GetCell() TableCellInterface
}

type TableColumn struct {
	cell TableCellInterface
}

// constructors

func NewTableColumn() *TableColumn {
	return &TableColumn{
		cell: NewTableCell(""),
	}
}

func MakeColumnFromString(cells string) *TableColumn {
	return NewTableColumn().setColumnFromString(cells)
}

// Implements TableColumnInterface.

var _ TableColumnInterface = (*TableColumn)(nil)

func (t *TableColumn) SetCell(cell TableCellInterface) *TableColumn {
	t.cell = cell
	return t
}

func (t *TableColumn) GetCell() TableCellInterface {
	return t.cell
}

// Data injections Helpers

func (t *TableColumn) setColumnFromString(row string) *TableColumn {
	t.cell = NewTableCell(row)
	return t
}
