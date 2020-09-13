import "sort"

func sortByBits(arr []int) []int {
	sort.Slice(arr, func(i, j int) bool {
		i_cnt, j_cnt := bitsNum(arr[i]), bitsNum(arr[j])
		if i_cnt == j_cnt {
			return arr[i] < arr[j]
		}
		return i_cnt < j_cnt
	})
	return arr
}

func bitsNum(a int) int {
	num_1s := 0
	for a != 0 {
		num_1s += a & 0x1
		a >>= 1
	}
	return num_1s
}