func findFrequentTreeSum(root *TreeNode) []int {
	counter := make(map[int]int)
	dfs(root, counter)
	max_cnt := 0
	for _, cnt := range counter {
		if cnt > max_cnt {
			max_cnt = cnt
		}
	}
	freq_sum := []int{}
	for sum, cnt := range counter {
		if cnt == max_cnt {
			freq_sum = append(freq_sum, sum)
		}
	}
	return freq_sum
}

func dfs(root *TreeNode, counter map[int]int) int {
	// assume sum does not exceeds max int32
	sum := 0
	if root != nil {
		sum = root.Val + dfs(root.Left, counter) + dfs(root.Right, counter)
		counter[sum] += 1
	}
	return sum
}