func maxSubArray(nums []int) int {
	cur_sum, max_sum := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		cur_sum = max(cur_sum+nums[i], nums[i])
		max_sum = max(max_sum, cur_sum)
	}
	return max_sum
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}