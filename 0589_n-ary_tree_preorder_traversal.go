func preorder(root *Node) []int {
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
		for i := len(top.Children) - 1; i >= 0; i-- {
			stack = append(stack, top.Children[i])
		}
	}
	return vals
}