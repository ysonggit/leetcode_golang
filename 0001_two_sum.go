func twoSum(nums []int, target int) []int {
	// https://gobyexample.com/maps
	dict := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		cur := nums[i]
		res := target - cur
		// https://stackoverflow.com/questions/2050391/how-to-check-if-a-map-contains-a-key-in-go
		if _, ok := dict[res]; ok {
			return []int{dict[res], i}
		}
		dict[cur] = i
	}
	return []int{}
}