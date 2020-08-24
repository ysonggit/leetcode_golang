/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumRootToLeaf(root *TreeNode) int {
	return dfs(root, 0)
}
func dfs(root *TreeNode, cur int) int {
	if root == nil {
		return cur
	}
	cur = (cur << 1) | root.Val
	if root.Left == nil && root.Right == nil {
		return cur
	}
	total := 0
	if root.Left != nil {
		total += dfs(root.Left, cur)
	}
	if root.Right != nil {
		total += dfs(root.Right, cur)
	}
	return total
}