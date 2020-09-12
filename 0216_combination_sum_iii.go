func combinationSum3(k int, n int) [][]int {
	sol := [][]int{}
	path := []int{}
	dfs(&sol, path, 1, k, n)
	return sol
}

func dfs(sol *[][]int, path []int, start int, count int, target int) {
	if target == 0 && count == 0 {
		cur := make([]int, len(path))
		copy(cur, path)
		*sol = append(*sol, cur)
		return
	}
	for i := start; i < 10; i++ {
		path = append(path, i)
		dfs(sol, path, i+1, count-1, target-i)
		path = path[:len(path)-1]
	}
}