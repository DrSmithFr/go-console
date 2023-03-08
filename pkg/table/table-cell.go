package table

type TableCellInterface interface {
	GetValue() string
	SetValue(value string) *TableCell

	GetColspan() int
	GetPadType() PaddingType

	getRowspan() int
}

type TableCell struct {
	Value   string
	Colspan int
	PadType PaddingType
	rowspan int
}

func NewTableCell(value string) *TableCell {
	t := new(TableCell)

	t.Value = value
	t.rowspan = 1
	t.Colspan = 1
	t.PadType = PadDefault

	return t
}

// Implement TableCellInterface

var _ TableCellInterface = (*TableCell)(nil)

func (t *TableCell) GetValue() string {
	return t.Value
}

func (t *TableCell) getRowspan() int {
	if t.rowspan == 0 {
		return 1
	}

	return t.rowspan
}

func (t *TableCell) GetColspan() int {
	if t.Colspan == 0 {
		return 1
	}

	return t.Colspan
}

func (t *TableCell) GetPadType() PaddingType {
	return t.PadType
}

// Implement TableCell fluent setters

func (t *TableCell) SetValue(value string) *TableCell {
	t.Value = value
	return t
}

func (t *TableCell) setRowspan(rowspan int) *TableCell {
	t.rowspan = rowspan
	return t
}

func (t *TableCell) SetColspan(colspan int) *TableCell {
	t.Colspan = colspan
	return t
}

func (t *TableCell) SetPadType(pad PaddingType) *TableCell {
	t.PadType = pad
	return t
}
