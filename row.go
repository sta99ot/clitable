package clitable

type row struct {
	height int
	values [][]string
}

func (r *row) getValueLine(c int, line int) string {
	if len(r.values[c]) <= line {
		return ""
	}
	return r.values[c][line]
}

func (r *row) getLine(line int) []string {
	result := make([]string, len(r.values))
	for i := 0; i < len(result); i++ {
		result[i] = r.getValueLine(i, line)
	}
	return result
}
