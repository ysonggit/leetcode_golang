func subtractProductAndSum(n int) int {
	prod, sum := 1, 0
	for n > 0 {
		d := n % 10
		prod *= d
		sum += d
		n = n / 10
	}
	return prod - sum
}