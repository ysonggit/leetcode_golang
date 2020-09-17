func countCharacters(words []string, chars string) int {
	char_cnts := make([]int, 26)
	for _, c := range chars {
		char_cnts[int(c-97)]++
	}
	len_sum := 0
	for _, cur := range words {
		good_word := true
		word_cnts := make([]int, 26)
		for _, c := range cur {
			word_cnts[int(c-97)]++
		}
		for i := 0; i < 26; i++ {
			if word_cnts[i] > char_cnts[i] {
				good_word = false
				break
			}
		}
		if good_word {
			len_sum += len(cur)
		}
	}
	return len_sum
}