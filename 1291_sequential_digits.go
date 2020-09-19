import (
	"sort"
	"strconv"
)

func sequentialDigits(low int, high int) []int {
	s := "123456789"
	res := make([]int, 0)
	for i := 0; i < 9; i++ {
		for j := i + 1; j < 9; j++ {
			num, _ := strconv.Atoi(s[i : i+(j-i+1)])
			if num >= low && num <= high {
				res = append(res, num)
			}
		}
	}
	sort.Ints(res)
	return res
}