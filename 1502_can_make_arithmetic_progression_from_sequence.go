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

func canMakeArithmeticProgression_2(arr []int) bool {
	// https://en.wikipedia.org/wiki/Arithmetic_progression
	min_val, max_val, n := 1000000, -1000000, len(arr)
	hashset := make(map[int]bool)
	for _, num := range arr {
		if min_val > num {
			min_val = num
		}
		if max_val < num {
			max_val = num
		}
		hashset[num] = true
	}
	delta := (max_val - min_val) / (n - 1)
	for i, cur := 0, min_val; i < n; i, cur = i+1, cur+delta {
		if _, ok := hashset[cur]; !ok {
			return false
		}
	}
	return true
}