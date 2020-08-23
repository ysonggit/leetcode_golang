import "unicode"

func makeGood(s string) string {
	stack := []rune{}
	for _, c := range s {
		if len(stack) == 0 {
			stack = append(stack, rune(c))
		} else {
			n := len(stack)
			top := stack[n-1]
			if unicode.IsLower(top) && unicode.IsUpper(c) && unicode.ToLower(c) == rune(top) {
				stack = stack[:n-1]
			} else if unicode.IsUpper(top) && unicode.IsLower(c) && unicode.ToLower(top) == rune(c) {
				stack = stack[:n-1]
			} else {
				stack = append(stack, rune(c))
			}
		}
	}
	return string(stack)
}