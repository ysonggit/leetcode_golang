func maxProduct(nums []int) int {
	max_1, max_2 := 0, 0
	for _, i := range nums {
		if i > max_1 {
			max_2 = max_1
			max_1 = i
		} else {
			if i > max_2 {
				max_2 = i
			}
		}
	}
	return (max_1 - 1) * (max_2 - 1)
}