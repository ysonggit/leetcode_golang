func shiftGrid(grid [][]int, k int) [][]int {
	for i := 0; i < k; i++ {
		grid = shiftOnce(grid)
	}
	return grid
}
func shiftOnce(grid [][]int) [][]int {
	m, n := len(grid), len(grid[0])
	first_col := make([]int, m)
	first_col[0] = grid[m-1][n-1]
	for i := 1; i < m; i++ {
		first_col[i] = grid[i-1][n-1]
	}
	for i := 0; i < m; i++ {
		for j := n - 1; j > 0; j-- {
			grid[i][j] = grid[i][j-1]
		}
	}
	for i := 0; i < m; i++ {
		grid[i][0] = first_col[i]
	}
	return grid
}