func sortArrayByParityII(A []int) []int {
	n := len(A)
	res := make([]int, n)
	i, j := 0, 1
	for _, v := range A {
		if v&1 == 0 {
			res[i] = v
			i += 2
		} else {
			res[j] = v
			j += 2
		}
	}
	return res
}