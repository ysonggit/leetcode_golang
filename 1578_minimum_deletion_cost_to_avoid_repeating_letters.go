func minCost(s string, cost []int) int {
	min_cost := 0
	for pre, i := 0, 1; i < len(s); i++ {
		if s[pre] != s[i] {
			pre = i
		} else {
			if cost[pre] < cost[i] {
				min_cost += cost[pre]
				pre = i
			} else {
				min_cost += cost[i]
			}
		}
	}
	return min_cost
}