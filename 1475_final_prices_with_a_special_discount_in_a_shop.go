func finalPrices(prices []int) []int {
	n := len(prices)
	final := prices
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if prices[j] <= prices[i] {
				final[i] = final[i] - prices[j]
				break
			}
		}
	}
	return final
}