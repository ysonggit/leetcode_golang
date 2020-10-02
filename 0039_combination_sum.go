func combinationSum(candidates []int, target int) [][]int {
	path := []int{}
	result := [][]int{}
	dfs(candidates, target, path, 0, &result)
	return result
}

func dfs(candidates []int, target int, path []int, curIdx int, result *[][]int) {
	if target < 0 {
		return
	}
	if target == 0 {
		cur := make([]int, len(path))
		copy(cur, path)
		*result = append(*result, cur)
		return
	}
	for i := curIdx; i < len(candidates); i++ {
		num := candidates[i]
		path = append(path, num)
		dfs(candidates, target-num, path, i, result)
		path = path[:len(path)-1]
	}
}