import "fmt"

func isPathCrossing(path string) bool {
	x, y := 0, 0
	visited := make(map[string]bool)
	key := fmt.Sprintf("%d_%d", x, y)
	visited[key] = true
	for _, step := range path {
		switch step {
		case 'N':
			y++
		case 'S':
			y--
		case 'E':
			x++
		default:
			x--
		}
		key = fmt.Sprintf("%d_%d", x, y)
		if _, ok := visited[key]; ok {
			return true
		}
		visited[key] = true
	}
	return false
}