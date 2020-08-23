func average(salary []int) float64 {
	min_s, max_s, sum := 1000000, 1000, 0
	for _, cur := range salary {
		sum += cur
		if cur > max_s {
			max_s = cur
		}
		if cur < min_s {
			min_s = cur
		}
	}
	return float64(sum-max_s-min_s) / float64(len(salary)-2)
}