package clitable

type ColumnOption func(c *Column)

func WithColumnAlignment(alignment Alignment) ColumnOption {
	return func(c *Column) {
		c.Alignment = alignment
	}
}

func WithColumnColor(color Color) ColumnOption {
	return func(c *Column) {
		c.Color = color
	}
}

type TableOption func(c *Table)

func WithTableBorders(style BorderStyle) TableOption {
	return func(t *Table) {
		t.bordersStyle = style
	}
}

func WithTableHeaders(headersEnabled bool) TableOption {
	return func(t *Table) {
		t.headersEnabled = headersEnabled
	}
}

func WithBordersColor(bordersColor Color) TableOption {
	return func(t *Table) {
		t.borderColor = bordersColor
	}
}


func WithTableHeadersColor(headersColor Color) TableOption {
	return func(t *Table) {
		t.headersColor = headersColor
	}
}
