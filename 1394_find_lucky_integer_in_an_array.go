func findLucky(arr []int) int {
	counter := make(map[int]int)
	for _, num := range arr {
		counter[num]++
	}
	max_lucky := -1
	for key, val := range counter {
		if key == val {
			if key > max_lucky {
				max_lucky = key
			}
		}
	}
	return max_lucky
}