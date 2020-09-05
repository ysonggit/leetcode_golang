import "strings"

func maxScore(s string) int {
	right_1 := strings.Count(s, "1")
	max_score := 0
	for i, c := range s {
		if i == len(s)-1 { // "00"
			continue
		}
		if c == '1' {
			right_1--
		} else {
			right_1++
		}
		if max_score < right_1 {
			max_score = right_1
		}
	}
	return max_score
}
