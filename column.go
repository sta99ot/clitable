package clitable

type Alignment int

var AlignLeft Alignment = 0
var AlignCenter Alignment = 1
var AlignRight Alignment = 2

type Column struct {
	Name      string
	Alignment Alignment
	Color     Color
	MaxWidth  int
	width     int
}

func NewColumn(name string, options ...ColumnOption) *Column {
	col := &Column{
		Name:     name,
		MaxWidth: -1,
		Color:    ColorNone,
		width:    len(name),
	}
	for _, opt := range options {
		opt(col)
	}
	return col
}

func (c *Column) getEffectiveWidth() int {
	if c.MaxWidth > 0 && c.width <= c.MaxWidth || c.MaxWidth <= 0 {
		return c.width
	}
	return c.MaxWidth
}
