package clitable

type ColumnOption func(c *Column)

func WithLeftAlignment() ColumnOption {
	return func(c *Column) {
		c.alignment = AlignLeft
	}
}

func WithCenterAlignment() ColumnOption {
	return func(c *Column) {
		c.alignment = AlignCenter
	}
}

func WithRightAlignment() ColumnOption {
	return func(c *Column) {
		c.alignment = AlignRight
	}
}

func WithColumnColor(color Color) ColumnOption {
	return func(c *Column) {
		c.color = color
	}
}

func WithColumnMaxWidth(maxWidth int) ColumnOption {
	return func(c *Column) {
		c.maxWidth = maxWidth
	}
}

func WithColumnMinWidth(minWidth int) ColumnOption {
	return func(c *Column) {
		c.width = minWidth
	}
}

type TableOption func(c *Table)

func WithDefaultBorders() TableOption {
	return func(t *Table) {
		t.bordersStyle = BordersDefault
	}
}

func WithMinimalBorders() TableOption {
	return func(t *Table) {
		t.bordersStyle = BordersMinimal
	}
}

func WithoutBorders() TableOption {
	return func(t *Table) {
		t.bordersStyle = BordersNone
	}
}

func WithoutTableHeaders() TableOption {
	return func(t *Table) {
		t.headersEnabled = false
	}
}

func WithBordersColor(bordersColor Color) TableOption {
	return func(t *Table) {
		t.borderColor = bordersColor
	}
}

func WithHeadersColor(headersColor Color) TableOption {
	return func(t *Table) {
		t.headersColor = headersColor
	}
}
