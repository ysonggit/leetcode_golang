func carPooling(trips [][]int, capacity int) bool {
	cap_at_stops := make([]int, 1000)
	for _, trip := range trips {
		cap_at_stops[trip[1]] += trip[0]
		cap_at_stops[trip[2]] -= trip[0]
	}
	cur := 0
	for i := 0; i < 1000; i++ {
		cur += cap_at_stops[i]
		if cur > capacity {
			return false
		}
	}
	return true
}