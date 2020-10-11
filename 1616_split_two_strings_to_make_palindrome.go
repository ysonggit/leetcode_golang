func checkPalindromeFormation(a string, b string) bool {
	n := len(a)
	if n <= 1 {
		return true
	}
	start, end := 0, n-1
	for start < end {
		if a[start] == b[end] || a[end] == b[start] || a[start] == a[end] || b[start] == b[end] {
			start++
			end--
		} else {
			return false
		}
	}
	return true
}