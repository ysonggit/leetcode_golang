import "sort"

func maxNumberOfBalloons(text string) int {
	// text consists of lower case English letters only.
	// a ~ z : 97 ~ 122
	dict := [26]int{}
	for _, c := range text {
		i := int(c) - 97
		dict[i]++
	}
	// balloon min(dict[0],dict[1],dict[13]) = x min(dict[11], dict[14]) = 2x
	abn := []int{dict[0], dict[1], dict[13]}
	sort.Ints(abn)
	min_x := abn[0]
	if min_x == 0 {
		return 0
	}
	min_y := dict[11]
	if dict[14] < min_y {
		min_y = dict[14]
	}
	if min_y/min_x >= 2 {
		return min_x
	}
	return min_y / 2
}