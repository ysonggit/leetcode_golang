import "strings"

func reorderSpaces(text string) string {
	total_len := len(text)
	char_len := 0
	for _, c := range text {
		if c != ' ' {
			char_len++
		}
	}
	space_len := total_len - char_len
	words := strings.Fields(text)
	if len(words) == 1 {
		return appendSpaces(words[0], space_len)
	}
	spaces := space_len / (len(words) - 1)
	res := words[0]
	for i := 1; i < len(words); i++ {
		res = appendSpaces(res, spaces)
		res += words[i]
	}
	spaces_left := space_len - spaces*(len(words)-1)
	if spaces_left > 0 {
		res = appendSpaces(res, spaces_left)
	}
	return res
}

func appendSpaces(s string, spaces_num int) string {
	for i := 0; i < spaces_num; i++ {
		s += " "
	}
	return s
}