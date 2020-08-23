import "sort"

func canMakeArithmeticProgression(arr []int) bool {
	sort.Ints(arr)
	n := len(arr)
	min_a, max_a := arr[0], arr[n-1]
	if (max_a-min_a)%(n-1) != 0 {
		return false
	}
	delta := (max_a - min_a) / (n - 1)
	for i := 0; i < n-1; i++ {
		if arr[i+1]-arr[i] != delta {
			return false
		}
	}
	return true
}