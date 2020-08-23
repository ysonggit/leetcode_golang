func numIdenticalPairs(nums []int) int {
	dups := make(map[int]int)
	for _, n := range nums {
		dups[n]++
	}
	sum := 0
	for _, num := range dups {
		sum += (num - 1) * num / 2
	}
	return sum
}