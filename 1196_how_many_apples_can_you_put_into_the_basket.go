import "sort"

func maxNumberOfApples(arr []int) int {
	sort.Ints(arr)
	weights := 0
	for i := 0; i < len(arr); i++ {
		weights += arr[i]
		if weights >= 5000 {
			return i
		}
	}
	return len(arr)
}