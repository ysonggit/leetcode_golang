func findNearestRightNode(root *TreeNode, u *TreeNode) *TreeNode {
	if root == nil || u == nil {
		return nil
	}
	queue := []*TreeNode{root}
	target_found := false
	for len(queue) > 0 {
		cur_level_size := len(queue)
		for i := 0; i < cur_level_size; i++ {
			cur := queue[i]
			if cur == u {
				target_found = true
			} else {
				if target_found {
					return cur
				}
			}
			if cur.Left != nil {
				queue = append(queue, cur.Left)
			}
			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}
		}
		if target_found {
			return nil
		}
		queue = queue[cur_level_size:]
	}
	return nil
}