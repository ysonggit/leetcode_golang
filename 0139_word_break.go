func wordBreak(s string, wordDict []string) bool {
	n := len(s)
	dp := make([]bool, n+1)
	words := make(map[string]bool)
	for _, t := range wordDict {
		words[t] = true
	}
	dp[0] = true
	for i := 1; i <= n; i++ {
		for k := 0; k < i; k++ {
			if _, ok := words[s[k:i]]; ok {
				if dp[k] {
					dp[i] = true
					break
				}
			}
		}
	}
	return dp[n]
}