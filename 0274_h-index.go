func hIndex(citations []int) int {
	m := make(map[int]int)
	n := len(citations)
	for _, v := range citations {
		if v > n {
			m[n]++
		} else {
			m[v]++
		}
	}
	h_idx := 0
	for i := len(citations); i >= 0; i-- {
		h_idx += m[i]
		if i <= h_idx {
			return i
		}
	}
	return 0
}
