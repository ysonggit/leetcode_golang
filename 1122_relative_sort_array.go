import "sort"

func relativeSortArray(arr1 []int, arr2 []int) []int {
	common := make([]int, 0)
	for _, num := range arr2 {
		for i := len(arr1) - 1; i >= 0; i-- {
			if arr1[i] == num {
				arr1 = append(arr1[:i], arr1[i+1:]...)
				common = append(common, num)
			}
		}
	}
	sort.Ints(arr1)
	return append(common, arr1...)
}