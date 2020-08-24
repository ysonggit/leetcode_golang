func removeDuplicates(S string) string {
	stack := make([]rune, 0)
	for _, c := range S {
		if len(stack) == 0 {
			stack = append(stack, rune(c))
		} else {
			n := len(stack)
			top := stack[n-1]
			if top == rune(c) {
				stack = stack[:n-1]
			} else {
				stack = append(stack, rune(c))
			}
		}
	}
	return string(stack)
}