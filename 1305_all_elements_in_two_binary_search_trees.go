func getAllElements(root1 *TreeNode, root2 *TreeNode) []int {
	// get all nodes values from 2 trees into 2 arrays
	arr1, arr2 := make([]int, 0), make([]int, 0)
	inorder(root1, &arr1)
	inorder(root2, &arr2)
	// merge array while sorting linearly
	return merge(arr1, arr2)
}

func inorder(root *TreeNode, vals *[]int) {
	if root == nil {
		return
	}
	inorder(root.Left, vals)
	*vals = append(*vals, root.Val)
	inorder(root.Right, vals)
}

func merge(arr1 []int, arr2 []int) []int {
	arr := make([]int, len(arr1)+len(arr2))
	i, j, k := 0, 0, 0
	for i < len(arr1) && j < len(arr2) {
		if arr1[i] < arr2[j] {
			arr[k] = arr1[i]
			i++
		} else {
			arr[k] = arr2[j]
			j++
		}
		k++
	}
	for i < len(arr1) {
		arr[k] = arr1[i]
		k++
		i++
	}
	for j < len(arr2) {
		arr[k] = arr2[j]
		k++
		j++
	}
	return arr
}