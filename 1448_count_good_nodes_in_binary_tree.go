func goodNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return dfs(root, root.Val)
}

func dfs(root *TreeNode, max_val int) int {
	if root == nil {
		return 0
	}
	if root.Val >= max_val {
		max_val = root.Val
		return 1 + dfs(root.Left, max_val) + dfs(root.Right, max_val)
	}
	return dfs(root.Left, max_val) + dfs(root.Right, max_val)
} 