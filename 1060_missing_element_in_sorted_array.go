func missingElement(nums []int, k int) int {
	for i := 0; i < len(nums)-1; i++ {
		gap := nums[i+1] - nums[i] - 1
		if gap < k {
			k -= gap
		} else {
			return k + nums[i]
		}
	}
	return nums[len(nums)-1] + k
}