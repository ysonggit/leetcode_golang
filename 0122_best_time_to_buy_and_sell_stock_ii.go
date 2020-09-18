func maxProfit(prices []int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}
	max_pro := 0
	for i := 1; i < n; i++ {
		if prices[i] > prices[i-1] {
			max_pro += prices[i] - prices[i-1]
		}
	}
	return max_pro
}