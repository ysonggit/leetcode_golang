func numPairsDivisibleBy60(time []int) int {
	m := make(map[int]int)
	res := 0
	for _, cur := range time {
		res += m[(60-cur%60)%60]
		m[cur%60] += 1
	}
	return res
}