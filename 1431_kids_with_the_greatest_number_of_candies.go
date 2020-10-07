func kidsWithCandies(candies []int, extraCandies int) []bool {
	max_exist := 0
	for _, i := range candies {
		if i > max_exist {
			max_exist = i
		}
	}
	res := make([]bool, len(candies))
	for i := 0; i < len(candies); i++ {
		gap := max_exist - candies[i]
		if extraCandies >= gap {
			res[i] = true
		}
	}
	return res
}