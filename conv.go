package clitable

import (
	"fmt"
	"strings"
)

func splitLines(text string, maxWidth int) []string {
	var result []string
	for _, l := range strings.Split(text, "\n") {
		if len(l) <= maxWidth || maxWidth <= 0 {
			result = append(result, l)
			continue
		}
		for j := 0; j < len(l); j += maxWidth {
			length := maxWidth
			if len(l) <= j+maxWidth {
				length = len(l) - j
			}
			result = append(result, l[j:j+length])
		}
	}
	return result
}

func toString(v interface{}) string {
	return fmt.Sprintf("%v", v)
}
