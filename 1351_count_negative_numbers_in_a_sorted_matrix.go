func countNegatives(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	// search from bottom left corner to upper right corner
	r, c, neg := m-1, 0, 0
	for ok := true; ok; ok = (r >= 0 && c < n) {
		if r >= 0 && c < n && grid[r][c] < 0 {
			neg += n - c
			r--
		} else {
			if c < n && r >= 0 {
				c++
			} else if c == n-1 && r >= 0 {
				r--
				c = 0
			}
		}
	}
	return neg
}