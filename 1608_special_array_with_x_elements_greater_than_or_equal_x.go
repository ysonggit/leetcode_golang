func specialArray(nums []int) int {
	counts := make(map[int]int)
	for _, i := range nums {
		if i > len(nums) {
			counts[len(nums)]++
		} else {
			counts[i] += 1
		}
	}
	total := 0
	for i := 1000; i >= 0; i-- {
		total += counts[i]
		if total == i {
			return i
		}
	}
	return -1
}