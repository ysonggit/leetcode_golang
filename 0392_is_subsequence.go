func isSubsequence(s string, t string) bool {
	m, n := len(s), len(t)
	i_s := 0
	for i_t := 0; i_s < m && i_t < n; i_t++ {
		if s[i_s] == t[i_t] {
			i_s++
		}
	}
	return i_s == m
}