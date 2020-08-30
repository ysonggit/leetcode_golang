func containsPattern(arr []int, m int, k int) bool {
	n := len(arr)
	cnt := 0
	for i := 0; i < n-m; i++ {
		if arr[i] != arr[i+m] {
			cnt = 0
		} else {
			cnt++
		}
		if cnt == (k-1)*m {
			return true
		}
	}
	return false
}