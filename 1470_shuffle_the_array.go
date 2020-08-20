func shuffle(nums []int, n int) []int {
	x := make([]int, n)
	y := make([]int, n)
	copy(x, nums[0:n])
	copy(y, nums[n:2*n])
	for i := 0; i < 2*n; i++ {
		if i&1 == 0 {
			nums[i] = x[i/2]
		} else {
			nums[i] = y[(i-1)/2]
		}
	}
	return nums
}