func kthSmallest(matrix [][]int, k int) int {
	n := len(matrix)
	if n*n < k {
		return -1
	}
	low, high := matrix[0][0], matrix[n-1][n-1]
	for low < high {
		mid := low + (high-low)/2
		cnt := 0
		for i := 0; i < n; i++ {
			col := binarySearch(matrix[i], mid)
			cnt += col + 1
		}
		if cnt < k {
			low = mid + 1
		} else {
			high = mid
		}
	}
	return low
}

func binarySearch(row []int, target int) int {
	if row[0] > target {
		return -1
	}
	l, r := 0, len(row)-1
	for l < r {
		m := l + (r-l+1)/2 // must be (r-l+1)
		if row[m] <= target {
			l = m
		} else {
			r = m - 1
		}
	}
	return l
}