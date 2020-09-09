func addOneRow(root *TreeNode, v int, d int) *TreeNode {
	if root == nil || d <= 0 {
		return nil // must be nil, can NOT be new(TreeNode), which is a empty ptr
	}
	if d == 1 {
		cur := new(TreeNode)
		cur.Val = v
		cur.Left = root
		return cur
	}
	if d == 2 {
		cur := new(TreeNode)
		cur.Val = v
		cur.Left = root.Left
		root.Left = cur
		cur_dup := new(TreeNode)
		cur_dup.Val = v
		cur_dup.Right = root.Right
		root.Right = cur_dup
		return root
	}
	root.Left = addOneRow(root.Left, v, d-1)
	root.Right = addOneRow(root.Right, v, d-1)
	return root
}