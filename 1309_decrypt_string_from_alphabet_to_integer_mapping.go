import "strconv"

func freqAlphabets(s string) string {
	t := []rune{}
	i, n := 0, len(s)
	for i < n {
		if i < n-2 && s[i+2] == '#' {
			val, _ := strconv.Atoi(s[i : i+2])
			t = append(t, rune(val+96))
			i += 3
		} else {
			val, _ := strconv.Atoi(s[i : i+1])
			t = append(t, rune(val+96))
			i++
		}
	}
	return string(t)
}