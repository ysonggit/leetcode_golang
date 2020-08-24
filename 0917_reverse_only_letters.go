func reverseOnlyLetters(S string) string {
	s := []rune(S)
	n := len(S)
	if n < 2 {
		return S
	}
	front, back := 0, n-1
	for ok := true; ok; ok = front < back {
		if isLetter(s[front]) && isLetter(s[back]) {
			s[front], s[back] = s[back], s[front]
			front++
			back--
		} else if !isLetter(s[front]) {
			front++
		} else if !isLetter(s[back]) {
			back--
		}
	}
	return string(s)
}

func isLetter(c rune) bool {
	return ('a' <= c && c <= 'z') || ('A' <= c && c <= 'Z')
}