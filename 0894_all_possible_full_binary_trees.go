func allPossibleFBT(N int) []*TreeNode {
	if N%2 == 0 {
		return []*TreeNode{}
	}
	if N == 1 {
		return []*TreeNode{{Val: 0}}
	}
	res := make([]*TreeNode, 0)
	for i := 1; i < N-1; i += 2 { // only odd numbers do
		left := allPossibleFBT(i)
		right := allPossibleFBT(N - i - 1)
		for _, l_node := range left {
			for _, r_node := range right {
				res = append(res, &TreeNode{Val: 0, Left: l_node, Right: r_node})
			}
		}
	}
	return res
}