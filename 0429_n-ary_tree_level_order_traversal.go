func levelOrder(root *Node) [][]int {
	queue := make([]*Node, 0)
	vals := make([][]int, 0)
	if root == nil {
		return vals
	}
	queue = append(queue, root)
	for len(queue) > 0 {
		n := len(queue)
		cur_vals := make([]int, 0)
		for i := 0; i < n; i++ {
			cur := queue[i]
			cur_vals = append(cur_vals, cur.Val)
			for _, child := range cur.Children {
				queue = append(queue, child)
			}
		}
		vals = append(vals, cur_vals)
		queue = queue[n:]
	}
	return vals
}