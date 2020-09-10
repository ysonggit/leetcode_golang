func isCompleteTree(root *TreeNode) bool {
	queue := make([]*TreeNode, 0)	
	queue = append(queue, root)
	end := false
	for len(queue) > 0 {
		top := queue[0]
		queue = queue[1:]
		if top != nil {
			if end {
				return false
			}
			queue = append(queue, top.Left)
			queue = append(queue, top.Right)
		} else {
			end = true
		}
	}
	return true
}
