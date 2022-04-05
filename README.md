# clitable

Generate simple text-based tables in your terminal

- Accepts primitive data types and interfaces which provides `String() string` functions
- Left/Right/Center alignment
- Colors!
- Minimum and Maximum column width
- Multi-line rows

**Example**
```go
package main

import (
	"github.com/sta99ot/clitable"
)

func main() {
	table, _ := clitable.NewTable(
		[]*clitable.Column{
			clitable.NewColumn("id", "ID", clitable.WithRightAlignment(), clitable.WithColumnColor(clitable.ColorCyan)),
			clitable.NewColumn("name", "NAME", clitable.WithCenterAlignment(), clitable.WithColumnColor(clitable.ColorRed)),
			clitable.NewColumn("desc", "DESCRIPTION", clitable.WithLeftAlignment(), clitable.WithColumnMaxWidth(16)),
		},
		clitable.WithMinimalBorders(),
		clitable.WithHeadersColor(clitable.ColorGreen),
		clitable.WithBordersColor(clitable.ColorGray),
	)
	_ = table.AddRow(1, "One", "Hello\nthere")
	_ = table.AddRow(2, "Two", "I'm your friend")
	_ = table.Print()
}
```

```
ID |  NAME | DESCRIPTION     
---+-------+-----------------
1  |  One  | Hello           
   |       | there           
2  |  Two  | I'm your friend 
```

**Future**

- Support for structure/map reflection on table
- Better line cell text splitting ig it doesn't fit into `maxWidth`