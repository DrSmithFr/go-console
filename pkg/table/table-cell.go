package table

type TableCellInterface interface {
	GetValue() string
	GetRowspan() int
	GetColspan() int
}

type TableCell struct {
	Value   string
	Rowspan int
	Colspan int
}

func NewTableCell(value string) *TableCell {
	t := new(TableCell)

	t.Value = value
	t.Rowspan = 1
	t.Colspan = 1

	return t
}

// Implement TableCellInterface

var _ TableCellInterface = (*TableCell)(nil)

func (t *TableCell) GetValue() string {
	return t.Value
}

func (t *TableCell) GetRowspan() int {
	if t.Rowspan == 0 {
		return 1
	}

	return t.Rowspan
}

func (t *TableCell) GetColspan() int {
	if t.Colspan == 0 {
		return 1
	}

	return t.Colspan
}

// Implement TableCell fluent setters

func (t *TableCell) SetValue(value string) *TableCell {
	t.Value = value
	return t
}

func (t *TableCell) SetRowspan(rowspan int) *TableCell {
	t.Rowspan = rowspan
	return t
}

func (t *TableCell) SetColspan(colspan int) *TableCell {
	t.Colspan = colspan
	return t
}
