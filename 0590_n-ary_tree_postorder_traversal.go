func postorder(root *Node) []int {
	vals := []int{}
	if root == nil {
		return vals
	}
	stack := make([]*Node, 0)
	stack = append(stack, root)
	for len(stack) > 0 {
		n := len(stack)
		top := stack[n-1]
		stack = stack[:n-1] //pop stack
		vals = append(vals, top.Val)
		for _, c := range top.Children {
			stack = append(stack, c)
		}
	}
	// do an inplace reverse array
	for i, j := 0, len(vals)-1; i < j; i, j = i+1, j-1 {
		vals[i], vals[j] = vals[j], vals[i]
	}
	return vals
}