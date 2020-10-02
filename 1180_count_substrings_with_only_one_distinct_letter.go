func countLetters(S string) int {
	unique_char_cnt := 1
	res := 0
	l, r := 0, 1
	for r < len(S) {
		if S[l] == S[r] {
			unique_char_cnt++
		} else {
			res += (unique_char_cnt + 1) * unique_char_cnt / 2
			unique_char_cnt = 1
			l = r
		}
		r++
	}
	res += (unique_char_cnt + 1) * unique_char_cnt / 2
	return res
}