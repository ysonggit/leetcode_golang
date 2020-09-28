func minOperations(logs []string) int {
	stack := make([]string, 0)
	for _, op := range logs {
		if op == "./" {
			continue
		} else if op == "../" {
			if len(stack) > 0 {
				stack = stack[0 : len(stack)-1]
			}
		} else {
			stack = append(stack, op)
		}
	}
	return len(stack)
}