package clitable

import (
	"fmt"
	"io"
	"os"
	"strings"
)

const colorReset = "\033[0m"

type BorderStyle int

var BordersDefault BorderStyle = 0
var BordersMinimal BorderStyle = 1
var BordersNone BorderStyle = 2

type Table struct {
	rows           []*row
	columns        []*Column
	columnsMap     map[string]int
	headersEnabled bool
	bordersStyle   BorderStyle
	headersColor   Color
	borderColor    Color
}

func NewTable(columns []*Column, options ...TableOption) (*Table, error) {
	t := &Table{
		columns:        make([]*Column, len(columns)),
		columnsMap:     make(map[string]int),
		headersEnabled: true,
		headersColor:   ColorNone,
		bordersStyle:   BordersDefault,
	}

	for i, c := range columns {
		_, ok := t.columnsMap[c.Name]
		if ok {
			return nil, fmt.Errorf("column %s already exists", c.Name)
		}
		t.columnsMap[c.Name] = i
		t.columns[i] = c
	}
	for _, opt := range options {
		opt(t)
	}
	return t, nil
}

func (t *Table) AddRow(values ...interface{}) error {
	if len(values) != len(t.columns) {
		return fmt.Errorf("values count don't match columns count")
	}
	r := &row{
		height: 1,
		values: make([][]string, len(t.columns)),
	}
	for i, v := range values {
		vStr := strings.Trim(strings.ReplaceAll(toString(v), "\t", "  "), "\n")
		vLines := strings.Split(vStr, "\n")
		if len(vLines) > r.height {
			r.height = len(vLines)
		}
		for _, line := range vLines {
			if len(line) > t.columns[i].width {
				t.columns[i].width = len(line)
			}
		}
		r.values[i] = vLines
	}
	t.rows = append(t.rows, r)
	return nil
}

func (t *Table) getColumnNames() []string {
	names := make([]string, len(t.columns))
	for i, col := range t.columns {
		names[i] = col.Name
	}
	return names
}

func (t *Table) printBorderRow(w io.Writer, edge bool) error {
	line := ""
	switch t.bordersStyle {
	case BordersDefault:
		lines := make([]string, len(t.columns))
		for i, c := range t.columns {
			lines[i] = strings.Repeat("-", c.getEffectiveWidth()+2)
		}
		line = "+" + strings.Join(lines, "+") + "+\n"
	case BordersMinimal:
		if edge {
			return nil
		}
		lines := make([]string, len(t.columns))
		for i, c := range t.columns {
			add := 1
			if i > 0 && i < len(lines)-1 {
				add = 2
			}
			lines[i] = strings.Repeat("-", c.getEffectiveWidth()+add)
		}
		line = strings.Join(lines, "+") + "\n"
	}

	if line != "" {
		if t.borderColor != ColorNone {
			line = string(t.borderColor) + line + colorReset
		}
		_, err := w.Write([]byte(line))
		if err != nil {
			return fmt.Errorf("failed to print border row: %v", err)
		}
	}
	return nil
}

func (t *Table) getBorderBegin() string {
	result := ""
	switch t.bordersStyle {
	case BordersDefault:
		result = "| "
	}
	if t.borderColor != ColorNone && result != "" {
		result = string(t.borderColor) + result + colorReset
	}
	return result
}

func (t *Table) getBorderEnd() string {
	result := ""
	switch t.bordersStyle {
	case BordersDefault:
		result = " |\n"
	case BordersMinimal:
		result = "\n"
	default:
		result = "\n"
	}
	if t.borderColor != ColorNone && result != "" {
		result = string(t.borderColor) + result + colorReset
	}
	return result
}

func (t *Table) getBorderDelimiter() string {
	result := ""
	switch t.bordersStyle {
	case BordersDefault:
		result = " | "
	case BordersMinimal:
		result = " | "
	default:
		result = " "
	}
	if t.borderColor != ColorNone && result != "" {
		result = string(t.borderColor) + result + colorReset
	}
	return result
}

func (t *Table) cellLine(line string, c *Column) string {
	switch c.Alignment {
	case AlignRight:
		return strings.Repeat(" ", c.getEffectiveWidth()-len(line)) + line
	case AlignCenter:
		right := (c.getEffectiveWidth() - len(line)) / 2
		left := c.getEffectiveWidth() - len(line) - right
		return strings.Repeat(" ", left) + line + strings.Repeat(" ", right)
	default:
		return line + strings.Repeat(" ", c.getEffectiveWidth()-len(line))
	}
}

func (t *Table) printRow(w io.Writer, lines []string, header bool) error {
	cells := make([]string, len(lines))
	for i, line := range lines {
		col := t.columns[i]
		cells[i] = t.cellLine(line, col)
		if !header && col.Color != ColorNone {
			cells[i] = string(col.Color) + cells[i] + colorReset
		}
		if header && t.headersColor != ColorNone {
			cells[i] = string(t.headersColor) + cells[i] + colorReset
		}
	}
	line := t.getBorderBegin() + strings.Join(cells, t.getBorderDelimiter()) + t.getBorderEnd()

	if _, err := w.Write([]byte(line)); err != nil {
		return fmt.Errorf("failed to print row: %v", err)
	}
	return nil
}

func (t *Table) Write(w io.Writer) error {
	if t.headersEnabled {
		if err := t.printBorderRow(w, true); err != nil {
			return err
		}
		columnNames := t.getColumnNames()
		if err := t.printRow(w, columnNames, true); err != nil {
			return err
		}
	}

	if err := t.printBorderRow(w, !t.headersEnabled); err != nil {
		return err
	}
	for _, r := range t.rows {
		for j := 0; j < r.height; j++ {
			cells := r.getLine(j)
			if err := t.printRow(w, cells, false); err != nil {
				return err
			}
		}
	}
	if err := t.printBorderRow(w, true); err != nil {
		return err
	}
	return nil
}

func (t *Table) Print() error {
	return t.Write(os.Stdout)
}
