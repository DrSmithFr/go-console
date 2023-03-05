package table

type TableCellInterface interface {
	GetValue() string
	GetRowspan() int
	GetColspan() int
}

type TableCell struct {
	value   string
	rowspan int
	colspan int
}

func NewTableCell(value string) *TableCell {
	t := new(TableCell)

	t.value = value
	t.rowspan = 1
	t.colspan = 1

	return t
}

// Implement TableCellInterface

var _ TableCellInterface = (*TableCell)(nil)

func (t *TableCell) GetValue() string {
	return t.value
}

func (t *TableCell) GetRowspan() int {
	return t.rowspan
}

func (t *TableCell) GetColspan() int {
	return t.colspan
}

// Implement TableCell fluent setters

func (t *TableCell) SetValue(value string) *TableCell {
	t.value = value
	return t
}

func (t *TableCell) SetRowspan(rowspan int) *TableCell {
	t.rowspan = rowspan
	return t
}

func (t *TableCell) SetColspan(colspan int) *TableCell {
	t.colspan = colspan
	return t
}
