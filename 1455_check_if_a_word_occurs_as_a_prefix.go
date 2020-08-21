import "strings"

func isPrefixOfWord(sentence string, searchWord string) int {
	words := strings.Fields(sentence)
	for i := 0; i < len(words); i++ {
		if len(words[i]) >= len(searchWord) && searchWord == words[i][:len(searchWord)] {
			return i + 1
		}
	}
	return -1
}