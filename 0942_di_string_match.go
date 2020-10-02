func diStringMatch(S string) []int {
	n := len(S)
	res := []int{}
	l, r := 0, n
	for i := 0; i < n; i++ {
		if S[i] == 'I' {
			res = append(res, l)
			l++
		} else {
			res = append(res, r)
			r--
		}
	}
	if S[n-1] == 'I' {
		res = append(res, l)
	} else {
		res = append(res, r)
	}
	return res
}