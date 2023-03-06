package table

type TableColumnInterface interface {
	SetCell(cell TableCellInterface) *TableColumn
	GetCell() TableCellInterface
}

type TableColumn struct {
	Cell TableCellInterface
}

// constructors

func NewTableColumn() *TableColumn {
	return &TableColumn{
		Cell: NewTableCell(""),
	}
}

func MakeColumnFromString(cells string) *TableColumn {
	return NewTableColumn().setColumnFromString(cells)
}

// Implements TableColumnInterface.

var _ TableColumnInterface = (*TableColumn)(nil)

func (t *TableColumn) SetCell(cell TableCellInterface) *TableColumn {
	t.Cell = cell
	return t
}

func (t *TableColumn) GetCell() TableCellInterface {
	return t.Cell
}

// Data injections Helpers

func (t *TableColumn) setColumnFromString(row string) *TableColumn {
	t.Cell = NewTableCell(row)
	return t
}
