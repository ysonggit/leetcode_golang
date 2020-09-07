import "strings"

func repeatedSubstringPattern(s string) bool {
	n := len(s)
	s2 := (s + s)[1 : 2*n-1]
	return strings.Contains(s2, s)
}