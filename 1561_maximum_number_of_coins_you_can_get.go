import "sort"

func maxCoins(piles []int) int {
	sort.Sort(sort.Reverse(sort.IntSlice(piles)))
	n, sum := len(piles), 0
	i, rounds := 1, n/3
	for rounds > 0 {
		sum += piles[i]
		rounds--
		i += 2
	}
	return sum
}