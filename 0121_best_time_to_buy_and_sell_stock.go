func maxProfit(prices []int) int {
	n := len(prices)
	max_pro := 0
	if n == 0 {
		return max_pro
	}
	min_pri := prices[0]
	for i := 1; i < n; i++ {
		if min_pri > prices[i] {
			min_pri = prices[i]
		}
		if max_pro < prices[i]-min_pri {
			max_pro = prices[i] - min_pri
		}
	}
	return max_pro
}