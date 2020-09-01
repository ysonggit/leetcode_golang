import "strconv"

func findNumbers(nums []int) int {
	cnt := 0
	for _, i := range nums {
		if hasEvenDigits(i) {
			cnt++
		}
	}
	return cnt
}
func hasEvenDigits(num int) bool {
	return len(strconv.Itoa(num))%2 == 0
}