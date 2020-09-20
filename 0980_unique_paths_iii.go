func uniquePathsIII(grid [][]int) int {
	// find start coord
	x, y, zeros := 0, 0, 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 0 {
				zeros++
			}
			if grid[i][j] == 1 {
				x, y = i, j
			}
		}
	}
	return dfs(&grid, x, y, zeros)
}

func dfs(grid *[][]int, x int, y int, zeros int) int {
	if x < 0 || y < 0 || x == len(*grid) || y == len((*grid)[0]) || (*grid)[x][y] == -1 {
		return 0
	}
	if (*grid)[x][y] == 2 {
		if zeros == -1 {
			return 1
		} else {
			return 0
		}
	}
	(*grid)[x][y] = -1
	zeros--
	total := dfs(grid, x+1, y, zeros) + dfs(grid, x-1, y, zeros) + dfs(grid, x, y+1, zeros) + dfs(grid, x, y-1, zeros)
	(*grid)[x][y] = 0
	zeros++
	return total
}