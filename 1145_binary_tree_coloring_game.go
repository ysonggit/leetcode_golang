/**
It is always best case for use if either choose y
a) left child of x
this case works when number of nodes in left subtree of x are > n/2

b) right chilf of x
this case works when number of nodes in right subtree of x are > n/2

c) parent of x
this case works when number in nodes in tree of x < (n+1)/2

any of the three condition is satisfied then return true. Otherwise, return false;
*/

func btreeGameWinningMove(root *TreeNode, n int, x int) bool {
	x_node := findX(root, x)
	left_sub_x_count := nodesCount(x_node.Left)
	if left_sub_x_count > n/2 {
		return true
	}
	right_sub_x_count := nodesCount(x_node.Right)
	if right_sub_x_count > n/2 {
		return true
	}
	sub_x_count := 1 + left_sub_x_count + right_sub_x_count
	if sub_x_count < (n+1)/2 {
		return true
	}
	return false
}

func findX(root *TreeNode, x int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == x {
		return root
	}
	left := findX(root.Left, x)
	if left != nil {
		return left
	}
	return findX(root.Right, x)
}

func nodesCount(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + nodesCount(root.Left) + nodesCount(root.Right)
}