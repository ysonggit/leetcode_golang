func minStartValue(nums []int) int {
	sum, min_prefix_sum := 0, 0
	for _, i := range nums {
		sum += i
		if sum < min_prefix_sum {
			min_prefix_sum = sum
		}
	}
	return 1 - min_prefix_sum
}