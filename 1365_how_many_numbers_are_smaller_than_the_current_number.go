func smallerNumbersThanCurrent(nums []int) []int {
	n, freq := len(nums), make([]int, 101)
	res := make([]int, n)
	for _, i := range nums {
		freq[i]++
	}
	for i := 1; i < len(freq); i++ {
		freq[i] += freq[i-1]
	}
	for i := 0; i < n; i++ {
		freq[nums[i]]--
	}
	for i := 0; i < n; i++ {
		res[i] = freq[nums[i]]
	}
	return res
}