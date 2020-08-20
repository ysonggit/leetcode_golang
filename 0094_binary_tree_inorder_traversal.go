func inorderTraversal(root *TreeNode) []int {
	stack := []*TreeNode{}
	res := []int{}
	if root == nil {
		return res
	}
	cur := root
	for {
		if cur == nil && len(stack) == 0 {
			break
		}
		if cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		} else {
			n := len(stack)
			cur = stack[n-1]
			stack = stack[:n-1]
			res = append(res, cur.Val)
			cur = cur.Right
		}
	}
	return res
}