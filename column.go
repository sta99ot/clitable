package clitable

type Alignment int

var AlignLeft Alignment = 0
var AlignCenter Alignment = 1
var AlignRight Alignment = 2

type Column struct {
	id        string
	name      string
	alignment Alignment
	color     Color
	maxWidth  int
	width     int
}

func NewColumn(id string, name string, options ...ColumnOption) *Column {
	col := &Column{
		id:       id,
		name:     name,
		maxWidth: -1,
		color:    ColorNone,
		width:    len(name),
	}
	for _, opt := range options {
		opt(col)
	}
	return col
}

func (c *Column) getEffectiveWidth() int {
	if c.maxWidth > 0 && c.width <= c.maxWidth || c.maxWidth <= 0 {
		return c.width
	}
	return c.maxWidth
}
