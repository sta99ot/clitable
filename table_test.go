package clitable

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func fillTable(t *testing.T, table *Table) {
	require.NoError(t, table.AddRow(1, "One", "Hello\nthere"))
	require.NoError(t, table.AddRow(2, "Two", "I'm your friend"))
	require.NoError(t, table.AddRow(3, "Three", []string{"one", "two", "three"}))
	require.NoError(t, table.AddRow(4, "Four", map[string]interface{}{"field_one": "one", "field_two": 2}))

}

var cols = []*Column{
	NewColumn("id", "ID", WithLeftAlignment(), WithColumnColor(ColorCyan)),
	NewColumn("name", "NAME", WithCenterAlignment(), WithColumnColor(ColorRed)),
	NewColumn("desc", "DESCRIPTION", WithLeftAlignment(), WithColumnMaxWidth(16)),
}

func TestTables(t *testing.T) {
	println("Default borders:")
	table, _ := NewTable(
		cols,
		WithDefaultBorders(),
		WithHeadersColor(ColorYellow),
		WithBordersColor(ColorGray),
	)
	fillTable(t, table)
	require.NoError(t, table.Print())
	println("Minimal borders:")

	table, _ = NewTable(
		cols,
		WithMinimalBorders(),
		WithHeadersColor(ColorYellow),
		WithBordersColor(ColorGray),
	)
	fillTable(t, table)
	require.NoError(t, table.Print())

	println("No borders:")
	table, _ = NewTable(
		cols,
		WithoutBorders(),
		WithHeadersColor(ColorYellow),
		WithBordersColor(ColorGray),
	)
	fillTable(t, table)
	require.NoError(t, table.Print())
}
