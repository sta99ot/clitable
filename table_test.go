package clitable

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func fillTable(t *testing.T, table *Table) {
	require.NoError(t, table.AddRow(1, "Ho", "Hello\n there"))
	require.NoError(t, table.AddRow(2, "Two", "Longer message"))
	require.NoError(t, table.AddRow(2, "Two", []string{"one", "two", "three"}))
	require.NoError(t, table.AddRow(2, "Two", map[string]interface{}{"field_one": "one", "field_two": 2}))

}

var cols = []*Column{
	NewColumn("id", WithColumnAlignment(AlignLeft), WithColumnColor(ColorCyan)),
	NewColumn("name", WithColumnAlignment(AlignRight), WithColumnColor(ColorGray)),
	NewColumn("desc", WithColumnAlignment(AlignLeft), WithColumnColor(ColorYellow)),
}

func TestTables(t *testing.T) {
	println("Default borders:")
	tableDefault, _ := NewTable(
		cols,
		WithTableBorders(BordersDefault),
		WithTableHeadersColor(ColorGreen),
		WithBordersColor(ColorRed),
	)
	fillTable(t, tableDefault)
	require.NoError(t, tableDefault.Print())
	println("Minimal borders:")

	tableMinimal, _ := NewTable(
		cols,
		WithTableBorders(BordersMinimal),
		WithTableHeadersColor(ColorGreen),
		WithBordersColor(ColorRed),
	)
	fillTable(t, tableMinimal)
	require.NoError(t, tableMinimal.Print())

	println("No borders:")
	tableNone, _ := NewTable(
		cols,
		WithTableBorders(BordersNone),
		WithTableHeadersColor(ColorGreen),
		WithBordersColor(ColorRed),
	)
	fillTable(t, tableNone)
	require.NoError(t, tableNone.Print())
}

