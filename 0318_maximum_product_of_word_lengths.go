func maxProduct(words []string) int {
	if len(words) == 0 {
		return 0
	}
	n, bitmap, max_prod := len(words), make([]int, len(words)), 0
	for i := 0; i < n; i++ {
		w := words[i]
		bitmap[i] = 0
		for j := 0; j < len(w); j++ {
			bitmap[i] |= 1 << int(w[j]-'a')
		}
	}
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			prod := len(words[i]) * len(words[j])
			if (bitmap[i]&bitmap[j]) == 0 && max_prod < prod {
				max_prod = prod
			}
		}
	}
	return max_prod
}