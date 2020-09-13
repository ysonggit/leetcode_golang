import "sort"

func intervalIntersection(A [][]int, B [][]int) [][]int {
	var res [][]int
	i, j := 0, 0
	for len(A) > i && len(B) > j {
		a, b := A[i], B[j]
		left, right := max(a[0], b[0]), min(a[1], b[1])
		if left <= right {
			res = append(res, []int{left, right})
		}
		if a[1] <= right {
			i++
		}
		if b[1] <= right {
			j++
		}
	}
	return res
}

func min(nums ...int) int {
	sort.Ints(nums)
	return nums[0]
}

func max(nums ...int) int {
	sort.Ints(nums)
	return nums[len(nums)-1]
}