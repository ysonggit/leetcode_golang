func reformat(s string) string {
	num_letter, num_digit, n := 0, 0, len(s)
	for i := 0; i < n; i++ {
		if isLetter(rune(s[i])) {
			num_letter++
		} else {
			num_digit++
		}
	}
	if num_letter-num_digit > 1 || num_letter-num_digit < -1 {
		return ""
	}
	if num_letter > num_digit {
		return reformatString(s, 1, 0)
	}
	return reformatString(s, 0, 1)
}
func isLetter(c rune) bool {
	return ('a' <= c && c <= 'z') || ('A' <= c && c <= 'Z')
}
func reformatString(s string, num_idx int, char_idx int) string {
	perm := make([]rune, len(s))
	for _, c := range s {
		if isLetter(c) {
			perm[char_idx] = c
			char_idx += 2
		} else {
			perm[num_idx] = c
			num_idx += 2
		}
	}
	return string(perm)
}