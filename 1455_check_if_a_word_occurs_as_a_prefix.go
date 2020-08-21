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

func isPrefixOfWord_2(sentence string, searchWord string) int {
	start := 0
	n := len(searchWord)
	word_idx := 0
	for i := 0; i < len(sentence); i++ {
		if ' ' == sentence[i] || i == len(sentence)-1 {
			word_idx++
			m := i - start + 1
			//fmt.Println(sentence[start:start+m])
			if m >= n && searchWord == sentence[start:start+n] {
				return word_idx
			} else {
				start = i + 1
			}
		}
	}
	return -1
}