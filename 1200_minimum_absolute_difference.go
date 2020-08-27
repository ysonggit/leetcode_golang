import "sort"

func minimumAbsDifference(arr []int) [][]int {
	sort.Ints(arr)
	min_diff := 2000000
	for i := 0; i < len(arr)-1; i++ {
		diff := arr[i+1] - arr[i]
		if diff < min_diff {
			min_diff = diff
		}
	}
	res := [][]int{}
	for i := 0; i < len(arr)-1; i++ {
		if arr[i+1]-arr[i] == min_diff {
			res = append(res, []int{arr[i], arr[i+1]})
		}
	}
	return res
}