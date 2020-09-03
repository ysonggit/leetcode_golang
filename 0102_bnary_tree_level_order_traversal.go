func levelOrder(root *TreeNode) [][]int {
	res, level := make([][]int, 0), []*TreeNode{root}
	if root == nil {
		return res
	}
	for len(level) > 0 {
		nextlevel := make([]*TreeNode, 0)
		curlevel := make([]int, 0)
		for i := 0; i < len(level); i++ {
			cur := level[i]
			curlevel = append(curlevel, cur.Val)
			if cur.Left != nil {
				nextlevel = append(nextlevel, cur.Left)
			}
			if cur.Right != nil {
				nextlevel = append(nextlevel, cur.Right)
			}
		}
		res = append(res, curlevel)
		level = nextlevel
	}
	return res
}