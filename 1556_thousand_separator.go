import "strconv"

func thousandSeparator(n int) string {
	num := strconv.Itoa(n)
	res := []rune{}
	for i := 0; i < len(num); i++ {
		if i > 0 && (len(num)-i)%3 == 0 {
			res = append(res, rune('.'))
		}
		res = append(res, rune(num[i]))
	}
	return string(res)
}