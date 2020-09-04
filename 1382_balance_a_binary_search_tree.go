func balanceBST(root *TreeNode) *TreeNode {
	nodes := make([]int, 0)
	inOrder(root, &nodes)
	//fmt.Println(nodes)
	return constructBST(nodes, 0, len(nodes)-1)
}

func inOrder(root *TreeNode, nodes *[]int) {
	if root != nil {
		inOrder(root.Left, nodes)
		*nodes = append(*nodes, root.Val)
		inOrder(root.Right, nodes)
	}
}

func constructBST(nodes []int, left int, right int) *TreeNode {
	if left > right {
		return nil
	}
	if left == right {
		root := new(TreeNode)
		root.Val = nodes[left]
		return root
	}
	mid := (left + right) / 2
	root := new(TreeNode)
	root.Val = nodes[mid]
	root.Left = constructBST(nodes, left, mid-1)
	root.Right = constructBST(nodes, mid+1, right)
	return root
}