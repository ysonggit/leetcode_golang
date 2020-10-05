func bitwiseComplement(N int) int {
	if N == 0 {
		return 1
	}
	loops := N
	mask := 1
	for loops > 0 {
		N = N ^ mask
		mask <<= 1
		loops >>= 1
	}
	return N
}