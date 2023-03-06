package table

type TableCellInterface interface {
	GetValue() string
	SetValue(value string) *TableCell

	GetRowspan() int
	GetColspan() int
	GetPadType() PaddingType
}

type TableCell struct {
	Value   string
	Rowspan int
	Colspan int
	PadType PaddingType
}

func NewTableCell(value string) *TableCell {
	t := new(TableCell)

	t.Value = value
	t.Rowspan = 1
	t.Colspan = 1
	t.PadType = PadDefault

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

func (t *TableCell) GetPadType() PaddingType {
	return t.PadType
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

func (t *TableCell) SetPadType(pad PaddingType) *TableCell {
	t.PadType = pad
	return t
}
