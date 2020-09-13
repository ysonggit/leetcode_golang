func insert(intervals [][]int, newInterval []int) [][]int {
	res := make([][]int, 0)
	for _, cur := range intervals {
		if cur[0] > newInterval[1] {
			itv := make([]int, 2)
			copy(itv, newInterval)
			res = append(res, itv)
			copy(newInterval, cur)
		} else if cur[1] < newInterval[0] {
			res = append(res, cur)
		} else if cur[1] >= newInterval[0] || cur[0] <= newInterval[1] {
			if cur[0] < newInterval[0] {
				newInterval[0] = cur[0]
			}
			if cur[1] > newInterval[1] {
				newInterval[1] = cur[1]
			}
		}
	}
	res = append(res, newInterval)
	return res
}