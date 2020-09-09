import "strconv"

func printTree(root *TreeNode) [][]string {
	rows := height(root)
	grid := make([][]string, rows)
	cols := 1<<rows - 1 // pow(2, rows)
	for i := 0; i < rows; i++ {
		row := make([]string, 0)
		for j := 0; j < cols; j++ {
			row = append(row, "")
		}
		grid[i] = row
	}
	preorder(root, &grid, 0, cols-1, 0)
	return grid
}

func height(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left, right := height(root.Left), height(root.Right)
	if left > right {
		return 1 + left
	}
	return 1 + right
}

func preorder(root *TreeNode, grid *[][]string, low int, high int, r int) {
	if root != nil {
		mid := (high + low) / 2
		(*grid)[r][mid] = strconv.Itoa(root.Val)
		preorder(root.Left, grid, low, mid, r+1)
		preorder(root.Right, grid, mid+1, high, r+1)
	}
}