func findErrorNums(nums []int) []int {
	dup, n := 0, len(nums)
	sum_tobe := n * (n + 1) / 2
	sum_actual := 0
	for i := 0; i < n; i++ {
		cur_pos := abs(nums[i])
		sum_actual += cur_pos
		if nums[cur_pos-1] < 0 {
			dup = cur_pos
		} else {
			nums[cur_pos-1] *= -1
		}
	}
	miss := sum_tobe - sum_actual + dup
	return []int{dup, miss}
}

func abs(x int) int {
	if x < 0 {
		return -1 * x
	}
	return x
}