func modifyString(s string) string {
	t := []rune(s)
	for i := 0; i < len(t); i++ {
		if t[i] == '?' {
			for c := 'a'; c <= 'c'; c++ {
				if (i == 0 || t[i-1] != c) && (i+1 == len(t) || t[i+1] != c) {
					t[i] = c
					break
				}
			}
		}
	}
	return string(t)
}