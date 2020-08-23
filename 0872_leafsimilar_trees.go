func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	vals1, vals2 := make([]int, 0), make([]int, 0)
	dfs(root1, &vals1)
	dfs(root2, &vals2)
	if len(vals1) != len(vals2) {
		return false
	}
	for i := 0; i < len(vals1); i++ {
		if vals1[i] != vals2[i] {
			return false
		}
	}
	return true
}

func dfs(root *TreeNode, vals *[]int) {
	if root == nil {
		return
	}
	if root.Left == nil && root.Right == nil {
		*vals = append(*vals, root.Val)
	}
	dfs(root.Left, vals)
	dfs(root.Right, vals)
} 