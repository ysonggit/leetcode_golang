import "strconv"

func simplifiedFractions(n int) []string {
	res := make([]string, 0)
	for i := 2; i <= n; i++ {
		for j := 1; j < i; j++ {
			if gcd(i, j) == 1 {
				res = append(res, strconv.Itoa(j)+"/"+strconv.Itoa(i))
			}
		}
	}
	return res
}

// https://play.golang.org/p/SmzvkDjYlb
// https://en.wikipedia.org/wiki/Euclidean_algorithm
func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}