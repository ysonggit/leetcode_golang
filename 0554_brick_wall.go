func leastBricks(wall [][]int) int {
	width_counts := make(map[int]int, 0)
	max_spaces := 0
	for _, row := range wall {
		width_sofar := 0
		for j := 0; j < len(row)-1; j++ { // last position is end of all bricks, doesn't count
			width_sofar += row[j]
			width_counts[width_sofar]++
			if max_spaces < width_counts[width_sofar] {
				max_spaces = width_counts[width_sofar]
			}
		}
	}
	return len(wall) - max_spaces
}