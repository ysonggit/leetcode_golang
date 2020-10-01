func replaceElements(arr []int) []int {
	n := len(arr)
	mx := -1
	for i := n - 1; i >= 0; i-- {
		arr[i], mx = mx, max(mx, arr[i])
	}
	return arr
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}