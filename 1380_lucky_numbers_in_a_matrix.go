func luckyNumbers(matrix [][]int) []int {
	all_max, all_min := make(map[int]bool), make(map[int]bool)
	m, n := len(matrix), len(matrix[0])
	for i := 0; i < m; i++ {
		all_min[minInRow(matrix[i])] = true
	}
	for j := 0; j < n; j++ {
		col := make([]int, m)
		for i := 0; i < m; i++ {
			col[i] = matrix[i][j]
		}
		all_max[maxInCol(col)] = true
	}
	lucky_nums := make([]int, 0)
	for key, _ := range all_max {
		if _, ok := all_min[key]; ok {
			lucky_nums = append(lucky_nums, key)
		}
	}
	return lucky_nums
}

func maxInCol(col []int) int {
	max_num := 0
	for _, cur := range col {
		if cur > max_num {
			max_num = cur
		}
	}
	return max_num
}

func minInRow(row []int) int {
	min_num := 100001
	for _, cur := range row {
		if cur < min_num {
			min_num = cur
		}
	}
	return min_num
}