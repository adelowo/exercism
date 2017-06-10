package pascal

const testVersion = 1

func Triangle(levels int) [][]int {
	if levels == 1 {
		return [][]int{{1}}
	}

	top := Triangle(levels - 1)
	currentLine := make([]int, levels)
	for i := 0; i < levels; i++ {
		if i == 0 || i == levels-1 {
			currentLine[i] = 1
		} else {
			currentLine[i] = top[levels-2][i-1] + top[levels-2][i]
		}
	}
	return append(top, currentLine)
}
