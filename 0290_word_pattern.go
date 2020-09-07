import "strings"

func wordPattern(pattern string, str string) bool {
	words := strings.Fields(str)
	dict1, dict2 := make(map[rune]string, 0), make(map[string]rune, 0)
	if len(words) != len(pattern) {
		return false
	}
	for i, c := range pattern {
		s := words[i]
		if val, ok := dict1[c]; ok {
			if val != s {
				return false
			}
		} else {
			dict1[c] = s
		}
		if val, ok := dict2[s]; ok {
			if val != c {
				return false
			}
		} else {
			dict2[s] = c
		}
	}
	return true
}